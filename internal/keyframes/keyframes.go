package keyframes

import (
	"fmt"
	"sort"
	"time"

	"gonum.org/v1/gonum/interp"
)

type Keyframe struct {
	Time  float64
	Value float64
}

type KeyframeSequence struct {
	Keyframes []Keyframe
	predictor interp.PiecewiseCubic
}

func NewKeyframeSequence(keyframes []Keyframe) (*KeyframeSequence, error) {
	kfs := &KeyframeSequence{
		Keyframes: keyframes,
	}

	err := kfs.initialize()
	if err != nil {
		return nil, err
	}

	return kfs, nil
}

func (k *KeyframeSequence) ValueAtTime(t float64) float64 {
	return k.predictor.Predict(t)
}

func (k *KeyframeSequence) Duration() time.Duration {
	lastTime := k.Keyframes[len(k.Keyframes)-1].Time

	return time.Duration(lastTime * float64(time.Second))
}

func (k *KeyframeSequence) sortAndValidate() error {
	sort.Slice(k.Keyframes, func(i, j int) bool {
		return k.Keyframes[i].Time < k.Keyframes[j].Time
	})

	for i := 1; i < len(k.Keyframes); i++ {
		if k.Keyframes[i].Time == k.Keyframes[i-1].Time {
			return fmt.Errorf("duplicate keyframe time: %f", k.Keyframes[i].Time)
		}
	}

	return nil
}

func (k *KeyframeSequence) initialize() error {
	if err := k.sortAndValidate(); err != nil {
		return err
	}

	k.predictor = interp.PiecewiseCubic{}

	times := make([]float64, len(k.Keyframes))
	values := make([]float64, len(k.Keyframes))
	for i, kf := range k.Keyframes {
		times[i] = kf.Time
		values[i] = kf.Value
	}

	derivatives := make([]float64, len(k.Keyframes))

	for i := 0; i < len(times); i++ {
		if i == 0 {
			// First point: use forward difference
			derivatives[i] = (values[i+1] - values[i]) / (times[i+1] - times[i])
		} else if i == len(times)-1 {
			// Last point: use backward difference
			derivatives[i] = (values[i] - values[i-1]) / (times[i] - times[i-1])
		} else {
			// Middle points: use central difference
			derivatives[i] = (values[i+1] - values[i-1]) / (times[i+1] - times[i-1])
		}
	}

	k.predictor.FitWithDerivatives(times, values, derivatives)

	return nil
}
