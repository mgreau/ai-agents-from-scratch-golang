package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// MockLLMRunnable - SOLUTION
type MockLLMRunnableSolution struct {
	*core.BaseRunnable
	temperature float32
	modelName   string
	responses   []string
}

func NewMockLLMRunnableSolution(temperature float32) *MockLLMRunnableSolution {
	return &MockLLMRunnableSolution{
		BaseRunnable: core.NewBaseRunnable("MockLLM"),
		temperature:  temperature,
		responses: []string{
			"That's an interesting question!",
			"Let me help you with that.",
			"Here's what I think...",
			"Based on my knowledge...",
			"Great question! The answer is...",
		},
	}
}

func (m *MockLLMRunnableSolution) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
	var prompt string
	
	// Handle different input types
	switch v := input.(type) {
	case string:
		prompt = v
	case []core.Message:
		// Convert messages to prompt
		var parts []string
		for _, msg := range v {
			parts = append(parts, fmt.Sprintf("%s: %s", msg.GetType(), msg.GetContent()))
		}
		prompt = strings.Join(parts, "\n")
	default:
		return nil, fmt.Errorf("input must be string or []core.Message, got %T", input)
	}
	
	// Generate mock response based on temperature
	var response string
	if m.temperature > 0.5 {
		// More random
		response = m.responses[rand.Intn(len(m.responses))]
	} else {
		// More deterministic
		response = m.responses[0]
	}
	
	// Add context from prompt
	response += fmt.Sprintf(" Regarding: %s", prompt[:min(50, len(prompt))])
	
	return response, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func runSolution() {
	llm := NewMockLLMRunnableSolution(0.7)
	ctx := context.Background()
	
	// Test 1: String input
	fmt.Println("=== Test 1: String Input ===")
	resp1, _ := llm.Invoke(ctx, "What is Go?", nil)
	fmt.Printf("Response: %s\n", resp1)
	
	// Test 2: Message input
	fmt.Println("\n=== Test 2: Message Input ===")
	messages := []core.Message{
		core.NewSystemMessage("You are helpful", nil),
		core.NewHumanMessage("Explain channels", nil),
	}
	resp2, _ := llm.Invoke(ctx, messages, nil)
	fmt.Printf("Response: %s\n", resp2)
	
	// Test 3: In pipeline
	fmt.Println("\n=== Test 3: In Pipeline ===")
	type UppercaseRunnable struct {
		*core.BaseRunnable
	}
	upper := &UppercaseRunnable{core.NewBaseRunnable("Upper")}
	
	pipeline := llm.Pipe(upper)
	resp3, _ := pipeline.Invoke(ctx, "test", nil)
	fmt.Printf("Pipeline: %s\n", resp3)
}
