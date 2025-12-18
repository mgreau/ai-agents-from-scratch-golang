.PHONY: help build run-intro run-translation run-coding run-agent run-react clean test

help:
	@echo "AI Agents From Scratch - Go Edition"
	@echo ""
	@echo "Available commands:"
	@echo "  make build              - Build all examples"
	@echo "  make run-intro          - Run intro example (basic LLM)"
	@echo "  make run-translation    - Run translation example (system prompts)"
	@echo "  make run-think          - Run think example (reasoning)"
	@echo "  make run-batch          - Run batch example (parallel processing)"
	@echo "  make run-coding         - Run coding example (streaming)"
	@echo "  make run-agent          - Run simple agent example (tools)"
	@echo "  make run-react          - Run ReAct agent example"
	@echo "  make run-all            - Run all examples in sequence"
	@echo "  make clean              - Clean build artifacts"
	@echo "  make test               - Run tests"
	@echo "  make deps               - Download dependencies"

deps:
	go mod download
	go mod tidy

build:
	@echo "Building examples..."
	@mkdir -p bin
	@cd examples-go/01_intro && go build -o ../../bin/intro .
	@cd examples-go/03_translation && go build -o ../../bin/translation .
	@cd examples-go/04_think && go build -o ../../bin/think .
	@cd examples-go/05_batch && go build -o ../../bin/batch .
	@cd examples-go/06_coding && go build -o ../../bin/coding .
	@cd examples-go/07_simple-agent && go build -o ../../bin/simple-agent .
	@cd examples-go/09_react-agent && go build -o ../../bin/react-agent .
	@echo "Build complete! Binaries in ./bin/"

run-intro: build
	@echo "Running intro example..."
	@./bin/intro

run-translation: build
	@echo "Running translation example..."
	@./bin/translation

run-think: build
	@echo "Running think example (reasoning)..."
	@./bin/think

run-batch: build
	@echo "Running batch example (parallel processing)..."
	@./bin/batch

run-coding: build
	@echo "Running coding example (streaming)..."
	@./bin/coding

run-agent: build
	@echo "Running simple agent example..."
	@./bin/simple-agent

run-react: build
	@echo "Running ReAct agent example..."
	@./bin/react-agent

run-all: build
	@echo "Running all examples..."
	@echo "\n=== 01: Intro ==="
	@./bin/intro
	@echo "\n=== 03: Translation ==="
	@./bin/translation
	@echo "\n=== 04: Think ==="
	@./bin/think
	@echo "\n=== 05: Batch ==="
	@./bin/batch
	@echo "\n=== 06: Coding ==="
	@./bin/coding
	@echo "\n=== 07: Simple Agent ==="
	@./bin/simple-agent
	@echo "\n=== 09: ReAct Agent ==="
	@./bin/react-agent

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@echo "Clean complete!"

test:
	@echo "Running tests..."
	@go test -v ./pkg/...

fmt:
	@echo "Formatting code..."
	@go fmt ./...

lint:
	@echo "Running linter..."
	@golangci-lint run ./...

install-tools:
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
