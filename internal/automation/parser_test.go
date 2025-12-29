package automation_test

import (
	"testing"

	"github.com/fruity-loozrz/go-scratchpad/internal/automation"

	"github.com/stretchr/testify/require"
)

func TestParserGreenPath(t *testing.T) {
	require := require.New(t)

	program, err := automation.Parse(`
	bpm 120
	+
	-
	+1
	-1
	+1 1
	-1 1
	+ 2
	- 2
	+1 2
	-1 2
	+2 2
	-2 2
	+ 3/4
	- 3/4
	`)
	require.NoError(err)
	require.Equal(
		&automation.Program{
			Bpm: 120.0,
			Moves: []automation.Move{
				{Dh: 1, Dt: 1},
				{Dh: -1, Dt: 1},
				{Dh: 1, Dt: 1},
				{Dh: -1, Dt: 1},
				{Dh: 1, Dt: 1},
				{Dh: -1, Dt: 1},
				{Dh: 1, Dt: 2},
				{Dh: -1, Dt: 2},
				{Dh: 1, Dt: 2},
				{Dh: -1, Dt: 2},
				{Dh: 2, Dt: 2},
				{Dh: -2, Dt: 2},
				{Dh: 1, Dt: 0.75},
				{Dh: -1, Dt: 0.75},
			},
		}, program)
}

func TestParserUnknownToken(t *testing.T) {
	_, err := automation.Parse("foo")
	require := require.New(t)
	require.Error(err)
}

func TestParserDefaultBpm(t *testing.T) {
	program, err := automation.Parse("+")
	require := require.New(t)
	require.NoError(err)
	require.Equal(
		&automation.Program{
			Bpm:   140,
			Moves: []automation.Move{{Dh: 1, Dt: 1}},
		}, program)
}

func TestParserWithComments(t *testing.T) {
	program, err := automation.Parse(`
# this is a comment
+
+ # this is an inline comment
+ + # this another inline comment
# this is an ending comment
	`)
	require := require.New(t)
	require.NoError(err)
	require.Equal(
		&automation.Program{
			Bpm: 140,
			Moves: []automation.Move{
				{Dh: 1, Dt: 1},
				{Dh: 1, Dt: 1},
				{Dh: 1, Dt: 1},
			},
		}, program)
}
