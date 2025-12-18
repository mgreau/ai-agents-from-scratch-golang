package agents

import (
	"context"
	"fmt"
	"strings"

	"github.com/mgreau/ai-agents-from-scratch-go/pkg/llm"
	"github.com/mgreau/ai-agents-from-scratch-go/pkg/tools"
)

// ReActAgent implements the ReAct (Reasoning + Acting) pattern
type ReActAgent struct {
	llm        *llm.LlamaCppLLM
	tools      *tools.ToolRegistry
	maxIter    int
	verbose    bool
	scratchpad []string
}

// NewReActAgent creates a new ReAct agent
func NewReActAgent(llm *llm.LlamaCppLLM, toolRegistry *tools.ToolRegistry, maxIter int, verbose bool) *ReActAgent {
	return &ReActAgent{
		llm:        llm,
		tools:      toolRegistry,
		maxIter:    maxIter,
		verbose:    verbose,
		scratchpad: []string{},
	}
}

// Run executes the ReAct loop
func (a *ReActAgent) Run(ctx context.Context, query string) (string, error) {
	if a.verbose {
		fmt.Printf("\n=== ReAct Agent Started ===\n")
		fmt.Printf("Query: %s\n\n", query)
	}

	// Build initial prompt with tools
	systemPrompt := a.buildSystemPrompt()
	prompt := fmt.Sprintf("%s\n\nQuestion: %s\n\nThought:", systemPrompt, query)

	for i := 0; i < a.maxIter; i++ {
		if a.verbose {
			fmt.Printf("--- Iteration %d ---\n", i+1)
		}

		// Get LLM response
		response, err := a.llm.Invoke(ctx, prompt, nil)
		if err != nil {
			return "", fmt.Errorf("LLM invocation failed: %w", err)
		}

		responseStr, ok := response.(string)
		if !ok {
			return "", fmt.Errorf("unexpected response type")
		}

		a.scratchpad = append(a.scratchpad, responseStr)

		if a.verbose {
			fmt.Printf("Response: %s\n", responseStr)
		}

		// Parse response for actions
		if strings.Contains(responseStr, "Action:") {
			action, actionInput := a.parseAction(responseStr)
			
			if a.verbose {
				fmt.Printf("Action: %s\n", action)
				fmt.Printf("Action Input: %s\n", actionInput)
			}

			// Execute tool
			observation, err := a.tools.ExecuteTool(ctx, action, actionInput)
			if err != nil {
				observation = fmt.Sprintf("Error: %v", err)
			}

			if a.verbose {
				fmt.Printf("Observation: %s\n\n", observation)
			}

			// Add observation to prompt for next iteration
			prompt = fmt.Sprintf("%s\nObservation: %s\n\nThought:", prompt, observation)
			a.scratchpad = append(a.scratchpad, observation)
			
		} else if strings.Contains(responseStr, "Final Answer:") {
			// Extract and return final answer
			answer := a.extractFinalAnswer(responseStr)
			if a.verbose {
				fmt.Printf("\n=== ReAct Agent Completed ===\n")
				fmt.Printf("Final Answer: %s\n", answer)
			}
			return answer, nil
		} else {
			// Continue reasoning
			prompt = prompt + responseStr + "\n\n"
		}
	}

	return "", fmt.Errorf("max iterations reached without final answer")
}

// buildSystemPrompt creates the system prompt with tool descriptions
func (a *ReActAgent) buildSystemPrompt() string {
	toolsDesc := ""
	for _, tool := range a.tools.GetAll() {
		toolsDesc += fmt.Sprintf("- %s: %s\n", tool.Name(), tool.Description())
	}

	return fmt.Sprintf(`You are a helpful assistant that can use tools to answer questions.

Available tools:
%s

Use the following format:

Question: the input question you must answer
Thought: you should always think about what to do
Action: the action to take, should be one of [%s]
Action Input: the input to the action
Observation: the result of the action
... (this Thought/Action/Action Input/Observation can repeat N times)
Thought: I now know the final answer
Final Answer: the final answer to the original input question

Begin!`, toolsDesc, strings.Join(a.getToolNames(), ", "))
}

// parseAction extracts action and action input from response
func (a *ReActAgent) parseAction(response string) (string, string) {
	action := ""
	actionInput := ""

	lines := strings.Split(response, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "Action:") {
			action = strings.TrimSpace(strings.TrimPrefix(line, "Action:"))
		} else if strings.HasPrefix(line, "Action Input:") {
			actionInput = strings.TrimSpace(strings.TrimPrefix(line, "Action Input:"))
		}
	}

	return action, actionInput
}

// extractFinalAnswer extracts the final answer from response
func (a *ReActAgent) extractFinalAnswer(response string) string {
	if idx := strings.Index(response, "Final Answer:"); idx != -1 {
		return strings.TrimSpace(response[idx+len("Final Answer:"):])
	}
	return response
}

// getToolNames returns list of tool names
func (a *ReActAgent) getToolNames() []string {
	names := []string{}
	for _, tool := range a.tools.GetAll() {
		names = append(names, tool.Name())
	}
	return names
}

// GetScratchpad returns the agent's reasoning history
func (a *ReActAgent) GetScratchpad() []string {
	return a.scratchpad
}
