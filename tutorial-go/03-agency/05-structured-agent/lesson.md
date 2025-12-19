# Structured Agent: JSON Mode for Reliability

**Part 3: Agency - Lesson 5**

> Use JSON output for 100% parseable tool calls.

## Overview

ReAct agents are powerful but have a problem: **parsing failures**. LLMs sometimes format outputs incorrectly, breaking the agent loop. **Structured Agents** solve this by enforcing JSON output.

```go
// ReAct Agent (text parsing - can fail):
response := "Thought: I should calculate\nActoin: calculator\n..." // Typo!
// Parser fails

// Structured Agent (JSON - always parseable):
response := `{
  "thought": "I should calculate",
  "action": "calculator",
  "action_input": {"operation": "add", "a": 5, "b": 3}
}`
// Always works!
```

## Why This Matters

### The Problem: Parsing Failures

ReAct agents fail when LLM output is malformed:

```
Common failures:
- "Acton:" instead of "Action:"
- Missing newlines
- Extra text before/after
- Inconsistent formatting
- Misspellings
```

### The Solution: JSON Schema

Force LLM to output valid JSON:

```go
type AgentStep struct {
    Thought     string                 `json:"thought"`
    Action      string                 `json:"action"`
    ActionInput map[string]interface{} `json:"action_input"`
    Done        bool                   `json:"done"`
    Answer      string                 `json:"answer,omitempty"`
}

// LLM must output this exact structure
```

## Core Concepts

### Structured Output

Instead of free text, LLM outputs JSON matching a schema:

**Prompt:**
```
Respond ONLY with valid JSON matching this schema:
{
  "thought": "your reasoning",
  "action": "tool_name or null",
  "action_input": {...} or null,
  "done": false,
  "answer": "final answer or null"
}
```

**LLM Output:**
```json
{
  "thought": "I need to calculate",
  "action": "calculator",
  "action_input": {"operation": "add", "a": 5, "b": 3},
  "done": false,
  "answer": null
}
```

### Benefits

1. **100% Parseable** - Always valid JSON
2. **Type Safe** - Unmarshal directly to structs
3. **Easier Debugging** - Clear structure
4. **Validation** - Check required fields
5. **Reliability** - No parsing failures

## Implementation

### Structured Agent

```go
type StructuredAgent struct {
    *core.BaseRunnable
    llm          core.Runnable
    executor     *ToolExecutor
    tools        []Tool
    maxIterations int
}

type AgentStep struct {
    Thought     string                 `json:"thought"`
    Action      string                 `json:"action"`
    ActionInput map[string]interface{} `json:"action_input"`
    Done        bool                   `json:"done"`
    Answer      string                 `json:"answer,omitempty"`
}

func NewStructuredAgent(llm core.Runnable, tools []Tool) *StructuredAgent {
    return &StructuredAgent{
        BaseRunnable: core.NewBaseRunnable("StructuredAgent"),
        llm:         llm,
        executor:    NewToolExecutor(tools, ExecutorConfig{}),
        tools:       tools,
        maxIterations: 10,
    }
}

func (a *StructuredAgent) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
    question := input.(string)
    history := []AgentStep{}
    
    for i := 0; i < a.maxIterations; i++ {
        // Build prompt
        prompt := a.buildJSONPrompt(question, history)
        
        // Get LLM response (JSON)
        response, err := a.llm.Invoke(ctx, prompt, config)
        if err != nil {
            return nil, err
        }
        
        // Parse JSON
        var step AgentStep
        if err := json.Unmarshal([]byte(response.(string)), &step); err != nil {
            return nil, fmt.Errorf("failed to parse JSON: %w", err)
        }
        
        // Check if done
        if step.Done {
            return step.Answer, nil
        }
        
        // Execute action
        toolCall := ToolCall{
            Name:      step.Action,
            Arguments: mustMarshal(step.ActionInput),
        }
        observation, err := a.executor.Execute(ctx, toolCall)
        if err != nil {
            observation = fmt.Sprintf("Error: %v", err)
        }
        
        // Add observation to step
        step.Answer = observation // Use Answer field for observation
        history = append(history, step)
    }
    
    return nil, fmt.Errorf("max iterations reached")
}

func (a *StructuredAgent) buildJSONPrompt(question string, history []AgentStep) string {
    toolsJSON := a.formatToolsJSON()
    historyJSON := mustMarshal(history)
    
    return fmt.Sprintf(`Answer this question using tools if needed.

Tools available:
%s

Question: %s

History:
%s

Respond with ONLY valid JSON matching this schema:
{
  "thought": "your reasoning about what to do next",
  "action": "tool_name or null if done",
  "action_input": {tool arguments} or null,
  "done": true if you have the final answer, false otherwise,
  "answer": "final answer to user or null"
}`, toolsJSON, question, historyJSON)
}
```

## Practical Examples

### Example 1: Structured Calculator

```go
func structuredCalculator() {
    llm, _ := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    defer llm.Close()
    
    tools := []Tool{NewCalculatorTool()}
    agent := NewStructuredAgent(llm, tools)
    
    ctx := context.Background()
    response, _ := agent.Invoke(ctx, "What is 456 * 789?", nil)
    
    fmt.Println("Answer:", response)
    
    // LLM outputs:
    // {"thought": "Calculate", "action": "calculator", ...}
    // Always parseable!
}
```

### Example 2: Compare Reliability

```go
func compareReliability() {
    // Run same task 100 times
    reactSuccess := 0
    structuredSuccess := 0
    
    for i := 0; i < 100; i++ {
        // ReAct agent
        _, err := reactAgent.Invoke(ctx, task, nil)
        if err == nil {
            reactSuccess++
        }
        
        // Structured agent
        _, err = structuredAgent.Invoke(ctx, task, nil)
        if err == nil {
            structuredSuccess++
        }
    }
    
    fmt.Printf("ReAct: %d%% success\n", reactSuccess)
    fmt.Printf("Structured: %d%% success\n", structuredSuccess)
    // Structured: 100% (always works)
}
```

## Exercises

### Exercise 53: JSON Schema Validation
Add schema validation before parsing.

### Exercise 54: Structured with Retry
Retry on JSON parse failures.

### Exercise 55: Complex Actions
Support multiple actions in one step.

### Exercise 56: Hybrid Agent
Combine ReAct and Structured approaches.

## Key Takeaways

1. ✅ **JSON Output** - Always parseable
2. ✅ **Type Safe** - Go structs
3. ✅ **Reliable** - No parsing failures
4. ✅ **Debuggable** - Clear structure
5. ✅ **Production Ready** - Use in real apps

## Congratulations!

You've completed **Part 3: Agency**! 🎉

You now know:
- ✅ Tools - Define functions with schemas
- ✅ Tool Executor - Safe execution
- ✅ Simple Agent - Single-step decisions
- ✅ ReAct Agent - Multi-step reasoning
- ✅ Structured Agent - JSON for reliability

## What's Next

**Part 4: Graphs** - Build complex workflows with state machines

Start with [Part 4 Overview](../../04-graphs/README.md) →

## Further Reading

- [ReAct Paper](https://arxiv.org/abs/2210.03629)
- [Function Calling](https://platform.openai.com/docs/guides/function-calling)
- [JSON Mode](https://platform.openai.com/docs/guides/text-generation/json-mode)
