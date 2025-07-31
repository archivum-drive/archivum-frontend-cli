package data

import (
	"archivum-frontend-cli/internal/models"

	"github.com/google/uuid"
)

func MockNode() models.Node {
	return models.Node{
		Id:         models.NodeId(uuid.New()),
		Name:       "Example File",
		BlobRef:    models.BlobId("example-blob-id"),
		NodeType:   models.NodeType(0),
		Attributes: nil,
		Members:    nil,
	}
}

func MockGroupNode(members []models.NodeId) models.Node {
	return models.Node{
		Id:         models.NodeId(uuid.New()),
		Name:       "Example Group",
		BlobRef:    models.BlobId(""),
		NodeType:   models.NodeType(1),
		Attributes: nil,
		Members:    members,
	}
}

func MockFileSystem() map[models.NodeId]models.Node {
	nodes := make(map[models.NodeId]models.Node)

	// Create some mock nodes
	node := MockNode()
	nodes[node.Id] = node
	node = MockNode()
	node.Name = "Another File"
	nodes[node.Id] = node

	node = MockGroupNode([]models.NodeId{nodes[node.Id].Id})
	nodes[node.Id] = node

	return nodes
}
