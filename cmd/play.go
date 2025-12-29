package cmd

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

var playCmd = &cobra.Command{
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

func init() {
	playCmd.Flags().StringVarP(&automationFile, "automation", "a", "", "automation file (required)")
	playCmd.MarkFlagRequired("automation")
}

func createPlate(fileName string) (*ring.Ring, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	plate, err := ring.NewRingFromWav(file)
	if err != nil {
		return nil, err
	}

	return plate, nil
}

func runPlay(soundFile, automationFile string) error {
	plate, err := createPlate(soundFile)
	if err != nil {
		return fmt.Errorf("failed to create plate: %w", err)
	}

	automationBytes, err := os.ReadFile(automationFile)
	if err != nil {
		return fmt.Errorf("failed to read automation file: %w", err)
	}

	program, err := automation.Parse(string(automationBytes))
	if err != nil {
		return fmt.Errorf("failed to parse automation: %w", err)
	}

	kfs := program.ToKeyframes()

	kfInterpolator, err := keyframes.NewKeyframeSequence(kfs)
	if err != nil {
		return fmt.Errorf("failed to create keyframe sequence: %w", err)
	}

	plate.SetHeadPositionFn(
		func(f float64) float64 {
			return kfInterpolator.ValueAtTime(f)
		},
	)

	plate.SetDuration(kfInterpolator.Duration())

	op := &oto.NewContextOptions{
		SampleRate:   plate.SampleRate(),
		ChannelCount: plate.NumChannels(),
		Format:       oto.FormatSignedInt16LE,
	}

	ctx, readyChan, err := oto.NewContext(op)
	if err != nil {
		return fmt.Errorf("failed to create audio context: %w", err)
	}
	<-readyChan

	player := ctx.NewPlayer(plate)
	player.Play()

	for {
		if !player.IsPlaying() {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
