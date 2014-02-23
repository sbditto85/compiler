package semantics

import (
	"fmt"
	sym "github.com/sbditto85/compiler/symbol_table"
)

func (s *SemanticManager) IPush(value, scope string) {
	s.sas.push(&Id_Sar{value:value, scope:scope})
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
	return fmt.Errorf("%s does not exist",sar.GetValue())
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
		return fmt.Errorf("%s doesn't exist",class_sar.GetValue())
	}
	scopeToTest := st.ScopeAbove("g",class_sar.GetType())

	value := class_sar.GetValue()+"."+var_sar.GetValue()
	ref_sar := &Ref_Sar{value:value, scope:scopeToTest}

	if ref_sar.InstExists(st,var_sar) {
		s.sas.push(ref_sar)
		return nil
	}
	
	return fmt.Errorf("%s does not have member %s",class_sar.GetValue(),var_sar.GetValue())
}

func (s *SemanticManager) BAL(scope string) {
	s.sas.push(&Bal_Sar{scope:scope})
}

func (s *SemanticManager) EAL(scope string) {
	len := s.sas.len()
	args := make([]SemanticActionRecord,len,len)
	len--
LOOP:	for arg := s.sas.pop(); arg != nil; arg = s.sas.pop() {
		switch arg.(type) {
		case *Bal_Sar:
			break LOOP
		}
		args[len] = arg
		len--
	}
	len++
	s.sas.push(&Al_Sar{scope:scope,args:args[len:]})
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

	s.debugMessage(fmt.Sprintf("Identifer: %s, with %d Arguments",id_sar.GetValue(),len(al_sar.GetArgs())))

	fun_val := id_sar.GetValue() + "("
	for i,a := range(al_sar.GetArgs()) {
		if i != 0 {
			fun_val += ", "
		}
		fun_val += a.GetType()
	}
	fun_val += ")"
	s.sas.push(&Func_Sar{value:fun_val, scope:scope, id_sar:id_sar, al_sar:al_sar})
	return nil
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
	s.ops.push(&Operator{value:value, precIn:precIn, precOn:precOn})
	//fmt.Printf("OPS: %#v\n",s.ops)
	return
}

func (s *SemanticManager) CloseParen() (err error) {
	err = s.Comma()
	if err != nil {
		return err
	}
	op := s.ops.pop()
	if op.value != "(" || op == nil {
		return fmt.Errorf("Close paren didn't find opening paren")
	}
	return
}

func (s *SemanticManager) Comma() (err error) {
	for op := s.ops.topElement(); op != nil && op.value != "("; op = s.ops.topElement() {
		s.ops.pop()
		s.debugMessage(fmt.Sprintf("Testing operation %s ...",op.value))
		err := op.Perform(s)
		if err != nil {
			return err
		}
		s.debugMessage(fmt.Sprintf("... finished operation %s",op.value))
	}
	return
}

func (s *SemanticManager) EoE() (err error) {
	for i := s.ops.len(); i > 0; i-- {
		op := s.ops.pop()
		s.debugMessage(fmt.Sprintf("Testing operation %s ...",op.value))
		err = op.Perform(s)
		if err != nil {
			//panic(err.Error())
			return err
		}
		s.debugMessage(fmt.Sprintf("... finished operation %s",op.value))
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
		return fmt.Errorf("Operator doesn't have type %#v\n",op1)
	}
	if op2Typ == "" {
		return fmt.Errorf("Operator doesn't have type %#v\n",op2)
	}
	s.debugMessage(fmt.Sprintf("Comparing %s(%s) to %s(%s) for %s",op1.GetValue(), op1Typ, op2.GetValue(), op2Typ,op))
	if op1Typ != op2Typ {
		return fmt.Errorf("Cann't assign operand '%s' (%s) to '%s' (%s) types mismatch",op1.GetValue(), op1Typ, op2.GetValue(), op2Typ)
	}

	data := make(map[string]interface{})
	data["type"] = op1Typ
	value := fmt.Sprintf("%s %s %s",op2.GetValue(),op,op1.GetValue())
	symId := s.st.AddElement(value,"Tvar",data,true)

	s.sas.push(&Tvar_Sar{value:value, typ:op1Typ, scope:s.st.GetScope(), symId:symId})
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
		return fmt.Errorf("Operator doesn't have type %#v\n",op1)
	}
	if op2Typ == "" {
		return fmt.Errorf("Operator doesn't have type %#v\n",op2)
	}
	s.debugMessage(fmt.Sprintf("Comparing %s(%s) to %s(%s)",op1.GetValue(), op1Typ, op2.GetValue(), op2Typ))
	if op1Typ != op2Typ {
		return fmt.Errorf("Cann't assign operand %s(%s) to %s(%s) types mismatch",op1.GetValue(), op1Typ, op2.GetValue(), op2Typ)
	}
	return nil
}
