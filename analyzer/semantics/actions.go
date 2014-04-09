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
	switch value {
	case "true":
		value = "1"
	case "false":
		value = "0"
	}

	//symbol table action
	//check if there
	elem := s.st.GetElementInScope("g", value)
	symId := elem.SymId
	if symId == "" {
		//add it
		data := make(map[string]interface{})
		data["type"] = typ
		data["scope"] = "g"
		symId = s.st.AddElement(value, "LitVar", data, true)
	}

	s.sas.push(&Lit_Sar{value: value, scope: scope, typ: typ, symId: symId})
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

	//symbol table action
	//check if its in the symbol table
	symId, err := s.st.GetTypeSymId(value)
	if err != nil {
		data := make(map[string]interface{})
		data["type"] = value
		data["scope"] = "g"
		symId = s.st.AddElement(value, "Type", data, true)
	}

	s.sas.push(&Type_Sar{value: value, scope: scope, symId: symId})
}

func (s *SemanticManager) IExist(st *sym.SymbolTable) error {
	sar := s.sas.pop()
	if sar == nil {
		return fmt.Errorf("No sar on the stack")
	}
	if sar.Exists(st) {

		//SymId should be already set in Exists function!

		//icode
		elem, _ := st.GetElement(sar.GetSymId())
		if elem.Kind == "Ivar" {
			//fmt.Printf("ELEM: %#v\n",elem)
			if cls, ok := elem.Data["this_class"]; ok {
				//push a class sar
				switch class := cls.(type) {
				case string:
					s.sas.push(&Id_Sar{value: "this", typ: class, symId: "this", exists: true, scope: "g" + class})
				default:
					panic("Ivar is messed up compiler error")
				}

				//push the var sar
				s.sas.push(sar)
				//call RExist()
				return s.RExist(st)

			}
		}

		s.sas.push(sar)
		return nil
	}
	return fmt.Errorf("%s does not exist", sar.GetValue())
}

func (s *SemanticManager) VPush(value, scope, typ string) {

	//Should already be in symbol table ... lets get the symId
	elems := s.st.GetScopeElements(scope)

	for _, elem := range elems {
		if value == elem.Value {
			if t, ok := elem.Data["type"]; ok && t == typ {
				if elem.Kind == "Ivar" {
					//fmt.Printf("ELEM: %#v\n",elem)
					if cls, ok := elem.Data["this_class"]; ok {
						//push a class sar
						switch class := cls.(type) {
						case string:
							s.sas.push(&Id_Sar{value: "this", typ: class, symId: "this", exists: true, scope: "g" + class})
						default:
							panic("Ivar is messed up compiler error")
						}

						//push the var sar
						s.sas.push(&Id_Sar{value: value, scope: scope, exists: true, typ: typ, symId: elem.SymId})
						//call RExist()
						s.RExist(s.st)

					}
				} else {
					s.sas.push(&Id_Sar{value: value, scope: scope, exists: true, typ: typ, symId: elem.SymId})
				}
				break
			}
		}
	}
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
	//class_sar now has symId set to the appropriate one
	scopeToTest := st.ScopeAbove("g", class_sar.GetType())

	value := class_sar.GetValue() + "." + var_sar.GetValue()
	ref_sar := &Ref_Sar{value: value, scope: scopeToTest, class_sar: class_sar, var_sar: var_sar}

	if ref_sar.InstExists(st, var_sar) {

		//symbol table action
		data := make(map[string]interface{})
		data["type"] = var_sar.GetType()
		data["class_symId"] = class_sar.GetSymId()
		data["var_symId"] = var_sar.GetSymId()
		switch var_sar.(type) {
		case *Func_Sar:
		default:
			data["indirect"] = true
		}
		value := fmt.Sprintf("%s.%s", class_sar.GetValue(), var_sar.GetValue())
		symId := s.st.AddElement(value, "Tvar", data, true)

		ref_sar.SetSymId(symId)

		s.sas.push(ref_sar)

		//icode
		switch my_sar := var_sar.(type) {
		case *Func_Sar:
			s.gen.AddRow("", "FRAME", class_sar.GetSymId(), my_sar.GetSymId(), "", s.lx.GetCurFullLine())
			for _, param := range my_sar.GetAlSar().GetArgs() {
				s.gen.AddRow("", "PUSH", param.GetSymId(), "", "", s.lx.GetCurFullLine())
			}
			s.gen.AddRow("", "CALL", my_sar.GetSymId(), "", "", s.lx.GetCurFullLine())
			s.gen.AddRow("", "PEEK", symId, "", "", s.lx.GetCurFullLine())
		default:
			s.gen.AddRow("", "REF", symId, var_sar.GetSymId(), class_sar.GetSymId(), s.lx.GetCurFullLine())
		}

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
		//fmt.Printf("%#v\n",type_sar)
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

func (s *SemanticManager) SetupFunc(st *sym.SymbolTable) {
	//icode
	curScope := st.GetScope()
	tmp := str.Split(curScope, ".")
	if len(tmp) < 1 {
		return
	}
	f := tmp[len(tmp)-1]
	tmp = tmp[:len(tmp)-1]
	searchScope := str.Join(tmp, ".")

	elems := st.GetScopeElements(searchScope)
	symId := ""
	var e sym.SymbolTableElement
LOOP:
	for _, elem := range elems {
		switch elem.Kind {
		case "Constructor", "Method", "Main":
			if elem.Value == f {
				symId = elem.SymId
				e = elem
				break LOOP
			}
		}
	}
	s.gen.SwitchToMain()
	s.gen.AddRow("", "FUNC", symId, "", "", s.lx.GetCurFullLine())
	switch e.Kind {
	case "Constructor":
		classStaticInit := st.GetElementInScope("g."+e.Value, e.Value+"StaticInit") //Later figure out scope class

		s.gen.AddRow("", "FRAME", "this", classStaticInit.SymId, "", s.lx.GetCurFullLine())
		s.gen.AddRow("", "CALL", classStaticInit.SymId, "", "", s.lx.GetCurFullLine())
	}
}

func (s *SemanticManager) StaticInit(className string, st *sym.SymbolTable) {
	s.gen.SwitchToMain()

	fElem := st.GetElementInScope("g."+className, className+"StaticInit")
	s.gen.AddRow("", "FUNC", fElem.SymId, "", "", s.lx.GetCurFullLine())
	s.gen.AddAndResetStatic()
	s.gen.AddRow("", "RTN", "", "", "", s.lx.GetCurFullLine())

	s.gen.SwitchToStatic()
}

func (s *SemanticManager) ReturnFunc(st *sym.SymbolTable) {
	//icode
	s.gen.AddRow("", "RTN", "", "", "", s.lx.GetCurFullLine())
	s.gen.SwitchToStatic()
}

func (s *SemanticManager) ReturnThisFunc(st *sym.SymbolTable) {
	//icode
	s.gen.SwitchToMain()
	s.gen.AddRow("", "RETURN", "this", "", "", s.lx.GetCurFullLine())
	s.gen.SwitchToStatic()
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
		return fmt.Errorf("Invalid array offset, should be int")
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

	if !id_sar.Exists(st) {
		return fmt.Errorf("Expected identifier for array, received %s", id.GetValue())
	}

	s.sas.push(id_sar)
	if err := s.IExist(st); err != nil {
		return fmt.Errorf("Expected identifier for array, received %s", id.GetValue())
	}

	base := s.sas.pop()

	if !exp.Exists(st) {
		return fmt.Errorf("Expected offset for array, received %s", exp.GetValue())
	}

	//symbol table
	data := make(map[string]interface{})
	data["arr_symId"] = id_sar.GetSymId()
	data["type"] = id_sar.GetType()
	data["exp_symId"] = exp.GetSymId()
	data["indirect"] = true
	data["isArray"] = false
	symId := s.st.AddElement(value, "Tvar", data, true)

	arr_sar := &Arr_Sar{value: value, typ: id_sar.GetType(), scope: st.GetScope(), id_sar: id_sar, exp: exp, symId: symId}

	if !arr_sar.Exists(st) {
		return fmt.Errorf("%s does not exists", value)
	}

	s.sas.push(arr_sar)

	s.debugMessage(fmt.Sprintf("Type: %s, with array size %s", id_sar.GetValue(), exp.GetValue()))

	//icode
	s.gen.AddRow("", "AEF", symId, exp.GetSymId(), base.GetSymId(), s.lx.GetCurFullLine())

	return
}

func (s *SemanticManager) If() (err error) {
	sar := s.sas.pop()
	if sar.GetType() != "bool" {
		err = fmt.Errorf("not a bool for if statement")
	}
	//icode
	ifLabel := s.st.GenSymId("If")
	s.gen.AddRow("", "BF", sar.GetSymId(), ifLabel, "", s.lx.GetCurFullLine())
	s.gen.AddLabel(ifLabel)
	return
}

func (s *SemanticManager) EndIf() {
	s.gen.LabelNextRow()
}

func (s *SemanticManager) Else() {
	//icode
	elseLabel := s.st.GenSymId("Else")
	s.gen.AddRow("", "JMP", elseLabel, "", "", s.lx.GetCurFullLine())
	s.gen.AddElseLabel(elseLabel)
	s.EndIf()
}

func (s *SemanticManager) EndElse() {
	s.gen.ElseLblNextRow()
}

func (s *SemanticManager) While() (err error) {
	sar := s.sas.pop()
	if sar.GetType() != "bool" {
		err = fmt.Errorf("not a bool for while statement")
	}
	//icode
	endLabel := s.st.GenSymId("End")
	s.gen.AddRow("", "BF", sar.GetSymId(), endLabel, "", s.lx.GetCurFullLine())
	s.gen.AddLabel(endLabel)
	return
}

func (s *SemanticManager) InitWhile() (initLabel string) {
	initLabel = s.st.GenSymId("While")
	s.gen.AddLabel(initLabel)
	s.gen.LabelNextRow()
	return
}

func (s *SemanticManager) EndWhile(initLabel string) {
	s.gen.AddRow("", "JMP", initLabel, "", "", s.lx.GetCurFullLine())
	s.gen.LabelNextRow()
}

func (s *SemanticManager) Return(st *sym.SymbolTable, isVoid bool) (err error) {
	if err := s.EoE(); err != nil {
		return err
	}

	funType := st.GetFunctionType()

	if isVoid {
		if funType != "void" {
			err = fmt.Errorf("Returning something from a void function")
		}
		s.debugMessage("Returning from a void function")
		return
	}

	sar := s.sas.pop()

	s.debugMessage(fmt.Sprintf("Expression return type (%s) expected (%s)", sar.GetType(), funType))
	switch funType {
	case "int", "char", "bool", "void":
		if sar.GetType() != funType {
			err = fmt.Errorf("Return type (%s) does not match declared return type (%s)", sar.GetType(), funType)
		}
	default:
		if sar.GetType() != funType && sar.GetValue() != "null" && sar.GetValue() != "this" {
			err = fmt.Errorf("Return type (%s) does not match declared return type (%s)", sar.GetType(), funType)
		}
	}

	//icode
	if sar.GetValue() == "this" {
		s.gen.AddRow("", "RETURN", "this", "", "", s.lx.GetCurFullLine())
	} else {
		s.gen.AddRow("", "RETURN", sar.GetSymId(), "", "", s.lx.GetCurFullLine())
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

	//icode
	s.gen.AddRow("", "WRITE", sar.GetSymId(), "", "", s.lx.GetCurFullLine())

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

	//icode
	s.gen.AddRow("", "READ", sar.GetSymId(), "", "", s.lx.GetCurFullLine())

	return
}

func (s *SemanticManager) Atoi(scope string) (err error) {
	sar := s.sas.pop()
	if sar.GetType() != "char" {
		err = fmt.Errorf("not a char for atoi")
	}

	value := "atoi(" + sar.GetValue() + ")"

	//symbol table action
	data := make(map[string]interface{})
	data["type"] = sar.GetType()
	symId := s.st.AddElement(value, "Tvar", data, true)

	s.sas.push(&Tvar_Sar{value: value, typ: "int", scope: scope, symId: symId})
	return
}

func (s *SemanticManager) Itoa(scope string) (err error) {
	sar := s.sas.pop()
	if sar.GetType() != "int" {
		err = fmt.Errorf("not a int for itoa")
	}

	value := "itoa(" + sar.GetValue() + ")"

	//symbol table action
	data := make(map[string]interface{})
	data["type"] = sar.GetType()
	symId := s.st.AddElement(value, "Tvar", data, true)

	s.sas.push(&Tvar_Sar{value: value, typ: "char", scope: scope, symId: symId})
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

	//icode (will get overwrote with constructor symId later)
	classSymId := type_sar.GetSymId()

	value := type_sar.GetValue() + "("
	for i, arg := range al_sar.GetArgs() {
		if i > 0 {
			value += ", "
		}
		value += arg.GetValue()
	}
	value += ")"

	//symbol table action
	data := make(map[string]interface{})
	data["type"] = type_sar.GetType()
	data["type_symId"] = type_sar.GetSymId()
	symId := s.st.AddElement(value, "Tvar", data, true)

	new_sar := &New_Sar{value: value, typ: type_sar.GetValue(), scope: "g." + type_sar.GetValue(), type_sar: type_sar, al_sar: al_sar, symId: symId}

	if !new_sar.ConstructorExists(st) {
		return fmt.Errorf("Constructor with %d arguments doens't exist", len(al_sar.GetArgs()))
	}

	//icode
	s.gen.AddRow("", "NEWI", classSymId, symId, "", s.lx.GetCurFullLine())
	s.gen.AddRow("", "FRAME", symId, type_sar.GetSymId(), "", s.lx.GetCurFullLine())
	for _, param := range al_sar.GetArgs() {
		s.gen.AddRow("", "PUSH", param.GetSymId(), "", "", s.lx.GetCurFullLine())
	}
	s.gen.AddRow("", "CALL", type_sar.GetSymId(), "", "", s.lx.GetCurFullLine())
	s.gen.AddRow("", "PEEK", symId, "", "", s.lx.GetCurFullLine())

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

	//symbol table
	data := make(map[string]interface{})
	data["type_symId"] = type_sar.GetSymId()
	data["type"] = type_sar.GetValue()
	data["isArray"] = true
	data["exp_symId"] = sar.GetSymId()
	symId := s.st.AddElement(value, "Tvar", data, true)

	new_sar := &New_Sar{value: value, typ: type_sar.GetValue(), scope: "g." + type_sar.GetValue(), type_sar: type_sar, al_sar: nil, symId: symId}

	if type_sar.GetType() == "void" {
		return fmt.Errorf("Array cannot be of type void")
	}

	s.sas.push(new_sar)

	s.debugMessage(fmt.Sprintf("Type: %s, with array size %s", type_sar.GetValue(), sar.GetValue()))

	//icode
	data = make(map[string]interface{})
	data["type_symId"] = type_sar.GetSymId()
	data["type"] = sar.GetType()
	symId = s.st.AddElement(value, "Tvar", data, true)

	s.gen.AddRow("", "MUL", symId, s.GetTypeSize(type_sar.GetValue()), sar.GetSymId(), s.lx.GetCurFullLine())
	s.gen.AddRow("", "NEW", symId, new_sar.GetSymId(), "", s.lx.GetCurFullLine())

	return
}

func (s *SemanticManager) Cd(st *sym.SymbolTable, className string) (err error) {
	scope := st.GetScope()
	tmp := str.Split(scope, ".")
	if len(tmp) != 2 {
		err = fmt.Errorf("Cannot get class name from scope %s", scope)
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
		return fmt.Errorf("Not enough operands to test arithmetic operator")
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

	//icode
	switch op {
	case "+":
		s.gen.AddRow("", "ADD", symId, op1.GetSymId(), op2.GetSymId(), s.lx.GetCurFullLine())
	case "-":
		s.gen.AddRow("", "SUB", symId, op1.GetSymId(), op2.GetSymId(), s.lx.GetCurFullLine())
	case "*":
		s.gen.AddRow("", "MUL", symId, op1.GetSymId(), op2.GetSymId(), s.lx.GetCurFullLine())
	case "/":
		s.gen.AddRow("", "DIV", symId, op1.GetSymId(), op2.GetSymId(), s.lx.GetCurFullLine())
	}

	return nil
}

func (s *SemanticManager) IsSarAnArray(sar SemanticActionRecord) (isArray bool) {
	//var err error
	switch stype := sar.(type) {
	case *Lit_Sar:
		isArray = false
	case *Ref_Sar:
		elem, _ := s.st.GetElement(stype.GetSymId())
		var_symId, _ := sym.StringFromData(elem.Data,"var_symId")
		varElem, _ := s.st.GetElement(var_symId)
		isArray, _ = sym.BoolFromData(varElem.Data,"isArray")
		//if err != nil { panic(fmt.Sprintf("Could not find isArray for %s",elem.SymId)) }
	default:
		elem, _ := s.st.GetElement(stype.GetSymId())
		isArray, _ = sym.BoolFromData(elem.Data,"isArray")
		//if err != nil { panic(fmt.Sprintf("Could not find isArray for %s",elem.SymId)) }
	}
	return
}

func (s *SemanticManager) AssignmentOperator() error {
	op1 := s.sas.pop()
	op2 := s.sas.pop()
	if op1 == nil || op2 == nil {
		return fmt.Errorf("Not enough operands to test assignment operator")
	}
	op1Typ := op1.GetType()
	op2Typ := op2.GetType()

	if s.IsSarAnArray(op1) != s.IsSarAnArray(op2) {
		return fmt.Errorf("Must assign arrays to arrays %#v, %#v\n", op1, op2)
	}

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

	//iCode
	s.gen.AddRow("", "MOV", op2.GetSymId(), op1.GetSymId(), "", s.lx.GetCurFullLine())

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

	//icode
	switch op {
	case "<":
		s.gen.AddRow("", "LT", symId, op2.GetSymId(), op1.GetSymId(), s.lx.GetCurFullLine())
	case "<=":
		s.gen.AddRow("", "LTE", symId, op2.GetSymId(), op1.GetSymId(), s.lx.GetCurFullLine())
	case ">":
		s.gen.AddRow("", "GT", symId, op2.GetSymId(), op1.GetSymId(), s.lx.GetCurFullLine())
	case ">=":
		s.gen.AddRow("", "GTE", symId, op2.GetSymId(), op1.GetSymId(), s.lx.GetCurFullLine())
	}

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

	//icode
	switch op {
	case "==":
		s.gen.AddRow("", "EQ", symId, op2.GetSymId(), op1.GetSymId(), s.lx.GetCurFullLine())
	case "!=":
		s.gen.AddRow("", "NEQ", symId, op2.GetSymId(), op1.GetSymId(), s.lx.GetCurFullLine())
	}
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

	//icode
	switch op {
	case "&&":
		s.gen.AddRow("", "AND", symId, op2.GetSymId(), op1.GetSymId(), s.lx.GetCurFullLine())
	case "||":
		s.gen.AddRow("", "OR", symId, op2.GetSymId(), op1.GetSymId(), s.lx.GetCurFullLine())
	}
	return nil
}
