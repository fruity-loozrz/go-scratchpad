package vnljs

import (
	"math/rand"

	"github.com/dop251/goja"
	"github.com/fruity-loozrz/go-scratchpad/internal/vnl"
)

type ActionApi struct {
	Action vnl.ScratchAction
	api    *Api
}

func (a *ActionApi) Dur(durationInBeats float64) *ActionApi {
	a.Action.DurationInBeats = durationInBeats
	return a
}

func (a *ActionApi) Easing(easing vnl.EasingType) *ActionApi {
	a.Action.Easing = easing
	return a
}

func (a *ActionApi) FaderPattern(pattern vnl.FaderPattern) *ActionApi {
	a.Action.FaderPattern = pattern
	return a
}

func (a *ActionApi) Platter(start, end float64) *ActionApi {
	a.Action.PlatterStart = start
	a.Action.PlatterEnd = end
	return a
}

//-----

type Api struct {
	actions            []*vnl.ScratchAction
	SampleFile         string
	BeatsPerMinute     float64
	RotationsPerMinute float64
	randSource         *rand.Rand
}

func (a *Api) Action() *ActionApi {
	actionApi := &ActionApi{api: a}
	a.actions = append(a.actions, &actionApi.Action)
	return actionApi
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

func ExecuteVnlJs(script string) (*Api, error) {
	api := &Api{
		BeatsPerMinute:     100.0,
		RotationsPerMinute: 33.0,
	}

	vm := goja.New()
	_ = vm.Set("api", api)
	_ = vm.Set("sample", api.Sample)
	_ = vm.Set("rpm", api.RPM)
	_ = vm.Set("bpm", api.BPM)
	_ = vm.Set("$", api.Action)
	_ = vm.Set("seed", api.Seed)
	_ = vm.Set("rand", api.Rand)

	_, err := vm.RunString(script)
	if err != nil {
		return nil, err
	}

	return api, nil
}
