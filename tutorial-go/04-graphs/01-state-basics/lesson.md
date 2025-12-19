# State Basics: Nodes, Edges, and State Machines

**Part 4: Graphs - Lesson 1**

> The foundation of workflow graphs.

## Overview

A **State Graph** is a workflow where:
- **Nodes** are operations that transform state
- **Edges** connect nodes
- **State** flows through the graph

```go
type StateGraph struct {
    nodes map[string]NodeFunc
    edges map[string][]string
}

type NodeFunc func(state State) (State, error)
type State map[string]interface{}
```

## Core Implementation

### Graph Structure

```go
package graphs

type StateGraph struct {
    nodes     map[string]NodeFunc
    edges     map[string][]string
    startNode string
}

func NewStateGraph() *StateGraph {
    return &StateGraph{
        nodes: make(map[string]NodeFunc),
        edges: make(map[string][]string),
    }
}

func (g *StateGraph) AddNode(name string, fn NodeFunc) {
    g.nodes[name] = fn
}

func (g *StateGraph) AddEdge(from, to string) {
    g.edges[from] = append(g.edges[from], to)
}

func (g *StateGraph) SetStart(node string) {
    g.startNode = node
}
```

### Execution

```go
func (g *StateGraph) Execute(ctx context.Context, initialState State) (State, error) {
    currentNode := g.startNode
    state := initialState
    
    for currentNode != "" {
        // Execute node
        nodeFn, exists := g.nodes[currentNode]
        if !exists {
            return state, fmt.Errorf("node not found: %s", currentNode)
        }
        
        newState, err := nodeFn(state)
        if err != nil {
            return state, err
        }
        state = newState
        
        // Get next node
        nextNodes := g.edges[currentNode]
        if len(nextNodes) == 0 {
            break // End of graph
        }
        currentNode = nextNodes[0]
    }
    
    return state, nil
}
```

## Examples

```go
// Linear workflow
graph := NewStateGraph()
graph.AddNode("load", loadData)
graph.AddNode("process", processData)
graph.AddNode("save", saveData)
graph.AddEdge("load", "process")
graph.AddEdge("process", "save")
graph.SetStart("load")

result, _ := graph.Execute(ctx, State{"path": "data.json"})
```

## Exercises

- Exercise 57: Linear Pipeline
- Exercise 58: Node Error Handling
- Exercise 59: State Validation
- Exercise 60: Graph Visualization

## Key Takeaways

1. ✅ Nodes transform state
2. ✅ Edges connect nodes
3. ✅ State flows through graph
4. ✅ Simple execution model

**Next**: [02-channels](../02-channels/lesson.md) - Parallel state routing
