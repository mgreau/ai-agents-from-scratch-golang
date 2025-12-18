package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// TrimRunnable removes leading/trailing whitespace - SOLUTION
type TrimRunnableSolution struct {
	*core.BaseRunnable
}

func NewTrimRunnableSolution() *TrimRunnableSolution {
	return &TrimRunnableSolution{
		BaseRunnable: core.NewBaseRunnable("Trim"),
	}
}

func (t *TrimRunnableSolution) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
	str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("input must be string, got %T", input)
	}
	return strings.TrimSpace(str), nil
}

// UppercaseRunnable converts text to uppercase - SOLUTION
type UppercaseRunnableSolution struct {
	*core.BaseRunnable
}

func NewUppercaseRunnableSolution() *UppercaseRunnableSolution {
	return &UppercaseRunnableSolution{
		BaseRunnable: core.NewBaseRunnable("Uppercase"),
	}
}

func (u *UppercaseRunnableSolution) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
	str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("input must be string, got %T", input)
	}
	return strings.ToUpper(str), nil
}

// PrefixRunnable adds a prefix to text - SOLUTION
type PrefixRunnableSolution struct {
	*core.BaseRunnable
	Prefix string
}

func NewPrefixRunnableSolution(prefix string) *PrefixRunnableSolution {
	return &PrefixRunnableSolution{
		BaseRunnable: core.NewBaseRunnable("Prefix"),
		Prefix:       prefix,
	}
}

func (p *PrefixRunnableSolution) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
	str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("input must be string, got %T", input)
	}
	return p.Prefix + str, nil
}

func runSolution() {
	// Create components
	trim := NewTrimRunnableSolution()
	upper := NewUppercaseRunnableSolution()
	prefix := NewPrefixRunnableSolution(">> ")
	
	ctx := context.Background()
	
	// Test 1: Individual components
	fmt.Println("=== Test 1: Individual Components ===")
	r1, _ := trim.Invoke(ctx, "  hello  ", nil)
	fmt.Printf("After Trim: '%s'\n", r1)
	
	r2, _ := upper.Invoke(ctx, "hello", nil)
	fmt.Printf("After Upper: '%s'\n", r2)
	
	r3, _ := prefix.Invoke(ctx, "hello", nil)
	fmt.Printf("After Prefix: '%s'\n", r3)
	
	// Test 2: Two-step pipeline
	fmt.Println("\n=== Test 2: Two-Step Pipeline ===")
	pipeline2 := trim.Pipe(upper)
	result2, _ := pipeline2.Invoke(ctx, "  hello world  ", nil)
	fmt.Printf("Trim->Upper: '%s'\n", result2)
	
	// Test 3: Three-step pipeline
	fmt.Println("\n=== Test 3: Three-Step Pipeline ===")
	pipeline3 := trim.Pipe(upper).Pipe(prefix)
	result3, _ := pipeline3.Invoke(ctx, "  hello world  ", nil)
	fmt.Printf("Trim->Upper->Prefix: '%s'\n", result3)
	
	// Test 4: Different inputs
	fmt.Println("\n=== Test 4: Different Inputs ===")
	tests := []string{
		"  go is awesome  ",
		"composability",
		"  MIXED Case  ",
	}
	
	for _, test := range tests {
		result, _ := pipeline3.Invoke(ctx, test, nil)
		fmt.Printf("'%s' => '%s'\n", test, result)
	}
	
	// Test 5: Batch processing
	fmt.Println("\n=== Test 5: Batch Processing ===")
	inputs := []interface{}{
		"  first  ",
		"  second  ",
		"  third  ",
	}
	results, _ := pipeline3.Batch(ctx, inputs, nil)
	for i, r := range results {
		fmt.Printf("Batch %d: %s\n", i+1, r)
	}
	
	// Test 6: Different pipeline order
	fmt.Println("\n=== Test 6: Different Order ===")
	// Try: Prefix -> Uppercase -> Trim (different order, different result!)
	pipeline4 := prefix.Pipe(upper).Pipe(trim)
	result4, _ := pipeline4.Invoke(ctx, "  hello  ", nil)
	fmt.Printf("Prefix->Upper->Trim: '%s'\n", result4)
}
