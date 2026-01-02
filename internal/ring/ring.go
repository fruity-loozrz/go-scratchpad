package ring

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
	"os"
	"time"

	"github.com/youpy/go-wav"
)

const SizeofFloat32 = 4

var ErrInvalidBufferSize = errors.New("invalid buffer size")

type Ring struct {
	buffer     []float64
	sampleRate uint32

	realTime float64
	// head position in seconds
	headPositionFn func(float64) float64

	// buffer            []int
	// bytesPerSample    int
	// sampleDuration    float64
	maxDuration float64
}

func NewRingFromWav(file *os.File) (*Ring, error) {
	ring := &Ring{}

	reader := wav.NewReader(file)
	format, err := reader.Format()
	if err != nil {
		return nil, err
	}

	ring.sampleRate = format.SampleRate

	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		for _, sample := range samples {
			// TODO: support stereo
			fSample := reader.FloatValue(sample, 0)
			ring.buffer = append(ring.buffer, fSample)
		}
	}

	return ring, nil
}

func (r *Ring) getSampleAtTime(t float64) float64 {
	// TODO: support stereo
	// TODO: consider interpolation
	sampleNum := int(t * float64(r.sampleRate))
	for sampleNum < 0 {
		sampleNum += len(r.buffer)
	}
	sampleNum %= len(r.buffer)
	return r.buffer[sampleNum]
}

func (r *Ring) Read(buf []byte) (int, error) {
	// reader MUST read float32 samples
	if len(buf)%SizeofFloat32 != 0 {
		return 0, ErrInvalidBufferSize
	}
	if r.realTime > r.maxDuration {
		return 0, io.EOF
	}

	bytesRequested := len(buf)
	samplesRequested := bytesRequested / SizeofFloat32
	bytesRead := 0

	for i := 0; i < samplesRequested; i++ {
		headTime := r.headPositionFn(r.realTime)
		sample := r.getSampleAtTime(headTime)

		binary.LittleEndian.PutUint32(
			buf[i*SizeofFloat32:],
			math.Float32bits(float32(sample)),
		)
		bytesRead += SizeofFloat32

		r.realTime += 1.0 / float64(r.sampleRate)
		if r.realTime > r.maxDuration {
			return bytesRead, io.EOF
		}
	}

	return bytesRead, nil
}

func (r *Ring) SampleRate() uint32                         { return r.sampleRate }
func (r *Ring) NumChannels() int                           { return 1 }
func (r *Ring) SetHeadPositionFn(fn func(float64) float64) { r.headPositionFn = fn }
func (r *Ring) SetDuration(d time.Duration)                { r.maxDuration = float64(d) / float64(time.Second) }
