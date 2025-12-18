package main

import (
	"context"
	"fmt"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// DoubleRunnable doubles numeric input - SOLUTION
type DoubleRunnableSolution struct {
	*core.BaseRunnable
}

// NewDoubleRunnableSolution creates a new DoubleRunnable
func NewDoubleRunnableSolution() *DoubleRunnableSolution {
	return &DoubleRunnableSolution{
		BaseRunnable: core.NewBaseRunnable("DoubleRunnable"),
	}
}

// Invoke implements the Runnable interface
// Note: We override Invoke to implement our custom logic
func (d *DoubleRunnableSolution) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
	// Try to handle different numeric types
	switch v := input.(type) {
	case int:
		return v * 2, nil
	case int64:
		return v * 2, nil
	case float64:
		return v * 2, nil
	case float32:
		return v * 2, nil
	default:
		return nil, fmt.Errorf("input must be a number, got %T", input)
	}
}

// Alternative implementation using the internal call pattern
// This would be used if BaseRunnable had a call() method to override
/*
func (d *DoubleRunnableSolution) call(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
	switch v := input.(type) {
	case int:
		return v * 2, nil
	case float64:
		return v * 2, nil
	default:
		return nil, fmt.Errorf("input must be a number, got %T", input)
	}
}
*/

func runSolution() {
	doubler := NewDoubleRunnableSolution()
	
	ctx := context.Background()
	
	// Test with int
	result, err := doubler.Invoke(ctx, 5, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("5 * 2 = %v\n", result)
	
	// Test with float
	result, err = doubler.Invoke(ctx, 3.5, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("3.5 * 2 = %v\n", result)
	
	// Test with invalid input
	result, err = doubler.Invoke(ctx, "hello", nil)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
	
	// Test batch processing (inherited from BaseRunnable)
	results, err := doubler.Batch(ctx, []interface{}{1, 2, 3, 4}, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Batch result: %v\n", results)
}
