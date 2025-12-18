# Memory: Conversation History Management

**Part 2: Composition - Lesson 5**

> Make your AI remember conversations.

## Overview

All the chains you've built are **stateless** - they forget everything after each call. For conversational AI, you need **memory** to maintain context across turns.

```go
// Without memory (stateless)
chain.Invoke(ctx, "My name is Alice", nil) // "Nice to meet you"
chain.Invoke(ctx, "What's my name?", nil)  // "I don't know" ❌

// With memory (stateful)
memoryChain.Invoke(ctx, "My name is Alice", nil) // "Nice to meet you"
memoryChain.Invoke(ctx, "What's my name?", nil)  // "Your name is Alice" ✅
```

## Why This Matters

### The Problem: Forgetful AI

Without memory, every conversation starts from scratch:

```go
// User: "I like pizza"
response1 := chain.Invoke(ctx, "I like pizza", nil)
// AI: "That's great!"

// User: "What food do I like?"
response2 := chain.Invoke(ctx, "What food do I like?", nil)  
// AI: "I don't have that information" ❌
```

### The Solution: Conversation Memory

```go
memory := NewConversationBufferMemory(maxMessages)
chain := NewConversationChain(llm, memory)

// Conversation 1
chain.Invoke(ctx, "I like pizza", nil)
// AI remembers: [Human: "I like pizza", AI: "That's great!"]

// Conversation 2
chain.Invoke(ctx, "What food do I like?", nil)
// AI: "You mentioned you like pizza" ✅
```

## Core Concepts

### Memory Types

1. **Buffer Memory** - Store last N messages
2. **Summary Memory** - Summarize old messages
3. **Token Buffer** - Limit by token count
4. **Entity Memory** - Extract and track entities

### ConversationBufferMemory

Simplest memory: sliding window of messages.

```go
type ConversationBufferMemory struct {
    messages   []core.Message
    maxSize    int
    systemMsg  core.Message
}

func (m *ConversationBufferMemory) AddMessage(msg core.Message) {
    if msg.GetType() == core.MessageTypeSystem {
        m.systemMsg = msg
        return
    }
    
    m.messages = append(m.messages, msg)
    
    // Enforce max size (sliding window)
    if len(m.messages) > m.maxSize {
        m.messages = m.messages[1:]
    }
}

func (m *ConversationBufferMemory) GetMessages() []core.Message {
    result := []core.Message{}
    if m.systemMsg != nil {
        result = append(result, m.systemMsg)
    }
    result = append(result, m.messages...)
    return result
}
```

## Implementation in Go

### ConversationChain

Chain that maintains conversation history:

```go
type ConversationChain struct {
    *core.BaseRunnable
    llm    core.Runnable
    memory Memory
}

func NewConversationChain(llm core.Runnable, memory Memory) *ConversationChain {
    return &ConversationChain{
        BaseRunnable: core.NewBaseRunnable("ConversationChain"),
        llm:         llm,
        memory:      memory,
    }
}

func (c *ConversationChain) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
    // Get user input
    userMsg, ok := input.(string)
    if !ok {
        return nil, fmt.Errorf("input must be string")
    }
    
    // Add user message to memory
    c.memory.AddMessage(core.NewHumanMessage(userMsg, nil))
    
    // Get conversation history
    history := c.memory.GetMessages()
    
    // Call LLM with history
    response, err := c.llm.Invoke(ctx, history, config)
    if err != nil {
        return nil, err
    }
    
    // Add AI response to memory
    aiMsg := core.NewAIMessage(response.(string), nil)
    c.memory.AddMessage(aiMsg)
    
    return response, nil
}
```

## Practical Examples

### Example 1: Basic Conversation

```go
func basicConversation() {
    memory := NewConversationBufferMemory(10)
    llm, _ := NewLlamaCppLLM(LlamaCppConfig{
        ModelPath: "models/Qwen3-1.7B-Q8_0.gguf",
    })
    defer llm.Close()
    
    chain := NewConversationChain(llm, memory)
    
    ctx := context.Background()
    
    // Turn 1
    resp1, _ := chain.Invoke(ctx, "My favorite color is blue", nil)
    fmt.Println("AI:", resp1)
    
    // Turn 2
    resp2, _ := chain.Invoke(ctx, "What's my favorite color?", nil)
    fmt.Println("AI:", resp2) // "Blue"
}
```

### Example 2: With System Message

```go
func withSystemMessage() {
    memory := NewConversationBufferMemory(10)
    memory.AddMessage(core.NewSystemMessage(
        "You are a helpful assistant that remembers user preferences",
        nil,
    ))
    
    chain := NewConversationChain(llm, memory)
    
    // System message is included in every call
}
```

### Example 3: Persistent Memory

```go
func persistentMemory() {
    memory := NewConversationBufferMemory(10)
    
    // Save to file
    memory.Save("conversation.json")
    
    // Load later
    memory2 := NewConversationBufferMemory(10)
    memory2.Load("conversation.json")
    
    // Continue conversation
}
```

## Exercises

### Exercise 33: Summary Memory
Implement memory that summarizes old messages.

### Exercise 34: Token Buffer Memory
Limit memory by token count, not message count.

### Exercise 35: Entity Memory
Extract and track entities (names, dates, etc.).

### Exercise 36: Multi-User Memory
Manage separate memories for different users.

## Key Takeaways

1. ✅ **Memory** - Maintain conversation context
2. ✅ **Buffer Memory** - Sliding window of messages
3. ✅ **Sliding Window** - Keep recent, discard old
4. ✅ **System Message** - Always preserved
5. ✅ **Persistence** - Save/load conversations
6. ✅ **Stateful Chains** - Chains with memory

## Congratulations!

You've completed **Part 2: Composition**! 🎉

You now know:
- ✅ Prompts - Template-driven inputs
- ✅ Parsers - Structured output extraction
- ✅ LLM Chains - Prompt + LLM + Parser
- ✅ Piping - Multi-step pipelines
- ✅ Memory - Conversation history

## What's Next

**Part 3: Agency** - Build autonomous agents that use tools and make decisions

Start with [Part 3 Overview](../../03-agency/README.md) →
