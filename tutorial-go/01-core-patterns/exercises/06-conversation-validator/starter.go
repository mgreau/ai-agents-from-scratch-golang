package main

import (
	"fmt"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// ValidationOptions controls validation behavior
type ValidationOptions struct {
	RequireSystemFirst bool // System message must be first
	RequireHuman       bool // At least one human message
	NoConsecutiveSame  bool // No consecutive same-type messages
}

// ValidateConversation validates message sequence
// TODO: Implement this function
//
// Requirements:
// 1. Check if empty (optional error)
// 2. If opts.RequireSystemFirst, check first message is System
// 3. If opts.RequireHuman, ensure at least one Human message exists
// 4. If opts.NoConsecutiveSame, ensure no consecutive same-type messages
// 5. Check Tool messages follow AI messages with tool calls
// 6. Return descriptive error or nil if valid
//
// Example validations:
//   - Empty conversation: error or ok?
//   - [System, Human, AI]: valid
//   - [Human, System]: invalid if RequireSystemFirst
//   - [System, AI]: invalid if RequireHuman
//   - [System, Human, Human]: invalid if NoConsecutiveSame

func ValidateConversation(messages []core.Message, opts ValidationOptions) error {
	// YOUR CODE HERE
	return nil
}

func main() {
	// Test 1: Valid conversation
	valid := []core.Message{
		core.NewSystemMessage("You are helpful", nil),
		core.NewHumanMessage("Hello", nil),
		core.NewAIMessage("Hi there!", nil),
	}
	
	opts := ValidationOptions{
		RequireSystemFirst: true,
		RequireHuman:       true,
		NoConsecutiveSame:  true,
	}
	
	if err := ValidateConversation(valid, opts); err != nil {
		fmt.Printf("Validation failed: %v\n", err)
	} else {
		fmt.Println("✓ Valid conversation")
	}
	
	// TODO: Test invalid conversations
}
