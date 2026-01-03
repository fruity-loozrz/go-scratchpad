package scratch

import (
	"fmt"
	"io"
	"os"

	"github.com/fruity-loozrz/go-scratchpad/internal/automation"
	"github.com/fruity-loozrz/go-scratchpad/internal/keyframes"
	"github.com/fruity-loozrz/go-scratchpad/internal/ring"
)

type Scratch struct {
	*ring.Ring

	automationReader io.ReadCloser
	wavReader        ring.Reader
}

func NewScratch() *Scratch {
	return &Scratch{}
}

func (s *Scratch) SetAutomationReader(source io.ReadCloser) error {
	if s.automationReader != nil {
		return fmt.Errorf("automation source already set")
	}
	s.automationReader = source
	return nil
}

func (s *Scratch) SetWavReader(source ring.Reader) error {
	if s.wavReader != nil {
		return fmt.Errorf("wav source already set")
	}
	s.wavReader = source
	return nil
}

func (s *Scratch) SetAutomationFileName(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	return s.SetAutomationReader(f)
}

func (s *Scratch) SetWavFileName(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	return s.SetWavReader(f)
}

func (s *Scratch) Init() error {
	ring, err := ring.NewRingFromWav(s.wavReader)
	if err != nil {
		return fmt.Errorf("unable to create ring: %w", err)
	}
	s.Ring = ring

	automationString, err := io.ReadAll(s.automationReader)
	if err != nil {
		return fmt.Errorf("unable to read automation: %w", err)
	}

	program, err := automation.Parse(string(automationString))
	if err != nil {
		return fmt.Errorf("unable to parse automation: %w", err)
	}

	kfPoints := program.ToKeyframes()
	kfSequence, err := keyframes.NewKeyframeSequence(program.Predictor, kfPoints)
	if err != nil {
		return fmt.Errorf("failed to create keyframe sequence: %w", err)
	}

	ring.SetHeadPositionFn(
		func(f float64) float64 {
			return kfSequence.ValueAtTime(f)
		},
	)

	ring.SetDuration(kfSequence.Duration())

	return nil
}

func (s *Scratch) Close() error {
	if s.automationReader != nil {
		err := s.automationReader.Close()
		if err != nil {
			return err
		}
	}

	if s.wavReader != nil {
		err := s.wavReader.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
