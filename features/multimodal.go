package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	openai "github.com/sashabaranov/go-openai"
)

func AnalyzeImage(client *openai.Client, imagePath string, prompt string) string {
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "Erreur de lecture de l'image: " + err.Error()
	}
	base64Image := base64.StdEncoding.EncodeToString(imageData)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4o, // Modèle avec vision
		Messages: []openai.ChatCompletionMessage{
			{Role: "user", Content: prompt},
			{Role: "user", Content: fmt.Sprintf("Analyse cette image: data:image/png;base64,%s", base64Image)},
		},
		MaxTokens: 300,
	}
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "Erreur API: " + err.Error()
	}
	return resp.Choices[0].Message.Content
}

// Dans main.go, ajoutez un flag --image <chemin>
// Exemple : txgpt --image erreur.png "Explique cette erreur Nmap"
