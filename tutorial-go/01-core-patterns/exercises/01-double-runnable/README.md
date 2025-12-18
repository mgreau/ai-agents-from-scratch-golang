# Exercise 01: Double Runnable

Build a simple Runnable that doubles numeric input.

## 🎯 Objective

Learn the Runnable pattern by implementing a minimal Runnable that doubles numbers.

## 📚 Concepts

- Runnable interface
- BaseRunnable embedding
- Implementing the call method
- Error handling

## 📝 Task

Implement `DoubleRunnable` that:

1. Takes a number as input
2. Doubles it
3. Returns the result
4. Handles invalid input gracefully

## ✅ Requirements

- [ ] Embed `*core.BaseRunnable`
- [ ] Implement `Invoke()` method
- [ ] Handle non-numeric input with error
- [ ] Pass all tests

## 🚀 Getting Started

```bash
# Review the starter code
cat starter.go

# Implement the TODOs
vim starter.go

# Run tests
go test -v

# Expected output:
# === RUN   TestDoubleRunnable
# --- PASS: TestDoubleRunnable (0.00s)
# === RUN   TestDoubleRunnableInvalidInput
# --- PASS: TestDoubleRunnableInvalidInput (0.00s)
# PASS
```

## 💡 Hints

1. Use type assertion to convert `interface{}` to `int` or `float64`
2. Return an error for invalid input types
3. The `Invoke` method is already implemented by `BaseRunnable`
4. You only need to implement the internal logic

## 📖 Example Usage

```go
doubler := NewDoubleRunnable()

// Use Invoke
result, err := doubler.Invoke(context.Background(), 5, nil)
// result: 10

// Use Batch
results, err := doubler.Batch(context.Background(), []interface{}{1, 2, 3}, nil)
// results: [2, 4, 6]

// Use Pipe
tripler := NewTripleRunnable()
result, err := doubler.Pipe(tripler).Invoke(context.Background(), 2, nil)
// result: 12 (2 * 2 * 3)
```

## 🧪 Tests

Your implementation should pass:

```go
func TestDoubleRunnable(t *testing.T) {
    // Test basic doubling
    // Test with different number types
    // Test with zero
    // Test with negative numbers
}

func TestDoubleRunnableInvalidInput(t *testing.T) {
    // Test with string input
    // Test with nil input
    // Test with struct input
}
```

## 🔍 Common Mistakes

❌ **Wrong:**
```go
func (d *DoubleRunnable) Invoke(...) {
    // Don't override Invoke directly
}
```

✅ **Right:**
```go
func (d *DoubleRunnable) call(...) {
    // Implement the internal logic
}
```

❌ **Wrong:**
```go
num := input.(int) // Will panic if input isn't int
```

✅ **Right:**
```go
num, ok := input.(int)
if !ok {
    return nil, fmt.Errorf("input must be int")
}
```

## 🎓 Learning Points

After completing this exercise, you should understand:

1. ✅ How to implement the Runnable interface
2. ✅ The difference between Invoke() and call()
3. ✅ Why BaseRunnable handles Invoke/Stream/Batch
4. ✅ How to use type assertions safely
5. ✅ Why consistent interfaces enable composition

## 🚀 Next Steps

1. Complete the implementation
2. Pass all tests
3. Read the solution if stuck
4. Move to Exercise 02: Conversation Manager

## 💻 Solution Available

If you're stuck, check `solution.go` for a reference implementation.

Don't look until you've tried!
