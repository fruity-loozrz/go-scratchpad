package vnl

import (
	"fmt"
	"sort"

	"github.com/fruity-loozrz/go-scratchpad/internal/math/interpolation"
)

// SmoothEnvelope represents an envelope with smooth interpolation between keyframes
// using configurable easing functions for each interval.
type SmoothEnvelope struct {
	keyframes []Keyframe
	easings   []func(float64) float64
}

// NewSmoothEnvelop creates a new SmoothEnvelop with the given keyframes and easing functions.
// The easings slice should have length len(keyframes)-1, with each easing function
// corresponding to the interval between consecutive keyframes.
// If easings is nil or shorter than needed, Linear easing will be used for missing intervals.
func NewSmoothEnvelop(keyframes []Keyframe, easings []func(float64) float64) (*SmoothEnvelope, error) {
	if len(keyframes) < 2 {
		return nil, fmt.Errorf("SmoothEnvelop requires at least 2 keyframes, got %d", len(keyframes))
	}

	// Create a copy and sort by position
	kfs := make([]Keyframe, len(keyframes))
	copy(kfs, keyframes)
	sort.Slice(kfs, func(i, j int) bool {
		return kfs[i].Pos < kfs[j].Pos
	})

	// Prepare easing functions
	numIntervals := len(kfs) - 1
	easingFuncs := make([]func(float64) float64, numIntervals)

	for i := 0; i < numIntervals; i++ {
		if easings != nil && i < len(easings) && easings[i] != nil {
			easingFuncs[i] = easings[i]
		} else {
			easingFuncs[i] = GetEasingFunc(Linear)
		}
	}

	return &SmoothEnvelope{
		keyframes: kfs,
		easings:   easingFuncs,
	}, nil
}

// NewSmoothEnvelopFromTypes creates a SmoothEnvelop using EasingType constants.
func NewSmoothEnvelopFromTypes(keyframes []Keyframe, easingTypes []EasingType) (*SmoothEnvelope, error) {
	easings := make([]func(float64) float64, len(easingTypes))
	for i, et := range easingTypes {
		easings[i] = GetEasingFunc(et)
	}
	return NewSmoothEnvelop(keyframes, easings)
}

// ValueAt returns the interpolated value at the given position.
func (se *SmoothEnvelope) ValueAt(pos float64) float64 {
	// Find the interval containing pos
	for i := 0; i < len(se.keyframes)-1; i++ {
		startKf := se.keyframes[i]
		endKf := se.keyframes[i+1]

		if pos >= startKf.Pos && pos <= endKf.Pos {
			intervalWidth := endKf.Pos - startKf.Pos
			if intervalWidth == 0 {
				// Zero-length interval represents an instant transition
				return endKf.Value
			}

			return interpolation.EaseBetween(
				startKf.Pos, endKf.Pos,
				startKf.Value, endKf.Value,
				pos,
				se.easings[i],
			)
		}
	}

	// Position is outside all intervals, return last value as fallback
	return se.keyframes[len(se.keyframes)-1].Value
}

// Keyframes returns a copy of the keyframes.
func (se *SmoothEnvelope) Keyframes() []Keyframe {
	kfs := make([]Keyframe, len(se.keyframes))
	copy(kfs, se.keyframes)
	return kfs
}

// NumIntervals returns the number of intervals between keyframes.
func (se *SmoothEnvelope) NumIntervals() int {
	return len(se.keyframes) - 1
}
