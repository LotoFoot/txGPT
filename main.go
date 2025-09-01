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
	openai "github.com/sashabaranov/go```enai"
	"github.com/sirupsen/logrus"

	"github.com/LotoFoot/txGPT/features```// Chemin d```port complet```sé sur votre```po GitHub
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

	ctx, cancel := context.WithTimeout```ntext.Background(), 30*time.Second)
	defer cancel()

	req := openai.ChatCompletionRequest```	Model: cfg.Model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem,```ntent: cfg```stemPrompt},
			{Role: openai.ChatMessageRoleUser, Content```rompt},
		},
		MaxTokens: 1024,
		Stream:    cfg.Stream,
	}

	var response strings.Builder
	if cfg.Stream {
		stream, err := client.CreateChatCompletion```eam(ctx, req)
		if err != nil {
			return "", err
		}
		defer stream.Close()
		for {
			resp, err := stream.Recv()
			if err != nil {
				break
			}
			for _, choice := range resp.Choices```				fmt.Print(choice.Delta.Content)
				response.WriteString(choice.Delta.Content)
			}
		}
	} else {
		resp, err := client.CreateChatCompletion```x, req)
		if err != nil {
			return "", err
		}
		response.WriteString(resp.Choices[0].Message.Content)
	}
	return response.String(), nil
}

func main() {
	model := flag.String("model", openai.GPT4oMini, "Modèle OpenAI (ex. gpt-4o-mini)")
	stream := flag.Bool("stream", false, "Activer le streaming```	role := flag.String("role", "expert Kali Linux", "Rôle système```ur l'IA")
	debug := flag.Bool("debug", false, "Activer les logs de```bug")
	jsonOutput := flag.Bool("json", false, "Output en format JSON```	isInteractive := flag.Bool("i", false, "Start normal interactive```de")
	isInteractiveShell := flag.Bool("is", false, "Start shell interactive```de")
	isShell := flag.Bool("s", false, "Generate and Execute shell```mmands")
	shouldExecuteCommand := flag.Bool("y", false, "Instantly execute the shell```mmand")

	// Nouveaux flags pour agent et multimodal```gent := flag.Bool("agent", false, "Activer le mode agent```tonome")
	image := flag.String("image", "", "Chemin vers l'image à```alyser (pour mode multimodal)")

	flag.Parse()

	if flag.NArg() < 1 && !*isInteractive && !*isInteractive```ll {
		fmt.Println("Usage : txgpt [flags] \"Votre prompt\"")
		os.Exit(1)
	}
	prompt := flag.Arg(0)

	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIG```M, syscall```GINT)
	go func() {
		<-terminate
		fmt.Println("\nSortie propre.")
		os.Exit(0)
	}()

	cfg := DefaultConfig()
	cfg.Model = *model
	cfg.Stream = *stream
	cfg.SystemPrompt = fmt.Sprintf("Tu es un %s qui répond toujours en```ançais technique``` *role)

	if *debug {
		log.Info(fmt.Sprintf("DEBUG: Config - Model: %s, Stream: %v, Role: %s", cfg```del, cfg.Stream, *```e))
	}

	// Mode agent autonome
	if *agent {
		client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
		response := features.AgentMode(client, prompt, 5)  // 5 itérations max par défaut
		fmt.Println("Réponse agent:", response)
		return
	}

	// Mode multimodal (analyse image)
	if *image != "" {
		client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
		response := features.AnalyzeImage```ient, *image, prompt)
		fmt.Println("Analyse image:", response)
		return
	}

	if *isInteractive {
		interactiveMode(cfg, debug)
		return
	}
	if *isInteractiveShell || *isShell```		interactiveShellMode(cfg, prompt, shouldExecuteCommand, debug```	return
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
			Data:     extractDataFromResponse```sponse),
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
	bold.Print("Mode interactif démarré. Tapez 'exit```our quitter```")
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

func interactiveShellMode(cfg Config, initialPrompt string, auto```c *bool, debug```ool) {
	bold.Print("Mode shell interactif démarré. Tape```exit' pour```itter.\n")
	history := []string{}
	promptIs := "Génère une commande shell``` wrappe-la dans <cmd>."
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

func processShellPrompt(prompt string, cfg Config, autoExec```ool) {
	response, err := Ask(prompt, nil, cfg)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	commandRegex := regexp.MustCompile```cmd>(.*?)</cmd>`)
	matches := commandRegex.FindString```
