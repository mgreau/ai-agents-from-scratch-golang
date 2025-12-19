# Tools: Functions for Agents

**Part 3: Agency - Lesson 1**

> Giving agents superpowers through external functions.

## Overview

Agents are only as powerful as the tools they can use. **Tools** are functions that extend an agent's capabilities beyond text generation - they can fetch data, perform calculations, interact with APIs, and manipulate the real world.

```go
// Without tools: Limited to what LLM knows
agent.Ask("What's the weather right now?")
// "I don't have access to real-time weather data"

// With tools: Can fetch live data
agent.WithTools(weatherTool).Ask("What's the weather right now?")
// Calls get_weather() → "It's 72°F and sunny"
```

## Why This Matters

### The Problem: LLMs Are Frozen in Time

LLMs only know what they were trained on:

```go
// LLM limitations
llm.Invoke(ctx, "What time is it?", nil)
// "I'm an AI and don't have access to current time"

llm.Invoke(ctx, "What's 987654 * 123456?", nil)
// May give wrong answer or refuse

llm.Invoke(ctx, "Send an email to john@example.com", nil)
// "I cannot send emails"
```

### The Solution: Tools

```go
tools := []Tool{
    NewTimeTool(),
    NewCalculatorTool(),
    NewEmailTool(),
}

agent := NewAgent(llm, tools)

agent.Invoke(ctx, "What time is it?", nil)
// Calls get_time() → "14:35:22 UTC"

agent.Invoke(ctx, "What's 987654 * 123456?", nil)
// Calls calculator() → "121,932,469,024"

agent.Invoke(ctx, "Send email to john@example.com saying hello", nil)
// Calls send_email() → "Email sent successfully"
```

## Learning Objectives

By the end of this lesson, you will:

- ✅ Understand tool anatomy (name, description, schema, function)
- ✅ Create tools with input/output schemas
- ✅ Write tool descriptions for LLMs
- ✅ Implement tool functions in Go
- ✅ Build a tool registry
- ✅ Handle tool errors gracefully
- ✅ Test tools independently

## Core Concepts

### What is a Tool?

A tool is a **structured function** that an agent can call. It has:

1. **Name** - Unique identifier (e.g., "calculator")
2. **Description** - What it does (for LLM understanding)
3. **Input Schema** - Expected parameters and types
4. **Function** - The actual Go function to execute
5. **Output Format** - What it returns

**Structure:**
```
Tool: "calculator"
├─ Description: "Performs basic math operations"
├─ Input Schema: {"operation": "add", "a": 5, "b": 3}
├─ Function: func(input) → result
└─ Output: "8"
```

### Tool Anatomy in Go

```go
type Tool struct {
    Name        string
    Description string
    InputSchema interface{} // Struct defining expected input
    Function    ToolFunc
}

type ToolFunc func(input interface{}) (string, error)
```

### Tool Call Flow

```
1. Agent decides to use tool
2. Agent generates tool call:
   {
       "name": "calculator",
       "arguments": {"operation": "add", "a": 5, "b": 3}
   }
3. Tool executor parses call
4. Tool executor validates arguments
5. Tool function executes
6. Result returned to agent
```

## Implementation in Go

### Step 1: Tool Interface

**Location:** `pkg/tools/base.go`

```go
package tools

import (
    "context"
    "encoding/json"
    "fmt"
)

// Tool represents a function an agent can call
type Tool interface {
    Name() string
    Description() string
    InputSchema() interface{}
    Execute(ctx context.Context, input interface{}) (string, error)
}

// BaseTool provides common functionality
type BaseTool struct {
    name        string
    description string
    schema      interface{}
    fn          ToolFunc
}

type ToolFunc func(ctx context.Context, input interface{}) (string, error)

func NewBaseTool(name, description string, schema interface{}, fn ToolFunc) *BaseTool {
    return &BaseTool{
        name:        name,
        description: description,
        schema:      schema,
        fn:          fn,
    }
}

func (t *BaseTool) Name() string {
    return t.name
}

func (t *BaseTool) Description() string {
    return t.description
}

func (t *BaseTool) InputSchema() interface{} {
    return t.schema
}

func (t *BaseTool) Execute(ctx context.Context, input interface{}) (string, error) {
    return t.fn(ctx, input)
}
```

### Step 2: Example Tool - Calculator

```go
// CalculatorInput defines the schema
type CalculatorInput struct {
    Operation string  `json:"operation"` // add, subtract, multiply, divide
    A         float64 `json:"a"`
    B         float64 `json:"b"`
}

// NewCalculatorTool creates a calculator tool
func NewCalculatorTool() Tool {
    return NewBaseTool(
        "calculator",
        "Performs basic arithmetic operations (add, subtract, multiply, divide). "+
            "Use this when you need to perform calculations.",
        CalculatorInput{},
        calculateFunc,
    )
}

// calculateFunc implements the calculator logic
func calculateFunc(ctx context.Context, input interface{}) (string, error) {
    // Parse input
    inputJSON, err := json.Marshal(input)
    if err != nil {
        return "", fmt.Errorf("failed to marshal input: %w", err)
    }
    
    var calcInput CalculatorInput
    if err := json.Unmarshal(inputJSON, &calcInput); err != nil {
        return "", fmt.Errorf("invalid input format: %w", err)
    }
    
    // Validate operation
    var result float64
    switch calcInput.Operation {
    case "add":
        result = calcInput.A + calcInput.B
    case "subtract":
        result = calcInput.A - calcInput.B
    case "multiply":
        result = calcInput.A * calcInput.B
    case "divide":
        if calcInput.B == 0 {
            return "", fmt.Errorf("division by zero")
        }
        result = calcInput.A / calcInput.B
    default:
        return "", fmt.Errorf("unknown operation: %s", calcInput.Operation)
    }
    
    return fmt.Sprintf("%.2f", result), nil
}
```

### Step 3: More Example Tools

**Time Tool:**
```go
type TimeInput struct {
    Timezone string `json:"timezone"` // e.g., "America/New_York", "UTC"
}

func NewTimeTool() Tool {
    return NewBaseTool(
        "get_time",
        "Returns the current time in a specified timezone. "+
            "Use this when users ask 'what time is it' or need current time.",
        TimeInput{},
        func(ctx context.Context, input interface{}) (string, error) {
            inputJSON, _ := json.Marshal(input)
            var timeInput TimeInput
            json.Unmarshal(inputJSON, &timeInput)
            
            loc, err := time.LoadLocation(timeInput.Timezone)
            if err != nil {
                loc = time.UTC
            }
            
            now := time.Now().In(loc)
            return now.Format("15:04:05 MST"), nil
        },
    )
}
```

**Weather Tool (mock):**
```go
type WeatherInput struct {
    Location string `json:"location"` // City name
}

func NewWeatherTool() Tool {
    return NewBaseTool(
        "get_weather",
        "Returns current weather for a location. "+
            "Use this when users ask about weather conditions.",
        WeatherInput{},
        func(ctx context.Context, input interface{}) (string, error) {
            inputJSON, _ := json.Marshal(input)
            var weatherInput WeatherInput
            json.Unmarshal(inputJSON, &weatherInput)
            
            // Mock response (in real app, call weather API)
            weather := map[string]interface{}{
                "location":    weatherInput.Location,
                "temperature": 72,
                "condition":   "sunny",
                "humidity":    "45%",
            }
            
            result, _ := json.Marshal(weather)
            return string(result), nil
        },
    )
}
```

### Step 4: Tool Registry

```go
// ToolRegistry manages available tools
type ToolRegistry struct {
    tools map[string]Tool
}

func NewToolRegistry() *ToolRegistry {
    return &ToolRegistry{
        tools: make(map[string]Tool),
    }
}

func (r *ToolRegistry) Register(tool Tool) {
    r.tools[tool.Name()] = tool
}

func (r *ToolRegistry) Get(name string) (Tool, bool) {
    tool, exists := r.tools[name]
    return tool, exists
}

func (r *ToolRegistry) All() []Tool {
    tools := make([]Tool, 0, len(r.tools))
    for _, tool := range r.tools {
        tools = append(tools, tool)
    }
    return tools
}

// ToPromptString formats tools for LLM prompts
func (r *ToolRegistry) ToPromptString() string {
    var builder strings.Builder
    builder.WriteString("You have access to the following tools:\n\n")
    
    for _, tool := range r.tools {
        builder.WriteString(fmt.Sprintf("- %s: %s\n", tool.Name(), tool.Description()))
        
        // Include schema information
        schemaJSON, _ := json.MarshalIndent(tool.InputSchema(), "  ", "  ")
        builder.WriteString(fmt.Sprintf("  Input schema: %s\n\n", schemaJSON))
    }
    
    return builder.String()
}
```

## Practical Examples

### Example 1: Create and Test Calculator

```go
func testCalculator() {
    calc := NewCalculatorTool()
    
    ctx := context.Background()
    
    // Test addition
    result, _ := calc.Execute(ctx, map[string]interface{}{
        "operation": "add",
        "a":         15.5,
        "b":         23.7,
    })
    fmt.Println("15.5 + 23.7 =", result) // 39.20
    
    // Test division
    result, _ = calc.Execute(ctx, map[string]interface{}{
        "operation": "divide",
        "a":         100.0,
        "b":         4.0,
    })
    fmt.Println("100 / 4 =", result) // 25.00
    
    // Test error handling
    _, err := calc.Execute(ctx, map[string]interface{}{
        "operation": "divide",
        "a":         10.0,
        "b":         0.0,
    })
    fmt.Println("Error:", err) // division by zero
}
```

### Example 2: Build Tool Registry

```go
func buildRegistry() {
    registry := NewToolRegistry()
    
    // Register tools
    registry.Register(NewCalculatorTool())
    registry.Register(NewTimeTool())
    registry.Register(NewWeatherTool())
    
    // List all tools
    fmt.Println("Available tools:")
    for _, tool := range registry.All() {
        fmt.Printf("- %s: %s\n", tool.Name(), tool.Description())
    }
    
    // Get specific tool
    calc, exists := registry.Get("calculator")
    if exists {
        result, _ := calc.Execute(context.Background(), map[string]interface{}{
            "operation": "multiply",
            "a":         7.0,
            "b":         8.0,
        })
        fmt.Println("7 * 8 =", result)
    }
}
```

### Example 3: Format Tools for LLM

```go
func formatForLLM() {
    registry := NewToolRegistry()
    registry.Register(NewCalculatorTool())
    registry.Register(NewWeatherTool())
    
    prompt := registry.ToPromptString()
    fmt.Println(prompt)
    
    // Output:
    // You have access to the following tools:
    //
    // - calculator: Performs basic arithmetic operations...
    //   Input schema: {
    //     "operation": "string",
    //     "a": "float64",
    //     "b": "float64"
    //   }
    //
    // - get_weather: Returns current weather...
    //   Input schema: {
    //     "location": "string"
    //   }
}
```

### Example 4: Custom Tool

```go
// Create a custom search tool
type SearchInput struct {
    Query string `json:"query"`
    NumResults int `json:"num_results"`
}

func NewSearchTool() Tool {
    return NewBaseTool(
        "search",
        "Search the web for information. Use when you need to find current information.",
        SearchInput{},
        func(ctx context.Context, input interface{}) (string, error) {
            inputJSON, _ := json.Marshal(input)
            var searchInput SearchInput
            json.Unmarshal(inputJSON, &searchInput)
            
            // Mock search (in real app, call search API)
            results := []string{
                "Result 1 about " + searchInput.Query,
                "Result 2 about " + searchInput.Query,
            }
            
            return strings.Join(results, "\n"), nil
        },
    )
}
```

## Exercises

### Exercise 37: File Reader Tool
Create a tool that reads file contents safely.
- [Starter Code](exercises/37-file-reader/starter.go)
- [Solution](exercises/37-file-reader/solution.go)

### Exercise 38: HTTP Request Tool
Build a tool that makes HTTP GET requests.
- [Starter Code](exercises/38-http-request/starter.go)
- [Solution](exercises/38-http-request/solution.go)

### Exercise 39: Database Query Tool
Implement a tool that queries a database.
- [Starter Code](exercises/39-database-query/starter.go)
- [Solution](exercises/39-database-query/solution.go)

### Exercise 40: Tool with Validation
Create a tool with complex input validation.
- [Starter Code](exercises/40-tool-validation/starter.go)
- [Solution](exercises/40-tool-validation/solution.go)

## Key Takeaways

1. ✅ **Tool = Name + Description + Schema + Function**
2. ✅ **Schemas** - Define expected inputs with Go structs
3. ✅ **Descriptions** - Critical for LLM understanding
4. ✅ **Tool Registry** - Manage available tools
5. ✅ **Error Handling** - Always return meaningful errors
6. ✅ **Testing** - Test tools independently before agents
7. ✅ **Type Safety** - Go's type system validates inputs

## What's Next

**Next Lesson**: [02-tool-executor](../02-tool-executor/lesson.md) - Safely execute tools with validation

**See it in action**: Check `pkg/tools/base.go` for implementation

**Practice**: Complete all 4 exercises before moving on

## Further Reading

- [Function Calling in LLMs](https://platform.openai.com/docs/guides/function-calling)
- [JSON Schema](https://json-schema.org/)
- [LangChain Tools](https://python.langchain.com/docs/modules/agents/tools/)
