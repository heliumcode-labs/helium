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

// Brand mark of HeliumCode: a compact wordmark. The periodic-table tile that
// used to sit here was a bordered box that overflowed narrow terminals, so it
// is gone.
const tagline = "agentic coding, right in your terminal"

// bannerHeader is the welcome variant of the header. It renders a compact
// wordmark plus the essentials and deliberately avoids the boxed element tile,
// which overflowed on narrow (phone) terminals.
func bannerHeader(width int) string {
	t := theme.CurrentTheme()
	base := styles.Regular()

	if width < 12 {
		width = 12
	}

	muted := func(s string) string {
		return base.Foreground(t.TextMuted()).Render(ansi.Truncate(s, width, "…"))
	}

	name := lipgloss.JoinHorizontal(
		lipgloss.Left,
		base.Foreground(t.Primary()).Bold(true).Render("Helium"),
		base.Foreground(t.Accent()).Bold(true).Render("Code"),
	)

	// Drop the version rather than let the wordmark wrap on very narrow screens.
	if lipgloss.Width(name)+1+len(version.Version) <= width {
		name = lipgloss.JoinHorizontal(
			lipgloss.Left,
			name,
			" ",
			base.Foreground(t.TextMuted()).Render(version.Version),
		)
	}

	repo := "https://github.com/heliumcode-labs/helium"

	lines := []string{name, muted(tagline)}

	// The repo link is the first thing to go when space is tight.
	if lipgloss.Width(repo) <= width {
		lines = append(lines, muted(repo))
	}

	lines = append(lines, muted(fmt.Sprintf("cwd: %s", config.WorkingDirectory())))

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
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
		key("/help"), hint(" shortcuts    "),
		key("tab"), hint(" autocomplete"),
	)

	// One muted line naming the slash commands, so they're discoverable
	// without needing a ctrl-based shortcut.
	commands := base.Foreground(t.TextMuted()).Render(
		ansi.Truncate("/help /commands /models /theme /sessions /quit", width, "…"),
	)

	lines := []string{intro}
	if !providerConfigured() {
		lines = append(lines, "",
			base.Foreground(t.Warning()).Render(
				"No provider configured — add an API key to a provider to start chatting.",
			),
		)
	}
	lines = append(lines, "", commands, "", shortcuts)

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
