# Exercise 06: Conversation Validator

Validate conversation structure and message flow.

## 🎯 Objective

Learn to validate message sequences and conversation rules.

## 📚 Concepts

- Message ordering rules
- State validation
- Business logic in Go
- Error reporting

## 📝 Task

Implement `ValidateConversation` that checks:

1. First message must be System (optional rule)
2. No consecutive messages of the same type (except AI/Human alternating)
3. Tool messages must follow AI messages with tool calls
4. At least one Human message exists

## ✅ Requirements

- [ ] Check message ordering
- [ ] Validate tool call flow
- [ ] Return descriptive errors
- [ ] Support optional strict mode
- [ ] Handle edge cases (empty, single message)

## 💡 Hints

1. Iterate through messages with index
2. Check msg.GetType() for each
3. Track previous message type
4. For tool messages, verify previous AI message had tool calls

## 🧪 Tests

```go
func TestValidator(t *testing.T) {
    // Valid conversation
    valid := []core.Message{
        core.NewSystemMessage("You are helpful", nil),
        core.NewHumanMessage("Hello", nil),
        core.NewAIMessage("Hi!", nil),
    }
    err := ValidateConversation(valid)
    // Should return nil
    
    // Invalid: no human message
    invalid := []core.Message{
        core.NewSystemMessage("You are helpful", nil),
        core.NewAIMessage("Hi!", nil),
    }
    err = ValidateConversation(invalid)
    // Should return error
}
```

## 🚀 Next Steps

After completing, move to Exercise 07: Chat History Manager
