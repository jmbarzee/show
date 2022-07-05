package show

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/common"
	"github.com/jmbarzee/show/common/node"
	"github.com/jmbarzee/space"
)

// TODO some sort of MarshallJSON

type Show struct {
	mu *sync.RWMutex
	// devices is the list of devices
	devices []common.Device
	// nodeTree is the idealogical hieracy of show nodes
	nodeTree common.Node
}

func NewShow() (*Show, error) {
	root := node.NewGroupOption()

	s := &Show{
		mu:       &sync.RWMutex{},
		devices:  []common.Device{},
		nodeTree: root,
	}
	return s, nil
}

// Allocate passes a vibe into the tree where it will be allocated to sub devices as it is Stabilized
func (s *Show) Allocate(vibe common.Vibe) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nodeTree.Allocate(vibe)
}

// DispatchRenders dispatches renders to all connected subs
func (s *Show) DispatchRenders(ctx context.Context, t time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, device := range s.devices {

		if err := device.DispatchRender(t); err != nil {
			// TODO handle errors
		}
	}
}

// InsertNode places a device into the tree underneath the device with parentID
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

// DeleteNode removes a device from the tree underneath the device with parentID
func (s *Show) DeleteNode(parentID, childID uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.nodeTree.Delete(parentID, childID)
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
