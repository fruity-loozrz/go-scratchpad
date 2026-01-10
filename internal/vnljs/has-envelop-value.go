package vnljs

import "github.com/fruity-loozrz/go-scratchpad/internal/vnl"

type HasEnvelopeValue interface {
	Envelope() *vnl.SmoothEnvelope
}
