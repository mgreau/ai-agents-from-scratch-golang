package main

import (
	"fmt"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// ValidationOptions controls validation behavior - SOLUTION
type ValidationOptionsSolution struct {
	RequireSystemFirst bool
	RequireHuman       bool
	NoConsecutiveSame  bool
}

// ValidateConversation validates message sequence - SOLUTION
func ValidateConversationSolution(messages []core.Message, opts ValidationOptionsSolution) error {
	// Check if empty
	if len(messages) == 0 {
		return fmt.Errorf("conversation is empty")
	}
	
	// Check first message is System
	if opts.RequireSystemFirst && messages[0].GetType() != core.MessageTypeSystem {
		return fmt.Errorf("first message must be System, got %s", messages[0].GetType())
	}
	
	// Check for at least one Human message
	if opts.RequireHuman {
		hasHuman := false
		for _, msg := range messages {
			if msg.GetType() == core.MessageTypeHuman {
				hasHuman = true
				break
			}
		}
		if !hasHuman {
			return fmt.Errorf("conversation must have at least one Human message")
		}
	}
	
	// Check for consecutive same-type messages
	if opts.NoConsecutiveSame && len(messages) > 1 {
		for i := 1; i < len(messages); i++ {
			prevType := messages[i-1].GetType()
			currType := messages[i].GetType()
			
			if prevType == currType {
				return fmt.Errorf("consecutive %s messages at positions %d and %d", currType, i-1, i)
			}
		}
	}
	
	// Check Tool messages follow AI messages
	for i, msg := range messages {
		if msg.GetType() == core.MessageTypeTool {
			if i == 0 {
				return fmt.Errorf("Tool message cannot be first")
			}
			
			// Previous message should be AI with tool calls
			prevMsg := messages[i-1]
			if prevMsg.GetType() != core.MessageTypeAI {
				return fmt.Errorf("Tool message at position %d must follow AI message", i)
			}
			
			// Check if AI message had tool calls
			if aiMsg, ok := prevMsg.(*core.AIMessage); ok {
				if !aiMsg.HasToolCalls() {
					return fmt.Errorf("Tool message at position %d follows AI without tool calls", i)
				}
			}
		}
	}
	
	return nil
}

func runSolution() {
	opts := ValidationOptionsSolution{
		RequireSystemFirst: true,
		RequireHuman:       true,
		NoConsecutiveSame:  true,
	}
	
	// Test 1: Valid conversation
	fmt.Println("=== Test 1: Valid Conversation ===")
	valid := []core.Message{
		core.NewSystemMessage("You are helpful", nil),
		core.NewHumanMessage("Hello", nil),
		core.NewAIMessage("Hi!", nil),
	}
	
	if err := ValidateConversationSolution(valid, opts); err != nil {
		fmt.Printf("✗ %v\n", err)
	} else {
		fmt.Println("✓ Valid")
	}
	
	// Test 2: No system message first
	fmt.Println("\n=== Test 2: No System First ===")
	noSystem := []core.Message{
		core.NewHumanMessage("Hello", nil),
		core.NewAIMessage("Hi!", nil),
	}
	
	if err := ValidateConversationSolution(noSystem, opts); err != nil {
		fmt.Printf("✗ Expected error: %v\n", err)
	}
	
	// Test 3: No human message
	fmt.Println("\n=== Test 3: No Human Message ===")
	noHuman := []core.Message{
		core.NewSystemMessage("You are helpful", nil),
		core.NewAIMessage("Hello!", nil),
	}
	
	if err := ValidateConversationSolution(noHuman, opts); err != nil {
		fmt.Printf("✗ Expected error: %v\n", err)
	}
	
	// Test 4: Consecutive same-type messages
	fmt.Println("\n=== Test 4: Consecutive Same Type ===")
	consecutive := []core.Message{
		core.NewSystemMessage("You are helpful", nil),
		core.NewHumanMessage("Hello", nil),
		core.NewHumanMessage("Are you there?", nil),
	}
	
	if err := ValidateConversationSolution(consecutive, opts); err != nil {
		fmt.Printf("✗ Expected error: %v\n", err)
	}
	
	// Test 5: Valid tool flow
	fmt.Println("\n=== Test 5: Valid Tool Flow ===")
	withTools := []core.Message{
		core.NewSystemMessage("You have a calculator", nil),
		core.NewHumanMessage("What's 5 + 3?", nil),
		core.NewAIMessage("Calculating...", map[string]interface{}{
			"tool_calls": []core.ToolCall{{
				ID:   "call_123",
				Type: "function",
				Function: core.ToolCallFunction{
					Name:      "calculator",
					Arguments: `{"a": 5, "b": 3}`,
				},
			}},
		}),
		core.NewToolMessage("8", "call_123", nil),
		core.NewAIMessage("The result is 8", nil),
	}
	
	if err := ValidateConversationSolution(withTools, opts); err != nil {
		fmt.Printf("✗ %v\n", err)
	} else {
		fmt.Println("✓ Valid tool flow")
	}
	
	// Test 6: Lenient validation
	fmt.Println("\n=== Test 6: Lenient Validation ===")
	lenientOpts := ValidationOptionsSolution{
		RequireSystemFirst: false,
		RequireHuman:       false,
		NoConsecutiveSame:  false,
	}
	
	lenient := []core.Message{
		core.NewAIMessage("I can start conversations", nil),
		core.NewAIMessage("And send multiple messages", nil),
	}
	
	if err := ValidateConversationSolution(lenient, lenientOpts); err != nil {
		fmt.Printf("✗ %v\n", err)
	} else {
		fmt.Println("✓ Valid with lenient options")
	}
}
