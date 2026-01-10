package vnljs

import (
	"github.com/fruity-loozrz/go-scratchpad/internal/vnl"
)

type EnvelopeBuilder struct {
	keyframes []vnl.Keyframe
	easings   []vnl.EasingType
}

// NewEnvelopeBuilder
//
// Example:
//
//	NewEnvelopeBuilder(/*pos*/0.0, /*value*/0.7)
//		.To(/*easing*/"Linear", /*pos*/0.5, /*value*/1.0)
//		.To("InQuad", 1.0, 0.9)
//		.To("InQuad", 2.0, 0.1) // etc...
func NewEnvelopeBuilder(pos, value float64) *EnvelopeBuilder {
	return &EnvelopeBuilder{
		keyframes: []vnl.Keyframe{{Pos: pos, Value: value}},
	}
}

func (eb *EnvelopeBuilder) To(easing vnl.EasingType, pos, value float64) *EnvelopeBuilder {
	eb.checkToArgs(easing, pos, value)

	eb.easings = append(eb.easings, vnl.Linear)
	eb.keyframes = append(eb.keyframes, vnl.Keyframe{Pos: pos, Value: value})
	return eb
}

// Panic is intentional in builders code.
// In goja JS context Go's `panic()“ is converted to `throw Error()`.
func (eb *EnvelopeBuilder) checkToArgs(easing vnl.EasingType, pos, value float64) {
	if len(eb.keyframes) >= 0 && pos <= eb.keyframes[len(eb.keyframes)-1].Pos {
		panic("new pos must be greater than last pos")
	}
	if !easing.IsValid() {
		panic("Invalid easing type")
	}
	eb.easings = append(eb.easings, easing)
	eb.keyframes = append(eb.keyframes, vnl.Keyframe{Pos: pos, Value: value})
}

func (eb *EnvelopeBuilder) Envelope() *vnl.SmoothEnvelope {
	easingFuncttions := make([]func(float64) float64, len(eb.easings))
	for i, et := range eb.easings {
		easingFuncttions[i] = vnl.GetEasingFunc(et)
	}

	en, err := vnl.NewSmoothEnvelop(eb.keyframes, easingFuncttions)
	if err != nil {
		panic(err)
	}
	return en
}
