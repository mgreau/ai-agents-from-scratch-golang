package main

import (
	"context"
	"fmt"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
	"github.com/mgreau/ai-agents-from-scratch-go/pkg/prompts"
)

func runSolution() {
	template := prompts.NewChatPromptTemplate([]prompts.MessageTemplate{
		{Role: core.MessageTypeSystem, Template: "You are a {role}"},
		{Role: core.MessageTypeHuman, Template: "{user_question}"},
	})
	
	ctx := context.Background()
	messages, _ := template.FormatMessages(map[string]string{
		"role": "helpful coding assistant",
		"user_question": "Explain Go interfaces",
	})
	
	for _, msg := range messages {
		fmt.Printf("%s: %s\n", msg.GetType(), msg.GetContent())
	}
}
