package bmp

import (
	"image"
	"math"
)
// BmpImage wraps [image.Image] adding a [BmpResolution] field.
type BmpImage struct {
	image.Image
	Resolution BmpResolution
}

// BmpResolution represent a BMP file resolution
// expressed in pixel per meter.
type BmpResolution struct {
	// XResolution is the horizontal component of the resolution.
	XResolution uint32
	// YResolution is the vertical component of the resolution.
	YResolution uint32
}

// FromDpi utility function building a [BmpResolution] data
// from a standard dot-per-inch value.
func FromDpi(dpi int) BmpResolution {
	ppm := uint32(math.Round((float64(dpi) * 1000) / 25.4))
	return BmpResolution{
		YResolution: ppm,
		XResolution: ppm,
	}
}
