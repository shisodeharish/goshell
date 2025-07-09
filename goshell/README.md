# 🧠 goshell – Staff-Level Go Shell

`goshell` (Coding Challenges Shell) is a custom UNIX-like shell built in **Go**, designed with a staff engineer's mindset: modular, testable, container-ready, and CI/CD integrated.

---

## 📦 Features

| Capability              | Status |
|-------------------------|--------|
| Custom Shell Prompt     | ✅     |
| Built-In Commands       | ✅ `cd`, `pwd`, `exit`, `history` |
| External Commands       | ✅ Full support |
| Pipe Support            | ✅ Multi-command `|` chaining |
| Signal Handling         | ✅ Handles `CTRL+C` |
| Command History         | ✅ Persistent & interactive |
| CI/CD                   | ✅ GitHub Actions |
| Dockerized              | ✅ Production-ready Dockerfile |
| Unit Tests              | ✅ Sample included |

---

## 🚀 Quick Start

```bash
make build
./goshell
```

### Docker

```bash
docker build -t goshell .
docker run -it goshell
```

---

## 📁 Directory Layout

```
goshell/
├── shell/         # Core modules
├── utils/         # Input parsing utilities
├── tests/         # Unit tests
├── Dockerfile     # Docker support
├── Makefile       # Build automation
├── .github/       # CI workflows
├── LICENSE
├── README.md
└── go.mod
```

---

## ✅ Staff Engineering Practices Used

- Clean architecture separation
- Reusable utility modules
- Signal-safe I/O handling
- History with arrow-key UX via `liner`
- GitHub Actions CI
- Docker for repeatable builds
- Documentation-first approach

---

## 🧠 Built by

**Harish Shisode**  
[GitHub](https://github.com/shisodeharish) • [Website](https://learningdiary.me/)
