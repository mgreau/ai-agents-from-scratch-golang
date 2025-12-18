# LLM Chain: Prompt + LLM + Parser

**Part 2: Composition - Lesson 3**

> The building block of AI applications.

## Overview

You've learned Prompts (Lesson 1) and Parsers (Lesson 2). Now you'll combine them with an LLM into the most fundamental pattern in AI applications: **the LLM Chain**.

```go
// Before: Manual composition
prompt, _ := template.Format(input)
response, _ := llm.Invoke(ctx, prompt, nil)
result, _ := parser.Parse(response.(string))

// After: LLMChain
chain := NewLLMChain(template, llm, parser)
result, _ := chain.Invoke(ctx, input, nil)
```

## Why This Matters

### The Problem: Repetitive Composition

Every AI feature follows the same pattern:
1. Format prompt with variables
2. Call LLM
3. Parse response

Without abstraction, you repeat this everywhere:

```go
// Feature 1: Translation
prompt1 := formatTranslationPrompt(text, lang)
response1 := callLLM(prompt1)
translation := parseString(response1)

// Feature 2: Sentiment
prompt2 := formatSentimentPrompt(text)
response2 := callLLM(prompt2)
sentiment := parseJSON(response2)

// Feature 3: Summarization  
prompt3 := formatSummaryPrompt(text)
response3 := callLLM(prompt3)
summary := parseString(response3)

// Same pattern, repeated code!
```

### The Solution: LLMChain

```go
// Reusable chains
translationChain := NewLLMChain(translationPrompt, llm, stringParser)
sentimentChain := NewLLMChain(sentimentPrompt, llm, jsonParser)
summaryChain := NewLLMChain(summaryPrompt, llm, stringParser)

// One-line usage
translation, _ := translationChain.Invoke(ctx, input, nil)
sentiment, _ := sentimentChain.Invoke(ctx, input, nil)
summary, _ := summaryChain.Invoke(ctx, input, nil)
```

## Core Concepts

### What is an LLMChain?

An LLMChain is a **Runnable that combines three components**:
1. **PromptTemplate** - Formats input
2. **LLM** - Generates response
3. **OutputParser** - Extracts structured data

**Flow:**
```
Input → PromptTemplate → LLM → OutputParser → Output
  ↓           ↓          ↓          ↓           ↓
{"text": "..."} → "Translate..." → "Hola" → "Hola" → "Hola"
```

### Key Benefits

1. **Encapsulation** - Hide complexity
2. **Reusability** - Use same chain many times
3. **Composability** - Chain can be part of larger pipeline
4. **Testability** - Test components independently
5. **Maintainability** - Change one place, affects all uses

## Implementation in Go

### LLMChain Struct

**Location:** `pkg/chains/llm_chain.go`

```go
package chains

import (
    "context"
    "fmt"

    "github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
    "github.com/mgreau/ai-agents-from-scratch-go/pkg/prompts"
    "github.com/mgreau/ai-agents-from-scratch-go/pkg/parsers"
)

// LLMChain combines prompt + LLM + parser
type LLMChain struct {
    *core.BaseRunnable
    prompt prompts.PromptTemplate
    llm    core.Runnable
    parser parsers.OutputParser
}

// NewLLMChain creates a new chain
func NewLLMChain(prompt prompts.PromptTemplate, llm core.Runnable, parser parsers.OutputParser) *LLMChain {
    return &LLMChain{
        BaseRunnable: core.NewBaseRunnable("LLMChain"),
        prompt:      prompt,
        llm:         llm,
        parser:      parser,
    }
}

// Invoke executes the chain
func (c *LLMChain) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
    // Step 1: Format prompt
    formattedPrompt, err := c.prompt.Invoke(ctx, input, config)
    if err != nil {
        return nil, fmt.Errorf("prompt formatting failed: %w", err)
    }
    
    // Step 2: Call LLM
    llmResponse, err := c.llm.Invoke(ctx, formattedPrompt, config)
    if err != nil {
        return nil, fmt.Errorf("LLM invocation failed: %w", err)
    }
    
    // Step 3: Parse output (optional)
    if c.parser != nil {
        parsed, err := c.parser.Invoke(ctx, llmResponse, config)
        if err != nil {
            return nil, fmt.Errorf("parsing failed: %w", err)
        }
        return parsed, nil
    }
    
    return llmResponse, nil
}

// Or use Pipe for simplicity:
func NewLLMChainWithPipe(prompt prompts.PromptTemplate, llm core.Runnable, parser parsers.OutputParser) core.Runnable {
    if parser != nil {
        return prompt.Pipe(llm).Pipe(parser)
    }
    return prompt.Pipe(llm)
}
```

## Practical Examples

### Example 1: Translation Chain

```go
func translationChain() {
    template := prompts.NewPromptTemplate(prompts.PromptTemplateConfig{
        Template: "Translate to {language}: {text}",
    })
    
    llm, _ := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    defer llm.Close()
    
    parser := parsers.NewStringOutputParser(parsers.StringOutputParserConfig{
        TrimSpace: true,
    })
    
    chain := NewLLMChain(template, llm, parser)
    
    ctx := context.Background()
    result, _ := chain.Invoke(ctx, map[string]string{
        "language": "Spanish",
        "text": "Hello, how are you?",
    }, nil)
    
    fmt.Println(result) // Hola, ¿cómo estás?
}
```

### Example 2: Sentiment Analysis Chain

```go
func sentimentChain() {
    template := prompts.NewPromptTemplate(prompts.PromptTemplateConfig{
        Template: `Analyze sentiment. Respond in JSON:
{"sentiment": "positive/negative/neutral", "confidence": 0.0-1.0}

Text: {text}`,
    })
    
    llm, _ := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    defer llm.Close()
    
    parser := parsers.NewJSONOutputParser(false)
    
    chain := NewLLMChain(template, llm, parser)
    
    ctx := context.Background()
    result, _ := chain.Invoke(ctx, map[string]string{
        "text": "This product is amazing!",
    }, nil)
    
    data := result.(map[string]interface{})
    fmt.Printf("Sentiment: %s (%.0f%%)\n", 
        data["sentiment"], data["confidence"].(float64)*100)
}
```

### Example 3: Chaining Chains

```go
func chainOfChains() {
    // Chain 1: Summarize
    summaryChain := NewLLMChain(summaryTemplate, llm, stringParser)
    
    // Chain 2: Translate summary
    translateChain := NewLLMChain(translateTemplate, llm, stringParser)
    
    // Compose chains
    pipeline := summaryChain.Pipe(translateChain)
    
    result, _ := pipeline.Invoke(ctx, longArticle, nil)
    // Summarizes then translates
}
```

## Exercises

### Exercise 25: Question Answering Chain
Build a Q&A chain with context injection.

### Exercise 26: Multi-Step Chain
Create a chain that calls LLM multiple times.

### Exercise 27: Conditional Chain
Implement branching based on LLM output.

### Exercise 28: Chain with Memory
Add conversation history to chain.

## Key Takeaways

1. ✅ **LLMChain** - Prompt + LLM + Parser combined
2. ✅ **Encapsulation** - Hide implementation details
3. ✅ **Reusability** - Define once, use everywhere
4. ✅ **Composability** - Chains can be chained
5. ✅ **Flexibility** - Parser is optional
6. ✅ **Error handling** - Proper error propagation

## What's Next

**Next Lesson**: [04-piping](../04-piping/lesson.md) - Multi-step pipelines

**See it in action**: Check `pkg/chains/llm_chain.go`

**Practice**: Complete all 4 exercises
