package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

const (
	TodoToolName    = "todo"
	todoDescription = `Manages a persistent, ordered task list for the current session.

Call this tool to break a request into explicit steps, track what you are working on right now and mark steps as you finish them. The list survives between tool calls in the same session, so it is the shared source of truth for progress.

WHEN TO USE THIS TOOL:
- The user asks for multi-step work: plan the steps first, then execute them one by one
- You are starting a step: mark it in_progress
- You finished a step: mark it completed and move the next one to in_progress
- The user changes direction or asks what is left

HOW TO USE:
- todos: the COMPLETE list in its new state (always send every item, never a partial diff)
- Each item needs:
  - content: short imperative description, e.g. "Add the websearch tool"
  - status: pending, in_progress or completed
  - active_form: present-continuous form used while the item runs, e.g. "Adding the websearch tool"
- Call the tool with no todos parameter to read the current list

RULES:
- Keep the list short: at most 10 items, one line each
- Keep exactly one item in_progress at a time
- Mark an item completed only when it really is; never batch-complete unfinished work
- Rewrite the whole list on every call`
)

// Supported todo item statuses.
const (
	TodoStatusPending    = "pending"
	TodoStatusInProgress = "in_progress"
	TodoStatusCompleted  = "completed"
)

const maxTodos = 50

type TodoItem struct {
	Content    string `json:"content"`
	Status     string `json:"status"`
	ActiveForm string `json:"active_form,omitempty"`
}

type TodoParams struct {
	Todos []TodoItem `json:"todos"`
}

type TodoResponseMetadata struct {
	Total      int `json:"total"`
	Completed  int `json:"completed"`
	InProgress int `json:"in_progress"`
	Pending    int `json:"pending"`
}

// todoStore keeps the list per session so the tool stays stateless for callers.
var todoStore = struct {
	sync.RWMutex
	items map[string][]TodoItem
}{
	items: make(map[string][]TodoItem),
}

type todoTool struct{}

func NewTodoTool() BaseTool {
	return &todoTool{}
}

func (t *todoTool) Info() ToolInfo {
	return ToolInfo{
		Name:        TodoToolName,
		Description: todoDescription,
		Parameters: map[string]any{
			"todos": map[string]any{
				"type":        "array",
				"description": "The complete task list in its new state. Omit to read the current list.",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"content": map[string]any{
							"type":        "string",
							"description": "Short imperative description of the task",
						},
						"status": map[string]any{
							"type":        "string",
							"enum":        []string{TodoStatusPending, TodoStatusInProgress, TodoStatusCompleted},
							"description": "Current state of the task",
						},
						"active_form": map[string]any{
							"type":        "string",
							"description": "Present-continuous form shown while the task is running",
						},
					},
					"required": []string{"content", "status"},
				},
			},
		},
	}
}

func (t *todoTool) Run(ctx context.Context, call ToolCall) (ToolResponse, error) {
	var params TodoParams
	if err := json.Unmarshal([]byte(call.Input), &params); err != nil {
		return NewTextErrorResponse(fmt.Sprintf("error parsing parameters: %s", err)), nil
	}

	sessionID, _ := GetContextValues(ctx)
	if sessionID == "" {
		sessionID = "default"
	}

	if params.Todos != nil {
		normalized, err := normalizeTodos(params.Todos)
		if err != nil {
			return NewTextErrorResponse(err.Error()), nil
		}

		todoStore.Lock()
		todoStore.items[sessionID] = normalized
		todoStore.Unlock()
	}

	todoStore.RLock()
	current := append([]TodoItem(nil), todoStore.items[sessionID]...)
	todoStore.RUnlock()

	metadata := TodoResponseMetadata{}
	for _, item := range current {
		metadata.Total++
		switch item.Status {
		case TodoStatusCompleted:
			metadata.Completed++
		case TodoStatusInProgress:
			metadata.InProgress++
		default:
			metadata.Pending++
		}
	}

	return WithResponseMetadata(NewTextResponse(formatTodos(current)), metadata), nil
}

func normalizeTodos(todos []TodoItem) ([]TodoItem, error) {
	if len(todos) > maxTodos {
		return nil, fmt.Errorf("too many todos: %d (max %d)", len(todos), maxTodos)
	}

	normalized := make([]TodoItem, 0, len(todos))
	for i, item := range todos {
		content := strings.TrimSpace(item.Content)
		if content == "" {
			return nil, fmt.Errorf("todos[%d]: content is required", i)
		}

		status := strings.ToLower(strings.TrimSpace(item.Status))
		if status == "" {
			status = TodoStatusPending
		}
		switch status {
		case TodoStatusPending, TodoStatusInProgress, TodoStatusCompleted:
		default:
			return nil, fmt.Errorf("todos[%d]: status must be one of %s, %s or %s",
				i, TodoStatusPending, TodoStatusInProgress, TodoStatusCompleted)
		}

		normalized = append(normalized, TodoItem{
			Content:    content,
			Status:     status,
			ActiveForm: strings.TrimSpace(item.ActiveForm),
		})
	}

	return normalized, nil
}

func formatTodos(todos []TodoItem) string {
	if len(todos) == 0 {
		return "Todo list is empty."
	}

	var done int
	for _, item := range todos {
		if item.Status == TodoStatusCompleted {
			done++
		}
	}

	var b strings.Builder
	fmt.Fprintf(&b, "Todo list (%d/%d completed):\n", done, len(todos))
	for i, item := range todos {
		marker := "[ ]"
		switch item.Status {
		case TodoStatusInProgress:
			marker = "[>]"
		case TodoStatusCompleted:
			marker = "[x]"
		}
		fmt.Fprintf(&b, "%d. %s %s\n", i+1, marker, item.Content)
	}

	return strings.TrimRight(b.String(), "\n")
}
