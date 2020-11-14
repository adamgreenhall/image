package tiff

import (
	"image"
	"image/color"
	"log"
)

// ColorNChannel is a bogus color model that doesn't convert to RGBA
type ColorNChannel struct {
	Values []uint8
}

// ImageNChannel is an in-memory image whose At method returns ColorNChannel values.
type ImageNChannel struct {
	// Pix holds the image's pixels, in channel order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*Nchannels].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
	// Nchannels is the number of channels
	Nchannels int
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *ImageNChannel) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*p.Nchannels
}

// needed to implement image.Image interface
func nChannelModel(c color.Color) color.Color {
	log.Fatal("cannot convert nchannel to RGBA")
	return color.RGBA{}
}

// NChannelModel is a bogus color model
var NChannelModel color.Model = color.ModelFunc(nChannelModel)

// ColorModel is bogus for NChannel
func (p *ImageNChannel) ColorModel() color.Model { return NChannelModel }

// Bounds is the image's rectangle
func (p *ImageNChannel) Bounds() image.Rectangle { return p.Rect }

// At is bogus
func (p *ImageNChannel) At(x, y int) color.Color {
	log.Fatal("cannot convert nchannel to RGBA")
	return color.RGBA{}
}

// NewNChannel creates a new ImageNChannel
func NewNChannel(r image.Rectangle, nChannels int) *ImageNChannel {
	w, h := r.Dx(), r.Dy()
	buf := make([]uint8, nChannels*w*h)
	return &ImageNChannel{buf, nChannels * w, r, nChannels}
}
