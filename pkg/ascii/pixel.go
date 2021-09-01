package ascii

import (
	"image/color"
	"math"
	"reflect"

	"github.com/aybabtme/rgbterm"
)

var (
	pixels = " .,:;i1tfLCG08@"
)

type PixelConverter struct {
}

func NewPixelConverter() *PixelConverter {
	return &PixelConverter{}
}

func (c *PixelConverter) ToASCIIString(pixel color.Color, options *Options) string {
	if options.Reversed {
		pixels = reverse(pixels)
	}

	r := reflect.ValueOf(pixel).FieldByName("R").Uint()
	g := reflect.ValueOf(pixel).FieldByName("G").Uint()
	b := reflect.ValueOf(pixel).FieldByName("B").Uint()
	a := reflect.ValueOf(pixel).FieldByName("A").Uint()
	value := (r + g + b) * a / 255

	precision := float64(255 * 3 / (len(pixels) - 1))
	roundValue := int(math.Floor(float64(value)/precision + 0.5))
	rawChar := pixels[roundValue]

	if options.Colored {
		return rgbterm.FgString(string([]byte{rawChar}), uint8(r), uint8(g), uint8(b))
	}
	return string([]byte{rawChar})
}

func reverse(s string) (ret string) {
	for _, v := range s {
		defer func(r rune) {
			ret += string(r)
		}(v)
	}
	return
}
