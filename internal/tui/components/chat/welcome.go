package chat

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/heliumcode-labs/helium/internal/config"
	"github.com/heliumcode-labs/helium/internal/tui/styles"
	"github.com/heliumcode-labs/helium/internal/tui/theme"
	"github.com/heliumcode-labs/helium/internal/version"
)

// Brand mark of HeliumCode: a periodic-table style tile carrying helium's
// atomic number, symbol and atomic weight.
const (
	elementNumber = "2"
	elementWeight = "4.0026"
	elementName   = "helium"
	elementTileW  = 14
)

const tagline = "agentic coding, right in your terminal"

// elementTile renders the bordered brand tile.
func elementTile() string {
	t := theme.CurrentTheme()
	base := styles.Regular()

	// The header line has to be exactly as wide as the symbol block below it
	// so the tile edges stay flush.
	header := base.
		Width(elementTileW).
		Foreground(t.TextMuted()).
		Render(fmt.Sprintf("%s      %s", elementNumber, elementWeight))

	symbol := base.
		Width(elementTileW).
		Align(lipgloss.Center).
		Bold(true).
		Foreground(t.Primary()).
		Render(styles.HeliumIcon)

	name := base.
		Width(elementTileW).
		Align(lipgloss.Center).
		Foreground(t.TextMuted()).
		Render(elementName)

	inner := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		base.Render(""),
		symbol,
		name,
	)

	return base.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(t.Primary()).
		Padding(0, 1).
		Render(inner)
}

// bannerHeader is the welcome variant of the header: the brand tile next to
// the product name, version, repository and working directory. It is only
// used on the initial (empty session) screen.
func bannerHeader(width int) string {
	t := theme.CurrentTheme()
	base := styles.Regular()

	tile := elementTile()

	versionText := base.
		Foreground(t.TextMuted()).
		Render(version.Version)

	name := lipgloss.JoinHorizontal(
		lipgloss.Left,
		base.Foreground(t.Text()).Bold(true).Render("Helium"),
		base.Foreground(t.Accent()).Bold(true).Render("Code"),
		" ",
		versionText,
	)

	// The tile and its gap take up space on the left, the info column gets
	// whatever is left.
	rightWidth := max(10, width-lipgloss.Width(tile)-2)
	truncate := func(s string) string { return ansi.Truncate(s, rightWidth, "…") }

	info := lipgloss.JoinVertical(
		lipgloss.Left,
		truncate(name),
		base.Foreground(t.TextMuted()).Render(truncate(tagline)),
		base.Foreground(t.TextMuted()).Render(truncate("https://github.com/heliumcode-labs/helium")),
		base.Foreground(t.TextMuted()).Render(truncate(fmt.Sprintf("cwd: %s", config.WorkingDirectory()))),
	)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		tile,
		" ",
		info,
	)
}

// providerConfigured reports whether at least one provider has a usable API key.
// When nothing is configured the welcome screen nudges the user to set one up
// instead of letting the first request fail with an opaque auth error.
func providerConfigured() bool {
	cfg := config.Get()
	if cfg == nil {
		return true // don't nag when config isn't loaded (tests, early startup)
	}
	for _, provider := range cfg.Providers {
		if provider.APIKey != "" && !provider.Disabled {
			return true
		}
	}
	return false
}

// welcomeHints renders the shortcut hints shown below the banner on the
// initial screen.
func welcomeHints(width int) string {
	t := theme.CurrentTheme()
	base := styles.Regular()

	key := func(label string) string {
		return base.
			Background(t.Primary()).
			Foreground(t.Background()).
			Bold(true).
			Padding(0, 1).
			Render(label)
	}
	hint := func(label string) string {
		return base.Foreground(t.TextMuted()).Render(label)
	}

	intro := base.Foreground(t.Text()).Render(
		"Describe a task — HeliumCode plans the work, edits files and runs commands.",
	)

	shortcuts := lipgloss.JoinHorizontal(
		lipgloss.Left,
		key("/help"), hint(" commands    "),
		key("ctrl+?"), hint(" shortcuts    "),
		key("tab"), hint(" autocomplete"),
	)

	lines := []string{intro}
	if !providerConfigured() {
		lines = append(lines, "",
			base.Foreground(t.Warning()).Render(
				"No provider configured — add an API key to a provider to start chatting.",
			),
		)
	}
	lines = append(lines, "", shortcuts)

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
