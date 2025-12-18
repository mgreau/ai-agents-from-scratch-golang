package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
	"github.com/mgreau/ai-agents-from-scratch-go/pkg/llm"
)

func main() {
	// Get model path
	modelPath, err := filepath.Abs("../../models/Qwen3-1.7B-Q8_0.gguf")
	if err != nil {
		log.Fatalf("Failed to get model path: %v", err)
	}

	// Create LLM instance
	systemPrompt := `You are a concise assistant. Provide brief, one-sentence answers.`

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

	// Multiple prompts to process in parallel
	prompts := []interface{}{
		"What is the capital of France?",
		"What is 2 + 2?",
		"What color is the sky?",
		"Who wrote Romeo and Juliet?",
		"What is the largest planet in our solar system?",
	}

	fmt.Println("Processing multiple queries in parallel using Batch...\n")

	// Process in batch (parallel execution)
	ctx := context.Background()
	start := time.Now()
	
	results, err := llamaLLM.Batch(ctx, prompts, nil)
	if err != nil {
		log.Fatalf("Batch processing failed: %v", err)
	}

	elapsed := time.Since(start)

	// Display results
	fmt.Println("=== Batch Results ===")
	for i, result := range results {
		fmt.Printf("\nQ%d: %s\n", i+1, prompts[i])
		fmt.Printf("A%d: %s\n", i+1, result)
	}

	fmt.Printf("\nTotal time: %v\n", elapsed)
	fmt.Printf("Average per query: %v\n", elapsed/time.Duration(len(prompts)))

	// Compare with sequential processing
	fmt.Println("\n--- Comparing with Sequential Processing ---")
	start = time.Now()
	
	for i, prompt := range prompts {
		_, err := llamaLLM.Invoke(ctx, prompt, nil)
		if err != nil {
			log.Printf("Sequential query %d failed: %v", i, err)
		}
	}
	
	seqElapsed := time.Since(start)
	fmt.Printf("Sequential time: %v\n", seqElapsed)
	fmt.Printf("Speedup: %.2fx faster\n", float64(seqElapsed)/float64(elapsed))
}

// Note: Actual speedup depends on:
// 1. Hardware capabilities (CPU cores, memory bandwidth)
// 2. Model size and quantization
// 3. Context size
// 4. Number of parallel requests
// In practice, Go's goroutines enable efficient parallel processing
