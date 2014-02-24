package semantics

import (
	"fmt"
	sym "github.com/sbditto85/compiler/symbol_table"
)

//////////////////////////////////
// SemanticManager
//////////////////////////////////
type SemanticManager struct {
	ops *OperatorStack
	sas *SemanticActionStack
	st *sym.SymbolTable
	debug bool
}
func NewSemanticManager(st *sym.SymbolTable, debug bool) *SemanticManager {
	ops := NewOperatorStack()
	sas := NewSemanticActionStack()
	return &SemanticManager{ops:ops, sas:sas, st:st, debug:debug}
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
func (s *OperatorStack) GetPrec(op string) (precIn, precOn int, err error) {
	switch op {
	case "=":
		precIn = 1
		precOn = 1
	case "+","-":
		precIn = 11
		precOn = 11
	case "*","/","%":
		precIn = 13
		precOn = 13
	case "(","[":
		precIn = 15
		precOn = -1
	case ")","]":
		precIn = 0
		precOn = 0
	default:
		err = fmt.Errorf("No Precidence for operator %s",op)
	}
	return
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
	case "+","-","*","/":
		return s.ArithmeticOperator(o.value)
	}
	//panic(fmt.Sprintf("Operator not found %s",o.value))
	return fmt.Errorf("Operator not found %s",o.value)
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
	GetScope() string
	IsSameType(other SemanticActionRecord) bool
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
}
func (i *Id_Sar) GetValue() string {
	return i.value
}
func (i *Id_Sar) GetType() string {
	return i.typ
}
func (i *Id_Sar) GetScope() string {
	return i.scope
}
func (i *Id_Sar) IsSameType(other SemanticActionRecord) bool {
	return i.typ == other.GetType()
}
func (i *Id_Sar) Exists(st *sym.SymbolTable) bool {
	scopeChecking := st.GetScope()
	var err error
	for scopeChecking != "" {
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
		scopeChecking,err = st.ScopeBelow(scopeChecking)
		if err != nil {
			return false
		}
	}
	return false
}

type Tvar_Sar struct {
	value string
	typ string
	scope string
	symId string
}
func (i *Tvar_Sar) GetValue() string {
	return i.value
}
func (i *Tvar_Sar) GetType() string {
	return i.typ
}
func (i *Tvar_Sar) GetScope() string {
	return i.scope
}
func (i *Tvar_Sar) IsSameType(other SemanticActionRecord) bool {
	return i.typ == other.GetType()
}
func (i *Tvar_Sar) Exists(st *sym.SymbolTable) bool {
	return true
}

type Ref_Sar struct {
	value string
	typ string
	scope string
	exists bool
	class_sar SemanticActionRecord
	var_sar SemanticActionRecord
}
func (r *Ref_Sar) GetValue() string {
	return r.value
}
func (r *Ref_Sar) GetType() string {
	return r.typ
}
func (r *Ref_Sar) GetScope() string {
	return r.scope
}
func (r *Ref_Sar) IsSameType(other SemanticActionRecord) bool {
	return r.typ == other.GetType()
}
func (r *Ref_Sar) Exists(st *sym.SymbolTable) bool {
	return r.exists
}
func (r *Ref_Sar) InstExists(st *sym.SymbolTable, inside SemanticActionRecord) bool {
	elems := st.GetScopeElements(r.scope)

	//fmt.Printf("INSIDE: %#v\n",inside)

	if !r.class_sar.Exists(st) {
		return false
	}

	switch sar := inside.(type) {
	case *Func_Sar:
		method_sar := sar.GetIdSar()
		method_sar.scope = r.scope

		for _,elem := range(elems) {
			if elem.Kind == "Method" && elem.Value == method_sar.GetValue() {
				//check modifier
				if mod, ok := elem.Data["accessMod"]; !ok || mod != "public" {
					continue;
				}
				if p, ok := elem.Data["parameters"]; ok {
					switch params := p.(type) {
					case []sym.Parameter:
						al := sar.GetAlSar().GetArgs()
						for i,a := range(al) {
							if params[i].Typ != a.GetType() {
								return false
							}
						}
					default:
						return false //only one type should be
						}
				} else {
					return false //gotta have params to be a method
				}
				
				//set type and exists on *Id_Sar
				if typ,ok := elem.Data["type"]; ok {
					r.typ = typ.(string)
				} else {
					return false //NEED TYPE, break? check other scopes?
				}
				r.exists = true
				return true
			}
		}
	default:
		for _,elem := range(elems) {
			switch elem.Kind {
			case "Ivar":
				if elem.Value == inside.GetValue() {
					//check modifier
					if mod, ok := elem.Data["accessMod"]; !ok || mod != "public" {
						continue;
					}
					
					//set type and exists on *Id_Sar
					if typ,ok := elem.Data["type"]; ok {
						r.typ = typ.(string)
					} else {
						return false //NEED TYPE, break? check other scopes?
					}
					r.exists = true
					return true
				}
			} 
		}
	}
	return false
}

type Bal_Sar struct {
	scope string
}
func (b *Bal_Sar) GetValue() string {
	return ""
}
func (b *Bal_Sar) GetType() string {
	return ""
}
func (b *Bal_Sar) GetScope() string {
	return b.scope
}
func (b *Bal_Sar) IsSameType(other SemanticActionRecord) bool {
	return false
}
func (b *Bal_Sar) Exists(st *sym.SymbolTable) bool {
	return true
}


type Al_Sar struct {
	scope string
	args []SemanticActionRecord
}
func (a *Al_Sar) GetValue() string {
	return ""
}
func (a *Al_Sar) GetType() string {
	return ""
}
func (a *Al_Sar) GetScope() string {
	return a.scope
}
func (a *Al_Sar) IsSameType(other SemanticActionRecord) bool {
	return false
}
func (a *Al_Sar) Exists(st *sym.SymbolTable) bool {
	return true
}
func (a *Al_Sar) GetArgs() []SemanticActionRecord {
	return a.args
}

type Func_Sar struct {
	value string
	typ string
	scope string
	exists bool
	id_sar *Id_Sar
	al_sar *Al_Sar
}
func (f *Func_Sar) GetValue() string {
	return f.value
}
func (f *Func_Sar) GetType() string {
	return f.typ
}
func (f *Func_Sar) GetScope() string {
	return f.scope
}
func (f *Func_Sar) IsSameType(other SemanticActionRecord) bool {
	return f.typ == other.GetType()
}
func (f *Func_Sar) Exists(st *sym.SymbolTable) bool {
	return f.exists
}
func (f *Func_Sar) GetIdSar() *Id_Sar {
	return f.id_sar
}
func (f *Func_Sar) GetAlSar() *Al_Sar {
	return f.al_sar
}
