package play

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/fruity-loozrz/go-scratchpad/internal/automation"
	"github.com/fruity-loozrz/go-scratchpad/internal/keyframes"
	"github.com/fruity-loozrz/go-scratchpad/internal/ring"
	"github.com/spf13/cobra"
)

var automationFile string

func NewPlayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "play [sound file]",
		Short: "Play a sound file with automation",
		Long:  `Play a sound file with automation from a specified automation file.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			soundFile := args[0]

			if err := runPlay(soundFile, automationFile); err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVarP(&automationFile, "automation", "a", "", "automation file (required)")
	cmd.MarkFlagRequired("automation")

	return cmd
}

func createRing(fileName string) (*ring.Ring, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ring, err := ring.NewRingFromWav(file)
	if err != nil {
		return nil, err
	}

	return ring, nil
}

func runPlay(soundFile, automationFile string) error {
	ring, err := createRing(soundFile)
	if err != nil {
		return fmt.Errorf("failed to create ring: %w", err)
	}

	automationBytes, err := os.ReadFile(automationFile)
	if err != nil {
		return fmt.Errorf("failed to read automation file: %w", err)
	}

	program, err := automation.Parse(string(automationBytes))
	if err != nil {
		return fmt.Errorf("failed to parse automation: %w", err)
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

	op := &oto.NewContextOptions{
		SampleRate:   int(ring.SampleRate()),
		ChannelCount: ring.NumChannels(),
		Format:       oto.FormatFloat32LE,
	}

	ctx, readyChan, err := oto.NewContext(op)
	if err != nil {
		return fmt.Errorf("failed to create audio context: %w", err)
	}
	<-readyChan

	player := ctx.NewPlayer(ring)
	player.Play()

	for {
		if !player.IsPlaying() {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
