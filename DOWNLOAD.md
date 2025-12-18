# Download Models for Go Edition

Download the GGUF models used in this repository. These models are compatible with llama.cpp and can be used with the Go implementation.

## Quantization Levels

You can adjust the quantization level to balance model precision and file size:
- `:Q8_0` - Higher precision and better output quality, requires more memory and storage
- `:Q6_K` - Good balance between size and accuracy (recommended default)
- `:Q5_K_S` - Smaller model, loads faster, uses less memory, slightly lower precision

## Download Options

### Option 1: Using huggingface-cli (Recommended for Go)

```bash
# Install huggingface-cli if you don't have it
pip install huggingface-hub

# Download Qwen3-1.7B (recommended for examples)
mkdir -p models
huggingface-cli download Qwen/Qwen3-1.7B-GGUF Qwen3-1.7B-Q8_0.gguf --local-dir models --local-dir-use-symlinks False

# Or download other models
huggingface-cli download giladgd/gpt-oss-20b-GGUF gpt-oss-20b.MXFP4.gguf --local-dir models --local-dir-use-symlinks False
huggingface-cli download unsloth/DeepSeek-R1-0528-Qwen3-8B-GGUF DeepSeek-R1-0528-Qwen3-8B-Q6_K.gguf --local-dir models --local-dir-use-symlinks False
```

### Option 2: Manual Download

Visit these URLs and download directly to the `models/` directory:
- [Qwen3-1.7B-Q8_0.gguf](https://huggingface.co/Qwen/Qwen3-1.7B-GGUF/resolve/main/Qwen3-1.7B-Q8_0.gguf)
- [Other models on Hugging Face](https://huggingface.co/models?library=gguf)

### Option 3: Using node-llama-cpp (if you have Node.js)

```bash
npx --no node-llama-cpp pull --dir ./models hf:Qwen/Qwen3-1.7B-GGUF:Q8_0
npx --no node-llama-cpp pull --dir ./models hf:giladgd/gpt-oss-20b-GGUF/gpt-oss-20b.MXFP4.gguf
npx --no node-llama-cpp pull --dir ./models hf:unsloth/DeepSeek-R1-0528-Qwen3-8B-GGUF:Q6_K --filename DeepSeek-R1-0528-Qwen3-8B-Q6_K.gguf
```

## Verify Installation

After downloading, verify your models:

```bash
ls -lh models/
```

You should see `.gguf` files in the models directory. The examples are configured to use `Qwen3-1.7B-Q8_0.gguf` by default.


