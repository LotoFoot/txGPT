package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"

	Prompt "github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	openai "github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

var (
	bold        = color.New(color.Bold)
	blue        = color.New(color.FgBlue)
	log         = logrus.New()
	cache       = make(map[string]string)
	cacheMu     sync.Mutex
	programLoop = true
)

type Config struct {
	Model        string
	Stream       bool
	SystemPrompt string
}

func DefaultConfig() Config {
	return Config{
		Model:        openai.GPT4oMini,
		Stream:       false,
		SystemPrompt: "",
	}
}

func Ask(prompt string, _ interface{}, cfg Config) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	if client == nil {
		return "", fmt.Errorf("OPENAI_API_KEY non définie")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req := openai.ChatCompletionRequest{
		Model: cfg.Model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: cfg.SystemPrompt},
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
		MaxTokens: 1024,
		Stream:    cfg.Stream,
	}

	var response strings.Builder
	if cfg.Stream {
		stream, err := client.CreateChatCompletionStream(ctx, req)
		if err != nil {
			return "", err
		}
		defer stream.Close()
		for {
			resp, err := stream.Recv()
			if err != nil {
				break
			}
			for _, choice := range resp.Choices {
				fmt.Print(choice.Delta.Content)
				response.WriteString(choice.Delta.Content)
			}
		}
	} else {
		resp, err := client.CreateChatCompletion(ctx, req)
		if err != nil {
			return "", err
		}
		response.WriteString(resp.Choices[0].Message.Content)
	}
	return response.String(), nil
}

func main() {
	model := flag.String("model", openai.GPT4oMini, "Modèle OpenAI (ex. gpt-4o-mini)")
	stream := flag.Bool("stream", false, "Activer le streaming")
	role := flag.String("role", "expert Kali Linux", "Rôle système pour l'IA")
	debug := flag.Bool("debug", false, "Activer les logs de debug")
	jsonOutput := flag.Bool("json", false, "Output en format JSON")
	isInteractive := flag.Bool("i", false, "Start normal interactive mode")
	isInteractiveShell := flag.Bool("is", false, "Start shell interactive mode")
	isShell := flag.Bool("s", false, "Generate and Execute shell commands")
	shouldExecuteCommand := flag.Bool("y", false, "Instantly execute the shell command")
	flag.Parse()

	if flag.NArg() < 1 && !*isInteractive && !*isInteractiveShell {
		fmt.Println("Usage : txgpt [flags] \"Votre prompt\"")
		os.Exit(1)
	}
	prompt := flag.Arg(0)

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-terminate
		fmt.Println("\nSortie propre.")
		os.Exit(0)
	}()

	cfg := DefaultConfig()
	cfg.Model = *model
	cfg.Stream = *stream
	cfg.SystemPrompt = fmt.Sprintf("Tu es un %s qui répond toujours en français technique.", *role)

	if *debug {
		log.Info(fmt.Sprintf("DEBUG: Config - Model: %s, Stream: %v, Role: %s", cfg.Model, cfg.Stream, *role))
	}

	if *isInteractive {
		interactiveMode(cfg, debug)
		return
	}
	if *isInteractiveShell || *isShell {
		interactiveShellMode(cfg, prompt, shouldExecuteCommand, debug)
		return
	}

	response, err := Ask(prompt, nil, cfg)
	if err != nil {
		fmt.Println("Erreur :", err)
		os.Exit(1)
	}

	if *jsonOutput {
		jsonResp := struct {
			Response string     `json:"response"`
			Data     [][]string `json:"data,omitempty"`
		}{
			Response: response,
			Data:     extractDataFromResponse(response),
		}
		jsonBytes, err := json.Marshal(jsonResp)
		if err != nil {
			fmt.Println("Erreur JSON :", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonBytes))
	} else if !*stream {
		fmt.Println("Réponse :", response)
	}
}

func interactiveMode(cfg Config, debug *bool) {
	bold.Print("Mode interactif démarré. Tapez 'exit' pour quitter.\n")
	history := []string{}
	for {
		blue.Println("╭─ You")
		input := Prompt.Input("╰─> ", func(d Prompt.Document) []Prompt.Suggest {
			return []Prompt.Suggest{}
		}, Prompt.OptionHistory(history))
		input = strings.TrimSpace(input)
		if input == "exit" {
			return
		}
		if len(input) < 1 {
			continue
		}
		history = append(history, input)
		response, err := Ask(input, nil, cfg)
		if err != nil {
			fmt.Println("Erreur :", err)
			continue
		}
		fmt.Println("Réponse :", response)
	}
}

func interactiveShellMode(cfg Config, initialPrompt string, autoExec *bool, debug *bool) {
	bold.Print("Mode shell interactif démarré. Tapez 'exit' pour quitter.\n")
	history := []string{}
	promptIs := "Génère une commande shell et wrappe-la dans <cmd>."
	cfg.SystemPrompt += promptIs
	if initialPrompt != "" {
		processShellPrompt(initialPrompt, cfg, autoExec)
	}
	for {
		blue.Println("╭─ You")
		input := Prompt.Input("╰─> ", func(d Prompt.Document) []Prompt.Suggest {
			return []Prompt.Suggest{}
		}, Prompt.OptionHistory(history))
		input = strings.TrimSpace(input)
		if input == "exit" {
			return
		}
		if len(input) < 1 {
			continue
		}
		history = append(history, input)
		processShellPrompt(input, cfg, autoExec)
	}
}

func processShellPrompt(prompt string, cfg Config, autoExec *bool) {
	response, err := Ask(prompt, nil, cfg)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	commandRegex := regexp.MustCompile(`<cmd>(.*?)</cmd>`)
	matches := commandRegex.FindStringSubmatch(response)
	if len(matches) > 1 {
		cmd := strings.TrimSpace(matches[1])
		if *autoExec {
			exec.Command("sh", "-c", cmd).Run()
		} else {
			fmt.Printf("Exécuter '%s' ? (y/n): ", cmd)
			var confirm string
			fmt.Scanln(&confirm)
			if confirm == "y" {
				exec.Command("sh", "-c", cmd).Run()
			}
		}
	} else {
		fmt.Println("Réponse :", response)
	}
}

func extractDataFromResponse(resp string) [][]string {
	var data [][]string
	re := regexp.MustCompile(`(\d+)/tcp\s+(open|closed)\s+(\w+)`)
	matches := re.FindAllStringSubmatch(resp, -1)
	for _, match := range matches {
		if len(match) == 4 {
			data = append(data, []string{match[1], match[2], match[3]})
		}
	}
	if len(data) == 0 {
		data = [][]string{{"Exemple", "Valeur1"}, {"Autre", "Valeur2"}}
	}
	return data
}
