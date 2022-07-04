package addrled

import (
	"encoding/binary"

	"github.com/jmbarzee/show/ifaces"
)

type instruction struct {
	lights []ifaces.Light
}

func (i instruction) Package() ([]byte, error) {
	bytes := make([]byte, 0, len(i.lights)*4)
	buffer := make([]byte, 4)
	for _, light := range i.lights {
		rgb := light.GetColor().RGB()
		uint32enc := rgb.ToUInt32RGBW()
		binary.LittleEndian.PutUint32(buffer, uint32enc)
		bytes = append(bytes, buffer...)
	}
	return bytes, nil
}
