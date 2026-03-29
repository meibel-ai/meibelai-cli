package output

import (
	"github.com/charmbracelet/lipgloss"
)

// Color palette
var (
	// Primary colors
	ColorPrimary   = lipgloss.Color("#7C3AED") // Purple
	ColorSecondary = lipgloss.Color("#06B6D4") // Cyan
	ColorSuccess   = lipgloss.Color("#10B981") // Green
	ColorWarning   = lipgloss.Color("#F59E0B") // Yellow
	ColorError     = lipgloss.Color("#EF4444") // Red
	ColorMuted     = lipgloss.Color("#6B7280") // Gray

	// Text colors
	ColorText     = lipgloss.Color("#F9FAFB") // White
	ColorSubtle   = lipgloss.Color("#9CA3AF") // Light gray
	ColorHighlight = lipgloss.Color("#FBBF24") // Gold
)

// Styles defines the visual styles for CLI output.
var Styles = struct {
	// Text styles
	Title     lipgloss.Style
	Subtitle  lipgloss.Style
	Bold      lipgloss.Style
	Muted     lipgloss.Style
	Success   lipgloss.Style
	Warning   lipgloss.Style
	Error     lipgloss.Style

	// Table styles
	TableHeader   lipgloss.Style
	TableCell     lipgloss.Style
	TableBorder   lipgloss.Style

	// Status indicators
	StatusActive  lipgloss.Style
	StatusPending lipgloss.Style
	StatusDone    lipgloss.Style

	// Misc
	Code    lipgloss.Style
	URL     lipgloss.Style
	Spinner lipgloss.Style
}{
	// Text styles
	Title: lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorText).
		MarginBottom(1),

	Subtitle: lipgloss.NewStyle().
		Foreground(ColorSubtle).
		Italic(true),

	Bold: lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorText),

	Muted: lipgloss.NewStyle().
		Foreground(ColorMuted),

	Success: lipgloss.NewStyle().
		Foreground(ColorSuccess),

	Warning: lipgloss.NewStyle().
		Foreground(ColorWarning),

	Error: lipgloss.NewStyle().
		Foreground(ColorError).
		Bold(true),

	// Table styles
	TableHeader: lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorPrimary).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(ColorMuted),

	TableCell: lipgloss.NewStyle().
		Foreground(ColorText).
		Padding(0, 1),

	TableBorder: lipgloss.NewStyle().
		Foreground(ColorMuted),

	// Status indicators
	StatusActive: lipgloss.NewStyle().
		Foreground(ColorSuccess).
		Bold(true),

	StatusPending: lipgloss.NewStyle().
		Foreground(ColorWarning),

	StatusDone: lipgloss.NewStyle().
		Foreground(ColorMuted),

	// Misc
	Code: lipgloss.NewStyle().
		Background(lipgloss.Color("#1F2937")).
		Foreground(ColorSecondary).
		Padding(0, 1),

	URL: lipgloss.NewStyle().
		Foreground(ColorSecondary).
		Underline(true),

	Spinner: lipgloss.NewStyle().
		Foreground(ColorPrimary),
}

// Icon constants
const (
	IconSuccess = "✓"
	IconError   = "✗"
	IconWarning = "⚠"
	IconInfo    = "ℹ"
	IconArrow   = "→"
	IconBullet  = "•"
)
