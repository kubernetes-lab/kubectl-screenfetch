package ascii

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"strings"

	"golang.org/x/image/draw"
)

type Options struct {
	Colored  bool
	Reversed bool
}

var DefaultOptions = Options{
	Colored:  true,
	Reversed: false,
}

type ImageConverter struct {
	imageResizer   *ImageResizer
	pixelConverter *PixelConverter
}

func NewImageConverter() *ImageConverter {
	return &ImageConverter{
		imageResizer:   NewImageResizer(),
		pixelConverter: NewPixelConverter(),
	}
}

func (c *ImageConverter) ToASCIIString(srcImage image.Image, options *Options) string {
	newImage := c.imageResizer.Scale(srcImage, options)
	newWidth := newImage.Bounds().Max.X
	newHeight := newImage.Bounds().Max.Y

	chars := make([]string, 0, newWidth*newHeight+newWidth)
	for i := 0; i < newHeight; i++ {
		for j := 0; j < newWidth; j++ {
			p := color.NRGBAModel.Convert(newImage.At(j, i))
			char := c.pixelConverter.ToASCIIString(p, options)
			chars = append(chars, char)
		}
		chars = append(chars, "\n")
	}

	return strings.Join(chars, "")
}

type ImageResizer struct {
	terminal *Terminal
}

func NewImageResizer() *ImageResizer {
	return &ImageResizer{
		terminal: NewTerminal(),
	}
}

func (r *ImageResizer) Scale(srcImage image.Image, options *Options) image.Image {
	sr := srcImage.Bounds()
	screenWidth, screenHeight, err := r.terminal.ScreenSize()
	if err != nil {
		log.Fatal(err)
	}

	ratio := float64(screenHeight) / float64(sr.Max.Y)
	scaledWidth := float64(sr.Max.X) * ratio / r.terminal.CharWidth()
	if scaledWidth < float64(screenWidth) {
		ratio = ratio / r.terminal.CharWidth()
	} else {
		ratio = float64(screenWidth) / float64(sr.Max.X)
	}

	newWidth := int(float64(sr.Max.X) * ratio)
	newHeight := int(float64(sr.Max.Y) * ratio * r.terminal.CharWidth())

	if newWidth > 50 {
		newWidth = 50
	}
	if newHeight > 25 {
		newHeight = 25
	}

	dr := image.Rect(0, 0, newWidth, newHeight)
	dstImage := image.NewNRGBA(dr)
	draw.BiLinear.Scale(dstImage, dr, srcImage, srcImage.Bounds(), draw.Over, nil)
	return dstImage
}
