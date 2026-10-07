# HeliumCode Configuration Schema Generator

This tool generates a JSON Schema for the HeliumCode configuration file. The schema can be used to validate configuration files and provide autocompletion in editors that support JSON Schema.

## Usage

```bash
go run cmd/schema/main.go > helium-schema.json
```

This will generate a JSON Schema file that can be used to validate configuration files.

## Schema Features

The generated schema includes:

- All configuration options with descriptions
- Default values where applicable
- Validation for enum values (e.g., model IDs, provider types)
- Required fields
- Type checking

## Using the Schema

You can use the generated schema in several ways:

1. **Editor Integration**: Many editors (VS Code, JetBrains IDEs, etc.) support JSON Schema for validation and autocompletion. You can configure your editor to use the generated schema for `.helium.json` files.

2. **Validation Tools**: You can use tools like [jsonschema](https://github.com/Julian/jsonschema) to validate your configuration files against the schema.

3. **Documentation**: The schema serves as documentation for the configuration options.

## Example Configuration

Here's an example configuration that conforms to the schema:

```json
{
  "data": {
    "directory": ".helium"
  },
  "debug": false,
  "providers": {
    "anthropic": {
      "apiKey": "your-api-key"
    }
  },
  "agents": {
    "coder": {
      "model": "claude-sonnet-5.5"
    },
    "task": {
      "model": "claude-sonnet-5.5"
    },
    "title": {
      "model": "claude-haiku-4.5",
      "maxTokens": 80
    }
  }
}
```

Model IDs must match one of the values in the schema's `enum`, which is generated from the live model catalog in `internal/llm/models`.
