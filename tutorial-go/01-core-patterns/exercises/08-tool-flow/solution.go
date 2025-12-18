package main

import (
	"encoding/json"
	"fmt"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// ToolFlowHandlerSolution manages tool call execution - SOLUTION
type ToolFlowHandlerSolution struct {
	tools map[string]ToolFuncSolution
}

// ToolFuncSolution is a function that executes a tool
type ToolFuncSolution func(args map[string]interface{}) (string, error)

// NewToolFlowHandlerSolution creates a new handler
func NewToolFlowHandlerSolution() *ToolFlowHandlerSolution {
	handler := &ToolFlowHandlerSolution{
		tools: make(map[string]ToolFuncSolution),
	}
	
	// Register mock tools
	handler.RegisterTool("calculator", mockCalculatorSolution)
	handler.RegisterTool("get_weather", mockWeatherSolution)
	handler.RegisterTool("get_time", mockTimeSolution)
	
	return handler
}

// RegisterTool adds a tool
func (h *ToolFlowHandlerSolution) RegisterTool(name string, fn ToolFuncSolution) {
	h.tools[name] = fn
}

// HandleToolCalls processes tool calls in conversation
func (h *ToolFlowHandlerSolution) HandleToolCalls(messages []core.Message) []core.Message {
	// Find the last message
	if len(messages) == 0 {
		return messages
	}
	
	lastMsg := messages[len(messages)-1]
	
	// Check if it's an AI message with tool calls
	aiMsg, ok := lastMsg.(*core.AIMessage)
	if !ok || !aiMsg.HasToolCalls() {
		return messages
	}
	
	// Process each tool call
	for _, toolCall := range aiMsg.ToolCalls {
		// Parse arguments
		var args map[string]interface{}
		if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args); err != nil {
			// Add error message
			errMsg := core.NewToolMessage(
				fmt.Sprintf("Error parsing arguments: %v", err),
				toolCall.ID,
				nil,
			)
			messages = append(messages, errMsg)
			continue
		}
		
		// Execute tool
		toolFn, exists := h.tools[toolCall.Function.Name]
		if !exists {
			// Tool not found
			errMsg := core.NewToolMessage(
				fmt.Sprintf("Unknown tool: %s", toolCall.Function.Name),
				toolCall.ID,
				nil,
			)
			messages = append(messages, errMsg)
			continue
		}
		
		result, err := toolFn(args)
		if err != nil {
			// Execution error
			errMsg := core.NewToolMessage(
				fmt.Sprintf("Execution error: %v", err),
				toolCall.ID,
				nil,
			)
			messages = append(messages, errMsg)
			continue
		}
		
		// Add successful result
		toolMsg := core.NewToolMessage(result, toolCall.ID, nil)
		messages = append(messages, toolMsg)
	}
	
	return messages
}

// Mock tools
func mockCalculatorSolution(args map[string]interface{}) (string, error) {
	a := args["a"].(float64)
	b := args["b"].(float64)
	op := args["operation"].(string)
	
	var result float64
	switch op {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	case "divide":
		if b == 0 {
			return "", fmt.Errorf("division by zero")
		}
		result = a / b
	default:
		return "", fmt.Errorf("unknown operation: %s", op)
	}
	
	return fmt.Sprintf("%.2f", result), nil
}

func mockWeatherSolution(args map[string]interface{}) (string, error) {
	location := args["location"].(string)
	return fmt.Sprintf(`{"location": "%s", "temperature": 72, "condition": "sunny"}`, location), nil
}

func mockTimeSolution(args map[string]interface{}) (string, error) {
	timezone := "UTC"
	if tz, ok := args["timezone"].(string); ok {
		timezone = tz
	}
	return fmt.Sprintf(`{"timezone": "%s", "time": "14:30:00"}`, timezone), nil
}

func runSolution() {
	handler := NewToolFlowHandlerSolution()
	
	// Test 1: Single tool call
	fmt.Println("=== Test 1: Single Tool Call ===")
	messages := []core.Message{
		core.NewSystemMessage("You have tools", nil),
		core.NewHumanMessage("What's 15 + 23?", nil),
		core.NewAIMessage("Calculating...", map[string]interface{}{
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
	
	updated := handler.HandleToolCalls(messages)
	for _, msg := range updated {
		fmt.Printf("%s: %s\n", msg.GetType(), msg.GetContent())
	}
	
	// Test 2: Multiple tool calls
	fmt.Println("\n=== Test 2: Multiple Tool Calls ===")
	messages2 := []core.Message{
		core.NewSystemMessage("You have tools", nil),
		core.NewHumanMessage("What's the weather and time?", nil),
		core.NewAIMessage("Let me check both", map[string]interface{}{
			"tool_calls": []core.ToolCall{
				{
					ID:   "call_weather",
					Type: "function",
					Function: core.ToolCallFunction{
						Name:      "get_weather",
						Arguments: `{"location": "Paris"}`,
					},
				},
				{
					ID:   "call_time",
					Type: "function",
					Function: core.ToolCallFunction{
						Name:      "get_time",
						Arguments: `{"timezone": "CET"}`,
					},
				},
			},
		}),
	}
	
	updated2 := handler.HandleToolCalls(messages2)
	for _, msg := range updated2 {
		fmt.Printf("%s: %s\n", msg.GetType(), msg.GetContent())
	}
	
	// Test 3: Tool error
	fmt.Println("\n=== Test 3: Tool Error (Division by Zero) ===")
	messages3 := []core.Message{
		core.NewAIMessage("Dividing...", map[string]interface{}{
			"tool_calls": []core.ToolCall{{
				ID:   "call_div",
				Type: "function",
				Function: core.ToolCallFunction{
					Name:      "calculator",
					Arguments: `{"operation": "divide", "a": 10, "b": 0}`,
				},
			}},
		}),
	}
	
	updated3 := handler.HandleToolCalls(messages3)
	for _, msg := range updated3 {
		fmt.Printf("%s: %s\n", msg.GetType(), msg.GetContent())
	}
	
	// Test 4: Unknown tool
	fmt.Println("\n=== Test 4: Unknown Tool ===")
	messages4 := []core.Message{
		core.NewAIMessage("Using unknown tool...", map[string]interface{}{
			"tool_calls": []core.ToolCall{{
				ID:   "call_unknown",
				Type: "function",
				Function: core.ToolCallFunction{
					Name:      "unknown_tool",
					Arguments: `{}`,
				},
			}},
		}),
	}
	
	updated4 := handler.HandleToolCalls(messages4)
	for _, msg := range updated4 {
		fmt.Printf("%s: %s\n", msg.GetType(), msg.GetContent())
	}
}
