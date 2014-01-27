package syntax

type SymbolTable struct {
	curId int
}

func NewSymbolTable() *SymbolTable {
	s := &SymbolTable{}
	return s
}

func (s *SymbolTable) AddSymbol() {}

func (s *SymbolTable) GetSymbol() {}

func (s *SymbolTable) GenId(str rune) string {
	t := s.curId
	s.curId++
	return string(str) + string(t)
}

func (s *SymbolTable) GetCurrScope() []string {
	return nil
}
