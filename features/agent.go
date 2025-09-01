package main

import (
	"context"
	"strings"

	openai "github.com/sashabaranov/go-openai" // Importez votre lib OpenAI
)

func AgentMode(client *openai.Client, prompt string, maxIterations int) string {
	var response string
	currentPrompt := prompt
	for i := 0; i < maxIterations; i++ {
		req := openai.ChatCompletionRequest{
			Model:     openai.GPT4o, // Ou GPT-3.5-turbo pour économie
			Messages:  []openai.ChatCompletionMessage{{Role: "user", Content: currentPrompt}},
			MaxTokens: 500,
		}
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			return "Erreur: " + err.Error()
		}
		response = resp.Choices[0].Message.Content

		// Logique simple de boucle : si la réponse indique "continue", itérez
		if !containsContinue(response) { // Fonction personnalisée pour vérifier si plus d'étapes sont nécessaires
			break
		}
		currentPrompt = "Continue la tâche: " + response // Boucle sur la réponse précédente
	}
	return response
}

func containsContinue(text string) bool {
	// Logique basique : cherchez des mots comme "continue" ou "étape suivante"
	return strings.Contains(text, "continue") || strings.Contains(text, "étape suivante")
}

// Dans main.go, ajoutez un flag --agent
// Exemple d'utilisation : txgpt --agent "Debug ce script Nmap"
