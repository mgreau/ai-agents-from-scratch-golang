package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// JSONParserRunnable parses JSON strings into Go data structures
type JSONParserRunnable struct {
	*core.BaseRunnable
}

// NewJSONParserRunnable creates a new JSON parser
func NewJSONParserRunnable() *JSONParserRunnable {
	return &JSONParserRunnable{
		BaseRunnable: core.NewBaseRunnable("JSONParser"),
	}
}

// TODO: Implement the Invoke method
// 
// Requirements:
// 1. Accept input as string or []byte
// 2. Parse JSON using json.Unmarshal
// 3. Return map[string]interface{} with parsed data
// 4. Return error for invalid JSON or wrong input type
//
// Example:
//   input: `{"name": "John", "age": 30}`
//   output: map[string]interface{}{"name": "John", "age": 30}
//
//   input: `{invalid}`
//   output: error("invalid JSON")

// YOUR CODE HERE

func main() {
	parser := NewJSONParserRunnable()
	
	ctx := context.Background()
	
	// Test with valid JSON
	result, err := parser.Invoke(ctx, `{"message": "Hello, Go!"}`, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Parsed: %+v\n", result)
	// Expected: Parsed: map[message:Hello, Go!]
}
