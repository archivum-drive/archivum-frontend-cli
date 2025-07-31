package models

import "github.com/google/uuid"

type BlobId string

type NodeId uuid.UUID

type NodeType int

const (
	NodeTypeFile NodeType = iota
	NodeTypeGroup
)

type Node struct {
	Id         NodeId
	Name       string
	BlobRef    BlobId
	NodeType   NodeType
	Attributes map[string]string
	Members    []NodeId
}
