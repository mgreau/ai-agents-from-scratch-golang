# Part 4: Graphs

> **Building complex workflows with state machines and conditional routing**

Welcome to Part 4, the final section! You've mastered Foundation, Composition, and Agency. Now you'll learn to build **complex workflows** using graphs - state machines that can branch, loop, and make decisions dynamically.

## What are Graphs?

Graphs are **state machines** where nodes are operations and edges are transitions. They enable:
1. Complex workflows with branching logic
2. Conditional routing based on state
3. Parallel execution paths
4. Checkpointing and recovery
5. Agents that make dynamic decisions

**Example:**
```go
graph := NewStateGraph()
graph.AddNode("analyze", analyzeNode)
graph.AddNode("approve", approveNode)
graph.AddNode("reject", rejectNode)

// Conditional routing based on analysis result
graph.AddConditionalEdge("analyze", func(state State) string {
    if state["score"].(float64) > 0.8 {
        return "approve"
    }
    return "reject"
})

result := graph.Execute(ctx, initialState)
```

## Why Graphs Matter

### Without Graphs: Hardcoded Logic

```go
func processDocument(doc string) string {
    result := analyze(doc)
    if result.score > 0.8 {
        if result.needsReview {
            review := humanReview(result)
            if review.approved {
                return publish(result)
            }
            return archive(result)
        }
        return publish(result)
    }
    return reject(result)
}
// Complex nested conditionals, hard to modify
```

### With Graphs: Declarative Workflows

```go
graph := NewStateGraph()
graph.AddNode("analyze", analyzeNode)
graph.AddNode("review", reviewNode)
graph.AddNode("publish", publishNode)
graph.AddNode("reject", rejectNode)

graph.AddConditionalEdge("analyze", routeAfterAnalysis)
graph.AddConditionalEdge("review", routeAfterReview)

// Clear, modifiable, testable
result := graph.Execute(ctx, doc)
```

## What You'll Learn

### 1. **State Basics** (Lesson 1)
Nodes, edges, and state machines.

**Skills:**
- Define graph nodes
- Connect with edges
- Manage state through graph
- Build simple linear workflows

**Example:**
```go
graph.AddNode("step1", func(state State) (State, error) {
    state["result"] = process(state["input"])
    return state, nil
})
graph.AddEdge("step1", "step2")
```

### 2. **Channels** (Lesson 2)
State management and data routing.

**Skills:**
- Channel-based state updates
- Parallel state branches
- Merge multiple states
- Isolate state per path

**Example:**
```go
graph.AddNode("parallel1", node1)
graph.AddNode("parallel2", node2)
graph.AddNode("merge", mergeNode)
graph.AddEdge("start", "parallel1")
graph.AddEdge("start", "parallel2")
graph.AddEdge("parallel1", "merge")
graph.AddEdge("parallel2", "merge")
```

### 3. **Conditional Edges** (Lesson 3)
Dynamic routing based on state.

**Skills:**
- Route based on state values
- Implement decision points
- Handle multiple branches
- Default routing

**Example:**
```go
graph.AddConditionalEdge("decision", func(state State) string {
    score := state["score"].(float64)
    if score > 0.9 {
        return "high_priority"
    } else if score > 0.5 {
        return "medium_priority"
    }
    return "low_priority"
})
```

### 4. **Executor** (Lesson 4)
Running complex workflows efficiently.

**Skills:**
- Execute graphs step-by-step
- Handle errors gracefully
- Support parallel execution
- Stream execution progress

**Example:**
```go
executor := NewGraphExecutor(graph)
result, err := executor.Execute(ctx, initialState, ExecutorConfig{
    Parallel: true,
    Stream:   true,
})
```

### 5. **Checkpointing** (Lesson 5)
Persistence and recovery for long-running workflows.

**Skills:**
- Save state at checkpoints
- Resume from checkpoint
- Handle interruptions
- Implement retry logic

**Example:**
```go
executor := NewCheckpointingExecutor(graph, storage)
result, err := executor.Execute(ctx, state)
// If interrupted, can resume:
result, err = executor.Resume(ctx, checkpointID)
```

### 6. **Agent Graphs** (Lesson 6)
Agents as graph nodes with complex workflows.

**Skills:**
- Use agents as nodes
- Multi-agent collaboration
- Conditional agent routing
- Agent feedback loops

**Example:**
```go
graph.AddNode("research", researchAgent)
graph.AddNode("write", writeAgent)
graph.AddNode("review", reviewAgent)

graph.AddConditionalEdge("review", func(state State) string {
    if state["approved"].(bool) {
        return "publish"
    }
    return "write" // Loop back to rewrite
})
```

## The Graph Patterns

```
State → Node → Node → Node → Result (linear)
                ↓
State → Node → [Branch] → Merge → Result (parallel)
                ↓
State → Node → Decision → [Multiple Paths] → Result (conditional)
                ↓
State → Agent → Agent → Agent → Result (agent collaboration)
```

## Go Advantages for Graphs

**Goroutines for Parallelism:**
```go
// Execute multiple nodes concurrently
var wg sync.WaitGroup
for _, node := range parallelNodes {
    wg.Add(1)
    go func(n Node) {
        defer wg.Done()
        executeNode(ctx, n, state)
    }(node)
}
wg.Wait()
```

**Channels for State:**
```go
// Stream state updates
stateChan := make(chan State, 10)
go graph.Execute(ctx, state, stateChan)
for state := range stateChan {
    fmt.Printf("State update: %+v\n", state)
}
```

**Context for Control:**
```go
// Cancel long-running graph
ctx, cancel := context.WithCancel(ctx)
go graph.Execute(ctx, state)
// Later: cancel()
```

## Prerequisites

- ✅ All of Part 1 (Foundation)
- ✅ All of Part 2 (Composition)
- ✅ All of Part 3 (Agency)

## Lessons

### [Lesson 1: State Basics](01-state-basics/lesson.md)
Nodes, edges, and simple state machines.

### [Lesson 2: Channels](02-channels/lesson.md)
State management and parallel routing.

### [Lesson 3: Conditional Edges](03-conditional-edges/lesson.md)
Dynamic routing based on state values.

### [Lesson 4: Executor](04-executor/lesson.md)
Running complex workflows efficiently.

### [Lesson 5: Checkpointing](05-checkpointing/lesson.md)
Persistence and recovery for long workflows.

### [Lesson 6: Agent Graphs](06-agent-graphs/lesson.md)
Multi-agent collaboration in workflows.

## Final Projects

After completing all lessons:

1. **Document Processing Pipeline**
   - Upload → Analyze → OCR/Parse → Review → Publish
   - Conditional routing based on document type
   - Human-in-the-loop for review
   - Checkpointing for long documents

2. **Multi-Agent Research System**
   - Research agent → Write agent → Review agent
   - Loop back if review fails
   - Parallel research on multiple topics
   - Checkpoint after each agent

3. **Customer Support Workflow**
   - Classify → Route → Handle → Follow-up
   - Different agents for different issue types
   - Escalation paths
   - Persistent state across interactions

4. **Code Generation Pipeline**
   - Plan → Generate → Test → Review → Deploy
   - Conditional routing based on test results
   - Parallel generation for multiple files
   - Checkpointing for long generations

## What's Next

This is the **final section** of the tutorial. After completing Part 4, you'll have mastered:
- ✅ Foundation patterns
- ✅ Composition techniques
- ✅ Autonomous agents
- ✅ Complex workflows

You'll be ready to build production AI applications in Go!

---

Ready? Start with [Lesson 1: State Basics](01-state-basics/lesson.md) →
