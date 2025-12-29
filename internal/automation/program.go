package automation

import kf "github.com/fruity-loozrz/go-scratchpad/internal/keyframes"

type Program struct {
	Bpm   float64
	Moves []Move
}

func (p *Program) ToKeyframes() []kf.Keyframe {
	beatDuration := 60.0 / p.Bpm
	realTime := 0.0
	playHeadTime := 0.0

	keyframes := []kf.Keyframe{}
	for _, move := range p.Moves {
		// validate the values:
		//  - Dt must by greater than 0

		realTime += move.Dt * beatDuration
		playHeadTime += move.Dh * beatDuration
		keyframes = append(keyframes, kf.Keyframe{
			Time:  realTime,
			Value: playHeadTime,
		})
	}

	return keyframes
}
