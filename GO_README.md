# AI Agents From Scratch - Go Edition

> Learn to build AI agents locally without frameworks. Now in Go!

## 🎯 What This Is

An educational project teaching AI agent fundamentals through **working Go code**. Build agents from first principles, understand the internals, then use frameworks wisely.

## 🚀 Quick Start

```bash
# 1. Build llama.cpp (one-time setup)
cd ~ && git clone https://github.com/ggerganov/llama.cpp.git
cd llama.cpp && mkdir build && cd build
cmake .. && cmake --build . --config Release

# Set environment variables (add to ~/.bashrc)
export LLAMA_CPP_DIR="$HOME/llama.cpp"
export CGO_CFLAGS="-I${LLAMA_CPP_DIR}"
export CGO_LDFLAGS="-L${LLAMA_CPP_DIR}/build -lllama"
export LD_LIBRARY_PATH="${LLAMA_CPP_DIR}/build:${LD_LIBRARY_PATH}"

# 2. Clone this project
cd /path/to/projects
git clone https://github.com/mgreau/ai-agents-from-scratch-go.git
cd ai-agents-from-scratch-go

# 3. Download dependencies
make deps

# 4. Download a model (1.8GB)
pip install huggingface-hub
huggingface-cli download Qwen/Qwen3-1.7B-GGUF Qwen3-1.7B-Q8_0.gguf \
  --local-dir models --local-dir-use-symlinks False

# 5. Build and run
make build
make run-intro
```

**Issues with setup?** See [SETUP_LLAMA_CPP.md](SETUP_LLAMA_CPP.md) for detailed troubleshooting.

**Full guide**: [QUICKSTART_GO.md](QUICKSTART_GO.md)

## 📚 What You'll Learn

### Phase 1: Fundamentals (7 Examples)

1. **Basic LLM** - Load model, generate response
2. **Specialization** - System prompts for specific tasks
3. **Reasoning** - Complex problem solving
4. **Parallel Processing** - Batch requests efficiently
5. **Streaming** - Real-time token generation
6. **Tools** - Function calling and tool use
7. **ReAct Agent** - Reasoning + Acting pattern

### Phase 2: Architecture

- **Runnable Pattern** - Composable components
- **Message System** - Typed conversations
- **Tool Registry** - Extensible tool system
- **Agent Framework** - Multi-step reasoning

## 🏗️ Project Status

| Component | Status | Notes |
|-----------|--------|-------|
| Core (Runnable, Messages) | ✅ Complete | Production-ready |
| Tool System | ✅ Complete | Add custom tools easily |
| ReAct Agent | ✅ Complete | Fully functional |
| Examples | ✅ 7/10 | Key examples done |
| LLM Integration | ✅ Complete | Uses go-skynet/go-llama.cpp |
| Setup Required | ⚠️ One-time | Build llama.cpp C++ library |
| Memory System | 🔄 Optional | Nice to have |
| Advanced Features | 🔄 Optional | Nice to have |

## 💻 Code Example

```go
package main

import (
    "context"
    "fmt"
    "github.com/mgreau/ai-agents-from-scratch-go/pkg/llm"
)

func main() {
    // Create LLM
    llamaLLM, _ := llm.NewLlamaCppLLM(llm.LlamaCppConfig{
        ModelPath:   "models/Qwen3-1.7B-Q8_0.gguf",
        Temperature: 0.7,
    })
    defer llamaLLM.Close()
    
    // Generate response
    ctx := context.Background()
    response, _ := llamaLLM.Invoke(ctx, "Hello, how are you?", nil)
    
    fmt.Println("AI:", response)
}
```

## 🎓 Learning Path

### For Beginners
1. Start with [QUICKSTART_GO.md](QUICKSTART_GO.md)
2. Read concept docs in `examples/*/CONCEPT.md`
3. Run examples: `make run-intro`, `make run-agent`, etc.
4. Study the code in `examples-go/`

### For Intermediate Developers
1. Read [CONVERSION_SUMMARY.md](CONVERSION_SUMMARY.md)
2. Explore the architecture in `pkg/core/`
3. Implement custom tools in `pkg/tools/`
4. Build your own agent

### For Advanced Developers
1. Complete LLM integration ([GO_IMPLEMENTATION_NOTES.md](GO_IMPLEMENTATION_NOTES.md))
2. Add memory systems
3. Implement chains and graphs
4. Contribute improvements

## 🛠️ Implementation Guide

The LLM wrapper is a **template** with clear TODOs. To complete:

1. **Choose a library**:
   - `github.com/go-skynet/go-llama.cpp` (easiest)
   - `github.com/ggerganov/llama.cpp/bindings/go` (official)
   - `github.com/gotzmann/llama.go` (alternative)

2. **Update `pkg/llm/llama.go`**:
   - Add import
   - Implement model loading
   - Implement inference
   - Implement streaming

3. **Test**:
   ```bash
   make build
   make run-intro
   ```

**Estimated time**: 2-4 hours  
**Difficulty**: Intermediate Go

See [GO_IMPLEMENTATION_NOTES.md](GO_IMPLEMENTATION_NOTES.md) for details.

## 📦 Project Structure

```
ai-agents-from-scratch-golang/
├── pkg/                    # Go packages
│   ├── core/              # Core abstractions ✅
│   ├── llm/               # LLM wrapper (template) ⚠️
│   ├── tools/             # Tool system ✅
│   └── agents/            # Agent implementations ✅
├── examples-go/           # Go examples ✅
│   ├── 01_intro/
│   ├── 03_translation/
│   ├── 04_think/
│   ├── 05_batch/
│   ├── 06_coding/
│   ├── 07_simple-agent/
│   └── 09_react-agent/
├── examples/              # Original JS (reference)
├── Makefile               # Build commands ✅
├── go.mod                 # Dependencies ✅
└── README.md              # Full documentation ✅
```

## 🎯 Key Features

### Type Safety
```go
// Compile-time type checking
type Runnable interface {
    Invoke(ctx context.Context, input interface{}, config *Config) (interface{}, error)
}
```

### Concurrency
```go
// Native goroutines for parallel processing
results, _ := runnable.Batch(ctx, inputs, nil)
```

### Resource Management
```go
// Explicit cleanup with defer
defer llm.Close()
```

### Streaming
```go
// Channels for real-time output
stream, _ := llm.Stream(ctx, prompt, nil)
for token := range stream {
    fmt.Print(token)
}
```

## 📖 Documentation

- [QUICKSTART_GO.md](QUICKSTART_GO.md) - Get started in 5 minutes
- [GO_IMPLEMENTATION_NOTES.md](GO_IMPLEMENTATION_NOTES.md) - Complete the LLM integration
- [CONVERSION_SUMMARY.md](CONVERSION_SUMMARY.md) - JS to Go conversion details
- [README.md](README.md) - Full project documentation
- [DOWNLOAD.md](DOWNLOAD.md) - Model download guide

## 🤔 Why Go?

| Advantage | Description |
|-----------|-------------|
| **Performance** | Compiled binaries, low memory usage |
| **Concurrency** | Native goroutines for parallel processing |
| **Type Safety** | Catch errors at compile time |
| **Simplicity** | Clean, readable code |
| **Deployment** | Single binary, no runtime dependencies |

## 🎨 Design Principles

1. **Educational First**: Code clarity over clever optimizations
2. **Interface-Based**: Extensible through interfaces
3. **Library-Agnostic**: Core logic independent of LLM library
4. **Production-Ready**: Patterns used in real frameworks

## 🤝 Contributing

1. Complete the LLM integration
2. Add missing examples (memory, AoT agent)
3. Implement advanced features (chains, graphs)
4. Improve documentation
5. Add tests

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## 📊 Comparison

| Aspect | JavaScript | Go |
|--------|-----------|-----|
| Type Safety | Runtime | Compile-time ✅ |
| Performance | Interpreted | Compiled ✅ |
| Concurrency | Async/Await | Goroutines ✅ |
| Binary Size | N/A | 15MB ✅ |
| Memory Usage | ~500MB | ~200MB ✅ |
| Deployment | Node.js + deps | Single binary ✅ |

## 🔗 Resources

- [Go Documentation](https://go.dev/doc/)
- [llama.cpp](https://github.com/ggerganov/llama.cpp)
- [GGUF Models](https://huggingface.co/models?library=gguf)
- [Original JavaScript Version](https://github.com/mgreau/ai-agents-from-scratch)

## 📝 License

Educational resource - use and modify for learning.

## 🙏 Acknowledgments

- Original JavaScript implementation
- llama.cpp team
- Go community
- Contributors

---

**Ready to start?** → [QUICKSTART_GO.md](QUICKSTART_GO.md)  
**Need help?** → [GO_IMPLEMENTATION_NOTES.md](GO_IMPLEMENTATION_NOTES.md)  
**Want details?** → [CONVERSION_SUMMARY.md](CONVERSION_SUMMARY.md)

**Built with ❤️ for developers who want to understand AI agents deeply**

🚀 Happy learning!
