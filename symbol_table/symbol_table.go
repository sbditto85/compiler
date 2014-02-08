package symbol_table

import (
	str "strings"
	"fmt"
	"strconv"
)

type Parameter struct {
	Typ string
	Identifier string
	IsArr bool
}

type SymbolTable struct {
	symIdNum int
	scope string
	elems map[string]SymbolTableElement
}

func NewSymbolTable() *SymbolTable {
	e := make(map[string]SymbolTableElement)
	return &SymbolTable{scope:"g",elems:e}
}

func (s *SymbolTable) GenSymId(kind string) string {
	s.symIdNum++
	return string(([]rune)(kind)[0]) + strconv.Itoa(s.symIdNum)
}

func (s *SymbolTable) AddElement(value string, kind string, data map[string]interface{}) (symId string) {

	curScope := s.scope

	switch kind {
	case "Class","Method","Main","Constructor":
		s.AddScope(value)
		if v,ok := data["parameters"]; ok {
			paramSymIds := make([]string,0) 
			if parameters,ok := v.([]Parameter); ok { 
				for _,p := range(parameters) {
					pmap := make(map[string]interface{})
					pmap["type"] = p.Typ
					pmap["isArray"] = p.IsArr
					paramSymId := s.AddElement(p.Identifier,"Parameter",pmap)
					paramSymIds = append(paramSymIds,paramSymId)
				}
			}
			data["paramSymIds"] = paramSymIds
		}
	}

	symId = s.GenSymId(kind)
	
	s.elems[symId] = SymbolTableElement{
		scope: curScope,
		symid: symId,
		value: value,
		kind: kind,
		data: data,
	}
	
	return symId
}

func (s *SymbolTable) AddScope(scope string) {
	s.scope += "." + scope
}

func (s *SymbolTable) DownScope() error {
	tmp := str.Split(s.scope,".")
	if len(tmp) < 1 {
		return fmt.Errorf("Can't drop scope, current scope is %s",s.scope)
	}
	tmp = tmp[:len(tmp)-1]
	s.scope = str.Join(tmp,".")
	return nil
}

func (s *SymbolTable) GetScope() string {
	return s.scope
}

func (s *SymbolTable) PrintTable() {
	fmt.Printf("Current Scope: %s\n",s.scope)
	fmt.Println("Elements:")
	for _,e := range(s.elems) {
		e.PrintElement()
	}
}

type SymbolTableElement struct {
	scope string
	symid string
	value string
	kind string
	data map[string]interface{}
}

func (s *SymbolTableElement) PrintElement() {
	fmt.Printf("Scope: %s, SymId: %s, Value: %s, Kind: %s\n",s.scope,s.symid,s.value,s.kind)
	fmt.Println("Extra Data:")
	for k,v := range(s.data) {
		fmt.Printf("Key: %s, Value: %v\n",k,v)
	}
}
