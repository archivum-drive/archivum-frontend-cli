package nodeDetails

import (
	"archivum-frontend-cli/internal/models"

	"github.com/charmbracelet/lipgloss"
)

func View(node models.Node, active bool) string {
	style := lipgloss.NewStyle().
		Padding(0, 1).
		MarginLeft(1).
		MarginBottom(1)

	styleActive := style.
		UnsetMarginLeft().
		BorderLeft(true).
		BorderStyle(lipgloss.ThickBorder())

	output := node.Name + "\n"

	if active {
		return styleActive.Render(output)
	}
	return style.Render(output)
}
