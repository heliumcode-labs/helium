package agent

import (
	"context"

	"github.com/heliumcode-labs/helium/internal/history"
	"github.com/heliumcode-labs/helium/internal/llm/tools"
	"github.com/heliumcode-labs/helium/internal/lsp"
	"github.com/heliumcode-labs/helium/internal/message"
	"github.com/heliumcode-labs/helium/internal/permission"
	"github.com/heliumcode-labs/helium/internal/session"
)

func CoderAgentTools(
	permissions permission.Service,
	sessions session.Service,
	messages message.Service,
	history history.Service,
	lspClients map[string]*lsp.Client,
) []tools.BaseTool {
	ctx := context.Background()
	otherTools := GetMcpTools(ctx, permissions)
	if len(lspClients) > 0 {
		otherTools = append(otherTools, tools.NewDiagnosticsTool(lspClients))
	}
	return append(
		[]tools.BaseTool{
			tools.NewBashTool(permissions),
			tools.NewEditTool(lspClients, permissions, history),
			tools.NewFetchTool(permissions),
			tools.NewGlobTool(),
			tools.NewGrepTool(),
			tools.NewLsTool(),
			tools.NewSourcegraphTool(),
			tools.NewTodoTool(),
			tools.NewViewTool(lspClients),
			tools.NewPatchTool(lspClients, permissions, history),
			tools.NewWebSearchTool(permissions),
			tools.NewWriteTool(lspClients, permissions, history),
			NewAgentTool(sessions, messages, lspClients),
		}, otherTools...,
	)
}

func TaskAgentTools(lspClients map[string]*lsp.Client) []tools.BaseTool {
	return []tools.BaseTool{
		tools.NewGlobTool(),
		tools.NewGrepTool(),
		tools.NewLsTool(),
		tools.NewSourcegraphTool(),
		tools.NewTodoTool(),
		tools.NewViewTool(lspClients),
	}
}
