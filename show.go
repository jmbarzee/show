package show

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jmbarzee/show/device"
	"github.com/jmbarzee/show/ifaces"
	"github.com/jmbarzee/show/node"
	"github.com/jmbarzee/space"
)

// TODO some sort of MarshallJSON

type Show struct {
	mu *sync.RWMutex
	// subs is the list of subscribers
	subs []Subscriber
	// nodeTree is the idealogical hieracy of show nodes
	nodeTree node.Node
}

func NewShow() (*Show, error) {
	root := node.NewGroupOption()

	s := &Show{
		mu:       &sync.RWMutex{},
		subs:     []Subscriber{},
		nodeTree: root,
	}
	return s, nil
}

// Allocate passes a vibe into the tree where it will be allocated to sub devices as it is Stabilized
func (s *Show) Allocate(vibe ifaces.Vibe) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nodeTree.Allocate(vibe)
}

// DispatchRenders dispatches renders to all connected subs
func (s *Show) DispatchRenders(ctx context.Context, t time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, sub := range s.subs {
		if !sub.IsConnected() {
			continue
		}

		if err := sub.DispatchRender(t); err != nil {
			// TODO handle errors
		}
	}
}

// InsertNode places a device into the tree underneath the device with parentID
func (s *Show) InsertNode(parentID, childID uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var childNode node.Node
	for _, sub := range s.subs {
		nodes := sub.Device.GetNodes()
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
func (s *Show) MoveDevice(deviceID uuid.UUID, loc space.Cartesian, ori space.Spherical) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, sub := range s.subs {
		if sub.GetID() != deviceID {
			continue
		}
		sub.SetLocation(loc)
		sub.SetOrientation(ori)
		return nil
	}
	return fmt.Errorf("Device %v not found", deviceID.String())
}

// ConnectDevice attaches a sender to an existing device or creates a new one if not existing
func (s *Show) ConnectDevice(dev device.Device, sender Sender) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, sub := range s.subs {
		if sub.GetID() == dev.GetID() && sub.GetType() == dev.GetType() {
			return sub.Connect(sender)
		}
	}

	// No previous device found, build new subscriber
	newSub := Subscriber{Device: dev, Sender: sender}
	s.subs = append(s.subs, newSub)
	return nil
}
