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
	table *quad
	static *quad
	st    *sym.SymbolTable
	quadSwitch QuadSwitch
	//stack for if/while labels
	//stack for else labels
}

func NewGenerator(st *sym.SymbolTable) *Generator {
	table := NewQuad()
	return &Generator{table: table, st: st}
}

func (g *Generator) SetQuadSwitch(to QuadSwitch) {
	g.quadSwitch = to
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
