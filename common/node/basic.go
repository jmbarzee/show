package node

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/common"
)

var ErrorFindParentNode = errors.New("Failed to find parent node with matching ID")
var ErrorFindChildNode = errors.New("Failed to find child node with matching ID")
var ErrorParentCantHaveChildren = errors.New("Found node with matching ID, but node can't hold children")

// Basic implements some traits and features which are shared between all nodes
type Basic struct {
	ID uuid.UUID
}

// NewBasic creates a Basic with a new ID
func NewBasic() Basic {
	return Basic{
		ID: uuid.New(),
	}
}

// GetChildren returns any children under the node
func (Basic) GetChildren() []common.Node {
	return nil
}

// GetChildrenInfo returns any children under the node
func (Basic) GetChildrenInfo() []common.NodeInfo {
	return nil
}

// Insert will insert a node underneath a parent node.
func (d Basic) Insert(parentID uuid.UUID, newNode common.Node) error {
	if parentID == d.ID {
		return ErrorParentCantHaveChildren
	}
	return ErrorFindParentNode
}

// Delete will delete a node underneath a parent node.
func (d Basic) Delete(parentID, childID uuid.UUID) error {
	if parentID == d.ID {
		return ErrorParentCantHaveChildren
	}
	return ErrorFindParentNode
}

// Get ID provides a wa
func (d Basic) GetID() uuid.UUID {
	return d.ID
}
