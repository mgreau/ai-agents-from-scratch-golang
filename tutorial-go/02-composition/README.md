# Part 2: Composition

> **Building complex behaviors from simple components**

Welcome to Part 2! In Foundation (Part 1), you learned the core patterns: Runnables, Messages, LLM Wrappers, and Config. Now you'll learn to **compose** these pieces into powerful applications.

## What is Composition?

Composition is the art of combining simple components into complex systems. Instead of building monolithic functions, you create reusable pieces that snap together like LEGO blocks.

**Example:**
```go
// Simple components
promptTemplate := NewPromptTemplate("Translate to {language}: {text}")
llm := NewLlamaCppLLM(config)
parser := NewStringOutputParser()

// Compose into pipeline
translator := promptTemplate.Pipe(llm).Pipe(parser)

// Use anywhere
result, _ := translator.Invoke(ctx, map[string]string{
    "language": "Spanish",
    "text": "Hello world",
}, nil)
// Output: "Hola mundo"
```

## Why Composition Matters

**Without composition:**
```go
// Hardcoded, inflexible
func translate(text, language string) string {
    prompt := fmt.Sprintf("Translate to %s: %s", language, text)
    response, _ := llm.Invoke(ctx, prompt, nil)
    cleaned := strings.TrimSpace(response)
    return cleaned
}
// Can't reuse prompt, can't swap LLM, can't test parts independently
```

**With composition:**
```go
// Reusable, testable, flexible
translator := promptTemplate.Pipe(llm).Pipe(parser)

// Easy to modify
translatorWithHistory := promptTemplate.
    Pipe(llm).
    Pipe(parser).
    Pipe(historyManager)

// Easy to test
mockLLM := NewMockLLM()
testTranslator := promptTemplate.Pipe(mockLLM).Pipe(parser)
```

## What You'll Learn

This section teaches five key composition patterns:

### 1. **Prompts** (Lesson 1)
Template-driven LLM inputs with variable substitution.

**Skills:**
- Create reusable prompt templates
- Inject variables safely
- Build chat prompt templates
- Compose prompts from pieces

**Example:**
```go
template := NewPromptTemplate("Summarize: {text}")
prompt, _ := template.Format(map[string]string{
    "text": "Long article...",
})
```

### 2. **Parsers** (Lesson 2)
Transform LLM text into structured data.

**Skills:**
- Parse JSON from LLM output
- Extract structured data
- Validate output format
- Handle parsing errors

**Example:**
```go
parser := NewJSONOutputParser()
result, _ := parser.Parse(`{"sentiment": "positive", "score": 0.95}`)
// map[string]interface{}{"sentiment": "positive", "score": 0.95}
```

### 3. **LLM Chain** (Lesson 3)
Combine prompts + LLM + parsers into single units.

**Skills:**
- Build Prompt → LLM → Parser chains
- Handle end-to-end flows
- Add error handling
- Create reusable chains

**Example:**
```go
chain := NewLLMChain(promptTemplate, llm, parser)
result, _ := chain.Invoke(ctx, input, nil)
```

### 4. **Piping** (Lesson 4)
Data transformation pipelines with multiple steps.

**Skills:**
- Chain multiple Runnables
- Transform data between steps
- Branch and merge flows
- Debug complex pipelines

**Example:**
```go
pipeline := step1.Pipe(step2).Pipe(step3).Pipe(step4)
result, _ := pipeline.Invoke(ctx, input, nil)
```

### 5. **Memory** (Lesson 5)
Add conversation history and context management.

**Skills:**
- Store conversation history
- Implement sliding windows
- Add memory to chains
- Manage context limits

**Example:**
```go
memory := NewConversationBufferMemory(10)
chain := NewConversationChain(llm, memory)
chain.Invoke(ctx, "Hello", nil)  // Remembers context
chain.Invoke(ctx, "Who am I?", nil)  // Uses history
```

## The Composition Patterns

These five lessons build on each other:

```
Prompts + LLM + Parsers = LLM Chain
         ↓
    Multiple Chains = Pipeline
         ↓
  Pipeline + Memory = Conversational Agent
```

## Go Advantages for Composition

**Type Safety:**
```go
// Compiler catches errors
var chain Runnable = promptTemplate.Pipe(llm).Pipe(parser)
// Type-safe composition
```

**Interfaces:**
```go
// Any Runnable works
func BuildPipeline(steps ...Runnable) Runnable {
    return NewRunnableSequence(steps)
}
```

**Channels for Streaming:**
```go
stream, _ := chain.Stream(ctx, input, nil)
for chunk := range stream {
    fmt.Print(chunk)
}
```

**Goroutines for Parallelism:**
```go
// Run multiple chains concurrently
results := make(chan interface{}, 3)
go func() { results <- chain1.Invoke(ctx, input, nil) }()
go func() { results <- chain2.Invoke(ctx, input, nil) }()
go func() { results <- chain3.Invoke(ctx, input, nil) }()
```

## Prerequisites

Before starting, ensure you understand:
- ✅ Runnable interface (Part 1, Lesson 1)
- ✅ Message types (Part 1, Lesson 2)
- ✅ LLM wrappers (Part 1, Lesson 3)
- ✅ Config and callbacks (Part 1, Lesson 4)

## Lessons

### [Lesson 1: Prompts](01-prompts/lesson.md)
Create reusable prompt templates with variable substitution.

### [Lesson 2: Parsers](02-parsers/lesson.md)
Extract structured data from LLM text output.

### [Lesson 3: LLM Chain](03-llm-chain/lesson.md)
Combine prompts, LLMs, and parsers into cohesive units.

### [Lesson 4: Piping](04-piping/lesson.md)
Build multi-step data transformation pipelines.

### [Lesson 5: Memory](05-memory/lesson.md)
Add conversation history and context management.

## Exercises

Each lesson includes 4 hands-on exercises (20 total).

**Practice path:**
1. Complete all 5 lessons in order
2. Do exercises for each lesson before moving on
3. Build a final project combining all concepts

## Final Project Ideas

After completing all lessons, try building:

1. **Smart Translator**
   - Prompt templates for different styles
   - JSON parser for structured output
   - Memory for context-aware translation

2. **Code Reviewer**
   - Parse code with templates
   - Extract issues with structured parser
   - Chain multiple review steps

3. **Research Assistant**
   - Query templates with examples
   - Parse search results
   - Memory for multi-turn research

## What's Next

After mastering Composition, you'll move to:

**Part 3: Agency** - Build autonomous agents that use tools and make decisions

**Part 4: Graphs** - Create complex workflows with state machines

---

Ready? Start with [Lesson 1: Prompts](01-prompts/lesson.md) →
