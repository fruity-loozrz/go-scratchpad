package vnl

type ScratchAction struct {
	PlatterStart    float64      // Starting position of the platter in revolutions of the platter
	PlatterEnd      float64      // Ending position of the platter in revolutions of the platter
	DurationInBeats float64      // Duration of the action in beats
	Easing          EasingType   //
	FaderPattern    FaderPattern //
	Envelope        *SmoothEnvelope
}

// GetEnvelope returns the envelope of the action,
// where the Pos is the time in beats and the Value is the platter position
func (a *ScratchAction) GetEnvelope() *SmoothEnvelope {
	if a.Envelope != nil {
		return a.Envelope
	}

	a.Envelope, _ = NewSmoothEnvelopFromKeyframesAndEasings(
		[]Keyframe{
			{
				Pos:   0.0,
				Value: a.PlatterStart,
			},
			{
				Pos:   a.DurationInBeats,
				Value: a.PlatterEnd,
			},
		},
		[]EasingType{
			a.Easing,
		},
	)

	return a.Envelope
}
