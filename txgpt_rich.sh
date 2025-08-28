#!/bin/bash
# Utilisation : ./txgpt_rich.sh "ton prompt"

# Vérifie si un prompt est fourni
if [ -z "$1" ]; then
    echo "Usage : ./txgpt_rich.sh \"Votre prompt\""
    exit 1
fi

PROMPT="$1"

# Exécute txGPT avec --json et capture l'output
TXGPT_OUTPUT=$(./txgpt --json "$PROMPT")

# Vérifie si txGPT a réussi
if [ $? -ne 0 ]; then
    echo "Erreur lors de l'exécution de txGPT."
    exit 1
fi

# Passe l'output à rich_display.py
python rich_display.py --prompt "$PROMPT" --output "$TXGPT_OUTPUT"
