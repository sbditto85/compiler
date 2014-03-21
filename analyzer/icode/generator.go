package icode

import (
	lex "github.com/sbditto85/compiler/lexer"
	sym "github.com/sbditto85/compiler/symbol_table"
)

type Generator struct {
	table *quad
}

func NewGenerator(st *sym.SymbolTable, l *lex.Lexer) *Generator {
	table := NewQuad()
	return &Generator{table: table}
}
