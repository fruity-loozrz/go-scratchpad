package microfmt

import (
	"errors"
	"math"
	"testing"

	"github.com/fruity-loozrz/go-scratchpad/internal/vnl"
)

func TestPatternToKeyframes(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		rangeMax float64
		want     []vnl.Keyframe
		wantErr  bool
	}{
		{
			name:     "example from comment: _-_-",
			pattern:  "_-_-",
			rangeMax: 1.0,
			want: []vnl.Keyframe{
				{Pos: 0.0, Value: 0},
				{Pos: 0.3333333333333333, Value: 1},
				{Pos: 0.6666666666666666, Value: 0},
				{Pos: 1.0, Value: 1},
			},
			wantErr: false,
		},
		{
			name:     "simple pattern __",
			pattern:  "__",
			rangeMax: 1.0,
			want: []vnl.Keyframe{
				{Pos: 0.0, Value: 0},
				{Pos: 1.0, Value: 0},
			},
			wantErr: false,
		},
		{
			name:     "simple pattern --",
			pattern:  "--",
			rangeMax: 1.0,
			want: []vnl.Keyframe{
				{Pos: 0.0, Value: 1},
				{Pos: 1.0, Value: 1},
			},
			wantErr: false,
		},
		{
			name:     "alternating pattern with range 2.0",
			pattern:  "_-_",
			rangeMax: 2.0,
			want: []vnl.Keyframe{
				{Pos: 0.0, Value: 0},
				{Pos: 1.0, Value: 1},
				{Pos: 2.0, Value: 0},
			},
			wantErr: false,
		},
		{
			name:     "single character underscore",
			pattern:  "_",
			rangeMax: 1.0,
			want: []vnl.Keyframe{
				{Pos: 0.0, Value: 0},
				{Pos: 1.0, Value: 0},
			},
			wantErr: false,
		},
		{
			name:     "single character dash",
			pattern:  "-",
			rangeMax: 1.0,
			want: []vnl.Keyframe{
				{Pos: 0.0, Value: 1},
				{Pos: 1.0, Value: 1},
			},
			wantErr: false,
		},
		{
			name:     "single character with range 2.0",
			pattern:  "_",
			rangeMax: 2.0,
			want: []vnl.Keyframe{
				{Pos: 0.0, Value: 0},
				{Pos: 2.0, Value: 0},
			},
			wantErr: false,
		},
		{
			name:     "longer pattern",
			pattern:  "_-_-_",
			rangeMax: 1.0,
			want: []vnl.Keyframe{
				{Pos: 0.0, Value: 0},
				{Pos: 0.25, Value: 1},
				{Pos: 0.5, Value: 0},
				{Pos: 0.75, Value: 1},
				{Pos: 1.0, Value: 0},
			},
			wantErr: false,
		},
		{
			name:     "invalid character x",
			pattern:  "_x_",
			rangeMax: 1.0,
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "invalid character number",
			pattern:  "_1_",
			rangeMax: 1.0,
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "invalid character space",
			pattern:  "_ _",
			rangeMax: 1.0,
			want:     nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PatternToKeyframes(tt.pattern, tt.rangeMax, vnl.Linear)
			if (err != nil) != tt.wantErr {
				t.Errorf("PatternToKeyframes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if !errors.Is(err, ErrInvalidPatternCharacter) {
					t.Errorf("PatternToKeyframes() expected ErrInvalidPatternCharacter, got %v", err)
				}
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("PatternToKeyframes() got %d keyframes, want %d", len(got), len(tt.want))
				return
			}
			for i := range got {
				if !floatEquals(got[i].Pos, tt.want[i].Pos) {
					t.Errorf("PatternToKeyframes() keyframe[%d].Pos = %v, want %v", i, got[i].Pos, tt.want[i].Pos)
				}
				if got[i].Value != tt.want[i].Value {
					t.Errorf("PatternToKeyframes() keyframe[%d].Value = %v, want %v", i, got[i].Value, tt.want[i].Value)
				}
			}
		})
	}
}

func TestNewEnvelopeFromPattern(t *testing.T) {
	tests := []struct {
		name       string
		pattern    string
		rangeMax   float64
		easingType vnl.EasingType
		wantErr    bool
	}{
		{
			name:       "valid pattern with linear easing",
			pattern:    "_-_-",
			rangeMax:   1.0,
			easingType: vnl.Linear,
			wantErr:    false,
		},
		{
			name:       "valid pattern with InQuad easing",
			pattern:    "_-_",
			rangeMax:   2.0,
			easingType: vnl.InQuad,
			wantErr:    false,
		},
		{
			name:       "invalid pattern",
			pattern:    "_x_",
			rangeMax:   1.0,
			easingType: vnl.Linear,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEnvelopeFromPattern(tt.rangeMax, tt.pattern, tt.easingType)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEnvelopeFromPattern() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("NewEnvelopeFromPattern() returned nil envelope without error")
			}
		})
	}
}

func floatEquals(a, b float64) bool {
	const epsilon = 1e-9
	return math.Abs(a-b) < epsilon
}
