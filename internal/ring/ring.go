package ring

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
	"time"

	"github.com/youpy/go-wav"
)

const SizeofFloat32 = 4

var ErrInvalidBufferSize = errors.New("invalid buffer size")

type Ring struct {
	buffers      [][]float64
	samplesCount uint32
	sampleRate   uint32
	numChannels  uint16

	realTime          float64
	positionAndGainFn func(float64) (float64, float64)
	maxDuration       float64
}

func NewRingFromWav(file Reader) (*Ring, error) {
	ring := &Ring{}

	if err := ring.initialize(file); err != nil {
		return nil, err
	}

	return ring, nil
}

func (r *Ring) initialize(file Reader) error {
	reader := wav.NewReader(file)
	format, err := reader.Format()
	if err != nil {
		return err
	}

	r.sampleRate = format.SampleRate
	r.numChannels = format.NumChannels

	for range r.numChannels {
		r.buffers = append(r.buffers, make([]float64, 0))
	}

	var samplesCount uint32
	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		for _, sample := range samples {
			for ch := range r.numChannels {
				r.buffers[ch] = append(r.buffers[ch], reader.FloatValue(sample, uint(ch)))
			}
			samplesCount++
		}
	}
	r.samplesCount = samplesCount

	r.positionAndGainFn = func(t float64) (float64, float64) { return t, 1.0 }

	return nil
}

func (r *Ring) getSampleAtTimeLinear(t float64, ch int) float64 {
	pos := t * float64(r.sampleRate)
	n := float64(r.samplesCount)

	pos = math.Mod(pos, n)
	if pos < 0 {
		pos += n
	}

	i0 := int(pos)
	i1 := i0 + 1
	if i1 >= int(r.samplesCount) {
		i1 = 0
	}

	frac := pos - float64(i0)
	s0 := r.buffers[ch][i0]
	s1 := r.buffers[ch][i1]
	return s0 + (s1-s0)*frac
}

func (r *Ring) Read(buf []byte) (int, error) {
	// reader MUST read float32 samples
	if len(buf)%SizeofFloat32 != 0 {
		return 0, ErrInvalidBufferSize
	}
	if r.realTime > r.maxDuration {
		return 0, io.EOF
	}

	numChannels := int(r.numChannels)

	bytesRequested := len(buf)
	samplesRequested := bytesRequested / SizeofFloat32 / numChannels
	bytesRead := 0

	for i := range samplesRequested {
		headTime, gain := r.positionAndGainFn(r.realTime)

		for currentChannel := 0; currentChannel < numChannels; currentChannel++ {
			sample := r.getSampleAtTimeLinear(headTime, currentChannel) * gain

			binary.LittleEndian.PutUint32(
				buf[(i*numChannels+currentChannel)*SizeofFloat32:],
				math.Float32bits(float32(sample)),
			)
			bytesRead += SizeofFloat32
		}

		r.realTime += 1.0 / float64(r.sampleRate)
		if r.realTime > r.maxDuration {
			return bytesRead, io.EOF
		}
	}

	return bytesRead, nil
}

// SetHeadPositionFn sets a function that returns the head position in seconds at a given time
func (r *Ring) SetPositionAndGainFn(fn func(float64) (float64, float64)) { r.positionAndGainFn = fn }
func (r *Ring) SetDuration(d time.Duration)                              { r.maxDuration = float64(d) / float64(time.Second) }
func (r *Ring) SampleRate() uint32                                       { return r.sampleRate }
func (r *Ring) NumChannels() int                                         { return int(r.numChannels) }
