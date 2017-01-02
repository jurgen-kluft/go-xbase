package color

import (
	"errors"
	"fmt"
	"image/color"
	"math"
)

// Min3u8 returns the smallest of three numbers
func min3u8(a, b, c uint8) uint8 {
	if (a < b) && (a < c) {
		return a
	} else if (b < a) && (b < c) {
		return b
	}
	return c
}

// Max3u8 returns the largest of three numbers
func max3u8(a, b, c uint8) uint8 {
	if (a >= b) && (a >= c) {
		return a
	} else if (b >= a) && (b >= c) {
		return b
	}
	return c
}

// HSV is a color representation in 8-bit
type HSV struct {
	Hue, Sat, Val uint8
}

// RGBtoHSV Convert RGB to HSV
func RGBtoHSV(cr color.RGBA) (hsv HSV) {
	rgbMin := min3u8(cr.R, cr.G, cr.B)
	rgbMax := max3u8(cr.R, cr.G, cr.B)

	hsv.Val = rgbMax
	if hsv.Val == 0 {
		hsv.Hue = 0
		hsv.Sat = 0
		return
	}

	hsv.Sat = 255 * (rgbMax - rgbMin) / hsv.Val
	if hsv.Sat == 0 {
		hsv.Hue = 0
		return
	}

	span := (rgbMax - rgbMin)
	if rgbMax == cr.R {
		hsv.Hue = 43 * (cr.G - cr.B) / span
	} else if rgbMax == cr.G {
		hsv.Hue = 85 + 43*(cr.B-cr.R)/span
	} else { /* rgbMax == cr.B */
		hsv.Hue = 171 + 43*(cr.R-cr.G)/span
	}
	return
}

// RGBf is a color representation in float64
type RGBf struct {
	R, G, B float64
}

// HSLf is a color representation in float64
type HSLf struct {
	H, S, L float64
}

// HTMLToRGB Takes a string like '#123456' or 'ABCDEF' and returns an RGB
func HTMLToRGB(in string) (RGBf, error) {
	if in[0] == '#' {
		in = in[1:]
	}

	if len(in) != 6 {
		return RGBf{}, errors.New("Invalid string length")
	}

	var r, g, b byte
	if n, err := fmt.Sscanf(in, "%2x%2x%2x", &r, &g, &b); err != nil || n != 3 {
		return RGBf{}, err
	}

	return RGBf{float64(r) / 255, float64(g) / 255, float64(b) / 255}, nil
}

// ToHSL converts a RGBf to a representation in HSLf
func (c RGBf) ToHSL() HSLf {
	var h, s, l float64

	r := c.R
	g := c.G
	b := c.B

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)

	// Luminosity is the average of the max and min rgb color intensities.
	l = (max + min) / 2

	// saturation
	delta := max - min
	if delta == 0 {
		// it's gray
		return HSLf{0, 0, l}
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

	return HSLf{h, s, l}
}

// A nudge to make truncation round to nearest number instead of flooring
const delta = 1 / 512.0

// ToHTML will convert a RGB float color into a HTML color string
func (c RGBf) ToHTML() string {
	return fmt.Sprintf("%02x%02x%02x", byte((c.R+delta)*255), byte((c.G+delta)*255), byte((c.B+delta)*255))
}

func hueToRGB(v1, v2, h float64) float64 {
	if h < 0 {
		h++
	}
	if h > 1 {
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

// ToRGB converts a HSL color to RGB
func (c HSLf) ToRGB() RGBf {
	h := c.H
	s := c.S
	l := c.L

	if s == 0 {
		// it's gray
		return RGBf{l, l, l}
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

	return RGBf{r, g, b}
}

// ToHTML converts a HSL color into a HTML color string
func (c HSLf) ToHTML() string {
	return c.ToRGB().ToHTML()
}
