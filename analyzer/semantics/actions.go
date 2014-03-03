package semantics

import (
	"fmt"
	sym "github.com/sbditto85/compiler/symbol_table"
	str "strings"
)

func (s *SemanticManager) IPush(value, scope string) {
	s.sas.push(&Id_Sar{value: value, scope: scope})
}

func (s *SemanticManager) LPush(value, scope, typ string) {
	s.sas.push(&Lit_Sar{value: value, scope: scope, typ: typ})
}

func (s *SemanticManager) OPush(value string) (err error) {
	topOp := s.ops.topElement()
	precIn, precOn, errPrec := s.ops.GetPrec(value)
	if errPrec != nil {
		return errPrec
	}
	for s.ops.len() != 0 && topOp.precOn >= precIn {
		op := s.ops.pop()
		err = op.Perform(s)
		if err != nil {
			//panic(err.Error())
			return err
		}
		topOp = s.ops.topElement()
	}
	s.ops.push(&Operator{value: value, precIn: precIn, precOn: precOn})
	//fmt.Printf("OPS: %#v\n",s.ops)
	return
}

func (s *SemanticManager) TPush(value, scope string) {
	s.sas.push(&Type_Sar{value: value, scope: scope})
}

func (s *SemanticManager) IExist(st *sym.SymbolTable) error {
	sar := s.sas.pop()
	if sar == nil {
		return fmt.Errorf("No sar on the stack")
	}
	if sar.Exists(st) {
		s.sas.push(sar)
		return nil
	}
	return fmt.Errorf("%s does not exist", sar.GetValue())
}

func (s *SemanticManager) VPush(value, scope, typ string) {
	s.sas.push(&Id_Sar{value: value, scope: scope, exists: true, typ: typ})
	return
}

func (s *SemanticManager) RExist(st *sym.SymbolTable) error {
	var_sar := s.sas.pop()
	if var_sar == nil {
		return fmt.Errorf("No sar on the stack")
	}
	class_sar := s.sas.pop()
	if class_sar == nil {
		return fmt.Errorf("No sar on the stack")
	}
	if !class_sar.Exists(st) {
		return fmt.Errorf("%s doesn't exist", class_sar.GetValue())
	}
	scopeToTest := st.ScopeAbove("g", class_sar.GetType())

	value := class_sar.GetValue() + "." + var_sar.GetValue()
	ref_sar := &Ref_Sar{value: value, scope: scopeToTest, class_sar: class_sar, var_sar: var_sar}

	if ref_sar.InstExists(st, var_sar) {

		//TODO: add temp in symbol table for ref
		//symbol table action
		//data := make(map[string]interface{})
		//data["type"] = op1Typ
		//value := fmt.Sprintf("%s %s %s",op2.GetValue(),op,op1.GetValue())
		//symId := s.st.AddElement(value,"Tvar",data,true)

		s.sas.push(ref_sar)
		return nil
	}

	return fmt.Errorf("%s does not have member %s", class_sar.GetValue(), var_sar.GetValue())
}

func (s *SemanticManager) TExists(st *sym.SymbolTable) (err error) {
	ts := s.sas.pop()
	var type_sar *Type_Sar
	switch t := ts.(type) {
	case *Type_Sar:
		type_sar = t
	default:
		err = fmt.Errorf("tExists expects Type_Sar, received something else")
	}
	if !type_sar.Exists(st) {
		err = fmt.Errorf("Type %s doesn't exist", type_sar.GetValue())
	}
	return
}

func (s *SemanticManager) BAL(scope string) {
	s.sas.push(&Bal_Sar{scope: scope})
}

func (s *SemanticManager) EAL(scope string) {
	len := s.sas.len()
	args := make([]SemanticActionRecord, len, len)
	len--
LOOP:
	for arg := s.sas.pop(); arg != nil; arg = s.sas.pop() {
		switch arg.(type) {
		case *Bal_Sar:
			break LOOP
		}
		args[len] = arg
		len--
	}
	len++
	s.sas.push(&Al_Sar{scope: scope, args: args[len:]})
}

func (s *SemanticManager) Func(scope string) (err error) {
	as := s.sas.pop()
	var al_sar *Al_Sar
	switch a := as.(type) {
	case *Al_Sar:
		al_sar = a
	default:
		return fmt.Errorf("Expected Argument List for function")
	}

	is := s.sas.pop()
	var id_sar *Id_Sar
	switch i := is.(type) {
	case *Id_Sar:
		id_sar = i
	default:
		return fmt.Errorf("Expected identifier for function")
	}

	s.debugMessage(fmt.Sprintf("Identifer: %s, with %d Arguments", id_sar.GetValue(), len(al_sar.GetArgs())))

	fun_val := id_sar.GetValue() + "("
	for i, a := range al_sar.GetArgs() {
		if i != 0 {
			fun_val += ", "
		}
		fun_val += a.GetType()
	}
	fun_val += ")"
	s.sas.push(&Func_Sar{value: fun_val, scope: scope, id_sar: id_sar, al_sar: al_sar})
	return nil
}

func (s *SemanticManager) Arr(st *sym.SymbolTable) (err error) {
	exp := s.sas.pop()
	if exp.GetType() != "int" {
		return fmt.Errorf("Invalid array count, should be int")
	}

	id := s.sas.pop()
	var id_sar *Id_Sar
	switch i := id.(type) {
	case *Id_Sar:
		id_sar = i
	default:
		return fmt.Errorf("Expected identifier for array, received %s", id.GetValue())
	}

	value := id_sar.GetValue() + "[" + exp.GetValue() + "]"

	arr_sar := &Arr_Sar{value: value, typ: id_sar.GetType(), scope: st.GetScope(), id_sar: id_sar, exp: exp}

	if !arr_sar.Exists(st) {
		return fmt.Errorf("%s does not exists", value)
	}

	s.sas.push(arr_sar)

	s.debugMessage(fmt.Sprintf("Type: %s, with array size %s", id_sar.GetValue(), exp.GetValue()))

	return
}

func (s *SemanticManager) If() (err error) {
	sar := s.sas.pop()
	if sar.GetType() != "bool" {
		err = fmt.Errorf("not a bool for if statement")
	}
	return
}

func (s *SemanticManager) While() (err error) {
	sar := s.sas.pop()
	if sar.GetType() != "bool" {
		err = fmt.Errorf("not a bool for while statement")
	}
	return
}

func (s *SemanticManager) Cout() (err error) {
	if err := s.EoE(); err != nil {
		return err
	}

	sar := s.sas.pop()
	if sar.GetType() != "char" && sar.GetType() != "int" {
		err = fmt.Errorf("not a char or int for cout")
	}
	return
}

func (s *SemanticManager) Cin() (err error) {
	if err := s.EoE(); err != nil {
		return err
	}

	sar := s.sas.pop()
	if sar.GetType() != "char" && sar.GetType() != "int" {
		err = fmt.Errorf("not a char or int for cin")
	}
	return
}

func (s *SemanticManager) Atoi(scope string) (err error) {
	sar := s.sas.pop()
	if sar.GetType() != "char" {
		err = fmt.Errorf("not a char for atoi")
	}

	//TODO: add new element to symbol table
	value := "atoi(" + sar.GetValue() + ")"
	s.sas.push(&Tvar_Sar{value: value, typ: "int", scope: scope})
	return
}

func (s *SemanticManager) Itoa(scope string) (err error) {
	sar := s.sas.pop()
	if sar.GetType() != "int" {
		err = fmt.Errorf("not a int for itoa")
	}

	//TODO: add new element to symbol table
	value := "itoa(" + sar.GetValue() + ")"
	s.sas.push(&Tvar_Sar{value: value, typ: "char", scope: scope})
	return
}

func (s *SemanticManager) NewObj(st *sym.SymbolTable) (err error) {
	as := s.sas.pop()
	var al_sar *Al_Sar
	switch a := as.(type) {
	case *Al_Sar:
		al_sar = a
	default:
		return fmt.Errorf("Expected Argument List for function")
	}

	ts := s.sas.pop()
	var type_sar *Type_Sar
	switch t := ts.(type) {
	case *Type_Sar:
		type_sar = t
	default:
		return fmt.Errorf("Expected identifier for function")
	}

	if !type_sar.Exists(st) {
		return fmt.Errorf("Type %s doesn't exist", type_sar.GetValue())
	}

	value := type_sar.GetValue() + "("
	for i, arg := range al_sar.GetArgs() {
		if i > 0 {
			value += ", "
		}
		value += arg.GetValue()
	}
	value += ")"

	new_sar := &New_Sar{value: value, typ: type_sar.GetValue(), scope: "g." + type_sar.GetValue(), type_sar: type_sar, al_sar: al_sar}

	if !new_sar.ConstructorExists(st) {
		return fmt.Errorf("Constructor with %d arguments doens't exist", len(al_sar.GetArgs()))
	}

	s.sas.push(new_sar)

	s.debugMessage(fmt.Sprintf("Type: %s, with %d Arguments", type_sar.GetValue(), len(al_sar.GetArgs())))

	return
}

func (s *SemanticManager) NewArray(st *sym.SymbolTable) (err error) {
	sar := s.sas.pop()
	if sar.GetType() != "int" {
		return fmt.Errorf("Invalid array count, should be int")
	}

	ts := s.sas.pop()
	var type_sar *Type_Sar
	switch t := ts.(type) {
	case *Type_Sar:
		type_sar = t
	default:
		return fmt.Errorf("Expected identifier for new array")
	}

	if !type_sar.Exists(st) {
		return fmt.Errorf("Type %s doesn't exist", type_sar.GetValue())
	}

	value := type_sar.GetValue() + "[" + sar.GetValue() + "]"

	new_sar := &New_Sar{value: value, typ: type_sar.GetValue(), scope: "g." + type_sar.GetValue(), type_sar: type_sar, al_sar: nil}

	if type_sar.GetType() == "void" {
		return fmt.Errorf("Array cannot be of type void")
	}

	s.sas.push(new_sar)

	s.debugMessage(fmt.Sprintf("Type: %s, with array size %s", type_sar.GetValue(), sar.GetValue()))

	return
}

func (s *SemanticManager) Cd(st *sym.SymbolTable, className string) (err error) {
	scope := st.GetScope()
	tmp := str.Split(scope, ".")
	if len(tmp) != 2 {
		err = fmt.Errorf("Cannot get class name from scope", scope)
	}

	if tmp[1] != className {
		err = fmt.Errorf("Constructor name (%s) does not match class name (%s)", tmp[1], className)
	}

	return
}

func (s *SemanticManager) CloseParen() (err error) {
	for op := s.ops.topElement(); op != nil && op.value != "("; op = s.ops.topElement() {
		s.ops.pop()
		s.debugMessage(fmt.Sprintf("Testing operation %s ...", op.value))
		err := op.Perform(s)
		if err != nil {
			return err
		}
		s.debugMessage(fmt.Sprintf("... finished operation %s", op.value))
	}
	op := s.ops.pop()
	if op.value != "(" || op == nil {
		return fmt.Errorf("Close paren didn't find opening paren")
	}
	s.debugMessage("Finished )")
	return
}
func (s *SemanticManager) CloseAngleBracket() (err error) {
	for op := s.ops.topElement(); op != nil && op.value != "["; op = s.ops.topElement() {
		s.ops.pop()
		s.debugMessage(fmt.Sprintf("Testing operation %s ...", op.value))
		err := op.Perform(s)
		if err != nil {
			return err
		}
		s.debugMessage(fmt.Sprintf("... finished operation %s", op.value))
	}
	op := s.ops.pop()
	if op.value != "[" || op == nil {
		return fmt.Errorf("Close paren didn't find opening paren")
	}
	s.debugMessage("Finished ]")
	return
}
func (s *SemanticManager) Comma() (err error) {
	for op := s.ops.topElement(); op != nil && op.value != "("; op = s.ops.topElement() {
		s.ops.pop()
		s.debugMessage(fmt.Sprintf("Testing operation %s ...", op.value))
		err := op.Perform(s)
		if err != nil {
			return err
		}
		s.debugMessage(fmt.Sprintf("... finished operation %s", op.value))
	}
	s.debugMessage("Finished ,")
	return
}

func (s *SemanticManager) EoE() (err error) {
	for i := s.ops.len(); i > 0; i-- {
		op := s.ops.pop()
		s.debugMessage(fmt.Sprintf("Testing operation %s ...", op.value))
		err = op.Perform(s)
		if err != nil {
			//panic(err.Error())
			return err
		}
		s.debugMessage(fmt.Sprintf("... finished operation %s", op.value))
	}
	return
}

func (s *SemanticManager) ArithmeticOperator(op string) error {
	op1 := s.sas.pop()
	op2 := s.sas.pop()
	if op1 == nil || op2 == nil {
		return fmt.Errorf("Not enough operands to test assignment operator")
	}
	op1Typ := op1.GetType()
	op2Typ := op2.GetType()
	if op1Typ == "" {
		return fmt.Errorf("Operand doesn't have type %#v", op1)
	}
	if op2Typ == "" {
		return fmt.Errorf("Operand doesn't have type %#v", op2)
	}
	if op1Typ != "int" {
		return fmt.Errorf("Operand of type %s cannot perform %s", op1, op)
	}
	if op2Typ != "int" {
		return fmt.Errorf("Operand of type %s cannot perform %s", op1, op)
	}
	s.debugMessage(fmt.Sprintf("Comparing %s(%s) to %s(%s) for %s", op1.GetValue(), op1Typ, op2.GetValue(), op2Typ, op))
	if op1Typ != op2Typ {
		return fmt.Errorf("Cann't assign operand '%s' (%s) to '%s' (%s) types mismatch", op1.GetValue(), op1Typ, op2.GetValue(), op2Typ)
	}

	data := make(map[string]interface{})
	data["type"] = op1Typ
	value := fmt.Sprintf("%s %s %s", op2.GetValue(), op, op1.GetValue())
	symId := s.st.AddElement(value, "Tvar", data, true)

	s.sas.push(&Tvar_Sar{value: value, typ: op1Typ, scope: s.st.GetScope(), symId: symId})
	return nil
}

func (s *SemanticManager) AssignmentOperator() error {
	op1 := s.sas.pop()
	op2 := s.sas.pop()
	if op1 == nil || op2 == nil {
		return fmt.Errorf("Not enough operands to test assignment operator")
	}
	op1Typ := op1.GetType()
	op2Typ := op2.GetType()
	if op1Typ == "" {
		return fmt.Errorf("Operand doesn't have type %#v", op1)
	}
	if op2Typ == "" {
		return fmt.Errorf("Operand doesn't have type %#v", op2)
	}
	s.debugMessage(fmt.Sprintf("Comparing %s(%s) to %s(%s)", op1.GetValue(), op1Typ, op2.GetValue(), op2Typ))
	if op1Typ != op2Typ {
		return fmt.Errorf("Cann't assign operand %s(%s) to %s(%s) types mismatch", op1.GetValue(), op1Typ, op2.GetValue(), op2Typ)
	}
	return nil
}

func (s *SemanticManager) GreaterLesser(op string) error {
	op1 := s.sas.pop()
	op2 := s.sas.pop()
	if op1 == nil || op2 == nil {
		return fmt.Errorf("Not enough operands to test assignment operator %s", op)
	}
	op1Typ := op1.GetType()
	op2Typ := op2.GetType()
	if op1Typ == "" {
		return fmt.Errorf("Operand doesn't have type %#v", op1)
	}
	if op2Typ == "" {
		return fmt.Errorf("Operand doesn't have type %#v", op2)
	}

	s.debugMessage(fmt.Sprintf("Comparing %s(%s) to %s(%s) for op %s", op1.GetValue(), op1Typ, op2.GetValue(), op2Typ, op))

	if op1Typ != op2Typ {
		return fmt.Errorf("Cann't assign operand %s(%s) to %s(%s) types mismatch", op1.GetValue(), op1Typ, op2.GetValue(), op2Typ)
	}

	if op1Typ != "int" && op1Typ != "char" {
		return fmt.Errorf("%s not of comparable type (%s)", op1.GetValue(), op1Typ)
	}
	if op2Typ != "int" && op2Typ != "char" {
		return fmt.Errorf("%s not of comparable type (%s)", op2.GetValue(), op2Typ)
	}

	data := make(map[string]interface{})
	data["type"] = "bool"
	value := fmt.Sprintf("%s %s %s", op2.GetValue(), op, op1.GetValue())
	symId := s.st.AddElement(value, "Tvar", data, true)

	s.sas.push(&Tvar_Sar{value: value, typ: "bool", scope: s.st.GetScope(), symId: symId})

	return nil
}

func (s *SemanticManager) EqualNot(op string) error {
	op1 := s.sas.pop()
	op2 := s.sas.pop()
	if op1 == nil || op2 == nil {
		return fmt.Errorf("Not enough operands to test assignment operator %s", op)
	}
	op1Typ := op1.GetType()
	op2Typ := op2.GetType()
	if op1Typ == "" {
		return fmt.Errorf("Operand doesn't have type %#v", op1)
	}
	if op2Typ == "" {
		return fmt.Errorf("Operand doesn't have type %#v", op2)
	}

	s.debugMessage(fmt.Sprintf("Comparing %s(%s) to %s(%s) for op %s", op1.GetValue(), op1Typ, op2.GetValue(), op2Typ, op))

	if op1Typ != op2Typ {
		return fmt.Errorf("Cann't assign operand %s(%s) to %s(%s) types mismatch", op1.GetValue(), op1Typ, op2.GetValue(), op2Typ)
	}

	if op1Typ == "void" {
		return fmt.Errorf("%s not of comparable type (%s)", op1.GetValue(), op1Typ)
	}
	if op2Typ == "void" {
		return fmt.Errorf("%s not of comparable type (%s)", op2.GetValue(), op2Typ)
	}

	data := make(map[string]interface{})
	data["type"] = "bool"
	value := fmt.Sprintf("%s %s %s", op2.GetValue(), op, op1.GetValue())
	symId := s.st.AddElement(value, "Tvar", data, true)

	s.sas.push(&Tvar_Sar{value: value, typ: "bool", scope: s.st.GetScope(), symId: symId})

	return nil
}

func (s *SemanticManager) IsBoolean(op string) error {
	op1 := s.sas.pop()
	if op1.GetType() != "bool" {
		return fmt.Errorf("%s is not bool for %s", op1.GetValue(), op)
	}

	op2 := s.sas.pop()
	if op2.GetType() != "bool" {
		return fmt.Errorf("%s is not bool for %s", op2.GetValue(), op)
	}

	//symbol table action
	data := make(map[string]interface{})
	data["type"] = "bool"
	value := fmt.Sprintf("%s %s %s", op2.GetValue(), op, op1.GetValue())
	symId := s.st.AddElement(value, "Tvar", data, true)

	s.debugMessage(fmt.Sprintf("Comparing %s and %s as bool for op %s", op1.GetValue(), op2.GetValue(), op))

	s.sas.push(&Tvar_Sar{value: value, typ: "bool", scope: "", symId: symId})

	return nil
}
