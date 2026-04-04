# txGPT – AI-Powered CLI Assistant for Kali Linux

![txGPT Logo](images/txGPT.png)

*txGPT est un outil CLI écrit en **Go** qui exploite l'API OpenAI pour générer scripts, commandes et explications techniques. Optimisé pour les flux de travail **Kali Linux** (pentest), il fonctionne également sous Linux, macOS et Windows.*

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
| Code improvement pipeline | `sgpt --code … \| tgpt "Improve this code"` |
| Shell automation | `sgpt --shell …` with optional sudo |
| BDD test generation | Edge-case tests via sgpt ("QA Tester" role) |
| Image generation hook | Custom `generate_image "<prompt>"` function |
| Piping & chaining | Chain sgpt ↔ tgpt for iterative workflows |

---

## 🌐 txGPT Pro – Interface Web

**txgpt-pro.html** est une interface web complète qui s'exécute directement dans le navigateur, sans serveur. Elle connecte l'API OpenAI avec une sidebar de raccourcis orientés cybersécurité.

### 🚀 Nouveautés

- **Icônes Font Awesome 6** sur tous les boutons de la sidebar pour une UI plus claire et professionnelle
- **Section Android** : 8 nouveaux outils de pentesting mobile

### 🗂️ Sections de la sidebar

| Section | Outils |
|---------|--------|
| 🔍 **Recon** | nmap, whois, theHarvester, shodan, whatweb |
| 🌐 **Web** | sqlmap, nikto, gobuster, wfuzz, curl |
| 🔐 **Exploit** | metasploit, searchsploit, hydra, john, hashcat |
| 🕵️ **Forensics** | volatility, binwalk, strings, wireshark, autopsy |
| 📡 **Réseau / Pentest** | netcat, tcpdump, aircrack-ng, ettercap |
| 🔎 **OSINT** | maltego, recon-ng, spiderfoot |
| 📜 **Scripts** | bash, python exploit, reverse shell |
| 💬 **Quick Prompts** | OWASP Top 10, Incident Response, etc. |
| 🤖 **Android** | adb, apktool, jadx, frida, objection, drozer, apksigner, MobSF |

### 🤖 Outils Android (nouveaux)

| Outil | Description |
|-------|-------------|
| `adb` | Android Debug Bridge – commandes de débogage/accès device |
| `apktool` | Décompilation, modification et recompilation d'APK |
| `jadx` | Décompilateur Java/Kotlin pour APK |
| `frida` | Instrumentation dynamique et hooking d'apps Android |
| `objection` | Exploration runtime, bypass SSL pinning & root detection |
| `drozer` | Audit de sécurité complet des applications Android |
| `apksigner` | Vérification et analyse de signature APK |
| `MobSF` | Mobile Security Framework – analyse statique APK |

### ⚡ Utilisation

```bash
# Ouvrir directement dans le navigateur
xdg-open txgpt-pro.html
# ou
firefox txgpt-pro.html
```

Ou télécharger la dernière version :
```bash
wget -O ~/txgpt-pro.html https://raw.githubusercontent.com/LotoFoot/txGPT/main/txgpt-pro.html
```

---

## 🖼️ Demo GIF

![Demo txGPT](GIF_txGPT.mp4)

*(Exemple de démo : exécution de commandes en mode streaming sur Kali Linux.)*

---

## 🧰 Prerequisites

| Tool | Version | Install (Kali) | Install (macOS) |
|------|---------|----------------|-----------------|
| **Go** | ≥ 1.22 | `sudo apt install golang-go` | `brew install go` (via Homebrew) |
| **Git** | latest | `sudo apt install git` | `brew install git` (via Homebrew) |
| **OpenAI API key** | active | create on platform.openai.com | create on platform.openai.com |
| Python 3 (optional) | ≥ 3.8 | `sudo apt install python3 python3-pip` + `pip install rich` | `brew install python` + `pip install rich` |
| **shell-gpt** (optional) | latest | `python3 -m venv ~/sgpt_env`<br>`source ~/sgpt_env/bin/activate`<br>`pip install shell-gpt` | Même commande (assurez-vous d'avoir Python installé) |
