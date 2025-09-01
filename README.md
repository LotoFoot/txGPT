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

---

## 🚀 Installation

txGPT est facile à installer. Suivez les instructions ci-dessous pour votre système d'exploitation. Assurez-vous d'avoir les prérequis installés (voir ci-dessus).

### Installation sur Kali Linux

txGPT est optimisé pour Kali Linux. Ouvrez un terminal et suivez ces étapes :

1. **Clonez le dépôt**  
git clone https://github.com/LotoFoot/txGPT.git
cd txGPT

text
- **Explication** : `git clone` télécharge le code source du dépôt GitHub. `cd txGPT` vous place dans le répertoire du projet pour les étapes suivantes.

2. **Gérez les dépendances Go**  
go mod tidy

text
- **Explication** : Cette commande met à jour et télécharge les modules Go nécessaires. Elle résout automatiquement les dépendances et versions compatibles pour éviter les erreurs de compilation.

3. **Compilez l'exécutable**  
go build -o txgpt

text
- **Explication** : `go build` compile le code source en un binaire exécutable nommé `txgpt`. L'option `-o` spécifie le nom de sortie. Vous pouvez ensuite tester avec `./txgpt`.

4. **Installation globale (optionnelle)**  
sudo mv txgpt /usr/local/bin/
sudo chmod +x /usr/local/bin/txgpt

text
- **Explication** : `sudo mv` déplace le binaire vers un répertoire système accessible globalement. `sudo chmod +x` rend le fichier exécutable par tous. Cela permet de lancer `txgpt` depuis n'importe où sans chemin absolu.

5. **Configurez la clé API OpenAI**  
echo 'export OPENAI_API_KEY="sk-proj-XXXXXXXXXXXX"' >> ~/.bashrc
source ~/.bashrc

text
- **Explication** : Remplacez la clé par la vôtre (de platform.openai.com). `echo` ajoute la variable d'environnement à votre fichier Bash. `source` recharge le fichier pour appliquer les changements immédiatement.

6. **Affichage riche (optionnel)**  
pip install rich # Installez dans un venv si possible
chmod +x txgpt_rich.sh
./txgpt_rich.sh "Votre requête ici"

text
- **Explication** : `pip install rich` ajoute la bibliothèque Python pour un affichage coloré. `chmod +x` rend le script shell exécutable. Utilisez-le pour un rendu visuel amélioré.

### Installation sur macOS

txGPT fonctionne bien sur macOS (via Terminal ou iTerm2). Utilisez Homebrew pour les prérequis si possible. Ouvrez un terminal et suivez ces étapes :

1. **Clonez le dépôt**  
git clone https://github.com/LotoFoot/txGPT.git
cd txGPT

text
- **Explication** : Identique à Kali – `git clone` récupère le code, et `cd` navigue dans le dossier.

2. **Gérez les dépendances Go**  
go mod tidy

text
- **Explication** : Met à jour les modules Go, comme sur Kali, pour assurer la compatibilité des dépendances.

3. **Compilez l'exécutable**  
go build -o txgpt

text
- **Explication** : Compile en binaire pour macOS. Testez avec `./txgpt`. (Note : macOS peut nécessiter des permissions ; utilisez `chmod +x txgpt` si besoin.)

4. **Installation globale (optionnelle)**  
mv txgpt /usr/local/bin/
chmod +x /usr/local/bin/txgpt

text
- **Explication** : `mv` déplace le binaire (pas de `sudo` nécessaire si `/usr/local/bin/` est dans votre PATH). `chmod +x` le rend exécutable. Ajoutez `/usr/local/bin` à votre PATH si ce n'est pas déjà fait (éditez `~/.zshrc` ou `~/.bash_profile`).

5. **Configurez la clé API OpenAI**  
echo 'export OPENAI_API_KEY="sk-proj-XXXXXXXXXXXX"' >> ~/.zshrc # Ou ~/.bash_profile si vous utilisez Bash
source ~/.zshrc

text
- **Explication** : Ajoute la clé à votre fichier de configuration shell (macOS utilise souvent Zsh). `source` applique les changements. Remplacez la clé par la vôtre.

6. **Affichage riche (optionnel)**  
pip install rich # Installez dans un venv si possible
chmod +x txgpt_rich.sh
./txgpt_rich.sh "Votre requête ici"

text
- **Explication** : Similaire à Kali – installe Rich pour un affichage amélioré.

### Vérification (pour les deux OS)
Testez avec :  
txgpt "Génère une commande Nmap pour scanner un réseau"

text
En cas d'erreur, voir **🛠️ Troubleshooting**.

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
| `shopt/complete not found` | `sudo apt install bash-completion` (Kali) ou équivalent sur macOS |
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