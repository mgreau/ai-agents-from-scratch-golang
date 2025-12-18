package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// FormatMessages formats messages for console display - SOLUTION
func FormatMessagesSolution(messages []core.Message) string {
	var sb strings.Builder
	
	for _, msg := range messages {
		// Format timestamp
		t := time.UnixMilli(msg.GetTimestamp())
		timestamp := t.Format("15:04:05")
		
		// Format based on type
		var prefix string
		switch msg.GetType() {
		case core.MessageTypeSystem:
			prefix = "⚙️  SYSTEM"
		case core.MessageTypeHuman:
			prefix = "👤 USER  "
		case core.MessageTypeAI:
			prefix = "🤖 AI    "
		case core.MessageTypeTool:
			prefix = "🔧 TOOL  "
		default:
			prefix = "❓ UNKNOWN"
		}
		
		// Build formatted line
		content := msg.GetContent()
		// Truncate if too long
		if len(content) > 100 {
			content = content[:97] + "..."
		}
		
		sb.WriteString(fmt.Sprintf("[%s] %s | %s\n", timestamp, prefix, content))
	}
	
	return sb.String()
}

// FormatMessagesWithColors adds ANSI color codes - BONUS
func FormatMessagesWithColors(messages []core.Message) string {
	const (
		colorReset  = "\033[0m"
		colorGray   = "\033[90m"
		colorCyan   = "\033[36m"
		colorGreen  = "\033[32m"
		colorYellow = "\033[33m"
	)
	
	var sb strings.Builder
	
	for _, msg := range messages {
		t := time.UnixMilli(msg.GetTimestamp())
		timestamp := t.Format("15:04:05")
		
		var prefix, color string
		switch msg.GetType() {
		case core.MessageTypeSystem:
			prefix = "⚙️  SYSTEM"
			color = colorGray
		case core.MessageTypeHuman:
			prefix = "👤 USER  "
			color = colorCyan
		case core.MessageTypeAI:
			prefix = "🤖 AI    "
			color = colorGreen
		case core.MessageTypeTool:
			prefix = "🔧 TOOL  "
			color = colorYellow
		}
		
		content := msg.GetContent()
		if len(content) > 100 {
			content = content[:97] + "..."
		}
		
		sb.WriteString(fmt.Sprintf("%s[%s] %s | %s%s\n", 
			color, timestamp, prefix, content, colorReset))
	}
	
	return sb.String()
}

func runSolution() {
	// Test 1: Basic formatting
	fmt.Println("=== Test 1: Basic Formatting ===")
	messages := []core.Message{
		core.NewSystemMessage("You are a helpful coding assistant", nil),
		core.NewHumanMessage("How do I create a slice in Go?", nil),
		core.NewAIMessage("You can create a slice using: s := make([]int, 0)", nil),
	}
	
	fmt.Print(FormatMessagesSolution(messages))
	
	// Test 2: With tool message
	fmt.Println("\n=== Test 2: With Tool Call ===")
	messages2 := []core.Message{
		core.NewSystemMessage("You have access to a calculator", nil),
		core.NewHumanMessage("What's 15 * 23?", nil),
		core.NewAIMessage("Let me calculate that...", nil),
		core.NewToolMessage("345", "call_123", nil),
		core.NewAIMessage("15 * 23 = 345", nil),
	}
	
	fmt.Print(FormatMessagesSolution(messages2))
	
	// Test 3: Long message truncation
	fmt.Println("\n=== Test 3: Long Message Truncation ===")
	longMsg := core.NewAIMessage(strings.Repeat("This is a very long message. ", 20), nil)
	messages3 := []core.Message{longMsg}
	
	fmt.Print(FormatMessagesSolution(messages3))
	
	// Test 4: Colored output (if terminal supports it)
	fmt.Println("\n=== Test 4: With Colors ===")
	fmt.Print(FormatMessagesWithColors(messages))
	
	// Test 5: Empty conversation
	fmt.Println("\n=== Test 5: Empty Conversation ===")
	emptyMessages := []core.Message{}
	result := FormatMessagesSolution(emptyMessages)
	if result == "" {
		fmt.Println("(empty conversation)")
	}
}
