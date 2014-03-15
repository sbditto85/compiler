package icode

type Quad struct {
	rows    []quadRow
	curRow  int
	numRows int
	labels  map[string][]int
}

func NewQuad() *Quad {
	r := make([]quadRow, 0)
	l := make(map[string][]int)
	return &Quad{rows: r, labels: l}
}

type quadRow struct {
	label   string
	command string
	op1     string
	op2     string
	op3     string
	comment string
}
