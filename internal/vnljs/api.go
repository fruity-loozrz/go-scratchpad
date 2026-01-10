package vnljs

import (
	"math/rand"

	"github.com/fruity-loozrz/go-scratchpad/internal/microfmt"
	"github.com/fruity-loozrz/go-scratchpad/internal/vnl"
)

type Api struct {
	actions            []*vnl.ScratchAction
	SampleFile         string
	SampleOffset       float64
	BeatsPerMinute     float64
	RotationsPerMinute float64
	randSource         *rand.Rand
}

type SampleBuilder struct {
	api *Api
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

func (a *Api) Sample(sampleFile string) *SampleBuilder {
	a.SampleFile = sampleFile
	return &SampleBuilder{api: a}
}

func (s *SampleBuilder) Offset(offset float64) *SampleBuilder {
	s.api.SampleOffset = offset
	return s
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

func (a *Api) EnvelopeMicro(microFmt string, easingType vnl.EasingType) *vnl.SmoothEnvelope {
	env, err := microfmt.NewEnvelopeFromPattern(1.0, microFmt, easingType)
	if err != nil {
		panic(err)
	}
	return env
}
