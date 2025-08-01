package allObjectsView

import (
	"archivum-frontend-cli/internal/models"
	"archivum-frontend-cli/internal/ui/components/nodeDetails"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	nodes    map[models.NodeId]models.Node
	nodeKeys []models.NodeId

	selectedNode models.NodeId
}

func NewModel(nodes map[models.NodeId]models.Node) Model {
	filtered := make(map[models.NodeId]models.Node)
	for id, node := range nodes {
		if node.NodeType == models.NodeTypeFile {
			filtered[id] = node
		}
	}
	nodes = filtered

	keys := make([]models.NodeId, 0, len(nodes))
	for id := range nodes {
		keys = append(keys, id)
	}

	var firstNode models.NodeId
	if len(keys) > 0 {
		firstNode = keys[0]
	}
	return Model{
		nodes:        nodes,
		nodeKeys:     keys,
		selectedNode: firstNode,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.handleKeyPress(msg)
		return m, nil
	}
	return m, nil
}

func (m *Model) handleKeyPress(msg tea.KeyMsg) {
	switch msg.String() {
	case "j", "down":
		m.selectNextNode()
	case "k", "up":
		m.selectPreviousNode()
	}
}

func (m *Model) selectNextNode() {
	var nextNode models.NodeId
	for i, id := range m.nodeKeys {
		if id == m.selectedNode {
			if i+1 < len(m.nodeKeys) {
				nextNode = m.nodeKeys[i+1]
				m.selectedNode = nextNode
			}
			break
		}
	}
}

func (m *Model) selectPreviousNode() {
	var prevNode models.NodeId
	for i, id := range m.nodeKeys {
		if id == m.selectedNode {
			if i > 0 {
				prevNode = m.nodeKeys[i-1]
				m.selectedNode = prevNode
			}
			break
		}
	}
}

func View(m Model) string {
	if len(m.nodes) == 0 {
		return "No objects found."
	}

	var output []string

	for _, nodeId := range m.nodeKeys {
		node := m.nodes[nodeId]
		if node.NodeType == models.NodeTypeFile {
			output = append(output, nodeDetails.View(node, node.Id == m.selectedNode))
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, output...)
}
