# Exercise 02: JSON Parser Runnable

Build a Runnable that parses JSON strings into Go structs.

## 🎯 Objective

Learn to handle errors and type conversions in Runnables.

## 📚 Concepts

- Error handling in Runnables
- JSON parsing in Go
- Type assertions
- Returning structured data

## 📝 Task

Implement `JSONParserRunnable` that:

1. Takes a JSON string as input
2. Parses it into a map[string]interface{}
3. Returns the parsed data
4. Handles invalid JSON gracefully

## ✅ Requirements

- [ ] Embed `*core.BaseRunnable`
- [ ] Parse JSON using encoding/json
- [ ] Return error for invalid JSON
- [ ] Handle both string and []byte input
- [ ] Pass all tests

## 💡 Hints

1. Use `json.Unmarshal()` to parse JSON
2. Check input type (string or []byte)
3. Return descriptive errors
4. Test with various JSON structures

## 🧪 Tests

```go
func TestJSONParser(t *testing.T) {
    parser := NewJSONParserRunnable()
    
    // Valid JSON
    result, err := parser.Invoke(ctx, `{"name": "John", "age": 30}`, nil)
    // Should return: map[string]interface{}{"name": "John", "age": 30}
    
    // Invalid JSON
    _, err = parser.Invoke(ctx, `{invalid json}`, nil)
    // Should return error
}
```

## 🚀 Next Steps

After completing, move to Exercise 03: Pipeline Composition
