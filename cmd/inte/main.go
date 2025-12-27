package main

import (
	"fmt"

	"gonum.org/v1/gonum/interp"
)

func main() {
	xs := []float64{0, 1, 2}
	ys := []float64{0, 2, 1}

	dydxs := make([]float64, len(xs))

	// Calculate derivatives at each point
	for i := 0; i < len(xs); i++ {
		if i == 0 {
			// First point: use forward difference
			dydxs[i] = (ys[i+1] - ys[i]) / (xs[i+1] - xs[i])
		} else if i == len(xs)-1 {
			// Last point: use backward difference
			dydxs[i] = (ys[i] - ys[i-1]) / (xs[i] - xs[i-1])
		} else {
			// Middle points: use central difference
			dydxs[i] = (ys[i+1] - ys[i-1]) / (xs[i+1] - xs[i-1])
		}
	}

	predictor := interp.PiecewiseCubic{}
	predictor.FitWithDerivatives(xs, ys, dydxs)

	for i := float64(0); i <= 2; i += 0.1 {
		fmt.Println(predictor.Predict(i))
	}
}
