package vnl

// Sample is a low-level scratchpad controls state
// Represents the immediate values after all interpolations
type Sample struct {
	// Pos represents the playhead position on the sample in seconds
	Pos float64
	// Vol represents the volume in the range [0.0, 1.0] (usually controlled by crossfader)
	Vol float64
}

// DSL Types

type EasingType string

const (
	// --- 1. Hand Mechanics ---

	// Linear: Robotic movement.
	// Used for: Transformer (slow drags), basic movement.
	// DSL symbol: "-"
	EaseLinear EasingType = "linear"

	// Smooth: Natural hand biomechanics.
	// Used for: Baby Scratch, simple back-and-forth motion.
	// Math: Sine InOut
	// DSL symbol: "~"
	EaseSmooth EasingType = "easeInOutSine"

	// Sharp: Aggressive, sharp movement.
	// Used for: Chirps, Stabs, Scribble. Fast acceleration/deceleration.
	// Math: Cubic InOut
	// DSL symbol: "^"
	EaseSharp EasingType = "easeInOutCubic"

	// Heavy: Feeling of heavy hand or high resistance.
	// Used for: Slow, "viscous" scratches.
	// Math: Quint InOut
	EaseHeavy EasingType = "easeInOutQuint"

	// --- 2. Platter Physics ---

	// Release: Hand releases the record, motor takes over.
	// Used for: Phrase endings, spin-ups.
	// Math: Expo Out (Sharp start, smooth stabilization)
	// DSL symbol: ">"
	EaseMotorStart EasingType = "easeOutExpo"

	// Spinback: Sharp throw backwards.
	// Used for: Backspin, rewinding.
	// Math: Cubic Out
	EaseSpinback EasingType = "easeOutCubic"

	// Stop: Simulates pressing the Stop button (Power Down).
	// Platter slows down to 0.
	// Math: Quad Out
	EasePowerDown EasingType = "easeOutQuad"

	// --- 3. Special FX ---

	// Bounce: Rebound effect.
	// Used for: Hitting a limiter or juggling simulation.
	// Math: EaseOutBounce
	EaseBounce EasingType = "easeOutBounce"

	// Elastic: Spring-like motion.
	// Used for: Glitch effects, digital scratches.
	// Math: EaseOutElastic
	EaseElastic EasingType = "easeOutElastic"
)

type Direction int

const (
	// Forward: Natural playback direction.
	// The sample plays "normally".
	DirFwd Direction = 1

	// Backward: Reverse playback.
	// The sample plays in reverse (rewind sound).
	DirBwd Direction = -1

	// Still: The record is held static (motor might be running or stopped).
	// Used for pauses or setting up the next cue point.
	DirStill Direction = 0
)

type FaderState int

const (
	// Open: Sound is audible (1.0).
	FaderOpen FaderState = 1

	// Closed: Sound is muted (0.0).
	FaderClosed FaderState = 0
)

type FaderCurve string

const (
	// SharpCut: Instant transition (binary).
	// Default for scratching (simulates minimal cut-in lag).
	// Uses a very fast ramp (e.g., 2-3ms) to avoid clicking.
	CurveSharp FaderCurve = "sharp"

	// Linear: Smooth transition (0.0 -> 1.0).
	// Used for fades, volume swells, or mixing tracks.
	CurveLinear FaderCurve = "linear"

	// ConstantPower: Maintains perceived loudness during the transition.
	// Standard for DJ mixers during blending.
	CurveConstantPower FaderCurve = "constantPower"
)

type ScratchAction struct {
	PlatterStart    float64      // Starting position of the platter in revolutions of the platter
	PlatterEnd      float64      // Ending position of the platter in revolutions of the platter
	DurationInBeats float64      // Duration of the action in beats
	Easing          EasingType   //
	FaderPattern    FaderPattern //
}

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
