package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// TrimRunnable removes leading/trailing whitespace
type TrimRunnable struct {
	*core.BaseRunnable
}

func NewTrimRunnable() *TrimRunnable {
	return &TrimRunnable{
		BaseRunnable: core.NewBaseRunnable("Trim"),
	}
}

// TODO: Implement Invoke for TrimRunnable
// Use strings.TrimSpace()

// UppercaseRunnable converts text to uppercase
type UppercaseRunnable struct {
	*core.BaseRunnable
}

func NewUppercaseRunnable() *UppercaseRunnable {
	return &UppercaseRunnable{
		BaseRunnable: core.NewBaseRunnable("Uppercase"),
	}
}

// TODO: Implement Invoke for UppercaseRunnable
// Use strings.ToUpper()

// PrefixRunnable adds a prefix to text
type PrefixRunnable struct {
	*core.BaseRunnable
	Prefix string
}

func NewPrefixRunnable(prefix string) *PrefixRunnable {
	return &PrefixRunnable{
		BaseRunnable: core.NewBaseRunnable("Prefix"),
		Prefix:       prefix,
	}
}

// TODO: Implement Invoke for PrefixRunnable
// Return: p.Prefix + input

func main() {
	// Create components
	trim := NewTrimRunnable()
	upper := NewUppercaseRunnable()
	prefix := NewPrefixRunnable(">> ")
	
	// TODO: Compose into pipeline using Pipe()
	// pipeline := trim.Pipe(upper).Pipe(prefix)
	
	// TODO: Test the pipeline
	ctx := context.Background()
	// result, _ := pipeline.Invoke(ctx, "  hello world  ", nil)
	// fmt.Printf("Result: %s\n", result)
	// Expected: ">> HELLO WORLD"
}
