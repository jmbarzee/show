package node

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/repeatable"
)

const GroupOptionType = "GroupOption"

// GroupOption represents a series of groups
type GroupOption struct {
	Basic
	Groups []*Group
}

var _ common.Node = (*Group)(nil)

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
	groupNum := repeatable.Option(vibe.Start(), len(n.Groups))
	fmt.Println("Using Group:", groupNum)
	n.Groups[groupNum].Allocate(vibe)
}

// Clean removes all stored resources which ended before a time
func (n GroupOption) Clean(t time.Time) {
	for _, group := range n.Groups {
		group.Clean(t)
	}
}

// GetChildren returns all groups under the GroupOption
func (n GroupOption) GetChildren() []common.Node {
	nodes := make([]common.Node, len(n.Groups))
	for i, group := range n.Groups {
		nodes[i] = group
	}
	return nodes
}

// Insert will insert a node underneath a parent node.
func (n *GroupOption) Insert(parentID uuid.UUID, newNode common.Node) error {
	if parentID == n.id {
		group := NewGroup(newNode)
		n.Groups = append(n.Groups, group)
		return nil
	}

	for _, group := range n.Groups {
		err := group.Insert(parentID, newNode)
		if err == nil {
			return nil
		} else if errors.Is(err, ParentCantHaveChildrenError) {
			return err
		}
	}
	return FindParentNodeError
}

// Delete will delete a node underneath a parent node.
func (n *GroupOption) Delete(parentID, childID uuid.UUID) error {
	if parentID == n.id {
		for i, group := range n.Groups {
			if group.GetID() == childID {
				length := len(n.Groups)
				n.Groups[i] = n.Groups[length-1] // Copy last element to index i.
				n.Groups[length-1] = nil         // Erase last element (write zero value).
				n.Groups = n.Groups[:length-1]
			}
		}
		return FindChildNodeError
	}
	for _, group := range n.Groups {
		err := group.Delete(parentID, childID)
		if err == nil {
			return nil
		} else if errors.Is(err, ParentCantHaveChildrenError) {
			return err
		}
	}
	return FindParentNodeError
}

// GetType returns the type
func (GroupOption) GetType() string {
	return GroupOptionType
}
