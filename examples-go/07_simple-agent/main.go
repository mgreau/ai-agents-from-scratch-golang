package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/llm"
	"github.com/mgreau/ai-agents-from-scratch-go/pkg/tools"
)

func main() {
	// Get model path
	modelPath, err := filepath.Abs("../../models/Qwen3-1.7B-Q8_0.gguf")
	if err != nil {
		log.Fatalf("Failed to get model path: %v", err)
	}

	// Create tool registry
	registry := tools.NewToolRegistry()
	registry.Register(tools.NewGetCurrentTimeTool())
	registry.Register(tools.NewCalculatorTool())

	// Create LLM instance with system prompt
	systemPrompt := `You are a helpful assistant with access to tools.
When you need to use a tool, respond with the tool name and its arguments.
Always convert times from 12-hour format to 24-hour format without seconds.`

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

	// User query
	prompt := "What time is it right now?"

	// Get response
	ctx := context.Background()
	response, err := llamaLLM.Invoke(ctx, prompt, nil)
	if err != nil {
		log.Fatalf("Failed to get response: %v", err)
	}

	fmt.Printf("User: %s\n", prompt)
	fmt.Printf("AI: %s\n", response)

	// In a real implementation, you would:
	// 1. Parse the response to detect tool calls
	// 2. Execute the requested tools
	// 3. Send tool results back to the LLM
	// 4. Get the final response

	// For demonstration, let's manually call the tool
	fmt.Println("\n--- Tool Execution ---")
	timeResult, err := registry.ExecuteTool(ctx, "getCurrentTime", "")
	if err != nil {
		log.Fatalf("Failed to execute tool: %v", err)
	}
	fmt.Printf("Tool result: %s\n", timeResult)
}
