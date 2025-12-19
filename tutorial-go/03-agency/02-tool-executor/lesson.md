# Tool Executor: Safe Tool Execution

**Part 3: Agency - Lesson 2**

> Parse, validate, and execute tool calls safely.

## Overview

In Lesson 1, you defined tools. Now you need to **execute** them based on LLM outputs. The Tool Executor is the bridge between the LLM's tool calls and actual function execution.

```go
// LLM outputs this:
toolCall := ToolCall{
    Name: "calculator",
    Arguments: `{"operation": "add", "a": 15, "b": 23}`,
}

// Tool Executor:
executor := NewToolExecutor(tools)
result, _ := executor.Execute(ctx, toolCall)
// Returns: "38.00"
```

## Why This Matters

### The Problem: Unsafe Execution

Without proper execution:

```go
// Dangerous: No validation
funcName := llmOutput["function"]
args := llmOutput["arguments"]
result := callFunction(funcName, args) // What if funcName is malicious?
```

Problems:
- No input validation
- No error handling
- No timeout protection
- No logging/observability
- Security risks

### The Solution: Tool Executor

```go
executor := NewToolExecutor(tools, ExecutorConfig{
    Timeout:     5 * time.Second,
    MaxRetries:  3,
    ValidateInputs: true,
})

result, err := executor.Execute(ctx, toolCall)
// Validates, times out, logs, handles errors
```

## Core Concepts

### Tool Call Structure

```go
type ToolCall struct {
    ID        string `json:"id"`        // Unique call ID
    Name      string `json:"name"`      // Tool to call
    Arguments string `json:"arguments"` // JSON string of args
}
```

### Execution Flow

```
1. Parse ToolCall from LLM output
2. Look up tool in registry
3. Validate tool exists
4. Parse arguments JSON
5. Validate against schema
6. Execute with timeout
7. Handle errors
8. Return result or error
```

## Implementation in Go

### Tool Executor

**Location:** `pkg/tools/executor.go`

```go
package tools

import (
    "context"
    "encoding/json"
    "fmt"
    "time"
)

// ToolExecutor executes tools safely
type ToolExecutor struct {
    registry *ToolRegistry
    timeout  time.Duration
}

// ExecutorConfig holds configuration
type ExecutorConfig struct {
    Timeout time.Duration
}

func NewToolExecutor(tools []Tool, config ExecutorConfig) *ToolExecutor {
    if config.Timeout == 0 {
        config.Timeout = 30 * time.Second
    }
    
    registry := NewToolRegistry()
    for _, tool := range tools {
        registry.Register(tool)
    }
    
    return &ToolExecutor{
        registry: registry,
        timeout:  config.Timeout,
    }
}

// Execute runs a tool call
func (e *ToolExecutor) Execute(ctx context.Context, call ToolCall) (string, error) {
    // 1. Get tool
    tool, exists := e.registry.Get(call.Name)
    if !exists {
        return "", fmt.Errorf("tool not found: %s", call.Name)
    }
    
    // 2. Parse arguments
    var args map[string]interface{}
    if err := json.Unmarshal([]byte(call.Arguments), &args); err != nil {
        return "", fmt.Errorf("invalid arguments JSON: %w", err)
    }
    
    // 3. Execute with timeout
    ctx, cancel := context.WithTimeout(ctx, e.timeout)
    defer cancel()
    
    resultChan := make(chan executionResult, 1)
    go func() {
        result, err := tool.Execute(ctx, args)
        resultChan <- executionResult{result, err}
    }()
    
    select {
    case res := <-resultChan:
        return res.result, res.err
    case <-ctx.Done():
        return "", fmt.Errorf("tool execution timeout: %s", call.Name)
    }
}

type executionResult struct {
    result string
    err    error
}

// ExecuteBatch runs multiple tool calls concurrently
func (e *ToolExecutor) ExecuteBatch(ctx context.Context, calls []ToolCall) ([]string, error) {
    results := make([]string, len(calls))
    errors := make([]error, len(calls))
    
    var wg sync.WaitGroup
    for i, call := range calls {
        wg.Add(1)
        go func(idx int, tc ToolCall) {
            defer wg.Done()
            result, err := e.Execute(ctx, tc)
            results[idx] = result
            errors[idx] = err
        }(i, call)
    }
    
    wg.Wait()
    
    // Check for errors
    for _, err := range errors {
        if err != nil {
            return results, fmt.Errorf("batch execution failed: %w", err)
        }
    }
    
    return results, nil
}
```

## Practical Examples

### Example 1: Basic Execution

```go
func basicExecution() {
    tools := []Tool{
        NewCalculatorTool(),
        NewTimeTool(),
    }
    
    executor := NewToolExecutor(tools, ExecutorConfig{
        Timeout: 5 * time.Second,
    })
    
    ctx := context.Background()
    
    // Execute calculator
    call := ToolCall{
        ID:   "call_123",
        Name: "calculator",
        Arguments: `{"operation": "multiply", "a": 7, "b": 8}`,
    }
    
    result, err := executor.Execute(ctx, call)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Println("Result:", result) // 56.00
}
```

### Example 2: Error Handling

```go
func errorHandling() {
    executor := NewToolExecutor(tools, ExecutorConfig{})
    ctx := context.Background()
    
    // Unknown tool
    call1 := ToolCall{Name: "unknown_tool", Arguments: "{}"}
    _, err := executor.Execute(ctx, call1)
    fmt.Println("Error:", err) // tool not found: unknown_tool
    
    // Invalid JSON
    call2 := ToolCall{Name: "calculator", Arguments: "{invalid}"}
    _, err = executor.Execute(ctx, call2)
    fmt.Println("Error:", err) // invalid arguments JSON
    
    // Division by zero
    call3 := ToolCall{
        Name: "calculator",
        Arguments: `{"operation": "divide", "a": 10, "b": 0}`,
    }
    _, err = executor.Execute(ctx, call3)
    fmt.Println("Error:", err) // division by zero
}
```

### Example 3: Batch Execution

```go
func batchExecution() {
    executor := NewToolExecutor(tools, ExecutorConfig{})
    ctx := context.Background()
    
    calls := []ToolCall{
        {Name: "calculator", Arguments: `{"operation": "add", "a": 5, "b": 3}`},
        {Name: "calculator", Arguments: `{"operation": "multiply", "a": 4, "b": 7}`},
        {Name: "get_time", Arguments: `{"timezone": "UTC"}`},
    }
    
    results, err := executor.ExecuteBatch(ctx, calls)
    if err != nil {
        fmt.Printf("Batch error: %v\n", err)
        return
    }
    
    for i, result := range results {
        fmt.Printf("Call %d: %s\n", i+1, result)
    }
}
```

## Exercises

### Exercise 41: Retry Logic
Add automatic retry with exponential backoff.

### Exercise 42: Tool Logging
Implement comprehensive logging for all tool executions.

### Exercise 43: Input Validation
Add schema-based input validation before execution.

### Exercise 44: Tool Caching
Cache tool results to avoid redundant executions.

## Key Takeaways

1. ✅ **Safe Execution** - Validate inputs, handle errors
2. ✅ **Timeouts** - Prevent hanging tools
3. ✅ **Error Messages** - Clear, actionable errors
4. ✅ **Batch Processing** - Execute multiple tools concurrently
5. ✅ **Goroutines** - Async execution with channels

## What's Next

**Next Lesson**: [03-simple-agent](../03-simple-agent/lesson.md) - Build single-step agents

**See it in action**: Check `pkg/tools/executor.go`

**Practice**: Complete all 4 exercises
