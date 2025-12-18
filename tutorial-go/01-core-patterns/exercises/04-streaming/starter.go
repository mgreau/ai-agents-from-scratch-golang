package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// StreamingTextRunnable streams text character by character
type StreamingTextRunnable struct {
	*core.BaseRunnable
	Delay time.Duration // Delay between characters
}

func NewStreamingTextRunnable(delay time.Duration) *StreamingTextRunnable {
	return &StreamingTextRunnable{
		BaseRunnable: core.NewBaseRunnable("StreamingText"),
		Delay:        delay,
	}
}

// TODO: Implement Stream method
// 
// Requirements:
// 1. Create a buffered channel: make(chan interface{}, 10)
// 2. Launch goroutine with defer close(out)
// 3. Convert input string to runes (for Unicode support)
// 4. For each rune:
//    - Check if context is cancelled (ctx.Done())
//    - Send rune to channel
//    - Sleep for s.Delay
// 5. Return channel
//
// Template:
// func (s *StreamingTextRunnable) Stream(ctx context.Context, input interface{}, config *core.Config) (<-chan interface{}, error) {
//     // YOUR CODE HERE
// }

func main() {
	// Create streaming runnable with 100ms delay
	streamer := NewStreamingTextRunnable(100 * time.Millisecond)
	
	ctx := context.Background()
	
	// TODO: Call Stream and print each character
	// stream, err := streamer.Stream(ctx, "Hello, Go!", nil)
	// if err != nil {
	//     fmt.Printf("Error: %v\n", err)
	//     return
	// }
	//
	// for char := range stream {
	//     fmt.Print(char)
	// }
	// fmt.Println()
}
