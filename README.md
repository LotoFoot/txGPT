# txGPT – AI-Powered CLI Assistant for Kali Linux

txGPT is a Go-based command-line tool (CLI) that leverages the OpenAI API to generate scripts, commands and technical explanations.  
It is tuned for Kali Linux and common penetration-testing workflows, but also works on any Linux, macOS or Windows host.
![txGPT demo](images/txGPT.png)


Bien sûr ! Puisque le push fonctionne maintenant (félicitations pour avoir résolu le 403 !), on va mettre à jour le fichier **README.md** de ton dépôt https://github.com/LotoFoot/txGPT. Je vais te fournir une version mise à jour du README (en anglais, comme l'original, pour cohérence), qui inclut les nouvelles fonctionnalités : output JSON, intégration Rich pour affichage enrichi, extraction dynamique de données (regex pour Nmap), et les scripts `rich_display.py` et `txgpt_rich.sh`.



- Remplace **tout le contenu** de README.md par cette version mise à jour (copie-colle-la). J'ai ajouté des sections pour les nouvelles features, l'installation de Rich, et des exemples.

```markdown
# txGPT – AI-Powered CLI Assistant for Kali Linux

txGPT is a Go-based command-line tool (CLI) that leverages the OpenAI API to generate scripts, commands, and technical explanations. It is tuned for Kali Linux and common penetration-testing workflows, but also works on any Linux, macOS, or Windows host.

## Features

- **English or French output** – default is English; change with `--lang`.
- **Streaming mode** – see long answers appear live.
- **Role presets** – e.g. `--role "kali expert"` for pentest-focused replies.
- **Interactive shell** – run `txgpt` without arguments for a REPL-like loop.
- **Safe execution flag** – optional `--exec` asks before running generated code.
- **JSON output** – use `--json` for structured responses (e.g., for integration with other tools).
- **Rich terminal display** – Optional Python integration with Rich library for colored, tabulated outputs (via `rich_display.py` and `txgpt_rich.sh`).
- **Dynamic data extraction** – Automatically parses responses (e.g., Nmap ports/states/services) into JSON data arrays.
- **Lightweight** – single static binary, no Python stack required (except for optional Rich features).

## Prerequisites

- Go ≥ 1.22 (`sudo apt install golang-go` on Kali Linux).
- A valid OpenAI API key – create one at https://platform.openai.com/account/api-keys.
- Git – to clone this repo.
- (Optional for Rich display) Python 3 and Rich library: `pip install rich`.

## Installation

1. Clone the repo:
   ```
   git clone https://github.com/LotoFoot/txGPT.git
   cd txGPT
   ```

2. Build the binary:
   ```
   go mod tidy
   go build -o txgpt  # Linux / macOS
   go build -o txgpt.exe  # Windows
   ```

3. (Unix) Install globally:
   ```
   sudo mv txgpt /usr/local/bin/
   sudo chmod +x /usr/local/bin/txgpt
   ```

4. Configure the API key (Unix):
   ```
   echo 'export OPENAI_API_KEY="sk-proj-YOUR_KEY"' >> ~/.bashrc
   source ~/.bashrc
   ```

   On Windows PowerShell:
   ```
   $env:OPENAI_API_KEY = "sk-proj-YOUR_KEY"
   # Permanent: Add to your profile script
   ```

5. (Optional) For Rich display:
   - Install Rich: `pip install rich`.
   - Use the provided `txgpt_rich.sh` wrapper for enriched outputs.

## Quick Start

Basic usage:
```
txgpt "Generate a Bash script that backs up /var/www to /tmp."
```

With Rich display (on Unix-like systems):
```
./txgpt_rich.sh "Écris un script Nmap"
```
This pipes JSON output to `rich_display.py` for colored tables and panels.

## Examples

- Basic host discovery:
  ```
  txgpt "Give me a one-liner with nmap to list live hosts on 10.0.0.0/24"
  ```

- Create a Python reverse shell:
  ```
  txgpt --role "red team" "Produce a Python3 reverse shell (no external libs) that connects to 10.10.10.5:9001"
  ```

- Explain a Metasploit module:
  ```
  txgpt "Describe how to use exploit/windows/smb/ms17_010_eternalblue step by step."
  ```

- JSON output with data extraction:
  ```
  ./txgpt --json "Écris un script Nmap avec des exemples de ports ouverts"
  # Outputs JSON like: {"response":"...", "data":[["80","open","http"],["443","open","https"]]}
  ```

- Enriched with Rich:
  ```
  ./txgpt_rich.sh "Génère un scan Nmap"  # Affiche en couleurs avec tableaux
  ```

## Troubleshooting

| Issue | Fix |
|-------|-----|
| **401 Unauthorized** | Key invalid/expired → regenerate on OpenAI dashboard. |
| `shopt/complete not found` | Install bash-completion: `sudo apt install bash-completion`. |
| Flags ignored (`--stream`) | Re-build after updating `main.go`; ensure you run the new binary. |
| **403 on push** | Use a valid PAT or SSH key for authentication. |
| Rich not displaying | Ensure `pip install rich` and run via `txgpt_rich.sh`. |

## Contributing

1. Fork this repo.
2. `git checkout -b feature/my-feature`.
3. `git commit -m "Add my feature"`.
4. `git push origin feature/my-feature`.
5. Open a Pull Request.

Please write comments and documentation in English.

## License

MIT License – see **LICENSE** file for details.

*Author: Lionel Oto – contact: you@example.com*
```
