package main

import (
	"io"
	"math"
	"sync"
	"time"

	"github.com/ebitengine/oto/v3"
)

const (
	SampleRate = 44100
)

// PCM 16-bit LE mono ring reader
type pcm16LeRingBufferReader struct {
	data       []int16
	currentPos int

	stopped chan struct{}
	once    sync.Once
}

func newRingBufferReader(data []int16) *pcm16LeRingBufferReader {
	return &pcm16LeRingBufferReader{
		data:    data,
		stopped: make(chan struct{}),
	}
}

func (r *pcm16LeRingBufferReader) Read(buf []byte) (int, error) {
	// Oto will ask for byte buffers; for int16 LE we need even count
	if len(buf)%2 != 0 {
		return 0, io.ErrUnexpectedEOF
	}
	if len(r.data) == 0 {
		return 0, io.EOF
	}

	select {
	case <-r.stopped:
		return 0, io.EOF
	default:
	}

	// Fill the entire byte buffer with little-endian int16 samples
	for i := 0; i < len(buf); i += 2 {
		s := r.data[r.currentPos]
		buf[i] = byte(s)
		buf[i+1] = byte(s >> 8)

		r.currentPos++
		if r.currentPos >= len(r.data) {
			r.currentPos = 0
		}
	}

	return len(buf), nil
}

func (r *pcm16LeRingBufferReader) Stop() {
	r.once.Do(func() { close(r.stopped) })
}

// ----

func populateBufferWithSineWave(buffer []int16, sampleRate int, freqHz float64, volume float64) {
	amp := 32767.0 * volume
	for i := 0; i < len(buffer); i++ {
		buffer[i] = int16(math.Sin(2*math.Pi*float64(i)*freqHz/float64(sampleRate)) * amp)
	}
}

func main() {
	otoCtx, ready, err := oto.NewContext(&oto.NewContextOptions{
		SampleRate:   SampleRate,
		ChannelCount: 1,
		Format:       oto.FormatSignedInt16LE,
		BufferSize:   1024, // bytes; internal buffer in oto
	})
	if err != nil {
		panic(err)
	}
	<-ready

	// One cycle buffer (ring)
	ring := make([]int16, 44100)
	populateBufferWithSineWave(ring, SampleRate, 880.0, 0.1)

	reader := newRingBufferReader(ring)
	player := otoCtx.NewPlayer(reader)

	player.Play()
	time.Sleep(2 * time.Second)

	reader.Stop()
}
