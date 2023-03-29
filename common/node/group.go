package node

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/common"
)

const GroupType = "Group"

// Group represents a group of devices who's effects will share traits
type Group struct {
	Basic
	Children []common.Node
}

var _ common.Node = (*Group)(nil)

// NewGroup creates a new Group with a unique ID
func NewGroup(deviceNodes ...common.Node) *Group {
	if deviceNodes == nil {
		deviceNodes = []common.Node{}
	}
	return &Group{
		Basic:    NewBasic(),
		Children: deviceNodes,
	}
}

// Allocate passes Vibe into this device and its children
// Allocate Stabilize the Vibe before passing it to children devices
func (n Group) Allocate(vibe common.Vibe) {
	newVibe := vibe.Stabilize()

	for _, child := range n.Children {
		child.Allocate(newVibe)
		newVibe = newVibe.Duplicate()
	}
}

// Clean removes all stored resources which ended before a time
func (n Group) Clean(t time.Time) {
	for _, child := range n.Children {
		child.Clean(t)
	}
}

// GetChildren returns all groups under the GroupOption
func (n Group) GetChildren() []common.Node {
	return n.Children
}

// Insert will insert a node underneath a parent node.
func (n *Group) Insert(parentID uuid.UUID, newNode common.Node) error {
	if parentID == n.id {
		n.Children = append(n.Children, newNode)
		return nil
	}
	for _, child := range n.Children {
		err := child.Insert(parentID, newNode)
		if err == nil {
			return nil
		} else if errors.Is(err, ParentCantHaveChildrenError) {
			return err
		}
	}
	return FindParentNodeError
}

// Delete will delete a node underneath a parent node.
func (n *Group) Delete(parentID, childID uuid.UUID) error {
	if parentID == n.id {
		for i, child := range n.Children {
			if child.GetID() == childID {
				length := len(n.Children)
				n.Children[i] = n.Children[length-1] // Copy last element to index i.
				n.Children[length-1] = nil           // Erase last element (write zero value).
				n.Children = n.Children[:length-1]
			}
		}
		return FindChildNodeError
	}
	for _, child := range n.Children {
		err := child.Delete(parentID, childID)
		if err == nil {
			return nil
		} else if errors.Is(err, ParentCantHaveChildrenError) {
			return err
		}
	}
	return FindParentNodeError
}

// GetType returns the type
func (Group) GetType() string {
	return GroupType
}
