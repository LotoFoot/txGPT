# txGPT – AI-Powered CLI Assistant for Kali Linux

txGPT is a Go-based command-line tool (CLI) that leverages the OpenAI API to generate scripts, commands and technical explanations.  
It is tuned for Kali Linux and common penetration-testing workflows, but also works on any Linux, macOS or Windows host.

---

## Features
* **English or French output** – default is English; change with `--lang`.
* **Streaming mode** – see long answers appear live.
* **Role presets** – e.g. `--role "kali expert"` for pentest-focused replies.
* **Interactive shell** – run `txgpt` without arguments for a REPL-like loop.
* **Safe execution flag** – optional `--exec` asks before running generated code.
* **Lightweight** – single static binary, no Python stack required.

---

## Prerequisites
* Go ≥ 1.22  
  `sudo apt install golang-go`    (on Kali Linux)
* A valid OpenAI API key – create one at <https://platform.openai.com/account/api-keys>
* Git – to clone this repo.

---

## Installation

1 – Clone
git clone https://github.com/LotoFoot/txGPT.git
cd txGPT

2 – Build
go mod tidy
go build -o txgpt # Linux / macOS

go build -o txgpt.exe # Windows
3 – (Unix) install globally
sudo mv txgpt /usr/local/bin/
sudo chmod +x /usr/local/bin/txgpt

text

### Configure the API key (Unix)

echo 'export OPENAI_API_KEY="sk-proj-YOUR_KEY"' >> ~/.bashrc
source ~/.bashrc

text

On Windows PowerShell:

$env:OPENAI_API_KEY = "sk-proj-YOUR_KEY"

permanent:
text

---

## Quick Start

txgpt "Generate a Bash script that backs up /var/www to /tmp."

text

Streaming example:

txgpt --stream --role "kali expert"
"Write a secure Bash script that scans 192.168.1.0/24 with nmap, detects hosts and open ports, and outputs XML and HTML."

text

Interactive mode:

txgpt
Prompt > Explain the difference between nmap -sS and -sT

text

---

## Examples

* Basic host discovery  
  `txgpt "Give me a one-liner with nmap to list live hosts on 10.0.0.0/24"`

* Create a Python reverse shell template  
  `txgpt --role "red team" "Produce a Python3 reverse shell (no external libs) that connects to 10.10.10.5:9001"`

* Explain a metasploit module  
  `txgpt "Describe how to use exploit/windows/smb/ms17_010_eternalblue step by step."`

---

## Troubleshooting

| Issue                          | Fix                                                                 |
| --------------------------------| --------------------------------------------------------------------|
| **401 Unauthorized**           | Key invalid/expired → regenerate on OpenAI dashboard.               |
| `shopt/complete not found`     | Install bash-completion `sudo apt install bash-completion`.         |
| Flags ignored (`--stream`)     | Re-build after updating `main.go`; ensure you run the new binary.    |

---

## Contributing

1. Fork this repo  
2. `git checkout -b feature/my-feature`  
3. `git commit -m "Add my feature"`  
4. `git push origin feature/my-feature`  
5. Open a Pull Request

Please write comments and documentation in English.

---

## License
MIT License – see **LICENSE** file for details.

---

*Author : Your Name – contact: you@example.com*
Commit and push
bash
git add README.md
git commit -m "Add English README"
git push origin main
