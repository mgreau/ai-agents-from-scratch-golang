# Setting Up go-llama.cpp

The Go implementation uses `github.com/go-skynet/go-llama.cpp` which requires building the underlying llama.cpp C++ library.

## Prerequisites

Before building the Go code, you need:

1. **C++ Compiler**
   - macOS: `xcode-select --install`
   - Linux: `sudo apt-get install build-essential`
   - Windows: MinGW-w64 or Visual Studio

2. **CMake** (version 3.14+)
   ```bash
   # macOS
   brew install cmake
   
   # Ubuntu/Debian
   sudo apt-get install cmake
   
   # Fedora
   sudo dnf install cmake
   ```

3. **Git** (for cloning llama.cpp)
   ```bash
   git --version
   ```

## Option 1: Build llama.cpp Manually (Recommended for Learning)

### Step 1: Clone and Build llama.cpp

```bash
# Clone llama.cpp
cd ~
git clone https://github.com/ggerganov/llama.cpp.git
cd llama.cpp

# Build the library
mkdir build
cd build
cmake ..
cmake --build . --config Release

# Install (optional, system-wide)
sudo cmake --install .
```

### Step 2: Set Environment Variables

```bash
# Add to ~/.bashrc, ~/.zshrc, or ~/.profile
export LLAMA_CPP_DIR="$HOME/llama.cpp"
export CGO_CFLAGS="-I${LLAMA_CPP_DIR}"
export CGO_LDFLAGS="-L${LLAMA_CPP_DIR}/build -lllama"
export LIBRARY_PATH="${LLAMA_CPP_DIR}/build:${LIBRARY_PATH}"
export LD_LIBRARY_PATH="${LLAMA_CPP_DIR}/build:${LD_LIBRARY_PATH}"

# Reload shell config
source ~/.bashrc  # or ~/.zshrc
```

### Step 3: Build Go Project

```bash
cd /path/to/ai-agents-from-scratch-go
make build
```

## Option 2: Use go-llama.cpp with Automatic Build

The `go-llama.cpp` library can attempt to build llama.cpp automatically, but this may not always work.

```bash
# Set build tags
export CGO_ENABLED=1
go build -tags llama_cpp_server ./examples-go/01_intro/
```

## Option 3: Use Pre-built Binaries (If Available)

Some distributions provide pre-built llama.cpp binaries:

```bash
# macOS with Homebrew (if available)
brew install llama.cpp

# Set paths to use Homebrew installation
export LLAMA_CPP_DIR="/opt/homebrew/opt/llama.cpp"
export CGO_CFLAGS="-I${LLAMA_CPP_DIR}/include"
export CGO_LDFLAGS="-L${LLAMA_CPP_DIR}/lib -lllama"
```

## Verification

Test if everything is set up correctly:

```bash
# Check if llama.cpp headers are found
echo $CGO_CFLAGS

# Try to build a simple example
cd examples-go/01_intro
go build .

# If successful, you should see the binary
ls -lh intro
```

## Troubleshooting

### Error: 'common.h' file not found

This means CGO can't find the llama.cpp headers.

**Solution:**
```bash
# Find where llama.cpp is installed
find ~ -name "common.h" -path "*/llama.cpp/*" 2>/dev/null

# Set CGO_CFLAGS to point to that directory
export CGO_CFLAGS="-I/path/to/llama.cpp"
```

### Error: library not found

This means the linker can't find the compiled llama.cpp library.

**Solution:**
```bash
# Find the compiled library
find ~ -name "libllama.*" 2>/dev/null

# Set CGO_LDFLAGS and library paths
export CGO_LDFLAGS="-L/path/to/llama.cpp/build -lllama"
export LD_LIBRARY_PATH="/path/to/llama.cpp/build:${LD_LIBRARY_PATH}"
```

### Error: undefined reference to llama_*

This means the library exists but isn't being linked properly.

**Solution:**
```bash
# Try static linking
export CGO_LDFLAGS="-L/path/to/llama.cpp/build -lllama -lstdc++"

# Or link all dependencies
export CGO_LDFLAGS="-L/path/to/llama.cpp/build -lllama -lm -lpthread"
```

## macOS Specific Issues

### Apple Silicon (M1/M2/M3)

```bash
# Build llama.cpp with Metal support for GPU acceleration
cd ~/llama.cpp
mkdir build
cd build
cmake .. -DLLAMA_METAL=ON
cmake --build . --config Release

# When building Go code
export CGO_LDFLAGS="-L${LLAMA_CPP_DIR}/build -lllama -framework Metal -framework Foundation -framework MetalKit"
```

### Homebrew Installation

```bash
# If using Homebrew
brew install llama.cpp

# Set paths (adjust based on your Homebrew installation)
export LLAMA_CPP_DIR="/opt/homebrew/opt/llama.cpp"
export CGO_CFLAGS="-I${LLAMA_CPP_DIR}/include"
export CGO_LDFLAGS="-L${LLAMA_CPP_DIR}/lib -lllama"
```

## Alternative: Use a Different Go Binding

If you encounter persistent issues with `go-skynet/go-llama.cpp`, consider these alternatives:

### 1. llama-cpp-go (Official bindings)

```bash
# Update go.mod
go get github.com/ggerganov/llama.cpp/bindings/go@latest

# Update pkg/llm/llama.go imports
# Change: github.com/go-skynet/go-llama.cpp
# To: github.com/ggerganov/llama.cpp/bindings/go
```

### 2. go-llama (Pure Go, slower but no CGO)

```bash
go get github.com/gotzmann/llama.go@latest
```

## Testing the Setup

Once everything is configured:

```bash
# Download a model if you haven't
make -C /path/to/project deps
huggingface-cli download Qwen/Qwen3-1.7B-GGUF Qwen3-1.7B-Q8_0.gguf \
  --local-dir models --local-dir-use-symlinks False

# Build and run
make build
make run-intro
```

If you see "Loading model from: ..." and then get a response, it's working!

## Docker Alternative (Easiest Setup)

If native setup is too complex, use Docker:

```dockerfile
# Dockerfile
FROM golang:1.23

RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    git

# Build llama.cpp
RUN git clone https://github.com/ggerganov/llama.cpp.git /llama.cpp && \
    cd /llama.cpp && \
    mkdir build && \
    cd build && \
    cmake .. && \
    cmake --build . --config Release

ENV LLAMA_CPP_DIR=/llama.cpp
ENV CGO_CFLAGS="-I${LLAMA_CPP_DIR}"
ENV CGO_LDFLAGS="-L${LLAMA_CPP_DIR}/build -lllama"
ENV LD_LIBRARY_PATH="${LLAMA_CPP_DIR}/build"

WORKDIR /app
COPY . .

RUN go mod download
RUN make build

CMD ["./bin/intro"]
```

```bash
# Build and run
docker build -t ai-agents-go .
docker run -v $(pwd)/models:/app/models ai-agents-go
```

## Summary

The key steps are:
1. ✅ Build llama.cpp C++ library
2. ✅ Set CGO environment variables
3. ✅ Build Go project
4. ✅ Run examples

**Estimated setup time:** 15-30 minutes (depending on download speed and build time)

## Need Help?

- Check llama.cpp documentation: https://github.com/ggerganov/llama.cpp
- Check go-llama.cpp issues: https://github.com/go-skynet/go-llama.cpp/issues
- Use Docker if native setup is problematic
