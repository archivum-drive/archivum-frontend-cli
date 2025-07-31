package ui

import (
	"archivum-frontend-cli/internal/data"
	allObjectsView "archivum-frontend-cli/internal/ui/views/allobjects"

	tea "github.com/charmbracelet/bubbletea"
)

type ViewType int

const (
	AllObjects ViewType = iota
)

type Model struct {
	currentView ViewType

	allObjectsView allObjectsView.Model
}

func NewModel() Model {
	return Model{
		currentView: AllObjects,

		allObjectsView: allObjectsView.NewModel(data.MockFileSystem()),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyPress(msg)
	}

	return m, nil
}

func (m Model) View() string {
	switch m.currentView {
	case AllObjects:
		return allObjectsView.View(m.allObjectsView)
	default:
		return "Unknown view"
	}
}

func (m Model) handleKeyPress(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit

	case "up", "k":
		// return m.moveUp(), nil

	case "down", "j":
		// return m.moveDown(), nil

	case "enter":
		// return m.selectItem(), nil
	}

	return m, nil
}
