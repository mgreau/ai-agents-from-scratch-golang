# Context & Configuration

**Part 1: Foundation - Lesson 4**

> Passing state, callbacks, and metadata through Runnable chains in Go

## Overview

You've learned Runnables (Lesson 1), Messages (Lesson 2), and LLM Wrappers (Lesson 3). Now we tackle a critical question: **How do we pass configuration, callbacks, and state through complex chains without cluttering our code?**

The answer is **Config** - a powerful pattern that threads context through every step of a pipeline, enabling logging, debugging, authentication, and more without changing your core logic.

**Note**: Go already has `context.Context` for cancellation and deadlines. Our `Config` complements it by carrying application-specific data (callbacks, metadata, tags).

## Why Does This Matter?

### The Problem: Configuration Chaos

Without a proper config system:

```go
// Bad: Configuration everywhere
func complexPipeline(
    ctx context.Context,
    input string,
    temperature float32,
    callbacks []Callback,
    debug bool,
    userID string,
) (string, error) {
    result1, _ := step1(ctx, input, temperature, debug)
    for _, cb := range callbacks {
        cb.OnStep("step1", result1)
    }
    
    result2, _ := step2(ctx, result1, userID, debug)
    for _, cb := range callbacks {
        cb.OnStep("step2", result2)
    }
    
    result3, _ := step3(ctx, result2, temperature, callbacks, debug)
    for _, cb := range callbacks {
        cb.OnStep("step3", result3)
    }
    
    return result3, nil
}

// Every function needs to know about every configuration option!
```

Problems:
- Every function signature becomes huge
- Adding new config requires changing every function
- Hard to add features like logging or metrics
- Can't intercept at specific points
- Passing user context is messy

### The Solution: Config Pattern

With Config:

```go
// Good: Config flows automatically
config := &Config{
    Temperature: 0.7,
    Callbacks:   []Callback{loggingCallback, metricsCallback},
    Metadata: map[string]interface{}{
        "userID":    "user_123",
        "sessionID": "sess_456",
    },
    Tags: []string{"production", "api-v2"},
}

result, _ := pipeline.Invoke(ctx, input, config)

// Every Runnable in the pipeline receives config automatically
// No need to pass it manually at each step!
```

Much cleaner! And infinitely extensible.

## Learning Objectives

By the end of this lesson, you will:

- ✅ Understand the Config pattern in Go
- ✅ Implement a callback system for monitoring
- ✅ Add metadata and tags for tracking
- ✅ Build configurable Runnables
- ✅ Create custom callbacks for logging and metrics
- ✅ Debug chains with visibility into each step
- ✅ Understand how Config complements context.Context

## Core Concepts

### What is Config?

Config is a struct that flows through your entire pipeline, carrying:

1. **Callbacks** - Hooks called at specific points (logging, metrics, debugging)
2. **Metadata** - Arbitrary data (user IDs, session info, request context)
3. **Tags** - Labels for filtering and organization
4. **MaxRetries** - Error handling configuration
5. **Timeout** - Operation timeout (in seconds)

### Go's Two Contexts

**context.Context (Go standard)**:
- Cancellation signals
- Deadlines
- Request-scoped values (sparingly)
- Propagates through call stack

**Config (Our pattern)**:
- Application configuration
- Callbacks and observability
- Business metadata
- More flexible than context.Context values

**Both work together**:
```go
func (r *Runnable) Invoke(
    ctx context.Context,  // Go's context for cancellation
    input interface{},
    config *Config,       // Our config for application data
) (interface{}, error)
```

### The Flow

```
User calls: runnable.Invoke(ctx, input, config)
                                        ↓
                        ┌───────────────┴─────────────┐
                        │                             │
                   Config passed to every step        │
                        │                             │
          ┌─────────────┼─────────┬──────────┬────────┼────┐
          │             │         │          │        │    │
        Step1        Step2      Step3      Step4    Step5  ...
          │             │         │          │        │
          └─────────────┴─────────┴──────────┴────────┘
                        │
                   All use same config
                   All trigger callbacks
                   All have access to metadata
```

## Implementation Deep Dive

### Step 1: The Config Struct

**Location:** `pkg/core/context.go`
```go
package core

import "context"

// Config holds configuration for Runnable execution
type Config struct {
    Callbacks  []Callback
    Tags       []string
    Metadata   map[string]interface{}
    MaxRetries int
    Timeout    int // seconds
}

// NewConfig creates a new Config with default values
func NewConfig() *Config {
    return &Config{
        Callbacks:  []Callback{},
        Tags:       []string{},
        Metadata:   make(map[string]interface{}),
        MaxRetries: 3,
        Timeout:    0, // no timeout
    }
}

// Builder pattern methods
func (c *Config) WithCallbacks(callbacks []Callback) *Config {
    c.Callbacks = callbacks
    return c
}

func (c *Config) WithTags(tags []string) *Config {
    c.Tags = tags
    return c
}

func (c *Config) WithMetadata(metadata map[string]interface{}) *Config {
    c.Metadata = metadata
    return c
}

func (c *Config) WithMaxRetries(maxRetries int) *Config {
    c.MaxRetries = maxRetries
    return c
}

func (c *Config) WithTimeout(timeout int) *Config {
    c.Timeout = timeout
    return c
}
```

**Why this design?**
- Simple struct, easy to extend
- Builder pattern for fluent API
- Sensible defaults
- No magic, just data

### Step 2: The Callback Interface

```go
// Callback interface for observability
type Callback interface {
    OnStart(ctx context.Context, runnable Runnable, input interface{}) error
    OnEnd(ctx context.Context, runnable Runnable, output interface{}) error
    OnError(ctx context.Context, runnable Runnable, err error) error
}

// CallbackManager manages multiple callbacks
type CallbackManager struct {
    callbacks []Callback
}

func NewCallbackManager(callbacks []Callback) *CallbackManager {
    if callbacks == nil {
        callbacks = []Callback{}
    }
    return &CallbackManager{
        callbacks: callbacks,
    }
}

func (cm *CallbackManager) HandleStart(ctx context.Context, runnable Runnable, input interface{}) error {
    for _, cb := range cm.callbacks {
        if err := cb.OnStart(ctx, runnable, input); err != nil {
            return err
        }
    }
    return nil
}

func (cm *CallbackManager) HandleEnd(ctx context.Context, runnable Runnable, output interface{}) error {
    for _, cb := range cm.callbacks {
        if err := cb.OnEnd(ctx, runnable, output); err != nil {
            return err
        }
    }
    return nil
}

func (cm *CallbackManager) HandleError(ctx context.Context, runnable Runnable, err error) error {
    for _, cb := range cm.callbacks {
        if cbErr := cb.OnError(ctx, runnable, err); cbErr != nil {
            return cbErr
        }
    }
    return nil
}
```

### Step 3: Example Callback - Logging

```go
// LoggingCallback logs all Runnable events
type LoggingCallback struct {
    Verbose bool
}

func (lc *LoggingCallback) OnStart(ctx context.Context, runnable Runnable, input interface{}) error {
    if lc.Verbose {
        fmt.Printf("[START] %s | Input: %v\n", runnable.Name(), input)
    }
    return nil
}

func (lc *LoggingCallback) OnEnd(ctx context.Context, runnable Runnable, output interface{}) error {
    if lc.Verbose {
        fmt.Printf("[END] %s | Output: %v\n", runnable.Name(), output)
    }
    return nil
}

func (lc *LoggingCallback) OnError(ctx context.Context, runnable Runnable, err error) error {
    fmt.Printf("[ERROR] %s | Error: %v\n", runnable.Name(), err)
    return nil
}
```

### Step 4: Integrating with Runnable

```go
// In BaseRunnable.Invoke:
func (r *BaseRunnable) Invoke(ctx context.Context, input interface{}, config *Config) (interface{}, error) {
    if config == nil {
        config = NewConfig()
    }
    
    // Create callback manager
    cm := NewCallbackManager(config.Callbacks)
    
    // Notify callbacks: starting
    if err := cm.HandleStart(ctx, r, input); err != nil {
        return nil, err
    }
    
    // Execute the runnable
    output, err := r.call(ctx, input, config)
    if err != nil {
        // Notify callbacks: error
        if cbErr := cm.HandleError(ctx, r, err); cbErr != nil {
            return nil, fmt.Errorf("callback error: %w, original error: %v", cbErr, err)
        }
        return nil, err
    }
    
    // Notify callbacks: success
    if err := cm.HandleEnd(ctx, r, output); err != nil {
        return nil, err
    }
    
    return output, nil
}
```

## Practical Examples

### Example 1: Basic Logging

```go
func withLogging() error {
    // Create callback
    logger := &LoggingCallback{Verbose: true}
    
    // Create config
    config := NewConfig().WithCallbacks([]Callback{logger})
    
    // Use with LLM
    llm, _ := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    defer llm.Close()
    
    ctx := context.Background()
    response, err := llm.Invoke(ctx, "Hello!", config)
    
    // Output:
    // [START] LlamaCppLLM | Input: Hello!
    // [END] LlamaCppLLM | Output: Hi there!
    
    return err
}
```

### Example 2: Metrics Tracking

```go
// MetricsCallback tracks performance
type MetricsCallback struct {
    StartTime time.Time
}

func (mc *MetricsCallback) OnStart(ctx context.Context, runnable Runnable, input interface{}) error {
    mc.StartTime = time.Now()
    return nil
}

func (mc *MetricsCallback) OnEnd(ctx context.Context, runnable Runnable, output interface{}) error {
    duration := time.Since(mc.StartTime)
    fmt.Printf("[METRICS] %s completed in %v\n", runnable.Name(), duration)
    return nil
}

func (mc *MetricsCallback) OnError(ctx context.Context, runnable Runnable, err error) error {
    duration := time.Since(mc.StartTime)
    fmt.Printf("[METRICS] %s failed after %v\n", runnable.Name(), duration)
    return nil
}

func withMetrics() error {
    metrics := &MetricsCallback{}
    config := NewConfig().WithCallbacks([]Callback{metrics})
    
    // Use with any Runnable
    llm, _ := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    defer llm.Close()
    
    ctx := context.Background()
    _, err := llm.Invoke(ctx, "Explain Go", config)
    
    // Output: [METRICS] LlamaCppLLM completed in 2.34s
    
    return err
}
```

### Example 3: User Context

```go
func withUserContext() error {
    config := NewConfig().
        WithMetadata(map[string]interface{}{
            "userID":    "user_123",
            "sessionID": "sess_456",
            "ip":        "192.168.1.1",
        }).
        WithTags([]string{"production", "api-v2"})
    
    // Later, callbacks can access this metadata
    // to add context to logs or metrics
    
    llm, _ := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    defer llm.Close()
    
    ctx := context.Background()
    _, err := llm.Invoke(ctx, "Hello", config)
    
    return err
}
```

### Example 4: Multiple Callbacks

```go
func multipleCallbacks() error {
    logger := &LoggingCallback{Verbose: true}
    metrics := &MetricsCallback{}
    
    // Chain multiple callbacks
    config := NewConfig().WithCallbacks([]Callback{logger, metrics})
    
    pipeline := step1.Pipe(step2).Pipe(step3)
    
    ctx := context.Background()
    result, err := pipeline.Invoke(ctx, input, config)
    
    // Both logger and metrics run at each step!
    
    return err
}
```

## Exercises

### Exercise 13: Simple Logger
Implement a callback that logs to a file.
- [Starter Code](exercises/13-simple-logger/starter.go)
- [Solution](exercises/13-simple-logger/solution.go)

### Exercise 14: Metrics with Metadata
Build a callback that tracks metrics and includes user context.
- [Starter Code](exercises/14-metrics-metadata/starter.go)
- [Solution](exercises/14-metrics-metadata/solution.go)

### Exercise 15: Config Inheritance
Create a system where child configs inherit parent settings.
- [Starter Code](exercises/15-config-inheritance/starter.go)
- [Solution](exercises/15-config-inheritance/solution.go)

### Exercise 16: Runtime Config
Build a Runnable that uses config to override runtime parameters.
- [Starter Code](exercises/16-runtime-config/starter.go)
- [Solution](exercises/16-runtime-config/solution.go)

## Key Takeaways

1. ✅ **Config pattern** - Passes application data through chains
2. ✅ **Complements context.Context** - Works alongside Go's context
3. ✅ **Callback interface** - OnStart, OnEnd, OnError hooks
4. ✅ **CallbackManager** - Manages multiple callbacks
5. ✅ **Metadata & Tags** - Arbitrary data and labels
6. ✅ **Builder pattern** - Fluent API for configuration
7. ✅ **Observability** - Logging, metrics, debugging built-in

## What's Next

**Foundation Complete!** 🎉

You've now mastered the four core patterns:
1. ✅ Runnable - Composability
2. ✅ Messages - Structured data
3. ✅ LLM Wrapper - Model integration
4. ✅ Config - Configuration and observability

**Next Section**: [02-composition](../../02-composition/) - Build prompt templates, parsers, and chains

**See it in action**: Check `pkg/core/context.go` for the full implementation

**Practice**: Complete all 4 exercises before moving on

## Further Reading

- [Context Package](https://pkg.go.dev/context)
- [Functional Options Pattern](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
- [Observer Pattern in Go](https://refactoring.guru/design-patterns/observer/go/example)
