package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// HeliumTheme implements the Theme interface with the HeliumCode brand colors.
// It provides both dark and light variants.
//
// The palette is built around a deep ink background with a neon-mint primary
// (the glow of an excited noble gas), a violet secondary and a coral accent.
type HeliumTheme struct {
	BaseTheme
}

// NewHeliumTheme creates a new instance of the HeliumCode theme.
func NewHeliumTheme() *HeliumTheme {
	// Dark mode colors
	darkBackground := "#0b0f17"
	darkCurrentLine := "#131a26"
	darkSelection := "#1e2637"
	darkForeground := "#e6edf3"
	darkComment := "#7d8590"
	darkPrimary := "#4fe3c1"   // neon mint
	darkSecondary := "#8b7bf7" // violet
	darkAccent := "#ff7ab6"    // coral pink
	darkRed := "#ff6b6b"       // Error red
	darkOrange := "#ffb020"    // Warning amber
	darkGreen := "#56d364"     // Success green
	darkCyan := "#58a6ff"      // Info blue
	darkYellow := "#ffd666"    // Emphasized text
	darkBorder := "#263145"    // Border color

	// Light mode colors
	lightBackground := "#f7f9fc"
	lightCurrentLine := "#eef1f6"
	lightSelection := "#e2e8f0"
	lightForeground := "#10161f"
	lightComment := "#5b6472"
	lightPrimary := "#0b8f79"   // deep mint
	lightSecondary := "#6d4fe0" // violet
	lightAccent := "#db2777"    // coral pink
	lightRed := "#d32f2f"       // Error red
	lightOrange := "#c47200"    // Warning amber
	lightGreen := "#1a7f37"     // Success green
	lightCyan := "#1668c7"      // Info blue
	lightYellow := "#9a6700"    // Emphasized text
	lightBorder := "#d5dce6"    // Border color

	theme := &HeliumTheme{}

	// Base colors
	theme.PrimaryColor = lipgloss.AdaptiveColor{
		Dark:  darkPrimary,
		Light: lightPrimary,
	}
	theme.SecondaryColor = lipgloss.AdaptiveColor{
		Dark:  darkSecondary,
		Light: lightSecondary,
	}
	theme.AccentColor = lipgloss.AdaptiveColor{
		Dark:  darkAccent,
		Light: lightAccent,
	}

	// Status colors
	theme.ErrorColor = lipgloss.AdaptiveColor{
		Dark:  darkRed,
		Light: lightRed,
	}
	theme.WarningColor = lipgloss.AdaptiveColor{
		Dark:  darkOrange,
		Light: lightOrange,
	}
	theme.SuccessColor = lipgloss.AdaptiveColor{
		Dark:  darkGreen,
		Light: lightGreen,
	}
	theme.InfoColor = lipgloss.AdaptiveColor{
		Dark:  darkCyan,
		Light: lightCyan,
	}

	// Text colors
	theme.TextColor = lipgloss.AdaptiveColor{
		Dark:  darkForeground,
		Light: lightForeground,
	}
	theme.TextMutedColor = lipgloss.AdaptiveColor{
		Dark:  darkComment,
		Light: lightComment,
	}
	theme.TextEmphasizedColor = lipgloss.AdaptiveColor{
		Dark:  darkYellow,
		Light: lightYellow,
	}

	// Background colors
	theme.BackgroundColor = lipgloss.AdaptiveColor{
		Dark:  darkBackground,
		Light: lightBackground,
	}
	theme.BackgroundSecondaryColor = lipgloss.AdaptiveColor{
		Dark:  darkCurrentLine,
		Light: lightCurrentLine,
	}
	theme.BackgroundDarkerColor = lipgloss.AdaptiveColor{
		Dark:  "#070a10", // deeper than the background
		Light: "#ffffff",
	}

	// Border colors
	theme.BorderNormalColor = lipgloss.AdaptiveColor{
		Dark:  darkBorder,
		Light: lightBorder,
	}
	theme.BorderFocusedColor = lipgloss.AdaptiveColor{
		Dark:  darkPrimary,
		Light: lightPrimary,
	}
	theme.BorderDimColor = lipgloss.AdaptiveColor{
		Dark:  darkSelection,
		Light: lightSelection,
	}

	// Diff view colors
	theme.DiffAddedColor = lipgloss.AdaptiveColor{
		Dark:  "#7ee787",
		Light: "#1a7f37",
	}
	theme.DiffRemovedColor = lipgloss.AdaptiveColor{
		Dark:  "#ff7b72",
		Light: "#cf222e",
	}
	theme.DiffContextColor = lipgloss.AdaptiveColor{
		Dark:  "#8b949e",
		Light: "#57606a",
	}
	theme.DiffHunkHeaderColor = lipgloss.AdaptiveColor{
		Dark:  "#79c0ff",
		Light: "#0969da",
	}
	theme.DiffHighlightAddedColor = lipgloss.AdaptiveColor{
		Dark:  "#b7f0b7",
		Light: "#1b5e20",
	}
	theme.DiffHighlightRemovedColor = lipgloss.AdaptiveColor{
		Dark:  "#ffc9c9",
		Light: "#7f1d1d",
	}
	theme.DiffAddedBgColor = lipgloss.AdaptiveColor{
		Dark:  "#12281a",
		Light: "#e6f4ea",
	}
	theme.DiffRemovedBgColor = lipgloss.AdaptiveColor{
		Dark:  "#2a1517",
		Light: "#fbe9e9",
	}
	theme.DiffContextBgColor = lipgloss.AdaptiveColor{
		Dark:  darkBackground,
		Light: lightBackground,
	}
	theme.DiffLineNumberColor = lipgloss.AdaptiveColor{
		Dark:  "#6e7681",
		Light: "#8c959f",
	}
	theme.DiffAddedLineNumberBgColor = lipgloss.AdaptiveColor{
		Dark:  "#16301c",
		Light: "#ccf0d3",
	}
	theme.DiffRemovedLineNumberBgColor = lipgloss.AdaptiveColor{
		Dark:  "#301b1c",
		Light: "#ffd9dc",
	}

	// Markdown colors
	theme.MarkdownTextColor = lipgloss.AdaptiveColor{
		Dark:  darkForeground,
		Light: lightForeground,
	}
	theme.MarkdownHeadingColor = lipgloss.AdaptiveColor{
		Dark:  darkSecondary,
		Light: lightSecondary,
	}
	theme.MarkdownLinkColor = lipgloss.AdaptiveColor{
		Dark:  darkPrimary,
		Light: lightPrimary,
	}
	theme.MarkdownLinkTextColor = lipgloss.AdaptiveColor{
		Dark:  darkCyan,
		Light: lightCyan,
	}
	theme.MarkdownCodeColor = lipgloss.AdaptiveColor{
		Dark:  darkGreen,
		Light: lightGreen,
	}
	theme.MarkdownBlockQuoteColor = lipgloss.AdaptiveColor{
		Dark:  darkComment,
		Light: lightComment,
	}
	theme.MarkdownEmphColor = lipgloss.AdaptiveColor{
		Dark:  darkYellow,
		Light: lightYellow,
	}
	theme.MarkdownStrongColor = lipgloss.AdaptiveColor{
		Dark:  darkAccent,
		Light: lightAccent,
	}
	theme.MarkdownHorizontalRuleColor = lipgloss.AdaptiveColor{
		Dark:  darkComment,
		Light: lightComment,
	}
	theme.MarkdownListItemColor = lipgloss.AdaptiveColor{
		Dark:  darkPrimary,
		Light: lightPrimary,
	}
	theme.MarkdownListEnumerationColor = lipgloss.AdaptiveColor{
		Dark:  darkSecondary,
		Light: lightSecondary,
	}
	theme.MarkdownImageColor = lipgloss.AdaptiveColor{
		Dark:  darkAccent,
		Light: lightAccent,
	}
	theme.MarkdownImageTextColor = lipgloss.AdaptiveColor{
		Dark:  darkCyan,
		Light: lightCyan,
	}
	theme.MarkdownCodeBlockColor = lipgloss.AdaptiveColor{
		Dark:  darkForeground,
		Light: lightForeground,
	}

	// Syntax highlighting colors
	theme.SyntaxCommentColor = lipgloss.AdaptiveColor{
		Dark:  darkComment,
		Light: lightComment,
	}
	theme.SyntaxKeywordColor = lipgloss.AdaptiveColor{
		Dark:  darkAccent,
		Light: lightAccent,
	}
	theme.SyntaxFunctionColor = lipgloss.AdaptiveColor{
		Dark:  darkPrimary,
		Light: lightPrimary,
	}
	theme.SyntaxVariableColor = lipgloss.AdaptiveColor{
		Dark:  darkForeground,
		Light: lightForeground,
	}
	theme.SyntaxStringColor = lipgloss.AdaptiveColor{
		Dark:  darkGreen,
		Light: lightGreen,
	}
	theme.SyntaxNumberColor = lipgloss.AdaptiveColor{
		Dark:  darkOrange,
		Light: lightOrange,
	}
	theme.SyntaxTypeColor = lipgloss.AdaptiveColor{
		Dark:  darkSecondary,
		Light: lightSecondary,
	}
	theme.SyntaxOperatorColor = lipgloss.AdaptiveColor{
		Dark:  darkCyan,
		Light: lightCyan,
	}
	theme.SyntaxPunctuationColor = lipgloss.AdaptiveColor{
		Dark:  darkComment,
		Light: lightComment,
	}

	return theme
}

func init() {
	// Register the HeliumCode theme with the theme manager
	RegisterTheme("helium", NewHeliumTheme())
}
