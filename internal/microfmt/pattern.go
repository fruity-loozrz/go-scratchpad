package microfmt

import (
	"errors"
	"fmt"

	"github.com/fruity-loozrz/go-scratchpad/internal/vnl"
)

const (
	char0 = '_'
	char1 = '-'
)

var ErrInvalidPatternCharacter = errors.New("invalid pattern character")

func NewEnvelopeFromPattern(
	rangeMax float64,
	pattern string,
	easingType vnl.EasingType,
) (*vnl.SmoothEnvelope, error) {
	kfs, err := PatternToKeyframes(pattern, rangeMax, easingType)
	if err != nil {
		return nil, err
	}

	easingFn := vnl.GetEasingFunc(easingType)

	easingFns := make([]func(float64) float64, len(kfs)-1)
	for i := range easingFns {
		easingFns[i] = easingFn
	}

	return vnl.NewSmoothEnvelop(kfs, easingFns)
}

func PatternToKeyframes(pattern string, rangeMax float64, easingType vnl.EasingType) ([]vnl.Keyframe, error) {
	kfs := []vnl.Keyframe{}

	// Special case: single character creates two keyframes at start and end with same value
	if len(pattern) == 1 {
		var value float64
		patternRune := rune(pattern[0])

		if patternRune == char0 {
			value = 0
		} else if patternRune == char1 {
			value = 1
		} else {
			return nil, fmt.Errorf("%w: %q", ErrInvalidPatternCharacter, pattern)
		}

		return []vnl.Keyframe{
			{Pos: 0, Value: value},
			{Pos: rangeMax, Value: value},
		}, nil
	}

	for patternIndex, patternRune := range pattern {
		var value float64

		if patternRune == char0 {
			value = 0
		} else if patternRune == char1 {
			value = 1
		} else {
			return nil, fmt.Errorf("%w: %q", ErrInvalidPatternCharacter, pattern)
		}

		pos := float64(patternIndex) / float64(len(pattern)-1) * rangeMax

		kfs = append(kfs, vnl.Keyframe{
			Pos:   pos,
			Value: value,
		})
	}

	return kfs, nil
}
/*

If input is:
pattern: "_-_-", range: 1, easing: Linear


The chart should look like:
"/\/\"

So the keyframes should be:
[
	[0, 0]
	[0.33, 1]
	[0.66, 0]
	[1, 1]
]

*/
