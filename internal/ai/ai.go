package ai

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

// defaultTimeout est le delai maximal d'une requete API.
const defaultTimeout = 30 * time.Second

// Config regroupe les parametres de personnalisation des requetes IA.
type Config struct {
	Model        string  // ex. openai.GPT4oMini
	Temperature  float32 // 0.0-2.0
	MaxTokens    int     // 0 = pas de limite explicite
	BaseURL      string  // proxy optionnel
	SystemPrompt string  // role systeme
	Stream       bool    // activer le streaming
}

// DefaultConfig retourne une configuration par defaut.
func DefaultConfig() Config {
	return Config{
		Model:        openai.GPT4oMini,
		Temperature:  0.9,
		MaxTokens:    0,
		SystemPrompt: "Tu es un assistant IA specialise dans le code, la cybersecurite et les outils Kali Linux. Reponds toujours en francais de facon concise et technique.",
		Stream:       false,
	}
}

// Ask envoie un prompt avec un historique de conversation a l'API OpenAI.
// Supporte le mode streaming et le mode standard (multi-tours).
func Ask(prompt string, history []openai.ChatCompletionMessage, cfg Config) (string, error) {
	if strings.TrimSpace(prompt) == "" {
		return "", fmt.Errorf("le prompt ne peut pas etre vide")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY non definie — definissez-la via une variable d'environnement")
	}

	clientCfg := openai.DefaultConfig(apiKey)
	if cfg.BaseURL != "" {
		clientCfg.BaseURL = cfg.BaseURL
	}
	client := openai.NewClientWithConfig(clientCfg)

	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: cfg.SystemPrompt},
	}
	messages = append(messages, history...)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})

	req := openai.ChatCompletionRequest{
		Model:       cfg.Model,
		Messages:    messages,
		Temperature: cfg.Temperature,
		MaxTokens:   cfg.MaxTokens,
	}

	if cfg.Stream {
		return streamAnswer(client, req)
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("erreur appel API: %w", err)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("reponse vide recue de l'API")
	}
	return resp.Choices[0].Message.Content, nil
}

// streamAnswer gere la reception progressive d'une reponse en streaming.
// Utilise strings.Builder pour eviter les concatenations couteuses.
func streamAnswer(client *openai.Client, req openai.ChatCompletionRequest) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return "", fmt.Errorf("erreur creation du stream: %w", err)
	}
	defer stream.Close()

	var full strings.Builder
	for {
		part, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("erreur lecture du stream: %w", err)
		}
		if len(part.Choices) > 0 {
			txt := part.Choices[0].Delta.Content
			fmt.Print(txt)
			full.WriteString(txt)
		}
	}
	fmt.Println()
	return full.String(), nil
}
