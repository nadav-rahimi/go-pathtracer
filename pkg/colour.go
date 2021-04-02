package pathtracer

// Struct to hold the RGB colour data where each component
// is in the range 0.0 to 1.0 corresponding to the range 0-255
type Colour struct {
	R, G, B float64
}

// Predefined Colours
var (
	Red          = Colour{1.0, 0, 0}
	Grey         = Colour{0.5, 0.5, 0.5}
	Blue         = Colour{0.5, 0.7, 1.0}
	Black        = Colour{0, 0, 0}
	White        = Colour{1.0, 1.0, 1.0}
	Orange       = Colour{0.9961, 0.4314, 0.}
	Salmon       = Colour{1, 0.5490196, 0.411764}
	Yellow       = Colour{1, 0.8745, 0}
	LightRed     = Colour{1, 0.62352, 0.5019}
	OffWhite     = Colour{0.992157, 0.960784, 0.9019608}
	Chestnut     = Colour{0.23529, 0.160784, 0.1490196}
	DarkBlue     = Colour{0.05882, 0.05882, 1}
	SortaBlue    = Colour{0.2, 0.2, 1}
	LightPink    = Colour{1, 0.7137, 0.7568}
	LightBlue    = Colour{0.7686, 0.7686, 1}
	LightGreen   = Colour{0.466666, 0.74901, 0.37647}
	PaleYellow   = Colour{0.9568, 0.8156, 0.24705}
	SpotifyGreen = Colour{0.07843, 0.843137, 0.3764706}
)

// Returns the red component of the colour in the range 0-255
func (c Colour) R256() uint8 {
	return uint8(255.99 * c.R)
}

// Returns the green component of the colour in the range 0-255
func (c Colour) G256() uint8 {
	return uint8(255.99 * c.G)
}

// Returns the blue component of the colour in the range 0-255
func (c Colour) B256() uint8 {
	return uint8(255.99 * c.B)
}

// Adds two colours
func (c Colour) Add(o Colour) Colour {
	return Colour{c.R + o.R, c.G + o.G, c.B + o.B}
}

// Multiplies two colours
func (c Colour) Mul(o Colour) Colour {
	return Colour{c.R * o.R, c.G * o.G, c.B * o.B}
}

// Multiplies colour by value "f"
func (c Colour) MulFloat(f float64) Colour {
	return Colour{c.R * f, c.G * f, c.B * f}
}

// Divides colour by value "f"
func (c Colour) DivFloat(f float64) Colour {
	return Colour{c.R / f, c.G / f, c.B / f}
}

// Returns the point at the gradient between two
// colours, point should be between 0.0 and 1.0
func Gradient(a, b Colour, f float64) Colour {
	// scale between 0.0 and 1.0
	f = 0.5 * (f + 1.0)

	// linear blend: blended_value = (1 - f) * a + f * b
	return a.MulFloat(1.0 - f).Add(b.MulFloat(f))
}
