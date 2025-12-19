# Checkpointing: Persistence and Recovery

**Part 4: Graphs - Lesson 5**

> Save and resume long-running workflows.

## Overview

**Checkpointing** saves graph state for recovery.

```go
executor := NewCheckpointingExecutor(graph, storage)
result, err := executor.Execute(ctx, state)

// If interrupted, resume:
result, err = executor.Resume(ctx, checkpointID)
```

## Implementation

```go
type CheckpointStorage interface {
    Save(id string, state State) error
    Load(id string) (State, error)
}

type CheckpointingExecutor struct {
    graph   *StateGraph
    storage CheckpointStorage
}

func (e *CheckpointingExecutor) Execute(ctx context.Context, state State) (State, error) {
    checkpointID := generateID()
    
    for currentNode := e.graph.startNode; currentNode != ""; {
        // Save checkpoint
        e.storage.Save(checkpointID, state)
        
        // Execute node
        nodeFn := e.graph.nodes[currentNode]
        newState, err := nodeFn(state)
        if err != nil {
            return state, err
        }
        state = newState
        
        // Get next node
        nextNodes := e.graph.edges[currentNode]
        if len(nextNodes) == 0 {
            break
        }
        currentNode = nextNodes[0]
    }
    
    return state, nil
}

func (e *CheckpointingExecutor) Resume(ctx context.Context, checkpointID string) (State, error) {
    state, err := e.storage.Load(checkpointID)
    if err != nil {
        return nil, err
    }
    return e.Execute(ctx, state)
}
```

## Exercises

- Exercise 73: File-Based Checkpoints
- Exercise 74: Resume Logic
- Exercise 75: Checkpoint Cleanup
- Exercise 76: Incremental Saves

**Next**: [06-agent-graphs](../06-agent-graphs/lesson.md)
