package keyframes

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/fruity-loozrz/go-scratchpad/internal/automation"
	"github.com/fruity-loozrz/go-scratchpad/internal/keyframes"
	"github.com/spf13/cobra"
)

var (
	automationFile string
	timeStep       float64
)

func NewKeyframesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keyframes",
		Short: "Convert automation file to CSV keyframes",
		Long:  `Convert an automation file to CSV format, either as raw keyframes or interpolated at regular intervals.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := runKeyframes(automationFile, timeStep); err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVarP(&automationFile, "automation", "a", "", "automation file (required)")
	cmd.Flags().Float64VarP(&timeStep, "timestep", "t", 0, "time step for interpolation in seconds (optional)")
	cmd.MarkFlagRequired("automation")

	return cmd
}

func runKeyframes(automationFile string, timeStep float64) error {
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

	// Create CSV writer
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	// Write header
	if err := writer.Write([]string{"Time", "Value"}); err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}

	// Check if we need interpolation
	if timeStep > 0 {
		// Create keyframe sequence for interpolation
		kfSeq, err := keyframes.NewKeyframeSequence(kfs)
		if err != nil {
			return fmt.Errorf("failed to create keyframe sequence: %w", err)
		}

		// Get duration in seconds
		duration := kfSeq.Duration().Seconds()

		// Interpolate and output
		for t := 0.0; t <= duration; t += timeStep {
			value := kfSeq.ValueAtTime(t)
			if err := writer.Write([]string{
				fmt.Sprintf("%.6f", t),
				fmt.Sprintf("%.6f", value),
			}); err != nil {
				return fmt.Errorf("failed to write CSV row: %w", err)
			}
		}
	} else {
		// Output raw keyframes
		for _, kf := range kfs {
			if err := writer.Write([]string{
				fmt.Sprintf("%.6f", kf.Time),
				fmt.Sprintf("%.6f", kf.Value),
			}); err != nil {
				return fmt.Errorf("failed to write CSV row: %w", err)
			}
		}
	}

	return nil
}
