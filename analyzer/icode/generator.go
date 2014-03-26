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
	//stack for if/while labels
	//stack for else labels
}

func NewGenerator(st *sym.SymbolTable) *Generator {
	table := NewQuad()
	static := NewQuad()
	return &Generator{table: table, static: static, st: st, quadSwitch: STATIC}
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

func (g *Generator) AddRow(label, command, op1, op2, op3, comment string) error {
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
