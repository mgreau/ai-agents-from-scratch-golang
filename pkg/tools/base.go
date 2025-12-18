package tools

import (
	"context"
	"encoding/json"
	"fmt"
)

// Tool represents a function that can be called by an agent
type Tool interface {
	Name() string
	Description() string
	ArgsSchema() map[string]interface{}
	Execute(ctx context.Context, args map[string]interface{}) (string, error)
}

// BaseTool provides common functionality for tools
type BaseTool struct {
	name        string
	description string
	argsSchema  map[string]interface{}
}

// NewBaseTool creates a new base tool
func NewBaseTool(name, description string, argsSchema map[string]interface{}) *BaseTool {
	return &BaseTool{
		name:        name,
		description: description,
		argsSchema:  argsSchema,
	}
}

// Name returns the tool name
func (t *BaseTool) Name() string {
	return t.name
}

// Description returns the tool description
func (t *BaseTool) Description() string {
	return t.description
}

// ArgsSchema returns the JSON schema for arguments
func (t *BaseTool) ArgsSchema() map[string]interface{} {
	return t.argsSchema
}

// Execute must be implemented by concrete tools
func (t *BaseTool) Execute(ctx context.Context, args map[string]interface{}) (string, error) {
	return "", fmt.Errorf("execute not implemented for %s", t.name)
}

// ToFunctionDefinition converts tool to OpenAI-style function definition
func ToFunctionDefinition(tool Tool) map[string]interface{} {
	return map[string]interface{}{
		"type": "function",
		"function": map[string]interface{}{
			"name":        tool.Name(),
			"description": tool.Description(),
			"parameters":  tool.ArgsSchema(),
		},
	}
}

// GetCurrentTimeTool returns the current time
type GetCurrentTimeTool struct {
	*BaseTool
}

// NewGetCurrentTimeTool creates a new GetCurrentTimeTool
func NewGetCurrentTimeTool() *GetCurrentTimeTool {
	return &GetCurrentTimeTool{
		BaseTool: NewBaseTool(
			"getCurrentTime",
			"Get the current time",
			map[string]interface{}{
				"type":       "object",
				"properties": map[string]interface{}{},
			},
		),
	}
}

// Execute returns the current time
func (t *GetCurrentTimeTool) Execute(ctx context.Context, args map[string]interface{}) (string, error) {
	return fmt.Sprintf("%s", "14:30"), nil // Simplified for demo
}

// CalculatorTool performs basic calculations
type CalculatorTool struct {
	*BaseTool
}

// NewCalculatorTool creates a new calculator tool
func NewCalculatorTool() *CalculatorTool {
	return &CalculatorTool{
		BaseTool: NewBaseTool(
			"calculator",
			"Perform basic arithmetic calculations. Supports +, -, *, / operations.",
			map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"expression": map[string]interface{}{
						"type":        "string",
						"description": "The mathematical expression to evaluate (e.g., '2 + 2', '10 * 5')",
					},
				},
				"required": []string{"expression"},
			},
		),
	}
}

// Execute evaluates the mathematical expression
func (t *CalculatorTool) Execute(ctx context.Context, args map[string]interface{}) (string, error) {
	expression, ok := args["expression"].(string)
	if !ok {
		return "", fmt.Errorf("expression must be a string")
	}

	// Simplified evaluation - in production use a proper math parser
	// For demo purposes, we'll just return a mock result
	result := fmt.Sprintf("Result of '%s' is 42", expression)
	return result, nil
}

// ToolRegistry manages available tools
type ToolRegistry struct {
	tools map[string]Tool
}

// NewToolRegistry creates a new tool registry
func NewToolRegistry() *ToolRegistry {
	return &ToolRegistry{
		tools: make(map[string]Tool),
	}
}

// Register adds a tool to the registry
func (r *ToolRegistry) Register(tool Tool) {
	r.tools[tool.Name()] = tool
}

// Get retrieves a tool by name
func (r *ToolRegistry) Get(name string) (Tool, bool) {
	tool, ok := r.tools[name]
	return tool, ok
}

// GetAll returns all registered tools
func (r *ToolRegistry) GetAll() []Tool {
	tools := make([]Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		tools = append(tools, tool)
	}
	return tools
}

// GetFunctionDefinitions returns all tools as function definitions
func (r *ToolRegistry) GetFunctionDefinitions() []map[string]interface{} {
	defs := make([]map[string]interface{}, 0, len(r.tools))
	for _, tool := range r.tools {
		defs = append(defs, ToFunctionDefinition(tool))
	}
	return defs
}

// ExecuteTool executes a tool by name with given arguments
func (r *ToolRegistry) ExecuteTool(ctx context.Context, name string, argsJSON string) (string, error) {
	tool, ok := r.Get(name)
	if !ok {
		return "", fmt.Errorf("tool not found: %s", name)
	}

	// Parse arguments
	var args map[string]interface{}
	if argsJSON != "" {
		if err := json.Unmarshal([]byte(argsJSON), &args); err != nil {
			return "", fmt.Errorf("failed to parse arguments: %w", err)
		}
	} else {
		args = make(map[string]interface{})
	}

	// Execute tool
	return tool.Execute(ctx, args)
}
