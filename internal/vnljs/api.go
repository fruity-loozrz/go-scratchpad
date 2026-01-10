package vnljs

import (
	"math/rand"

	"github.com/fruity-loozrz/go-scratchpad/internal/vnl"
)

type Api struct {
	actions            []*vnl.ScratchAction
	SampleFile         string
	BeatsPerMinute     float64
	RotationsPerMinute float64
	randSource         *rand.Rand
}

func (a *Api) Action() *ActionBuilder {
	actionBuilder := &ActionBuilder{api: a}
	a.actions = append(a.actions, &actionBuilder.Action)
	return actionBuilder
}

func (a *Api) Actions() []vnl.ScratchAction {
	actions := make([]vnl.ScratchAction, len(a.actions))
	for i, action := range a.actions {
		actions[i] = *action
	}
	return actions
}

func (a *Api) BPM(bpm float64) *Api {
	a.BeatsPerMinute = bpm
	return a
}

func (a *Api) RPM(rpm float64) *Api {
	a.RotationsPerMinute = rpm
	return a
}

func (a *Api) Sample(sampleFile string) *Api {
	a.SampleFile = sampleFile
	return a
}

func (a *Api) Seed(seed int64) {
	a.randSource = rand.New(rand.NewSource(seed))
}

func (a *Api) Rand() float64 {
	if a.randSource == nil {
		a.Seed(0)
	}
	return a.randSource.Float64()
}

func (a *Api) Envelope(pos, value float64) *EnvelopeBuilder {
	return NewEnvelopeBuilder(pos, value)
}
