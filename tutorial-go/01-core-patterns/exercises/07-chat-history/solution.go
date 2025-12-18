package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// ChatHistoryManager manages conversation history - SOLUTION
type ChatHistoryManagerSolution struct {
	messages  []core.Message
	maxSize   int
	systemMsg core.Message
}

// NewChatHistoryManagerSolution creates a new history manager
func NewChatHistoryManagerSolution(maxSize int) *ChatHistoryManagerSolution {
	return &ChatHistoryManagerSolution{
		messages: make([]core.Message, 0, maxSize),
		maxSize:  maxSize,
	}
}

// AddMessage adds a message to history
func (h *ChatHistoryManagerSolution) AddMessage(msg core.Message) {
	// Handle system message separately
	if msg.GetType() == core.MessageTypeSystem {
		h.systemMsg = msg
		return
	}
	
	// Add message
	h.messages = append(h.messages, msg)
	
	// Enforce max size (sliding window)
	if len(h.messages) > h.maxSize {
		// Remove oldest message
		h.messages = h.messages[1:]
	}
}

// GetMessages returns all messages including system
func (h *ChatHistoryManagerSolution) GetMessages() []core.Message {
	result := make([]core.Message, 0, len(h.messages)+1)
	
	// Add system message first if exists
	if h.systemMsg != nil {
		result = append(result, h.systemMsg)
	}
	
	// Add all other messages
	result = append(result, h.messages...)
	
	return result
}

// Clear removes all messages
func (h *ChatHistoryManagerSolution) Clear() {
	h.messages = make([]core.Message, 0, h.maxSize)
	h.systemMsg = nil
}

// Count returns number of messages (excluding system)
func (h *ChatHistoryManagerSolution) Count() int {
	return len(h.messages)
}

// Save persists history to JSON file
func (h *ChatHistoryManagerSolution) Save(filename string) error {
	messages := h.GetMessages()
	
	// Convert to JSON
	data := make([]json.RawMessage, len(messages))
	for i, msg := range messages {
		msgJSON, err := msg.ToJSON()
		if err != nil {
			return fmt.Errorf("failed to marshal message %d: %w", i, err)
		}
		data[i] = msgJSON
	}
	
	// Write to file
	output, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal messages: %w", err)
	}
	
	return os.WriteFile(filename, output, 0644)
}

// Load restores history from JSON file
func (h *ChatHistoryManagerSolution) Load(filename string) error {
	// Read file
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	
	// Parse JSON
	var rawMessages []map[string]interface{}
	if err := json.Unmarshal(data, &rawMessages); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	
	// Clear current history
	h.Clear()
	
	// Convert to Message types
	for _, raw := range rawMessages {
		msgType := core.MessageType(raw["type"].(string))
		content := raw["content"].(string)
		
		var msg core.Message
		switch msgType {
		case core.MessageTypeSystem:
			msg = core.NewSystemMessage(content, nil)
		case core.MessageTypeHuman:
			msg = core.NewHumanMessage(content, nil)
		case core.MessageTypeAI:
			msg = core.NewAIMessage(content, nil)
		case core.MessageTypeTool:
			toolCallID := raw["tool_call_id"].(string)
			msg = core.NewToolMessage(content, toolCallID, nil)
		default:
			return fmt.Errorf("unknown message type: %s", msgType)
		}
		
		h.AddMessage(msg)
	}
	
	return nil
}

func runSolution() {
	// Test 1: Basic usage
	fmt.Println("=== Test 1: Basic Usage ===")
	history := NewChatHistoryManagerSolution(5)
	
	history.AddMessage(core.NewSystemMessage("You are helpful", nil))
	history.AddMessage(core.NewHumanMessage("Hello", nil))
	history.AddMessage(core.NewAIMessage("Hi!", nil))
	
	messages := history.GetMessages()
	fmt.Printf("Count: %d messages\n", len(messages))
	for _, msg := range messages {
		fmt.Printf("  %s: %s\n", msg.GetType(), msg.GetContent())
	}
	
	// Test 2: Sliding window (max size)
	fmt.Println("\n=== Test 2: Sliding Window ===")
	history2 := NewChatHistoryManagerSolution(3)
	
	history2.AddMessage(core.NewSystemMessage("System msg (always kept)", nil))
	history2.AddMessage(core.NewHumanMessage("Message 1", nil))
	history2.AddMessage(core.NewAIMessage("Message 2", nil))
	history2.AddMessage(core.NewHumanMessage("Message 3", nil))
	history2.AddMessage(core.NewAIMessage("Message 4", nil))
	history2.AddMessage(core.NewHumanMessage("Message 5 (oldest should be removed)", nil))
	
	messages2 := history2.GetMessages()
	fmt.Printf("Count: %d messages (max 3 + system)\n", len(messages2))
	for _, msg := range messages2 {
		fmt.Printf("  %s: %s\n", msg.GetType(), msg.GetContent())
	}
	
	// Test 3: Save and load
	fmt.Println("\n=== Test 3: Save and Load ===")
	history3 := NewChatHistoryManagerSolution(10)
	history3.AddMessage(core.NewSystemMessage("Saved system", nil))
	history3.AddMessage(core.NewHumanMessage("Saved message", nil))
	history3.AddMessage(core.NewAIMessage("Saved response", nil))
	
	filename := "chat_history_test.json"
	if err := history3.Save(filename); err != nil {
		fmt.Printf("Save error: %v\n", err)
		return
	}
	fmt.Println("✓ Saved to", filename)
	
	// Load into new manager
	history4 := NewChatHistoryManagerSolution(10)
	if err := history4.Load(filename); err != nil {
		fmt.Printf("Load error: %v\n", err)
		return
	}
	fmt.Println("✓ Loaded from", filename)
	
	messages4 := history4.GetMessages()
	fmt.Printf("Loaded %d messages:\n", len(messages4))
	for _, msg := range messages4 {
		fmt.Printf("  %s: %s\n", msg.GetType(), msg.GetContent())
	}
	
	// Cleanup
	os.Remove(filename)
	
	// Test 4: Clear
	fmt.Println("\n=== Test 4: Clear ===")
	history.Clear()
	fmt.Printf("After clear: %d messages\n", history.Count())
}
