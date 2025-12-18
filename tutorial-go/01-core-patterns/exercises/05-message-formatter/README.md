# Exercise 05: Message Formatter

Build a beautiful console formatter for displaying messages.

## 🎯 Objective

Learn to work with Message interfaces and format output.

## 📚 Concepts

- Interface type assertions
- Message type switching
- Console color output (optional)
- String formatting

## 📝 Task

Implement `MessageFormatter` that displays messages nicely:

- System messages in gray/dim
- Human messages with "User:" prefix
- AI messages with "Assistant:" prefix
- Tool messages with "Tool:" prefix and result preview

## ✅ Requirements

- [ ] Accept []Message as input
- [ ] Format each message based on type
- [ ] Include timestamps
- [ ] Truncate long messages (optional)
- [ ] Return formatted string

## 💡 Hints

1. Use type assertion: `msg.GetType()`
2. Format timestamps with time.Format()
3. Use fmt.Sprintf() for formatting
4. Consider ANSI colors for terminal output

## 🧪 Tests

```go
func TestFormatter(t *testing.T) {
    messages := []core.Message{
        core.NewSystemMessage("You are helpful", nil),
        core.NewHumanMessage("Hello!", nil),
        core.NewAIMessage("Hi there!", nil),
    }
    
    formatted := FormatMessages(messages)
    // Should display nicely formatted conversation
}
```

## 🚀 Next Steps

After completing, move to Exercise 06: Conversation Validator
