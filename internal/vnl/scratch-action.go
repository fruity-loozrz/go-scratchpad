package vnl

type ScratchAction struct {
	PlatterStart    float64         // Starting position of the platter in revolutions of the platter
	PlatterEnd      float64         // Ending position of the platter in revolutions of the platter
	DurationInBeats float64         // Duration of the action in beats
	Easing          EasingType      //
	FaderPattern    FaderPattern    //
	FaderEnvelope   *SmoothEnvelope //
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

// GetFaderEnvelope returns the fader envelope of the action,
// where the Pos is the time in beats and the Value is the gain (0...1).
// Prefers FaderEnvelope if set, falls back to FaderPattern lookup.
func (a *ScratchAction) GetFaderEnvelope() *SmoothEnvelope {
	if a.FaderEnvelope != nil {
		return a.FaderEnvelope
	}

	// Backward compatibility: look up pattern in map
	faderEnv := FaderPatterns[a.FaderPattern]
	return &faderEnv
}
