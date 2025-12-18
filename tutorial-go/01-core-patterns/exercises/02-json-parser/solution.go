package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// JSONParserRunnable parses JSON strings into Go data structures - SOLUTION
type JSONParserRunnableSolution struct {
	*core.BaseRunnable
}

// NewJSONParserRunnableSolution creates a new JSON parser
func NewJSONParserRunnableSolution() *JSONParserRunnableSolution {
	return &JSONParserRunnableSolution{
		BaseRunnable: core.NewBaseRunnable("JSONParser"),
	}
}

// Invoke parses JSON input
func (j *JSONParserRunnableSolution) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
	var jsonBytes []byte
	
	// Handle different input types
	switch v := input.(type) {
	case string:
		jsonBytes = []byte(v)
	case []byte:
		jsonBytes = v
	default:
		return nil, fmt.Errorf("input must be string or []byte, got %T", input)
	}
	
	// Parse JSON
	var result map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	
	return result, nil
}

func runSolution() {
	parser := NewJSONParserRunnableSolution()
	
	ctx := context.Background()
	
	// Test 1: Valid JSON
	result, err := parser.Invoke(ctx, `{"name": "Alice", "age": 25, "active": true}`, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Test 1 - Parsed: %+v\n", result)
	
	// Test 2: Valid JSON with []byte input
	result, err = parser.Invoke(ctx, []byte(`{"city": "Paris", "country": "France"}`), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Test 2 - Parsed: %+v\n", result)
	
	// Test 3: Invalid JSON
	result, err = parser.Invoke(ctx, `{invalid json}`, nil)
	if err != nil {
		fmt.Printf("Test 3 - Expected error: %v\n", err)
	}
	
	// Test 4: Wrong input type
	result, err = parser.Invoke(ctx, 12345, nil)
	if err != nil {
		fmt.Printf("Test 4 - Expected error: %v\n", err)
	}
	
	// Test 5: Use in a pipeline
	// JSON Parser -> Extract field
	type FieldExtractorRunnable struct {
		*core.BaseRunnable
		Field string
	}
	
	extractor := &FieldExtractorRunnable{
		BaseRunnable: core.NewBaseRunnable("FieldExtractor"),
		Field:        "name",
	}
	
	// Pipeline: parse JSON then extract name field
	pipeline := parser.Pipe(extractor)
	
	result, err = pipeline.Invoke(ctx, `{"name": "Bob", "age": 30}`, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Test 5 - Pipeline result: %v\n", result)
}
