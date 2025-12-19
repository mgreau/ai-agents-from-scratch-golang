# ReAct Agent: Reasoning + Acting

**Part 3: Agency - Lesson 4**

> Multi-step agents that think, act, observe, and repeat.

## Overview

**ReAct (Reasoning + Acting)** is a powerful agent pattern that alternates between:
- **Thought** - Reasoning about what to do next
- **Action** - Executing a tool
- **Observation** - Seeing the result

This loop continues until the agent has enough information to answer.

```go
// User: "What's the weather in Paris and convert temp to Fahrenheit?"

// Agent Loop:
// Thought: I need weather data for Paris
// Action: get_weather("Paris")
// Observation: {"temperature": 20, "unit": "celsius"}

// Thought: Now I need to convert 20°C to Fahrenheit
// Action: calculator("20 * 9/5 + 32")
// Observation: 68

// Thought: I have all the information
// Answer: It's 68°F in Paris
```

## Why This Matters

### The Problem: Complex Tasks Need Multiple Steps

Simple agents fail on multi-step tasks:

```go
simpleAgent.Invoke(ctx, 
    "Search for Go programming and summarize the top result", nil)
// Can only do one: search OR summarize, not both
```

### The Solution: ReAct Loop

```go
reactAgent.Invoke(ctx, 
    "Search for Go programming and summarize the top result", nil)

// Thought: Need to search first
// Action: search("Go programming")
// Observation: [search results]

// Thought: Now summarize the top result
// Action: summarize([top result])
// Observation: [summary]

// Thought: I have the answer
// Answer: [final response]
```

## Core Concepts

### ReAct Loop

```
┌─────────────────────────────────────┐
│ User Question                       │
└─────────────┬───────────────────────┘
              │
              ↓
        ┌─────────┐
        │ Thought │ ← "What should I do?"
        └────┬────┘
             │
             ↓
        ┌────────┐
        │ Action │ ← Execute tool
        └────┬───┘
             │
             ↓
      ┌──────────────┐
      │ Observation  │ ← Tool result
      └──────┬───────┘
             │
             ↓
    ┌────────────────┐
    │ Done?          │
    └───┬────────┬───┘
        │ No     │ Yes
        │        └──────→ Answer
        │
        └───→ Back to Thought
```

### Prompt Format

```
Question: {user_question}

You can use these tools:
- calculator: ...
- search: ...

Think step by step using this format:

Thought: [your reasoning]
Action: [tool_name]
Action Input: [json_args]
Observation: [tool result - will be filled automatically]

... repeat Thought/Action/Observation as needed ...

Thought: I have enough information
Answer: [final answer to user]
```

## Implementation in Go

### ReAct Agent

**Location:** `pkg/agents/react.go`

```go
type ReActAgent struct {
    *core.BaseRunnable
    llm          core.Runnable
    executor     *ToolExecutor
    tools        []Tool
    maxIterations int
}

func NewReActAgent(llm core.Runnable, tools []Tool, maxIter int) *ReActAgent {
    if maxIter == 0 {
        maxIter = 10
    }
    return &ReActAgent{
        BaseRunnable: core.NewBaseRunnable("ReActAgent"),
        llm:         llm,
        executor:    NewToolExecutor(tools, ExecutorConfig{}),
        tools:       tools,
        maxIterations: maxIter,
    }
}

func (a *ReActAgent) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
    question := input.(string)
    scratchpad := ""
    
    for i := 0; i < a.maxIterations; i++ {
        // Build prompt with history
        prompt := a.buildPrompt(question, scratchpad)
        
        // Get LLM response
        response, err := a.llm.Invoke(ctx, prompt, config)
        if err != nil {
            return nil, err
        }
        
        responseStr := response.(string)
        
        // Check if agent is done
        if strings.Contains(responseStr, "Answer:") {
            answer := a.extractAnswer(responseStr)
            return answer, nil
        }
        
        // Parse thought and action
        thought, action, actionInput := a.parseStep(responseStr)
        
        // Execute action
        toolCall := ToolCall{
            Name:      action,
            Arguments: actionInput,
        }
        observation, err := a.executor.Execute(ctx, toolCall)
        if err != nil {
            observation = fmt.Sprintf("Error: %v", err)
        }
        
        // Add to scratchpad
        scratchpad += fmt.Sprintf("Thought: %s\nAction: %s\nAction Input: %s\nObservation: %s\n\n",
            thought, action, actionInput, observation)
    }
    
    return nil, fmt.Errorf("max iterations reached")
}

func (a *ReActAgent) parseStep(response string) (thought, action, actionInput string) {
    // Extract Thought:
    if idx := strings.Index(response, "Thought:"); idx != -1 {
        thought = extractUntil(response[idx+8:], "\n")
    }
    
    // Extract Action:
    if idx := strings.Index(response, "Action:"); idx != -1 {
        action = strings.TrimSpace(extractUntil(response[idx+7:], "\n"))
    }
    
    // Extract Action Input:
    if idx := strings.Index(response, "Action Input:"); idx != -1 {
        actionInput = strings.TrimSpace(extractUntil(response[idx+13:], "\n"))
    }
    
    return
}
```

## Practical Examples

### Example 1: Multi-Step Calculation

```go
func multiStepCalc() {
    llm, _ := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    defer llm.Close()
    
    tools := []Tool{NewCalculatorTool()}
    agent := NewReActAgent(llm, tools, 10)
    
    ctx := context.Background()
    
    response, _ := agent.Invoke(ctx, 
        "What is (15 * 23) + (45 / 3)?", nil)
    
    fmt.Println("Agent:", response)
    
    // Agent process:
    // Thought: Calculate 15 * 23 first
    // Action: calculator
    // Observation: 345
    //
    // Thought: Now calculate 45 / 3
    // Action: calculator
    // Observation: 15
    //
    // Thought: Add them together
    // Action: calculator
    // Observation: 360
    //
    // Answer: 360
}
```

### Example 2: Research Task

```go
func researchTask() {
    tools := []Tool{
        NewSearchTool(),
        NewSummarizeTool(),
    }
    agent := NewReActAgent(llm, tools, 15)
    
    response, _ := agent.Invoke(ctx, 
        "Find information about Go concurrency and summarize", nil)
    
    // Agent process:
    // Thought: Search for Go concurrency
    // Action: search
    // Observation: [search results]
    //
    // Thought: Summarize the results
    // Action: summarize
    // Observation: [summary]
    //
    // Answer: [final summary]
}
```

## Exercises

### Exercise 49: ReAct with Memory
Add conversation memory to ReAct agent.

### Exercise 50: Streaming ReAct
Stream thoughts and actions in real-time.

### Exercise 51: ReAct with Validation
Validate tool calls before execution.

### Exercise 52: Multi-Tool Planning
Agent that uses 3+ tools in sequence.

## Key Takeaways

1. ✅ **Multi-step** - Loop until done
2. ✅ **Reasoning** - Explicit thoughts
3. ✅ **Observable** - See agent's process
4. ✅ **Flexible** - Adapts to task complexity
5. ✅ **Max iterations** - Prevent infinite loops

## What's Next

**Next Lesson**: [05-structured-agent](../05-structured-agent/lesson.md) - JSON mode for reliability

**See it in action**: Check `pkg/agents/react.go` (already implemented!)

**Practice**: Complete all 4 exercises
