package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// ChatHistoryManager manages conversation history with size limits
type ChatHistoryManager struct {
	messages   []core.Message
	maxSize    int
	systemMsg  core.Message // Preserved system message
}

// NewChatHistoryManager creates a new history manager
func NewChatHistoryManager(maxSize int) *ChatHistoryManager {
	return &ChatHistoryManager{
		messages: make([]core.Message, 0),
		maxSize:  maxSize,
	}
}

// AddMessage adds a message to history
// TODO: Implement this method
//
// Requirements:
// 1. If message is System, store it separately (systemMsg field)
// 2. Add message to messages slice
// 3. If exceeds maxSize:
//    - Remove oldest message (but never remove systemMsg)
//    - Keep most recent messages
// 4. System message doesn't count toward maxSize

func (h *ChatHistoryManager) AddMessage(msg core.Message) {
	// YOUR CODE HERE
}

// GetMessages returns all messages (including system)
// TODO: Implement this method
//
// Requirements:
// 1. If systemMsg exists, it should be first
// 2. Return all stored messages in order
// 3. Return a copy, not original slice

func (h *ChatHistoryManager) GetMessages() []core.Message {
	// YOUR CODE HERE
	return nil
}

// Save persists history to a JSON file
// TODO: Implement this method

func (h *ChatHistoryManager) Save(filename string) error {
	// YOUR CODE HERE
	return nil
}

// Load restores history from a JSON file
// TODO: Implement this method

func (h *ChatHistoryManager) Load(filename string) error {
	// YOUR CODE HERE
	return nil
}

func main() {
	history := NewChatHistoryManager(5)
	
	// Add messages
	history.AddMessage(core.NewSystemMessage("You are a helpful assistant", nil))
	history.AddMessage(core.NewHumanMessage("Hello", nil))
	history.AddMessage(core.NewAIMessage("Hi!", nil))
	
	messages := history.GetMessages()
	fmt.Printf("History has %d messages\n", len(messages))
	for _, msg := range messages {
		fmt.Printf("  - %s: %s\n", msg.GetType(), msg.GetContent())
	}
}
