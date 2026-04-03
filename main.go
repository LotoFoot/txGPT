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
	"syscall"
	"time"

	Prompt "github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	openai "github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

const (
	defaultTimeout   = 30 * time.Second
	defaultMaxTokens = 1024
)

var (
	bold = color.New(color.Bold)
	blue = color.New(color.FgBlue)
	log  = logrus.New()
)

type Config struct {
	Model        string
	Stream       bool
	SystemPrompt string
}

func DefaultConfig() Config {
	return Config{Model: openai.GPT4oMini, Stream: false, SystemPrompt: ""}
}

func systemPromptForRole(role, lang string) string {
	if strings.EqualFold(lang, "en") {
		return fmt.Sprintf("You are a %s who always responds in concise technical English.", role)
	}
	return fmt.Sprintf("Tu es un %s qui repond toujours en francais technique.", role)
}

func Ask(prompt string, cfg Config) (string, error) {
	if strings.TrimSpace(prompt) == "" {
		return "", fmt.Errorf("le prompt ne peut pas etre vide")
	}
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY non definie")
	}
	client := openai.NewClient(apiKey)
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	req := openai.ChatCompletionRequest{
		Model: cfg.Model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: cfg.SystemPrompt},
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
		MaxTokens: defaultMaxTokens,
		Stream:    cfg.Stream,
	}
	var response strings.Builder
	if cfg.Stream {
		stream, err := client.CreateChatCompletionStream(ctx, req)
		if err != nil {
			return "", fmt.Errorf("erreur stream: %w", err)
		}
		defer stream.Close()
		for {
			resp, err := stream.Recv()
			if err != nil {
				log.WithError(err).Debug("fin du stream")
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
			return "", fmt.Errorf("erreur API: %w", err)
		}
		if len(resp.Choices) == 0 {
			return "", fmt.Errorf("reponse vide de l'API")
		}
		response.WriteString(resp.Choices[0].Message.Content)
	}
	return response.String(), nil
}

func main() {
	model              := flag.String("model", openai.GPT4oMini, "Modele OpenAI")
	stream             := flag.Bool("stream", false, "Activer le streaming")
	role               := flag.String("role", "expert Kali Linux", "Role systeme")
	lang               := flag.String("lang", "fr", "Langue: fr ou en")
	debug              := flag.Bool("debug", false, "Logs debug")
	jsonOutput         := flag.Bool("json", false, "Sortie JSON")
	isInteractive      := flag.Bool("i", false, "Mode interactif")
	isInteractiveShell := flag.Bool("is", false, "Mode shell interactif")
	isShell            := flag.Bool("s", false, "Generer commandes shell")
	autoExec           := flag.Bool("y", false, "Executer automatiquement")
	flag.Parse()
	if *debug {
		log.SetLevel(logrus.DebugLevel)
	}
	if flag.NArg() < 1 && !*isInteractive && !*isInteractiveShell {
		fmt.Fprintln(os.Stderr, "Usage: txgpt [flags] prompt")
		flag.Usage()
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
	cfg.SystemPrompt = systemPromptForRole(*role, *lang)
	log.WithFields(logrus.Fields{"model": cfg.Model, "stream": cfg.Stream, "role": *role, "lang": *lang}).Debug("config chargee")
	if *isInteractive {
		interactiveMode(cfg)
		return
	}
	if *isInteractiveShell || *isShell {
		interactiveShellMode(cfg, prompt, autoExec)
		return
	}
	response, err := Ask(prompt, cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Erreur:", err)
		os.Exit(1)
	}
	if *jsonOutput {
		jsonResp := struct {
			Response string     `json:"response"`
			Data     [][]string `json:"data,omitempty"`
		}{Response: response, Data: extractNmapData(response)}
		jsonBytes, err := json.Marshal(jsonResp)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Erreur JSON:", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonBytes))
	} else if !*stream {
		fmt.Println("Reponse:", response)
	}
}

func interactiveMode(cfg Config) {
	bold.Print("Mode interactif. Tapez exit.\n")
	history := []string{}
	for {
		blue.Println("You")
		input := Prompt.Input("> ", func(d Prompt.Document) []Prompt.Suggest { return []Prompt.Suggest{} }, Prompt.OptionHistory(history))
		input = strings.TrimSpace(input)
		if input == "exit" { return }
		if input == "" { continue }
		history = append(history, input)
		response, err := Ask(input, cfg)
		if err != nil { fmt.Fprintln(os.Stderr, "Erreur:", err); continue }
		fmt.Println("Reponse:", response)
	}
}

func interactiveShellMode(cfg Config, initialPrompt string, autoExec *bool) {
	bold.Print("Mode shell. Tapez exit.\n")
	history := []string{}
	cfg.SystemPrompt += " Genere une commande shell dans <cmd>.</cmd>"
	if initialPrompt != "" { processShellPrompt(initialPrompt, cfg, autoExec) }
	for {
		blue.Println("You")
		input := Prompt.Input("> ", func(d Prompt.Document) []Prompt.Suggest { return []Prompt.Suggest{} }, Prompt.OptionHistory(history))
		input = strings.TrimSpace(input)
		if input == "exit" { return }
		if input == "" { continue }
		history = append(history, input)
		processShellPrompt(input, cfg, autoExec)
	}
}

func processShellPrompt(prompt string, cfg Config, autoExec *bool) {
	response, err := Ask(prompt, cfg)
	if err != nil { fmt.Fprintln(os.Stderr, "Erreur:", err); return }
	re := regexp.MustCompile(`<cmd>([\s\S]*?)<\/cmd>`)
	matches := re.FindStringSubmatch(response)
	if len(matches) > 1 {
		cmd := strings.TrimSpace(matches[1])
		log.WithField("cmd", cmd).Debug("commande extraite")
		if *autoExec {
			if err := exec.Command("sh", "-c", cmd).Run(); err != nil { fmt.Fprintln(os.Stderr, "Erreur exec:", err) }
		} else {
			fmt.Printf("Executer '%s'? (o/n): ", cmd)
			var confirm string
			fmt.Scanln(&confirm)
			if confirm == "o" || confirm == "y" {
				if err := exec.Command("sh", "-c", cmd).Run(); err != nil { fmt.Fprintln(os.Stderr, "Erreur exec:", err) }
			}
		}
	} else {
		fmt.Println("Reponse:", response)
	}
}

func extractNmapData(resp string) [][]string {
	re := regexp.MustCompile(`(\d+)/tcp\s+(open|closed)\s+(\w+)`)
	matches := re.FindAllStringSubmatch(resp, -1)
	if len(matches) == 0 { return nil }
	data := make([][]string, 0, len(matches))
	for _, match := range matches {
		if len(match) == 4 { data = append(data, []string{match[1], match[2], match[3]}) }
	}
	return data
}
