package device

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/common/ifaces"
	"github.com/jmbarzee/show/common/node"
)

// Device represents a physical device with lights
// A device is made up of atleast a single Node
type Device interface {
	// GetNodes returns all the Nodes which the device holds
	GetNodes() []node.Node
	// GetType returns the type
	GetType() string
	// GetID will return the ID of a device node.
	GetID() uuid.UUID

	// DispatchRender produces and dispatches Instructions
	DispatchRender(time.Time) error

	ifaces.Tangible
}
