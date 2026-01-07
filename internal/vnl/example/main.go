package main

import (
	"fmt"
	"strings"

	"github.com/fruity-loozrz/go-scratchpad/internal/vnl"
)

func main() {
	beatsPerMinute := 100.0
	rpm := 33.0

	// Pattern: Baby Scratch (1/8 notes)

	// 1. The record moves forward a 45-degree angle (1/8 of a rotation)
	action0 := vnl.ScratchAction{
		PlatterStart:    0,
		PlatterEnd:      1.0 / 8,
		DurationInBeats: 1.0,
		Easing:          vnl.EaseSmooth,  // Smooth acceleration and deceleration ("~" easing)
		FaderPattern:    vnl.PatternOpen, // Fader is open for the entire distance
	}

	// 2. The record moves backward to the starting point
	action1 := vnl.ScratchAction{
		PlatterStart:    1.0 / 8,
		PlatterEnd:      0,
		DurationInBeats: 1.0,
		Easing:          vnl.EaseSmooth, // The same smooth acceleration and deceleration
		FaderPattern:    vnl.PatternOpen,
	}

	seq, err := vnl.NewSequencerFromBpmRpm(
		[]vnl.ScratchAction{action0, action1},
		beatsPerMinute,
		rpm,
	)
	if err != nil {
		panic(err)
	}

	totalDuration := seq.GetTotalSequenceDurationInSeconds()

	viewResolutionInSeconds := 0.01

	for currentTime := 0.0; currentTime < totalDuration; currentTime += viewResolutionInSeconds {
		sample, err := seq.GetPositionAndGainAtTime(currentTime)
		if err != nil {
			panic(err)
		}

		barScale := 40.0

		posBar := strings.Repeat("^", int(sample.Pos*barScale))
		gainBar := strings.Repeat("|", int(sample.Vol*barScale))

		// fmt.Println(currentTime)
		fmt.Println(posBar)
		fmt.Println(gainBar)
		fmt.Println()
	}
}
