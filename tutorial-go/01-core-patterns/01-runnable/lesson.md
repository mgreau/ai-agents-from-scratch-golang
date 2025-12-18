# The Runnable Contract

**Part 1: Foundation - Lesson 1**

> Understanding the single pattern that powers AI agent frameworks

## Overview

The `Runnable` is the fundamental building block of our framework. It's a simple yet powerful abstraction that allows us to build complex AI systems from composable parts. Think of it as the "contract" that every component in the framework must follow.

By the end of this lesson, you'll understand why frameworks like LangChain built everything around this single interface, and you'll implement your own Runnable components in Go.

## Why Does This Matter?

Imagine you're building with LEGO blocks. Each block has the same connection mechanism (those little bumps), which means any block can connect to any other block. The `Runnable` interface is exactly that for AI agents.

### The Problem Without Runnable

```go
// Without a common interface, every component is different:
llmResponse, _ := llm.Generate(prompt)
parsedOutput, _ := parser.Parse(llmResponse)
memorySaved, _ := memory.Store(parsedOutput)

// Different methods: Generate(), Parse(), Store()
// Hard to compose, hard to test, hard to maintain
```

### The Solution With Runnable

```go
// With Runnable, everything uses the same interface:
result, _ := prompt.
    Pipe(llm).
    Pipe(parser).
    Pipe(memory).
    Invoke(ctx, input, nil)

// Same method everywhere: Invoke()
// Easy to compose, test, and maintain
```

## Learning Objectives

By the end of this lesson, you will:

- ✅ Understand what makes a good abstraction in Go
- ✅ Implement the base `Runnable` interface
- ✅ Create custom Runnable components
- ✅ Know the three core execution patterns: `Invoke`, `Stream`, `Batch`
- ✅ Understand why this abstraction is powerful for AI systems

## Core Concepts

### What is a Runnable?

A `Runnable` is any component that can:
1. **Take input**
2. **Do something with it**
3. **Return output**

That's it! But this simplicity is what makes it powerful.

### The Three Execution Patterns

Every Runnable supports three ways of execution:

#### 1. `Invoke()` - Single Execution
Run once with one input, get one output.

```go
result, err := runnable.Invoke(ctx, input, nil)
// Input: "Hello"
// Output: "Hello, World!"
```

**Use case**: Normal execution, when you have one thing to process.

#### 2. `Stream()` - Streaming Execution
Process input and receive output in chunks as it's generated.

```go
stream, _ := runnable.Stream(ctx, input, nil)
for chunk := range stream {
    fmt.Print(chunk) // Print each piece as it arrives
}
// Output: "H", "e", "l", "l", "o", "..."
```

**Use case**: LLM text generation, where you want to show results in real-time.

#### 3. `Batch()` - Parallel Execution
Process multiple inputs at once using goroutines.

```go
results, _ := runnable.Batch(ctx, []interface{}{input1, input2, input3}, nil)
// Input: ["Hello", "Hi", "Hey"]
// Output: ["Hello, World!", "Hi, World!", "Hey, World!"]
```

**Use case**: Processing many items efficiently with Go's concurrency.

### The Magic: Composition

The real power comes from combining Runnables:

```go
pipeline := runnableA.Pipe(runnableB).Pipe(runnableC)
```

Because everything is a Runnable, you can chain them together infinitely!

## Implementation Deep Dive

Let's build the `Runnable` interface and base implementation step by step.

### Step 1: The Interface

**Location:** `pkg/core/runnable.go`

```go
package core

import "context"

// Runnable is the fundamental interface for all composable components.
// Every component in the framework implements this interface.
type Runnable interface {
    // Invoke processes a single input and returns output
    Invoke(ctx context.Context, input interface{}, config *Config) (interface{}, error)
    
    // Stream processes input and streams output via a channel
    Stream(ctx context.Context, input interface{}, config *Config) (<-chan interface{}, error)
    
    // Batch processes multiple inputs in parallel
    Batch(ctx context.Context, inputs []interface{}, config *Config) ([]interface{}, error)
    
    // Pipe composes this Runnable with another
    Pipe(other Runnable) Runnable
    
    // Name returns the name of this Runnable
    Name() string
}
```

**Key Go patterns used:**
- `context.Context` for cancellation and timeouts
- `interface{}` for flexible input/output (or use generics)
- `<-chan` for streaming results
- Method chaining with `Pipe()`

### Step 2: The Base Implementation

```go
// BaseRunnable provides default implementations
type BaseRunnable struct {
    name string
}

func NewBaseRunnable(name string) *BaseRunnable {
    return &BaseRunnable{name: name}
}

func (r *BaseRunnable) Name() string {
    return r.name
}

// Invoke is the main execution method
func (r *BaseRunnable) Invoke(ctx context.Context, input interface{}, config *Config) (interface{}, error) {
    if config == nil {
        config = NewConfig()
    }
    
    // Execute the actual logic (must be overridden)
    return r.call(ctx, input, config)
}

// call is the internal method subclasses override
func (r *BaseRunnable) call(ctx context.Context, input interface{}, config *Config) (interface{}, error) {
    return nil, fmt.Errorf("%s must implement call() method", r.name)
}
```

**Why this design?**
- `BaseRunnable` provides common functionality
- Subclasses embed `*BaseRunnable` and override `call()`
- Go's embedding provides a form of inheritance
- `context.Context` enables cancellation

### Step 3: Implementing Stream

```go
// Stream provides default streaming by yielding the full result
func (r *BaseRunnable) Stream(ctx context.Context, input interface{}, config *Config) (<-chan interface{}, error) {
    out := make(chan interface{}, 1)
    
    go func() {
        defer close(out)
        
        result, err := r.Invoke(ctx, input, config)
        if err != nil {
            return
        }
        
        select {
        case <-ctx.Done():
            return
        case out <- result:
        }
    }()
    
    return out, nil
}
```

**Go patterns:**
- Goroutines for concurrency
- Channels for communication
- `defer close(out)` ensures cleanup
- `select` for context cancellation

### Step 4: Implementing Batch

```go
// Batch processes multiple inputs in parallel using goroutines
func (r *BaseRunnable) Batch(ctx context.Context, inputs []interface{}, config *Config) ([]interface{}, error) {
    results := make([]interface{}, len(inputs))
    errors := make([]error, len(inputs))
    
    done := make(chan bool, len(inputs))
    
    for i, input := range inputs {
        go func(idx int, inp interface{}) {
            results[idx], errors[idx] = r.Invoke(ctx, inp, config)
            done <- true
        }(i, input)
    }
    
    // Wait for all goroutines
    for range inputs {
        <-done
    }
    
    // Check for errors
    for _, err := range errors {
        if err != nil {
            return nil, err
        }
    }
    
    return results, nil
}
```

**Go advantages:**
- Native parallelism with goroutines
- Efficient with many inputs
- Cancellable via context

### Step 5: Implementing Pipe

```go
// Pipe composes this Runnable with another
func (r *BaseRunnable) Pipe(other Runnable) Runnable {
    return NewRunnableSequence([]Runnable{r, other})
}

// RunnableSequence chains multiple Runnables
type RunnableSequence struct {
    *BaseRunnable
    steps []Runnable
}

func NewRunnableSequence(steps []Runnable) *RunnableSequence {
    return &RunnableSequence{
        BaseRunnable: NewBaseRunnable("RunnableSequence"),
        steps:        steps,
    }
}

func (rs *RunnableSequence) Invoke(ctx context.Context, input interface{}, config *Config) (interface{}, error) {
    output := input
    
    for _, step := range rs.steps {
        var err error
        output, err = step.Invoke(ctx, output, config)
        if err != nil {
            return nil, err
        }
    }
    
    return output, nil
}
```

**Composition pattern:**
- Each step's output becomes the next step's input
- Errors stop the pipeline
- Context propagates through all steps

## Example: Building a Custom Runnable

Let's build a practical example - a Runnable that uppercases text:

```go
// UppercaseRunnable converts text to uppercase
type UppercaseRunnable struct {
    *BaseRunnable
}

func NewUppercaseRunnable() *UppercaseRunnable {
    return &UppercaseRunnable{
        BaseRunnable: NewBaseRunnable("Uppercase"),
    }
}

// Override Invoke to implement our logic
func (u *UppercaseRunnable) Invoke(ctx context.Context, input interface{}, config *Config) (interface{}, error) {
    str, ok := input.(string)
    if !ok {
        return nil, fmt.Errorf("input must be string, got %T", input)
    }
    
    return strings.ToUpper(str), nil
}

// Usage:
func main() {
    upper := NewUppercaseRunnable()
    
    result, _ := upper.Invoke(context.Background(), "hello", nil)
    fmt.Println(result) // "HELLO"
    
    // Compose with other Runnables
    pipeline := NewPrefixRunnable(">> ").Pipe(upper)
    result, _ = pipeline.Invoke(context.Background(), "hello", nil)
    fmt.Println(result) // ">> HELLO"
}
```

## Why This Pattern is Powerful

### 1. Composability
Build complex systems from simple parts:
```go
complex := step1.Pipe(step2).Pipe(step3).Pipe(step4)
```

### 2. Testability
Test each component in isolation:
```go
func TestUppercase(t *testing.T) {
    u := NewUppercaseRunnable()
    result, _ := u.Invoke(context.Background(), "test", nil)
    assert.Equal(t, "TEST", result)
}
```

### 3. Reusability
Use the same component in different contexts:
```go
// In a pipeline
pipeline1 := promptFormat.Pipe(llm).Pipe(uppercase)

// Standalone
result := uppercase.Invoke(ctx, input, nil)

// In parallel
results := uppercase.Batch(ctx, inputs, nil)
```

### 4. Type Safety (with Generics)
Optional: Use Go generics for type-safe Runnables:
```go
type Runnable[I, O any] interface {
    Invoke(ctx context.Context, input I, config *Config) (O, error)
    // ...
}

type UppercaseRunnable struct {
    *BaseRunnable[string, string]
}
```

## Exercises

### Exercise 1: Double Runnable (Warmup)
Implement a Runnable that doubles numbers.
- [Starter Code](exercises/01-double-runnable/starter.go)
- [Solution](exercises/01-double-runnable/solution.go)

### Exercise 2: JSON Parser Runnable
Build a Runnable that parses JSON strings.
- [Starter Code](exercises/02-json-parser/starter.go)
- [Solution](exercises/02-json-parser/solution.go)

### Exercise 3: Pipeline Composition
Create a multi-step text processing pipeline.
- [Starter Code](exercises/03-pipeline/starter.go)
- [Solution](exercises/03-pipeline/solution.go)

### Exercise 4: Streaming Implementation
Implement true character-by-character streaming.
- [Starter Code](exercises/04-streaming/starter.go)
- [Solution](exercises/04-streaming/solution.go)

## Key Takeaways

1. ✅ **Runnable is a universal interface** - Everything can be a Runnable
2. ✅ **Three execution patterns** - Invoke, Stream, Batch
3. ✅ **Composition through Pipe** - Build complex from simple
4. ✅ **Context for control** - Cancellation, timeouts, metadata
5. ✅ **Goroutines for concurrency** - Native parallel processing
6. ✅ **Type safety with interfaces** - Compile-time guarantees

## What's Next

**Next Lesson**: [02-messages](../02-messages/lesson.md) - Structured conversation data

**See it in action**: Check `pkg/core/runnable.go` for the full implementation

**Practice**: Complete all 4 exercises before moving on

## Further Reading

- [Go Interfaces](https://go.dev/tour/methods/9)
- [Context Package](https://pkg.go.dev/context)
- [Concurrency Patterns](https://go.dev/blog/pipelines)
- [LangChain Runnable Docs](https://python.langchain.com/docs/expression_language/)
