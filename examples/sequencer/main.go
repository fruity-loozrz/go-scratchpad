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
	}

	// Create charts directory if it doesn't exist
	chartsDir := "examples/sequencer/charts"
	if err := os.MkdirAll(chartsDir, 0755); err != nil {
		panic(fmt.Errorf("failed to create charts directory: %w", err))
	}

	fmt.Printf("Generating %d scratch pattern examples...\n\n", len(examples))

	for _, ex := range examples {
		fmt.Printf("Generating %s: %s\n", ex.Number, ex.Name)
		if err := generateExample(ex, chartsDir); err != nil {
			fmt.Printf("  ERROR: %v\n", err)
		} else {
			fmt.Printf("  Success!\n")
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

func generateExample(ex Example, chartsDir string) error {
	seq, err := vnl.NewSequencerFromBpmRpm(ex.Actions, ex.BeatsPerMinute, ex.RPM)
	if err != nil {
		return fmt.Errorf("failed to create sequencer: %w", err)
	}

	times, positions, gains := sampleSequence(seq, resolution)
	filename := fmt.Sprintf("%s-%s.png", ex.Number, slugify(ex.Name))
	title := fmt.Sprintf("%s - %s", ex.Number, ex.Name)
	return drawChart(times, positions, gains, filename, title, chartsDir)
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

func drawChart(times, positions, gains []float64, filename, title, chartsDir string) error {
	xLabels := charts.FormatFloatLabels(times, "%.2f")
	seriesData := [][]float64{positions, gains}
	outputPath := fmt.Sprintf("%s/%s", chartsDir, filename)

	return charts.DrawLineChart(xLabels, seriesData, charts.LineChartOptions{
		Title:        title,
		OutputPath:   outputPath,
		LegendLabels: []string{"Position (revolutions)", "Gain (volume)"},
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
