package semantics

import (
	"fmt"
	ic "github.com/sbditto85/compiler/analyzer/icode"
	lex "github.com/sbditto85/compiler/lexer"
	sym "github.com/sbditto85/compiler/symbol_table"
)

//////////////////////////////////
// SemanticManager
//////////////////////////////////
type SemanticManager struct {
	ops   *OperatorStack
	sas   *SemanticActionStack
	st    *sym.SymbolTable
	lx    *lex.Lexer
	gen   *ic.Generator
	debug bool
}

func NewSemanticManager(st *sym.SymbolTable, lx *lex.Lexer, gen *ic.Generator, debug bool) *SemanticManager {
	ops := NewOperatorStack()
	sas := NewSemanticActionStack()
	return &SemanticManager{ops: ops, sas: sas, st: st, lx: lx, gen: gen, debug: debug}
}
func (s *SemanticManager) SetLexer(l *lex.Lexer) {
	s.lx = l
}
func (s *SemanticManager) debugMessage(msg string) {
	if s.debug {
		fmt.Printf("SM: %s\n", msg)
	}
}
func (s *SemanticManager) SetDebug(debug bool) {
	s.debug = debug
}

type IdentifierAssignment struct {
	identSymId  string
	assignSymId string
}

//////////////////////////////////
// OpS
//////////////////////////////////
type OperatorStack struct {
	top  *OpElement
	size int
}

func NewOperatorStack() *OperatorStack {
	return &OperatorStack{}
}

type OpElement struct {
	value *Operator
	next  *OpElement
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
	case "+", "-":
		precIn = 11
		precOn = 11
	case "*", "/", "%":
		precIn = 13
		precOn = 13
	case "(", "[":
		precIn = 15
		precOn = -1
	case ")", "]":
		precIn = 0
		precOn = 0
	case "<", ">", "<=", ">=":
		precIn = 9
		precOn = 9
	case "==", "!=":
		precIn = 7
		precOn = 7
	case "&&":
		precIn = 5
		precOn = 5
	case "||":
		precIn = 3
		precOn = 3
	default:
		err = fmt.Errorf("No Precidence for operator %s", op)
	}
	return
}

type Operator struct {
	value  string
	precIn int
	precOn int
}

func (o *Operator) Perform(s *SemanticManager) error {
	switch o.value {
	case "=":
		return s.AssignmentOperator()
	case "+", "-", "*", "/":
		return s.ArithmeticOperator(o.value)
	case "<", ">", "<=", ">=":
		return s.GreaterLesser(o.value)
	case "==", "!=":
		return s.EqualNot(o.value)
	case "&&", "||":
		return s.IsBoolean(o.value)
	}
	//panic(fmt.Sprintf("Operator not found %s",o.value))
	return fmt.Errorf("Operator not found %s", o.value)
}

//////////////////////////////////
// SAS
//////////////////////////////////
type SemanticActionStack struct {
	top  *Element
	size int
}

func NewSemanticActionStack() *SemanticActionStack {
	return &SemanticActionStack{}
}

type Element struct {
	value SemanticActionRecord
	next  *Element
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

func (s *SemanticActionStack) print() {
	elem := s.top
	for elem != nil {
		fmt.Printf("%#v\n", elem.value)
		elem = elem.next
	}
}

type SemanticActionRecord interface {
	GetValue() string
	GetType() string
	GetScope() string
	IsSameType(other SemanticActionRecord) bool
	Exists(st *sym.SymbolTable) bool
	GetSymId() string
	SetSymId(symId string) error
}

//////////////////////////////////
// SARS
//////////////////////////////////
type Id_Sar struct {
	value  string
	typ    string
	scope  string
	exists bool
	symId  string
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
	if i.value == "this" {
		return true
	}
	scopeChecking := st.GetScope()
	var err error
	for scopeChecking != "" {
		elems := st.GetScopeElements(scopeChecking)
		for _, elem := range elems {
			switch elem.Kind {
			case "Lvar", "Ivar", "Parameter":
				if elem.Value == i.value {
					//fmt.Printf("elem: %#v\n",elem)
					//fmt.Printf("i: %#v\n",i)
					//set type and exists on *Id_Sar
					if typ, ok := elem.Data["type"]; ok {
						i.typ = typ.(string)
					} else {
						return false //NEED TYPE, break? check other scopes?
					}
					i.symId = elem.SymId
					i.exists = true
					return true
				}
			}
		}
		scopeChecking, err = st.ScopeBelow(scopeChecking)
		if err != nil {
			return false
		}
	}
	return false
}
func (i *Id_Sar) GetSymId() string {
	return i.symId
}
func (i *Id_Sar) SetSymId(symId string) error {
	i.symId = symId
	return nil
}

type Tvar_Sar struct {
	value string
	typ   string
	scope string
	symId string
}

func (t *Tvar_Sar) GetValue() string {
	return t.value
}
func (t *Tvar_Sar) GetType() string {
	return t.typ
}
func (t *Tvar_Sar) GetScope() string {
	return t.scope
}
func (t *Tvar_Sar) IsSameType(other SemanticActionRecord) bool {
	return t.typ == other.GetType()
}
func (t *Tvar_Sar) Exists(st *sym.SymbolTable) bool {
	return true
}
func (t *Tvar_Sar) GetSymId() string {
	return t.symId
}
func (t *Tvar_Sar) SetSymId(symId string) error {
	t.symId = symId
	return nil
}

type Ref_Sar struct {
	value     string
	typ       string
	scope     string
	exists    bool
	class_sar SemanticActionRecord
	var_sar   SemanticActionRecord
	symId     string
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

	//fmt.Printf("CLASS_SAR: %#v\n",r.class_sar)
	//fmt.Printf("INSIDE: %#v\n",inside)

	if r.class_sar.GetValue() != "this" && !r.class_sar.Exists(st) {
		return false
	}

	switch sar := inside.(type) {
	case *Func_Sar:
		method_sar := sar.GetIdSar()
		method_sar.scope = r.scope

		for _, elem := range elems {
			if elem.Kind == "Method" && elem.Value == method_sar.GetValue() {
				//check modifier
				if mod, ok := elem.Data["accessMod"]; !ok || mod != "public" {
					continue
				}
				if p, ok := elem.Data["parameters"]; ok {
					switch params := p.(type) {
					case []sym.Parameter:
						al := sar.GetAlSar().GetArgs()
						if len(params) != len(al) {
							return false
						}
						for i, a := range al {
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
				if typ, ok := elem.Data["type"]; ok {
					r.typ = typ.(string)
				} else {
					return false //NEED TYPE, break? check other scopes?
				}
				inside.SetSymId(elem.SymId)
				r.exists = true
				return true
			}
		}
	default:
		for _, elem := range elems {
			switch elem.Kind {
			case "Ivar":
				if elem.Value == inside.GetValue() {
					//check modifier
					if mod, ok := elem.Data["accessMod"]; !ok || mod != "public" {
						continue
					}

					//set type and exists on *Id_Sar
					if typ, ok := elem.Data["type"]; ok {
						r.typ = typ.(string)
					} else {
						return false //NEED TYPE, break? check other scopes?
					}
					inside.SetSymId(elem.SymId)
					r.exists = true
					return true
				}
			}
		}
	}
	return false
}
func (r *Ref_Sar) GetSymId() string {
	return r.symId
}
func (r *Ref_Sar) SetSymId(symId string) error {
	r.symId = symId
	return nil
}

type Bal_Sar struct {
	scope string
	symId string
}

func (b *Bal_Sar) GetSymId() string {
	return b.symId
}
func (b *Bal_Sar) SetSymId(symId string) error {
	b.symId = symId
	return nil
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
	symId string
	args  []SemanticActionRecord
}

func (a *Al_Sar) GetSymId() string {
	return a.symId
}
func (a *Al_Sar) SetSymId(symId string) error {
	a.symId = symId
	return nil
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
	value  string
	typ    string
	scope  string
	symId  string
	exists bool
	id_sar *Id_Sar
	al_sar *Al_Sar
}

func (f *Func_Sar) GetSymId() string {
	return f.symId
}
func (f *Func_Sar) SetSymId(symId string) error {
	f.symId = symId
	return nil
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

	scope, err := st.ScopeBelow(f.id_sar.GetScope())
	if err != nil {
		return false
	}
	elems := st.GetScopeElements(scope)

	for _, elem := range elems {
		if elem.Kind == "Method" && elem.Value == f.id_sar.GetValue() {
			var pSymIds []string
			if ps, ok := elem.Data["paramSymIds"]; ok {
				switch psConv := ps.(type) {
				case []string:
					pSymIds = psConv
				}
			}
			//Check params
			if p, ok := elem.Data["parameters"]; ok {
				switch params := p.(type) {
				case []sym.Parameter:
					al := f.al_sar.GetArgs()
					if len(params) != len(al) {
						return false
					}
					for i, a := range al {
						if params[i].Typ != a.GetType() {
							if i < len(pSymIds) {
								a.SetSymId(pSymIds[i])
							}
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
			if typ, ok := elem.Data["type"]; ok {
				f.typ = typ.(string)
			} else {
				return false //NEED TYPE, break? check other scopes?
			}
			f.symId = elem.SymId
			f.exists = true
			return true
		}
	}
	return f.exists
}
func (f *Func_Sar) GetIdSar() *Id_Sar {
	return f.id_sar
}
func (f *Func_Sar) GetAlSar() *Al_Sar {
	return f.al_sar
}

type Type_Sar struct {
	value  string
	typ    string
	scope  string
	symId  string
	exists bool
}

func (t *Type_Sar) GetSymId() string {
	return t.symId
}
func (t *Type_Sar) SetSymId(symId string) error {
	t.symId = symId
	return nil
}
func (t *Type_Sar) GetValue() string {
	return t.value
}
func (t *Type_Sar) GetType() string {
	return t.typ
}
func (t *Type_Sar) GetScope() string {
	return t.scope
}
func (t *Type_Sar) IsSameType(other SemanticActionRecord) bool {
	return t.typ == other.GetType()
}
func (t *Type_Sar) Exists(st *sym.SymbolTable) bool {
	switch t.value {
	case "int", "char", "bool", "void":

		_, err := st.GetTypeSymId(t.value)
		if err != nil {
			data := make(map[string]interface{})
			data["type"] = t.value
			data["scope"] = "g"
			st.AddElement(t.value, "Type", data, true)
		}
		return true
	}

	elems := st.GetScopeElements("g")

	for _, elem := range elems {
		if elem.Kind == "Class" && elem.Value == t.GetValue() {
			t.exists = true
			t.symId = elem.SymId
			t.typ = elem.Value
			return true
		}
	}

	return false
}

type New_Sar struct {
	value    string
	typ      string
	scope    string
	symId    string
	exists   bool
	type_sar *Type_Sar
	al_sar   *Al_Sar
}

func (n *New_Sar) GetSymId() string {
	return n.symId
}
func (n *New_Sar) SetSymId(symId string) error {
	n.symId = symId
	return nil
}
func (n *New_Sar) GetValue() string {
	return n.value
}
func (n *New_Sar) GetType() string {
	return n.typ
}
func (n *New_Sar) GetScope() string {
	return n.scope
}
func (n *New_Sar) IsSameType(other SemanticActionRecord) bool {
	return n.typ == other.GetType()
}
func (n *New_Sar) Exists(st *sym.SymbolTable) bool {
	return n.exists
}
func (n *New_Sar) ConstructorExists(st *sym.SymbolTable) bool {
	elems := st.GetScopeElements("g." + n.type_sar.GetValue())

	for _, elem := range elems {
		if elem.Kind == "Constructor" && elem.Value == n.type_sar.GetValue() {
			//check modifier
			if mod, ok := elem.Data["accessMod"]; !ok || mod != "public" {
				continue
			}
			if p, ok := elem.Data["parameters"]; ok {
				var pSymIds []string
				if ps, ok := elem.Data["paramSymIds"]; ok {
					switch psConv := ps.(type) {
					case []string:
						pSymIds = psConv
					}
				}
				switch params := p.(type) {
				case []sym.Parameter:
					al := n.al_sar.GetArgs()
					if len(params) != len(al) {
						return false
					}
					for i, a := range al {
						if params[i].Typ != a.GetType() {
							if i < len(pSymIds) {
								a.SetSymId(pSymIds[i])
							}
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
			if typ, ok := elem.Data["type"]; ok {
				n.typ = typ.(string)
			} else {
				return false //NEED TYPE, break? check other scopes?
			}
			n.type_sar.SetSymId(elem.SymId)
			n.exists = true
			return true
		}
	}
	return false
}

type Lit_Sar struct {
	value string
	typ   string
	scope string
	symId string
}

func (l *Lit_Sar) GetSymId() string {
	return l.symId
}
func (l *Lit_Sar) SetSymId(symId string) error {
	l.symId = symId
	return nil
}
func (l *Lit_Sar) GetValue() string {
	return l.value
}
func (l *Lit_Sar) GetType() string {
	return l.typ
}
func (l *Lit_Sar) GetScope() string {
	return l.scope
}
func (l *Lit_Sar) IsSameType(other SemanticActionRecord) bool {
	return l.typ == other.GetType()
}
func (l *Lit_Sar) Exists(st *sym.SymbolTable) bool {
	return true
}

type Arr_Sar struct {
	value  string
	typ    string
	scope  string
	symId  string
	exp    SemanticActionRecord
	id_sar *Id_Sar
}

func (a *Arr_Sar) GetSymId() string {
	return a.symId
}
func (a *Arr_Sar) SetSymId(symId string) error {
	a.symId = symId
	return nil
}
func (a *Arr_Sar) GetValue() string {
	return a.value
}
func (a *Arr_Sar) GetType() string {
	return a.typ
}
func (a *Arr_Sar) GetScope() string {
	return a.scope
}
func (a *Arr_Sar) IsSameType(other SemanticActionRecord) bool {
	return a.typ == other.GetType()
}
func (a *Arr_Sar) Exists(st *sym.SymbolTable) bool {
	if a.exp.GetType() != "int" {
		return false
	}

	if !a.id_sar.Exists(st) {
		return false
	}

	a.typ = a.id_sar.GetType()

	symId := a.id_sar.GetSymId()

	if elem, err := st.GetElement(symId); err == nil {
		if val, ok := elem.Data["isArray"]; ok {
			switch v := val.(type) {
			case bool:
				a.symId = elem.SymId
				return v
			}
		}
	}

	return false
}
