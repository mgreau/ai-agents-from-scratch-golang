package core

import (
	"context"
	"fmt"
)

// Runnable is the base interface for all composable components.
// Every Runnable must implement the Call method.
// This provides invoke, stream, batch, and pipe capabilities.
type Runnable interface {
	Invoke(ctx context.Context, input interface{}, config *Config) (interface{}, error)
	Stream(ctx context.Context, input interface{}, config *Config) (<-chan interface{}, error)
	Batch(ctx context.Context, inputs []interface{}, config *Config) ([]interface{}, error)
	Pipe(other Runnable) Runnable
	Name() string
}

// BaseRunnable provides default implementations for Runnable interface
type BaseRunnable struct {
	name string
}

// NewBaseRunnable creates a new BaseRunnable with the given name
func NewBaseRunnable(name string) *BaseRunnable {
	return &BaseRunnable{name: name}
}

// Name returns the name of this runnable
func (r *BaseRunnable) Name() string {
	return r.name
}

// Invoke processes a single input
func (r *BaseRunnable) Invoke(ctx context.Context, input interface{}, config *Config) (interface{}, error) {
	if config == nil {
		config = NewConfig()
	}

	// Create callback manager
	cm := NewCallbackManager(config.Callbacks)

	// Notify callbacks: starting
	if err := cm.HandleStart(ctx, r, input); err != nil {
		return nil, err
	}

	// Execute the runnable - this should be overridden by subclasses
	output, err := r.call(ctx, input, config)
	if err != nil {
		// Notify callbacks: error
		if cbErr := cm.HandleError(ctx, r, err); cbErr != nil {
			return nil, fmt.Errorf("callback error: %w, original error: %v", cbErr, err)
		}
		return nil, err
	}

	// Notify callbacks: success
	if err := cm.HandleEnd(ctx, r, output); err != nil {
		return nil, err
	}

	return output, nil
}

// call is the internal method that subclasses must implement
func (r *BaseRunnable) call(ctx context.Context, input interface{}, config *Config) (interface{}, error) {
	return nil, fmt.Errorf("%s must implement call() method", r.name)
}

// Stream outputs in chunks
func (r *BaseRunnable) Stream(ctx context.Context, input interface{}, config *Config) (<-chan interface{}, error) {
	out := make(chan interface{}, 1)
	go func() {
		defer close(out)
		result, err := r.Invoke(ctx, input, config)
		if err != nil {
			return
		}
		out <- result
	}()
	return out, nil
}

// Batch processes multiple inputs in parallel
func (r *BaseRunnable) Batch(ctx context.Context, inputs []interface{}, config *Config) ([]interface{}, error) {
	results := make([]interface{}, len(inputs))
	errors := make([]error, len(inputs))

	// Process all inputs concurrently
	done := make(chan bool, len(inputs))
	for i, input := range inputs {
		go func(idx int, inp interface{}) {
			results[idx], errors[idx] = r.Invoke(ctx, inp, config)
			done <- true
		}(i, input)
	}

	// Wait for all to complete
	for range inputs {
		<-done
	}

	// Check for errors
	for _, err := range errors {
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}

// Pipe composes this Runnable with another
func (r *BaseRunnable) Pipe(other Runnable) Runnable {
	return NewRunnableSequence([]Runnable{r, other})
}

// RunnableSequence chains multiple Runnables together
type RunnableSequence struct {
	*BaseRunnable
	steps []Runnable
}

// NewRunnableSequence creates a new sequence of runnables
func NewRunnableSequence(steps []Runnable) *RunnableSequence {
	return &RunnableSequence{
		BaseRunnable: NewBaseRunnable("RunnableSequence"),
		steps:        steps,
	}
}

// Invoke runs through each step sequentially
func (rs *RunnableSequence) Invoke(ctx context.Context, input interface{}, config *Config) (interface{}, error) {
	if config == nil {
		config = NewConfig()
	}

	output := input
	var err error

	// Run through each step sequentially
	for _, step := range rs.steps {
		output, err = step.Invoke(ctx, output, config)
		if err != nil {
			return nil, err
		}
	}

	return output, nil
}

// Stream streams through all steps
func (rs *RunnableSequence) Stream(ctx context.Context, input interface{}, config *Config) (<-chan interface{}, error) {
	if config == nil {
		config = NewConfig()
	}

	// Process all steps except the last one normally
	output := input
	var err error
	for i := 0; i < len(rs.steps)-1; i++ {
		output, err = rs.steps[i].Invoke(ctx, output, config)
		if err != nil {
			out := make(chan interface{})
			close(out)
			return out, err
		}
	}

	// Stream only the last step
	return rs.steps[len(rs.steps)-1].Stream(ctx, output, config)
}

// Pipe adds a step to the sequence
func (rs *RunnableSequence) Pipe(other Runnable) Runnable {
	steps := make([]Runnable, len(rs.steps)+1)
	copy(steps, rs.steps)
	steps[len(rs.steps)] = other
	return NewRunnableSequence(steps)
}

// RunnableParallel executes multiple runnables in parallel
type RunnableParallel struct {
	*BaseRunnable
	runnables map[string]Runnable
}

// NewRunnableParallel creates a new parallel runnable
func NewRunnableParallel(runnables map[string]Runnable) *RunnableParallel {
	return &RunnableParallel{
		BaseRunnable: NewBaseRunnable("RunnableParallel"),
		runnables:    runnables,
	}
}

// Invoke runs all runnables in parallel
func (rp *RunnableParallel) Invoke(ctx context.Context, input interface{}, config *Config) (interface{}, error) {
	if config == nil {
		config = NewConfig()
	}

	type result struct {
		key   string
		value interface{}
		err   error
	}

	results := make(chan result, len(rp.runnables))

	// Run all runnables in parallel
	for key, runnable := range rp.runnables {
		go func(k string, r Runnable) {
			output, err := r.Invoke(ctx, input, config)
			results <- result{key: k, value: output, err: err}
		}(key, runnable)
	}

	// Collect results
	output := make(map[string]interface{})
	for i := 0; i < len(rp.runnables); i++ {
		res := <-results
		if res.err != nil {
			return nil, res.err
		}
		output[res.key] = res.value
	}

	return output, nil
}
