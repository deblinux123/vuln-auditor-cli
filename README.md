# Vuln-auditor-cli
# 🛡️ VulnAuditor CLI

An advanced, AI-powered Cyber Security TUI (Terminal User Interface) assistant designed for Pentesters, Bug Bounty Hunters, and DevSecOps engineers. It runs **100% locally and offline** to analyze security logs (like `sqlmap` output) and source code to instantly identify injection vectors, generate Proof of Concepts (PoC), and provide secure, ready-to-paste code patches.

Built with 🖤 using **Go (Golang)**, **Ollama**, and the premium **Charm.sh** terminal ecosystem.

---

## 🇮🇷 راهنمای فارسی (Persian Quick Overview)
**ابزار VulnAuditor CLI** یک دستیار امنیتی پیشرفته متنی در ترمینال است که به صورت کاملاً محلی و ۱۰۰٪ آفلاین اجرا می‌شود. این ابزار لاگ‌های اسکنرها (مثل `sqlmap`) یا کدهای منبع پروژه را دریافت کرده، آسیب‌پذیری‌های خطرناک تزریق (مثل SQL Injection و OS Command Injection) را موشکافی می‌کند، سناریوی اثبات ادعا (PoC) می‌سازد و کدهای کاملاً امن و اصلاح‌شده (Patch) را به شما تحویل می‌دهد تا حریم خصوصی کدهای شما کاملاً حفظ شود.

---

## ✨ Features

*   **🔒 100% Local & Private:** No API keys, no cloud dependencies, and zero data leakage. Your sensitive code blocks and scan logs are processed entirely offline on your local machine using GGUF models.
*   **🎮 Premium Keyboard-Driven TUI:** Beautiful terminal forms (`huh`), responsive progress bars (`bubbles`), and structural logging (`log`) that blend seamlessly into an elegant CLI experience.
*   **🎨 Styled Markdown Reports:** Renders rich text, code syntax highlighting, and colored tables right inside your terminal using `glamour`.
*   **🎯 Target Specific (Injection Focus):** Fine-tuned workflow specifically for critical vulnerabilities like SQL Injection (SQLi) and OS Command Injection.
*   **📂 Multi-Input Architecture:** Supports auditing via direct text pasting, file path loading, or pipeline streaming (`stdin`).

---

## 📦 Tech Stack & Packages

VulnAuditor stands on the shoulders of giants in the Go and TUI space:
*   **CLI Framework / Forms:** [Charm.sh Huh](https://github.com/charmbracelet/huh) - Keyboard-centric interactive forms.
*   **Progress Animation:** [Charm.sh Bubbles/Progress](https://github.com/charmbracelet/bubbles) - Smooth visual tracking for heavy AI tasks.
*   **Markdown Renderer:** [Charm.sh Glamour](https://github.com/charmbracelet/glamour) - High-fidelity terminal markdown rendering.
*   **Structured Logs:** [Charm.sh Log](https://github.com/charmbracelet/log) - Colorful, production-grade logging.
*   **AI Engine Backend:** [Ollama API](https://ollama.com/) - Orchestrates local GGUF model execution (Default port: `11434`).

---

## 🛠️ Prerequisites / پیش‌نیازها

Before launching VulnAuditor, ensure your local environment is ready:

1.  **Install Ollama:** Follow the setup guide for your OS at [ollama.com](https://ollama.com/).
2.  **Pull a Compatible Model:** We highly recommend models optimized for code comprehension and security logic:
```bash
    # Excellent for 8B/7B hardware setups
    ollama pull qwen2.5-coder
    
    # Alternative choice
    ollama pull gemma2
    ```
3.  **Run Ollama Service:** Ensure the backend service is actively running on your local machine (`localhost:11434`).

---

## 🚀 Installation & Usage

### 1. Fast Global Install (Via Go)
```bash
go install [github.com/your_username/vuln-auditor-cli@latest](https://github.com/your_username/vuln-auditor-cli@latest)
