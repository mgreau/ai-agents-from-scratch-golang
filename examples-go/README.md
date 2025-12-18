# Go Examples - AI Agents From Scratch

This directory contains Go implementations of the AI agent examples. Each example demonstrates progressively more advanced concepts in building AI agents.

## Prerequisites

1. **Install Go 1.23+**
   ```bash
   go version
   ```

2. **Install GCC** (required for CGO and go-llama.cpp)
   - **macOS**: `xcode-select --install`
   - **Linux**: `sudo apt-get install build-essential`
   - **Windows**: Install MinGW-w64

3. **Download Models**
   Follow instructions in [DOWNLOAD.md](../DOWNLOAD.md) to download GGUF models.

## Quick Start

From the project root:

```bash
# Build all examples
make build

# Run specific examples
make run-intro
make run-translation
make run-coding
make run-agent
make run-react
```

Or run directly:

```bash
cd examples-go/01_intro
go run main.go
```

## Examples Overview

### 01_intro - Basic LLM Interaction
**What you'll learn:**
- Loading and running a local LLM with Go
- Basic prompt/response cycle
- Resource management (defer cleanup)

**Run:**
```bash
make run-intro
# or
cd 01_intro && go run main.go
```

### 03_translation - System Prompts & Specialization
**What you'll learn:**
- Using system prompts to specialize agents
- Configuring LLM parameters (temperature, etc.)
- Creating focused, task-specific agents

**Run:**
```bash
make run-translation
```

### 06_coding - Streaming Responses
**What you'll learn:**
- Real-time streaming responses
- Go channels for streaming data
- Progressive output display

**Run:**
```bash
make run-coding
```

### 07_simple-agent - Function Calling (Tools)
**What you'll learn:**
- Tool/function calling fundamentals
- Defining tools the LLM can use
- Tool registry pattern
- How agents decide when to use tools

**Run:**
```bash
make run-agent
```

### 09_react-agent - Reasoning + Acting
**What you'll learn:**
- ReAct pattern (Reason → Act → Observe)
- Iterative problem solving
- Multi-step tool use
- Self-correction loops

**Run:**
```bash
make run-react
```

## Code Structure

Each example follows this pattern:

```go
package main

import (
    "context"
    "github.com/mgreau/ai-agents-from-scratch-go/pkg/llm"
    "github.com/mgreau/ai-agents-from-scratch-go/pkg/tools"
    // ...
)

func main() {
    // 1. Load model
    llamaLLM, err := llm.NewLlamaCppLLM(config)
    defer llamaLLM.Close()  // Always clean up
    
    // 2. Setup tools/prompts
    // ...
    
    // 3. Execute
    ctx := context.Background()
    response, err := llamaLLM.Invoke(ctx, prompt, nil)
    
    // 4. Handle response
    fmt.Println(response)
}
```

## Key Go Patterns Used

### 1. Context for Cancellation
```go
ctx := context.Background()
// or with timeout:
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
```

### 2. Resource Cleanup with Defer
```go
llamaLLM, err := llm.NewLlamaCppLLM(config)
if err != nil {
    log.Fatal(err)
}
defer llamaLLM.Close()  // Ensures cleanup
```

### 3. Channels for Streaming
```go
stream, err := llamaLLM.Stream(ctx, prompt, nil)
for token := range stream {
    fmt.Print(token)
}
```

### 4. Error Handling
```go
if err != nil {
    log.Fatalf("Failed: %v", err)
}
```

## Customizing Examples

### Change Model

Edit the `modelPath` in any example:

```go
modelPath, err := filepath.Abs("../../models/YOUR-MODEL.gguf")
```

### Adjust Parameters

Modify the `LlamaCppConfig`:

```go
llamaLLM, err := llm.NewLlamaCppLLM(llm.LlamaCppConfig{
    ModelPath:   modelPath,
    ContextSize: 4096,      // Larger context
    Temperature: 0.9,       // More creative
    TopP:        0.95,
    TopK:        50,
    Threads:     8,         // More CPU threads
    SystemPrompt: "...",
})
```

## Troubleshooting

### CGO Errors

If you see CGO-related errors:
```bash
# Ensure GCC is installed
gcc --version

# Set CGO flags if needed
export CGO_ENABLED=1
```

### Model Not Found

```bash
# Verify model exists
ls -lh ../../models/

# Check path in code matches actual filename
```

### Memory Issues

If the model won't load:
- Try a smaller quantization (Q5_K_S instead of Q8_0)
- Reduce context size in config
- Close other applications

### Slow Performance

- Increase `Threads` in config
- Use a smaller model
- Enable GPU acceleration (if available)

## Next Steps

1. **Understand the Code**: Read through each example carefully
2. **Experiment**: Modify parameters and prompts
3. **Build Your Own**: Create custom tools and agents
4. **Read the Packages**: Explore `pkg/` directory for implementations
5. **Check Original Examples**: See JavaScript versions in `examples/` for comparison

## Contributing

Found a bug or want to improve an example? See [CONTRIBUTING.md](../CONTRIBUTING.md).

## Resources

- [Go Documentation](https://go.dev/doc/)
- [go-llama.cpp](https://github.com/go-skynet/go-llama.cpp)
- [GGUF Models on Hugging Face](https://huggingface.co/models?library=gguf)
