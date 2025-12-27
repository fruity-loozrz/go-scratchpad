package automationast

type Program struct {
	Lines []Stmt
}

type Stmt interface {
	isStmt()
}

type BPM struct {
	Value Number
}

func (BPM) isStmt() {}

type Move struct {
	DH Number
	DT *Number // nil => dt not specified
}

func (Move) isStmt() {}

type Number struct {
	Kind  string // "int" | "float" | "frac"
	Int   int64
	Float float64
	Num   int64 // numerator for frac
	Den   int64 // denominator for frac
}
