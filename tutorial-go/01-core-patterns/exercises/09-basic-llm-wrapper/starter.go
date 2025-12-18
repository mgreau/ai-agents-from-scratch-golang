package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// MockLLMConfig holds configuration
type MockLLMConfig struct {
	Temperature float32
	ModelName   string
}

// MockLLMRunnable is a mock LLM for testing
type MockLLMRunnable struct {
	*core.BaseRunnable
	temperature float32
	modelName   string
}

// NewMockLLMRunnable creates a new mock LLM
func NewMockLLMRunnable(config MockLLMConfig) *MockLLMRunnable {
	if config.Temperature == 0 {
		config.Temperature = 0.7
	}
	if config.ModelName == "" {
		config.ModelName = "mock-llm"
	}
	
	return &MockLLMRunnable{
		BaseRunnable: core.NewBaseRunnable("MockLLM"),
		temperature:  config.Temperature,
		modelName:    config.ModelName,
	}
}

// TODO: Implement Invoke method
//
// Requirements:
// 1. Accept input as string or []core.Message
// 2. If []Message, convert to simple prompt string
// 3. Return a mock response (can be random or based on input)
// 4. Temperature affects randomness (higher = more random)
// 5. Return error for invalid input types

func (m *MockLLMRunnable) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
	// YOUR CODE HERE
	return "", nil
}

func main() {
	llm := NewMockLLMRunnable(MockLLMConfig{
		Temperature: 0.7,
	})
	
	ctx := context.Background()
	
	// Test with string
	response, _ := llm.Invoke(ctx, "What is Go?", nil)
	fmt.Printf("Response: %s\n", response)
}
