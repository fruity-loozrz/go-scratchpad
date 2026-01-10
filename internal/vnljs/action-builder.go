package vnljs

import (
	"fmt"

	"github.com/fruity-loozrz/go-scratchpad/internal/microfmt"
	"github.com/fruity-loozrz/go-scratchpad/internal/vnl"
)

type ActionBuilder struct {
	Action vnl.ScratchAction
	api    *Api
}

func (a *ActionBuilder) Dur(durationInBeats float64) *ActionBuilder {
	a.Action.DurationInBeats = durationInBeats
	return a
}

func (a *ActionBuilder) Easing(easing vnl.EasingType) *ActionBuilder {
	a.Action.Easing = easing
	return a
}

func (a *ActionBuilder) FaderPattern(pattern vnl.FaderPattern) *ActionBuilder {
	a.Action.FaderPattern = pattern
	return a
}

func (a *ActionBuilder) FaderEnvelope(env *vnl.SmoothEnvelope) *ActionBuilder {
	a.Action.FaderEnvelope = env
	return a
}

func (a *ActionBuilder) FaderMicro(microformat string, easing vnl.EasingType) *ActionBuilder {
	envelope, err := microfmt.NewEnvelopeFromPattern(1.0, microformat, easing)
	if err != nil {
		panic(fmt.Sprintf("FaderMicro: invalid pattern %q: %v", microformat, err))
	}
	a.Action.FaderEnvelope = envelope
	return a
}

func (a *ActionBuilder) Platter(start, end float64) *ActionBuilder {
	a.Action.PlatterStart = start
	a.Action.PlatterEnd = end
	return a
}

func (a *ActionBuilder) PlatterEnvelopInBeats(env HasEnvelopeValue) *ActionBuilder {
	e := env.Envelope()
	a.Action.Envelope = e
	// TODO: add a method for that
	a.Action.DurationInBeats = e.Keyframes()[len(e.Keyframes())-1].Pos
	return a
}
