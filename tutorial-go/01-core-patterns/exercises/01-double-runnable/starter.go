package main

import (
	"context"
	"fmt"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// DoubleRunnable doubles numeric input
type DoubleRunnable struct {
	*core.BaseRunnable
}

// NewDoubleRunnable creates a new DoubleRunnable
func NewDoubleRunnable() *DoubleRunnable {
	return &DoubleRunnable{
		BaseRunnable: core.NewBaseRunnable("DoubleRunnable"),
	}
}

// TODO: Implement the Invoke method
// Hint: Override the call method from BaseRunnable
// This is where you implement the actual doubling logic
//
// The method signature should be:
// func (d *DoubleRunnable) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error)
//
// Steps:
// 1. Check if input is a number (int or float64)
// 2. Double it
// 3. Return the result
// 4. Return an error if input is not a number
//
// Example:
//   input: 5
//   output: 10
//
//   input: "hello"
//   output: error("input must be a number")

// YOUR CODE HERE

func main() {
	// Example usage
	doubler := NewDoubleRunnable()
	
	ctx := context.Background()
	result, err := doubler.Invoke(ctx, 5, nil)
	
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Result: %v\n", result)
	// Expected: Result: 10
}
