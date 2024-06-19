package node

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/common"
)

const GroupOptionType = "GroupOption"

// GroupOption represents a series of groups
type GroupOption struct {
	Basic
	Groups []*Group
}

var _ common.Node = (*GroupOption)(nil)

// NewGroupOption creates a new GroupOption with a unique ID
func NewGroupOption(groups ...*Group) *GroupOption {
	if groups == nil {
		groups = []*Group{}
	}
	return &GroupOption{
		Basic:  NewBasic(),
		Groups: groups,
	}
}

// Allocate passes Vibe into this device and a single child group
// Allocate Stabilize the Vibe before passing it to a child group
func (n GroupOption) Allocate(vibe common.Vibe) {
	if len(n.Groups) == 0 {
		return
	}
	groupNum := vibe.Option(len(n.Groups))
	n.Groups[groupNum].Allocate(vibe)
}

// Clean removes all stored resources which ended before a time
func (n GroupOption) Clean(t time.Time) {
	for _, group := range n.Groups {
		group.Clean(t)
	}
}

// GetNodeInfo returns the NodeInfo of this Node or a child node,
// if the given ID is a match
func (n GroupOption) GetNodeInfo(nodeID uuid.UUID) common.NodeInfo {
	if n.ID == nodeID {
		return n
	}
	for _, child := range n.Groups {
		nodeInfo := child.GetNodeInfo(nodeID)
		if nodeInfo != nil {
			return nodeInfo
		}
	}
	return nil
}

// GetChildren returns all groups under the GroupOption
func (n GroupOption) GetChildren() []common.Node {
	nodes := make([]common.Node, len(n.Groups))
	for i, group := range n.Groups {
		nodes[i] = group
	}
	return nodes
}

// GetChildrenInfo returns all groups under the GroupOption
func (n GroupOption) GetChildrenInfo() []common.NodeInfo {
	nodes := make([]common.NodeInfo, len(n.Groups))
	for i, group := range n.Groups {
		nodes[i] = group
	}
	return nodes
}

// Insert will insert a node underneath a parent node.
func (n *GroupOption) Insert(parentID uuid.UUID, newNode common.Node) error {
	if parentID == n.ID {
		group := NewGroup(newNode)
		n.Groups = append(n.Groups, group)
		return nil
	}

	for _, group := range n.Groups {
		err := group.Insert(parentID, newNode)
		if err == nil {
			return nil
		} else if errors.Is(err, ErrorParentCantHaveChildren) {
			return err
		}
	}
	return ErrorFindParentNode
}

// Delete will delete a node underneath a parent node.
func (n *GroupOption) Delete(parentID, childID uuid.UUID) error {
	if parentID == n.ID {
		for i, group := range n.Groups {
			if group.GetID() == childID {
				length := len(n.Groups)
				n.Groups[i] = n.Groups[length-1] // Copy last element to index i.
				n.Groups[length-1] = nil         // Erase last element (write zero value).
				n.Groups = n.Groups[:length-1]
			}
		}
		return ErrorFindChildNode
	}
	for _, group := range n.Groups {
		err := group.Delete(parentID, childID)
		if err == nil {
			return nil
		} else if errors.Is(err, ErrorParentCantHaveChildren) {
			return err
		}
	}
	return ErrorFindParentNode
}

// GetType returns the type
func (GroupOption) GetType() string {
	return GroupOptionType
}

type groupOptionJSON struct {
	ID     uuid.UUID
	Type   string
	Groups []*Group
}

func (n *GroupOption) MarshalJSON() ([]byte, error) {
	temp := &groupOptionJSON{}

	temp.ID = n.ID
	temp.Type = n.GetType()
	temp.Groups = n.Groups

	return json.Marshal(temp)
}

func (n *GroupOption) UnmarshalJSON(data []byte) error {
	temp := &groupOptionJSON{}

	if err := json.Unmarshal(data, temp); err != nil {
		return err
	}

	n.ID = temp.ID
	n.Groups = temp.Groups

	return nil
}
