txGPT – AI-Powered CLI Assistant for Kali Linux
![txGPT Logo](https://github.com/LotoFoot/txGPT/blobeverages the OpenAI API to generate scripts, commands and technical explanations. It is tuned for Kali Linux and common penetration-testing workflows, but also works on any Linux, macOS or Windows host.

Table of Contents
Features

Demo GIF

Prerequisites

Installation

Clone & Build

Global Install (Unix)

API Key Setup

Optional Rich Display

Quick Start

Usage Examples

Troubleshooting

Contributing

License

Features
Capability	Notes
English or French output	Default: English · Switch with --lang fr
Streaming mode	Long answers appear live
Role presets	e.g. --role "kali expert" for pentest-focused replies
Interactive shell	Run txgpt without arguments for a REPL-like loop
Safe execution flag	Optional --exec asks before running generated code
JSON output	Use --json for structured responses (easy post-processing)
Rich terminal display	Optional Python integration with the Rich library for coloured panels and tables (rich_display.py + txgpt_rich.sh)
Dynamic data extraction	Automatically parses responses (e.g. Nmap port/state/service) into JSON arrays
Lightweight	Single static binary; Python only needed for Rich display
Integration with shell-gpt (sgpt)	Extend txGPT with advanced AI features like code generation, shell commands, and QA tests via sgpt (install in a virtual environment for compatibility)
Code generation and improvement	Use sgpt for generating Python/Bash scripts and pipe to txGPT for iterative enhancements (e.g., error handling, comments)
Shell command automation	Generate and execute shell commands with sgpt --shell (e.g., system updates), with optional sudo support
BDD test case generation	Create Behavior-Driven Development (BDD) tests for user stories, including edge cases (e.g., via sgpt with roles like "QA Tester")
Image generation support	Add custom commands for AI image generation (e.g., using DALL-E CLI; see contributing for integration)
Piping and chaining	Pipe outputs between txGPT and sgpt for workflows like code refinement or data analysis
Demo
[Ajoutez un GIF de démo ici si disponible]

![Demo GIF](https://github.com/LotoFoot/txGPT/blobcommand |
|--|--|--|
| Go | ≥ 1.22 | sudo apt install golang-go |
| Git | latest | sudo apt install git |
| OpenAI API key | active | https://platform.openai.com/account/api-keys |
| Optional Python 3 | 3.8 + | sudo apt install python3 python3-pip
pip install rich |
| Optional shell-gpt | latest | Install via virtual environment: python3 -m venv ~/sgpt_env
source ~/sgpt_env/bin/activate
pip install shell-gpt |

Installation
Clone & Build
text
git clone https://github.com/LotoFoot/txGPT.git
cd txGPT
go mod tidy
go build -o txgpt # Linux / macOS
go build -o txgpt.exe # Windows
Global Install (Unix)
text
sudo mv txgpt /usr/local/bin/
sudo chmod +x /usr/local/bin/txgpt
API Key Setup
Bash / Zsh:

text
echo 'export OPENAI_API_KEY="sk-proj-XXXXXXXXXXXX"' >> ~/.bashrc
source ~/.bashrc
Windows PowerShell:

text
$Env:OPENAI_API_KEY = "sk-proj-XXXXXXXXXXXX"
Permanent:

text
notepad $PROFILE # ajoutez la ligne ci-dessus dans le profil
Optional Rich Display
text
pip install rich # dans un venv ou --user
chmod +x txgpt_rich.sh
./txgpt_rich.sh "Écris un script Nmap"
Quick Start
text
txgpt "Generate a Bash script that backs up /var/www to /tmp."
Add --json for structured output or --exec to run with confirmation.

Usage Examples
Scenario	Command
Host discovery one-liner	txgpt "Give me a one-liner with nmap to list live hosts on 10.0.0.0/24"
Python reverse shell	txgpt --role "red team" "Produce a Python3 reverse shell (no external libs) that connects to 10.10.10.5:9001"
Metasploit walkthrough	txgpt "Describe how to use exploit/windows/smb/ms17_010_eternalblue step by step."
JSON + port parsing	./txgpt --json "Écris un script Nmap avec des exemples de ports ouverts"
Rich display	./txgpt_rich.sh "Génère un scan Nmap"
Code generation with sgpt	sgpt --code "Écris un script Python pour analyser des logs d'erreurs"
Code improvement via piping	sgpt --code "Script pour lister fichiers" | tgpt "Améliore ce code"
Shell update automation	sgpt --shell "Mettre à jour Kali sans confirmation en utilisant sudo" (confirmez avec [E] pour exécuter)
BDD test cases	sgpt "Génère des cas de tests BDD pour une user story : Ajouter un produit au panier, en incluant edge cases"
Add image generation	sgpt --chat "txgpt_improve" "Ajoute une nouvelle commande pour générer des images AI dans txGPT" (intégrez via custom function in bash)
JSON output with parsed ports: ./txgpt --json "Écris un script Nmap avec des exemples de ports ouverts" → {"response":"…","data":[["80","open","http"],["443","open","https"]]}

Troubleshooting
Issue	Fix
401 Unauthorized	Regenerate/replace your API key.
shopt/complete not found	sudo apt install bash-completion.
Flags ignored (--stream)	Re-build and run the updated binary.
403 on push	Use a valid PAT or SSH key for GitHub.
Rich not displaying	pip install rich then use txgpt_rich.sh.
sgpt API key prompt	Set export OPENAI_API_KEY=sk-your-key and reactivate venv.
Pip externally-managed error	Use a virtual environment as shown in prerequisites.
Contributing
Fork this repo

git checkout -b feature/my-feature

git commit -m "Add my feature"

git push origin feature/my-feature

Open a Pull Request

Please keep comments and documentation in English.

License
MIT License – see LICENSE.

Copyright (c) 2025 Rabzouz

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

Author : Rabzouz – contact: you@example.com
Happy hacking – always with permission!

About
Interface IA dans le terminal pour prédictions et chat