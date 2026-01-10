package vnljs

import (
	"github.com/dop251/goja"
)

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
	_ = vm.Set("from", api.Envelope)
	_ = vm.Set("micro", api.EnvelopeMicro)

	vm.SetFieldNameMapper(goja.UncapFieldNameMapper())

	_, err := vm.RunString(script)
	if err != nil {
		return nil, err
	}

	return api, nil
}
