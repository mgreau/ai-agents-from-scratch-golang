package core

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// MessageType represents the type of message
type MessageType string

const (
	MessageTypeSystem MessageType = "system"
	MessageTypeHuman  MessageType = "human"
	MessageTypeAI     MessageType = "ai"
	MessageTypeTool   MessageType = "tool"
)

// Message is the base interface for all message types
type Message interface {
	GetContent() string
	GetType() MessageType
	GetID() string
	GetTimestamp() int64
	ToJSON() ([]byte, error)
	ToPromptFormat() map[string]interface{}
	String() string
}

// BaseMessage contains common functionality for all message types
type BaseMessage struct {
	ID                string                 `json:"id"`
	Content           string                 `json:"content"`
	Timestamp         int64                  `json:"timestamp"`
	AdditionalKwargs  map[string]interface{} `json:"additional_kwargs,omitempty"`
}

// NewBaseMessage creates a new base message
func NewBaseMessage(content string, kwargs map[string]interface{}) *BaseMessage {
	if kwargs == nil {
		kwargs = make(map[string]interface{})
	}
	return &BaseMessage{
		ID:               generateMessageID(),
		Content:          content,
		Timestamp:        time.Now().UnixMilli(),
		AdditionalKwargs: kwargs,
	}
}

// GetContent returns the message content
func (m *BaseMessage) GetContent() string {
	return m.Content
}

// GetID returns the message ID
func (m *BaseMessage) GetID() string {
	return m.ID
}

// GetTimestamp returns the message timestamp
func (m *BaseMessage) GetTimestamp() int64 {
	return m.Timestamp
}

// generateMessageID generates a unique message ID
func generateMessageID() string {
	timestamp := time.Now().UnixMilli()
	random := rand.Intn(1000000)
	return fmt.Sprintf("msg_%d_%d", timestamp, random)
}

// SystemMessage represents instructions for the AI
type SystemMessage struct {
	*BaseMessage
}

// NewSystemMessage creates a new system message
func NewSystemMessage(content string, kwargs map[string]interface{}) *SystemMessage {
	return &SystemMessage{
		BaseMessage: NewBaseMessage(content, kwargs),
	}
}

// GetType returns the message type
func (m *SystemMessage) GetType() MessageType {
	return MessageTypeSystem
}

// ToPromptFormat converts to prompt format
func (m *SystemMessage) ToPromptFormat() map[string]interface{} {
	return map[string]interface{}{
		"role":    "system",
		"content": m.Content,
	}
}

// ToJSON converts to JSON
func (m *SystemMessage) ToJSON() ([]byte, error) {
	data := map[string]interface{}{
		"id":        m.ID,
		"type":      m.GetType(),
		"content":   m.Content,
		"timestamp": m.Timestamp,
	}
	for k, v := range m.AdditionalKwargs {
		data[k] = v
	}
	return json.Marshal(data)
}

// String returns a string representation
func (m *SystemMessage) String() string {
	t := time.UnixMilli(m.Timestamp)
	return fmt.Sprintf("[%s] system: %s", t.Format("15:04:05"), m.Content)
}

// HumanMessage represents user input
type HumanMessage struct {
	*BaseMessage
}

// NewHumanMessage creates a new human message
func NewHumanMessage(content string, kwargs map[string]interface{}) *HumanMessage {
	return &HumanMessage{
		BaseMessage: NewBaseMessage(content, kwargs),
	}
}

// GetType returns the message type
func (m *HumanMessage) GetType() MessageType {
	return MessageTypeHuman
}

// ToPromptFormat converts to prompt format
func (m *HumanMessage) ToPromptFormat() map[string]interface{} {
	return map[string]interface{}{
		"role":    "user",
		"content": m.Content,
	}
}

// ToJSON converts to JSON
func (m *HumanMessage) ToJSON() ([]byte, error) {
	data := map[string]interface{}{
		"id":        m.ID,
		"type":      m.GetType(),
		"content":   m.Content,
		"timestamp": m.Timestamp,
	}
	for k, v := range m.AdditionalKwargs {
		data[k] = v
	}
	return json.Marshal(data)
}

// String returns a string representation
func (m *HumanMessage) String() string {
	t := time.UnixMilli(m.Timestamp)
	return fmt.Sprintf("[%s] human: %s", t.Format("15:04:05"), m.Content)
}

// AIMessage represents assistant responses
type AIMessage struct {
	*BaseMessage
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
}

// ToolCall represents a request to execute a function
type ToolCall struct {
	ID       string                 `json:"id"`
	Type     string                 `json:"type"`
	Function ToolCallFunction       `json:"function"`
	Args     map[string]interface{} `json:"args,omitempty"`
}

// ToolCallFunction represents the function to be called
type ToolCallFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

// NewAIMessage creates a new AI message
func NewAIMessage(content string, kwargs map[string]interface{}) *AIMessage {
	msg := &AIMessage{
		BaseMessage: NewBaseMessage(content, kwargs),
		ToolCalls:   []ToolCall{},
	}
	
	// Extract tool calls from kwargs if present
	if toolCalls, ok := kwargs["tool_calls"]; ok {
		if tc, ok := toolCalls.([]ToolCall); ok {
			msg.ToolCalls = tc
		}
	}
	
	return msg
}

// GetType returns the message type
func (m *AIMessage) GetType() MessageType {
	return MessageTypeAI
}

// HasToolCalls checks if this message requests tool execution
func (m *AIMessage) HasToolCalls() bool {
	return len(m.ToolCalls) > 0
}

// GetToolCall gets a specific tool call by index
func (m *AIMessage) GetToolCall(index int) *ToolCall {
	if index < 0 || index >= len(m.ToolCalls) {
		return nil
	}
	return &m.ToolCalls[index]
}

// ToPromptFormat converts to prompt format
func (m *AIMessage) ToPromptFormat() map[string]interface{} {
	result := map[string]interface{}{
		"role":    "assistant",
		"content": m.Content,
	}
	
	if m.HasToolCalls() {
		result["tool_calls"] = m.ToolCalls
	}
	
	return result
}

// ToJSON converts to JSON
func (m *AIMessage) ToJSON() ([]byte, error) {
	data := map[string]interface{}{
		"id":        m.ID,
		"type":      m.GetType(),
		"content":   m.Content,
		"timestamp": m.Timestamp,
	}
	if len(m.ToolCalls) > 0 {
		data["tool_calls"] = m.ToolCalls
	}
	for k, v := range m.AdditionalKwargs {
		data[k] = v
	}
	return json.Marshal(data)
}

// String returns a string representation
func (m *AIMessage) String() string {
	t := time.UnixMilli(m.Timestamp)
	return fmt.Sprintf("[%s] ai: %s", t.Format("15:04:05"), m.Content)
}

// ToolMessage represents tool execution results
type ToolMessage struct {
	*BaseMessage
	ToolCallID string `json:"tool_call_id"`
}

// NewToolMessage creates a new tool message
func NewToolMessage(content string, toolCallID string, kwargs map[string]interface{}) *ToolMessage {
	return &ToolMessage{
		BaseMessage: NewBaseMessage(content, kwargs),
		ToolCallID:  toolCallID,
	}
}

// GetType returns the message type
func (m *ToolMessage) GetType() MessageType {
	return MessageTypeTool
}

// ToPromptFormat converts to prompt format
func (m *ToolMessage) ToPromptFormat() map[string]interface{} {
	return map[string]interface{}{
		"role":         "tool",
		"content":      m.Content,
		"tool_call_id": m.ToolCallID,
	}
}

// ToJSON converts to JSON
func (m *ToolMessage) ToJSON() ([]byte, error) {
	data := map[string]interface{}{
		"id":           m.ID,
		"type":         m.GetType(),
		"content":      m.Content,
		"timestamp":    m.Timestamp,
		"tool_call_id": m.ToolCallID,
	}
	for k, v := range m.AdditionalKwargs {
		data[k] = v
	}
	return json.Marshal(data)
}

// String returns a string representation
func (m *ToolMessage) String() string {
	t := time.UnixMilli(m.Timestamp)
	return fmt.Sprintf("[%s] tool: %s", t.Format("15:04:05"), m.Content)
}

// MessagesToPromptFormat converts messages to prompt format
func MessagesToPromptFormat(messages []Message) []map[string]interface{} {
	result := make([]map[string]interface{}, len(messages))
	for i, msg := range messages {
		result[i] = msg.ToPromptFormat()
	}
	return result
}

// FilterMessagesByType filters messages by type
func FilterMessagesByType(messages []Message, msgType MessageType) []Message {
	var filtered []Message
	for _, msg := range messages {
		if msg.GetType() == msgType {
			filtered = append(filtered, msg)
		}
	}
	return filtered
}

// GetLastMessages returns the last N messages
func GetLastMessages(messages []Message, n int) []Message {
	if n >= len(messages) {
		return messages
	}
	return messages[len(messages)-n:]
}
