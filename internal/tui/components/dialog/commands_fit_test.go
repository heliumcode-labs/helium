package dialog

import (
	"regexp"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/require"
)

var fitANSI = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func fitStripANSI(s string) string { return fitANSI.ReplaceAllString(s, "") }

func TestCommandDialogFitsTerminal(t *testing.T) {
	cmds := []Command{
		{ID: "a", Title: "New Session", Description: "Start a fresh conversation and keep the current one around for later"},
		{ID: "b", Title: "Compact Session", Description: "Summarize the current session and create a new one with the summary"},
		{ID: "c", Title: "Switch Model", Description: "Pick a different model for this session"},
		{ID: "d", Title: "Theme", Description: "Change the colour theme"},
		{ID: "e", Title: "Init", Description: "Create a project instructions file"},
		{ID: "f", Title: "One more", Description: "Another entry so the list scrolls"},
		{ID: "g", Title: "Last", Description: "The final entry"},
	}

	for _, size := range []struct{ w, h int }{
		{36, 14}, {40, 18}, {40, 30}, {60, 24}, {100, 40}, {120, 60},
	} {
		d := NewCommandDialogCmp()
		_, _ = d.Update(tea.WindowSizeMsg{Width: size.w, Height: size.h})
		d.SetCommands(cmds)
		plain := fitStripANSI(d.View())
		lines := strings.Split(plain, "\n")

		maxLine := 0
		for _, l := range lines {
			if w := lipgloss.Width(l); w > maxLine {
				maxLine = w
			}
		}

		require.LessOrEqual(t, maxLine, size.w,
			"commands dialog wider than terminal (%dx%d)", size.w, size.h)
		require.LessOrEqual(t, len(lines), size.h,
			"commands dialog taller than terminal (%dx%d): %d lines", size.w, size.h, len(lines))
	}
}
