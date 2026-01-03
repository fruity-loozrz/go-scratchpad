package play

import (
	"fmt"
	"log"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/fruity-loozrz/go-scratchpad/internal/scratch"
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

func runPlay(wavFileName, automationFileName string) error {
	scr := scratch.NewScratch()
	defer scr.Close()
	if err := scr.SetWavFileName(wavFileName); err != nil {
		return err
	}
	if err := scr.SetAutomationFileName(automationFileName); err != nil {
		return err
	}
	if err := scr.Init(); err != nil {
		return err
	}

	op := &oto.NewContextOptions{
		SampleRate:   int(scr.SampleRate()),
		ChannelCount: scr.NumChannels(),
		Format:       oto.FormatFloat32LE,
	}

	ctx, readyChan, err := oto.NewContext(op)
	if err != nil {
		return fmt.Errorf("failed to create audio context: %w", err)
	}
	<-readyChan

	player := ctx.NewPlayer(scr)
	player.Play()

	for {
		if !player.IsPlaying() {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
