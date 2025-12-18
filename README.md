# AI Agents From Scratch - Go Edition

Learn to build AI agents locally without frameworks. Understand what happens under the hood before using production frameworks.

## Purpose

This repository teaches you to build AI agents from first principles using **local LLMs** and **Go**. By working through these examples, you'll understand:

- How LLMs work at a fundamental level
- What agents really are (LLM + tools + patterns)
- How different agent architectures function
- Why frameworks make certain design choices

**Philosophy**: Learn by building. Understand deeply, then use frameworks wisely.

## Why Go?

This is a Go port of the original Node.js implementation. Go offers:
- **Performance**: Compiled binaries with low memory footprint
- **Concurrency**: Native goroutines for parallel processing
- **Type Safety**: Strong typing catches errors at compile time
- **Simplicity**: Clean, readable code that's easy to understand
- **Deployment**: Single binary deployment with no runtime dependencies

## Next Phase: Build LangChain & LangGraph Concepts From Scratch

> After mastering the fundamentals, the next stage of this project walks you through **re-implementing the core parts of LangChain and LangGraph** in plain JavaScript using local models.
> This is **not** about building a new framework, it’s about understanding *how frameworks work*.  

## Phase 1: Agent Fundamentals - From LLMs to ReAct

### Quick Start

**New to this project?** Start here: [Quick Start Guide (Go)](QUICKSTART_GO.md)

### Prerequisites
- Go 1.23 or later
- At least 8GB RAM (16GB recommended)
- GCC or compatible C compiler (for CGO)
- Download models and place in `./models/` folder, details in [DOWNLOAD.md](DOWNLOAD.md)

### Installation
```bash
# Clone the repository
git clone https://github.com/mgreau/ai-agents-from-scratch-go.git
cd ai-agents-from-scratch-go

# Download dependencies
make deps

# Build all examples
make build
```

### Run Examples
```bash
# Run individual examples
make run-intro          # Basic LLM interaction
make run-translation    # System prompts & specialization
make run-coding         # Streaming responses
make run-agent          # Simple agent with tools
make run-react          # ReAct agent (reasoning + acting)

# Or run binaries directly
./bin/intro
./bin/simple-agent
./bin/react-agent
```

## Learning Path

Follow these examples in order to build understanding progressively:

### 1. **Introduction** - Basic LLM Interaction
[Go Code](examples-go/01_intro/main.go)

**What you'll learn:**
- Loading and running a local LLM
- Basic prompt/response cycle

**Key concepts**: Model loading, context, inference pipeline, token generation

---

### 2. **Translation** - System Prompts & Specialization
[Go Code](examples-go/03_translation/main.go)

**What you'll learn:**
- Using system prompts to specialize agents
- Output format control
- Role-based behavior
- Chat wrappers for different models

**Key concepts**: System prompts, agent specialization, behavioral constraints, prompt engineering

---

### 3. **Think** - Reasoning & Problem Solving
[Go Code](examples-go/04_think/main.go)

**What you'll learn:**
- Configuring LLMs for logical reasoning
- Complex quantitative problems
- Limitations of pure LLM reasoning
- When to use external tools

**Key concepts**: Reasoning agents, problem decomposition, cognitive tasks, reasoning limitations

---

### 4. **Batch** - Parallel Processing
[Go Code](examples-go/05_batch/main.go)

**What you'll learn:**
- Processing multiple requests concurrently
- Context sequences for parallelism
- GPU batch processing
- Performance optimization

**Key concepts**: Parallel execution, sequences, batch size, throughput optimization

---

### 5. **Coding** - Streaming & Response Control
[Go Code](examples-go/06_coding/main.go)

**What you'll learn:**
- Real-time streaming responses
- Token limits and budget management
- Progressive output display
- User experience optimization

**Key concepts**: Streaming, token-by-token generation, response control, real-time feedback

---

### 6. **Simple Agent** - Function Calling (Tools)
[Go Code](examples-go/07_simple-agent/main.go)

**What you'll learn:**
- Function calling / tool use fundamentals
- Defining tools the LLM can use
- JSON Schema for parameters
- How LLMs decide when to use tools

**Key concepts**: Function calling, tool definitions, agent decision making, action-taking

**This is where text generation becomes agency!**

---

### 7. **ReAct Agent** - Reasoning + Acting
[Go Code](examples-go/09_react-agent/main.go)

**What you'll learn:**
- ReAct pattern (Reason → Act → Observe)
- Iterative problem solving
- Step-by-step tool use
- Self-correction loops

**Key concepts**: ReAct pattern, iterative reasoning, observation-action cycles, multi-step agents

**This is the foundation of modern agent frameworks!**

---



## Documentation Structure

Each example includes:

- **`main.go`** - Complete working Go code
- Inline comments explaining key concepts
- Clean, idiomatic Go patterns

Additional documentation:
- **[examples-go/README.md](examples-go/README.md)** - Overview of all examples
- **[QUICKSTART_GO.md](QUICKSTART_GO.md)** - Setup guide
- **[SETUP_LLAMA_CPP.md](SETUP_LLAMA_CPP.md)** - Troubleshooting

## Core Concepts

### What is an AI Agent?

```
AI Agent = LLM + System Prompt + Tools + Memory + Reasoning Pattern
           ─┬─   ──────┬──────   ──┬──   ──┬───   ────────┬────────
            │          │           │       │              │
         Brain      Identity    Hands   State         Strategy
```

### Evolution of Capabilities

```
1. intro          → Basic LLM usage
2. translation    → Specialized behavior (system prompts)
3. think          → Reasoning ability
4. batch          → Parallel processing (goroutines)
5. coding         → Streaming & channels
6. simple-agent   → Tool use (function calling)
7. react-agent    → Strategic reasoning + tool use
```

### Architecture Patterns

**Simple Agent (Steps 1-5)**
```
User → LLM → Response
```

**Tool-Using Agent (Step 6)**
```
User → LLM ⟷ Tools → Response
```

**ReAct Agent (Step 7)**
```
User → LLM → Think → Act → Observe
       ↑      ↓      ↓      ↓
       └──────┴──────┴──────┘
           Iterate until solved
```

## ️ Debugging and Logging

Go provides excellent built-in tools for debugging:

- Use `fmt.Printf()` for basic debugging
- Enable verbose mode in agents (see ReAct example)
- Use Go's `log` package for structured logging
- Examine context and message flow in your code

## ️ Project Structure - Go Edition

```
ai-agents-from-scratch-go/
├── README.md                           ← You are here
├── go.mod                              ← Go module definition
├── Makefile                            ← Build and run commands
├── examples-go/                        ← Go examples
│   ├── 01_intro/
│   │   └── main.go
│   ├── 03_translation/
│   │   └── main.go
│   ├── 06_coding/
│   │   └── main.go
│   ├── 07_simple-agent/
│   │   └── main.go
│   └── 09_react-agent/
│       └── main.go
├── pkg/                                ← Go packages
│   ├── core/                           ← Core abstractions
│   │   ├── runnable.go
│   │   ├── message.go
│   │   └── context.go
│   ├── llm/                            ← LLM implementations
│   │   └── llama.go
│   ├── tools/                          ← Tool definitions
│   │   └── base.go
│   ├── agents/                         ← Agent implementations
│   │   └── react.go
│   ├── chains/                         ← Chain implementations
│   ├── memory/                         ← Memory implementations
│   ├── prompts/                        ← Prompt templates
│   ├── parsers/                        ← Output parsers
│   └── graph/                          ← Graph implementations
├── examples/                           ← Original JS examples (reference)
├── models/                             ← Place your GGUF models here
├── bin/                                ← Compiled binaries
└── logs/                               ← Debug outputs
```

### Go vs JavaScript Structure

The Go version follows idiomatic Go patterns:
- **Packages**: Code organized in `pkg/` directory with clear package boundaries
- **Examples**: Each example is a standalone `main.go` that can be compiled
- **Binaries**: Compiled executables in `bin/` directory
- **Interfaces**: Go interfaces for extensibility (Runnable, Tool, Message, etc.)
- **Concurrency**: Goroutines for parallel processing instead of Promises







## Key Takeaways

### What You'll Learn:

1. **LLMs are stateless**: Context must be managed explicitly in Go
2. **System prompts shape behavior**: Same model, different roles
3. **Function calling enables agency**: Tools transform text generators into agents
4. **Goroutines for parallelism**: Native concurrency for batch processing
5. **Channels for streaming**: Real-time token generation with Go channels
6. **Reasoning patterns matter**: ReAct > simple prompting for complex tasks
7. **Type safety catches errors**: Compile-time checking prevents runtime issues
8. **Interface-based design**: Extensible, testable architecture

### Go-Specific Advantages:

1. **The Runnable pattern**: Composable components with a unified interface
2. **Type-safe messages**: Compile-time guarantees for conversation flow
3. **Context for cancellation**: Proper resource management and timeouts
4. **Defer for cleanup**: Automatic resource cleanup (model.Close())
5. **Goroutines for concurrency**: Efficient parallel processing
6. **Single binary deployment**: No runtime dependencies

### Using Frameworks:

Now that you understand the fundamentals, frameworks like LangChainGo or similar provide:
- Pre-built reasoning patterns and agent templates
- Extensive tool libraries and integrations
- Production-ready error handling and retries
- Multi-agent orchestration
- Observability and monitoring

**You'll use them better because you know what they're doing under the hood.**

## Additional Resources

- **node-llama-cpp**: [GitHub](https://github.com/withcatai/node-llama-cpp)
- **Model Hub**: [Hugging Face](https://huggingface.co/models?library=gguf)
- **GGUF Format**: Quantized models for local inference

## Contributing

This is a learning resource. Feel free to:
- Suggest improvements to documentation
- Add more example patterns
- Fix bugs or unclear explanations
- Share what you built!

## License

Educational resource - use and modify as needed for learning.

---

## Go Conversion Status

This repository has been successfully converted to Go! 🎉

**What's Complete:**
- ✅ Core architecture (Runnable, Messages, Context)
- ✅ Tool system with registry
- ✅ ReAct agent implementation  
- ✅ 7 working examples
- ✅ LLM integration (go-skynet/go-llama.cpp)
- ✅ Comprehensive documentation
- ✅ Build system (Makefile)

**Setup Required:**
- ⚠️ Build llama.cpp C++ library (one-time setup)
- ⚠️ Configure CGO environment variables
- See [SETUP_LLAMA_CPP.md](SETUP_LLAMA_CPP.md) for detailed instructions

**Optional Enhancements:**
- 🔄 Additional examples (memory-agent, aot-agent)
- 🔄 Advanced components (chains, graphs, parsers)

**Quick Setup:**
```bash
# 1. Build llama.cpp (one-time setup)
cd ~ && git clone https://github.com/ggerganov/llama.cpp.git
cd llama.cpp && mkdir build && cd build
cmake .. && cmake --build . --config Release

# 2. Set environment variables (add to ~/.bashrc or ~/.zshrc)
export LLAMA_CPP_DIR="$HOME/llama.cpp"
export CGO_CFLAGS="-I${LLAMA_CPP_DIR}"
export CGO_LDFLAGS="-L${LLAMA_CPP_DIR}/build -lllama"
export LD_LIBRARY_PATH="${LLAMA_CPP_DIR}/build:${LD_LIBRARY_PATH}"
source ~/.bashrc  # or ~/.zshrc

# 3. Build and run examples
cd /path/to/ai-agents-from-scratch-go
make deps
make build
make run-intro
```

**Setup Guide:**  
Having issues? See [SETUP_LLAMA_CPP.md](SETUP_LLAMA_CPP.md) for detailed troubleshooting.

**Detailed Instructions:**  
Step-by-step guide: [QUICKSTART_GO.md](QUICKSTART_GO.md)

**Conversion Details:**  
See [CONVERSION_SUMMARY.md](CONVERSION_SUMMARY.md) for full conversion details.

---

**Built with ❤️ for people who want to truly understand AI agents**

This Go implementation provides:
- Type safety and compile-time checks
- Better performance and lower memory usage
- Native concurrency with goroutines
- Single binary deployment
- Clean, idiomatic Go code

Start with [QUICKSTART_GO.md](QUICKSTART_GO.md) to begin building AI agents in Go!

Happy learning! 🚀 
