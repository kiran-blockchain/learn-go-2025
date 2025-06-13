# Go Profiling Guide

## 🔧 Setup

Add this to your `main.go` to enable pprof:
```go
import _ "net/http/pprof"
go http.ListenAndServe("localhost:6060", nil)
```

## 🚀 Run the App

```bash
go run main.go
```

## 📊 Available Profiling Endpoints (localhost:6060/debug/pprof/)

- /debug/pprof/
- /debug/pprof/goroutine
- /debug/pprof/heap
- /debug/pprof/threadcreate
- /debug/pprof/block

## 📤 Generate CPU Profile

```bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

### Options:

- `top` — Show top functions
- `list <func>` — Show source code
- `web` — Open graph in browser
- `svg` — Save as SVG
- `pdf` — Save as PDF

Example:
```bash
(pprof) web
(pprof) svg > output.svg
(pprof) pdf > output.pdf
```

## 🧠 Install Graphviz if needed

```bash
brew install graphviz  # macOS
sudo apt install graphviz  # Ubuntu
```

## 🧹 Clean Up

Stop the program and pprof server after profiling.

---

Happy Profiling! 🎯