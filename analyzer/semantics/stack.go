package semantics

import (
	//"fmt"
	tok "github.com/sbditto85/compiler/token"
	sym "github.com/sbditto85/compiler/symbol_table"
)

//////////////////////////////////
// SemanticManager
//////////////////////////////////
type SemanticManager struct {
	ops *OperatorStack
	sas *SemanticActionStack
}
func NewSemanticManager() *SemanticManager {
	ops := NewOperatorStack()
	sas := NewSemanticActionStack()
	return &SemanticManager{ops:ops,sas:sas}
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
type Operator struct{
	value string
	precidence int
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
	exits bool
	token *tok.Token
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
	
	return true
}
