package ring

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/go-audio/wav"
)

type Ring struct {
	buffer            []int
	sampleRate        int
	sampleDuration    float64
	totalPlaybackTime float64
	headPositionFn    func(float64) float64
	maxDuration       float64
}

func NewRingFromWav(file *os.File) (*Ring, error) {
	decoder := wav.NewDecoder(file)
	if !decoder.IsValidFile() {
		return nil, fmt.Errorf("invalid WAV file")
	}

	buf, err := decoder.FullPCMBuffer()
	if err != nil {
		return nil, err
	}

	numChannels := int(decoder.Format().NumChannels)
	bitDepth := int(decoder.SampleBitDepth())

	if numChannels != 1 {
		return nil, fmt.Errorf("error: audio must be mono (1 channel), got %d channels", numChannels)
	}
	if bitDepth != 16 {
		return nil, fmt.Errorf("error: audio must be 16-bit, got %d-bit", bitDepth)
	}

	sampleRate := int(decoder.Format().SampleRate)

	log.Printf("Channels: %d, Sample Rate: %d Hz, Bit Depth: %d\n", numChannels, sampleRate, bitDepth)
	log.Printf("Playing %d samples...\n", len(buf.Data))

	return &Ring{
		buffer:         buf.Data,
		sampleRate:     sampleRate,
		sampleDuration: 1.0 / float64(sampleRate),
	}, nil
}

func (r *Ring) timeToSample(t float64) int {
	bufLen := len(r.buffer)
	i := int(t*float64(r.sampleRate)) % bufLen
	if i < 0 {
		i += bufLen
	}
	return i
}

func (r *Ring) getSampleByTime(t float64) int {
	return r.buffer[r.timeToSample(t)]
}

func (r *Ring) Read(buf []byte) (int, error) {
	bufLen := len(buf)
	// TODO: unhardcode the 2
	for i := 0; i < bufLen; i += 2 {
		r.totalPlaybackTime += r.sampleDuration
		if r.totalPlaybackTime >= r.maxDuration {
			return i, io.EOF
		}
		headPositionTime := r.headPositionFn(r.totalPlaybackTime)
		sample := r.getSampleByTime(headPositionTime)
		buf[i] = byte(sample)
		buf[i+1] = byte(sample >> 8)
	}
	return bufLen, nil
}

func (r *Ring) SampleRate() int                            { return r.sampleRate }
func (r *Ring) NumChannels() int                           { return 1 }
func (r *Ring) BitDepth() int                              { return 16 }
func (r *Ring) TotalPlaybackTime() float64                 { return r.totalPlaybackTime }
func (r *Ring) SetHeadPositionFn(fn func(float64) float64) { r.headPositionFn = fn }
func (r *Ring) SetDuration(d time.Duration)                { r.maxDuration = float64(d) / float64(time.Second) }
