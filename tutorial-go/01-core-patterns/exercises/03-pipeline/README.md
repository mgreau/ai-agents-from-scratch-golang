# Exercise 03: Pipeline Composition

Create a multi-step text processing pipeline using Pipe.

## 🎯 Objective

Learn to compose multiple Runnables into a pipeline.

## 📚 Concepts

- Pipeline composition
- Pipe() method
- Data transformation chains
- Input/output compatibility

## 📝 Task

Build a text processing pipeline with three steps:

1. **TrimRunnable** - Removes leading/trailing whitespace
2. **UppercaseRunnable** - Converts to uppercase
3. **PrefixRunnable** - Adds a prefix

Compose them: `Trim -> Uppercase -> Prefix`

## ✅ Requirements

- [ ] Implement all three Runnables
- [ ] Each embeds `*core.BaseRunnable`
- [ ] Compose using Pipe()
- [ ] Test the complete pipeline
- [ ] Ensure output of one matches input of next

## 💡 Hints

1. Each Runnable should do ONE thing
2. Use `strings` package for text operations
3. Chain with `.Pipe()`
4. Test each step individually first

## 🧪 Tests

```go
func TestPipeline(t *testing.T) {
    trim := NewTrimRunnable()
    upper := NewUppercaseRunnable()
    prefix := NewPrefixRunnable(">> ")
    
    pipeline := trim.Pipe(upper).Pipe(prefix)
    
    result, _ := pipeline.Invoke(ctx, "  hello world  ", nil)
    // Should return: ">> HELLO WORLD"
}
```

## 🚀 Next Steps

After completing, move to Exercise 04: Streaming Implementation
