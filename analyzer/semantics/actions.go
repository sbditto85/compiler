package semantics

import (
	"fmt"
	tok "github.com/sbditto85/compiler/token"
	sym "github.com/sbditto85/compiler/symbol_table"
)

func (s *SemanticManager) IPush(value, scope string, curTok *tok.Token) {
	s.sas.push(&Id_Sar{value:value, scope:scope, token:curTok})
}

func (s *SemanticManager) IExist(st *sym.SymbolTable) error {
	sar := s.sas.pop()
	if sar.Exists(st) {
		return nil
	}
	tok := sar.GetToken()
	return fmt.Errorf("%s does not exist on line %d\n",tok.Lexeme,tok.Linenum)
}

func (s *SemanticManager) OPush() {
	
}
