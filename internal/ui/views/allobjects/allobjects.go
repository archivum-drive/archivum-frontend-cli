package allObjectsView

import "archivum-frontend-cli/internal/models"

type Model struct {
	nodes map[models.NodeId]models.Node
}

func NewModel(nodes map[models.NodeId]models.Node) Model {
	return Model{
		nodes: nodes,
	}
}

func View(m Model) string {
	if len(m.nodes) == 0 {
		return "No objects found."
	}

	var output string
	for _, node := range m.nodes {
		if node.NodeType == models.NodeTypeGroup {
			continue
		}
		output += node.Name + "\n"
	}

	return output
}
