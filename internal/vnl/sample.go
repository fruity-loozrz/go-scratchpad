package vnl

// Sample is a low-level scratchpad controls state
// Represents the immediate values after all interpolations
type Sample struct {
	// Pos represents the playhead position on the sample in seconds
	Pos float64
	// Vol represents the volume in the range [0.0, 1.0] (usually controlled by crossfader)
	Vol float64
}
