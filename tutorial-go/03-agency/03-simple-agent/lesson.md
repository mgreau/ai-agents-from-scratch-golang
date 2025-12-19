# Simple Agent: Single-Step Decision Making

**Part 3: Agency - Lesson 3**

> The simplest agent: decide once, act once, respond.

## Overview

A **Simple Agent** is the most basic autonomous agent. It:
1. Receives a question/task
2. Decides which tool to use (or no tool)
3. Executes the tool once
4. Returns the final answer

```go
agent := NewSimpleAgent(llm, tools)
response, _ := agent.Invoke(ctx, "What's 15 * 23?", nil)
// Agent: "I'll use calculator" → Executes → "The answer is 345"
```

## Why This Matters

### When to Use Simple Agents

**Good for:**
- Questions needing one tool call
- Straightforward tasks
- Fast responses needed
- Predictable workflows

**Not good for:**
- Multi-step reasoning
- Complex problem-solving
- Tasks needing multiple tools

### The Simple Agent Pattern

```
User Question
    ↓
LLM: "Should I use a tool? If so, which one?"
    ↓
Tool Execution (if needed)
    ↓
LLM: "Here's the final answer"
    ↓
Response
```

## Implementation

### Simple Agent Structure

```go
type SimpleAgent struct {
    *core.BaseRunnable
    llm      core.Runnable
    executor *ToolExecutor
    tools    []Tool
}

func NewSimpleAgent(llm core.Runnable, tools []Tool) *SimpleAgent {
    return &SimpleAgent{
        BaseRunnable: core.NewBaseRunnable("SimpleAgent"),
        llm:         llm,
        executor:    NewToolExecutor(tools, ExecutorConfig{}),
        tools:       tools,
    }
}

func (a *SimpleAgent) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
    question := input.(string)
    
    // Step 1: Build prompt with tools
    prompt := a.buildPrompt(question)
    
    // Step 2: Get LLM decision
    response, err := a.llm.Invoke(ctx, prompt, config)
    if err != nil {
        return nil, err
    }
    
    // Step 3: Parse tool call (if any)
    toolCall, needsTool := a.parseToolCall(response.(string))
    
    // Step 4: Execute tool if needed
    var toolResult string
    if needsTool {
        toolResult, err = a.executor.Execute(ctx, toolCall)
        if err != nil {
            return nil, err
        }
        
        // Step 5: Get final answer with tool result
        finalPrompt := fmt.Sprintf("%s\nTool result: %s\nFinal answer:", prompt, toolResult)
        response, err = a.llm.Invoke(ctx, finalPrompt, config)
        if err != nil {
            return nil, err
        }
    }
    
    return response, nil
}

func (a *SimpleAgent) buildPrompt(question string) string {
    toolsDesc := a.formatTools()
    return fmt.Sprintf(`Answer the question. You can use tools if needed.

%s

Question: %s

If you need a tool, respond with:
TOOL: <tool_name>
ARGS: <json_arguments>

Otherwise, respond directly with the answer.`, toolsDesc, question)
}
```

## Practical Examples

### Example 1: Calculator Agent

```go
func calculatorAgent() {
    llm, _ := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    defer llm.Close()
    
    tools := []Tool{NewCalculatorTool()}
    agent := NewSimpleAgent(llm, tools)
    
    ctx := context.Background()
    
    // Test
    response, _ := agent.Invoke(ctx, "What is 456 * 789?", nil)
    fmt.Println("Agent:", response)
    // "The answer is 359,784"
}
```

### Example 2: Multi-Tool Agent

```go
func multiToolAgent() {
    llm, _ := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    defer llm.Close()
    
    tools := []Tool{
        NewCalculatorTool(),
        NewTimeTool(),
        NewWeatherTool(),
    }
    agent := NewSimpleAgent(llm, tools)
    
    ctx := context.Background()
    
    // Different questions
    fmt.Println(agent.Invoke(ctx, "What time is it in UTC?", nil))
    fmt.Println(agent.Invoke(ctx, "What's 100 / 4?", nil))
    fmt.Println(agent.Invoke(ctx, "What's the weather in Paris?", nil))
}
```

## Exercises

### Exercise 45: Simple Q&A Agent
Build an agent with search tool for questions.

### Exercise 46: File Operations Agent  
Create agent with file read/write tools.

### Exercise 47: Tool Selection Prompt
Optimize prompt for better tool selection.

### Exercise 48: Agent with Fallback
Add fallback when tool fails.

## Key Takeaways

1. ✅ **One decision** - Choose tool once
2. ✅ **Simple prompts** - Clear instructions
3. ✅ **Fast** - No multi-step reasoning
4. ✅ **Predictable** - Easy to debug
5. ✅ **Limited** - Can't solve complex problems

## What's Next

**Next Lesson**: [04-react-agent](../04-react-agent/lesson.md) - Multi-step reasoning loops

**See it in action**: Check `pkg/agents/simple.go`

**Practice**: Complete all 4 exercises
