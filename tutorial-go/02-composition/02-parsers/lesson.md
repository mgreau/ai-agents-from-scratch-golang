# Output Parsers: Structured Output Extraction

**Part 2: Composition - Lesson 2**

> LLMs return text. You need data.

## Overview

You've learned to create great prompts. LLMs return unstructured text, but you often need structured data:

```go
// LLM returns this:
"The sentiment is positive with a confidence of 0.92"

// You need this:
map[string]interface{}{
    "sentiment": "positive",
    "confidence": 0.92,
}
```

**Output parsers** transform LLM text into structured data you can use in your applications.

## Why This Matters

### The Problem: Parsing Chaos

Without parsers, your code is full of brittle string manipulation:

```go
response, _ := llm.Invoke(ctx, "Classify: I love this product!", nil)

// Fragile parsing everywhere
responseStr := response.(string)
if strings.Contains(responseStr, "positive") {
    sentiment = "positive"
} else if strings.Contains(responseStr, "negative") {
    sentiment = "negative"
}

// What if format changes?
// What if LLM adds extra text?
// How do you handle errors?
```

Problems:
- Brittle regex and string matching
- No validation of output format
- Hard to test parsing logic
- Inconsistent error handling
- Parser code duplicated everywhere

### The Solution: Output Parsers

```go
parser := NewJSONOutputParser()

template := NewPromptTemplate(PromptTemplateConfig{
    Template: `Classify the sentiment. Respond in JSON:
{"sentiment": "positive/negative/neutral", "confidence": 0.0-1.0}

Text: {text}`,
})

chain := template.Pipe(llm).Pipe(parser)

result, _ := chain.Invoke(ctx, map[string]string{
    "text": "I love this!",
}, nil)
// map[string]interface{}{"sentiment": "positive", "confidence": 0.95}
```

Benefits:
- ✅ Reliable structured extraction
- ✅ Format validation
- ✅ Error handling built-in
- ✅ Reusable parsing logic
- ✅ Type-safe outputs

## Learning Objectives

By the end of this lesson, you will:

- ✅ Build a BaseOutputParser abstraction
- ✅ Create StringOutputParser for text cleanup
- ✅ Implement JSONOutputParser for JSON extraction
- ✅ Build ListOutputParser for arrays
- ✅ Handle parsing errors gracefully
- ✅ Use parsers in chains with prompts

## Core Concepts

### What is an Output Parser?

An output parser **transforms LLM text output into structured data**.

**Flow:**
```
LLM Output (text) → Parser → Structured Data
    ↓                ↓              ↓
"positive: 0.95"  Parse()    {"sentiment": "positive", "confidence": 0.95}
```

### The Parser Hierarchy

```
BaseOutputParser (interface)
    ├── StringOutputParser (clean text)
    ├── JSONOutputParser (extract JSON)
    ├── ListOutputParser (extract lists)
    └── RegexOutputParser (regex patterns)
```

### Key Operations

1. **Parse**: Extract structured data from text
2. **GetFormatInstructions**: Tell LLM how to format response
3. **Validate**: Check output matches expected structure
4. **Handle Errors**: Gracefully handle malformed outputs

## Implementation in Go

### Step 1: BaseOutputParser Interface

**Location:** `pkg/parsers/base_parser.go`

```go
package parsers

import (
    "context"
    "fmt"

    "github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// OutputParser transforms LLM text into structured data
type OutputParser interface {
    core.Runnable
    Parse(text string) (interface{}, error)
    GetFormatInstructions() string
}

// ParseError represents a parsing failure
type ParseError struct {
    Message   string
    LLMOutput string
    Err       error
}

func (e *ParseError) Error() string {
    return fmt.Sprintf("parse error: %s (output: %s)", e.Message, e.LLMOutput)
}

func (e *ParseError) Unwrap() error {
    return e.Err
}
```

**Why this design:**
- Interface for flexibility
- ParseError with context
- Embeds Runnable for composability

### Step 2: StringOutputParser

Simplest parser - cleans up text:

```go
// StringOutputParser cleans text output
type StringOutputParser struct {
    *core.BaseRunnable
    trimSpace     bool
    removeMarkdown bool
}

// StringOutputParserConfig holds configuration
type StringOutputParserConfig struct {
    TrimSpace     bool
    RemoveMarkdown bool
}

// NewStringOutputParser creates a string parser
func NewStringOutputParser(config StringOutputParserConfig) *StringOutputParser {
    return &StringOutputParser{
        BaseRunnable:   core.NewBaseRunnable("StringOutputParser"),
        trimSpace:     config.TrimSpace,
        removeMarkdown: config.RemoveMarkdown,
    }
}

// Parse cleans the text
func (p *StringOutputParser) Parse(text string) (interface{}, error) {
    result := text
    
    if p.removeMarkdown {
        // Remove markdown code blocks
        result = removeCodeBlocks(result)
    }
    
    if p.trimSpace {
        result = strings.TrimSpace(result)
    }
    
    return result, nil
}

// GetFormatInstructions returns empty (no special format needed)
func (p *StringOutputParser) GetFormatInstructions() string {
    return ""
}

// Invoke implements Runnable interface
func (p *StringOutputParser) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
    text, ok := input.(string)
    if !ok {
        // Try to get content from Message
        if msg, ok := input.(core.Message); ok {
            text = msg.GetContent()
        } else {
            return nil, fmt.Errorf("input must be string or Message, got %T", input)
        }
    }
    
    return p.Parse(text)
}

// removeCodeBlocks strips markdown code blocks
func removeCodeBlocks(text string) string {
    // Remove ```language\n...\n```
    re := regexp.MustCompile("```[a-z]*\\n([\\s\\S]*?)```")
    matches := re.FindStringSubmatch(text)
    if len(matches) > 1 {
        return matches[1]
    }
    return text
}
```

### Step 3: JSONOutputParser

Extracts and validates JSON:

```go
// JSONOutputParser extracts JSON from text
type JSONOutputParser struct {
    *core.BaseRunnable
    strict bool // Require valid JSON only
}

// NewJSONOutputParser creates a JSON parser
func NewJSONOutputParser(strict bool) *JSONOutputParser {
    return &JSONOutputParser{
        BaseRunnable: core.NewBaseRunnable("JSONOutputParser"),
        strict:      strict,
    }
}

// Parse extracts JSON from text
func (p *JSONOutputParser) Parse(text string) (interface{}, error) {
    // Try to extract JSON from text
    jsonStr := p.extractJSON(text)
    if jsonStr == "" {
        if p.strict {
            return nil, &ParseError{
                Message:   "no JSON found in output",
                LLMOutput: text,
            }
        }
        jsonStr = text // Try parsing whole text
    }
    
    // Parse JSON
    var result interface{}
    if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
        return nil, &ParseError{
            Message:   "invalid JSON",
            LLMOutput: text,
            Err:       err,
        }
    }
    
    return result, nil
}

// extractJSON finds JSON object or array in text
func (p *JSONOutputParser) extractJSON(text string) string {
    // Find JSON object {...}
    if start := strings.Index(text, "{"); start != -1 {
        // Find matching closing brace
        depth := 0
        for i := start; i < len(text); i++ {
            switch text[i] {
            case '{':
                depth++
            case '}':
                depth--
                if depth == 0 {
                    return text[start : i+1]
                }
            }
        }
    }
    
    // Find JSON array [...]
    if start := strings.Index(text, "["); start != -1 {
        depth := 0
        for i := start; i < len(text); i++ {
            switch text[i] {
            case '[':
                depth++
            case ']':
                depth--
                if depth == 0 {
                    return text[start : i+1]
                }
            }
        }
    }
    
    return ""
}

// GetFormatInstructions tells LLM how to format output
func (p *JSONOutputParser) GetFormatInstructions() string {
    return `Respond with valid JSON only. Example:
{"key": "value", "number": 123}`
}

// Invoke implements Runnable
func (p *JSONOutputParser) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
    text, ok := input.(string)
    if !ok {
        if msg, ok := input.(core.Message); ok {
            text = msg.GetContent()
        } else {
            return nil, fmt.Errorf("input must be string or Message")
        }
    }
    
    return p.Parse(text)
}
```

### Step 4: ListOutputParser

Extracts lists/arrays:

```go
// ListOutputParser extracts lists from text
type ListOutputParser struct {
    *core.BaseRunnable
    separator string
}

// NewListOutputParser creates a list parser
func NewListOutputParser(separator string) *ListOutputParser {
    if separator == "" {
        separator = "\n"
    }
    return &ListOutputParser{
        BaseRunnable: core.NewBaseRunnable("ListOutputParser"),
        separator:   separator,
    }
}

// Parse extracts list items
func (p *ListOutputParser) Parse(text string) (interface{}, error) {
    // First try JSON array
    var jsonArray []string
    if err := json.Unmarshal([]byte(text), &jsonArray); err == nil {
        return jsonArray, nil
    }
    
    // Fall back to separator-based splitting
    lines := strings.Split(text, p.separator)
    
    var result []string
    for _, line := range lines {
        line = strings.TrimSpace(line)
        // Remove common list markers
        line = strings.TrimPrefix(line, "- ")
        line = strings.TrimPrefix(line, "* ")
        line = strings.TrimPrefix(line, "• ")
        
        // Remove numbered list markers (1., 2., etc.)
        re := regexp.MustCompile(`^\d+\.\s*`)
        line = re.ReplaceAllString(line, "")
        
        if line != "" {
            result = append(result, line)
        }
    }
    
    return result, nil
}

// GetFormatInstructions tells LLM how to format
func (p *ListOutputParser) GetFormatInstructions() string {
    return fmt.Sprintf(`Respond with a list of items, one per line.
Separate items with "%s"
Example:
Item 1
Item 2
Item 3`, p.separator)
}

// Invoke implements Runnable
func (p *ListOutputParser) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
    text, ok := input.(string)
    if !ok {
        if msg, ok := input.(core.Message); ok {
            text = msg.GetContent()
        } else {
            return nil, fmt.Errorf("input must be string or Message")
        }
    }
    
    return p.Parse(text)
}
```

## Practical Examples

### Example 1: String Parser

```go
func stringParserExample() {
    parser := NewStringOutputParser(StringOutputParserConfig{
        TrimSpace:     true,
        RemoveMarkdown: true,
    })
    
    text := "```go\nfunc main() {}\n```"
    result, _ := parser.Parse(text)
    
    fmt.Println(result)
    // Output: func main() {}
}
```

### Example 2: JSON Parser

```go
func jsonParserExample() {
    parser := NewJSONOutputParser(true)
    
    text := `The sentiment analysis shows: {"sentiment": "positive", "score": 0.95}`
    result, _ := parser.Parse(text)
    
    data := result.(map[string]interface{})
    fmt.Printf("Sentiment: %s, Score: %.2f\n", data["sentiment"], data["score"])
    // Output: Sentiment: positive, Score: 0.95
}
```

### Example 3: List Parser

```go
func listParserExample() {
    parser := NewListOutputParser("\n")
    
    text := `Here are the steps:
1. Install Go
2. Write code
3. Run tests
4. Deploy`
    
    result, _ := parser.Parse(text)
    items := result.([]string)
    
    for i, item := range items {
        fmt.Printf("%d: %s\n", i+1, item)
    }
    // Output:
    // 1: Install Go
    // 2: Write code
    // 3: Run tests
    // 4: Deploy
}
```

### Example 4: In a Pipeline

```go
func parserInPipeline() {
    // Create components
    template := NewPromptTemplate(PromptTemplateConfig{
        Template: `Analyze sentiment and respond in JSON:
{"sentiment": "positive/negative/neutral", "confidence": 0.0-1.0}

Text: {text}`,
    })
    
    llm, _ := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    defer llm.Close()
    
    parser := NewJSONOutputParser(false)
    
    // Compose: Template -> LLM -> Parser
    chain := template.Pipe(llm).Pipe(parser)
    
    // Use chain
    ctx := context.Background()
    result, _ := chain.Invoke(ctx, map[string]string{
        "text": "I absolutely love this product!",
    }, nil)
    
    data := result.(map[string]interface{})
    fmt.Printf("Sentiment: %s (%.0f%% confidence)\n", 
        data["sentiment"], 
        data["confidence"].(float64)*100)
}
```

### Example 5: Error Handling

```go
func errorHandling() {
    parser := NewJSONOutputParser(true)
    
    // Invalid JSON
    text := "This is not JSON at all"
    result, err := parser.Parse(text)
    
    if parseErr, ok := err.(*ParseError); ok {
        fmt.Printf("Parse failed: %s\n", parseErr.Message)
        fmt.Printf("LLM output was: %s\n", parseErr.LLMOutput)
        
        // Could retry with different prompt
        // or fall back to string parsing
    }
}
```

## Exercises

### Exercise 21: CSV Parser
Build a parser that extracts CSV data from text.
- [Starter Code](exercises/21-csv-parser/starter.go)
- [Solution](exercises/21-csv-parser/solution.go)

### Exercise 22: Structured Output
Create a parser with schema validation (required fields).
- [Starter Code](exercises/22-structured-output/starter.go)
- [Solution](exercises/22-structured-output/solution.go)

### Exercise 23: Regex Parser
Implement a parser using regular expressions.
- [Starter Code](exercises/23-regex-parser/starter.go)
- [Solution](exercises/23-regex-parser/solution.go)

### Exercise 24: Fallback Parser
Build a parser that tries multiple strategies.
- [Starter Code](exercises/24-fallback-parser/starter.go)
- [Solution](exercises/24-fallback-parser/solution.go)

## Key Takeaways

1. ✅ **BaseOutputParser** - Interface for all parsers
2. ✅ **StringOutputParser** - Clean text output
3. ✅ **JSONOutputParser** - Extract and validate JSON
4. ✅ **ListOutputParser** - Extract arrays/lists
5. ✅ **GetFormatInstructions** - Guide LLM output format
6. ✅ **ParseError** - Structured error handling
7. ✅ **Composable** - Use in pipelines with Pipe()

## What's Next

**Next Lesson**: [03-llm-chain](../03-llm-chain/lesson.md) - Combine prompts, LLMs, and parsers

**See it in action**: Check `pkg/parsers/` for implementations

**Practice**: Complete all 4 exercises before moving on

## Further Reading

- [JSON in Go](https://go.dev/blog/json)
- [Regex in Go](https://pkg.go.dev/regexp)
- [LangChain Output Parsers](https://python.langchain.com/docs/modules/model_io/output_parsers/)
