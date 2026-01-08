package vnl

import (
	"github.com/fogleman/ease"
	"github.com/fruity-loozrz/go-scratchpad/internal/math/interpolation"
)

type ActionWithAbsoluteTime struct {
	ScratchAction ScratchAction
	StartTime     float64
	EndTime       float64
}

type Sequencer struct {
	actions                            []ScratchAction
	actionsWithAbsoluteTime            []ActionWithAbsoluteTime
	beatDurationInSeconds              float64
	platterRevolutionDurationInSeconds float64
}

// Sequencer provides a pure function of time in seconds
// returning the time position on sample (vinyl position in seconds) and gain ([0..1])
// Use cases:
//
// Private:
//   - Given time, return ponter to action and linear progress (percentage in [0..1]) within the action itself
//
// Public:
//   - Given time, return position and gain
func NewSequencer(actions []ScratchAction, platterRevolutionDurationInSeconds, beatDurationInSeconds float64) (*Sequencer, error) {
	if platterRevolutionDurationInSeconds <= 0 {
		return nil, ErrSmallPlatterDuration
	}
	if beatDurationInSeconds <= 0 {
		return nil, ErrSmallBeatDuration
	}
	if len(actions) == 0 {
		return nil, ErrEmptyActions
	}
	return (&Sequencer{
		actions:                            actions,
		beatDurationInSeconds:              beatDurationInSeconds,
		platterRevolutionDurationInSeconds: platterRevolutionDurationInSeconds,
	}).initialize(), nil
}

func NewSequencerFromBpmRpm(actions []ScratchAction, bpm, rpm float64) (*Sequencer, error) {
	beatDurationInSeconds := 60 / bpm
	platterRevolutionDurationInSeconds := 60 / rpm
	return NewSequencer(actions, platterRevolutionDurationInSeconds, beatDurationInSeconds)
}

// initialize fills in a structure, where each action is given
// start and end time in seconds in the sequencer's coordinate system
func (s *Sequencer) initialize() *Sequencer {
	currentTime := 0.0

	s.actionsWithAbsoluteTime = make([]ActionWithAbsoluteTime, len(s.actions))

	for i, action := range s.actions {
		newAction := ActionWithAbsoluteTime{
			ScratchAction: action,
			StartTime:     currentTime,
			EndTime:       currentTime + s.beatDurationInSeconds*action.DurationInBeats,
		}

		s.actionsWithAbsoluteTime[i] = newAction
		currentTime = newAction.EndTime
	}

	return s
}

func (s *Sequencer) getActionAndProgressAtTime(timeInSeconds float64) (action *ActionWithAbsoluteTime, progressPercentageWithinAction float64, inRange bool) {
	if timeInSeconds < 0 {
		return &s.actionsWithAbsoluteTime[0], 0, false
	}

	for _, action := range s.actionsWithAbsoluteTime {
		if action.StartTime <= timeInSeconds && action.EndTime >= timeInSeconds {
			return &action, (timeInSeconds - action.StartTime) / (action.EndTime - action.StartTime), true
		}
	}

	return &s.actionsWithAbsoluteTime[len(s.actionsWithAbsoluteTime)-1], 1, false
}

func (s *Sequencer) getGainAtTime(timeInSeconds float64) (float64, error) {
	action, progress, _ := s.getActionAndProgressAtTime(timeInSeconds)

	faderEnvelope := FaderPatterns[action.ScratchAction.FaderPattern]

	faderPositions, faderValues := faderEnvelope.Unzip()
	return interpolation.PieceWiseEaseInterpolateWithSingleEasingFn(
		faderPositions,
		faderValues,
		ease.InOutQuint,
		progress,
	)
}

func (s *Sequencer) getPlatterPositionAtTimeInSeconds(timeInSeconds float64) float64 {
	action, progress, _ := s.getActionAndProgressAtTime(timeInSeconds)

	// TODO: the math here is disgusting.
	// Find better measurement units alignment for this, or incapulate it in the action

	progressInBeats := progress * action.ScratchAction.DurationInBeats
	platterPositionEnvelope := action.ScratchAction.GetEnvelope()
	platterPositionInRevoilutions := platterPositionEnvelope.ValueAt(progressInBeats)
	platterPositionInSeconds := platterPositionInRevoilutions * s.platterRevolutionDurationInSeconds
	return platterPositionInSeconds
}

func (s *Sequencer) GetTotalSequenceDurationInSeconds() float64 {
	return s.actionsWithAbsoluteTime[len(s.actionsWithAbsoluteTime)-1].EndTime
}

func (s *Sequencer) GetPositionAndGainAtTime(timeInSeconds float64) (sample Sample, err error) {
	sample.Pos = s.getPlatterPositionAtTimeInSeconds(timeInSeconds)
	sample.Vol, err = s.getGainAtTime(timeInSeconds)
	return
}
