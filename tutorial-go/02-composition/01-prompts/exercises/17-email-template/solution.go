package main

import (
	"context"
	"fmt"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/prompts"
)

// EmailPromptTemplate - SOLUTION
func NewEmailPromptTemplateSolution() *prompts.PromptTemplate {
	template := `To: {recipient}
From: {sender}
Subject: {subject}

Dear {recipient},

{message}

{call_to_action}

{signature}`
	
	return prompts.NewPromptTemplate(prompts.PromptTemplateConfig{
		Template: template,
		InputVariables: []string{
			"recipient", "sender", "subject", 
			"message", "call_to_action", "signature",
		},
	})
}

func runSolution() {
	template := NewEmailPromptTemplateSolution()
	ctx := context.Background()
	
	// Test 1: Formal email
	fmt.Println("=== Test 1: Formal Email ===")
	formal, _ := template.Invoke(ctx, map[string]string{
		"recipient":      "John Smith",
		"sender":         "Alice Johnson",
		"subject":        "Q4 Project Review",
		"message":        "I hope this email finds you well. I wanted to provide an update on our Q4 project status.",
		"call_to_action": "Please review the attached report at your earliest convenience.",
		"signature":      "Best regards,\nAlice Johnson\nProject Manager",
	}, nil)
	fmt.Println(formal)
	
	// Test 2: Casual email
	fmt.Println("\n=== Test 2: Casual Email ===")
	casual, _ := template.Invoke(ctx, map[string]string{
		"recipient":      "Mike",
		"sender":         "Sarah",
		"subject":        "Quick Update",
		"message":        "Hey! Just wanted to let you know the feature is ready for testing.",
		"call_to_action": "Let me know when you have time to check it out!",
		"signature":      "Thanks,\nSarah",
	}, nil)
	fmt.Println(casual)
	
	// Test 3: With partial variables
	fmt.Println("\n=== Test 3: With Default Signature ===")
	templateWithDefaults := prompts.NewPromptTemplate(prompts.PromptTemplateConfig{
		Template: template.(*prompts.PromptTemplate).Template,
		InputVariables: []string{
			"recipient", "sender", "subject", 
			"message", "call_to_action",
		},
		PartialVariables: map[string]string{
			"signature": "Best regards,\nThe Team",
		},
	})
	
	defaultSig, _ := templateWithDefaults.Invoke(ctx, map[string]string{
		"recipient":      "Team",
		"sender":         "Manager",
		"subject":        "Meeting Reminder",
		"message":        "Don't forget our meeting tomorrow at 10 AM.",
		"call_to_action": "See you there!",
	}, nil)
	fmt.Println(defaultSig)
}
