package main

import (
	"log"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/fruity-loozrz/go-scratchpad/internal/automation"
	"github.com/fruity-loozrz/go-scratchpad/internal/keyframes"
	"github.com/fruity-loozrz/go-scratchpad/internal/ring"
)

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

func main() {
	plate, err := createPlate("voice.wav")
	if err != nil {
		log.Fatal(err)
	}

	automationBytes, err := os.ReadFile("automation.txt")
	if err != nil {
		log.Fatal(err)
	}

	program, err := automation.Parse(string(automationBytes))
	if err != nil {
		log.Fatal(err)
	}

	kfs := program.ToKeyframes()

	kfInterpolator, err := keyframes.NewKeyframeSequence(kfs)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
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
}
