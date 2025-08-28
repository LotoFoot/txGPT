from rich.console import Console
from rich.table import Table
from rich.panel import Panel
from rich.markdown import Markdown
import argparse
import json
import html  # Pour déséchapper les entités HTML comme \u003c

# Parser pour arguments
parser = argparse.ArgumentParser()
parser.add_argument('--prompt', help='Le prompt original')
parser.add_argument('--output', help='La sortie à afficher (JSON ou texte)')
args = parser.parse_args()

console = Console()

# Essayer de parser l'output comme JSON
try:
    output_data = json.loads(args.output)
    response_text = output_data.get('response', 'Aucune réponse')
    data_rows = output_data.get('data', [])
except:
    response_text = args.output or 'Exemple de données'
    data_rows = []

# Nettoyer les échappements (remplace \\n par \n, déséchappe HTML/Unicode)
response_text = response_text.replace('\\n', '\n').replace('\\r', '')
response_text = html.unescape(response_text)  # Gère \u003c -> <, etc.

# Markdown dynamique pour le header
markdown_text = f"""
# Résultats pour : {args.prompt or 'Test'}
- Prompt original : {args.prompt or 'N/A'}
"""

console.print(Markdown(markdown_text), style="bold blue")

# Panel pour la réponse textuelle, rendue comme Markdown
console.print(Panel(Markdown(response_text), title="Réponse txGPT", style="green", expand=True))

# Tableau si des données sont présentes
if data_rows:
    table = Table(title="Données Extraites")
    # Ajouter des colonnes dynamiquement basées sur la première ligne
    if data_rows and data_rows[0]:
        for i in range(len(data_rows[0])):
            table.add_column(f"Col{i+1}", style="cyan" if i % 2 == 0 else "magenta")
    for row in data_rows:
        table.add_row(*row)
    console.print(table)

console.print("[bold red]Fin ! Prêt pour le prochain prompt ?[/bold red]")

