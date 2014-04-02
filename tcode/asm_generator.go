package tcode

import (
	ic "github.com/sbditto85/compiler/analyzer/icode"
	sym "github.com/sbditto85/compiler/symbol_table"
)

func GenerateASM(table *ic.Quad, st *sym.SymbolTable) (asm []string) {
	asm = make([]string, 0, table.Size())

	return
}
