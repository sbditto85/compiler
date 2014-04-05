package symbol_table

import (
	"fmt"
	"sort"
	"strconv"
	str "strings"
)

type Parameter struct {
	Typ        string
	Identifier string
	IsArr      bool
}

type SymbolTable struct {
	symIdNum      int
	scope         string
	elems         map[string]SymbolTableElement
	symIds        []string
	scopeElements map[string][]string
	funcType      string
}

func NewSymbolTable() *SymbolTable {
	e := make(map[string]SymbolTableElement)
	s := make([]string, 0)
	se := make(map[string][]string)
	return &SymbolTable{scope: "g", elems: e, symIds: s, scopeElements: se}
}

func (s *SymbolTable) GetFunctionType() string {
	return s.funcType
}

func (s *SymbolTable) GenSymId(kind string) string {
	if kind == "Main" {
		s.symIds = append(s.symIds, "MAIN")
		return "MAIN"
	}
	s.symIdNum++
	rs := ([]rune)(kind)
	symId := string(rs[0:2]) + strconv.Itoa(s.symIdNum)
	s.symIds = append(s.symIds, symId)
	return symId
}

func (s *SymbolTable) AddElement(value string, kind string, data map[string]interface{}, addToHash bool) (symId string) {

	curScope := s.scope

	switch kind {
	case "Ivar":
		tmp := str.Split(s.scope, ".")
		data["this_class"] = tmp[len(tmp)-1]
	case "Method", "Main", "Constructor":
		if _, ok := data["scope"]; ok {
			break
		}
		if typ, ok := data["type"]; ok {
			switch t := typ.(type) {
			case string:
				s.funcType = t
			default:
				s.funcType = kind
			}
		} else {
			s.funcType = kind
		}
		fallthrough
	case "Class":
		s.AddScope(value)
		//fmt.Printf("added scope to %s for value %s of kind %s\n",s.scope,value,kind)
		if v, ok := data["parameters"]; ok {
			paramSymIds := make([]string, 0)
			if parameters, ok := v.([]Parameter); ok {
				for _, p := range parameters {
					pmap := make(map[string]interface{})
					pmap["type"] = p.Typ
					pmap["isArray"] = p.IsArr
					paramSymId := s.AddElement(p.Identifier, "Parameter", pmap, addToHash)
					paramSymIds = append(paramSymIds, paramSymId)
				}
			}
			data["paramSymIds"] = paramSymIds
		}
	}

	if !addToHash {
		return ""
	}

	symId = s.GenSymId(kind)

	//fmt.Printf("scope %s for value %s for kind %s\n",curScope,value,kind)

	//check to see if we want to override the scope (used for types)
	if scp, ok := data["scope"]; ok {
		switch useScope := scp.(type) {
		case string:
			curScope = useScope
		}
	}

	s.scopeElements[curScope] = append(s.scopeElements[curScope], symId)

	s.elems[symId] = SymbolTableElement{
		Scope: curScope,
		SymId: symId,
		Value: value,
		Kind:  kind,
		Data:  data,
	}

	switch kind {
	case "Method", "Main", "Constructor":
		elem := s.elems[symId]
		offset, _ := IntFromData(data, "size")
		if params, err := StringSliceFromData(elem.Data, "paramSymIds"); err == nil {
			for _, paramSymId := range params {
				param, _ := s.GetElement(paramSymId)
				typ, _ := StringFromData(param.Data, "type")
				isArr, _ := BoolFromData(param.Data, "isArray")
				param.Data["offset"] = offset
				offset += SizeOfType(typ, isArr)
			}
		}
		data["size"] = offset
	case "Lvar", "Tvar":
		scopeCheck, method, _ := s.ScopeBelowWithCurr(curScope)
		methodElem := s.GetElementInScope(scopeCheck, method)
		offset, _ := IntFromData(methodElem.Data, "size")

		data["offset"] = offset

		typ, _ := StringFromData(data, "type")
		isArr, _ := BoolFromData(data, "isArray")
		if methodElem.Data != nil {
			methodElem.Data["size"] = offset + SizeOfType(typ, isArr)
		}
	case "Ivar":
		class_scope, _ := s.ScopeBelow(curScope)
		elems := s.GetScopeElements(class_scope)
		for _, elem := range elems {
			if elem.Kind == "Class" && s.ScopeAbove(class_scope, elem.Value) == curScope {
				toAdd := 0
				typ, _ := StringFromData(data, "type")
				isArr, _ := BoolFromData(data, "isArray")
				toAdd = SizeOfType(typ, isArr)

				classSize := 0
				if s, ok := elem.Data["size"]; ok {
					switch size := s.(type) {
					case int:
						classSize = size
					}
				}
				elem.Data["size"] = classSize + toAdd
				data["offset"] = classSize
				break
			}
		}
	}

	return symId
}

func (s *SymbolTable) AddScope(scope string) {
	s.scope = s.ScopeAbove(s.scope, scope)
	if _, ok := s.scopeElements[s.scope]; !ok {
		s.scopeElements[s.scope] = make([]string, 0)
	}
}
func (s *SymbolTable) ScopeAbove(scope, add string) string {
	return scope + "." + add
}

func (s *SymbolTable) DownScope() error {
	newscope, err := s.ScopeBelow(s.scope)
	if err == nil {
		s.scope = newscope
		//fmt.Printf("Scope dropped down to %s now \n",s.scope)
	}
	if s.funcType != "" {
		s.funcType = ""
	}
	return err
}
func (s *SymbolTable) ScopeBelow(scope string) (string, error) {
	tmp := str.Split(scope, ".")
	if len(tmp) < 1 {
		return "", fmt.Errorf("Can't drop scope, current scope is %s", s.scope)
	}
	tmp = tmp[:len(tmp)-1]
	newscope := str.Join(tmp, ".")
	return newscope, nil
}
func (s *SymbolTable) ScopeBelowWithCurr(scope string) (string, string, error) {
	tmp := str.Split(scope, ".")
	if len(tmp) < 1 {
		return "", "", fmt.Errorf("Can't drop scope, current scope is %s", s.scope)
	}
	head := tmp[len(tmp)-1]
	tmp = tmp[:len(tmp)-1]
	newscope := str.Join(tmp, ".")
	return newscope, head, nil
}

func (s *SymbolTable) GetScope() string {
	return s.scope
}

func (s *SymbolTable) GetElement(symId string) (SymbolTableElement, error) {
	if elem, ok := s.elems[symId]; ok {
		return elem, nil
	}
	return SymbolTableElement{}, fmt.Errorf("Element doesn't exists")
}

func (s *SymbolTable) GetTypeSymId(typ string) (symId string, err error) {
	for sId, symTabElem := range s.elems {
		if symTabElem.Kind == "Type" && symTabElem.Value == typ {
			symId = sId
			return
		}
	}
	err = fmt.Errorf("Type does not exist")
	return
}

func (s *SymbolTable) GetAllOfKind(kind string) (elems []SymbolTableElement) {
	elems = make([]SymbolTableElement, 0)
	for _, symTabElem := range s.elems {
		if symTabElem.Kind == kind {
			elems = append(elems, symTabElem)
		}
	}
	return
}

func (s *SymbolTable) GetScopeElements(scope string) []SymbolTableElement {
	if elemsSymIds, ok := s.scopeElements[scope]; ok {
		elems := make([]SymbolTableElement, 0, len(elemsSymIds))
		for _, symId := range elemsSymIds {
			elems = append(elems, s.elems[symId])
		}
		return elems
	}
	return make([]SymbolTableElement, 0) //if we aint got nut'n for the scope they get nut'n
}

func (s *SymbolTable) GetElementInScope(scope, value string) SymbolTableElement {
	if elemsSymIds, ok := s.scopeElements[scope]; ok {
		for _, symId := range elemsSymIds {
			if s.elems[symId].Value == value {
				return s.elems[symId]
			}

		}
	}
	return SymbolTableElement{}
}

func (s *SymbolTable) PrintTable() {
	fmt.Printf("Current Scope: %s\n", s.scope)
	fmt.Println("=================")
	fmt.Println("Elements:")
	keys := make([]string, 0, len(s.elems))
	for k := range s.elems {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		e := s.elems[key]
		e.PrintElement()
		fmt.Println("--------------")
	}
}

func (s *SymbolTable) PrintTableInAddOrder() {
	fmt.Printf("Current Scope: %s\n", s.scope)
	fmt.Println("=================")
	fmt.Println("Elements:")
	for _, key := range s.symIds {
		if e, ok := s.elems[key]; ok {
			e.PrintElement()
			fmt.Println("--------------")
		}
	}
}

type SymbolTableElement struct {
	Scope string
	SymId string
	Value string
	Kind  string
	Data  map[string]interface{}
}

func (s *SymbolTableElement) PrintElement() {
	fmt.Printf("Scope: %s, SymId: %s, Value: %s, Kind: %s\n", s.Scope, s.SymId, s.Value, s.Kind)
	fmt.Println("Extra Data:")
	for k, v := range s.Data {
		fmt.Printf("Key: %s, Value: %v\n", k, v)
	}
}

//Helper functions
func SizeOfType(typ string, isArr bool) (size int) {
	if isArr {
		return 4
	}
	size = 4
	switch typ {
	case "char", "bool":
		size = 1
	}
	return
}

func StringFromData(m map[string]interface{}, s string) (ret string, err error) {
	if e, ok := m[s]; ok {
		switch elem := e.(type) {
		case string:
			ret = elem
			return
		}
	}
	err = fmt.Errorf("Problem getting string (%s) from map", s)
	return
}

func StringSliceFromData(m map[string]interface{}, s string) (ret []string, err error) {
	if e, ok := m[s]; ok {
		switch elem := e.(type) {
		case []string:
			ret = elem
			return
		}
	}
	err = fmt.Errorf("Problem getting []string (%s) from map", s)
	return
}

func BoolFromData(m map[string]interface{}, s string) (ret bool, err error) {
	if e, ok := m[s]; ok {
		switch elem := e.(type) {
		case bool:
			ret = elem
			return
		}
	}
	err = fmt.Errorf("Problem getting bool (%s) from map", s)
	return
}

func IntFromData(m map[string]interface{}, s string) (ret int, err error) {
	if e, ok := m[s]; ok {
		switch elem := e.(type) {
		case int:
			ret = elem
			return
		}
	}
	err = fmt.Errorf("Problem getting int (%s) from map", s)
	return
}
