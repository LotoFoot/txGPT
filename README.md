text
# txGPT – AI-Powered CLI Assistant for Kali Linux

txGPT is a Go-based command-line tool (CLI) that leverages the OpenAI API to generate scripts, commands and technical explanations. It is tuned for Kali Linux and common penetration-testing workflows, but also works on any Linux, macOS or Windows host.  
![txGPT demo](images/txGPT.png)

## Features
- **English or French output** – default is English; change with `--lang`.
- **Streaming mode** – see long answers appear live.
- **Role presets** – e.g. `--role "kali expert"` for pentest-focused replies.
- **Interactive shell** – run `txgpt` without arguments for a REPL-like loop.
- **Safe execution flag** – optional `--exec` asks before running generated code.
- **JSON output** – use `--json` for structured responses (easy post-processing).
- **Rich terminal display** – optional Python integration with the Rich library for coloured panels and tables (`rich_display.py` + `txgpt_rich.sh`).
- **Dynamic data extraction** – automatically parses responses (e.g. Nmap `port/state/service`) into JSON arrays.
- **Lightweight** – single static binary; Python only needed for Rich display.

## Prerequisites
- Go ≥ 1.22 (`sudo apt install golang-go` on Kali).
- OpenAI API key – create one at <https://platform.openai.com/account/api-keys>.
- Git – to clone this repo.
- (Optional) Python 3 + Rich: `pip install rich` for colourised output.

## Installation
Clone
git clone https://github.com/LotoFoot/txGPT.git
cd txGPT

Build
go mod tidy
go build -o txgpt # Linux / macOS
go build -o txgpt.exe # Windows

text
Unix global install :
sudo mv txgpt /usr/local/bin/
sudo chmod +x /usr/local/bin/txgpt

text

### Configure the API key
Unix :
echo 'export OPENAI_API_KEY="sk-proj-YOUR_KEY"' >> ~/.bashrc
source ~/.bashrc

text
Windows PowerShell :
$env:OPENAI_API_KEY = "sk-proj-YOUR_KEY"

permanent : ajoute la ligne à ton $PROFILE
text

### Optional Rich display
pip install rich
./txgpt_rich.sh "Écris un script Nmap"

text

## Quick Start
txgpt "Generate a Bash script that backs up /var/www to /tmp."

text

## Examples
Host discovery
txgpt "Give me a one-liner with nmap to list live hosts on 10.0.0.0/24"

Python reverse shell
txgpt --role "red team" "Produce a Python3 reverse shell (no external libs) that connects to 10.10.10.5:9001"

Metasploit module
txgpt "Describe how to use exploit/windows/smb/ms17_010_eternalblue step by step."

JSON output with parsed ports
./txgpt --json "Écris un script Nmap avec des exemples de ports ouverts"

→ {"response":"…","data":[["80","open","http"],["443","open","https"]]}
Rich display
./txgpt_rich.sh "Génère un scan Nmap" # coloured panels + tables

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

Please write comments and documentation in English.

## License
MIT – see **LICENSE**.

*Author : Lionel Oto – contact: you@example.com*
Commandes Git pour terminer
powershell
# 1. Ouvre le README, colle le contenu ci-dessus et sauvegarde
code README.md   # ou notepad README.md

# 2. Marque le conflit comme résolu
git add README.md

# 3. Poursuis le rebase
git rebase --continue

# 4. Pousse la branche mise à jour
git push origin main
