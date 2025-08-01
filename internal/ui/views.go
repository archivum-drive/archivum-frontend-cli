package ui

import (
	"strings"

	"archivum-frontend-cli/internal/data"
	allObjectsView "archivum-frontend-cli/internal/ui/views/allobjects"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

type View struct {
	name string
	id   int
}

var views = []View{
	{name: "All Objects", id: 0},
	{name: "Groups", id: 1},
}

var (
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       " ",
		TopLeft:     "╭",
		TopRight:    "┘",
		BottomLeft:  "╰",
		BottomRight: "┐",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "┤",
		BottomLeft:  "╰",
		BottomRight: "┤",
	}

	containerStyle = lipgloss.NewStyle().
			Padding(1, 2)
)

type Model struct {
	currentView    View
	viewportHeight int

	allObjectsView allObjectsView.Model
}

func NewModel() Model {
	return Model{
		currentView: views[0],

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
	case tea.WindowSizeMsg:
		m.viewportHeight = msg.Height
	}
	return m, nil
}

func (m Model) View() string {
	view := ""

	switch m.currentView.id {
	case 0:
		view = allObjectsView.View(m.allObjectsView)
	default:
		view = "Unknown view"
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.renderMenue(),
		containerStyle.Render(view),
	)
}

func (m Model) handleKeyPress(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "tab":
		if m.currentView.id >= (len(views) - 1) {
			m.currentView = views[0]
		} else {
			m.currentView = views[m.currentView.id+1]
		}
		return m, nil
	}

	return m, nil
}

func (m Model) renderMenue() string {
	menueStyle := lipgloss.NewStyle().BorderRight(true).BorderStyle(tabBorder)

	tabs := list.New()
	tabs.Enumerator(func(l list.Items, i int) string { return "" })

	for view := range views {
		tabs.Item(
			renderTab(views[view].name, views[view] == m.currentView))
	}
	lines := []string{menueStyle.Render("\nArchivum CLI  \n"), tabs.String()}

	filler := []string{}
	numberOfFillers := m.viewportHeight - lipgloss.Height(strings.Join(lines, "\n"))
	for i := 0; i < numberOfFillers; i++ {
		filler = append(filler, menueStyle.Render(""))
	}

	lines = append(lines, filler...)

	menue := lipgloss.JoinVertical(lipgloss.Right, lines...)

	return menue
}

func renderTab(name string, active bool) string {
	style := lipgloss.NewStyle().
		Border(tabBorder).
		Faint(true).
		Padding(0, 1).
		Width(16).
		Align(lipgloss.Right)

	activeStyle := style.
		Border(activeTabBorder).
		Faint(false)

	if active {
		return activeStyle.Render(name)
	}
	return style.Render(name)
}
