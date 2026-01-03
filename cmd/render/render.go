package render

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math"
	"os"

	"github.com/fruity-loozrz/go-scratchpad/internal/scratch"
	"github.com/spf13/cobra"
	"github.com/youpy/go-wav"
)

var (
	automationFile string
	outputFile     string
)

func NewRenderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "render [sound file]",
		Short: "Render a sound file with automation to WAV",
		Long:  `Render a sound file with automation from a specified automation file and save the output as a WAV file.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			soundFile := args[0]

			if err := runRender(soundFile, automationFile, outputFile); err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVarP(&automationFile, "automation", "a", "", "automation file (required)")
	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "output WAV file (required)")
	cmd.MarkFlagRequired("automation")
	cmd.MarkFlagRequired("output")

	return cmd
}

func runRender(wavFileName, automationFileName, outputFileName string) error {
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

	sampleRate := scr.SampleRate()
	numChannels := scr.NumChannels()

	// Read all audio data from the scratch buffer
	const bufferSize = 4096
	readBuffer := make([]byte, bufferSize)
	var allSamples []wav.Sample

	for {
		n, err := scr.Read(readBuffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read audio data: %w", err)
		}

		// Convert float32 LE bytes to 16-bit PCM samples
		samples := convertFloat32ToInt16(readBuffer[:n], numChannels)
		allSamples = append(allSamples, samples...)
	}

	// Create output file
	outFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	// Write WAV file
	numSamples := uint32(len(allSamples))
	writer := wav.NewWriter(outFile, numSamples, uint16(numChannels), sampleRate, 16)
	if err := writer.WriteSamples(allSamples); err != nil {
		return fmt.Errorf("failed to write samples: %w", err)
	}

	fmt.Printf("Rendered %d samples to %s (%.2f seconds)\n",
		numSamples, outputFileName, float64(numSamples)/float64(sampleRate))

	return nil
}

// convertFloat32ToInt16 converts a buffer of float32 LE samples to 16-bit PCM samples
func convertFloat32ToInt16(buffer []byte, numChannels int) []wav.Sample {
	const sizeofFloat32 = 4
	numSamples := len(buffer) / sizeofFloat32 / numChannels
	samples := make([]wav.Sample, numSamples)

	for i := range numSamples {
		for ch := range min(numChannels, 2) {
			offset := (i*numChannels + ch) * sizeofFloat32

			// Read float32 value (Little Endian)
			bits := binary.LittleEndian.Uint32(buffer[offset : offset+sizeofFloat32])
			floatValue := math.Float32frombits(bits)

			// Convert float32 [-1.0, 1.0] to 16-bit int [-32768, 32767]
			scaled := floatValue * 32767.0

			// Clamp to prevent overflow
			if scaled > 32767.0 {
				scaled = 32767.0
			} else if scaled < -32768.0 {
				scaled = -32768.0
			}

			samples[i].Values[ch] = int(scaled)
		}
	}

	return samples
}
