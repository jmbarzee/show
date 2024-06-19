package node

import (
	"encoding/json"
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

// GetNodeInfo returns the NodeInfo of this Node or a child node,
// if the given ID is a match
func (n Group) GetNodeInfo(nodeID uuid.UUID) common.NodeInfo {
	if n.ID == nodeID {
		return n
	}
	for _, child := range n.Children {
		nodeInfo := child.GetNodeInfo(nodeID)
		if nodeInfo != nil {
			return nodeInfo
		}
	}
	return nil
}

// GetChildren returns all groups under the GroupOption
func (n Group) GetChildren() []common.Node {
	return n.Children
}

// GetChildren returns all groups under the GroupOption
func (n Group) GetChildrenInfo() []common.NodeInfo {
	nodes := make([]common.NodeInfo, len(n.Children))
	for i, child := range n.Children {
		nodes[i] = child
	}
	return nodes
}

// Insert will insert a node underneath a parent node.
func (n *Group) Insert(parentID uuid.UUID, newNode common.Node) error {
	if parentID == n.ID {
		n.Children = append(n.Children, newNode)
		return nil
	}
	for _, child := range n.Children {
		err := child.Insert(parentID, newNode)
		if err == nil {
			return nil
		} else if errors.Is(err, ErrorParentCantHaveChildren) {
			return err
		}
	}
	return ErrorFindParentNode
}

// Delete will delete a node underneath a parent node.
func (n *Group) Delete(parentID, childID uuid.UUID) error {
	if parentID == n.ID {
		for i, child := range n.Children {
			if child.GetID() == childID {
				length := len(n.Children)
				n.Children[i] = n.Children[length-1] // Copy last element to index i.
				n.Children[length-1] = nil           // Erase last element (write zero value).
				n.Children = n.Children[:length-1]
				return nil
			}
		}
		return ErrorFindChildNode
	}

	for _, child := range n.Children {
		err := child.Delete(parentID, childID)
		if err == nil {
			return nil
		} else if errors.Is(err, ErrorParentCantHaveChildren) {
			return err
		}
	}
	return ErrorFindParentNode
}

// GetType returns the type
func (Group) GetType() string {
	return GroupType
}

func (n *Group) MarshalJSON() ([]byte, error) {
	temp := &struct {
		ID       uuid.UUID
		Type     string
		Children []common.Node
	}{}

	temp.ID = n.ID
	temp.Type = n.GetType()
	temp.Children = n.Children

	return json.Marshal(temp)
}

func (n *Group) UnmarshalJSON(data []byte) error {
	temp := &struct {
		ID       uuid.UUID
		Children []json.RawMessage
	}{}

	err := json.Unmarshal(data, temp)
	if err != nil {
		return err
	}
	children, err := UnmarshalJSONs(temp.Children)
	if err != nil {
		return err
	}

	n.ID = temp.ID
	n.Children = children

	return nil
}
