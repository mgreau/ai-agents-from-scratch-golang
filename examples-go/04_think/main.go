package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/llm"
)

func main() {
	// Get model path
	modelPath, err := filepath.Abs("../../models/Qwen3-1.7B-Q8_0.gguf")
	if err != nil {
		log.Fatalf("Failed to get model path: %v", err)
	}

	// Create LLM instance optimized for reasoning
	systemPrompt := `You are a helpful assistant skilled at logical reasoning and problem solving.
Break down complex problems step by step and show your reasoning process.`

	llamaLLM, err := llm.NewLlamaCppLLM(llm.LlamaCppConfig{
		ModelPath:    modelPath,
		ContextSize:  2048,
		Temperature:  0.3, // Lower temperature for more focused reasoning
		Threads:      4,
		SystemPrompt: systemPrompt,
	})
	if err != nil {
		log.Fatalf("Failed to create LLM: %v", err)
	}
	defer llamaLLM.Close()

	// Complex reasoning problem
	prompt := `Solve this problem step by step:

If a train travels 120 km in 2 hours, and then travels another 90 km in 1.5 hours, 
what is its average speed for the entire journey?

Show your reasoning.`

	// Get response
	ctx := context.Background()
	response, err := llamaLLM.Invoke(ctx, prompt, nil)
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}

	fmt.Printf("Problem: %s\n\n", prompt)
	fmt.Printf("AI Response:\n%s\n", response)

	// Note: Pure LLM reasoning has limitations
	// For accurate calculations, you should use tools (like calculator)
	// See example 07_simple-agent for tool usage
}
