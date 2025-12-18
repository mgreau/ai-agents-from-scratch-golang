# Exercise 09: Basic LLM Wrapper

Build a minimal LLM wrapper that implements the Runnable interface.

## 🎯 Objective

Learn to wrap an external library as a Runnable.

## 📚 Concepts

- Wrapper pattern
- Interface implementation
- Resource management
- Type conversion

## 📝 Task

Implement `MockLLMRunnable` that:

1. Embeds `*core.BaseRunnable`
2. Accepts string or []Message input
3. Returns mock responses (no real LLM needed)
4. Implements temperature parameter
5. Properly handles errors

## ✅ Requirements

- [ ] Embed `*core.BaseRunnable`
- [ ] Implement `Invoke()` method
- [ ] Support both string and []Message inputs
- [ ] Return string responses
- [ ] Add temperature field (0.0-1.0)
- [ ] Handle invalid inputs

## 💡 Hints

1. Create config struct for initialization
2. Use type switch for input handling
3. For messages, convert to simple prompt
4. Mock responses can be deterministic or random
5. No need for actual LLM - focus on interface

## 🧪 Tests

```go
func TestMockLLM(t *testing.T) {
    llm := NewMockLLMRunnable(MockLLMConfig{
        Temperature: 0.7,
    })
    
    // Test with string
    response, _ := llm.Invoke(ctx, "Hello", nil)
    
    // Test with messages
    messages := []core.Message{
        core.NewHumanMessage("Hello", nil),
    }
    response, _ = llm.Invoke(ctx, messages, nil)
}
```

## 🚀 Next Steps

After completing, move to Exercise 10: Batch Processing
