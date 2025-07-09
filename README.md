# 🐚 goshell - Coding Challenges Shell

`goshell` is a minimal shell implementation written in Go that mimics core functionality of Unix shells like `bash` or `zsh`, with a focus on education, architecture, and system-level control.

---

## 🚀 Features

✅ Custom Prompt (`goshell>`)  
✅ Built-in Commands: `cd`, `pwd`, `exit`, `history`  
✅ External Commands Execution  
✅ Pipe Support: `cmd1 | cmd2 | cmd3`  
✅ Signal Handling: Ctrl+C won't terminate shell  
✅ Command History with Arrow Key Support  
✅ Persistent `.goshell_history`  
✅ Dockerized & CI/CD via GitHub Actions

---

## 🛠️ Getting Started

### 🔧 Prerequisites

- Go 1.22+
- (Optional) Docker

### 🧪 Run Locally

```bash
make build
./goshell
```

### 🐳 Docker Run

```bash
docker build -t goshell .
docker run -it goshell
```

---

## 🧪 Run Tests

```bash
go test ./...
```

---

## 📁 Project Structure

```
goshell/
│
├── main.go
├── shell/
│   ├── shell.go
│   ├── executor.go
│   ├── builtin.go
│   ├── pipes.go
│   ├── signal.go
│   └── history.go
├── tests/
│   └── shell_test.go
├── Makefile
├── Dockerfile
├── .dockerignore
├── .goshell_history
└── go.mod
```

---

## 📜 License

MIT – use it freely for learning or real-world applications.

---

## 🧠 Author

Built by [Harish Shisode](https://github.com/shisodeharish) | Inspired by Unix shell internals.
