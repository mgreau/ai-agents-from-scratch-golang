# Quick Start Guide - Go Edition

Get up and running with AI agents in Go in 5 minutes.

## Step 1: Install Prerequisites

### Go 1.23+
```bash
# Check if Go is installed
go version

# If not, download from https://go.dev/dl/
```

### GCC and CMake (for CGO and llama.cpp)
```bash
# macOS
xcode-select --install
brew install cmake

# Ubuntu/Debian
sudo apt-get install build-essential cmake

# Fedora/RHEL
sudo dnf install gcc gcc-c++ make cmake
```

## Step 2: Build llama.cpp (One-Time Setup)

The Go bindings require the llama.cpp C++ library:

```bash
# Clone and build llama.cpp
cd ~
git clone https://github.com/ggerganov/llama.cpp.git
cd llama.cpp
mkdir build && cd build
cmake .. && cmake --build . --config Release

# Set environment variables (add to ~/.bashrc or ~/.zshrc)
export LLAMA_CPP_DIR="$HOME/llama.cpp"
export CGO_CFLAGS="-I${LLAMA_CPP_DIR}"
export CGO_LDFLAGS="-L${LLAMA_CPP_DIR}/build -lllama"
export LD_LIBRARY_PATH="${LLAMA_CPP_DIR}/build:${LD_LIBRARY_PATH}"

# Reload shell
source ~/.bashrc  # or ~/.zshrc
```

**Having issues?** See [SETUP_LLAMA_CPP.md](SETUP_LLAMA_CPP.md) for detailed troubleshooting.

## Step 3: Clone and Setup Go Project

```bash
# Clone the repository
git clone https://github.com/mgreau/ai-agents-from-scratch-go.git
cd ai-agents-from-scratch-go

# Download Go dependencies
make deps
```

## Step 4: Download a Model

```bash
# Install huggingface-cli (requires Python/pip)
pip install huggingface-hub

# Download recommended model (1.7B, ~1.8GB)
mkdir -p models
huggingface-cli download Qwen/Qwen3-1.7B-GGUF Qwen3-1.7B-Q8_0.gguf \
  --local-dir models --local-dir-use-symlinks False
```

**Alternative**: Download manually from [Hugging Face](https://huggingface.co/Qwen/Qwen3-1.7B-GGUF) and place in `models/` folder.

## Step 5: Build Examples

```bash
# Build all examples
make build

# This creates binaries in ./bin/
```

## Step 6: Run Your First Agent

```bash
# Run the intro example
make run-intro

# Or run the binary directly
./bin/intro
```

You should see output like:
```
AI: Yes, I'm familiar with node-llama-cpp! It's a Node.js...
```

## Step 7: Try More Examples

```bash
# System prompts for specialized behavior
make run-translation

# Streaming responses
make run-coding

# Agent with tools (function calling)
make run-agent

# ReAct agent (reasoning + acting)
make run-react
```

## What's Next?

### Learn the Concepts
1. Read the examples code in `examples-go/`
2. Understand the Runnable pattern in `pkg/core/`
3. Explore tool definitions in `pkg/tools/`
4. Study the ReAct agent in `pkg/agents/`

### Customize
Edit any example to:
- Change the model
- Adjust temperature/parameters
- Add new tools
- Modify prompts

### Build Your Own
```go
package main

import (
    "context"
    "fmt"
    "log"
    "path/filepath"
    
    "github.com/mgreau/ai-agents-from-scratch-go/pkg/llm"
)

func main() {
    modelPath, _ := filepath.Abs("models/Qwen3-1.7B-Q8_0.gguf")
    
    llm, err := llm.NewLlamaCppLLM(llm.LlamaCppConfig{
        ModelPath:   modelPath,
        ContextSize: 2048,
        Temperature: 0.7,
        Threads:     4,
    })
    if err != nil {
        log.Fatal(err)
    }
    defer llm.Close()
    
    ctx := context.Background()
    response, err := llm.Invoke(ctx, "Hello! Who are you?", nil)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("AI:", response)
}
```

## Troubleshooting

### "Model file not found"
- Verify: `ls -lh models/`
- Check the filename matches in your code

### CGO errors
```bash
# Verify GCC is installed
gcc --version

# Set CGO flag explicitly
export CGO_ENABLED=1
```

### Out of memory
- Use a smaller model (Q5_K_S quantization)
- Reduce context size in config
- Close other applications

### Slow performance
- Increase threads in config
- Use a smaller model
- Check CPU usage

## Community & Support

- **Issues**: [GitHub Issues](https://github.com/mgreau/ai-agents-from-scratch-go/issues)
- **Discussions**: Check the original repo for concepts
- **Contributing**: See [CONTRIBUTING.md](CONTRIBUTING.md)

## Learning Path

1. ✅ **You are here** - Basic setup
2. 📖 **Examples** - Work through all examples in order
3. 🏗️ **Packages** - Study the implementations in `pkg/`
4. 🔨 **Build** - Create your own tools and agents
5. 🚀 **Production** - Learn frameworks like LangChain

Happy building! 🎉
