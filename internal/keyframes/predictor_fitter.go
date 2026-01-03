package keyframes

import "gonum.org/v1/gonum/interp"

type PredictorFitter interface {
	interp.Predictor
	interp.Fitter
}
