package show

import (
	"fmt"
	"time"

	"github.com/jmbarzee/show/device"
)

type Sender interface {
	Send(t time.Time, i device.Instruction) error
	IsConnected() bool
	Disconnect() error
}

// Subscriber represents a light service which has subscribed to light updates
// from the LightOrchestrator
type Subscriber struct {
	device.Device
	Sender
}

// IsConnected returns true if the server and kill functions are non-nil
func (s Subscriber) IsConnected() bool {
	return s.Sender != nil && s.IsConnected()
}

// Connect adds a connection to an existing subscriber
func (s *Subscriber) Connect(sender Sender) error {
	var err error
	if s.IsConnected() {
		err = s.Disconnect()
	}
	s.Sender = sender
	return err
}

// Disconnect Ends a subscribers connection
func (s *Subscriber) Disconnect() error {
	if s.Sender == nil {
		return fmt.Errorf("couldn't disconect sub: sender is nil")
	}

	var err error
	if s.IsConnected() {
		err = s.Disconnect()
	}

	s.Sender = nil
	return err
}

// DispatchRender sends lights after a subscriber's device renders them based on t
func (s Subscriber) DispatchRender(t time.Time) error {
	if !s.IsConnected() {
		return fmt.Errorf("couldn't dispatch render: sender not connected")
	}
	return s.Send(t, s.Render(t))
}
