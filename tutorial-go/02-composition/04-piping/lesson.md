# Piping: Multi-Step Data Transformation

**Part 2: Composition - Lesson 4**

> Build complex workflows from simple steps.

## Overview

You've mastered individual chains. Now you'll learn to combine multiple Runnables into sophisticated **pipelines** using the Pipe pattern.

```go
// Simple chain
result := step1.Pipe(step2).Pipe(step3).Invoke(ctx, input, nil)

// Data flows through: input → step1 → step2 → step3 → output
```

## Why This Matters

### The Problem: Complex Workflows

Real applications need multi-step processing:

```go
// Manual pipeline (fragile)
prompt := formatPrompt(input)
response := llm.Call(prompt)
parsed := parser.Parse(response)
validated := validator.Check(parsed)
transformed := transformer.Transform(validated)
saved := database.Save(transformed)

// Error handling at every step? Context passing? Cancellation?
```

### The Solution: Pipes

```go
// Declarative pipeline
pipeline := promptTemplate.
    Pipe(llm).
    Pipe(parser).
    Pipe(validator).
    Pipe(transformer).
    Pipe(dbSaver)

result, err := pipeline.Invoke(ctx, input, config)
// Automatic error propagation, context passing, observability
```

## Core Concepts

### What is Piping?

**Piping** connects Runnables so output of one becomes input of next:

```
RunA.Pipe(RunB).Pipe(RunC)
  ↓      ↓        ↓
input → RunA → RunB → RunC → output
```

### RunnableSequence

Implementation of piped Runnables:

```go
type RunnableSequence struct {
    *BaseRunnable
    steps []Runnable
}

func (rs *RunnableSequence) Invoke(ctx context.Context, input interface{}, config *Config) (interface{}, error) {
    current := input
    for i, step := range rs.steps {
        result, err := step.Invoke(ctx, current, config)
        if err != nil {
            return nil, fmt.Errorf("step %d (%s) failed: %w", i, step.Name(), err)
        }
        current = result
    }
    return current, nil
}
```

### Parallel Pipes

Execute multiple steps concurrently:

```go
// Sequential: A → B → C (takes 3 seconds)
sequential := stepA.Pipe(stepB).Pipe(stepC)

// Parallel: All run simultaneously (takes 1 second)
parallel := NewRunnableParallel(stepA, stepB, stepC)
```

## Practical Examples

### Example 1: Document Processing Pipeline

```go
func documentPipeline() {
    // Step 1: Extract text
    extractor := NewTextExtractor()
    
    // Step 2: Summarize
    summaryChain := NewLLMChain(summaryTemplate, llm, stringParser)
    
    // Step 3: Translate
    translateChain := NewLLMChain(translateTemplate, llm, stringParser)
    
    // Step 4: Format
    formatter := NewMarkdownFormatter()
    
    // Compose pipeline
    pipeline := extractor.
        Pipe(summaryChain).
        Pipe(translateChain).
        Pipe(formatter)
    
    // Process document
    result, _ := pipeline.Invoke(ctx, "document.pdf", nil)
    fmt.Println(result)
}
```

### Example 2: Branching Pipeline

```go
func branchingPipeline() {
    // Router decides which path to take
    router := NewRouter(func(input interface{}) string {
        text := input.(string)
        if len(text) > 1000 {
            return "long"
        }
        return "short"
    })
    
    // Different pipelines for different inputs
    longPipeline := summaryChain.Pipe(analyzeChain)
    shortPipeline := directAnalyzeChain
    
    // Conditional execution
    pipeline := router.Branch(map[string]Runnable{
        "long":  longPipeline,
        "short": shortPipeline,
    })
    
    result, _ := pipeline.Invoke(ctx, input, nil)
}
```

### Example 3: Parallel Processing

```go
func parallelProcessing() {
    // Run multiple chains concurrently
    parallel := NewRunnableParallel([]Runnable{
        sentimentChain,
        keywordsChain,
        categorizeChain,
    })
    
    results, _ := parallel.Invoke(ctx, text, nil)
    
    // Results is []interface{} with all outputs
    sentiment := results.([]interface{})[0]
    keywords := results.([]interface{})[1]
    category := results.([]interface{})[2]
}
```

## Exercises

### Exercise 29: ETL Pipeline
Build Extract-Transform-Load pipeline with multiple steps.

### Exercise 30: Retry Wrapper
Create a Runnable that retries failed steps.

### Exercise 31: Caching Layer
Add caching to pipeline to avoid redundant LLM calls.

### Exercise 32: Parallel Map
Implement map operation that processes list items concurrently.

## Key Takeaways

1. ✅ **Pipe()** - Connect Runnables sequentially
2. ✅ **RunnableSequence** - Execute steps in order
3. ✅ **RunnableParallel** - Execute concurrently
4. ✅ **Error propagation** - Failures bubble up
5. ✅ **Context passing** - Config flows through
6. ✅ **Composability** - Pipelines are Runnables too

## What's Next

**Next Lesson**: [05-memory](../05-memory/lesson.md) - Add conversation history

**See it in action**: Check `pkg/core/runnable.go`

**Practice**: Complete all 4 exercises
