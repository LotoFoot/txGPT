package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/LotoFoot/txGPT/internal/ai"
)

func main() {
	stream := flag.Bool("stream", false, "Activer le streaming")
	flag.Parse()

	cfg := ai.DefaultConfig()
	cfg.Stream = *stream

	if flag.NArg() < 1 {
		fmt.Println("Usage : ./txgpt.exe [--stream] \"Votre prompt\"")
		os.Exit(1)
	}
	prompt := flag.Arg(0)

	// Historique vide pour ce exemple ; étendez si needed
	response, err := ai.Ask(prompt, nil, cfg)
	if err != nil {
		fmt.Println("Erreur :", err)
		os.Exit(1)
	}
	fmt.Println("Réponse de l'IA :", response)
}
