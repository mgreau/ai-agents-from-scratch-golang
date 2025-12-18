package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// StreamingTextRunnable streams text character by character - SOLUTION
type StreamingTextRunnableSolution struct {
	*core.BaseRunnable
	Delay time.Duration
}

func NewStreamingTextRunnableSolution(delay time.Duration) *StreamingTextRunnableSolution {
	return &StreamingTextRunnableSolution{
		BaseRunnable: core.NewBaseRunnable("StreamingText"),
		Delay:        delay,
	}
}

// Stream sends each character individually
func (s *StreamingTextRunnableSolution) Stream(ctx context.Context, input interface{}, config *core.Config) (<-chan interface{}, error) {
	str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("input must be string, got %T", input)
	}
	
	// Create buffered channel
	out := make(chan interface{}, 10)
	
	go func() {
		defer close(out)
		
		// Convert to runes for proper Unicode handling
		runes := []rune(str)
		
		for _, r := range runes {
			select {
			case <-ctx.Done():
				// Context cancelled, stop streaming
				return
			case out <- string(r):
				// Character sent successfully
			}
			
			// Add delay between characters
			if s.Delay > 0 {
				time.Sleep(s.Delay)
			}
		}
	}()
	
	return out, nil
}

func runSolution() {
	// Test 1: Basic streaming
	fmt.Println("=== Test 1: Basic Streaming ===")
	streamer := NewStreamingTextRunnableSolution(50 * time.Millisecond)
	
	ctx := context.Background()
	stream, err := streamer.Stream(ctx, "Hello!", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Print("Streaming: ")
	for char := range stream {
		fmt.Print(char)
	}
	fmt.Println()
	
	// Test 2: With Unicode
	fmt.Println("\n=== Test 2: Unicode Support ===")
	stream, _ = streamer.Stream(ctx, "Hello 世界! 🎉", nil)
	
	fmt.Print("Unicode: ")
	for char := range stream {
		fmt.Print(char)
	}
	fmt.Println()
	
	// Test 3: With cancellation
	fmt.Println("\n=== Test 3: Context Cancellation ===")
	ctxCancel, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	
	stream, _ = streamer.Stream(ctxCancel, "This is a long text that will be interrupted", nil)
	
	fmt.Print("Cancelled: ")
	for char := range stream {
		fmt.Print(char)
	}
	fmt.Println(" [STOPPED]")
	
	// Test 4: Fast streaming (no delay)
	fmt.Println("\n=== Test 4: Fast Streaming (No Delay) ===")
	fastStreamer := NewStreamingTextRunnableSolution(0)
	stream, _ = fastStreamer.Stream(context.Background(), "Fast!", nil)
	
	fmt.Print("Fast: ")
	start := time.Now()
	for char := range stream {
		fmt.Print(char)
	}
	fmt.Printf(" (took %v)\n", time.Since(start))
	
	// Test 5: Collect streaming output
	fmt.Println("\n=== Test 5: Collecting Output ===")
	stream, _ = streamer.Stream(context.Background(), "Collect", nil)
	
	var collected string
	for char := range stream {
		collected += char.(string)
	}
	fmt.Printf("Collected: '%s'\n", collected)
}
