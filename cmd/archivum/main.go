package main

import (
	"fmt"
	"os"

	"archivum-frontend-cli/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Create the application model
	model := ui.NewModel()

	// Create a new Bubble Tea program with the model
	program := tea.NewProgram(model, tea.WithAltScreen())

	// Run the program
	if _, err := program.Run(); err != nil {
		fmt.Printf("Error running application: %v\n", err)
		os.Exit(1)
	}
}
