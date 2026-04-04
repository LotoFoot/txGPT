# txGPT â AI-Powered CLI Assistant for Kali Linux

![txGPT Logo](images/txGPT.png)

*txGPT est un outil CLI Ã©crit en **Go** qui exploite l'API OpenAI pour gÃ©nÃ©rer scripts, commandes et explications techniques. OptimisÃ© pour les flux de travail **Kali Linux** (pentest), il fonctionne Ã©galement sous Linux, macOS et Windows.*

---

## â¨ Features

| Capability | Notes |
|------------|-------|
| English **or** French output | Default : EN Â· `--lang fr` |
| Streaming mode | Answers appear live |
| Role presets | `--role "kali expert"` â¦ |
| Interactive shell | `txgpt` â¢ REPL |
| Safe execution | `--exec` asks before running generated code |
| JSON output | `--json` for structured responses |
| Rich display | Python + Rich (`rich_display.py` / `txgpt_rich.sh`) |
| Dynamic data extraction | Parses Nmap `port/state/service` blocks to JSON |
| Lightweight | Static binary; Python only for Rich mode |
| **shell-gpt integration** | Advanced code, shell & QA generation |
| Code improvement pipeline | `sgpt --code â¦ \| tgpt "Improve this code"` |
| Shell automation | `sgpt --shell â¦` with optional sudo |
| BDD test generation | Edge-case tests via sgpt ("QA Tester" role) |
| Image generation hook | Custom `generate_image "<prompt>"` function |
| Piping & chaining | Chain sgpt â tgpt for iterative workflows |

---

## ð txGPT Pro â Interface Web

**txgpt-pro.html** est une interface web complÃ¨te qui s'exÃ©cute directement dans le navigateur, sans serveur. Elle connecte l'API OpenAI avec une sidebar de raccourcis orientÃ©s cybersÃ©curitÃ©.

### ð NouveautÃ©s

- **IcÃ´nes Font Awesome 6** sur tous les boutons de la sidebar pour une UI plus claire et professionnelle
- **Section Android** : 8 nouveaux outils de pentesting mobile

### ðï¸ Sections de la sidebar

| Section | Outils |
|---------|--------|
| ð **Recon** | nmap, whois, theHarvester, shodan, whatweb |
| ð **Web** | sqlmap, nikto, gobuster, wfuzz, curl |
| ð **Exploit** | metasploit, searchsploit, hydra, john, hashcat |
| ðµï¸ **Forensics** | volatility, binwalk, strings, wireshark, autopsy |
| ð¡ **RÃ©seau / Pentest** | netcat, tcpdump, aircrack-ng, ettercap |
| ð **OSINT** | maltego, recon-ng, spiderfoot |
| ð **Scripts** | bash, python exploit, reverse shell |
| ð¬ **Quick Prompts** | OWASP Top 10, Incident Response, etc. |
| ð¤ **Android** | adb, apktool, jadx, frida, objection, drozer, apksigner, MobSF |

### ð¤ Outils Android (nouveaux)

| Outil | Description |
|-------|-------------|
| `adb` | Android Debug Bridge â commandes de dÃ©bogage/accÃ¨s device |
| `apktool` | DÃ©compilation, modification et recompilation d'APK |
| `jadx` | DÃ©compilateur Java/Kotlin pour APK |
| `frida` | Instrumentation dynamique et hooking d'apps Android |
| `objection` | Exploration runtime, bypass SSL pinning & root detection |
| `drozer` | Audit de sÃ©curitÃ© complet des applications Android |
| `apksigner` | VÃ©rification et analyse de signature APK |
| `MobSF` | Mobile Security Framework â analyse statique APK |

### â¡ Utilisation

```bash
# Ouvrir directement dans le navigateur
xdg-open txgpt-pro.html
# ou
firefox txgpt-pro.html
```

Ou tÃ©lÃ©charger la derniÃ¨re version :
```bash
wget -O ~/txgpt-pro.html https://raw.githubusercontent.com/LotoFoot/txGPT/main/txgpt-pro.html
```

---

## ð¼ï¸ Demo GIF

![Demo txGPT](GIF_txGPT.mp4)

*(Exemple de dÃ©mo : exÃ©cution de commandes en mode streaming sur Kali Linux.)*

---

## ð§° Prerequisites

| Tool | Version | Install (Kali) | Install (macOS) |
|------|---------|----------------|-----------------|
| **Go** | â¥ 1.22 | `sudo apt install golang-go` | `brew install go` (via Homebrew) |
| **Git** | latest | `sudo apt install git` | `brew install git` (via Homebrew) |
| **OpenAI API key** | active | create on platform.openai.com | create on platform.openai.com |
| Python 3 (optional) | â¥ 3.8 | `sudo apt install python3 python3-pip` + `pip install rich` | `brew install python` + `pip install rich` |
| **shell-gpt** (optional) | latest | `python3 -m venv ~/sgpt_env`<br>`source ~/sgpt_env/bin/activate`<br>`pip install shell-gpt` | MÃªme commande (assurez-vous d'avoir Python installÃ©) |
