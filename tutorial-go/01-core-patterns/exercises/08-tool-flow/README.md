# Exercise 08: Tool Flow Handler

Build a handler for AI tool call → tool execution → AI response flow.

## 🎯 Objective

Learn to manage tool calls and build agent loops.

## 📚 Concepts

- Tool call detection
- Message flow control
- Tool execution simulation
- Multi-step agent loops

## 📝 Task

Implement `ToolFlowHandler` that:

1. Detects when AI message has tool calls
2. "Executes" the tools (simulated)
3. Adds ToolMessage with results
4. Continues conversation
5. Handles multiple tool calls

## ✅ Requirements

- [ ] Parse tool calls from AIMessage
- [ ] Execute tools (use mock for now)
- [ ] Create ToolMessage with results
- [ ] Link tool call ID to result
- [ ] Support multiple tools in one turn
- [ ] Return updated conversation

## 💡 Hints

1. Check `aiMsg.HasToolCalls()`
2. Iterate through `aiMsg.ToolCalls`
3. Execute each tool (mock with simple logic)
4. Create ToolMessage with same ID
5. Append to conversation

## 🧪 Tests

```go
func TestToolFlow(t *testing.T) {
    handler := NewToolFlowHandler()
    
    aiMsg := core.NewAIMessage("Let me calculate", map[string]interface{}{
        "tool_calls": []core.ToolCall{{
            ID: "call_123",
            Function: core.ToolCallFunction{
                Name: "calculator",
                Arguments: `{"a": 5, "b": 3}`,
            },
        }},
    })
    
    messages := []core.Message{aiMsg}
    updated := handler.HandleToolCalls(messages)
    
    // Should add ToolMessage with result
}
```

## 🚀 Next Steps

After completing, you've finished all Messages exercises!
Move to Lesson 03: LLM Wrapper
