package show

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/node"
	"github.com/jmbarzee/show/common/space"
)

// TODO some sort of MarshallJSON

type Show struct {
	mu *sync.RWMutex
	// devices is the list of devices
	devices map[uuid.UUID]common.Device
	// nodeTree is the ideological hierarchy of show nodes
	nodeTree common.Node
}

func New() *Show {
	return &Show{
		mu:       &sync.RWMutex{},
		devices:  map[uuid.UUID]common.Device{},
		nodeTree: node.NewGroupOption(),
	}
}

// Allocate passes a vibe into the tree where it will be allocated to sub devices as it is Stabilized
func (s *Show) Allocate(vibe common.Vibe) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nodeTree.Allocate(vibe)
}

// DispatchRenders dispatches renders to all connected subs
func (s *Show) DispatchRenders(t time.Time) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, device := range s.devices {

		if err := device.DispatchRender(t); err != nil {
			// TODO handle errors
		}
	}
}

// InsertNode places a node, found from existing devices,
// into the tree underneath the node with parentID
func (s *Show) InsertNode(parentID, childID uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var childNode common.Node
	for _, device := range s.devices {
		nodes := device.GetNodes()
		for _, n := range nodes {
			if n.GetID() == childID {
				childNode = n
			}
		}

		if childNode != nil {
			break
		}
	}

	if childNode == nil {
		return errors.New("Could not find specified Child")
	}

	return s.nodeTree.Insert(parentID, childNode)
}

// NewNode creates a new node of the given type
// and inserts it into the tree underneath the node with parentID
// the id of the new node is returned
func (s *Show) NewNode(parentID uuid.UUID, nodeType string) (uuid.UUID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var childNode common.Node
	switch nodeType {
	case node.GroupOptionType:
		childNode = node.NewGroupOption()
	case node.GroupType:
		childNode = node.NewGroup()
	default:
		return uuid.UUID{}, errors.New("Could not find specified nodeType")
	}

	return childNode.GetID(), s.nodeTree.Insert(parentID, childNode)
}

// Returns the info of the root node
func (s Show) GetRootNodeInfo() common.NodeInfo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.nodeTree
}

// GetNodeInfo returns the info about a node, if found in the tree
func (s *Show) GetNodeInfo(nodeID uuid.UUID) (common.NodeInfo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	nodeInfo := s.nodeTree.GetNodeInfo(nodeID)
	if nodeInfo == nil {
		return nil, fmt.Errorf("Node %v not found", nodeID.String())
	}
	return nodeInfo, nil
}

// AddDevice add a device to the list of devices which is used for dispatching renders
func (s *Show) AddDevice(device common.Device) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	deviceID := device.GetID()
	if _, ok := s.devices[deviceID]; ok {
		return fmt.Errorf("Device %v was already added", deviceID)
	}
	s.devices[deviceID] = device
	return nil
}

// MoveDevice changes a devices location and orientation (rotation implicitly also)
func (s *Show) MoveDevice(deviceID uuid.UUID, bearing space.Object) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, device := range s.devices {
		if device.GetID() != deviceID {
			continue
		}
		device.Move(bearing.GetBearings())
		return nil
	}
	return fmt.Errorf("Device %v not found", deviceID.String())
}

// GetDeviceInfo returns the info about a device, if found
func (s *Show) GetDeviceInfo(deviceID uuid.UUID) (common.DeviceInfo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, device := range s.devices {
		if device.GetID() != deviceID {
			continue
		}
		return device, nil
	}
	return nil, fmt.Errorf("Device %v not found", deviceID.String())
}

// GetDeviceInfoAll returns the info for all devices
func (s *Show) GetDeviceInfoAll() []common.DeviceInfo {
	s.mu.Lock()
	defer s.mu.Unlock()

	devices := make([]common.DeviceInfo, len(s.devices))

	i := 0
	for _, device := range s.devices {
		devices[i] = device
		i++
	}
	return devices
}

// DeleteNode removes a device from the tree underneath the device with parentID
func (s *Show) DeleteNode(parentID, childID uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.nodeTree.Delete(parentID, childID)
}

// Clean removes all stored resources which ended before t
func (s *Show) Clean(t time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nodeTree.Clean(t)
}
