package keyframes

import (
	"fmt"

	"gonum.org/v1/gonum/interp"
)

type PiecewiseCubicPredictor struct {
	derivativePredictor interp.PiecewiseCubic
}

var _ interp.Predictor = (*PiecewiseCubicPredictor)(nil)
var _ interp.Fitter = (*PiecewiseCubicPredictor)(nil)

func (p *PiecewiseCubicPredictor) Predict(t float64) float64 {
	return p.derivativePredictor.Predict(t)
}

func (p *PiecewiseCubicPredictor) Fit(xs, ys []float64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("FitWithDerivatives panicked: %v", r)
		}
	}()

	derivatives := p.computeDerivatives(xs, ys)
	p.derivativePredictor.FitWithDerivatives(xs, ys, derivatives)

	return nil
}

func (p *PiecewiseCubicPredictor) computeDerivatives(xs, ys []float64) []float64 {
	if len(xs) == 0 {
		return []float64{}
	}

	if len(xs) == 1 {
		return []float64{0}
	}

	derivatives := make([]float64, len(xs))

	for i := 0; i < len(xs); i++ {
		if i == 0 {
			// First point: use forward difference
			derivatives[i] = (ys[i+1] - ys[i]) / (xs[i+1] - xs[i])
		} else if i == len(xs)-1 {
			// Last point: use backward difference
			derivatives[i] = (ys[i] - ys[i-1]) / (xs[i] - xs[i-1])
		} else {
			// Middle points: use central difference
			derivatives[i] = (ys[i+1] - ys[i-1]) / (xs[i+1] - xs[i-1])
		}
	}

	return derivatives
}
