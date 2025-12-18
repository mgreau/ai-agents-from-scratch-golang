# Part 1: Foundation - Core Patterns

**From Scripts to Abstractions**

Transform the patterns you already use into reusable components.

## Before You Start: Why This Tutorial Exists

**You've just built AI agents with Go and llama.cpp.** You know how to call LLMs, format prompts, parse responses, and create agent loops. That's awesome—you understand the fundamentals!

**But you probably noticed some friction:**
- Copy-pasting prompt formatting everywhere
- Manually building message arrays each time
- Hard to test individual components
- Difficult to swap out models or reuse patterns
- Agent code that works but feels messy

**This tutorial fixes those problems.** You'll transform the script-style code into clean, composable abstractions.

## 🎯 Learning Objectives

By the end of this tutorial, you'll understand:

1. **The Runnable Pattern** - Why and how to make everything composable
2. **Message System** - Type-safe conversation structures
3. **Context Management** - Configuration, callbacks, and cancellation
4. **Go Interfaces** - Extensibility through interface-based design

## 📚 Topics Covered

### 1. The Runnable Pattern

**Concept:** A unified interface for all components that allows composition and chaining.

**Why it matters:**
- Consistent API across all components
- Easy to compose complex workflows
- Simple to test and mock
- Familiar to users of LangChain/LangGraph

**Core interface:**
```go
type Runnable interface {
    Invoke(ctx context.Context, input interface{}, config *Config) (interface{}, error)
    Stream(ctx context.Context, input interface{}, config *Config) (<-chan interface{}, error)
    Batch(ctx context.Context, inputs []interface{}, config *Config) ([]interface{}, error)
    Pipe(other Runnable) Runnable
}
```

**Exercise:** Implement a simple Runnable that doubles numbers.

**Location:** [exercises/01-double-runnable/](exercises/01-double-runnable/)

---

### 2. Message System

**Concept:** Type-safe structures for different message types in conversations.

**Why it matters:**
- Clear separation of concerns
- Type safety prevents errors
- Easy to serialize/deserialize
- Matches LLM API requirements

**Key types:**
```go
type Message interface {
    GetContent() string
    GetType() MessageType
    ToPromptFormat() map[string]interface{}
}

// Implementations:
// - SystemMessage
// - HumanMessage
// - AIMessage
// - ToolMessage
```

**Exercise:** Build a conversation history manager.

**Location:** [exercises/02-conversation-manager/](exercises/02-conversation-manager/)

---

### 3. Context and Configuration

**Concept:** Managing execution context, configuration, and callbacks.

**Why it matters:**
- Cancellation and timeouts
- Observability through callbacks
- Runtime configuration
- Metadata passing

**Key structures:**
```go
type Config struct {
    Callbacks  []Callback
    Tags       []string
    Metadata   map[string]interface{}
    MaxRetries int
    Timeout    int
}

type Callback interface {
    OnStart(ctx context.Context, runnable Runnable, input interface{}) error
    OnEnd(ctx context.Context, runnable Runnable, output interface{}) error
    OnError(ctx context.Context, runnable Runnable, err error) error
}
```

**Exercise:** Implement a logging callback.

**Location:** [exercises/03-logging-callback/](exercises/03-logging-callback/)

---

### 4. Composition and Piping

**Concept:** Chaining Runnables together to create workflows.

**Why it matters:**
- Build complex logic from simple parts
- Reusable components
- Clear data flow
- Easy to test each stage

**Pattern:**
```go
// Sequential execution
result := runnable1.Pipe(runnable2).Pipe(runnable3).Invoke(ctx, input, nil)

// Parallel execution
parallel := NewRunnableParallel(map[string]Runnable{
    "task1": runnable1,
    "task2": runnable2,
})
results := parallel.Invoke(ctx, input, nil)
```

**Exercise:** Build a multi-stage text processor.

**Location:** [exercises/04-text-pipeline/](exercises/04-text-pipeline/)

---

## 🚀 Getting Started

### Prerequisites

- Completed examples in `examples-go/`
- Read `pkg/core/runnable.go`, `message.go`, `context.go`
- Basic understanding of Go interfaces

### Structure

Each exercise has:
```
exercises/XX-name/
├── README.md          # Detailed instructions
├── starter.go         # Starting point with TODOs
├── starter_test.go    # Tests to verify your solution
└── solution.go        # Reference implementation
```

### Workflow

1. **Read the exercise README**
2. **Study the starter code**
3. **Implement the TODOs**
4. **Run tests:** `go test`
5. **Compare with solution**
6. **Experiment and extend**

### Running Exercises

```bash
cd tutorial-go/01-core-patterns/exercises/01-double-runnable

# Read instructions
cat README.md

# Implement starter.go
vim starter.go

# Test your implementation
go test -v

# Compare with solution
diff starter.go solution.go
```

## 📝 Exercises Summary

| # | Name | Difficulty | Concepts |
|---|------|------------|----------|
| 01 | Double Runnable | Easy | Basic Runnable implementation |
| 02 | Conversation Manager | Medium | Message system, state management |
| 03 | Logging Callback | Easy | Callbacks, observability |
| 04 | Text Pipeline | Medium | Composition, piping |

## 💡 Key Takeaways

1. **Runnable is the foundation** - Everything implements this interface
2. **Type safety matters** - Use interfaces to catch errors at compile time
3. **Composition over inheritance** - Build complex behavior from simple parts
4. **Context is essential** - Always use context.Context for cancellation
5. **Defer for cleanup** - Always defer resource cleanup

## 🎯 Next Steps

After completing these exercises:

1. **Move to Tutorial 02** - [LLM Integration](../02-llm-integration/)
2. **Study pkg/llm/** - See how Runnable works with LLMs
3. **Experiment** - Combine concepts in your own project

## 🤔 Common Questions

**Q: Why use interface{} instead of generics?**
A: The original design predates Go generics. You can refactor to use generics for type safety.

**Q: Why not use channels for everything?**
A: Channels are great for streaming, but Invoke is simpler for single-shot operations.

**Q: How do I handle errors in Pipe?**
A: Errors stop the pipeline and propagate up. Add error handling runnables if needed.

**Q: Can I make Runnable generic?**
A: Yes! Exercise: Refactor to `Runnable[I, O]` for input/output types.

## 📚 Further Reading

- [pkg/core/runnable.go](../../pkg/core/runnable.go) - Full implementation
- [examples-go/](../../examples-go/) - Real usage examples
- [Go Interfaces](https://go.dev/tour/methods/9) - Go interface tutorial
- [Context Package](https://pkg.go.dev/context) - Using context.Context

## 🤝 Contributing

Improve this tutorial:
- Add more exercises
- Clarify explanations
- Fix errors
- Add diagrams

Submit a PR!
