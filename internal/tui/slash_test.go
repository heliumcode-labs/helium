package tui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/heliumcode-labs/helium/internal/tui/components/dialog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The welcome screen advertises "/help"; these guard that a leading slash is
// dispatched as a command instead of being sent to the model.

func TestSlashCommandHelp(t *testing.T) {
	a := appModel{}
	out, _ := a.runSlashCommand("/help")

	m, ok := out.(appModel)
	require.True(t, ok)
	assert.True(t, m.showHelp)
}

func TestSlashCommandQuitIsCaseInsensitive(t *testing.T) {
	a := appModel{}
	out, _ := a.runSlashCommand("  /Quit ")

	m, ok := out.(appModel)
	require.True(t, ok)
	assert.True(t, m.showQuit)
}

func TestSlashCommandUnknownWarns(t *testing.T) {
	a := appModel{}
	out, cmd := a.runSlashCommand("/nope")

	_, ok := out.(appModel)
	require.True(t, ok)
	require.NotNil(t, cmd, "unknown commands must surface a warning")
}

func TestSlashCommandFallsBackToRegisteredCommand(t *testing.T) {
	type marker struct{}
	ran := false

	a := appModel{commands: []dialog.Command{{
		ID:    "mycmd",
		Title: "My Command",
		Handler: func(dialog.Command) tea.Cmd {
			return func() tea.Msg {
				ran = true
				return marker{}
			}
		},
	}}}

	out, cmd := a.runSlashCommand("/mycmd")
	_, ok := out.(appModel)
	require.True(t, ok)
	require.NotNil(t, cmd)

	_, isMarker := cmd().(marker)
	assert.True(t, isMarker)
	assert.True(t, ran)
}
