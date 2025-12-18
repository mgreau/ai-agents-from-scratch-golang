package main

import (
	"encoding/json"
	"fmt"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// ToolFlowHandler manages tool call execution flow
type ToolFlowHandler struct {
	tools map[string]ToolFunc
}

// ToolFunc is a function that executes a tool
type ToolFunc func(args map[string]interface{}) (string, error)

// NewToolFlowHandler creates a new handler
func NewToolFlowHandler() *ToolFlowHandler {
	handler := &ToolFlowHandler{
		tools: make(map[string]ToolFunc),
	}
	
	// Register some mock tools
	handler.RegisterTool("calculator", mockCalculator)
	handler.RegisterTool("get_weather", mockWeather)
	
	return handler
}

// RegisterTool adds a tool
func (h *ToolFlowHandler) RegisterTool(name string, fn ToolFunc) {
	h.tools[name] = fn
}

// HandleToolCalls processes tool calls in conversation
// TODO: Implement this method
//
// Requirements:
// 1. Find the last AI message
// 2. Check if it has tool calls (aiMsg.HasToolCalls())
// 3. For each tool call:
//    a. Parse arguments JSON
//    b. Execute tool function
//    c. Create ToolMessage with result and tool call ID
//    d. Append to conversation
// 4. Return updated conversation

func (h *ToolFlowHandler) HandleToolCalls(messages []core.Message) []core.Message {
	// YOUR CODE HERE
	return messages
}

// Mock tool functions
func mockCalculator(args map[string]interface{}) (string, error) {
	a := args["a"].(float64)
	b := args["b"].(float64)
	op := args["operation"].(string)
	
	var result float64
	switch op {
	case "add":
		result = a + b
	case "multiply":
		result = a * b
	default:
		return "", fmt.Errorf("unknown operation: %s", op)
	}
	
	return fmt.Sprintf("%.0f", result), nil
}

func mockWeather(args map[string]interface{}) (string, error) {
	location := args["location"].(string)
	return fmt.Sprintf(`{"location": "%s", "temperature": 72, "condition": "sunny"}`, location), nil
}

func main() {
	handler := NewToolFlowHandler()
	
	// Simulate conversation with tool call
	messages := []core.Message{
		core.NewSystemMessage("You have access to tools", nil),
		core.NewHumanMessage("What's 15 + 23?", nil),
		core.NewAIMessage("Let me calculate that", map[string]interface{}{
			"tool_calls": []core.ToolCall{{
				ID:   "call_123",
				Type: "function",
				Function: core.ToolCallFunction{
					Name:      "calculator",
					Arguments: `{"operation": "add", "a": 15, "b": 23}`,
				},
			}},
		}),
	}
	
	// Process tool calls
	updated := handler.HandleToolCalls(messages)
	
	fmt.Printf("Updated conversation has %d messages\n", len(updated))
	for _, msg := range updated {
		fmt.Printf("  %s: %s\n", msg.GetType(), msg.GetContent())
	}
}
