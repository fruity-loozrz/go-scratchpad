package automation

import kf "github.com/fruity-loozrz/go-scratchpad/internal/keyframes"

type Program struct {
	Bpm       float64
	Predictor kf.PredictorFitter
	Moves     []Move
}

func (p *Program) SetInterpolationType(interpolationType interpolationType) {
	switch interpolationType {
	case interpolationTypeCubic:
		p.Predictor = &kf.PiecewiseCubicPredictor{}
	case interpolationTypeLinear:
		p.Predictor = &kf.PiecewiseLinearPredictor{}
	}
}

func (p *Program) ToKeyframes() []kf.Keyframe {
	beatDuration := 60.0 / p.Bpm
	realTime := 0.0
	playHeadTime := 0.0

	keyframes := []kf.Keyframe{}
	for _, move := range p.Moves {
		realTime += move.Dt * beatDuration
		playHeadTime += move.Dh * beatDuration
		keyframes = append(keyframes, kf.Keyframe{
			Time:  realTime,
			Value: playHeadTime,
		})
	}

	return keyframes
}
