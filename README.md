# txGPT – AI-Powered CLI Assistant for Kali Linux

txGPT is a Go-based command-line tool (CLI) that leverages the OpenAI API to generate scripts, commands and technical explanations. It is tuned for Kali Linux and common penetration-testing workflows, but also works on any Linux, macOS or Windows host.

![txGPT demo](images/txGPT.png)

## Table of Contents

1. [Features](#features)
2. [Demo GIF](#demo)
3. [Prerequisites](#prerequisites)
4. [Installation](#installation)
   1. [Clone & Build](#clone--build)
   2. [Global Install (Unix)](#global-install-unix)
   3. [API Key Setup](#api-key-setup)
   4. [Optional Rich Display](#optional-rich-display)
5. [Quick Start](#quick-start)
6. [Usage Examples](#usage-examples)
7. [Troubleshooting](#troubleshooting)
8. [Contributing](#contributing)
9. [License](#license)

## Features

| Capability | Notes |
|------------|-------|
| **English or French output** | Default: English · Switch with `--lang fr` |
| **Streaming mode** | Long answers appear live |
| **Role presets** | e.g. `--role "kali expert"` for pentest-focused replies |
| **Interactive shell** | Run `txgpt` without arguments for a REPL-like loop |
| **Safe execution flag** | Optional `--exec` asks before running generated code |
| **JSON output** | Use `--json` for structured responses (easy post-processing) |
| **Rich terminal display** | Optional Python integration with the Rich library for coloured panels and tables (`rich_display.py` + `txgpt_rich.sh`) |
| **Dynamic data extraction** | Automatically parses responses (e.g. Nmap `port/state/service`) into JSON arrays |
| **Lightweight** | Single static binary; Python only needed for Rich display |

## Demo

[Ajoutez un GIF de démo ici si disponible]

<p align="center">
  <img src="images/demo.gif" alt="animated demo" width="650">
</p>

## Prerequisites

| Tool | Version | Kali Linux command |
|------|---------|-------------------|
| Go | **≥ 1.22** | `sudo apt install golang-go` |
| Git | latest | `sudo apt install git` |
| OpenAI API key | active | <https://platform.openai.com/account/api-keys> |
| **Optional** Python 3 | 3.8 + | `sudo apt install python3 python3-pip` <br/>`pip install rich` |

## Installation

### Clone & Build

git clone https://github.com/LotoFoot/txGPT.git
cd txGPT
go mod tidy
go build -o txgpt # Linux / macOS
go build -o txgpt.exe # Windows

text

### Global Install (Unix)

sudo mv txgpt /usr/local/bin/
sudo chmod +x /usr/local/bin/txgpt

text

### API Key Setup

Bash / Zsh:
echo 'export OPENAI_API_KEY="sk-proj-XXXXXXXXXXXX"' >> ~/.bashrc
source ~/.bashrc

text

<details>
<summary>Windows PowerShell</summary>

$Env:OPENAI_API_KEY = "sk-proj-XXXXXXXXXXXX"

text

Permanent :
notepad $PROFILE # ajoutez la ligne ci-dessus dans le profil

text

</details>

### Optional Rich Display

pip install rich # dans un venv ou --user
chmod +x txgpt_rich.sh
./txgpt_rich.sh "Écris un script Nmap"

text

## Quick Start

txgpt "Generate a Bash script that backs up /var/www to /tmp."

text

Add `--json` for structured output or `--exec` to run with confirmation.

## Usage Examples

| Scenario | Command |
|----------|---------|
| Host discovery one-liner | `txgpt "Give me a one-liner with nmap to list live hosts on 10.0.0.0/24"` |
| Python reverse shell | `txgpt --role "red team" "Produce a Python3 reverse shell (no external libs) that connects to 10.10.10.5:9001"` |
| Metasploit walkthrough | `txgpt "Describe how to use exploit/windows/smb/ms17_010_eternalblue step by step."` |
| JSON + port parsing | `./txgpt --json "Écris un script Nmap avec des exemples de ports ouverts"` |
| Rich display | `./txgpt_rich.sh "Génère un scan Nmap"` |

JSON output with parsed ports:
./txgpt --json "Écris un script Nmap avec des exemples de ports ouverts"
→ {"response":"…","data":[["80","open","http"],["443","open","https"]]}

text

## Troubleshooting

| Issue | Fix |
|-------|-----|
| **401 Unauthorized** | Regenerate/replace your API key. |
| `shopt/complete not found` | `sudo apt install bash-completion`. |
| Flags ignored (`--stream`) | Re-build and run the updated binary. |
| **403 on push** | Use a valid PAT or SSH key for GitHub. |
| Rich not displaying | `pip install rich` then use `txgpt_rich.sh`. |

## Contributing

1. Fork this repo
2. `git checkout -b feature/my-feature`
3. `git commit -m "Add my feature"`
4. `git push origin feature/my-feature`
5. Open a Pull Request

Please keep comments and documentation in English.

## License

MIT License – see [LICENSE](LICENSE).

Copyright (c) 2025 Rabzouz

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

**Author** : Rabzouz – contact: you@example.com  
Happy hacking – always with permission!

## About

Interface IA dans le terminal pour prédictions et chat