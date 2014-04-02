package icode

import (
	"fmt"
	sym "github.com/sbditto85/compiler/symbol_table"
)

type QuadSwitch int

const (
	MAIN QuadSwitch = iota
	STATIC
)

type Generator struct {
	table       *Quad
	static      *Quad
	st          *sym.SymbolTable
	quadSwitch  QuadSwitch
	labelStk    []string
	elseLblStk  []string
	lblNext     int
	elseLblNext int
}

func NewGenerator(st *sym.SymbolTable) *Generator {
	table := NewQuad()
	static := NewQuad()
	labelStk := make([]string, 0)
	elseLblStk := make([]string, 0)
	return &Generator{table: table, static: static, st: st, quadSwitch: STATIC, labelStk: labelStk, elseLblStk: elseLblStk}
}

func (g *Generator) GetQuad() *Quad {
	return g.table
}

func (g *Generator) SwitchToMain() {
	g.quadSwitch = MAIN
}

func (g *Generator) SwitchToStatic() {
	g.quadSwitch = STATIC
}

func (g *Generator) AddAndResetStatic() {
	for _, qr := range g.static.rows {
		g.AddRow(qr.label, qr.command, qr.op1, qr.op2, qr.op3, qr.comment)
	}

	g.static = NewQuad()
}

func (g *Generator) PrintSwitch() {
	switch g.quadSwitch {
	case MAIN:
		fmt.Print("MAIN")
	case STATIC:
		fmt.Print("STATIC")
	}
}

func (g *Generator) AddLabel(lbl string) {
	g.labelStk = append(g.labelStk, lbl)
}

func (g *Generator) AddElseLabel(lbl string) {
	g.elseLblStk = append(g.elseLblStk, lbl)
}

func (g *Generator) LabelNextRow() {
	g.lblNext += 1
}

func (g *Generator) ElseLblNextRow() {
	g.elseLblNext += 1
	//fmt.Printf("LABELING NEXT ROW: %d %#v\n",g.elseLblNext, g.elseLblStk)
}

func (g *Generator) AddRow(label, command, op1, op2, op3, comment string) error {
	/*if g.lblNext > 0  && g.elseLblNext > 0 {
		fmt.Printf("labelStk: %#v\n",g.labelStk)
		fmt.Printf("elseLblStk: %#v\n",g.elseLblStk)
		panic("ICODE: Trying to double label a row")
	}*/
	for ; g.lblNext > 0; g.lblNext-- {
		if label == "" {
			label = g.labelStk[len(g.labelStk)-1]
			g.labelStk = g.labelStk[:len(g.labelStk)-1]
		} else {
			//handle back patching
			replaceLabel := g.labelStk[len(g.labelStk)-1]
			g.labelStk = g.labelStk[:len(g.labelStk)-1]

			//fmt.Printf("REPLACE: %s WITH %s\n",replaceLabel,label)

			g.table.ReplaceLabel(replaceLabel, label)
			g.static.ReplaceLabel(replaceLabel, label)
		}
	}
	for ; g.elseLblNext > 0; g.elseLblNext-- {
		if label == "" {
			label = g.elseLblStk[len(g.elseLblStk)-1]
			g.elseLblStk = g.elseLblStk[:len(g.elseLblStk)-1]
		} else {
			//handle back patching
			replaceLabel := g.elseLblStk[len(g.elseLblStk)-1]
			g.elseLblStk = g.elseLblStk[:len(g.elseLblStk)-1]

			//fmt.Printf("REPLACE: %s WITH %s\n",replaceLabel,label)

			g.table.ReplaceLabel(replaceLabel, label)
			g.static.ReplaceLabel(replaceLabel, label)
		}
	}
	switch g.quadSwitch {
	case MAIN:
		return g.table.AddQuadRow(label, command, op1, op2, op3, comment)
	case STATIC:
		return g.static.AddQuadRow(label, command, op1, op2, op3, comment)
	}
	return fmt.Errorf("Could not write to quad")
}

func (g *Generator) PrintQuadTable() {
	g.table.Print()
}

func (g *Generator) PrintQuadStatic() {
	g.static.Print()
}
