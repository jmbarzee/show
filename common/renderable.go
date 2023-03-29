package common

import (
	"github.com/jmbarzee/show/common/color"
	"github.com/jmbarzee/show/common/space"
)

type Item interface {
	Indexed
	Orientable
}

type Renderable interface {
	Item
	Colorable
}

// Indexed is the interface to something which is in an ordered set
type Indexed interface {
	// GetPosition returns the position and total
	// i.e. #3 of 10
	GetPosition() (int, int)
}

type (
	// Colored is the interface to something which is colored
	Colored interface {
		// GetColor returns the color of the light
		GetColor() color.Color
	}
	// Colorable is the interface to something which is colorable
	Colorable interface {
		Colored

		// SetColor changes the color of the light
		SetColor(newColor color.Color)
	}
)

type (
	// Located is the interface to something which exists in space
	Located interface {
		// GetLocation returns the location of the objcet
		GetLocation() space.Vector
	}

	// Locatable is the interface to something which can be positioned in space
	Locatable interface {
		Located

		// SetLocation changes the location of the objcet
		SetLocation(space.Vector)
	}
)

type (
	// Oriented is the interface to something which is oriented in space
	Oriented interface {
		Located

		// GetOrientation returns the Orientation of the object
		GetOrientation() space.Quaternion
	}

	// Orientable is the interface to something which is orientable in space
	Orientable interface {
		Locatable

		// GetOrientation returns the Orientation of the object
		GetOrientation() space.Quaternion
		// SetOrientation changes the Orientation of the object
		SetOrientation(space.Quaternion)
	}
)

var _ Orientable = (*space.Object)(nil)

type (
	// Moveable is the interface to something which is moveable in space
	// (Should have no symmetry)
	Moveable interface {
		// GetBearings returns all spacial properties of the device
		GetBearings() (location space.Vector, orientation space.Quaternion)
		// Move changes all properties of the device
		Move(location space.Vector, orientation space.Quaternion)
	}
)

var _ Moveable = (*space.Object)(nil)
