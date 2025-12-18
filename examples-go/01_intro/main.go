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

	// Create LLM instance
	llamaLLM, err := llm.NewLlamaCppLLM(llm.LlamaCppConfig{
		ModelPath:   modelPath,
		ContextSize: 2048,
		Temperature: 0.7,
		Threads:     4,
	})
	if err != nil {
		log.Fatalf("Failed to create LLM: %v", err)
	}
	defer llamaLLM.Close()

	// Create prompt
	prompt := "do you know node-llama-cpp"

	// Get response
	ctx := context.Background()
	response, err := llamaLLM.Invoke(ctx, prompt, nil)
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}

	fmt.Printf("AI: %s\n", response)
}
