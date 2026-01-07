package scratch

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fruity-loozrz/go-scratchpad/internal/ring"
	"github.com/fruity-loozrz/go-scratchpad/internal/vnl"
	"github.com/fruity-loozrz/go-scratchpad/internal/vnljs"
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
	automationString, err := io.ReadAll(s.automationReader)
	if err != nil {
		return fmt.Errorf("unable to read automation: %w", err)
	}

	vnlScriptApi, err := vnljs.ExecuteVnlJs(string(automationString))
	if err != nil {
		return fmt.Errorf("unable to execute vnl script: %w", err)
	}

	bpm := vnlScriptApi.BeatsPerMinute
	rpm := vnlScriptApi.RotationsPerMinute
	wavFile := vnlScriptApi.SampleFile
	if err := s.SetWavFileName(wavFile); err != nil {
		return err
	}

	seqr, err := vnl.NewSequencerFromBpmRpm(vnlScriptApi.Actions(), bpm, rpm)
	if err != nil {
		return fmt.Errorf("unable to create sequencer: %w", err)
	}

	ring, err := ring.NewRingFromWav(s.wavReader)
	if err != nil {
		return fmt.Errorf("unable to create ring: %w", err)
	}
	s.Ring = ring

	ring.SetPositionAndGainFn(
		func(t float64) (float64, float64) {
			sample, _ := seqr.GetPositionAndGainAtTime(t)
			return sample.Pos, sample.Vol
		},
	)

	// TODO: simplify the code, use seconds, not time.Duration everywhere
	ring.SetDuration(time.Duration(seqr.GetTotalSequenceDurationInSeconds() * float64(time.Second)))

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
