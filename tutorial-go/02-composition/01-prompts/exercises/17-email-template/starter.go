package main

import (
	"context"
	"fmt"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/prompts"
)

// EmailPromptTemplate generates professional emails
type EmailPromptTemplate struct {
	// TODO: Add fields
}

// NewEmailPromptTemplate creates email template
// TODO: Implement constructor
// 
// Template should include:
// - To: {recipient}
// - From: {sender}
// - Subject: {subject}
// - Tone: {tone} (formal/casual)
// - Body: {message}
// - CTA: {call_to_action}
// - Signature: {signature}

func NewEmailPromptTemplate() *EmailPromptTemplate {
	// YOUR CODE HERE
	return nil
}

func main() {
	template := NewEmailPromptTemplate()
	
	ctx := context.Background()
	
	// Test formal email
	formal, _ := template.Invoke(ctx, map[string]string{
		"recipient":      "John Smith",
		"sender":         "Alice Johnson",
		"subject":        "Project Update",
		"tone":           "formal",
		"message":        "I wanted to update you on our progress.",
		"call_to_action": "Please review the attached document.",
		"signature":      "Best regards,\nAlice Johnson\nProject Manager",
	}, nil)
	
	fmt.Println("=== Formal Email ===")
	fmt.Println(formal)
}
