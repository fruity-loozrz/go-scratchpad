package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fruity-loozrz/go-scratchpad/examples/charts"
	"github.com/fruity-loozrz/go-scratchpad/internal/vnl"
)

const resolution = 0.01

type Example struct {
	Number         string
	Name           string
	Description    string
	Actions        []vnl.ScratchAction
	BeatsPerMinute float64
	RPM            float64
}

func main() {
	examples := []Example{
		example001BabyScratch(),
		example002Transform(),
		example003Scribble(),
		example004Chirp(),
		example005Flare1(),
		example006Flare2(),
		example007Crab(),
		example008Tear(),
		example009Hydroplane(),
	}

	// Create charts directory if it doesn't exist
	chartsDir := "examples/sequencer/charts"
	if err := os.MkdirAll(chartsDir, 0755); err != nil {
		panic(fmt.Errorf("failed to create charts directory: %w", err))
	}

	fmt.Printf("Generating %d scratch pattern examples...\n\n", len(examples))

	for _, ex := range examples {
		fmt.Printf("Generating %s: %s (time & beats charts)\n", ex.Number, ex.Name)
		if err := generateExample(ex, chartsDir); err != nil {
			fmt.Printf("  ERROR: %v\n", err)
		} else {
			fmt.Printf("  Success! (2 charts)\n")
		}
	}

	fmt.Printf("\nAll charts generated in %s/\n", chartsDir)
}

func example001BabyScratch() Example {
	return Example{
		Number:         "001",
		Name:           "Baby Scratch",
		Description:    "Basic forward and backward motion with smooth easing",
		BeatsPerMinute: 100.0,
		RPM:            33.0,
		Actions: []vnl.ScratchAction{
			{
				PlatterStart:    0,
				PlatterEnd:      1.0 / 8,
				DurationInBeats: 1.0,
				Easing:          vnl.EaseSmooth,
				FaderPattern:    vnl.PatternOpen,
			},
			{
				PlatterStart:    1.0 / 8,
				PlatterEnd:      0,
				DurationInBeats: 1.0,
				Easing:          vnl.EaseSmooth,
				FaderPattern:    vnl.PatternOpen,
			},
		},
	}
}

func example002Transform() Example {
	return Example{
		Number:         "002",
		Name:           "Transform",
		Description:    "Forward motion with fader clicks on/off",
		BeatsPerMinute: 120.0,
		RPM:            33.0,
		Actions: []vnl.ScratchAction{
			{
				PlatterStart:    0,
				PlatterEnd:      1.0 / 4,
				DurationInBeats: 2.0,
				Easing:          vnl.EaseLinear,
				FaderPattern:    vnl.PatternTransform,
			},
		},
	}
}

func example003Scribble() Example {
	return Example{
		Number:         "003",
		Name:           "Scribble",
		Description:    "Rapid back-and-forth movements with linear easing",
		BeatsPerMinute: 140.0,
		RPM:            33.0,
		Actions: []vnl.ScratchAction{
			{
				PlatterStart:    0,
				PlatterEnd:      1.0 / 16,
				DurationInBeats: 0.25,
				Easing:          vnl.EaseLinear,
				FaderPattern:    vnl.PatternOpen,
			},
			{
				PlatterStart:    1.0 / 16,
				PlatterEnd:      0,
				DurationInBeats: 0.25,
				Easing:          vnl.EaseLinear,
				FaderPattern:    vnl.PatternOpen,
			},
			{
				PlatterStart:    0,
				PlatterEnd:      1.0 / 16,
				DurationInBeats: 0.25,
				Easing:          vnl.EaseLinear,
				FaderPattern:    vnl.PatternOpen,
			},
			{
				PlatterStart:    1.0 / 16,
				PlatterEnd:      0,
				DurationInBeats: 0.25,
				Easing:          vnl.EaseLinear,
				FaderPattern:    vnl.PatternOpen,
			},
		},
	}
}

func example004Chirp() Example {
	return Example{
		Number:         "004",
		Name:           "Chirp",
		Description:    "Forward scratch with fader cut pattern",
		BeatsPerMinute: 110.0,
		RPM:            33.0,
		Actions: []vnl.ScratchAction{
			{
				PlatterStart:    0,
				PlatterEnd:      1.0 / 8,
				DurationInBeats: 1.0,
				Easing:          vnl.EaseLinear,
				FaderPattern:    vnl.PatternCut,
			},
			{
				PlatterStart:    1.0 / 8,
				PlatterEnd:      0,
				DurationInBeats: 1.0,
				Easing:          vnl.EaseSmooth,
				FaderPattern:    vnl.PatternClosed,
			},
		},
	}
}

func example005Flare1() Example {
	return Example{
		Number:         "005",
		Name:           "1-Click Flare",
		Description:    "Single fader click during forward motion (Orbit)",
		BeatsPerMinute: 100.0,
		RPM:            33.0,
		Actions: []vnl.ScratchAction{
			{
				PlatterStart:    0,
				PlatterEnd:      1.0 / 8,
				DurationInBeats: 1.5,
				Easing:          vnl.EaseSmooth,
				FaderPattern:    vnl.PatternFlare1,
			},
			{
				PlatterStart:    1.0 / 8,
				PlatterEnd:      0,
				DurationInBeats: 1.5,
				Easing:          vnl.EaseSmooth,
				FaderPattern:    vnl.PatternOpen,
			},
		},
	}
}

func example006Flare2() Example {
	return Example{
		Number:         "006",
		Name:           "2-Click Flare",
		Description:    "Two fader clicks during forward motion",
		BeatsPerMinute: 95.0,
		RPM:            33.0,
		Actions: []vnl.ScratchAction{
			{
				PlatterStart:    0,
				PlatterEnd:      1.0 / 6,
				DurationInBeats: 2.0,
				Easing:          vnl.EaseSmooth,
				FaderPattern:    vnl.PatternFlare2,
			},
			{
				PlatterStart:    1.0 / 6,
				PlatterEnd:      0,
				DurationInBeats: 2.0,
				Easing:          vnl.EaseSmooth,
				FaderPattern:    vnl.PatternOpen,
			},
		},
	}
}

func example007Crab() Example {
	return Example{
		Number:         "007",
		Name:           "Crab",
		Description:    "Three rapid fader clicks (3-click technique)",
		BeatsPerMinute: 90.0,
		RPM:            33.0,
		Actions: []vnl.ScratchAction{
			{
				PlatterStart:    0,
				PlatterEnd:      1.0 / 8,
				DurationInBeats: 2.0,
				Easing:          vnl.EaseLinear,
				FaderPattern:    vnl.PatternCrab,
			},
			{
				PlatterStart:    1.0 / 8,
				PlatterEnd:      0,
				DurationInBeats: 2.0,
				Easing:          vnl.EaseLinear,
				FaderPattern:    vnl.PatternOpen,
			},
		},
	}
}

func example008Tear() Example {
	return Example{
		Number:         "008",
		Name:           "Tear",
		Description:    "Reverse scratch - sound only on the backward motion",
		BeatsPerMinute: 105.0,
		RPM:            33.0,
		Actions: []vnl.ScratchAction{
			{
				PlatterStart:    0,
				PlatterEnd:      1.0 / 8,
				DurationInBeats: 1.0,
				Easing:          vnl.EaseSmooth,
				FaderPattern:    vnl.PatternClosed,
			},
			{
				PlatterStart:    1.0 / 8,
				PlatterEnd:      0,
				DurationInBeats: 1.0,
				Easing:          vnl.EaseSmooth,
				FaderPattern:    vnl.PatternOpen,
			},
		},
	}
}

func example009Hydroplane() Example {
	return Example{
		Number:         "009",
		Name:           "Hydroplane",
		Description:    "Slow forward, fast backward with asymmetric timing",
		BeatsPerMinute: 115.0,
		RPM:            33.0,
		Actions: []vnl.ScratchAction{
			{
				PlatterStart:    0,
				PlatterEnd:      1.0 / 6,
				DurationInBeats: 2.0,
				Easing:          vnl.EaseLinear,
				FaderPattern:    vnl.PatternOpen,
			},
			{
				PlatterStart:    1.0 / 6,
				PlatterEnd:      0,
				DurationInBeats: 0.5,
				Easing:          vnl.EaseLinear,
				FaderPattern:    vnl.PatternOpen,
			},
		},
	}
}

func generateExample(ex Example, chartsDir string) error {
	seq, err := vnl.NewSequencerFromBpmRpm(ex.Actions, ex.BeatsPerMinute, ex.RPM)
	if err != nil {
		return fmt.Errorf("failed to create sequencer: %w", err)
	}

	beatDuration := 60.0 / ex.BeatsPerMinute
	platterRevDuration := 60.0 / ex.RPM

	times, positions, gains := sampleSequence(seq, resolution)

	// Generate time-based chart
	filenameTime := fmt.Sprintf("%s-%s-time.png", ex.Number, slugify(ex.Name))
	titleTime := fmt.Sprintf("%s - %s (Time)", ex.Number, ex.Name)
	if err := drawChart(times, positions, gains, filenameTime, titleTime, "time", chartsDir, beatDuration, platterRevDuration); err != nil {
		return err
	}

	// Generate beat-based chart
	filenameBeat := fmt.Sprintf("%s-%s-beats.png", ex.Number, slugify(ex.Name))
	titleBeat := fmt.Sprintf("%s - %s (Beats)", ex.Number, ex.Name)
	if err := drawChart(times, positions, gains, filenameBeat, titleBeat, "beats", chartsDir, beatDuration, platterRevDuration); err != nil {
		return err
	}

	return nil
}

func sampleSequence(seq *vnl.Sequencer, resolution float64) ([]float64, []float64, []float64) {
	totalDuration := seq.GetTotalSequenceDurationInSeconds()

	var times []float64
	var positions []float64
	var gains []float64

	for currentTime := 0.0; currentTime < totalDuration; currentTime += resolution {
		sample, err := seq.GetPositionAndGainAtTime(currentTime)
		if err != nil {
			panic(err)
		}

		times = append(times, currentTime)
		positions = append(positions, sample.Pos)
		gains = append(gains, sample.Vol)
	}

	return times, positions, gains
}

func drawChart(times, positions, gains []float64, filename, title, chartType, chartsDir string, beatDuration, platterRevDuration float64) error {
	var xLabels []string
	var yPositions, yGains []float64
	var legendPos, legendGain string

	if chartType == "time" {
		// Time-based chart (X = seconds, Y = seconds)
		xLabels = charts.FormatFloatLabels(times, "%.2f")
		yPositions = positions
		yGains = gains
		legendPos = "Position (seconds)"
		legendGain = "Gain (volume)"
	} else {
		// Beat-based chart (X = beats, Y = revolutions)
		beats := make([]float64, len(times))
		revolutions := make([]float64, len(positions))
		for i := range times {
			beats[i] = times[i] / beatDuration
			revolutions[i] = positions[i] / platterRevDuration
		}
		xLabels = charts.FormatFloatLabels(beats, "%.2f")
		yPositions = revolutions
		yGains = gains
		legendPos = "Position (revolutions)"
		legendGain = "Gain (volume)"
	}

	seriesData := [][]float64{yPositions, yGains}
	outputPath := fmt.Sprintf("%s/%s", chartsDir, filename)

	return charts.DrawLineChart(xLabels, seriesData, charts.LineChartOptions{
		Title:        title,
		OutputPath:   outputPath,
		LegendLabels: []string{legendPos, legendGain},
	})
}

func slugify(name string) string {
	var result strings.Builder
	for _, c := range name {
		if c >= 'A' && c <= 'Z' {
			result.WriteRune(c + 32)
		} else if c >= 'a' && c <= 'z' {
			result.WriteRune(c)
		} else if c >= '0' && c <= '9' {
			result.WriteRune(c)
		} else if c == ' ' {
			result.WriteRune('-')
		}
	}
	return result.String()
}
