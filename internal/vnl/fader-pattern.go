package vnl

type FaderPattern string

const (
	// Open: Fader is open for the entire duration.
	// Used for: Baby scratch, Drags, Releases.
	PatternOpen FaderPattern = "open"

	// Closed: Fader is closed (silence).
	// Used for: Pauses, ghost scratching.
	PatternClosed FaderPattern = "closed"

	// Cut: Starts open, closes briefly at the end (or reverse).
	// Used for: Chirps (cutting the sound at the turnaround point).
	PatternCut FaderPattern = "cut"

	// Transform: Rhythmic gating (on-off-on-off).
	// Used for: Transformer scratch. Usually requires a frequency param.
	PatternTransform FaderPattern = "transform"

	// Flare1: 1 click (cut) in the middle of the sound.
	// Result: 2 sounds from 1 movement.
	PatternFlare1 FaderPattern = "flare1"

	// Flare2: 2 clicks in the middle.
	// Result: 3 sounds from 1 movement.
	PatternFlare2 FaderPattern = "flare2"

	// Crab: Rapid 3-4 finger taps. Very fast sequence of cuts.
	PatternCrab FaderPattern = "crab"
)
