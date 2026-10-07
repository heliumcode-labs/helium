package tools

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTodoTool_RunStoresAndListsTodos(t *testing.T) {
	tool := NewTodoTool()

	write, err := tool.Run(context.Background(), ToolCall{
		Name:  TodoToolName,
		Input: `{"todos":[{"content":"Read the config package","status":"in_progress","active_form":"Reading the config package"},{"content":"Rewrite the README","status":"pending"}]}`,
	})
	require.NoError(t, err)
	require.False(t, write.IsError, write.Content)

	assert.Contains(t, write.Content, "Todo list (0/2 completed):")
	assert.Contains(t, write.Content, "1. [>] Read the config package")
	assert.Contains(t, write.Content, "2. [ ] Rewrite the README")

	metadata := TodoResponseMetadata{}
	require.NoError(t, json.Unmarshal([]byte(write.Metadata), &metadata))
	assert.Equal(t, 2, metadata.Total)
	assert.Equal(t, 1, metadata.InProgress)
	assert.Equal(t, 1, metadata.Pending)

	// A read call (no todos) returns the stored list.
	read, err := tool.Run(context.Background(), ToolCall{Name: TodoToolName, Input: `{}`})
	require.NoError(t, err)
	assert.Contains(t, read.Content, "1. [>] Read the config package")
}

func TestTodoTool_RunCompletesTodos(t *testing.T) {
	tool := NewTodoTool()

	_, err := tool.Run(context.Background(), ToolCall{
		Name:  TodoToolName,
		Input: `{"todos":[{"content":"step one","status":"completed"},{"content":"step two","status":"pending"}]}`,
	})
	require.NoError(t, err)

	resp, err := tool.Run(context.Background(), ToolCall{
		Name:  TodoToolName,
		Input: `{"todos":[{"content":"step one","status":"completed"},{"content":"step two","status":"completed"}]}`,
	})
	require.NoError(t, err)
	require.False(t, resp.IsError, resp.Content)

	metadata := TodoResponseMetadata{}
	require.NoError(t, json.Unmarshal([]byte(resp.Metadata), &metadata))
	assert.Equal(t, 2, metadata.Completed)
	assert.Contains(t, resp.Content, "(2/2 completed)")
}

func TestTodoTool_RunValidatesInput(t *testing.T) {
	tool := NewTodoTool()

	t.Run("rejects unknown status", func(t *testing.T) {
		resp, err := tool.Run(context.Background(), ToolCall{
			Name:  TodoToolName,
			Input: `{"todos":[{"content":"x","status":"done"}]}`,
		})
		require.NoError(t, err)
		assert.True(t, resp.IsError)
		assert.Contains(t, resp.Content, "status must be one of")
	})

	t.Run("rejects empty content", func(t *testing.T) {
		resp, err := tool.Run(context.Background(), ToolCall{
			Name:  TodoToolName,
			Input: `{"todos":[{"content":"   ","status":"pending"}]}`,
		})
		require.NoError(t, err)
		assert.True(t, resp.IsError)
		assert.Contains(t, resp.Content, "content is required")
	})

	t.Run("rejects malformed json", func(t *testing.T) {
		resp, err := tool.Run(context.Background(), ToolCall{
			Name:  TodoToolName,
			Input: `{"todos":`,
		})
		require.NoError(t, err)
		assert.True(t, resp.IsError)
	})
}

func TestTodoTool_EmptyList(t *testing.T) {
	tool := NewTodoTool()

	resp, err := tool.Run(context.Background(), ToolCall{Name: TodoToolName, Input: `{"todos":[]}`})
	require.NoError(t, err)
	require.False(t, resp.IsError, resp.Content)
	assert.Equal(t, "Todo list is empty.", resp.Content)
}

func TestFormatTodos(t *testing.T) {
	assert.Equal(t, "Todo list is empty.", formatTodos(nil))
	assert.Equal(t, "Todo list (1/1 completed):\n1. [x] ship it", formatTodos([]TodoItem{
		{Content: "ship it", Status: TodoStatusCompleted},
	}))
}
