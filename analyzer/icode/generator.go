package icode

import (
	lex "github.com/sbditto85/compiler/lexer"
	sym "github.com/sbditto85/compiler/symbol_table"
)

type Generator struct {
	table *quad
	//stack for if/while labels
	//stack for else labels
}

func NewGenerator(st *sym.SymbolTable, l *lex.Lexer) *Generator {
	table := NewQuad()
	return &Generator{table: table}
}

func (g *Generator) AddRow(label, command, op1, op2, op3, comment string) error {
	return g.table.AddQuadRow(label, command, op1, op2, op3, comment)
}
