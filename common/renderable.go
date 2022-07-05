package common

import (
	"github.com/jmbarzee/color"
	"github.com/jmbarzee/space"
)

type Item interface {
	Indexed
	Pointed
}

type Renderable interface {
	Item
	Colorable
}

// Colorable is the interface to something which is colorable
type Colorable interface {
	// GetColor returns the color of the light
	GetColor() color.Color

	// SetColor changes the color of the light
	SetColor(newColor color.Color)
}

// Indexed is the interface to something which is in an ordered set
type Indexed interface {
	// GetPosition returns the position and total
	// i.e. #3 of 10
	GetPosition() (int, int)
}

type (
	// Located is the interface to something which exists in space
	// (should have point symmetry)
	Located interface {
		// GetLocation returns the physical location of the device
		GetLocation() space.Cartesian
	}

	// Locatable is the interface to something which is in space
	// (should have point symmetry)
	Locatable interface {
		Located

		// SetLocation changes the physical location of the device
		SetLocation(space.Cartesian)
	}
)

type (
	// Pointed is the interface to something which is Pointed in space
	// (should have rotational symmetry about an axis)
	Pointed interface {
		Located

		// GetOrientation returns the physical orientation of the device
		GetOrientation() space.Spherical
	}

	// Tangible is the interface to something which is pointable in space
	// (should have rotational symmetry about an axis)
	Pointable interface {
		Locatable

		// GetOrientation returns the physical orientation of the device
		GetOrientation() space.Spherical
		// SetOrientation changes the physical orientation of the device
		SetOrientation(space.Spherical)
	}
)

type (
	// Oriented is the interface to something which is oriented in space
	// (Should have no symmetry)
	Oriented interface {
		Pointed

		// GetRotation returns the physical rotation of the device
		GetRotation() space.Spherical
	}

	// Orientable is the interface to something which is orientable in space
	// (Should have no symmetry)
	Orientable interface {
		Pointable

		// GetRotation returns the physical rotation of the device
		GetRotation() space.Spherical
		// SetRotationchanges the physical rotation of the device
		SetRotation(space.Spherical)
	}
)

var _ Orientable = (*space.Object)(nil)

type (
	// Moveable is the interface to something which is moveable in space
	// (Should have no symmetry)
	Moveable interface {
		// GetBearings returns all spacial properties of the device
		GetBearings() (location space.Cartesian, orientation, rotation space.Spherical)
		// Move changes all properties of the device
		Move(location space.Cartesian, orientation, rotation space.Spherical)
	}
)

var _ Moveable = (*space.Object)(nil)
