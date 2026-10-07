package chat

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestElementTile(t *testing.T) {
	tile := elementTile()

	assert.Contains(t, tile, "He")
	assert.Contains(t, tile, "4.0026")
	assert.Contains(t, tile, "helium")
	require.Equal(t, 6, lipgloss.Height(tile))
}

func TestBannerHeader(t *testing.T) {
	out := bannerHeader(100)

	assert.Contains(t, out, "HeliumCode")
	assert.Contains(t, out, "https://github.com/heliumcode-labs/helium")
	assert.Contains(t, out, "cwd:")
}

func TestWelcomeHints(t *testing.T) {
	out := welcomeHints(100)

	assert.Contains(t, out, "/help")
	assert.Contains(t, out, "ctrl+?")
	assert.Contains(t, out, "tab")
}

// The welcome screen must never paint outside the terminal, even on narrow
// terminals where the tile and the info column do not fit side by side.
func TestWelcomeScreenFitsWidth(t *testing.T) {
	for _, width := range []int{40, 60, 80, 120} {
		joined := lipgloss.JoinVertical(
			lipgloss.Top,
			bannerHeader(width),
			"",
			welcomeHints(width),
		)
		rendered := lipgloss.NewStyle().Width(width).Render(joined)

		require.LessOrEqual(t, lipgloss.Width(rendered), width,
			"rendered welcome screen is wider than the terminal (%d)", width)
	}
}
