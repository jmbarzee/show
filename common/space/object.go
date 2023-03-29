package space

// An Object is something which exists in space
type Object struct {
	// location is the location of the Object
	location Vector
	// orientation is Orientation of the Object
	orientation Quaternion
}

// NewObject creates an Object
func NewObject(location Vector, orientation Quaternion) *Object {
	return &Object{
		location:    location,
		orientation: orientation,
	}
}

// GetLocation returns the location of the Object
func (o Object) GetLocation() Vector {
	return o.location
}

// SetLocation changes the location of the Object
func (o *Object) SetLocation(newLocation Vector) {
	o.location = newLocation
}

// GetOrientation returns the Orientation of the Object
func (o Object) GetOrientation() Quaternion {
	return o.orientation
}

// SetOrientation changes the Orientation of the Object
func (o *Object) SetOrientation(newOrientation Quaternion) {
	o.orientation = newOrientation
}

// GetBearings returns all properties of the Object
func (o Object) GetBearings() (location Vector, orientation Quaternion) {
	return o.location, o.orientation
}

// Move changes all properties of the Object
func (o *Object) Move(location Vector, orientation Quaternion) {
	o.location = location
	o.orientation = orientation
}
