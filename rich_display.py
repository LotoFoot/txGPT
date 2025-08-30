#!/usr/bin/env python3
from rich.console import Console
from rich.table import Table
from rich.panel import Panel
from rich.markdown import Markdown
from rich.spinner import Spinner
import argparse, json, html

# ──────────────────── Parsing des arguments ────────────────────
parser = argparse.ArgumentParser(description="Affiche joliment la sortie JSON de txGPT")
parser.add_argument("--prompt",  required=True, help="Prompt original")
parser.add_argument("--output",  required=True, help="Sortie JSON ou texte brut de txGPT")
args = parser.parse_args()

console = Console()

# ──────────────────── Décodage JSON sûr ────────────────────
try:
    output_data = json.loads(args.output)
    response_text = output_data.get("response", "Aucune réponse")
    data_rows    = output_data.get("data", [])
except json.JSONDecodeError:
    response_text = args.output or "Aucune réponse"
    data_rows    = []

# Nettoyage des échappements
response_text = html.unescape(response_text.replace("\\n", "\n").replace("\\r", ""))

# ──────────────────── Affichage Rich ────────────────────
console.print(Markdown(f"# Résultats pour : {args.prompt}"), style="bold blue")
console.print(Panel(Markdown(response_text), title="Réponse txGPT", style="green"))

# Tableau si données présentes
if data_rows:
    table = Table(title="Données extraites")
    # Colonnes explicites (Port / État / Service) si cela correspond
    for col, style in zip(("Port", "État", "Service"), ("cyan", "magenta", "yellow")):
        table.add_column(col, style=style)
    for row in data_rows:
        table.add_row(*row[:3])           # Sécurise la longueur
    console.print(table)
else:
    console.print("[italic]Aucune donnée structurée trouvée.[/italic]")

console.print("[bold red]Fin ![/bold red]")
