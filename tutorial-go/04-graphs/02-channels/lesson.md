# Channels: State Management and Routing

**Part 4: Graphs - Lesson 2**

> Parallel execution and state merging.

## Overview

**Channels** enable parallel node execution and state merging in graphs.

```go
// Parallel execution
graph.AddNode("parallel1", node1)
graph.AddNode("parallel2", node2)
graph.AddEdge("start", "parallel1")
graph.AddEdge("start", "parallel2")
graph.AddNode("merge", mergeResults)
graph.AddEdge("parallel1", "merge")
graph.AddEdge("parallel2", "merge")
```

## Implementation

### Parallel Execution

```go
func (g *StateGraph) ExecuteParallel(ctx context.Context, state State) (State, error) {
    currentNode := g.startNode
    
    for currentNode != "" {
        nodeFn := g.nodes[currentNode]
        state, _ = nodeFn(state)
        
        nextNodes := g.edges[currentNode]
        if len(nextNodes) == 0 {
            break
        }
        
        // If multiple edges, execute in parallel
        if len(nextNodes) > 1 {
            results := make([]State, len(nextNodes))
            var wg sync.WaitGroup
            
            for i, next := range nextNodes {
                wg.Add(1)
                go func(idx int, node string) {
                    defer wg.Done()
                    results[idx], _ = g.executeSubgraph(ctx, node, state)
                }(i, next)
            }
            wg.Wait()
            
            // Merge results
            state = mergeStates(results)
            break
        }
        
        currentNode = nextNodes[0]
    }
    
    return state, nil
}
```

## Exercises

- Exercise 61: Parallel Processing
- Exercise 62: State Merging
- Exercise 63: Channel Buffering
- Exercise 64: Error Aggregation

**Next**: [03-conditional-edges](../03-conditional-edges/lesson.md)
