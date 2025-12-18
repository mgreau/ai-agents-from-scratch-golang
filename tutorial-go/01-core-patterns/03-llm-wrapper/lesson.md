# The LLM Wrapper

**Part 1: Foundation - Lesson 3**

> Wrapping go-llama.cpp as a Runnable for seamless integration

## Overview

In Lesson 1, you learned about Runnables - the composable interface. In Lesson 2, you mastered Messages - the data structures. Now we'll connect these concepts by wrapping **go-llama.cpp** (our local LLM library) as a Runnable that understands Messages.

By the end of this lesson, you'll have an LLM wrapper that can generate text, handle conversations, stream responses, and integrate seamlessly with chains.

## Why Does This Matter?

### The Problem: LLMs Don't Compose

go-llama.cpp is excellent at what it does - running local LLMs efficiently. But when you're building agents, you need more than just an LLM. You need components that work together seamlessly.

**Without a composable framework:**
```go
// Each component is isolated
func myAgent(ctx context.Context, userInput string) (string, error) {
    // Step 1: Format the prompt
    prompt := myCustomFormatter(userInput)
    
    // Step 2: Call the LLM
    model, _ := llama.New("./model.gguf", llama.SetContext(2048))
    defer model.Free()
    
    response, _ := model.Predict(prompt)
    
    // Step 3: Parse the response
    parsed := myCustomParser(response)
    
    // Step 4: Maybe call a tool?
    if parsed.needsTool {
        toolResult := myTool(parsed.args)
        // Now what? Call LLM again? How do we loop?
    }
    
    return parsed, nil
}

// Problems:
// - Can't reuse components
// - Can't chain operations
// - Hard to add logging or metrics
// - Complex control flow for agents
// - Every new feature requires changing everything
```

**With a composable framework:**
```go
// Components that work together
llm, _ := NewLlamaCppLLM(LlamaCppConfig{
    ModelPath: "./model.gguf",
})

// Simple usage
messages := []Message{
    NewSystemMessage("You are helpful", nil),
    NewHumanMessage("Hi", nil),
}
response, _ := llm.Invoke(ctx, messages, nil)
// Returns: AIMessage("Hello! How can I help you?")

// But the real power is composition
agent := promptTemplate.
    Pipe(llm).
    Pipe(outputParser).
    Pipe(toolExecutor)

// Now you can:
// ✅ Reuse components in different chains
// ✅ Add logging with callbacks (no code changes)
// ✅ Build complex agents that use tools
// ✅ Test each component independently
// ✅ Swap LLMs without rewriting everything
```

### What the Wrapper Provides

The LLM wrapper isn't about making go-llama.cpp easier - it's about making it **work with everything else**:

1. **Common Interface**: Same `Invoke()` / `Stream()` / `Batch()` as every other component
2. **Message Support**: Understands HumanMessage, AIMessage, SystemMessage
3. **Composability**: Works with `.Pipe()` to chain operations
4. **Context Management**: Proper use of context.Context for cancellation
5. **Resource Cleanup**: Automatic model management with defer
6. **Configuration**: Runtime settings (temperature, tokens, etc.)

Think of it as an adapter that lets go-llama.cpp play nicely with the rest of your agent system.

## Learning Objectives

By the end of this lesson, you will:

- ✅ Understand how to wrap complex libraries as Runnables
- ✅ Convert Messages to LLM prompt format
- ✅ Handle model loading and lifecycle
- ✅ Implement streaming for real-time output
- ✅ Add temperature and other generation parameters
- ✅ Manage context windows and chat history
- ✅ Handle batch processing with goroutines

## Core Concepts

### What is an LLM Wrapper?

An LLM wrapper is an abstraction layer that:
1. **Hides complexity** - No need to manage model lifecycle manually
2. **Provides a standard interface** - Same API regardless of underlying model
3. **Handles conversion** - Transforms Messages into model-specific format
4. **Manages resources** - Automatic initialization and cleanup
5. **Enables composition** - Works seamlessly in chains
6. **Isolates state** - Prevents history contamination in batch processing

### The Wrapper's Responsibilities

```
Input (Messages or string)
      ↓
[1. Convert to Prompt Format]
      ↓
[2. Manage System Prompt]
      ↓
[3. Call LLM]
      ↓
[4. Parse Response]
      ↓
Output (string or AIMessage)
```

### Key Challenges in Go

1. **Model Loading**: Models are large, need careful memory management
2. **Prompt Formatting**: Convert Messages to text format
3. **System Prompt Management**: Inject into prompt correctly
4. **Context Cancellation**: Respect context.Context throughout
5. **Streaming**: Real-time output with channels
6. **Batch Isolation**: Each batch call needs independent state
7. **Error Handling**: Models can fail in various ways
8. **Resource Cleanup**: Always defer model.Free()

## Implementation Deep Dive

Let's build the LLM wrapper step by step.

### Step 1: The Base Structure

**Location:** `pkg/llm/llama.go`
```go
package llm

import (
    "context"
    "fmt"
    "os"
    
    llama "github.com/go-skynet/go-llama.cpp"
    "github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// LlamaCppLLM wraps go-llama.cpp for local inference
type LlamaCppLLM struct {
    *core.BaseRunnable
    model          *llama.LLama
    modelPath      string
    contextSize    int
    temperature    float32
    topP           float32
    topK           int
    threads        int
    systemPrompt   string
}

// LlamaCppConfig holds configuration
type LlamaCppConfig struct {
    ModelPath    string
    ContextSize  int
    Temperature  float32
    TopP         float32
    TopK         int
    Threads      int
    SystemPrompt string
}

func NewLlamaCppLLM(config LlamaCppConfig) (*LlamaCppLLM, error) {
    // Set defaults
    if config.ContextSize == 0 {
        config.ContextSize = 2048
    }
    if config.Temperature == 0 {
        config.Temperature = 0.7
    }
    if config.TopP == 0 {
        config.TopP = 0.9
    }
    if config.TopK == 0 {
        config.TopK = 40
    }
    if config.Threads == 0 {
        config.Threads = 4
    }
    
    // Check model exists
    if _, err := os.Stat(config.ModelPath); os.IsNotExist(err) {
        return nil, fmt.Errorf("model file not found: %s", config.ModelPath)
    }
    
    l := &LlamaCppLLM{
        BaseRunnable: core.NewBaseRunnable("LlamaCppLLM"),
        modelPath:    config.ModelPath,
        contextSize:  config.ContextSize,
        temperature:  config.Temperature,
        topP:         config.TopP,
        topK:         config.TopK,
        threads:      config.Threads,
        systemPrompt: config.SystemPrompt,
    }
    
    // Load model
    model, err := llama.New(
        config.ModelPath,
        llama.SetContext(config.ContextSize),
        llama.SetThreads(config.Threads),
    )
    if err != nil {
        return nil, fmt.Errorf("failed to load model: %w", err)
    }
    l.model = model
    
    return l, nil
}

// Close releases model resources
func (l *LlamaCppLLM) Close() {
    if l.model != nil {
        l.model.Free()
    }
}
```

**Key Go patterns:**
- Config struct for initialization
- Explicit resource management with Close()
- Error checking for file existence
- Sensible defaults

### Step 2: Converting Messages to Prompt

```go
// messagesToPrompt converts messages to a prompt string
func (l *LlamaCppLLM) messagesToPrompt(messages []core.Message) string {
    var prompt string
    
    for _, msg := range messages {
        switch msg.GetType() {
        case core.MessageTypeSystem:
            prompt += fmt.Sprintf("System: %s\n\n", msg.GetContent())
        case core.MessageTypeHuman:
            prompt += fmt.Sprintf("User: %s\n\n", msg.GetContent())
        case core.MessageTypeAI:
            prompt += fmt.Sprintf("Assistant: %s\n\n", msg.GetContent())
        case core.MessageTypeTool:
            prompt += fmt.Sprintf("Tool: %s\n\n", msg.GetContent())
        }
    }
    
    // Add final prompt for assistant to respond
    prompt += "Assistant:"
    
    return prompt
}
```

**Why this format?**
- Clear role labels help the model understand context
- Double newlines separate messages
- Final "Assistant:" prompts the model to respond

### Step 3: Implementing Invoke

```go
// Invoke generates a response for the given prompt
func (l *LlamaCppLLM) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
    var prompt string
    
    // Handle different input types
    switch v := input.(type) {
    case string:
        prompt = v
    case []core.Message:
        prompt = l.messagesToPrompt(v)
    default:
        return nil, fmt.Errorf("input must be string or []core.Message, got %T", input)
    }
    
    // Add system prompt if set
    if l.systemPrompt != "" {
        prompt = fmt.Sprintf("System: %s\n\nUser: %s\n\nAssistant:", l.systemPrompt, prompt)
    }
    
    // Generate response
    result, err := l.model.Predict(
        prompt,
        llama.SetTemperature(float64(l.temperature)),
        llama.SetTopP(float64(l.topP)),
        llama.SetTopK(l.topK),
        llama.SetThreads(l.threads),
        llama.SetTokens(l.contextSize),
    )
    if err != nil {
        return nil, fmt.Errorf("prediction failed: %w", err)
    }
    
    return result, nil
}
```

**Features:**
- Accepts both string and []Message inputs
- Injects system prompt automatically
- Respects temperature and other parameters
- Returns string (could be AIMessage)

### Step 4: Implementing Stream

```go
// Stream generates a response and streams tokens via channel
func (l *LlamaCppLLM) Stream(ctx context.Context, input interface{}, config *core.Config) (<-chan interface{}, error) {
    var prompt string
    
    // Convert input to prompt
    switch v := input.(type) {
    case string:
        prompt = v
    case []core.Message:
        prompt = l.messagesToPrompt(v)
    default:
        return nil, fmt.Errorf("input must be string or []core.Message")
    }
    
    // Add system prompt
    if l.systemPrompt != "" {
        prompt = fmt.Sprintf("System: %s\n\nUser: %s\n\nAssistant:", l.systemPrompt, prompt)
    }
    
    out := make(chan interface{}, 10)
    
    go func() {
        defer close(out)
        
        // Stream response using callback
        err := l.model.Predict(
            prompt,
            func(token string) bool {
                select {
                case <-ctx.Done():
                    return false // Stop generation
                case out <- token:
                    return true // Continue
                }
            },
            llama.SetTemperature(float64(l.temperature)),
            llama.SetTopP(float64(l.topP)),
            llama.SetTopK(l.topK),
            llama.SetThreads(l.threads),
            llama.SetTokens(l.contextSize),
        )
        if err != nil {
            out <- fmt.Errorf("streaming failed: %w", err)
        }
    }()
    
    return out, nil
}
```

**Go streaming patterns:**
- Goroutine for async generation
- Channel for communication
- `defer close(out)` ensures cleanup
- Context cancellation support

## Practical Examples

### Example 1: Simple Q&A

```go
func simpleQA() error {
    llm, err := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    if err != nil {
        return err
    }
    defer llm.Close()
    
    ctx := context.Background()
    response, err := llm.Invoke(ctx, "What is Go?", nil)
    if err != nil {
        return err
    }
    
    fmt.Println("AI:", response)
    return nil
}
```

### Example 2: Conversation with Messages

```go
func conversation() error {
    llm, err := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath:    "models/Qwen3-1.7B-Q8_0.gguf",
        SystemPrompt: "You are a helpful coding assistant",
    })
    if err != nil {
        return err
    }
    defer llm.Close()
    
    messages := []core.Message{
        core.NewHumanMessage("How do I read a file in Go?", nil),
    }
    
    ctx := context.Background()
    response, err := llm.Invoke(ctx, messages, nil)
    if err != nil {
        return err
    }
    
    fmt.Println("AI:", response)
    return nil
}
```

### Example 3: Streaming Response

```go
func streamingResponse() error {
    llm, err := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    if err != nil {
        return err
    }
    defer llm.Close()
    
    ctx := context.Background()
    stream, err := llm.Stream(ctx, "Tell me a story about Go", nil)
    if err != nil {
        return err
    }
    
    fmt.Print("AI: ")
    for token := range stream {
        if err, ok := token.(error); ok {
            return err
        }
        fmt.Print(token)
    }
    fmt.Println()
    
    return nil
}
```

### Example 4: With Cancellation

```go
func withCancellation() error {
    llm, err := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    if err != nil {
        return err
    }
    defer llm.Close()
    
    // Context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    stream, err := llm.Stream(ctx, "Write a long essay", nil)
    if err != nil {
        return err
    }
    
    for token := range stream {
        if err, ok := token.(error); ok {
            if ctx.Err() == context.DeadlineExceeded {
                fmt.Println("\nTimeout reached")
                return nil
            }
            return err
        }
        fmt.Print(token)
    }
    
    return nil
}
```

## Exercises

### Exercise 9: Basic LLM Wrapper
Build a minimal LLM wrapper with Invoke only.
- [Starter Code](exercises/09-basic-llm-wrapper/starter.go)
- [Solution](exercises/09-basic-llm-wrapper/solution.go)

### Exercise 10: Batch Processing
Implement efficient batch processing with goroutines.
- [Starter Code](exercises/10-batch-processing/starter.go)
- [Solution](exercises/10-batch-processing/solution.go)

### Exercise 11: Streaming Implementation
Add proper streaming support with buffering.
- [Starter Code](exercises/11-streaming/starter.go)
- [Solution](exercises/11-streaming/solution.go)

### Exercise 12: Composition
Create a pipeline: prompt template → LLM → parser.
- [Starter Code](exercises/12-composition/starter.go)
- [Solution](exercises/12-composition/solution.go)

## Key Takeaways

1. ✅ **LLM as Runnable** - Same interface as other components
2. ✅ **Message support** - Converts Messages to prompts
3. ✅ **Resource management** - defer model.Close()
4. ✅ **Context support** - Cancellation and timeouts
5. ✅ **Streaming** - Real-time with channels and goroutines
6. ✅ **Configuration** - Temperature, tokens, etc.
7. ✅ **Type flexibility** - Accepts string or []Message

## What's Next

**Next Lesson**: [04-context](../04-context/lesson.md) - Context & Configuration

**See it in action**: Check `pkg/llm/llama.go` for the full implementation

**Practice**: Complete all 4 exercises before moving on

## Further Reading

- [go-llama.cpp Docs](https://github.com/go-skynet/go-llama.cpp)
- [Context Package](https://pkg.go.dev/context)
- [GGUF Format](https://github.com/ggerganov/ggml/blob/master/docs/gguf.md)
