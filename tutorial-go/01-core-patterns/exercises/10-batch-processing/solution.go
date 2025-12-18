package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// BatchLLMRunnable - SOLUTION
type BatchLLMRunnableSolution struct {
	*core.BaseRunnable
}

func NewBatchLLMRunnableSolution() *BatchLLMRunnableSolution {
	return &BatchLLMRunnableSolution{
		BaseRunnable: core.NewBaseRunnable("BatchLLM"),
	}
}

func (b *BatchLLMRunnableSolution) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
	time.Sleep(100 * time.Millisecond) // Simulate LLM call
	return fmt.Sprintf("Response to: %v", input), nil
}

func (b *BatchLLMRunnableSolution) Batch(ctx context.Context, inputs []interface{}, config *core.Config) ([]interface{}, error) {
	results := make([]interface{}, len(inputs))
	var wg sync.WaitGroup
	var mu sync.Mutex
	
	for i, input := range inputs {
		wg.Add(1)
		go func(idx int, inp interface{}) {
			defer wg.Done()
			
			result, _ := b.Invoke(ctx, inp, config)
			
			mu.Lock()
			results[idx] = result
			mu.Unlock()
		}(i, input)
	}
	
	wg.Wait()
	return results, nil
}

func runSolution() {
	llm := NewBatchLLMRunnableSolution()
	ctx := context.Background()
	
	inputs := []interface{}{"Q1", "Q2", "Q3", "Q4", "Q5"}
	
	start := time.Now()
	results, _ := llm.Batch(ctx, inputs, nil)
	duration := time.Since(start)
	
	fmt.Printf("Processed %d inputs in %v\n", len(results), duration)
	for i, r := range results {
		fmt.Printf("%d: %s\n", i+1, r)
	}
}
