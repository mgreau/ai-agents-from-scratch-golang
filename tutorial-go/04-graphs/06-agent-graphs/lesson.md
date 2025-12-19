# Agent Graphs: Multi-Agent Workflows

**Part 4: Graphs - Lesson 6**

> Use agents as graph nodes for complex workflows.

## Overview

**Agent graphs** combine agents with workflow logic.

```go
graph.AddNode("research", researchAgent)
graph.AddNode("write", writeAgent)
graph.AddNode("review", reviewAgent)

graph.AddConditionalEdge("review", func(state State) string {
    if state["approved"].(bool) {
        return "publish"
    }
    return "write" // Loop back
})
```

## Implementation

```go
func agentAsNode(agent Agent) NodeFunc {
    return func(state State) (State, error) {
        input := state["input"].(string)
        result, err := agent.Invoke(context.Background(), input, nil)
        if err != nil {
            return state, err
        }
        state["result"] = result
        return state, nil
    }
}

// Build multi-agent workflow
graph := NewStateGraph()
graph.AddNode("agent1", agentAsNode(agent1))
graph.AddNode("agent2", agentAsNode(agent2))
graph.AddNode("agent3", agentAsNode(agent3))

graph.AddConditionalEdge("agent1", func(state State) string {
    if needsAgent2(state) {
        return "agent2"
    }
    return "agent3"
})
```

## Examples

```go
// Research → Write → Review loop
researchAgent := NewReActAgent(llm, researchTools)
writeAgent := NewReActAgent(llm, writeTools)
reviewAgent := NewSimpleAgent(llm, reviewTools)

graph := NewStateGraph()
graph.AddNode("research", agentAsNode(researchAgent))
graph.AddNode("write", agentAsNode(writeAgent))
graph.AddNode("review", agentAsNode(reviewAgent))

graph.AddEdge("research", "write")
graph.AddEdge("write", "review")

graph.AddConditionalEdge("review", func(state State) string {
    approved := state["approved"].(bool)
    attempts := state["attempts"].(int)
    
    if approved {
        return "publish"
    }
    if attempts >= 3 {
        return "escalate"
    }
    state["attempts"] = attempts + 1
    return "write"
})
```

## Exercises

- Exercise 77: Multi-Agent Pipeline
- Exercise 78: Agent Feedback Loop
- Exercise 79: Conditional Agent Routing
- Exercise 80: Agent Collaboration

## Congratulations!

You've completed the **entire tutorial**! 🎉

You now know:
- ✅ Foundation (Runnables, Messages, LLMs, Config)
- ✅ Composition (Prompts, Parsers, Chains, Piping, Memory)
- ✅ Agency (Tools, Executors, Agents)
- ✅ Graphs (State Machines, Workflows)

**You're ready to build production AI applications in Go!**

## Further Reading

- [LangGraph Documentation](https://langchain-ai.github.io/langgraph/)
- [State Machines in Go](https://github.com/looplab/fsm)
- [Workflow Engines](https://temporal.io/)
