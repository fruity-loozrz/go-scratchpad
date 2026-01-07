package interpolation

import (
	"errors"
	"fmt"
)

var ErrTooFewPoints = errors.New("too few points")
var ErrArgumentsNotSameLength = errors.New("argument lists not same length")
var ErrWrongEasingFnsLength = errors.New("wrong easing functions length")

func EaseBetween(x0, x1, y0, y1, x float64, easingFn func(float64) float64) float64 {
	return y0 + (y1-y0)*easingFn((x-x0)/(x1-x0))
}

func PieceWiseEaseInterpolate(
	xs []float64,
	ys []float64,
	easingFns []func(float64) float64,
	x float64,
) (float64, error) {
	if len(xs) != len(ys) {
		return 0, fmt.Errorf("argument lists not same length: xs=%d, ys=%d %w", len(xs), len(ys), ErrArgumentsNotSameLength)
	}

	if len(xs)-1 != len(easingFns) {
		return 0, fmt.Errorf("wrong easing functions length: xs=%d, easingFns=%d (expected %d) %w", len(xs), len(easingFns), len(xs)-1, ErrWrongEasingFnsLength)
	}

	if len(xs) < 2 {
		return 0, ErrTooFewPoints
	}

	for i := 0; i < len(xs)-1; i++ {
		if x <= xs[i+1] {
			return EaseBetween(xs[i], xs[i+1], ys[i], ys[i+1], x, easingFns[i]), nil
		}
	}

	return ys[len(ys)-1], nil
}

func PieceWiseEaseInterpolateWithSingleEasingFn(
	xs []float64,
	ys []float64,
	easingFn func(float64) float64,
	x float64,
) (float64, error) {
	if len(xs) != len(ys) {
		return 0, fmt.Errorf("argument lists not same length: xs=%d, ys=%d %w", len(xs), len(ys), ErrArgumentsNotSameLength)
	}

	if len(xs) < 2 {
		return 0, ErrTooFewPoints
	}

	for i := 0; i < len(xs)-1; i++ {
		if x <= xs[i+1] {
			return EaseBetween(xs[i], xs[i+1], ys[i], ys[i+1], x, easingFn), nil
		}
	}

	return ys[len(ys)-1], nil
}
