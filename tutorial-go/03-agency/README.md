# Part 3: Agency

> **Building autonomous agents that use tools and make decisions**

Welcome to Part 3! You've mastered Foundation (Runnables, Messages, LLMs, Config) and Composition (Prompts, Parsers, Chains, Piping, Memory). Now you'll learn to build **autonomous agents** that can use tools, reason about problems, and make decisions.

## What is Agency?

Agency is the ability to **act autonomously** to achieve goals. An agent:
1. Receives a task or question
2. Decides what tools to use
3. Executes tools to gather information
4. Reasons about results
5. Repeats until task is complete

**Example:**
```go
agent := NewReActAgent(llm, tools)

// User: "What's the weather in Paris and what time is it there?"

// Agent thinks:
// 1. "I need weather data" → Calls get_weather("Paris")
// 2. "I need current time" → Calls get_time("Europe/Paris")
// 3. "I have all info" → Responds to user
```

## Why Agency Matters

### Without Agents: Limited AI

```go
// Hardcoded logic - inflexible
func handleRequest(question string) string {
    if strings.Contains(question, "weather") {
        return getWeather("New York")
    } else if strings.Contains(question, "time") {
        return getTime("UTC")
    }
    return "I don't know"
}
```

Problems:
- Can't handle unexpected requests
- No reasoning or planning
- Hardcoded tool selection
- Can't combine tools
- No learning or adaptation

### With Agents: Autonomous AI

```go
// Agent decides what to do
agent := NewReActAgent(llm, []Tool{
    weatherTool,
    timeTool,
    calculatorTool,
    searchTool,
})

result, _ := agent.Invoke(ctx, "What's the weather in Paris?", nil)
// Agent autonomously:
// - Chooses weather tool
// - Extracts "Paris" as location
// - Calls tool with correct args
// - Formats response
```

Benefits:
- ✅ Handles any question
- ✅ Reasons about tool usage
- ✅ Combines multiple tools
- ✅ Adapts to context
- ✅ Explains reasoning

## What You'll Learn

This section teaches five key patterns for building agents:

### 1. **Tools** (Lesson 1)
Define functions that agents can call.

**Skills:**
- Create tool schemas
- Define input/output types
- Add descriptions for LLMs
- Register tools in registry
- Handle tool errors

**Example:**
```go
tool := Tool{
    Name: "calculator",
    Description: "Performs basic math operations",
    InputSchema: CalculatorInput{},
    Function: func(input interface{}) (string, error) {
        // Implementation
    },
}
```

### 2. **Tool Executor** (Lesson 2)
Safely execute tools with validation and error handling.

**Skills:**
- Parse tool calls from LLM
- Validate inputs against schemas
- Execute with timeouts
- Handle errors gracefully
- Log tool execution

**Example:**
```go
executor := NewToolExecutor(tools)
result, err := executor.Execute(ctx, toolCall)
```

### 3. **Simple Agent** (Lesson 3)
Single-step agents that choose and execute one tool.

**Skills:**
- Prompt engineering for tool selection
- Parse tool choice from LLM
- Execute single tool
- Format final response
- Handle "no tool needed" cases

**Example:**
```go
agent := NewSimpleAgent(llm, tools)
response, _ := agent.Invoke(ctx, "What's 15 * 23?", nil)
// Agent: Thinks → calculator → 345
```

### 4. **ReAct Agent** (Lesson 4)
Multi-step agents that Reason and Act in a loop.

**Skills:**
- Implement ReAct pattern
- Parse Thought/Action/Observation
- Multi-step reasoning loops
- Know when to stop
- Handle max iterations

**Example:**
```go
agent := NewReActAgent(llm, tools)
response, _ := agent.Invoke(ctx, 
    "What's the weather in Paris and convert temp to Fahrenheit?", nil)
// Agent:
// Thought: Need weather data
// Action: get_weather("Paris")
// Observation: 20°C
// Thought: Need to convert
// Action: calculator("20 * 9/5 + 32")
// Observation: 68°F
// Thought: I have the answer
// Answer: 68°F
```

### 5. **Structured Agent** (Lesson 5)
Agents using JSON mode for reliable tool calling.

**Skills:**
- JSON mode for tool selection
- Schema-driven tool calls
- Validate JSON output
- Handle malformed responses
- Improve reliability

**Example:**
```go
agent := NewStructuredAgent(llm, tools)
// Uses JSON output for 100% parseable tool calls
```

## The Agency Patterns

These five lessons build on each other:

```
Tools → Tool Executor → Simple Agent
                           ↓
                     ReAct Agent
                           ↓
                  Structured Agent (most reliable)
```

## Go Advantages for Agents

**Type Safety:**
```go
// Tool inputs are type-checked
type CalculatorInput struct {
    Operation string  `json:"operation"`
    A         float64 `json:"a"`
    B         float64 `json:"b"`
}
```

**Goroutines for Timeouts:**
```go
// Execute tools with timeout
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
result, _ := tool.Execute(ctx, input)
```

**Error Handling:**
```go
// Explicit error handling at each step
if err := agent.executeAction(action); err != nil {
    return agent.handleToolError(err)
}
```

**Channels for Streaming:**
```go
// Stream agent thoughts in real-time
thoughts := agent.StreamThoughts(ctx, input)
for thought := range thoughts {
    fmt.Println("Agent thinking:", thought)
}
```

## Prerequisites

Before starting, ensure you understand:
- ✅ Runnable interface and Pipe (Part 1, Lesson 1)
- ✅ Message types (Part 1, Lesson 2)
- ✅ LLM wrappers (Part 1, Lesson 3)
- ✅ Prompts (Part 2, Lesson 1)
- ✅ Parsers (Part 2, Lesson 2)
- ✅ LLM Chains (Part 2, Lesson 3)

## Lessons

### [Lesson 1: Tools](01-tools/lesson.md)
Define functions with schemas that agents can call.

### [Lesson 2: Tool Executor](02-tool-executor/lesson.md)
Safely execute tools with validation and error handling.

### [Lesson 3: Simple Agent](03-simple-agent/lesson.md)
Build single-step agents that choose and execute one tool.

### [Lesson 4: ReAct Agent](04-react-agent/lesson.md)
Implement multi-step reasoning and acting loops.

### [Lesson 5: Structured Agent](05-structured-agent/lesson.md)
Use JSON mode for reliable tool calling.

## Exercises

Each lesson includes 4 hands-on exercises (20 total).

**Practice path:**
1. Complete all 5 lessons in order
2. Do exercises for each lesson before moving on
3. Build final project combining all patterns

## Final Project Ideas

After completing all lessons, try building:

1. **Personal Assistant**
   - Weather, time, calculator, search tools
   - ReAct pattern for complex queries
   - Memory for conversation context

2. **Code Assistant**
   - File operations, code search, linter tools
   - Structured output for reliability
   - Multi-step code generation

3. **Research Agent**
   - Web search, summarization, citation tools
   - Multi-step research planning
   - Entity extraction and tracking

4. **Customer Support Bot**
   - Database lookup, ticket creation, email tools
   - Simple agent for straightforward queries
   - Escalation to human when needed

## Real-World Applications

Agents are used for:

**Software Development:**
- Code generation and review
- Bug finding and fixing
- Documentation generation
- Test writing

**Business:**
- Customer support automation
- Data analysis and reporting
- Meeting scheduling
- Email management

**Research:**
- Literature review
- Data collection
- Hypothesis testing
- Report writing

**Personal:**
- Task management
- Travel planning
- Shopping research
- Learning assistance

## What's Next

After mastering Agency, you'll move to:

**Part 4: Graphs** - Build complex workflows with state machines and conditional routing

---

Ready? Start with [Lesson 1: Tools](01-tools/lesson.md) →
