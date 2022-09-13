package image_processing

import (
	"image"
	"image/color"
)

type ImageItem struct {
	Path             string
	Filename         string
	OutputDirectory  string
	ImageData        image.Image
	PointToFloodFill image.Point
	ColorToFloodFill color.Color
	ColorTolerance   uint8
	Quality          int
	Histogram        image.Image
}
