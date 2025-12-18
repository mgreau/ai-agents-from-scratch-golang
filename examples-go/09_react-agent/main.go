package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/agents"
	"github.com/mgreau/ai-agents-from-scratch-go/pkg/llm"
	"github.com/mgreau/ai-agents-from-scratch-go/pkg/tools"
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

	// Create tool registry with various tools
	registry := tools.NewToolRegistry()
	registry.Register(tools.NewCalculatorTool())
	registry.Register(tools.NewGetCurrentTimeTool())

	// Create ReAct agent
	agent := agents.NewReActAgent(llamaLLM, registry, 5, true)

	// Run agent with a complex query
	ctx := context.Background()
	query := "What is 15 multiplied by 23?"

	answer, err := agent.Run(ctx, query)
	if err != nil {
		log.Fatalf("Agent failed: %v", err)
	}

	fmt.Printf("\n\nFinal Result: %s\n", answer)
}
