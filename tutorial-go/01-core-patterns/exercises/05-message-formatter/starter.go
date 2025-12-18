package main

import (
	"fmt"
	"time"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// FormatMessages formats a slice of messages for console display
// TODO: Implement this function
//
// Requirements:
// 1. Iterate through messages
// 2. Format each based on type:
//    - System: "⚙️  SYSTEM | content"
//    - Human:  "👤 USER   | content"
//    - AI:     "🤖 AI     | content"
//    - Tool:   "🔧 TOOL   | content"
// 3. Include timestamp (HH:MM:SS format)
// 4. Return formatted string with newlines
//
// Example output:
//   [14:32:15] ⚙️  SYSTEM | You are helpful
//   [14:32:16] 👤 USER   | Hello!
//   [14:32:17] 🤖 AI     | Hi there!

func FormatMessages(messages []core.Message) string {
	// YOUR CODE HERE
	return ""
}

func main() {
	// Create sample conversation
	messages := []core.Message{
		core.NewSystemMessage("You are a helpful assistant", nil),
		core.NewHumanMessage("What is Go?", nil),
		core.NewAIMessage("Go is a programming language created by Google.", nil),
		core.NewHumanMessage("Thanks!", nil),
		core.NewAIMessage("You're welcome!", nil),
	}
	
	formatted := FormatMessages(messages)
	fmt.Println(formatted)
}
