# Exercise 07: Chat History Manager

Build a conversation history manager with size limits and persistence.

## 🎯 Objective

Learn to manage conversation state and implement sliding windows.

## 📚 Concepts

- State management
- Sliding window buffer
- Message persistence
- Token/message counting

## 📝 Task

Implement `ChatHistoryManager` that:

1. Stores conversation messages
2. Maintains a maximum size (by message count or tokens)
3. Uses sliding window (removes oldest when full)
4. Always keeps System message
5. Can save/load from JSON

## ✅ Requirements

- [ ] Add messages with `AddMessage()`
- [ ] Enforce max size (keep most recent)
- [ ] Always preserve System message
- [ ] Get messages with `GetMessages()`
- [ ] Implement `Save()` and `Load()`
- [ ] Support token-based limits (bonus)

## 💡 Hints

1. Use a slice to store messages
2. When adding, check if at max capacity
3. If full, remove oldest (but keep System)
4. Use json.Marshal for persistence
5. Consider using a circular buffer for efficiency

## 🧪 Tests

```go
func TestChatHistory(t *testing.T) {
    history := NewChatHistoryManager(5) // max 5 messages
    
    history.AddMessage(NewSystemMessage("You are helpful", nil))
    history.AddMessage(NewHumanMessage("Hello", nil))
    // ... add more messages
    
    messages := history.GetMessages()
    // Should have at most 5 messages, with System always present
}
```

## 🚀 Next Steps

After completing, move to Exercise 08: Tool Flow Handler
