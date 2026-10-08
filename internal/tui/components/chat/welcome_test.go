package chat

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBannerHeader(t *testing.T) {
	out := bannerHeader(100)

	assert.Contains(t, out, "HeliumCode")
	assert.Contains(t, out, "https://github.com/heliumcode-labs/helium")
	assert.Contains(t, out, "cwd:")
}

// The repo link is dropped rather than wrapped when the terminal is too narrow
// for it.
func TestBannerHeaderDropsRepoOnNarrowTerminal(t *testing.T) {
	out := bannerHeader(40)

	assert.Contains(t, out, "HeliumCode")
	assert.NotContains(t, out, "https://github.com/")
}

func TestWelcomeHints(t *testing.T) {
	out := welcomeHints(100)

	assert.Contains(t, out, "/help")
	assert.Contains(t, out, "/commands")
	assert.Contains(t, out, "tab")
}

// The welcome screen must never paint outside the terminal, even on a
// phone-sized one.
func TestWelcomeScreenFitsWidth(t *testing.T) {
	for _, width := range []int{36, 40, 60, 80, 120} {
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
