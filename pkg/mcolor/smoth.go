package mcolor

import (
	"image"
	"image/color"
	"math"
	"sort"

	"github.com/dontomato/mandelbrot/pkg/mcalc"
)

func createPaletteRGBA(data []mcalc.MandelIterationResult, settings *mcalc.MandelSettings, pal []color.Color) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, settings.Width, settings.Height))

	singleValues := make([]float64, settings.Width*settings.Height)
	i := 0
	for py := 0; py < settings.Height; py++ {
		for px := 0; px < settings.Width; px++ {
			singleValues[i] = float64(settings.IterationCount-data[i].Iteration) + math.Log(data[i].Z)
			i++
		}
	}

	sortedValues := make([]float64, len(singleValues))
	for i := range sortedValues {
		sortedValues[i] = singleValues[i]
	}

	sort.Float64s(sortedValues)

	splitValues := make([]float64, len(pal)-1)

	factor := .98
	start := .9

	for i := range splitValues {
		index := (i + 1) * len(sortedValues) / len(pal)
		splitValues[i] = sortedValues[index]
		start *= factor
	}

	sort.Float64s(splitValues)

	i = 0
	for py := 0; py < settings.Height; py++ {
		for px := 0; px < settings.Width; px++ {
			colorIndex := sort.Search(len(splitValues), func(j int) bool { return singleValues[i] < splitValues[j] })
			img.Set(px, py, pal[colorIndex])
			i++
		}
	}

	return img
}
