# txGPT – AI-Powered CLI Assistant for Kali Linux

![txGPT Logo](images/txGPT.png)

*txGPT est un outil CLI écrit en **Go** qui exploite l’API OpenAI pour générer scripts, commandes et explications techniques. Optimisé pour les flux de travail **Kali Linux** (pentest), il fonctionne également sous Linux, macOS et Windows.*

---

## ✨ Features

| Capability | Notes |
|------------|-------|
| English **or** French output | Default : EN · `--lang fr` |
| Streaming mode | Answers appear live |
| Role presets | `--role "kali expert"` … |
| Interactive shell | `txgpt` ⇢ REPL |
| Safe execution | `--exec` asks before running generated code |
| JSON output | `--json` for structured responses |
| Rich display | Python + Rich (`rich_display.py` / `txgpt_rich.sh`) |
| Dynamic data extraction | Parses Nmap `port/state/service` blocks to JSON |
| Lightweight | Static binary; Python only for Rich mode |
| **shell-gpt integration** | Advanced code, shell & QA generation |
| Code improvement pipeline | `sgpt --code … | tgpt "Improve this code"` |
| Shell automation | `sgpt --shell …` with optional sudo |
| BDD test generation | Edge-case tests via sgpt (“QA Tester” role) |
| Image generation hook | Custom `generate_image "<prompt>"` function |
| Piping & chaining | Chain sgpt ↔ tgpt for iterative workflows |

---

## 🖼️ Demo GIF

*(Insérez un GIF ou une vidéo de démo ici, ex: ![Demo](GIF_txGPT.mp4))*

---

## 🧰 Prerequisites

| Tool | Version | Install (Kali) |
|------|---------|----------------|
| **Go** | ≥ 1.22 | `sudo apt install golang-go` |
| **Git** | latest | `sudo apt install git` |
| **OpenAI API key** | active | create on platform.openai.com |
| Python 3 (optional) | ≥ 3.8 | `sudo apt install python3 python3-pip` + `pip install rich` |
| **shell-gpt** (optional) | latest | `python3 -m venv ~/sgpt_env`<br>`source ~/sgpt_env/bin/activate`<br>`pip install shell-gpt` |

---

## 🚀 Installation

### Clone & Build

git clone https://github.com/LotoFoot/txGPT.git
cd txGPT
go mod tidy
go build -o txgpt # Linux/macOS
go build -o txgpt.exe # Windows

text

### Global Install (Unix)

sudo mv txgpt /usr/local/bin/
sudo chmod +x /usr/local/bin/txgpt

text

### API Key Setup

**Bash / Zsh**

echo 'export OPENAI_API_KEY="sk-proj-XXXXXXXXXXXX"' >> ~/.bashrc
source ~/.bashrc

text

**PowerShell**

$Env:OPENAI_API_KEY = "sk-proj-XXXXXXXXXXXX"

text

### Optional Rich Display

pip install rich # in venv or --user
chmod +x txgpt_rich.sh
./txgpt_rich.sh "Écris un script Nmap"

text

---

## ⚡ Quick Start

txgpt "Generate a Bash script that backs up /var/www to /tmp."

text

- Add `--json` for structured output
- Add `--exec` to run the generated code with confirmation

---

## 📖 Usage Examples

| Scenario | One-liner |
|----------|-----------|
| Host discovery | `txgpt "Give me a one-liner with nmap to list live hosts on 10.0.0.0/24"` |
| Python reverse shell | `txgpt --role "red team" "Produce a Python3 reverse shell (no external libs) that connects to 10.10.10.5:9001"` |
| Metasploit walkthrough | `txgpt "Describe how to use exploit/windows/smb/ms17_010_eternalblue step by step."` |
| JSON + port parsing | `txgpt --json "Écris un script Nmap avec des exemples de ports ouverts"` |
| Rich display | `./txgpt_rich.sh "Génère un scan Nmap"` |
| Code generation (sgpt) | `sgpt --code "Écris un script Python pour analyser des logs d'erreurs"` |
| Code improvement | `sgpt --code "Script pour lister fichiers" | tgpt "Améliore ce code"` |
| System update | `sgpt --shell "Mettre à jour Kali sans confirmation en utilisant sudo"` |
| BDD tests | `sgpt "Génère des cas de tests BDD pour « Ajouter un produit au panier » (edge cases)"` |
| Image creation | `generate_image "cat in cyberpunk city"` (custom bash func) |

---

## 🛠️ Troubleshooting

| Issue | Fix |
|-------|-----|
| `401 Unauthorized` | regenerate your API key |
| `shopt/complete not found` | `sudo apt install bash-completion` |
| Flags ignored (`--stream`) | rebuild the binary |
| `403` on push | use a PAT or SSH key |
| Rich display blank | `pip install rich` then use the wrapper |
| sgpt asks for key | `export OPENAI_API_KEY=…` inside venv |
| PEP 668 pip error | always install extra packages inside a virtual-env |

---

## 🤝 Contributing

git fork
git checkout -b feature/my-feature

hack hack
git commit -m "Add my feature"
git push origin feature/my-feature

text

Open a Pull Request – please keep code & comments in **English**.

**Contributors:** LotoFoot, rabzouz, Moonwalkeuse

---

## 📝 License

MIT © 2025 Rabzouz. See `LICENSE` for details.

---

> **Author:** Rabzouz · you@example.com  
> Happy hacking – always with permission!

---

### About

CLI AI interface for predictions & chat · Interface IA dans le terminal pour prédictions et chat