
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
=======

# txGPT – AI-Powered CLI Assistant for Kali Linux

txGPT is a **Go-based command-line assistant** that leverages the OpenAI API to generate scripts, commands and technical explanations.  
Originally tuned for **Kali Linux** and common penetration-testing workflows, it also works on any Linux, macOS or Windows host.

<p align="center">
  <img src="images/txGPT.png" alt="txGPT demo" width="650">
</p>

---

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

---

## Features<a id="features"></a>

| Capability | Notes |
|------------|-------|
| **English or French output** | Default : English · Switch with `--lang fr` |
| **Streaming mode** | Long answers appear live |
| **Role presets** | e.g. `--role "kali expert"` for pentest-focused replies |
| **Interactive REPL** | Run `txgpt` with no args |
| **Safe execution** | `--exec` asks confirmation before running code |
| **JSON output** | `--json` returns structured responses for post-processing |
| **Rich display** | Python + Rich ➜ coloured panels & tables (`txgpt_rich.sh`) |
| **Dynamic data extraction** | Parses outputs (Nmap port/state/service) into JSON |
| **Lightweight** | Single static Go binary; Python only for Rich display |

---

## Demo<a id="demo"></a>

<p align="center">
  <img src="images/demo.gif" alt="animated demo" width="650">
</p>

---

## Prerequisites<a id="prerequisites"></a>

| Tool | Version | Kali Linux command |
|------|---------|-------------------|
| Go | **≥ 1.22** | `sudo apt install golang-go` |
| Git | latest | `sudo apt install git` |
| OpenAI API key | active | <https://platform.openai.com/account/api-keys> |
| **Optional** Python 3 | 3.8 + | `sudo apt install python3 python3-pip` <br/>`pip install rich` |

---

## Installation<a id="installation"></a>

### Clone & Build<a id="clone--build"></a>

git clone https://github.com/LotoFoot/txGPT.git
cd txGPT
go mod tidy
go build -o txgpt # Linux / macOS
go build -o txgpt.exe # Windows

text

### Global Install (Unix)<a id="global-install-unix"></a>

sudo mv txgpt /usr/local/bin/
sudo chmod +x /usr/local/bin/txgpt

text

### API Key Setup<a id="api-key-setup"></a>

Bash / Zsh
echo 'export OPENAI_API_KEY="sk-proj-XXXXXXXXXXXX"' >> ~/.bashrc
source ~/.bashrc

text

<details>
<summary>Windows PowerShell</summary>

$Env:OPENAI_API_KEY = "sk-proj-XXXXXXXXXXXX"

Permanent :
notepad $PROFILE # ajoutez la ligne ci-dessus dans le profil

text
</details>

### Optional Rich Display<a id="optional-rich-display"></a>

pip install rich # dans un venv ou --user
chmod +x txgpt_rich.sh
./txgpt_rich.sh "Écris un script Nmap"

text

---

## Quick Start<a id="quick-start"></a>

txgpt "Generate a Bash script that backs up /var/www to /tmp."

text

Add `--json` for structured output or `--exec` to run the generated code with confirmation.

---

## Usage Examples<a id="usage-examples"></a>

| Scenario | Command |
|----------|---------|
| Host discovery one-liner | `txgpt "Give me a one-liner with nmap to list live hosts on 10.0.0.0/24"` |
| Python reverse shell | `txgpt --role "red team" "Produce a Python3 reverse shell (no external libs) that connects to 10.10.10.5:9001"` |
| Metasploit walkthrough | `txgpt "Describe how to use exploit/windows/smb/ms17_010_eternalblue step by step."` |
| JSON + port parsing | `./txgpt --json "Écris un script Nmap avec des exemples de ports ouverts"` |
| Rich display | `./txgpt_rich.sh "Génère un scan Nmap"` |

---

## Troubleshooting<a id="troubleshooting"></a>

| Issue | Fix |
|-------|-----|
| **401 Unauthorized** | Regenerate your OpenAI key and export it again |
| `shopt/complete not found` | `sudo apt install bash-completion` |
| Flags ignored (`--stream`) | Re-build (`go build`) and run the new binary |
| **403 on git push** | Use a valid PAT or SSH key for GitHub |
| No Rich colours | `pip install rich` then use `txgpt_rich.sh` |

---

## Contributing<a id="contributing"></a>
>>>>>>> 0914bfd7f9d7cea24aaa694511fe3c4629cd4f89

1. Fork the repo  
2. `git checkout -b feature/my-feature`  
3. `git commit -m "Add my feature"`  
4. `git push origin feature/my-feature`  
5. Open a Pull Request

<<<<<<< HEAD
## License
MIT – see **LICENSE**.

*Author : rabzouz– contact: you@example.com*
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
=======
Please keep comments and documentation in English.

---

## License<a id="license"></a>

**MIT** – see [`LICENSE`](LICENSE).

---

> **Author** : rabzouz · you@example.com  
> *Happy hacking – always with permission!*

