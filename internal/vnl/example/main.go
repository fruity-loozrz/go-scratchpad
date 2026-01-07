package main

import (
	"fmt"

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
		DurationInBeats: 1.0 / 8,         // The duration of the move - an eighth note
		Easing:          vnl.EaseSmooth,  // Smooth acceleration and deceleration ("~" easing)
		FaderPattern:    vnl.PatternOpen, // Fader is open for the entire distance
	}

	// 2. The record moves backward to the starting point
	action1 := vnl.ScratchAction{
		PlatterStart:    1.0 / 8,
		PlatterEnd:      0,
		DurationInBeats: 1.0 / 8,
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

	for currentTime := 0.0; currentTime < totalDuration; currentTime++ {
		sample, err := seq.GetPositionAndGainAtTime(currentTime)
		if err != nil {
			panic(err)
		}
		fmt.Println(sample.Pos, sample.Vol)
	}
}
