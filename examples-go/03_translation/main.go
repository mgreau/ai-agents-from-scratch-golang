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

	// Create LLM instance with system prompt
	systemPrompt := `You are a professional translator who specializes in technical documentation.
	
Always translate text accurately while preserving technical terms and maintaining the original tone.
Provide only the translation without any additional explanation.`

	llamaLLM, err := llm.NewLlamaCppLLM(llm.LlamaCppConfig{
		ModelPath:    modelPath,
		ContextSize:  2048,
		Temperature:  0.3, // Lower temperature for more consistent translations
		Threads:      4,
		SystemPrompt: systemPrompt,
	})
	if err != nil {
		log.Fatalf("Failed to create LLM: %v", err)
	}
	defer llamaLLM.Close()

	// Create translation prompt
	prompt := "Translate the following English text to French: 'Hello, how are you today?'"

	// Get response
	ctx := context.Background()
	response, err := llamaLLM.Invoke(ctx, prompt, nil)
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}

	fmt.Printf("Translation: %s\n", response)
}
