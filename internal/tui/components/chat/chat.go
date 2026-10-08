package chat

import (
	"fmt"
	"sort"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/heliumcode-labs/helium/internal/config"
	"github.com/heliumcode-labs/helium/internal/message"
	"github.com/heliumcode-labs/helium/internal/session"
	"github.com/heliumcode-labs/helium/internal/tui/styles"
	"github.com/heliumcode-labs/helium/internal/tui/theme"
	"github.com/heliumcode-labs/helium/internal/version"
)

type SendMsg struct {
	Text        string
	Attachments []message.Attachment
}

type SessionSelectedMsg = session.Session

type SessionClearedMsg struct{}

type EditorFocusMsg bool

// SlashCommandMsg is emitted when the user submits a line that begins with
// "/". The shell runs a built-in command instead of sending the text to the
// model, which makes the commands discoverable on touch keyboards where the
// ctrl-based shortcuts are hard to reach.
type SlashCommandMsg struct {
	Input string
}

func header(width int) string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		logo(width),
		repo(width),
		"",
		cwd(width),
	)
}

func lspsConfigured(width int) string {
	cfg := config.Get()
	title := "LSP Configuration"
	title = ansi.Truncate(title, width, "…")

	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	lsps := baseStyle.
		Width(width).
		Foreground(t.Primary()).
		Bold(true).
		Render(title)

	// Get LSP names and sort them for consistent ordering
	var lspNames []string
	for name := range cfg.LSP {
		lspNames = append(lspNames, name)
	}
	sort.Strings(lspNames)

	// Nothing configured yet: don't leave a dangling heading on the welcome
	// screen.
	if len(lspNames) == 0 {
		return ""
	}

	var lspViews []string
	for _, name := range lspNames {
		lsp := cfg.LSP[name]
		lspName := baseStyle.
			Foreground(t.Text()).
			Render(fmt.Sprintf("• %s", name))

		cmd := lsp.Command
		cmd = ansi.Truncate(cmd, width-lipgloss.Width(lspName)-3, "…")

		lspPath := baseStyle.
			Foreground(t.TextMuted()).
			Render(fmt.Sprintf(" (%s)", cmd))

		lspViews = append(lspViews,
			baseStyle.
				Width(width).
				Render(
					lipgloss.JoinHorizontal(
						lipgloss.Left,
						lspName,
						lspPath,
					),
				),
		)
	}

	return baseStyle.
		Width(width).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				lsps,
				lipgloss.JoinVertical(
					lipgloss.Left,
					lspViews...,
				),
			),
		)
}

func logo(width int) string {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	versionText := baseStyle.
		Foreground(t.TextMuted()).
		Render(version.Version)

	// Element badge + split wordmark: "Helium" in text, "Code" in accent.
	badge := baseStyle.
		Background(t.Primary()).
		Foreground(t.Background()).
		Bold(true).
		Padding(0, 1).
		Render(styles.HeliumIcon)

	name := lipgloss.JoinHorizontal(
		lipgloss.Left,
		baseStyle.Foreground(t.Text()).Bold(true).Render("Helium"),
		baseStyle.Foreground(t.Accent()).Bold(true).Render("Code"),
	)

	return baseStyle.
		Width(width).
		Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				badge,
				" ",
				name,
				"  ",
				versionText,
			),
		)
}

func repo(width int) string {
	repo := "https://github.com/heliumcode-labs/helium"
	t := theme.CurrentTheme()

	return styles.BaseStyle().
		Foreground(t.TextMuted()).
		Width(width).
		Render(repo)
}

func cwd(width int) string {
	cwd := fmt.Sprintf("cwd: %s", config.WorkingDirectory())
	t := theme.CurrentTheme()

	return styles.BaseStyle().
		Foreground(t.TextMuted()).
		Width(width).
		Render(cwd)
}
