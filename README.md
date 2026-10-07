# He HeliumCode

> **He** — the noble gas that lifts everything. HeliumCode is a fast, terminal-native AI coding agent.

A powerful terminal-based AI assistant for developers, providing intelligent coding assistance directly in your terminal.

## Overview

HeliumCode is a Go-based CLI application that brings AI assistance to your terminal. It provides a TUI (Terminal User Interface) for interacting with various AI models to help with coding tasks, debugging, and more.

## Features

- **Interactive TUI**: Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) with the `helium` theme — a deep-ink and neon-mint palette designed for long sessions — plus a welcome screen, live context meter and cost tracking in the status bar
- **Multiple AI Providers**: Anthropic, OpenAI, Google Gemini, GitHub Copilot, xAI, Groq, OpenRouter, AWS Bedrock, Azure OpenAI, Google Cloud VertexAI and any OpenAI-compatible local endpoint
- **Current Model Catalog**: Claude Fable/Sonnet 5.5, GPT-6.1 & GPT-5.6 families, Gemini 3.x, Grok 4.7 and more, with fresh context windows, output limits and pricing
- **Session Management**: Save and manage multiple conversation sessions
- **Tool Integration**: AI can execute commands, search files and the web, track its own task list, and modify code
- **Task Tracking**: Built-in `todo` tool lets the assistant plan multi-step work and report progress
- **Web Search**: Keyless `websearch` tool (DuckDuckGo + Wikipedia fallback) for current information
- **Vim-like Editor**: Integrated editor with text input capabilities
- **Persistent Storage**: SQLite database for storing conversations and sessions
- **LSP Integration**: Language Server Protocol support for code intelligence
- **File Change Tracking**: Track and visualize file changes during sessions
- **External Editor Support**: Open your preferred editor for composing messages
- **Named Arguments for Custom Commands**: Create powerful custom commands with multiple named placeholders

## Installation

### Using the Install Script

```bash
# Install the latest version
curl -fsSL https://raw.githubusercontent.com/heliumcode-labs/helium/refs/heads/main/install | bash

# Install a specific version
curl -fsSL https://raw.githubusercontent.com/heliumcode-labs/helium/refs/heads/main/install | VERSION=0.1.0 bash
```

> The installer downloads a signed-checksum-verified archive into `~/.helium/bin` and
> adds it to your `PATH` for bash, zsh, fish, and sh.

### Using Go

```bash
go install github.com/heliumcode-labs/helium@latest
```

## Configuration

HeliumCode looks for configuration in the following locations:

- `$HOME/.helium.json`
- `$XDG_CONFIG_HOME/helium/.helium.json`
- `./.helium.json` (local directory)

### Auto Compact Feature

HeliumCode includes an auto compact feature that automatically summarizes your conversation when it approaches the model's context window limit. When enabled (default setting), this feature:

- Monitors token usage during your conversation
- Automatically triggers summarization when usage reaches 95% of the model's context window
- Creates a new session with the summary, allowing you to continue your work without losing context
- Helps prevent "out of context" errors that can occur with long conversations

You can enable or disable this feature in your configuration file:

```json
{
  "autoCompact": true // default is true
}
```

### Environment Variables

You can configure HeliumCode using environment variables:

| Environment Variable       | Purpose                                                                          |
| -------------------------- | -------------------------------------------------------------------------------- |
| `ANTHROPIC_API_KEY`        | For Claude models                                                                |
| `OPENAI_API_KEY`           | For OpenAI models                                                                |
| `GEMINI_API_KEY`           | For Google Gemini models                                                         |
| `GITHUB_TOKEN`             | For Github Copilot models (see [Using Github Copilot](#using-github-copilot))    |
| `VERTEXAI_PROJECT`         | For Google Cloud VertexAI (Gemini)                                               |
| `VERTEXAI_LOCATION`        | For Google Cloud VertexAI (Gemini)                                               |
| `GROQ_API_KEY`             | For Groq models                                                                  |
| `OPENROUTER_API_KEY`       | For OpenRouter models                                                            |
| `XAI_API_KEY`              | For xAI Grok models                                                              |
| `AWS_ACCESS_KEY_ID`        | For AWS Bedrock (Claude)                                                         |
| `AWS_SECRET_ACCESS_KEY`    | For AWS Bedrock (Claude)                                                         |
| `AWS_REGION`               | For AWS Bedrock (Claude)                                                         |
| `AZURE_OPENAI_ENDPOINT`    | For Azure OpenAI models                                                          |
| `AZURE_OPENAI_API_KEY`     | For Azure OpenAI models (optional when using Entra ID)                           |
| `AZURE_OPENAI_API_VERSION` | For Azure OpenAI models                                                          |
| `LOCAL_ENDPOINT`           | For self-hosted models                                                           |
| `SHELL`                    | Default shell to use (if not specified in config)                                |

### Shell Configuration

HeliumCode allows you to configure the shell used by the bash tool. By default, it uses the shell specified in the `SHELL` environment variable, or falls back to `/bin/bash` if not set.

You can override this in your configuration file:

```json
{
  "shell": {
    "path": "/bin/zsh",
    "args": ["-l"]
  }
}
```

This is useful if you want to use a different shell than your default system shell, or if you need to pass specific arguments to the shell.

### Configuration File Structure

```json
{
  "data": {
    "directory": ".helium"
  },
  "providers": {
    "openai": {
      "apiKey": "your-api-key",
      "disabled": false
    },
    "anthropic": {
      "apiKey": "your-api-key",
      "disabled": false
    },
    "copilot": {
      "disabled": false
    },
    "groq": {
      "apiKey": "your-api-key",
      "disabled": false
    },
    "openrouter": {
      "apiKey": "your-api-key",
      "disabled": false
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
  },
  "shell": {
    "path": "/bin/bash",
    "args": ["-l"]
  },
  "mcpServers": {
    "example": {
      "type": "stdio",
      "command": "path/to/mcp-server",
      "env": [],
      "args": []
    }
  },
  "lsp": {
    "go": {
      "disabled": false,
      "command": "gopls"
    }
  },
  "debug": false,
  "debugLSP": false,
  "autoCompact": true
}
```

## Supported AI Models

HeliumCode supports a variety of AI models from different providers. Every model carries its context window, output limit and pricing, which the status bar uses for the live context meter and cost display.

Run `ctrl+o` in the TUI to browse the models available for your configured providers.

### Anthropic

- Claude Fable 5.1 / 5
- Claude Opus 5.5 / 5 / 4.8 / 4.5
- Claude Sonnet 5.5 / 5 / 4.6 / 4.5
- Claude Haiku 4.5

### OpenAI

- GPT-6.1 Sol, GPT-6 Astra, GPT-6 Luna
- GPT-5.6 Sol / Terra / Luna
- GPT-5.5, GPT-5.4 (mini, nano), GPT-5, GPT-5 mini/nano
- GPT-4.1 family (gpt-4.1, gpt-4.1-mini, gpt-4.1-nano)
- GPT-4o family (gpt-4o, gpt-4o-mini)
- O3, O4 Mini

### Google Gemini

- Gemini 3.1 Pro
- Gemini 3.8 / 3.5 / 3.1 Flash (and Flash-Lite variants)
- Gemini 3 Flash
- Gemini 2.5 Pro / Flash / Flash Lite

### GitHub Copilot

- GPT-6.1 Sol, GPT-6 Astra/Luna, GPT-5.6 family, GPT-5.5/5.4/5, GPT-4.1, GPT-4o, o4-mini
- Claude Fable, Opus, Sonnet and Haiku (current generations)
- Gemini 3.8 / 3.5 Flash
- Grok 4.7 / 4.6 / 4.5

### xAI

- Grok 4.7 / 4.6 / 4.5 / 4.3

### Groq

- OpenAI GPT-OSS 120B / 20B
- Qwen3 6.27B / 8.27B
- Llama 3.3 70B Versatile
- Llama 3.1 8B Instant

### OpenRouter

- Anthropic Claude (Fable 5.1, Opus 5.5, Sonnet 5.5/5/4.6, Haiku 4.5)
- OpenAI GPT-6.1 Sol / GPT-6 Luna / GPT-5.6 / GPT-5 / GPT-4.1
- Google Gemini 3.1 Pro, 3.8 Flash, 2.5 Pro
- xAI Grok 4.7
- DeepSeek V3.2

### AWS Bedrock

- Claude Fable 5.1, Opus 5.5, Sonnet 5.5 / 4.6, Haiku 4.5

### Azure OpenAI

Mirrors the OpenAI catalog (GPT-6.1, GPT-5.6, GPT-5.5/5.4/5, GPT-4.1, GPT-4o, O3, O4 Mini).

### Google Cloud VertexAI

Mirrors the Gemini catalog (3.1 Pro, 3.8/3.5 Flash, 2.5 Pro/Flash/Flash Lite).

## Usage

```bash
# Start HeliumCode
helium

# Start with debug logging
helium -d

# Start with a specific working directory
helium -c /path/to/project
```

## Non-interactive Prompt Mode

You can run HeliumCode in non-interactive mode by passing a prompt directly as a command-line argument. This is useful for scripting, automation, or when you want a quick answer without launching the full TUI.

```bash
# Run a single prompt and print the AI's response to the terminal
helium -p "Explain the use of context in Go"

# Get response in JSON format
helium -p "Explain the use of context in Go" -f json

# Run without showing the spinner (useful for scripts)
helium -p "Explain the use of context in Go" -q
```

In this mode, HeliumCode will process your prompt, print the result to standard output, and then exit. All permissions are auto-approved for the session.

By default, a spinner animation is displayed while the model is processing your query. You can disable this spinner with the `-q` or `--quiet` flag, which is particularly useful when running HeliumCode from scripts or automated workflows.

### Output Formats

HeliumCode supports the following output formats in non-interactive mode:

| Format | Description                     |
| ------ | ------------------------------- |
| `text` | Plain text output (default)     |
| `json` | Output wrapped in a JSON object |

The output format is implemented as a strongly-typed `OutputFormat` in the codebase, ensuring type safety and validation when processing outputs.

## Command-line Flags

| Flag              | Short | Description                                         |
| ----------------- | ----- | --------------------------------------------------- |
| `--help`          | `-h`  | Display help information                            |
| `--debug`         | `-d`  | Enable debug mode                                   |
| `--cwd`           | `-c`  | Set current working directory                       |
| `--prompt`        | `-p`  | Run a single prompt in non-interactive mode         |
| `--output-format` | `-f`  | Output format for non-interactive mode (text, json) |
| `--quiet`         | `-q`  | Hide spinner in non-interactive mode                |

## Keyboard Shortcuts

### Global Shortcuts

| Shortcut | Action                                                  |
| -------- | ------------------------------------------------------- |
| `Ctrl+C` | Quit application                                        |
| `Ctrl+?` | Toggle help dialog                                      |
| `?`      | Toggle help dialog (when not in editing mode)           |
| `Ctrl+L` | View logs                                               |
| `Ctrl+S` | Switch session                                          |
| `Ctrl+K` | Command dialog                                          |
| `Ctrl+F` | Select files to upload                                  |
| `Ctrl+O` | Toggle model selection dialog                           |
| `Ctrl+T` | Switch theme                                            |
| `Esc`    | Close current overlay/dialog or return to previous mode |

### Chat Page Shortcuts

| Shortcut   | Action                                  |
| ---------- | --------------------------------------- |
| `Ctrl+N`   | Create new session                      |
| `Esc`      | Cancel current operation/generation     |
| `@`        | Open file/model completion              |
| `i`        | Focus editor (when not in writing mode) |
| `Esc`      | Exit writing mode and focus messages    |

### Editor Shortcuts

| Shortcut            | Action                                    |
| ------------------- | ----------------------------------------- |
| `Ctrl+S`            | Send message (when editor is focused)     |
| `Enter` or `Ctrl+S` | Send message (when editor is not focused) |
| `Ctrl+E`            | Open external editor                      |
| `Esc`               | Blur editor and focus messages            |

### Session Dialog Shortcuts

| Shortcut   | Action           |
| ---------- | ---------------- |
| `↑` or `k` | Previous session |
| `↓` or `j` | Next session     |
| `Enter`    | Select session   |
| `Esc`      | Close dialog     |

### Model Dialog Shortcuts

| Shortcut   | Action            |
| ---------- | ----------------- |
| `↑` or `k` | Move up           |
| `↓` or `j` | Move down         |
| `←` or `h` | Previous provider |
| `→` or `l` | Next provider     |
| `Esc`      | Close dialog      |

### Permission Dialog Shortcuts

| Shortcut                | Action                       |
| ----------------------- | ---------------------------- |
| `←` or `left`           | Switch options left          |
| `→` or `right` or `tab` | Switch options right         |
| `Enter` or `space`      | Confirm selection            |
| `a`                     | Allow permission             |
| `A`                     | Allow permission for session |
| `d`                     | Deny permission              |

### Logs Page Shortcuts

| Shortcut           | Action              |
| ------------------ | ------------------- |
| `Backspace` or `q` | Return to chat page |

## AI Assistant Tools

HeliumCode's AI assistant has access to various tools to help with coding tasks:

### File and Code Tools

| Tool          | Description                 | Parameters                                                                               |
| ------------- | --------------------------- | ---------------------------------------------------------------------------------------- |
| `glob`        | Find files by pattern       | `pattern` (required), `path` (optional)                                                  |
| `grep`        | Search file contents        | `pattern` (required), `path` (optional), `include` (optional), `literal_text` (optional) |
| `ls`          | List directory contents     | `path` (optional), `ignore` (optional array of patterns)                                 |
| `view`        | View file contents          | `file_path` (required), `offset` (optional), `limit` (optional)                          |
| `write`       | Write to files              | `file_path` (required), `content` (required)                                             |
| `edit`        | Edit files                  | Various parameters for file editing                                                      |
| `patch`       | Apply patches to files      | `file_path` (required), `diff` (required)                                                |
| `diagnostics` | Get diagnostics information | `file_path` (optional)                                                                   |

### Other Tools

| Tool          | Description                            | Parameters                                                                                |
| ------------- | -------------------------------------- | ----------------------------------------------------------------------------------------- |
| `bash`        | Execute shell commands                 | `command` (required), `timeout` (optional)                                                |
| `fetch`       | Fetch data from URLs                   | `url` (required), `format` (required), `timeout` (optional)                               |
| `websearch`   | Search the web (no API key required)   | `query` (required), `max_results` (optional)                                              |
| `todo`        | Track multi-step tasks                 | `todos` (array of `{content, status, active_form}`, omit to read the list)                 |
| `sourcegraph` | Search code across public repositories | `query` (required), `count` (optional), `context_window` (optional), `timeout` (optional) |
| `agent`       | Run sub-tasks with the AI agent        | `prompt` (required)                                                                       |

## Themes

HeliumCode ships with a fresh visual design: the `helium` theme (deep-ink background, neon-mint primary, violet and coral accents) is the default, alongside a set of community favorites.

Switch themes at any time with `ctrl+t`, or set one permanently in `.helium.json`:

```json
{
  "tui": {
    "theme": "helium"
  }
}
```

Available themes: `helium` (default), `catppuccin`, `dracula`, `flexoki`, `gruvbox`, `monokai`, `onedark`, `tokyonight`, `tron`.

## Architecture

HeliumCode is built with a modular architecture:

- **cmd**: Command-line interface using Cobra
- **internal/app**: Core application services
- **internal/config**: Configuration management
- **internal/db**: Database operations and migrations
- **internal/llm**: LLM providers and tools integration
- **internal/tui**: Terminal UI components and layouts
- **internal/logging**: Logging infrastructure
- **internal/message**: Message handling
- **internal/session**: Session management
- **internal/lsp**: Language Server Protocol integration

## Custom Commands

HeliumCode supports custom commands that can be created by users to quickly send predefined prompts to the AI assistant.

### Creating Custom Commands

Custom commands are predefined prompts stored as Markdown files in one of three locations:

1. **User Commands** (prefixed with `user:`):

   ```
   $XDG_CONFIG_HOME/helium/commands/
   ```

   (typically `~/.config/helium/commands/` on Linux/macOS)

   or

   ```
   $HOME/.helium/commands/
   ```

2. **Project Commands** (prefixed with `project:`):

   ```
   <PROJECT DIR>/.helium/commands/
   ```

Each `.md` file in these directories becomes a custom command. The file name (without extension) becomes the command ID.

For example, creating a file at `~/.config/helium/commands/prime-context.md` with content:

```markdown
RUN git ls-files
READ README.md
```

This creates a command called `user:prime-context`.

### Command Arguments

HeliumCode supports named arguments in custom commands using placeholders in the format `$NAME` (where NAME consists of uppercase letters, numbers, and underscores, and must start with a letter).

For example:

```markdown
# Fetch Context for Issue $ISSUE_NUMBER

RUN gh issue view $ISSUE_NUMBER --json title,body,comments
RUN git grep --author="$AUTHOR_NAME" -n .
RUN grep -R "$SEARCH_PATTERN" $DIRECTORY
```

When you run a command with arguments, HeliumCode will prompt you to enter values for each unique placeholder. Named arguments provide several benefits:

- Clear identification of what each argument represents
- Ability to use the same argument multiple times
- Better organization for commands with multiple inputs

### Organizing Commands

You can organize commands in subdirectories:

```
~/.config/helium/commands/git/commit.md
```

This creates a command with ID `user:git:commit`.

### Using Custom Commands

1. Press `Ctrl+K` to open the command dialog
2. Select your custom command (prefixed with either `user:` or `project:`)
3. Press Enter to execute the command

The content of the command file will be sent as a message to the AI assistant.

### Built-in Commands

HeliumCode includes several built-in commands:

| Command            | Description                                                                                         |
| ------------------ | --------------------------------------------------------------------------------------------------- |
| Initialize Project | Creates or updates the HELIUM.md memory file with project-specific information                    |
| Compact Session    | Manually triggers the summarization of the current session, creating a new session with the summary |

## MCP (Model Context Protocol)

HeliumCode implements the Model Context Protocol (MCP) to extend its capabilities through external tools. MCP provides a standardized way for the AI assistant to interact with external services and tools.

### MCP Features

- **External Tool Integration**: Connect to external tools and services via a standardized protocol
- **Tool Discovery**: Automatically discover available tools from MCP servers
- **Multiple Connection Types**:
  - **Stdio**: Communicate with tools via standard input/output
  - **SSE**: Communicate with tools via Server-Sent Events
- **Security**: Permission system for controlling access to MCP tools

### Configuring MCP Servers

MCP servers are defined in the configuration file under the `mcpServers` section:

```json
{
  "mcpServers": {
    "example": {
      "type": "stdio",
      "command": "path/to/mcp-server",
      "env": [],
      "args": []
    },
    "web-example": {
      "type": "sse",
      "url": "https://example.com/mcp",
      "headers": {
        "Authorization": "Bearer token"
      }
    }
  }
}
```

### MCP Tool Usage

Once configured, MCP tools are automatically available to the AI assistant alongside built-in tools. They follow the same permission model as other tools, requiring user approval before execution.

## LSP (Language Server Protocol)

HeliumCode integrates with Language Server Protocol to provide code intelligence features across multiple programming languages.

### LSP Features

- **Multi-language Support**: Connect to language servers for different programming languages
- **Diagnostics**: Receive error checking and linting information
- **File Watching**: Automatically notify language servers of file changes

### Configuring LSP

Language servers are configured in the configuration file under the `lsp` section:

```json
{
  "lsp": {
    "go": {
      "disabled": false,
      "command": "gopls"
    },
    "typescript": {
      "disabled": false,
      "command": "typescript-language-server",
      "args": ["--stdio"]
    }
  }
}
```

### LSP Integration with AI

The AI assistant can access LSP features through the `diagnostics` tool, allowing it to:

- Check for errors in your code
- Suggest fixes based on diagnostics

While the LSP client implementation supports the full LSP protocol (including completions, hover, definition, etc.), currently only diagnostics are exposed to the AI assistant.

## Using Github Copilot

_Copilot support is currently experimental._

### Requirements
- [Copilot chat in the IDE](https://github.com/settings/copilot) enabled in GitHub settings
- One of:
  - VSCode Github Copilot chat extension
  - Github `gh` CLI
  - Neovim Github Copilot plugin (`copilot.vim` or `copilot.lua`)
  - Github token with copilot permissions

If using one of the above plugins or cli tools, make sure you use the authenticate
the tool with your github account. This should create a github token at one of the following locations:
- ~/.config/github-copilot/[hosts,apps].json
- $XDG_CONFIG_HOME/github-copilot/[hosts,apps].json

If using an explicit github token, you may either set the $GITHUB_TOKEN environment variable or add it to the helium.json config file at `providers.copilot.apiKey`.

## Using a self-hosted model provider

HeliumCode can also load and use models from a self-hosted (OpenAI-like) provider.
This is useful for developers who want to experiment with custom models.

### Configuring a self-hosted provider

You can use a self-hosted model by setting the `LOCAL_ENDPOINT` environment variable.
This will cause HeliumCode to load and use the models from the specified endpoint.

```bash
LOCAL_ENDPOINT=http://localhost:1235/v1
```

### Configuring a self-hosted model

You can also configure a self-hosted model in the configuration file under the `agents` section:

```json
{
  "agents": {
    "coder": {
      "model": "local.granite-3.3-2b-instruct@q8_0",
      "reasoningEffort": "high"
    }
  }
}
```

## Development

### Prerequisites

- Go 1.24.0 or higher

### Building from Source

```bash
# Clone the repository
git clone https://github.com/heliumcode-labs/helium.git
cd helium

# Build
go build -o helium

# Run
./helium
```

## Acknowledgments

HeliumCode gratefully acknowledges the contributions and support from these key individuals:

- [@isaacphi](https://github.com/isaacphi) - For the [mcp-language-server](https://github.com/isaacphi/mcp-language-server) project which provided the foundation for our LSP client implementation
- [@adamdottv](https://github.com/adamdottv) - For the design direction and UI/UX architecture

Special thanks to the broader open source community whose tools and libraries have made this project possible.

## License

HeliumCode is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Here's how you can contribute:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please make sure to update tests as appropriate and follow the existing code style.
