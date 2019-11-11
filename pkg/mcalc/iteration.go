package mcalc

// MandelIterationResult describes result for calculations for point
type MandelIterationResult struct {
	Z         float64
	Iteration int
}

// IsBlack ckeck point is included in the Mandelbrot set
func (r *MandelIterationResult) IsBlack(settings *MandelSettings) bool {
	return r.Iteration == settings.IterationCount && r.Z == 0
}
