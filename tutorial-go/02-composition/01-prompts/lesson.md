# Prompts: Template-Driven LLM Inputs

**Part 2: Composition - Lesson 1**

> Stop hardcoding prompts. Start composing them.

## Overview

You've been writing prompts like this:

```go
prompt := fmt.Sprintf("You are a helpful assistant. The user asked: %s", userInput)
response, _ := llm.Invoke(ctx, prompt, nil)
```

This works, but it's fragile. What if you need:
- Different system messages for different use cases?
- To inject multiple variables safely?
- To reuse prompt patterns across your app?
- To validate inputs before sending to the LLM?
- To compose prompts from smaller pieces?

**PromptTemplates** solve all of these problems.

## Why This Matters

### The Problem: Prompt Chaos

Without templates, your code becomes a mess:

```go
// Scattered throughout your codebase:
prompt1 := fmt.Sprintf("Translate to %s: %s", lang, text)
prompt2 := "Translate to " + language + ": " + input
prompt3 := fmt.Sprintf("Translate to %s: %s", targetLanguage, userText)

// Same logic, different implementations everywhere!
```

Problems:
- No consistency in prompt format
- Hard to test prompts in isolation
- Can't reuse prompt patterns
- SQL-injection-like issues with direct string formatting
- No validation of variables

### The Solution: PromptTemplate

```go
translatePrompt := NewPromptTemplate(PromptTemplateConfig{
    Template: "Translate to {language}: {text}",
    InputVariables: []string{"language", "text"},
})

prompt, _ := translatePrompt.Format(map[string]string{
    "language": "Spanish",
    "text": "Hello, world!",
})
// "Translate to Spanish: Hello, world!"
```

Benefits:
- ✅ Reusable prompt patterns
- ✅ Variable validation
- ✅ Testable in isolation
- ✅ Composable with other Runnables
- ✅ Type-safe variable injection

## Learning Objectives

By the end of this lesson, you will:

- ✅ Build a PromptTemplate that replaces variables
- ✅ Create ChatPromptTemplate for structured messages
- ✅ Implement Few-Shot prompts with examples
- ✅ Use prompts as Runnables in chains
- ✅ Validate template variables
- ✅ Compose templates from pieces

## Core Concepts

### What is a PromptTemplate?

A PromptTemplate is a **reusable prompt pattern** with placeholders for variables.

**Structure:**
```
Template String: "Translate to {language}: {text}"
                          ↓
Variables Injected: {"language": "Spanish", "text": "Hello"}
                          ↓
Output: "Translate to Spanish: Hello"
```

### Template Syntax

We use `{variable}` syntax (inspired by Python's f-strings and Go's templates):

```go
template := "Hello {name}, you are {age} years old in {city}."
// Variables: name, age, city
```

### Key Operations

1. **Format**: Replace variables with values
2. **Validate**: Check required variables are provided
3. **FormatMessages**: Return as Message objects
4. **Invoke**: Use as a Runnable (returns formatted prompt)

## Implementation in Go

### Step 1: PromptTemplate Struct

**Location:** `pkg/prompts/prompt_template.go`

```go
package prompts

import (
    "context"
    "fmt"
    "regexp"
    "strings"

    "github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// PromptTemplateConfig holds configuration
type PromptTemplateConfig struct {
    Template        string
    InputVariables  []string
    PartialVariables map[string]string
}

// PromptTemplate is a template with {variable} placeholders
type PromptTemplate struct {
    *core.BaseRunnable
    template        string
    inputVariables  []string
    partialVariables map[string]string
}

// NewPromptTemplate creates a new prompt template
func NewPromptTemplate(config PromptTemplateConfig) *PromptTemplate {
    pt := &PromptTemplate{
        BaseRunnable:    core.NewBaseRunnable("PromptTemplate"),
        template:        config.Template,
        inputVariables:  config.InputVariables,
        partialVariables: config.PartialVariables,
    }
    
    if pt.partialVariables == nil {
        pt.partialVariables = make(map[string]string)
    }
    
    // Auto-detect variables if not provided
    if len(pt.inputVariables) == 0 {
        pt.inputVariables = pt.extractVariables()
    }
    
    return pt
}

// extractVariables finds all {variable} placeholders
func (pt *PromptTemplate) extractVariables() []string {
    re := regexp.MustCompile(`\{([a-zA-Z_][a-zA-Z0-9_]*)\}`)
    matches := re.FindAllStringSubmatch(pt.template, -1)
    
    seen := make(map[string]bool)
    var variables []string
    
    for _, match := range matches {
        if len(match) > 1 {
            varName := match[1]
            if !seen[varName] {
                seen[varName] = true
                variables = append(variables, varName)
            }
        }
    }
    
    return variables
}
```

**Why this design:**
- Config struct for clean initialization
- Auto-detection of variables from template
- Partial variables for pre-filled defaults
- Embeds BaseRunnable for composability

### Step 2: Format Method

```go
// Format replaces variables in template
func (pt *PromptTemplate) Format(values map[string]string) (string, error) {
    // Validate required variables
    if err := pt.validate(values); err != nil {
        return "", err
    }
    
    // Merge partial and provided variables
    allValues := make(map[string]string)
    for k, v := range pt.partialVariables {
        allValues[k] = v
    }
    for k, v := range values {
        allValues[k] = v
    }
    
    // Replace all variables
    result := pt.template
    for varName, value := range allValues {
        placeholder := fmt.Sprintf("{%s}", varName)
        result = strings.ReplaceAll(result, placeholder, value)
    }
    
    return result, nil
}

// validate checks all required variables are provided
func (pt *PromptTemplate) validate(values map[string]string) error {
    // Merge with partials for validation
    allValues := make(map[string]string)
    for k, v := range pt.partialVariables {
        allValues[k] = v
    }
    for k, v := range values {
        allValues[k] = v
    }
    
    // Check for missing variables
    var missing []string
    for _, varName := range pt.inputVariables {
        if _, exists := allValues[varName]; !exists {
            missing = append(missing, varName)
        }
    }
    
    if len(missing) > 0 {
        return fmt.Errorf("missing required variables: %v", missing)
    }
    
    return nil
}
```

**Key features:**
- Validation before formatting
- Merges partial and provided variables
- Clear error messages for missing variables

### Step 3: Runnable Integration

```go
// Invoke implements the Runnable interface
func (pt *PromptTemplate) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
    // Input should be map[string]string
    values, ok := input.(map[string]string)
    if !ok {
        return nil, fmt.Errorf("input must be map[string]string, got %T", input)
    }
    
    return pt.Format(values)
}
```

**This makes prompts composable:**
```go
// Use in a pipeline
chain := promptTemplate.Pipe(llm).Pipe(parser)
result, _ := chain.Invoke(ctx, variables, nil)
```

### Step 4: ChatPromptTemplate

For structured conversations with messages:

```go
// ChatPromptTemplate creates message sequences
type ChatPromptTemplate struct {
    *core.BaseRunnable
    messages []MessageTemplate
}

// MessageTemplate represents a single message in the chat
type MessageTemplate struct {
    Role     core.MessageType
    Template string
}

// NewChatPromptTemplate creates a chat prompt
func NewChatPromptTemplate(messages []MessageTemplate) *ChatPromptTemplate {
    return &ChatPromptTemplate{
        BaseRunnable: core.NewBaseRunnable("ChatPromptTemplate"),
        messages:     messages,
    }
}

// FormatMessages returns a slice of Messages
func (cpt *ChatPromptTemplate) FormatMessages(values map[string]string) ([]core.Message, error) {
    result := make([]core.Message, len(cpt.messages))
    
    for i, msgTemplate := range cpt.messages {
        // Format the template
        content := msgTemplate.Template
        for varName, value := range values {
            placeholder := fmt.Sprintf("{%s}", varName)
            content = strings.ReplaceAll(content, placeholder, value)
        }
        
        // Create appropriate message type
        switch msgTemplate.Role {
        case core.MessageTypeSystem:
            result[i] = core.NewSystemMessage(content, nil)
        case core.MessageTypeHuman:
            result[i] = core.NewHumanMessage(content, nil)
        case core.MessageTypeAI:
            result[i] = core.NewAIMessage(content, nil)
        default:
            return nil, fmt.Errorf("unsupported message type: %s", msgTemplate.Role)
        }
    }
    
    return result, nil
}

// Invoke for Runnable interface
func (cpt *ChatPromptTemplate) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
    values, ok := input.(map[string]string)
    if !ok {
        return nil, fmt.Errorf("input must be map[string]string")
    }
    
    return cpt.FormatMessages(values)
}
```

## Practical Examples

### Example 1: Simple Translation Template

```go
func simpleTemplate() {
    template := NewPromptTemplate(PromptTemplateConfig{
        Template: "Translate to {language}: {text}",
    })
    
    prompt, _ := template.Format(map[string]string{
        "language": "French",
        "text": "Hello, how are you?",
    })
    
    fmt.Println(prompt)
    // Output: Translate to French: Hello, how are you?
}
```

### Example 2: With Partial Variables

```go
func withPartials() {
    template := NewPromptTemplate(PromptTemplateConfig{
        Template: "You are a {role}. {instruction}",
        PartialVariables: map[string]string{
            "role": "helpful assistant",
        },
    })
    
    prompt, _ := template.Format(map[string]string{
        "instruction": "Explain quantum computing simply.",
    })
    
    fmt.Println(prompt)
    // Output: You are a helpful assistant. Explain quantum computing simply.
}
```

### Example 3: ChatPromptTemplate

```go
func chatTemplate() {
    template := NewChatPromptTemplate([]MessageTemplate{
        {Role: core.MessageTypeSystem, Template: "You are a {role}"},
        {Role: core.MessageTypeHuman, Template: "{user_input}"},
    })
    
    messages, _ := template.FormatMessages(map[string]string{
        "role": "coding tutor",
        "user_input": "Explain Go interfaces",
    })
    
    for _, msg := range messages {
        fmt.Printf("%s: %s\n", msg.GetType(), msg.GetContent())
    }
    // Output:
    // system: You are a coding tutor
    // human: Explain Go interfaces
}
```

### Example 4: In a Pipeline

```go
func inPipeline() {
    template := NewPromptTemplate(PromptTemplateConfig{
        Template: "Summarize this in one sentence: {text}",
    })
    
    llm, _ := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    defer llm.Close()
    
    // Compose: Template -> LLM
    chain := template.Pipe(llm)
    
    ctx := context.Background()
    result, _ := chain.Invoke(ctx, map[string]string{
        "text": "Long article about Go programming...",
    }, nil)
    
    fmt.Println("Summary:", result)
}
```

### Example 5: Few-Shot Template

```go
type FewShotPromptTemplate struct {
    *core.BaseRunnable
    exampleTemplate string
    examples        []map[string]string
    prefix          string
    suffix          string
}

func NewFewShotPromptTemplate(config FewShotPromptTemplateConfig) *FewShotPromptTemplate {
    return &FewShotPromptTemplate{
        BaseRunnable:    core.NewBaseRunnable("FewShotPrompt"),
        exampleTemplate: config.ExampleTemplate,
        examples:        config.Examples,
        prefix:          config.Prefix,
        suffix:          config.Suffix,
    }
}

func (fpt *FewShotPromptTemplate) Format(values map[string]string) (string, error) {
    var parts []string
    
    // Add prefix
    if fpt.prefix != "" {
        parts = append(parts, fpt.prefix)
    }
    
    // Add examples
    for _, example := range fpt.examples {
        exampleText := fpt.exampleTemplate
        for k, v := range example {
            placeholder := fmt.Sprintf("{%s}", k)
            exampleText = strings.ReplaceAll(exampleText, placeholder, v)
        }
        parts = append(parts, exampleText)
    }
    
    // Add suffix with user input
    if fpt.suffix != "" {
        suffixText := fpt.suffix
        for k, v := range values {
            placeholder := fmt.Sprintf("{%s}", k)
            suffixText = strings.ReplaceAll(suffixText, placeholder, v)
        }
        parts = append(parts, suffixText)
    }
    
    return strings.Join(parts, "\n\n"), nil
}

// Usage:
func fewShotExample() {
    template := NewFewShotPromptTemplate(FewShotPromptTemplateConfig{
        Prefix: "Classify sentiment as positive, negative, or neutral:",
        ExampleTemplate: "Text: {text}\nSentiment: {sentiment}",
        Examples: []map[string]string{
            {"text": "I love this!", "sentiment": "positive"},
            {"text": "This is terrible", "sentiment": "negative"},
            {"text": "It's okay", "sentiment": "neutral"},
        },
        Suffix: "Text: {text}\nSentiment:",
    })
    
    prompt, _ := template.Format(map[string]string{
        "text": "This product is amazing!",
    })
    
    fmt.Println(prompt)
    // Output includes all examples + new input
}
```

## Exercises

### Exercise 17: Email Template
Build a template for professional emails with multiple variables.
- [Starter Code](exercises/17-email-template/starter.go)
- [Solution](exercises/17-email-template/solution.go)

### Exercise 18: ChatPromptTemplate
Create a multi-message chat template with system, human, and AI messages.
- [Starter Code](exercises/18-chat-template/starter.go)
- [Solution](exercises/18-chat-template/solution.go)

### Exercise 19: Few-Shot Learning
Implement a few-shot template with examples.
- [Starter Code](exercises/19-few-shot/starter.go)
- [Solution](exercises/19-few-shot/solution.go)

### Exercise 20: Template Validation
Add advanced validation (type checking, value constraints).
- [Starter Code](exercises/20-template-validation/starter.go)
- [Solution](exercises/20-template-validation/solution.go)

## Key Takeaways

1. ✅ **PromptTemplate** - Reusable prompt patterns with variables
2. ✅ **ChatPromptTemplate** - Structured message sequences
3. ✅ **Validation** - Check variables before formatting
4. ✅ **Partial variables** - Pre-filled defaults
5. ✅ **Few-shot** - Include examples in prompts
6. ✅ **Composable** - Use in pipelines with Pipe()
7. ✅ **Type-safe** - Go's type system catches errors

## What's Next

**Next Lesson**: [02-parsers](../02-parsers/lesson.md) - Extract structured data from LLM output

**See it in action**: Check `pkg/prompts/prompt_template.go` for implementation

**Practice**: Complete all 4 exercises before moving on

## Further Reading

- [Go text/template](https://pkg.go.dev/text/template)
- [LangChain Prompts](https://python.langchain.com/docs/modules/model_io/prompts/)
- [Prompt Engineering Guide](https://www.promptingguide.ai/)
