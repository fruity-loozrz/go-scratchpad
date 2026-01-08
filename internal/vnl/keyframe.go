package vnl

// Keyframe represents a point in time within an envelope.
type Keyframe struct {
	Pos   float64 // 0.0 to 1.0 (Percentage of the duration)
	Value float64 // 0.0 (Closed) or 1.0 (Open)
}
