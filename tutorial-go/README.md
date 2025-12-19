# AI Agents From Scratch - Complete Tutorial

> **A comprehensive guide to building production-ready AI agents in Go**

## 🎓 Welcome!

This tutorial teaches you to build AI agents in Go from first principles. No prior AI experience required - just Go knowledge and curiosity.

## 📖 Tutorial Structure

The tutorial is organized into 4 progressive parts:

```
Foundation → Composition → Agency → Graphs
   ↓            ↓            ↓         ↓
Patterns    Components    Agents   Workflows
```

### 🏗️ Part 1: Foundation
**Learn the core patterns** (4 lessons, 16 exercises, 135 pages)

Start here to understand the fundamental abstractions:
- [01-core-patterns/](01-core-patterns/) - Runnables, Messages, LLMs, Config

**Time:** 1-2 weeks

**You'll build:** Simple LLM chat application with callbacks

---

### 🔧 Part 2: Composition  
**Build reusable components** (5 lessons, 20 exercises, 175 pages)

Learn to compose complex behaviors from simple pieces:
- [02-composition/](02-composition/) - Prompts, Parsers, Chains, Piping, Memory

**Time:** 2-3 weeks

**You'll build:** Translation service with conversation memory

---

### 🤖 Part 3: Agency
**Create autonomous agents** (5 lessons, 20 exercises, 175 pages)

Build agents that use tools and make decisions:
- [03-agency/](03-agency/) - Tools, Executors, Simple/ReAct/Structured Agents

**Time:** 2-3 weeks

**You'll build:** Multi-tool assistant with ReAct reasoning

---

### 🌐 Part 4: Graphs
**Design complex workflows** (6 lessons, 24 exercises, 180 pages)

Master state machines and multi-agent workflows:
- [04-graphs/](04-graphs/) - State Machines, Channels, Conditional Routing, Checkpointing

**Time:** 2-3 weeks

**You'll build:** Document processing pipeline with checkpoints

---

## 🚀 Getting Started

### Prerequisites

- **Go 1.21+** installed
- Basic Go knowledge (interfaces, goroutines, channels)
- Text editor or IDE
- (Optional) Local LLM model for running examples

### Start Learning

1. **Clone this repository**
   ```bash
   git clone https://github.com/mgreau/ai-agents-from-scratch-golang.git
   cd ai-agents-from-scratch-golang/tutorial-go
   ```

2. **Start with Part 1**
   ```bash
   cd 01-core-patterns
   # Read README.md, then start with lesson 01-runnable
   ```

3. **Do the exercises**
   - Each lesson has 4 exercises
   - Start with `starter.go`, try to solve it
   - Check `solution.go` when stuck or to compare approaches

4. **Run the examples**
   ```bash
   cd ../../examples-go/01_intro
   go run main.go
   ```

---

## 📊 Full Contents

### Part 1: Foundation (Weeks 1-2)

| Lesson | Topic | Exercises | Pages |
|--------|-------|-----------|-------|
| 1 | [Runnable Pattern](01-core-patterns/01-runnable/lesson.md) | 4 | 30 |
| 2 | [Messages & Types](01-core-patterns/02-messages/lesson.md) | 4 | 40 |
| 3 | [LLM Wrapper](01-core-patterns/03-llm-wrapper/lesson.md) | 4 | 35 |
| 4 | [Context & Config](01-core-patterns/04-context/lesson.md) | 4 | 30 |

**Total:** 16 exercises, 135 pages

---

### Part 2: Composition (Weeks 3-4)

| Lesson | Topic | Exercises | Pages |
|--------|-------|-----------|-------|
| 1 | [Prompts](02-composition/01-prompts/lesson.md) | 4 | 45 |
| 2 | [Parsers](02-composition/02-parsers/lesson.md) | 4 | 40 |
| 3 | [LLM Chain](02-composition/03-llm-chain/lesson.md) | 4 | 30 |
| 4 | [Piping](02-composition/04-piping/lesson.md) | 4 | 30 |
| 5 | [Memory](02-composition/05-memory/lesson.md) | 4 | 30 |

**Total:** 20 exercises, 175 pages

---

### Part 3: Agency (Weeks 5-6)

| Lesson | Topic | Exercises | Pages |
|--------|-------|-----------|-------|
| 1 | [Tools](03-agency/01-tools/lesson.md) | 4 | 35 |
| 2 | [Tool Executor](03-agency/02-tool-executor/lesson.md) | 4 | 30 |
| 3 | [Simple Agent](03-agency/03-simple-agent/lesson.md) | 4 | 30 |
| 4 | [ReAct Agent](03-agency/04-react-agent/lesson.md) | 4 | 40 |
| 5 | [Structured Agent](03-agency/05-structured-agent/lesson.md) | 4 | 40 |

**Total:** 20 exercises, 175 pages

---

### Part 4: Graphs (Weeks 7-8)

| Lesson | Topic | Exercises | Pages |
|--------|-------|-----------|-------|
| 1 | [State Basics](04-graphs/01-state-basics/lesson.md) | 4 | 30 |
| 2 | [Channels](04-graphs/02-channels/lesson.md) | 4 | 30 |
| 3 | [Conditional Edges](04-graphs/03-conditional-edges/lesson.md) | 4 | 30 |
| 4 | [Executor](04-graphs/04-executor/lesson.md) | 4 | 30 |
| 5 | [Checkpointing](04-graphs/05-checkpointing/lesson.md) | 4 | 30 |
| 6 | [Agent Graphs](04-graphs/06-agent-graphs/lesson.md) | 4 | 30 |

**Total:** 24 exercises, 180 pages

---

## 🎯 Learning Objectives

By completing this tutorial, you will:

### Technical Skills
- ✅ Build composable AI components with Go interfaces
- ✅ Manage conversation state with typed messages
- ✅ Integrate LLMs (go-llama.cpp) with proper resource management
- ✅ Create prompt templates with variable substitution
- ✅ Parse structured data from LLM outputs
- ✅ Implement multi-step reasoning agents (ReAct)
- ✅ Build complex workflows with state machines
- ✅ Use goroutines and channels for concurrency

### Production Patterns
- ✅ Error handling and recovery
- ✅ Timeout protection
- ✅ Resource cleanup
- ✅ Observability (logging, metrics, tracing)
- ✅ Checkpointing for long workflows
- ✅ Type-safe tool definitions

### Real-World Applications
- ✅ Document processing pipelines
- ✅ Customer support chatbots
- ✅ Code generation assistants
- ✅ Research agents
- ✅ Multi-agent collaboration systems

---

## 💡 How to Use This Tutorial

### 1. **Linear Path (Recommended)**
Follow the parts in order: Foundation → Composition → Agency → Graphs

Each part builds on previous concepts, so skipping ahead may be confusing.

### 2. **Topic-Based Path**
If you're experienced, jump to specific topics:
- Need agents? Start at Part 3
- Need workflows? Start at Part 4
- But review Foundation first for core concepts

### 3. **Exercise-Driven Path**
Read lesson → Do exercises immediately → Check solutions → Move on

Don't skip exercises - they reinforce learning.

---

## 🔨 Practice Tips

1. **Type the code yourself** - Don't copy/paste, type it out
2. **Modify examples** - Change parameters, add features
3. **Break things** - See what errors look like
4. **Read solutions** - Even if you solved it, compare approaches
5. **Build projects** - Apply concepts to real problems

---

## 📚 Additional Resources

### In This Repository
- [COMPLETION_SUMMARY.md](COMPLETION_SUMMARY.md) - Comprehensive overview
- [../examples-go/](../examples-go/) - 7 working examples
- [../pkg/](../pkg/) - Complete implementations

### External Resources
- [Go Documentation](https://go.dev/doc/)
- [go-llama.cpp](https://github.com/go-skynet/go-llama.cpp)
- [LangChain Concepts](https://python.langchain.com/docs/concepts)
- [ReAct Paper](https://arxiv.org/abs/2210.03629)

---

## 🤝 Contributing

Found a typo? Have a suggestion? Contributions welcome!

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

---

## 📄 License

See LICENSE file in root directory.

---

## 🎉 Ready to Start?

Begin your journey: **[Part 1: Foundation →](01-core-patterns/)**

Questions? Issues? Open an issue on GitHub!

**Happy Learning! 🚀**
