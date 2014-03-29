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
	table      *quad
	static     *quad
	st         *sym.SymbolTable
	quadSwitch QuadSwitch
	labelStk   []string
	elseLblStk []string
	lblNext    bool
	elseLblNext bool
}

func NewGenerator(st *sym.SymbolTable) *Generator {
	table := NewQuad()
	static := NewQuad()
	labelStk := make([]string,0)
	elseLblStk := make([]string,0)
	return &Generator{table: table, static: static, st: st, quadSwitch: STATIC,labelStk: labelStk, elseLblStk: elseLblStk}
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
	g.labelStk = append(g.labelStk,lbl)
}

func (g *Generator) AddElseLabel(lbl string) {
	g.elseLblStk = append(g.elseLblStk,lbl)
}

func (g *Generator) LabelNextRow() {
	g.lblNext = true
}

func (g *Generator) ElseLblNextRow() {
	g.elseLblNext = true
}

func (g *Generator) AddRow(label, command, op1, op2, op3, comment string) error {
	if g.lblNext && g.elseLblNext {
		panic("ICODE: Trying to double label a row")
	}
	if g.lblNext {
		if label == "" {
			//fmt.Printf("labelStk: %#v\n",g.labelStk)
			label = g.labelStk[len(g.labelStk) - 1]
			g.labelStk = g.labelStk[:len(g.labelStk) - 1]
		} else {
			//handle back patching
		}
		g.lblNext = false
	}
	if g.elseLblNext {
		if label == "" {
			label = g.elseLblStk[len(g.elseLblStk) - 1]
			g.labelStk = g.elseLblStk[:len(g.elseLblStk) - 1]
		} else {
			//handle back patching
		}
		g.elseLblNext = false
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
