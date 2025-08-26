package ai

import (
	"context"
	"fmt"
	"io"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

// Config pour personnaliser les requêtes.
type Config struct {
	Model        string  // ex. openai.GPT4oMini
	Temperature  float32 // 0.0-2.0
	MaxTokens    int     // 0 = illimité
	BaseURL      string  // proxy éventuel
	SystemPrompt string  // rôle système
	Stream       bool    // streaming
}

// DefaultConfig avec français par défaut.
func DefaultConfig() Config {
	return Config{
		Model:        openai.GPT4oMini,
		Temperature:  0.9,
		MaxTokens:    0,
		SystemPrompt: "Tu es un assistant IA spécialisé dans le code, la cybersécurité et les outils Kali Linux. Réponds toujours en français de façon concise et technique.",
		Stream:       false,
	}
}

// Ask avec config et historique (multi-tours).
func Ask(prompt string, history []openai.ChatCompletionMessage, cfg Config) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY non défini. Définissez-la via une variable d'environnement")
	}

	clientCfg := openai.DefaultConfig(apiKey)
	if cfg.BaseURL != "" {
		clientCfg.BaseURL = cfg.BaseURL
	}
	client := openai.NewClientWithConfig(clientCfg)

	messages := []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleSystem, Content: cfg.SystemPrompt}}
	messages = append(messages, history...)
	messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: prompt})

	req := openai.ChatCompletionRequest{
		Model:       cfg.Model,
		Messages:    messages,
		Temperature: cfg.Temperature,
		MaxTokens:   cfg.MaxTokens,
	}

	if cfg.Stream {
		return streamAnswer(client, req)
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("réponse vide")
	}
	return resp.Choices[0].Message.Content, nil
}

// streamAnswer pour affichage progressif.
func streamAnswer(client *openai.Client, req openai.ChatCompletionRequest) (string, error) {
	stream, err := client.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		return "", err
	}
	defer stream.Close()

	full := ""
	for {
		part, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if len(part.Choices) > 0 {
			txt := part.Choices[0].Delta.Content
			fmt.Print(txt)
			full += txt
		}
	}
	fmt.Println()
	return full, nil
}
