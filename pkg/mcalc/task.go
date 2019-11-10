package mcalc

import (
	"image"
)

// MandelConstParameters describes parameters for task which are constant
type MandelConstParameters struct {
	Width          int
	Height         int
	IterationCount int
}

// MandelSetParameters describes parameters for task which creates set of Masdelbrot set images
type MandelSetParameters struct {
	Settings        *MandelConstParameters
	SetCount        int
	GetInitialValue func(n int) complex128
	GetRealWidth    func(n int) float64
	GetCoordinates  func(c int) (float64, float64)
	CreateRGBA      func(result []MandelIterationResult, settings *MandelConstParameters) *image.RGBA
}

// MandelPictureParameters describes parameters for calculation of Masdelbrot set
type MandelPictureParameters struct {
	Settings   *MandelConstParameters
	Z0         complex128
	RealWidth  float64
	X, Y       float64
	CreateRGBA func(result []MandelIterationResult, settings *MandelConstParameters) *image.RGBA
}

// MandelIterationResult describes result for calculations for point
type MandelIterationResult struct {
	Z         float64
	Iteration int
}
