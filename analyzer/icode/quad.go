package icode

import (
	"fmt"
	"sort"
)

type Quad struct {
	rows    []*QuadRow
	curRow  int
	numRows int
	labels  map[string][]int
}

func NewQuad() *Quad {
	r := make([]*QuadRow, 0)
	l := make(map[string][]int)
	return &Quad{rows: r, labels: l, curRow: -1} //start at -1 to be idx of slice
}

func (q *Quad) Size() int {
	return q.numRows
}

func (q *Quad) GetRows() []*QuadRow {
	return q.rows
}

func (q *Quad) Print() {
	fmt.Printf("Num Rows: %d curRow: %d\n", q.numRows, q.curRow)
	fmt.Println("Lables:")

	keys := make([]string, 0, len(q.labels))
	for k, _ := range q.labels {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s: %#v\n", k, q.labels[k])
	}
	fmt.Println("Rows:")
	for _, row := range q.rows {
		row.Print()
	}
}

func (q *Quad) ReplaceLabel(from, to string) {
	for _, v := range q.labels[from] {
		row := q.rows[v]
		var replaced bool
		if row.label == from {
			row.label = to
			replaced = true
		}
		if row.op1 == from {
			row.op1 = to
			replaced = true
		}
		if row.op2 == from {
			row.op2 = to
			replaced = true
		}
		if row.op3 == from {
			row.op3 = to
			replaced = true
		}
		if replaced {
			lines := q.labels[to]
			q.labels[to] = append(lines, v)
		}
		q.rows[v] = row
	}
}

func (q *Quad) AddQuadRow(label, command, op1, op2, op3, comment string) error {
	q.numRows++
	q.curRow++
	if label != "" {
		lines := q.labels[label]
		q.labels[label] = append(lines, q.curRow)
	}
	if op1 != "" {
		lines := q.labels[op1]
		q.labels[op1] = append(lines, q.curRow)
	}
	if op2 != "" {
		lines := q.labels[op2]
		q.labels[op2] = append(lines, q.curRow)
	}
	if op3 != "" {
		lines := q.labels[op3]
		q.labels[op3] = append(lines, q.curRow)
	}
	r := &QuadRow{label: label, command: command, op1: op1, op2: op2, op3: op3, comment: comment}
	q.rows = append(q.rows, r)
	return nil
}

type QuadRow struct {
	label   string
	command string
	op1     string
	op2     string
	op3     string
	comment string
}

func NewQuadRow(lbl, cmd, op1, op2, op3, comment string) *QuadRow {
	return &QuadRow{label: lbl, command: cmd, op1: op1, op2: op2, op3: op3, comment: comment}
}

func (q *QuadRow) GetLabel() string {
	return q.label
}

func (q *QuadRow) GetCommand() string {
	return q.command
}

func (q *QuadRow) GetOp1() string {
	return q.op1
}

func (q *QuadRow) GetOp2() string {
	return q.op2
}

func (q *QuadRow) GetOp3() string {
	return q.op3
}

func (q *QuadRow) GetComment() string {
	return q.comment
}

func (q *QuadRow) Print() {
	if q.label != "" {
		fmt.Printf("%s: ", q.label)
	}
	fmt.Printf("%s %s", q.command, q.op1)
	if q.op2 != "" {
		fmt.Printf(", %s", q.op2)
	}
	if q.op3 != "" {
		fmt.Printf(", %s", q.op3)
	}
	fmt.Printf(" ;%s\n", q.comment)
}
