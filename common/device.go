package common

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// DeviceInfo is the read-only portion of a Device
type DeviceInfo interface {
	Moved

	// GetType returns the type
	GetType() string
	// GetID will return the ID of a device node.
	GetID() uuid.UUID

	// GetNodeInfos returns all the Nodes which the device holds
	GetNodeInfos() []NodeInfo
}

// Device represents a physical device with lights
// A device is made up of at least a single Node
type Device interface {
	Moveable
	DeviceInfo

	// GetNodes returns all the Nodes which the device holds
	GetNodes() []Node
	// DispatchRender produces and dispatches Instructions
	DispatchRender(time.Time) error

	// JSON for persistance
	json.Marshaler
	json.Unmarshaler
}

// NodeInfo is the read-only portion of a Node
type NodeInfo interface {
	// GetType returns the type
	GetType() string
	// GetID will return the ID of a device node.
	GetID() uuid.UUID

	// GetChildrenInfo returns any children under the node
	GetChildrenInfo() []NodeInfo

	// GetNodeInfo finds a given NodeInfo with a matching nodeID
	// through tree traversal
	GetNodeInfo(nodeID uuid.UUID) NodeInfo
}

// A Node is a node in the tree
// Nodes can reference an object which is also a Device or part of a Device
// Node can also be an abstraction which has a Device as a parent or child
type Node interface {
	NodeInfo

	// Allocate passes Vibe into this device and its children
	// Allocate typically Stabilize the Vibe before passing it to children devices
	Allocate(Vibe)
	// Clean removes all stored resources which ended before a time
	Clean(time.Time)

	// GetChildren returns any children under the node
	GetChildren() []Node
	// Insert will insert a node underneath a parent node.
	Insert(parentID uuid.UUID, newNode Node) error
	// Delete will delete a node underneath a parent node.
	Delete(parentID, childID uuid.UUID) error

	// JSON for persistance
	json.Marshaler
	json.Unmarshaler
}
