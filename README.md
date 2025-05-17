
---
# 📘 Hind Runtime Documentation

The **Hind Runtime** is the official interpreter for the `.hl` Hinglish Programming Language — a simple, fun scripting language using Hindi-English syntax.

---

## 🔧 Overview

Hind is a fast and lightweight command-line interpreter for Hinglish scripts. It is written in **Go (Golang)** for performance, cross-platform support, and easy deployment.

Supported Environments:
- ✅ Linux
- ✅ macOS
- ✅ Termux (Android)
- ✅ Replit (manual setup)
- ✅ Windows (via WSL or `go build`)

---

## 🚀 Built with Golang

Hind is written in [Go (Golang)](https://golang.org), enabling:

- ⚡ High-performance execution
- 🔁 Simple cross-compilation
- 🛠️ Easy static binary creation (no dependencies)

### 🔨 Build from Source

To build the runtime manually, follow these steps:

#### 1. Install Golang

Install Go from: https://golang.org/dl/

#### 2. Clone the Repo

```bash
git clone https://github.com/pjdeveloper896/Hinglish-language.git
cd Hinglish-language/hind
````

#### 3. Build the Binary

```bash
go build -o hinglish main.go
```

This creates a local `hinglish` executable.

#### 4. (Optional) Install Globally

```bash
chmod +x hinglish
sudo mv hinglish /usr/local/bin
```

Now you can run:

```bash
hinglish myfile.hl
```

---

## 📥 Quick Installation (Recommended)

Use this one-liner for auto-setup:

```bash
curl -fsSL https://pjdeveloper896.github.io/Hind/Install.sh | bash
```

---

## 💻 Manual Setup (Replit, Termux, No `sudo`)

```bash
curl -LO https://pjdeveloper896.github.io/Hind/hinglish
chmod +x hinglish
./hinglish mycode.hl
```

Optional alias:

```bash
alias hinglish="./hinglish"
```

---

## 📄 Example Hinglish Script

```hl
kaam bol() {
  likho "Hello Hinglish!"
}

bol()
```

---

## 📚 Full Documentation

* 🔗 [Live Docs](https://pjdeveloper896.github.io/Hinglish-language/docs/)
* 📥 [PDF Version](https://pjdeveloper896.github.io/Hinglish-language/docs/Hinglish_Documentation.pdf)
* 💾 [Source Code on GitHub](https://github.com/pjdeveloper896/Hinglish-language)

---

## 🧠 Features

* Hindi-style keywords (e.g. `agar`, `warna`, `ghoom`)
* Functions, variables, loops, conditionals
* Async/Await & API fetch support
* Modules: DOM, DateTime, Async
* Easy integration with terminal or web-based UIs

---

## 🤝 Contribute

Want to add something?
Fork → Edit → PR → Done!

GitHub: [pjdeveloper896](https://github.com/pjdeveloper896)



## 🇮🇳 जय हिंद | Code karo desi style mein!


