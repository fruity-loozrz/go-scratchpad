package keyframes

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/fruity-loozrz/go-scratchpad/internal/automation"
	"github.com/fruity-loozrz/go-scratchpad/internal/keyframes"
	"github.com/guptarohit/asciigraph"
	"github.com/spf13/cobra"
)

var (
	automationFile string
	timeStep       float64
	chart          bool
)

func NewKeyframesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keyframes",
		Short: "Convert automation file to CSV keyframes",
		Long:  `Convert an automation file to CSV format, either as raw keyframes or interpolated at regular intervals. Can also render as an ASCII chart.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runKeyframes(automationFile, timeStep, chart); err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVarP(&automationFile, "automation", "a", "", "automation file (required)")
	cmd.Flags().Float64VarP(&timeStep, "timestep", "t", 0, "time step for interpolation in seconds (optional)")
	cmd.Flags().BoolVarP(&chart, "chart", "c", false, "render as ASCII chart instead of CSV")
	cmd.MarkFlagRequired("automation")

	return cmd
}

func runKeyframes(automationFile string, timeStep float64, showChart bool) error {
	// Validate timestep if provided
	if timeStep < 0 {
		return fmt.Errorf("timestep must be positive")
	}

	// Read and parse automation file
	automationBytes, err := os.ReadFile(automationFile)
	if err != nil {
		return fmt.Errorf("failed to read automation file: %w", err)
	}

	program, err := automation.Parse(string(automationBytes))
	if err != nil {
		return fmt.Errorf("failed to parse automation: %w", err)
	}

	// Convert to keyframes
	kfs := program.ToKeyframes()

	// Collect data points
	var times []float64
	var values []float64

	// Check if we need interpolation
	if timeStep > 0 {
		// Create keyframe sequence for interpolation
		kfSeq, err := keyframes.NewKeyframeSequence(kfs)
		if err != nil {
			return fmt.Errorf("failed to create keyframe sequence: %w", err)
		}

		// Get duration in seconds
		duration := kfSeq.Duration().Seconds()

		// Interpolate and collect data
		for t := 0.0; t <= duration; t += timeStep {
			value := kfSeq.ValueAtTime(t)
			times = append(times, t)
			values = append(values, value)
		}
	} else {
		// Collect raw keyframes
		for _, kf := range kfs {
			times = append(times, kf.Time)
			values = append(values, kf.Value)
		}
	}

	// Output based on chart flag
	if showChart {
		// Calculate derivatives (rate of change)
		derivatives := make([]float64, len(values))
		for i := range values {
			if i == 0 {
				derivatives[i] = 0
			} else {
				dt := times[i] - times[i-1]
				dv := values[i] - values[i-1]
				if dt > 0 {
					derivatives[i] = dv / dt
				} else {
					derivatives[i] = 0
				}
			}
		}

		// Render ASCII chart with both value and derivative
		graph := asciigraph.PlotMany(
			[][]float64{values, derivatives},
			asciigraph.Height(20),
			asciigraph.Width(80),
			asciigraph.Caption("Keyframes: Value (green) and Derivative (red)"),
			asciigraph.SeriesColors(
				asciigraph.Green,
				asciigraph.Red,
			),
		)
		fmt.Println(graph)
	} else {
		// Output CSV
		writer := csv.NewWriter(os.Stdout)
		defer writer.Flush()

		// Write header
		if err := writer.Write([]string{"Time", "Value"}); err != nil {
			return fmt.Errorf("failed to write CSV header: %w", err)
		}

		// Write data
		for i := range times {
			if err := writer.Write([]string{
				fmt.Sprintf("%.6f", times[i]),
				fmt.Sprintf("%.6f", values[i]),
			}); err != nil {
				return fmt.Errorf("failed to write CSV row: %w", err)
			}
		}
	}

	return nil
}
