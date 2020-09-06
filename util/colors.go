package util

import (
	"image"
	"image/color"
	"math"
)

// http://www.glennchan.info/articles/technical/rec709rec601/rec709rec601.html
// https://onlineimagetools.com/grayscale-image

type rgbWeights struct {
	r float64
	g float64
	b float64
}

var BT601Weights = rgbWeights{
	r: 0.299,
	g: 0.587,
	b: 0.114,
}

var BT709Weights = rgbWeights{
	r: 0.21,
	g: 0.72,
	b: 0.07,
}

var AVGWeights = rgbWeights{
	r: 0.333333333333,
	g: 0.333333333333,
	b: 0.333333333333,
}

func NewBT601Weights() rgbWeights {
	return BT601Weights
}

func NewBT709Weights() rgbWeights {
	return BT709Weights
}
func NewAvgWeights() rgbWeights {
	return AVGWeights
}

type ColorIndex uint16

// GrayScaleMatrix is a matrix containing colorIndices
type GrayScaleMatrix struct {
	values []ColorIndex
	width  int
	height int
}

// NewMatrix creates a new GrayScaleMatrix with specified sizes
func NewMatrix(width int, height int) GrayScaleMatrix {
	m := GrayScaleMatrix{width: width, height: height, values: make([]ColorIndex, width*height)}
	return m
}

// GrayScaleMatrix.SetAt sets a  grayscale index at a matrix position
func (m *GrayScaleMatrix) SetAt(x int, y int, idx ColorIndex) {
	m.values[x+(y*m.width)] = idx
}

// GrayScaleMatrix.GetAt gets a grayscale index at a matrix position
func (m *GrayScaleMatrix) GetAt(x int, y int) ColorIndex {
	return m.values[x+(y*m.width)]
}

func convertToGrayScaleIndex(color color.Color, weights rgbWeights) ColorIndex {
	rr, gg, bb, _ := color.RGBA()
	gamma := 2.2
	r := math.Pow(float64(rr), gamma)
	g := math.Pow(float64(gg), gamma)
	b := math.Pow(float64(bb), gamma)
	m := math.Pow(weights.r*r+weights.g*g+weights.b*b, 1/gamma)
	Y := ColorIndex(m + 0.5)
	return Y >> 8
}

func CreateGrayscaleMap(image image.Image, weights rgbWeights) GrayScaleMatrix {
	bounds := image.Bounds()
	matrix := NewMatrix(bounds.Size().X, bounds.Size().Y)
	for y := 0; y < matrix.height; y++ {
		for x := 0; x < matrix.width; x++ {
			pixelColor := image.At(x, y)
			Y := convertToGrayScaleIndex(pixelColor, weights)
			matrix.SetAt(x, y, Y)
		}
	}
	return matrix
}
