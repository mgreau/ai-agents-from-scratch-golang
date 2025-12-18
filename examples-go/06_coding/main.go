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
	systemPrompt := `You are an expert Go programmer. Provide clear, concise, and idiomatic Go code.`

	llamaLLM, err := llm.NewLlamaCppLLM(llm.LlamaCppConfig{
		ModelPath:    modelPath,
		ContextSize:  2048,
		Temperature:  0.7,
		Threads:      4,
		SystemPrompt: systemPrompt,
	})
	if err != nil {
		log.Fatalf("Failed to create LLM: %v", err)
	}
	defer llamaLLM.Close()

	// Create coding prompt
	prompt := "Write a simple Go function that calculates the factorial of a number using recursion."

	// Stream response
	ctx := context.Background()
	fmt.Println("AI: ")
	
	stream, err := llamaLLM.Stream(ctx, prompt, nil)
	if err != nil {
		log.Fatalf("Failed to stream response: %v", err)
	}

	// Print tokens as they arrive
	for token := range stream {
		if err, ok := token.(error); ok {
			log.Fatalf("Stream error: %v", err)
		}
		fmt.Print(token)
	}
	fmt.Println()
}
