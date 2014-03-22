package icode

type quad struct {
	rows    []*quadRow
	curRow  int
	numRows int
	labels  map[string][]int
}

func NewQuad() *quad {
	r := make([]*quadRow, 0)
	l := make(map[string][]int)
	return &quad{rows: r, labels: l}
}

func (q *quad) AddQuadRow(label, command, op1, op2, op3, comment string) error {
	if label != "" {
		lines := q.labels[label]
		q.labels[label] = append(lines, q.curRow)
	}
	r := &quadRow{label: label, command: command, op1: op1, op2: op2, op3: op3, comment: comment}
	q.rows = append(q.rows, r)
	q.numRows++
	q.curRow++
	return nil
}

type quadRow struct {
	label   string
	command string
	op1     string
	op2     string
	op3     string
	comment string
}
