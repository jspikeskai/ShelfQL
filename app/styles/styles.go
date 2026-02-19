package styles

import "github.com/charmbracelet/lipgloss"

var TitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#9E4D9A")).
	Background(lipgloss.Color("#EBEBEB")).
	Width(50).
	Align(lipgloss.Center)

var InfoStyle = lipgloss.NewStyle().
	Italic(true).
	Foreground(lipgloss.Color("#42FFFF"))
