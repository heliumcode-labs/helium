package dialog

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	utilComponents "github.com/heliumcode-labs/helium/internal/tui/components/util"
	"github.com/heliumcode-labs/helium/internal/tui/layout"
	"github.com/heliumcode-labs/helium/internal/tui/styles"
	"github.com/heliumcode-labs/helium/internal/tui/theme"
	"github.com/heliumcode-labs/helium/internal/tui/util"
)

// Command represents a command that can be executed
type Command struct {
	ID          string
	Title       string
	Description string
	Handler     func(cmd Command) tea.Cmd
}

func (ci Command) Render(selected bool, width int) string {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	descStyle := baseStyle.Width(width).Foreground(t.TextMuted())
	itemStyle := baseStyle.Width(width).
		Foreground(t.Text()).
		Background(t.Background())

	if selected {
		itemStyle = itemStyle.
			Background(t.Primary()).
			Foreground(t.Background()).
			Bold(true)
		descStyle = descStyle.
			Background(t.Primary()).
			Foreground(t.Background())
	}

	title := itemStyle.Padding(0, 1).Render(ci.Title)
	if ci.Description != "" {
		description := descStyle.Padding(0, 1).Render(ci.Description)
		return lipgloss.JoinVertical(lipgloss.Left, title, description)
	}
	return title
}

// CommandSelectedMsg is sent when a command is selected
type CommandSelectedMsg struct {
	Command Command
}

// CloseCommandDialogMsg is sent when the command dialog is closed
type CloseCommandDialogMsg struct{}

// CommandDialog interface for the command selection dialog
type CommandDialog interface {
	tea.Model
	layout.Bindings
	SetCommands(commands []Command)
}

type commandDialogCmp struct {
	listView utilComponents.SimpleList[Command]
	width    int
	height   int
}

type commandKeyMap struct {
	Enter  key.Binding
	Escape key.Binding
}

var commandKeys = commandKeyMap{
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select command"),
	),
	Escape: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "close"),
	),
}

func (c *commandDialogCmp) Init() tea.Cmd {
	return c.listView.Init()
}

func (c *commandDialogCmp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, commandKeys.Enter):
			selectedItem, idx := c.listView.GetSelectedItem()
			if idx != -1 {
				return c, util.CmdHandler(CommandSelectedMsg{
					Command: selectedItem,
				})
			}
		case key.Matches(msg, commandKeys.Escape):
			return c, util.CmdHandler(CloseCommandDialogMsg{})
		}
	case tea.WindowSizeMsg:
		c.width = msg.Width
		c.height = msg.Height
	}

	u, cmd := c.listView.Update(msg)
	c.listView = u.(utilComponents.SimpleList[Command])
	cmds = append(cmds, cmd)

	return c, tea.Batch(cmds...)
}

func (c *commandDialogCmp) View() string {
	t := theme.CurrentTheme()
	baseStyle := styles.BaseStyle()

	maxWidth := 40

	commands := c.listView.GetItems()

	for _, cmd := range commands {
		if len(cmd.Title) > maxWidth-4 {
			maxWidth = len(cmd.Title) + 4
		}
		if cmd.Description != "" {
			if len(cmd.Description) > maxWidth-4 {
				maxWidth = len(cmd.Description) + 4
			}
		}
	}

	// Never render wider than the terminal. Padding (4) and the border (2) are
	// added on top of the content, so the content itself has to stop at
	// width-6 — otherwise the long descriptions spill past the right edge on a
	// phone and the list can no longer be scrolled.
	if c.width > 0 {
		capW := c.width - 6
		if capW < 10 {
			capW = 10
		}
		if maxWidth > capW {
			maxWidth = capW
		}
	}
	if maxWidth < 24 && (c.width <= 0 || 24 <= c.width-6) {
		maxWidth = 24
	}

	// Budget for the list itself: total height is the list plus the border (2),
	// the padding (2), the title and two spacer lines.
	budget := 0
	if c.height > 0 {
		budget = c.height - 7
		if budget < 1 {
			budget = 1
		}

		// Descriptions wrap, so a row is not a fixed height — shrink the
		// visible window until the rendered list actually fits.
		visible := len(commands)
		if visible > budget {
			visible = budget
		}
		if visible < 1 {
			visible = 1
		}
		c.listView.SetMaxVisibleItems(visible)
		for visible > 1 && lipgloss.Height(c.listView.View()) > budget {
			visible--
			c.listView.SetMaxVisibleItems(visible)
		}
	}

	c.listView.SetMaxWidth(maxWidth)

	title := baseStyle.
		Foreground(t.Primary()).
		Bold(true).
		Width(maxWidth).
		Padding(0, 1).
		Render("Commands")

	listView := c.listView.View()
	if budget > 0 {
		// Last-resort clip so one unusually tall entry can't push the dialog
		// off the bottom of the screen.
		if rows := strings.Split(listView, "\n"); len(rows) > budget {
			listView = strings.Join(rows[:budget], "\n")
		}
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		baseStyle.Width(maxWidth).Render(""),
		baseStyle.Width(maxWidth).Render(listView),
		baseStyle.Width(maxWidth).Render(""),
	)

	return baseStyle.Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderBackground(t.Background()).
		BorderForeground(t.TextMuted()).
		Width(lipgloss.Width(content) + 4).
		Render(content)
}

func (c *commandDialogCmp) BindingKeys() []key.Binding {
	return layout.KeyMapToSlice(commandKeys)
}

func (c *commandDialogCmp) SetCommands(commands []Command) {
	c.listView.SetItems(commands)
}

// NewCommandDialogCmp creates a new command selection dialog
func NewCommandDialogCmp() CommandDialog {
	listView := utilComponents.NewSimpleList[Command](
		[]Command{},
		10,
		"No commands available",
		true,
	)
	return &commandDialogCmp{
		listView: listView,
	}
}
