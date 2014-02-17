package semantics

import (
	"fmt"
	tok "github.com/sbditto85/compiler/token"
	sym "github.com/sbditto85/compiler/symbol_table"
)

//////////////////////////////////
// SemanticManager
//////////////////////////////////
type SemanticManager struct {
	ops *OperatorStack
	sas *SemanticActionStack
	debug bool
}
func NewSemanticManager(debug bool) *SemanticManager {
	ops := NewOperatorStack()
	sas := NewSemanticActionStack()
	return &SemanticManager{ops:ops, sas:sas, debug:debug}
}
func (s *SemanticManager) debugMessage(msg string) {
	if s.debug {
		fmt.Printf("SM: %s\n",msg)
	} 
}
func (s *SemanticManager) SetDebug(debug bool) {
	s.debug = debug
}

//////////////////////////////////
// OpS
//////////////////////////////////
type OperatorStack struct {
	top *OpElement
	size int
}
func NewOperatorStack() *OperatorStack {
	return &OperatorStack{}
}
type OpElement struct {
	value *Operator
	next *OpElement
}
// Return the stack's length
func (s *OperatorStack) len() int {
	return s.size
}
// Push a new element onto the stack
func (s *OperatorStack) push(value *Operator) {
	s.top = &OpElement{value, s.top}
	s.size++
}
// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *OperatorStack) pop() (value *Operator) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}
func (s *OperatorStack) topElement() (value *Operator) {
	if s.size > 0 {
		value = s.top.value
		return
	}
	return nil
}

type Operator struct{
	value string
	precIn int
	precOn int
}
func (o *Operator) Perform(s *SemanticManager) error {
	switch o.value {
	case "=":
		return s.AssignmentOperator()
	}
	return fmt.Errorf("Operator not found")
}

//////////////////////////////////
// SAS
//////////////////////////////////
type SemanticActionStack struct {
	top *Element
	size int
}
func NewSemanticActionStack() *SemanticActionStack {
	return &SemanticActionStack{}
}
type Element struct {
	value SemanticActionRecord
	next *Element
}
// Return the stack's length
func (s *SemanticActionStack) len() int {
	return s.size
}
// Push a new element onto the stack
func (s *SemanticActionStack) push(value SemanticActionRecord) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *SemanticActionStack) pop() (value SemanticActionRecord) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

type SemanticActionRecord interface {
	GetValue() string
	GetType() string
	IsSameType(other SemanticActionRecord) bool
	GetToken() *tok.Token
	Exists(st *sym.SymbolTable) bool
}

//////////////////////////////////
// SARS
//////////////////////////////////
type Id_Sar struct {
	value string
	typ string
	scope string
	exists bool
	token *tok.Token
}
func (i *Id_Sar) GetValue() string {
	return i.value
}
func (i *Id_Sar) GetType() string {
	return i.typ
}
func (i *Id_Sar) IsSameType(other SemanticActionRecord) bool {
	return i.typ == other.GetType()
}
func (i *Id_Sar) GetToken() *tok.Token {
	return i.token
}
func (i *Id_Sar) Exists(st *sym.SymbolTable) bool {
	scopeChecking := st.GetScope()

	elems := st.GetScopeElements(scopeChecking)
	for _,elem := range(elems) {
		switch elem.Kind {
		case "Lvar","Ivar":
			if elem.Value == i.value {
				//fmt.Printf("elem: %#v\n",elem)
				//fmt.Printf("i: %#v\n",i)
				//set type and exists on *Id_Sar
				if typ,ok := elem.Data["type"]; ok {
					i.typ = typ.(string)
				} else {
					return false //NEED TYPE, break? check other scopes?
				}
				i.exists = true
				return true
			}
		} 
	}
	return true
}
