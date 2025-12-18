package llm

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	llama "github.com/go-skynet/go-llama.cpp"
	"github.com/mgreau/ai-agents-from-scratch-go/pkg/core"
)

// LlamaCppLLM wraps go-llama.cpp for local inference
type LlamaCppLLM struct {
	*core.BaseRunnable
	model          *llama.LLama
	modelPath      string
	contextSize    int
	temperature    float32
	topP           float32
	topK           int
	threads        int
	systemPrompt   string
}

// LlamaCppConfig holds configuration for LlamaCpp LLM
type LlamaCppConfig struct {
	ModelPath    string
	ContextSize  int
	Temperature  float32
	TopP         float32
	TopK         int
	Threads      int
	SystemPrompt string
}

// NewLlamaCppLLM creates a new LlamaCpp LLM instance
func NewLlamaCppLLM(config LlamaCppConfig) (*LlamaCppLLM, error) {
	// Set defaults
	if config.ContextSize == 0 {
		config.ContextSize = 2048
	}
	if config.Temperature == 0 {
		config.Temperature = 0.7
	}
	if config.TopP == 0 {
		config.TopP = 0.9
	}
	if config.TopK == 0 {
		config.TopK = 40
	}
	if config.Threads == 0 {
		config.Threads = 4
	}

	// Check if model file exists
	if _, err := os.Stat(config.ModelPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("model file not found: %s", config.ModelPath)
	}

	l := &LlamaCppLLM{
		BaseRunnable: core.NewBaseRunnable("LlamaCppLLM"),
		modelPath:    config.ModelPath,
		contextSize:  config.ContextSize,
		temperature:  config.Temperature,
		topP:         config.TopP,
		topK:         config.TopK,
		threads:      config.Threads,
		systemPrompt: config.SystemPrompt,
	}

	// Load the model with go-llama.cpp
	fmt.Printf("Loading model from: %s\n", config.ModelPath)
	model, err := llama.New(
		config.ModelPath,
		llama.SetContext(config.ContextSize),
		llama.SetThreads(config.Threads),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load model: %w", err)
	}
	l.model = model

	fmt.Println("LlamaCppLLM initialized successfully!")
	return l, nil
}

// Invoke generates a response for the given prompt
func (l *LlamaCppLLM) Invoke(ctx context.Context, input interface{}, config *core.Config) (interface{}, error) {
	prompt, ok := input.(string)
	if !ok {
		// Try to convert from messages
		if messages, ok := input.([]core.Message); ok {
			prompt = l.messagesToPrompt(messages)
		} else {
			return nil, fmt.Errorf("input must be a string or []core.Message")
		}
	}

	// Add system prompt if set
	if l.systemPrompt != "" {
		prompt = fmt.Sprintf("System: %s\n\nUser: %s\n\nAssistant:", l.systemPrompt, prompt)
	}

	// Generate response using go-llama.cpp
	result, err := l.model.Predict(
		prompt,
		llama.SetTemperature(float64(l.temperature)),
		llama.SetTopP(float64(l.topP)),
		llama.SetTopK(l.topK),
		llama.SetThreads(l.threads),
		llama.SetTokens(l.contextSize),
	)
	if err != nil {
		return nil, fmt.Errorf("prediction failed: %w", err)
	}

	return result, nil
}

// Stream generates a response and streams tokens
func (l *LlamaCppLLM) Stream(ctx context.Context, input interface{}, config *core.Config) (<-chan interface{}, error) {
	prompt, ok := input.(string)
	if !ok {
		if messages, ok := input.([]core.Message); ok {
			prompt = l.messagesToPrompt(messages)
		} else {
			return nil, fmt.Errorf("input must be a string or []core.Message")
		}
	}

	// Add system prompt if set
	if l.systemPrompt != "" {
		prompt = fmt.Sprintf("System: %s\n\nUser: %s\n\nAssistant:", l.systemPrompt, prompt)
	}

	out := make(chan interface{}, 10)

	go func() {
		defer close(out)

		// Stream response using go-llama.cpp
		err := l.model.Predict(
			prompt,
			func(token string) bool {
				select {
				case <-ctx.Done():
					return false
				case out <- token:
					return true
				}
			},
			llama.SetTemperature(float64(l.temperature)),
			llama.SetTopP(float64(l.topP)),
			llama.SetTopK(l.topK),
			llama.SetThreads(l.threads),
			llama.SetTokens(l.contextSize),
		)
		if err != nil {
			out <- fmt.Errorf("streaming failed: %w", err)
		}
	}()

	return out, nil
}

// messagesToPrompt converts messages to a prompt string
func (l *LlamaCppLLM) messagesToPrompt(messages []core.Message) string {
	prompt := ""
	for _, msg := range messages {
		switch msg.GetType() {
		case core.MessageTypeSystem:
			prompt += fmt.Sprintf("System: %s\n\n", msg.GetContent())
		case core.MessageTypeHuman:
			prompt += fmt.Sprintf("User: %s\n\n", msg.GetContent())
		case core.MessageTypeAI:
			prompt += fmt.Sprintf("Assistant: %s\n\n", msg.GetContent())
		case core.MessageTypeTool:
			prompt += fmt.Sprintf("Tool: %s\n\n", msg.GetContent())
		}
	}
	return prompt + "Assistant:"
}

// Close releases model resources
func (l *LlamaCppLLM) Close() {
	if l.model != nil {
		l.model.Free()
		fmt.Println("LlamaCppLLM closed and resources freed")
	}
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// GetModelPath returns the absolute path to the model
func GetModelPath(relativePath string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cwd, relativePath), nil
}
