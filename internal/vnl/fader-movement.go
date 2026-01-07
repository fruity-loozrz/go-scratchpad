package vnl

import (
	"fmt"

	"github.com/fogleman/ease"
	"github.com/fruity-loozrz/go-scratchpad/internal/math/interpolation"
)

type FaderMovement struct {
}

// GetGainAtMovementProgress returns the gain at a given movement progress
func (f *FaderMovement) GetGainAtMovementProgress(envelope FaderEnvelope, movementProgress float64) (float64, error) {
	envelopePositions := make([]float64, len(envelope))
	envelopeValues := make([]float64, len(envelope))
	for i, kf := range envelope {
		envelopePositions[i] = kf.Pos
		envelopeValues[i] = kf.Value
	}

	gain, err := interpolation.PieceWiseEaseInterpolateWithSingleEasingFn(
		envelopePositions,
		envelopeValues,
		ease.InOutQuint,
		movementProgress,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get gain at movement progress: %w", err)
	}
	return gain, nil
}
