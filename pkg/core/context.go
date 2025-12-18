package core

import "context"

// Config holds configuration for Runnable execution
type Config struct {
	Callbacks []Callback
	Tags      []string
	Metadata  map[string]interface{}
	MaxRetries int
	Timeout   int
}

// NewConfig creates a new Config with default values
func NewConfig() *Config {
	return &Config{
		Callbacks:  []Callback{},
		Tags:       []string{},
		Metadata:   make(map[string]interface{}),
		MaxRetries: 3,
		Timeout:    0,
	}
}

// WithCallbacks sets callbacks
func (c *Config) WithCallbacks(callbacks []Callback) *Config {
	c.Callbacks = callbacks
	return c
}

// WithTags sets tags
func (c *Config) WithTags(tags []string) *Config {
	c.Tags = tags
	return c
}

// WithMetadata sets metadata
func (c *Config) WithMetadata(metadata map[string]interface{}) *Config {
	c.Metadata = metadata
	return c
}

// WithMaxRetries sets max retries
func (c *Config) WithMaxRetries(maxRetries int) *Config {
	c.MaxRetries = maxRetries
	return c
}

// WithTimeout sets timeout in seconds
func (c *Config) WithTimeout(timeout int) *Config {
	c.Timeout = timeout
	return c
}

// Callback interface for observability
type Callback interface {
	OnStart(ctx context.Context, runnable Runnable, input interface{}) error
	OnEnd(ctx context.Context, runnable Runnable, output interface{}) error
	OnError(ctx context.Context, runnable Runnable, err error) error
}

// CallbackManager manages multiple callbacks
type CallbackManager struct {
	callbacks []Callback
}

// NewCallbackManager creates a new callback manager
func NewCallbackManager(callbacks []Callback) *CallbackManager {
	if callbacks == nil {
		callbacks = []Callback{}
	}
	return &CallbackManager{
		callbacks: callbacks,
	}
}

// HandleStart notifies all callbacks of start
func (cm *CallbackManager) HandleStart(ctx context.Context, runnable Runnable, input interface{}) error {
	for _, cb := range cm.callbacks {
		if err := cb.OnStart(ctx, runnable, input); err != nil {
			return err
		}
	}
	return nil
}

// HandleEnd notifies all callbacks of end
func (cm *CallbackManager) HandleEnd(ctx context.Context, runnable Runnable, output interface{}) error {
	for _, cb := range cm.callbacks {
		if err := cb.OnEnd(ctx, runnable, output); err != nil {
			return err
		}
	}
	return nil
}

// HandleError notifies all callbacks of error
func (cm *CallbackManager) HandleError(ctx context.Context, runnable Runnable, err error) error {
	for _, cb := range cm.callbacks {
		if cbErr := cb.OnError(ctx, runnable, err); cbErr != nil {
			return cbErr
		}
	}
	return nil
}

// LoggingCallback is a simple callback that logs events
type LoggingCallback struct {
	Verbose bool
}

// OnStart logs the start event
func (lc *LoggingCallback) OnStart(ctx context.Context, runnable Runnable, input interface{}) error {
	if lc.Verbose {
		println("[START]", runnable.Name(), "Input:", input)
	}
	return nil
}

// OnEnd logs the end event
func (lc *LoggingCallback) OnEnd(ctx context.Context, runnable Runnable, output interface{}) error {
	if lc.Verbose {
		println("[END]", runnable.Name(), "Output:", output)
	}
	return nil
}

// OnError logs the error event
func (lc *LoggingCallback) OnError(ctx context.Context, runnable Runnable, err error) error {
	if lc.Verbose {
		println("[ERROR]", runnable.Name(), "Error:", err.Error())
	}
	return nil
}
