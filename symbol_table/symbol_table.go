package symbol_table

import (
	str "strings"
	"fmt"
	"strconv"
	"sort"
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
	symIds []string
}

func NewSymbolTable() *SymbolTable {
	e := make(map[string]SymbolTableElement)
	s := make([]string,0)
	return &SymbolTable{scope:"g",elems:e,symIds:s}
}

func (s *SymbolTable) GenSymId(kind string) string {
	s.symIdNum++
	symId := string(([]rune)(kind)[0]) + strconv.Itoa(s.symIdNum)
	s.symIds = append(s.symIds,symId)
	return symId
}

func (s *SymbolTable) AddElement(value string, kind string, data map[string]interface{}) (symId string) {

	curScope := s.scope

	switch kind {
	case "Class","Method","Main","Constructor":
		s.AddScope(value)
		//fmt.Printf("added scope to %s for value %s of kind %s\n",s.scope,value,kind)
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

	//fmt.Printf("scope %s for value %s for kind %s\n",curScope,value,kind)
	
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
	//fmt.Printf("Scope dropped down to %s now \n",s.scope)
	return nil
}

func (s *SymbolTable) GetScope() string {
	return s.scope
}

func (s *SymbolTable) PrintTable() {
	fmt.Printf("Current Scope: %s\n",s.scope)
	fmt.Println("=================")
	fmt.Println("Elements:")
	keys := make([]string,0,len(s.elems))
	for k := range(s.elems) {
		keys = append(keys,k)
	}
	sort.Strings(keys)
	for _,key := range(keys) {
		e := s.elems[key]
		e.PrintElement()
		fmt.Println("--------------")
	}
}

func (s *SymbolTable) PrintTableInAddOrder() {
	fmt.Printf("Current Scope: %s\n",s.scope)
	fmt.Println("=================")
	fmt.Println("Elements:")
	for _,key := range(s.symIds) {
		e := s.elems[key]
		e.PrintElement()
		fmt.Println("--------------")
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
