package color

import "math"

// RGB is a color represented by Red, Green, Blue
type RGB struct {
	// R, G, B all range between (0, 1)
	R, G, B, A float64
}

var _ Color = (*RGB)(nil)

// HSL returns a color represented by Hue, Saturation, and Lightness
func (c RGB) HSL() HSL {
	var h, s, l float64

	r := c.R
	g := c.G
	b := c.B

	max := max(r, g, b)
	min := min(r, g, b)

	// Luminosity is the average of the max and min rgb color intensities.
	l = (max + min) / 2

	// saturation
	delta := max - min
	if delta == 0 {
		// it's gray
		return HSL{0, 0, l, c.A}
	}

	// it's not gray
	if l < 0.5 {
		s = delta / (max + min)
	} else {
		s = delta / (2 - max - min)
	}

	// hue
	r2 := (((max - r) / 6) + (delta / 2)) / delta
	g2 := (((max - g) / 6) + (delta / 2)) / delta
	b2 := (((max - b) / 6) + (delta / 2)) / delta
	switch {
	case r == max:
		h = b2 - g2
	case g == max:
		h = (1.0 / 3.0) + r2 - b2
	case b == max:
		h = (2.0 / 3.0) + g2 - r2
	}

	// fix wraparounds
	switch {
	case h < 0:
		h++
	case h > 1:
		h--
	}

	return HSL{h, s, l, c.A}
}

// RGB returns a color represented by Red, Green, Blue
func (c RGB) RGB() RGB {
	return c
}

// ToUInt32RGBW converts an RGBA to a uint32.
// Each portion of the color is represented by 8 bits
func (c RGB) ToUInt32RGBW() uint32 {
	val := uint32(c.R*math.MaxUint8) << 0
	val |= uint32(c.G*math.MaxUint8) << 8
	val |= uint32(c.B*math.MaxUint8) << 16
	val |= uint32(c.A*math.MaxUint8) << 24
	return val
}

// FromUInt32RGBW converts and uint32 to an RGBA.
// Each portion of the color is represented by 8 bits
func FromUInt32RGBW(wgrb uint32) RGB {
	mask := uint32(0x000000ff)

	uint8r := mask & (wgrb >> 0)
	r := float64(uint8r) / math.MaxUint8
	uint8g := mask & (wgrb >> 8)
	g := float64(uint8g) / math.MaxUint8
	uint8b := mask & (wgrb >> 16)
	b := float64(uint8b) / math.MaxUint8
	uint8a := mask & (wgrb >> 24)
	a := float64(uint8a) / math.MaxUint8
	return RGB{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
