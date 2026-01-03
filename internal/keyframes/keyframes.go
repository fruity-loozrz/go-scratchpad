package keyframes

import (
	"fmt"
	"sort"
	"time"
)

type KeyframeSequence struct {
	Keyframes []Keyframe
	predictor PredictorFitter
}

func NewKeyframeSequence(predictor PredictorFitter, keyframes []Keyframe) (*KeyframeSequence, error) {
	kfs := &KeyframeSequence{
		Keyframes: keyframes,
		predictor: predictor,
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

	times := make([]float64, len(k.Keyframes))
	values := make([]float64, len(k.Keyframes))
	for i, kf := range k.Keyframes {
		times[i] = kf.Time
		values[i] = kf.Value
	}

	k.predictor.Fit(times, values)

	return nil
}
