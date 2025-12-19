# Executor: Running Complex Workflows

**Part 4: Graphs - Lesson 4**

> Execute graphs efficiently with streaming and parallelism.

## Overview

The **GraphExecutor** manages graph execution with advanced features.

```go
executor := NewGraphExecutor(graph, ExecutorConfig{
    Parallel: true,
    Stream:   true,
    MaxSteps: 100,
})

result, err := executor.Execute(ctx, state)
```

## Implementation

```go
type GraphExecutor struct {
    graph    *StateGraph
    parallel bool
    stream   bool
    maxSteps int
}

func NewGraphExecutor(graph *StateGraph, config ExecutorConfig) *GraphExecutor {
    return &GraphExecutor{
        graph:    graph,
        parallel: config.Parallel,
        stream:   config.Stream,
        maxSteps: config.MaxSteps,
    }
}

func (e *GraphExecutor) Execute(ctx context.Context, state State) (State, error) {
    if e.stream {
        return e.executeWithStreaming(ctx, state)
    }
    if e.parallel {
        return e.graph.ExecuteParallel(ctx, state)
    }
    return e.graph.Execute(ctx, state)
}

func (e *GraphExecutor) executeWithStreaming(ctx context.Context, state State) (State, error) {
    stateChan := make(chan State, 10)
    go func() {
        defer close(stateChan)
        e.graph.Execute(ctx, state)
    }()
    
    for s := range stateChan {
        state = s
    }
    return state, nil
}
```

## Exercises

- Exercise 69: Streaming Executor
- Exercise 70: Parallel Optimization
- Exercise 71: Progress Tracking
- Exercise 72: Error Recovery

**Next**: [05-checkpointing](../05-checkpointing/lesson.md)
