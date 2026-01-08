package vnl

// Envelope is the sequence of states for one action.
type Envelope []Keyframe

func (fe Envelope) Unzip() (positions []float64, values []float64) {
	positions = make([]float64, len(fe))
	values = make([]float64, len(fe))
	for i, kf := range fe {
		positions[i] = kf.Pos
		values[i] = kf.Value
	}
	return
}
