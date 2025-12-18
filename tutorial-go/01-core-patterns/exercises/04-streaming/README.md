# Exercise 04: Streaming Implementation

Implement true character-by-character streaming in a Runnable.

## 🎯 Objective

Learn to implement streaming with Go channels and goroutines.

## 📚 Concepts

- Channel-based streaming
- Goroutines for async processing
- `defer close(channel)`
- Context cancellation

## 📝 Task

Implement `StreamingTextRunnable` that:

1. Takes a string as input
2. Streams each character individually via a channel
3. Adds a configurable delay between characters
4. Respects context cancellation

## ✅ Requirements

- [ ] Override `Stream()` method
- [ ] Use a goroutine to send characters
- [ ] Use `defer close(out)` for cleanup
- [ ] Check `ctx.Done()` for cancellation
- [ ] Add delay between characters
- [ ] Buffered channel for performance

## 💡 Hints

1. Create buffered channel: `make(chan interface{}, 10)`
2. Use `time.Sleep()` for delay
3. Convert string to runes for proper Unicode handling
4. Use `select` to check context cancellation

## 🧪 Tests

```go
func TestStreaming(t *testing.T) {
    streamer := NewStreamingTextRunnable(100 * time.Millisecond)
    
    stream, _ := streamer.Stream(ctx, "Hello", nil)
    
    for char := range stream {
        fmt.Print(char) // Prints: H e l l o (with delays)
    }
}
```

## 🚀 Next Steps

After completing, you've finished all Runnable exercises!
Move to Lesson 02: Messages & Types
