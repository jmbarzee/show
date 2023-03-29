package color

import "math"

// Color is any representation of a color
type Color interface {
	// HSL returns a color represented by Hue, Saturation, and Lightness
	HSL() HSL
	// RGB returns a color represented by Red, Green, Blue
	RGB() RGB
}

var (
	Min  = math.SmallestNonzeroFloat64
	Max  = 1.0 - Min
	Half = 0.5
)

var (
	Black        = HSL{Min, Min, Min, Min}
	Grey         = HSL{Min, Min, Half, Half}
	GreyNatural  = HSL{Min, Min, Min, Half}
	White        = HSL{Min, Min, Max, Max}
	WhiteNatural = HSL{Min, Min, Min, Max}

	Red         = HSL{Min, Max, Half, Min}
	WarmRed     = HSL{1.0 / 24, Max, Half, Min}
	Orange      = HSL{2.0 / 24, Max, Half, Min}
	WarmYellow  = HSL{3.0 / 24, Max, Half, Min}
	Yellow      = HSL{4.0 / 24, Max, Half, Min}
	CoolYellow  = HSL{5.0 / 24, Max, Half, Min}
	YellowGreen = HSL{6.0 / 24, Max, Half, Min}
	WarmGreen   = HSL{7.0 / 24, Max, Half, Min}
	Green       = HSL{8.0 / 24, Max, Half, Min}
	CoolGreen   = HSL{9.0 / 24, Max, Half, Min}
	GreenCyan   = HSL{10.0 / 24, Max, Half, Min}
	WarmCyan    = HSL{11.0 / 24, Max, Half, Min}
	Cyan        = HSL{12.0 / 24, Max, Half, Min}
	CoolCyan    = HSL{13.0 / 24, Max, Half, Min}
	BlueCyan    = HSL{14.0 / 24, Max, Half, Min}
	CoolBlue    = HSL{15.0 / 24, Max, Half, Min}
	Blue        = HSL{16.0 / 24, Max, Half, Min}
	WarmBlue    = HSL{17.0 / 24, Max, Half, Min}
	Violet      = HSL{18.0 / 24, Max, Half, Min}
	CoolMagenta = HSL{19.0 / 24, Max, Half, Min}
	Magenta     = HSL{20.0 / 24, Max, Half, Min}
	WarmMagenta = HSL{21.0 / 24, Max, Half, Min}
	RedMagenta  = HSL{22.0 / 24, Max, Half, Min}
	CoolRed     = HSL{23.0 / 24, Max, Half, Min}

	AllColors = []Color{
		Red,
		WarmRed,
		Orange,
		WarmYellow,
		Yellow,
		CoolYellow,
		YellowGreen,
		WarmGreen,
		Green,
		CoolGreen,
		GreenCyan,
		WarmCyan,
		Cyan,
		CoolCyan,
		BlueCyan,
		CoolBlue,
		Blue,
		WarmBlue,
		Violet,
		CoolMagenta,
		Magenta,
		WarmMagenta,
		RedMagenta,
		CoolRed,
	}
)

func max(a, b, c float64) float64 {
	m := a
	if b > m {
		m = b
	}
	if c > m {
		m = c
	}
	return m
}

func min(a, b, c float64) float64 {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}
func modOne(val float64) float64 {
	return math.Mod(val, 1.0)
}
