package color

type (
	// HSL is a color represented by Hue, Saturation, and Lightness
	HSL struct {
		// H, S, L, A all range between (0, 1)
		H, S, L, A float64
	}
)

var _ Color = (*HSL)(nil)

// HSL returns a color represented by Hue, Saturation, and Lightness
func (c HSL) HSL() HSL {
	return c
}

// RGB returns a color represented by Red, Green, Blue
func (c HSL) RGB() RGB {

	hueToRGB := func(v1, v2, h float64) float64 {
		if h < 0 {
			h++
		} else if h > 1 {
			h--
		}
		switch {
		case 6*h < 1:
			return (v1 + (v2-v1)*6*h)
		case 2*h < 1:
			return v2
		case 3*h < 2:
			return v1 + (v2-v1)*((2.0/3.0)-h)*6
		}
		return v1
	}

	h := c.H
	s := c.S
	l := c.L

	if s == 0 {
		// it's gray
		return RGB{l, l, l, c.A}
	}

	var v1, v2 float64
	if l < 0.5 {
		v2 = l * (1 + s)
	} else {
		v2 = (l + s) - (s * l)
	}

	v1 = 2*l - v2

	r := hueToRGB(v1, v2, h+(1.0/3.0))
	g := hueToRGB(v1, v2, h)
	b := hueToRGB(v1, v2, h-(1.0/3.0))

	return RGB{r, g, b, c.A}
}

// SetHue will change hue to h (with wrapping).
func (c *HSL) SetHue(h float64) {
	hue := modOne(h)
	if hue < 0 {
		c.H = 1.0 + hue
	} else {
		c.H = hue
	}
}

// ShiftHue will shift hue by h (with wrapping).
func (c *HSL) ShiftHue(h float64) {
	c.SetHue(c.H + h)
}

// SetSaturation will change saturation to s (with bounding).
func (c *HSL) SetSaturation(s float64) {
	if s > Max {
		c.S = Max
	} else if s < Min {
		c.S = Min
	} else {
		c.S = s
	}
}

// SetLightness will change lightness to l (with bounding).
func (c *HSL) SetLightness(l float64) {
	if l > Max {
		c.L = Max
	} else if l < Min {
		c.L = Min
	} else {
		c.L = l
	}
}
