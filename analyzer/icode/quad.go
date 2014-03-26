package icode

import (
	"fmt"
)

type quad struct {
	rows    []*quadRow
	curRow  int
	numRows int
	labels  map[string][]int
}

func NewQuad() *quad {
	r := make([]*quadRow, 0)
	l := make(map[string][]int)
	return &quad{rows: r, labels: l, curRow: -1} //start at -1 to be idx of slice
}

func (q *quad) Print() {
	fmt.Printf("Num Rows: %d curRow: %d\n", q.numRows, q.curRow)
	fmt.Println("Lables:")
	for k, v := range q.labels {
		fmt.Printf("%s: %#v\n", k, v)
	}
	fmt.Println("Rows:")
	for _, row := range q.rows {
		row.Print()
	}
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

func (q *quadRow) Print() {
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
