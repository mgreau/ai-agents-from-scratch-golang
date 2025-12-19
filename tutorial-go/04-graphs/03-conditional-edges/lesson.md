# Conditional Edges: Dynamic Routing

**Part 4: Graphs - Lesson 3**

> Route based on state values.

## Overview

**Conditional edges** route to different nodes based on state.

```go
graph.AddConditionalEdge("decision", func(state State) string {
    if state["score"].(float64) > 0.8 {
        return "approve"
    }
    return "reject"
})
```

## Implementation

```go
type ConditionFunc func(state State) string

type StateGraph struct {
    // ... existing fields
    conditionalEdges map[string]ConditionFunc
}

func (g *StateGraph) AddConditionalEdge(from string, condition ConditionFunc) {
    g.conditionalEdges[from] = condition
}

func (g *StateGraph) Execute(ctx context.Context, state State) (State, error) {
    currentNode := g.startNode
    
    for currentNode != "" {
        // Execute node
        nodeFn := g.nodes[currentNode]
        state, _ = nodeFn(state)
        
        // Check for conditional edge
        if condition, exists := g.conditionalEdges[currentNode]; exists {
            nextNode := condition(state)
            currentNode = nextNode
            continue
        }
        
        // Use regular edge
        nextNodes := g.edges[currentNode]
        if len(nextNodes) == 0 {
            break
        }
        currentNode = nextNodes[0]
    }
    
    return state, nil
}
```

## Exercises

- Exercise 65: Score-Based Routing
- Exercise 66: Multi-Branch Decisions
- Exercise 67: Default Routes
- Exercise 68: Loop Detection

**Next**: [04-executor](../04-executor/lesson.md)
