# Contributing to MCP Go Utilities Server

First off, thank you for considering contributing to this project! 🎉

## Code of Conduct

This project aims to be welcoming to everyone. Please be respectful and constructive in your interactions.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check existing issues. When creating a bug report, include:

* Use a clear and descriptive title
* Describe the exact steps to reproduce the problem
* Provide specific examples
* Describe the behavior you observed and what you expected
* Include Go version, OS, and any relevant details

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion:

* Use a clear and descriptive title
* Provide a detailed description of the suggested enhancement
* Explain why this enhancement would be useful
* List any examples of similar features in other projects

### Pull Requests

1. Fork the repo and create your branch from `main`
2. Add tests for any new functionality
3. Ensure the test suite passes (`make test`)
4. Format your code (`make fmt`)
5. Run the linter (`make lint`)
6. Update documentation as needed
7. Write a clear commit message

## Development Setup

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/mcpserver.git
cd mcpserver

# Install dependencies
go mod download

# Run tests
make test

# Run the server
make run
```

## Coding Standards

* Follow standard Go conventions and idioms
* Write clear, self-documenting code
* Add comments for complex logic
* Keep functions focused and small
* Write tests for new functionality
* Maintain test coverage above 80%

## Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run linter
make lint
```

## Adding New Tools

To add a new MCP tool:

1. Add the tool implementation in `tools/tools.go`
2. Add corresponding tests in `tools/tools_test.go`
3. Register the tool in `main.go`
4. Update the README with the new tool documentation

Example:

```go
// In tools/tools.go
type MyToolInput struct {
    Value string `json:"value" jsonschema:"required,description=Input value"`
}

type MyToolOutput struct {
    Result string `json:"result"`
}

func MyTool(_ context.Context, _ mcp.CallToolRequest, in MyToolInput) (MyToolOutput, error) {
    // Implementation
    return MyToolOutput{Result: in.Value}, nil
}

// In main.go
myTool := mcp.NewTool(
    "mytool",
    mcp.WithDescription("Description of my tool"),
    mcp.WithInputSchema[tools.MyToolInput](),
    mcp.WithOutputSchema[tools.MyToolOutput](),
    mcp.WithString("value", mcp.Required(), mcp.Description("Input value")),
)
srv.AddTool(myTool, mcp.NewStructuredToolHandler(tools.MyTool))
```

## Commit Messages

* Use the present tense ("Add feature" not "Added feature")
* Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
* Limit the first line to 72 characters or less
* Reference issues and pull requests after the first line

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
