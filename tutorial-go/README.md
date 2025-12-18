# Go Tutorial - Building AI Agents Step by Step

This tutorial walks you through building AI agent components in Go, from basic concepts to advanced patterns.

## 🎯 Learning Path

### Phase 1: Foundations (Start Here)

Work through the examples in order:

1. **[01_intro](../examples-go/01_intro/)** - Basic LLM interaction
2. **[03_translation](../examples-go/03_translation/)** - System prompts
3. **[04_think](../examples-go/04_think/)** - Reasoning
4. **[05_batch](../examples-go/05_batch/)** - Parallel processing with goroutines
5. **[06_coding](../examples-go/06_coding/)** - Streaming with channels
6. **[07_simple-agent](../examples-go/07_simple-agent/)** - Tools and function calling
7. **[09_react-agent](../examples-go/09_react-agent/)** - ReAct pattern

### Phase 2: Deep Dive (Advanced)

Learn the architecture by building components:

1. **[01-core-patterns](01-core-patterns/)** - Runnable, Messages, Context
2. **[02-llm-integration](02-llm-integration/)** - LLM wrapper patterns
3. **[03-tools-system](03-tools-system/)** - Building custom tools
4. **[04-agent-patterns](04-agent-patterns/)** - Agent architectures
5. **[05-advanced-concepts](05-advanced-concepts/)** - Memory, chains, graphs

## 📚 How to Use This Tutorial

### For Beginners

1. **Start with examples-go/** - Run and study the working code
2. **Read inline comments** - Each example has detailed explanations
3. **Experiment** - Modify parameters and prompts
4. **Build something** - Create your own tool or agent

### For Intermediate Developers

1. **Study pkg/ implementations** - See how components work
2. **Read the advanced tutorials** - Deep dive into patterns
3. **Complete exercises** - Practice implementing features
4. **Extend the framework** - Add new capabilities

### For Advanced Developers

1. **Implement missing features** - Memory systems, chains, graphs
2. **Optimize performance** - Benchmarking and profiling
3. **Production deployment** - Docker, monitoring, scaling
4. **Contribute back** - Share your improvements

## 🛠️ Tutorial Structure

Each advanced tutorial includes:

- **README.md** - Concept explanation and learning objectives
- **starter/** - Starting point with TODO markers
- **solution/** - Complete implementation
- **tests/** - Test cases to verify your implementation
- **examples/** - Practical usage examples

## 🎓 Learning Objectives

By completing this tutorial, you'll understand:

### Go-Specific Skills
- ✅ Interface-based design in Go
- ✅ Context for cancellation and timeouts
- ✅ Channels for streaming data
- ✅ Goroutines for concurrency
- ✅ Error handling patterns
- ✅ Resource management with defer

### AI Agent Concepts
- ✅ LLM interaction and prompting
- ✅ Tool/function calling
- ✅ Agent reasoning patterns (ReAct)
- ✅ Message-driven architecture
- ✅ State management
- ✅ Composable components

### Architecture Patterns
- ✅ Runnable interface pattern
- ✅ Builder pattern for configuration
- ✅ Registry pattern for tools
- ✅ Strategy pattern for agents
- ✅ Observer pattern for callbacks

## 📖 Recommended Reading Order

### Week 1-2: Basics
- Run all examples in examples-go/
- Read pkg/core/ implementations
- Understand Runnable pattern

### Week 3-4: Tools & Agents
- Study pkg/tools/ and pkg/agents/
- Implement custom tools
- Build a simple agent

### Week 5-6: Advanced
- Complete advanced tutorials
- Implement memory systems
- Build chains and graphs

### Week 7-8: Project
- Build a complete application
- Deploy to production
- Optimize performance

## 🚀 Getting Started

```bash
# 1. Complete setup first (see QUICKSTART_GO.md)
make deps
make build

# 2. Run examples to understand concepts
make run-intro
make run-agent
make run-react

# 3. Study the implementations
cd pkg/core
# Read runnable.go, message.go, context.go

# 4. Try advanced tutorials
cd tutorial-go/01-core-patterns
# Follow the README.md
```

## 💡 Tips for Success

1. **Understand before implementing** - Read the concept first
2. **Type the code yourself** - Don't just copy/paste
3. **Break things** - Experiment to learn
4. **Read error messages** - Go's compiler helps you learn
5. **Use the debugger** - Step through code to understand flow
6. **Write tests** - Verify your understanding
7. **Ask questions** - Check issues or discussions

## 🤝 Contributing

Found a better way to explain something? Want to add a tutorial?

1. Create a new tutorial in `tutorial-go/XX-topic-name/`
2. Include starter, solution, and tests
3. Update this README
4. Submit a pull request

## 📚 Additional Resources

### Go Learning
- [Go Tour](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)

### AI/LLM Concepts
- [Prompt Engineering Guide](https://www.promptingguide.ai/)
- [ReAct Paper](https://arxiv.org/abs/2210.03629)
- [LangChain Concepts](https://docs.langchain.com/docs/)

### Architecture Patterns
- [Design Patterns in Go](https://refactoring.guru/design-patterns/go)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)

## 🎯 Next Steps

After completing this tutorial:

1. **Build a real project** - Apply what you learned
2. **Explore frameworks** - Try LangChainGo, ChromaDB, etc.
3. **Optimize** - Profile and improve performance
4. **Deploy** - Put your agent in production
5. **Share** - Write about your experience

Happy learning! 🚀
