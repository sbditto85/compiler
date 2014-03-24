package icode

import (
	sym "github.com/sbditto85/compiler/symbol_table"
)

type Generator struct {
	table *quad
	st    *sym.SymbolTable
	//stack for if/while labels
	//stack for else labels
}

func NewGenerator(st *sym.SymbolTable) *Generator {
	table := NewQuad()
	return &Generator{table: table, st: st}
}

func (g *Generator) AddRow(label, command, op1, op2, op3, comment string) error {
	return g.table.AddQuadRow(label, command, op1, op2, op3, comment)
}

func (g *Generator) PrintQuadTable() {
	g.table.Print()
}
