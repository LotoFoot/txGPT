#!/bin/bash
# txgpt_rich.sh  –  Usage : ./txgpt_rich.sh "Votre prompt"
set -e  # stoppe au 1ᵉʳ échec

if [ -z "$1" ]; then
    echo "Usage : $0 \"Votre prompt\""
    exit 1
fi

PROMPT="$1"

# 1) exécute txgpt et écrit la réponse JSON dans un fichier temporaire
TMP_JSON="$(mktemp)"
./txgpt --json "$PROMPT" > "$TMP_JSON"

# 2) vérifie qu’il y a bien quelque chose
if [ ! -s "$TMP_JSON" ]; then
    echo "Erreur : txgpt n’a rien renvoyé."
    cat "$TMP_JSON"   # log éventuel
    rm "$TMP_JSON"
    exit 1
fi

# 3) affiche joliment avec Rich
python3 rich_display.py \
        --prompt "$PROMPT" \
        --output "$(cat "$TMP_JSON")"

rm "$TMP_JSON"
