package vnl

func newFaderPatternEnvelope(kfs []Keyframe) *SmoothEnvelope {
	env, _ := NewSmoothEnvelopFromKeyframesAndEasings(
		kfs,
		[]EasingType{InOutQuint},
	)

	return env
}

var FaderPatterns = map[FaderPattern]SmoothEnvelope{

	// --- Basics ---

	// PatternOpen: Always on.
	// Used for: Baby Scratch, Drag.
	PatternOpen: *newFaderPatternEnvelope([]Keyframe{{0.0, 1.0}, {1.0, 1.0}}),

	// PatternClosed: Always off.
	// Used for: Pauses, Ghosting.
	PatternClosed: *newFaderPatternEnvelope([]Keyframe{{0.0, 0.0}, {1.0, 0.0}}),

	// --- Cuts (Gating) ---

	// PatternCut (Chirp logic):
	// Starts Open, cuts sharply at the very end to hide the turnaround.
	// The "0" at 0.95 ensures silence before the direction change.
	PatternCut: *newFaderPatternEnvelope([]Keyframe{
		{0.0, 1.0},
		{0.9, 1.0},  // Stay open until 90%
		{0.95, 0.0}, // Quick cut close
		{1.0, 0.0},
	}),

	// PatternTransform (2 clicks):
	// Rhythmic gating: On - Off - On - Off.
	// Assumes a 50% duty cycle (equal sound and silence).
	PatternTransform: *newFaderPatternEnvelope([]Keyframe{
		{0.0, 1.0}, {0.25, 1.0}, // Sound 1
		{0.25, 0.0}, {0.5, 0.0}, // Silence 1
		{0.5, 1.0}, {0.75, 1.0}, // Sound 2
		{0.75, 0.0}, {1.0, 0.0}, // Silence 2
	}),

	// --- Flares (The Clicking math) ---

	// PatternFlare1 (1-Click Flare / Orbit):
	// The fader starts Open, clicks Closed briefly in the middle, ends Open.
	// Creates 2 sounds.
	// Note: The "Cut" (0.0) is very narrow (0.45 to 0.55).
	PatternFlare1: *newFaderPatternEnvelope([]Keyframe{
		{0.0, 1.0},
		{0.45, 1.0}, // Sound 1 ends
		{0.46, 0.0}, // Click (Silence) starts
		{0.54, 0.0}, // Click ends
		{0.55, 1.0}, // Sound 2 starts
		{1.0, 1.0},
	}),

	// PatternFlare2 (2-Click Flare):
	// Starts Open, clicks twice. Creates 3 sounds.
	// Cuts are at ~33% and ~66%.
	PatternFlare2: *newFaderPatternEnvelope([]Keyframe{
		{0.0, 1.0},
		// Sound 1
		{0.28, 1.0}, {0.30, 0.0}, // 1st Cut
		// Sound 2 (Middle)
		{0.36, 0.0}, {0.38, 1.0},
		{0.62, 1.0}, {0.64, 0.0}, // 2nd Cut
		// Sound 3
		{0.70, 0.0}, {0.72, 1.0},
		{1.0, 1.0},
	}),

	// PatternCrab (3-Click / 4 Sounds):
	// Extremely fast taps. The gaps are tiny.
	PatternCrab: *newFaderPatternEnvelope([]Keyframe{
		{0.0, 1.0},
		{0.20, 1.0}, {0.22, 0.0}, {0.25, 1.0}, // Cut 1
		{0.45, 1.0}, {0.47, 0.0}, {0.50, 1.0}, // Cut 2
		{0.70, 1.0}, {0.72, 0.0}, {0.75, 1.0}, // Cut 3
		{1.0, 1.0},
	}),
}
