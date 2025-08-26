package main

import (
	"fmt"
	"os"

	"github.com/LotoFoot/txGPT/internal/ai" // Importe le package ai (adaptez si nécessaire)
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Utilisation : ./txgpt.exe \"Votre question ou prompt\"")
		return
	}
	prompt := os.Args[1] // Accès sécurisé au premier argument
	response, err := ai.Ask(prompt)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println("Réponse de l'IA :", response)
}
