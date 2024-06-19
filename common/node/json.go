package node

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/common"
)

type ErrorNodeTypeNotRegistered struct {
	nodeType string
}

func (e ErrorNodeTypeNotRegistered) Error() string {
	return fmt.Sprintf("Failed to find registered node for UnmarshallingJSON: %s", e.nodeType)
}

type typed struct {
	Type string
}

type builder func() common.Node

var register = map[string]builder{
	GroupType:       func() common.Node { return &Group{} },
	GroupOptionType: func() common.Node { return &GroupOption{} },
}

// Register can be called to register nodes for unmarshalling.
// IMPORTANT: only register nodes which are abstract and not tied to a device
func Register(nb builder) {
	n := nb()
	register[n.GetType()] = nb
}

// Generic is a type of placeholder node, used by UnmarshalJSON
// so that an ID can be stored and later linked to a device provided node
type Generic struct {
	Basic
}

// GetType returns the type
func (_ Generic) GetType() string { return "generic" }

// Allocate passes Vibe into this device and its children
// Allocate typically Stabilize the Vibe before passing it to children devices
func (_ Generic) Allocate(common.Vibe) {}

// Clean removes all stored resources which ended before a time
func (_ Generic) Clean(time.Time) {}

// GetNodeInfo finds a given NodeInfo with a matching nodeID
// through tree traversal
func (_ Generic) GetNodeInfo(uuid.UUID) common.NodeInfo { return nil }

func (d *Generic) MarshalJSON() ([]byte, error) {
	type nodeJSON Generic
	temp := (*nodeJSON)(d)

	return json.Marshal(temp)
}

func (d *Generic) UnmarshalJSON(data []byte) error {
	type nodeJSON Generic
	temp := (*nodeJSON)(d)

	return json.Unmarshal(data, temp)
}

func UnmarshalJSON(data json.RawMessage) (common.Node, error) {
	temp := &typed{}

	err := json.Unmarshal(data, temp)
	if err != nil {
		return nil, err
	}

	n := (common.Node)(&Generic{})
	nb, ok := register[temp.Type]
	if ok {
		n = nb()
	}

	err = n.UnmarshalJSON(data)
	return n, err
}

func UnmarshalJSONs(datas []json.RawMessage) ([]common.Node, error) {
	nodes := make([]common.Node, len(datas))
	for i, data := range datas {
		node, err := UnmarshalJSON(data)
		if err != nil {
			return nil, err
		}

		nodes[i] = node
	}
	return nodes, nil
}
