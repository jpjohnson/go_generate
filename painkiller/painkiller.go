package painkiller

//go:generate stringer -type=Pill

// Pill - type
type Pill int

const (
	// Placebo - default
	Placebo Pill = iota
	// Aspirin -
	Aspirin
	// Ibuprofen -
	Ibuprofen
	// Paracetamol -
	Paracetamol
	// Acetaminohpen - is Paracetamol
	Acetaminohpen = Paracetamol
)
