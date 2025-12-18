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
`intro/` | [Code](examples/01_intro/intro.js) | [Code Explanation](examples/01_intro/CODE.md) | [Concepts](examples/01_intro/CONCEPT.md)

**What you'll learn:**
- Loading and running a local LLM
- Basic prompt/response cycle

**Key concepts**: Model loading, context, inference pipeline, token generation

---

### 2. (Optional) **OpenAI Intro** - Using Proprietary Models
`openai-intro/` | [Code](examples/02_openai-intro/openai-intro.js) | [Code Explanation](examples/02_openai-intro/CODE.md) | [Concepts](examples/02_openai-intro/CONCEPT.md)

**What you'll learn:**
- How to call hosted LLMs (like GPT-4)
- Temperature Control
- Token Usage

**Key concepts**: Inference endpoints, network latency, cost vs control, data privacy, vendor dependence

---

### 3. **Translation** - System Prompts & Specialization
`translation/` | [Code](examples/03_translation/translation.js) | [Code Explanation](examples/03_translation/CODE.md) | [Concepts](examples/03_translation/CONCEPT.md)

**What you'll learn:**
- Using system prompts to specialize agents
- Output format control
- Role-based behavior
- Chat wrappers for different models

**Key concepts**: System prompts, agent specialization, behavioral constraints, prompt engineering

---

### 4. **Think** - Reasoning & Problem Solving
`think/` | [Code](examples/04_think/think.js) | [Code Explanation](examples/04_think/CODE.md) | [Concepts](examples/04_think/CONCEPT.md)

**What you'll learn:**
- Configuring LLMs for logical reasoning
- Complex quantitative problems
- Limitations of pure LLM reasoning
- When to use external tools

**Key concepts**: Reasoning agents, problem decomposition, cognitive tasks, reasoning limitations

---

### 5. **Batch** - Parallel Processing
`batch/` | [Code](examples/05_batch/batch.js) | [Code Explanation](examples/05_batch/CODE.md) | [Concepts](examples/05_batch/CONCEPT.md)

**What you'll learn:**
- Processing multiple requests concurrently
- Context sequences for parallelism
- GPU batch processing
- Performance optimization

**Key concepts**: Parallel execution, sequences, batch size, throughput optimization

---

### 6. **Coding** - Streaming & Response Control
`coding/` | [Code](examples/06_coding/coding.js) | [Code Explanation](examples/06_coding/CODE.md) | [Concepts](examples/06_coding/CONCEPT.md)

**What you'll learn:**
- Real-time streaming responses
- Token limits and budget management
- Progressive output display
- User experience optimization

**Key concepts**: Streaming, token-by-token generation, response control, real-time feedback

---

### 7. **Simple Agent** - Function Calling (Tools)
`simple-agent/` | [Code](examples/07_simple-agent/simple-agent.js) | [Code Explanation](examples/07_simple-agent/CODE.md) | [Concepts](examples/07_simple-agent/CONCEPT.md)

**What you'll learn:**
- Function calling / tool use fundamentals
- Defining tools the LLM can use
- JSON Schema for parameters
- How LLMs decide when to use tools

**Key concepts**: Function calling, tool definitions, agent decision making, action-taking

**This is where text generation becomes agency!**

---

### 8. **Simple Agent with Memory** - Persistent State
`simple-agent-with-memory/` | [Code](examples/08_simple-agent-with-memory/simple-agent-with-memory.js) | [Code Explanation](examples/08_simple-agent-with-memory/CODE.md) | [Concepts](examples/08_simple-agent-with-memory/CONCEPT.md)

**What you'll learn:**
- Persisting information across sessions
- Long-term memory management
- Facts and preferences storage
- Memory retrieval strategies

**Key concepts**: Persistent memory, state management, memory systems, context augmentation

---

### 9. **ReAct Agent** - Reasoning + Acting
`react-agent/` | [Code](examples/09_react-agent/react-agent.js) | [Code Explanation](examples/09_react-agent/CODE.md) | [Concepts](examples/09_react-agent/CONCEPT.md)

**What you'll learn:**
- ReAct pattern (Reason → Act → Observe)
- Iterative problem solving
- Step-by-step tool use
- Self-correction loops

**Key concepts**: ReAct pattern, iterative reasoning, observation-action cycles, multi-step agents

**This is the foundation of modern agent frameworks!**

---

### 10. **AoT Agent** - Atom of Thought Planning
`aot-agent/` | [Code](examples/10_aot-agent/aot-agent.js) | [Code Explanation](examples/10_aot-agent/CODE.md) | [Concepts](examples/10_aot-agent/CONCEPT.md)

**What you'll learn:**
- Atom of Thought methodology
- Atomic planning for multi-step computations
- Dependency management between operations
- Structured JSON output for reasoning plans
- Deterministic execution of plans

**Key concepts**: AoT planning, atomic operations, dependency resolution, plan validation, structured reasoning

**This is the foundation for agents that separate planning from execution, enabling precise and auditable multi-step reasoning!**

---

## Documentation Structure

Each example folder contains:

- **`<name>.js`** - The working code example
- **`CODE.md`** - Step-by-step code explanation
- Line-by-line breakdowns
- What each part does
- How it works
- **`CONCEPT.md`** - High-level concepts
- Why it matters for agents
- Architectural patterns
- Real-world applications
- Simple diagrams

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
4. batch          → Parallel processing
5. coding         → Streaming & control
6. simple-agent   → Tool use (function calling)
7. memory-agent   → Persistent state
8. react-agent    → Strategic reasoning + tool use
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

**Memory Agent (Step 7)**
```
User → LLM ⟷ Tools → Response
       ↕
     Memory
```

**ReAct Agent (Step 8)**
```
User → LLM → Think → Act → Observe
       ↑      ↓      ↓      ↓
       └──────┴──────┴──────┘
           Iterate until solved
```

## ️ Helper Utilities

### PromptDebugger
`helper/prompt-debugger.js`

Utility for debugging prompts sent to the LLM. Shows exactly what the model sees, including:
- System prompts
- Function definitions
- Conversation history
- Context state

Usage example in `simple-agent/simple-agent.js`

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

## Phase 2: Building a Production Framework (Tutorial)

After mastering the fundamentals above, **Phase 2** takes you from scratch examples to production-grade framework design. You'll rebuild core concepts from **LangChain** and **LangGraph** to understand how real frameworks work internally.

### What You'll Build

A lightweight but complete agent framework with:
- **Runnable Interface**, The composability pattern that powers everything
- **Message System**, Typed conversation structures (Human, AI, System, Tool)
- **Chains**, Composing multiple operations into pipelines
- **Memory**, Persistent state across conversations
- **Tools**, Function calling and external integrations
- **Agents**, Decision-making loops (ReAct, Tool-calling)
- **Graphs**, State machines for complex workflows (LangGraph concepts)

### Learning Approach

**Tutorial-first**: Step-by-step lessons with exercises  
**Implementation-driven**: Build each component yourself  
**Framework-compatible**: Learn patterns used in LangChain.js

### Structure Overview

```
tutorial/
├── 01-foundation/              # 1. Core Abstractions
│   ├── 01-runnable/
│   │   ├── lesson.md           # Why Runnable matters
│   │   ├── exercises/          # Hands-on practice
│   │   └── solutions/          # Reference implementations
│   ├── 02-messages/            # Structuring conversations
│   ├── 03-llm-wrapper/         # Wrapping node-llama-cpp
│   └── 04-context/             # Configuration & callbacks
│
├── 02-composition/             # 2. Building Chains
│   ├── 01-prompts/             # Template system
│   ├── 02-parsers/             # Structured outputs
│   ├── 03-llm-chain/           # Your first chain
│   ├── 04-piping/              # Composition patterns
│   └── 05-memory/              # Conversation state
│
├── 03-agency/                  # 3. Tools & Agents
│   ├── 01-tools/               # Function definitions
│   ├── 02-tool-executor/       # Safe execution
│   ├── 03-simple-agent/        # Basic agent loop
│   ├── 04-react-agent/         # Reasoning + Acting
│   └── 05-structured-agent/    # JSON mode
│
└── 04-graphs/                  # 4. State Machines
    ├── 01-state-basics/        # Nodes & edges
    ├── 02-channels/            # State management
    ├── 03-conditional-edges/   # Dynamic routing
    ├── 04-executor/            # Running workflows
    ├── 05-checkpointing/       # Persistence
    └── 06-agent-graph/         # Agents as graphs

src/
├── core/                       # Runnable, Messages, Context
├── llm/                        # LlamaCppLLM wrapper
├── prompts/                    # Template system
├── chains/                     # LLMChain, SequentialChain
├── tools/                      # BaseTool, built-in tools
├── agents/                     # AgentExecutor, ReActAgent
├── memory/                     # BufferMemory, WindowMemory
└── graph/                      # StateGraph, CompiledGraph
```

### Why This Matters

**Understanding beats using**: When you know how frameworks work internally, you can:
- Debug issues faster
- Customize behavior confidently
- Make architectural decisions wisely
- Build your own extensions
- Read framework source code fluently

**Learn once, use everywhere**: The patterns you'll learn (Runnable, composition, state machines) apply to:
- LangChain.js - You'll understand their abstractions
- LangGraph.js - You'll grasp state management
- Any agent framework - Same core concepts
- Your own projects - Build custom solutions

### Getting Started with Phase 2

After completing the fundamentals (intro → react-agent), start the tutorial:

[Overview](tutorial/README.md)

```bash
# Start with the foundation
cd tutorial/01-foundation/01-runnable
lesson.md                    # Read the lesson
node exercises/01-*.js           # Complete exercises
node solutions/01-*-solution.js  # Check your work
```

Each lesson includes:
- **Conceptual explanation**, Why it matters
- **Code walkthrough**, How to build it
- **Exercises**, Practice implementing
- **Solutions**, Reference code
- **Real-world examples**, Practical usage

**Time commitment**: ~8 weeks, 3-5 hours/week

### What You'll Achieve

By the end, you'll have:
1. Built a working agent framework from scratch
2. Understood how LangChain/LangGraph work internally
3. Mastered composability patterns
4. Created reusable components (tools, chains, agents)
5. Implemented state machines for complex workflows
6. Gained confidence to use or extend any framework

**Then**: Use LangChain.js in production, knowing exactly what happens under the hood.

---

## Key Takeaways

### After Phase 1 (Fundamentals), you'll understand:

1. **LLMs are stateless**: Context must be managed explicitly
2. **System prompts shape behavior**: Same model, different roles
3. **Function calling enables agency**: Tools transform text generators into agents
4. **Memory is essential**: Agents need to remember across sessions
5. **Reasoning patterns matter**: ReAct > simple prompting for complex tasks
6. **Performance matters**: Parallel processing, streaming, token limits
7. **Debugging is crucial**: See exactly what the model receives

### After Phase 2 (Framework Tutorial), you'll master:

1. **The Runnable pattern**: Why everything in frameworks uses one interface
2. **Composition over configuration**: Building complex systems from simple parts
3. **Message-driven architecture**: How frameworks structure conversations
4. **Chain abstraction**: Connecting prompts, LLMs, and parsers seamlessly
5. **Tool orchestration**: Safe execution with timeouts and error handling
6. **Agent execution loops**: The mechanics of decision-making agents
7. **State machines**: Managing complex workflows with graphs
8. **Production patterns**: Error handling, retries, streaming, and debugging

### What frameworks give you:

Now that you understand the fundamentals, frameworks like LangChain, CrewAI, or AutoGPT provide:
- Pre-built reasoning patterns and agent templates
- Extensive tool libraries and integrations
- Production-ready error handling and retries
- Multi-agent orchestration
- Observability and monitoring
- Community extensions and plugins

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

The Go version preserves all the educational value of the original while adding:
- Type safety and compile-time checks
- Better performance and lower memory usage
- Native concurrency with goroutines
- Single binary deployment

Start with [QUICKSTART_GO.md](QUICKSTART_GO.md) for the Go version, or explore the original JavaScript examples in `examples/` for comparison.

Happy learning! 🚀 
