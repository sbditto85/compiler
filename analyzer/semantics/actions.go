package semantics

import (
	"fmt"
	tok "github.com/sbditto85/compiler/token"
	sym "github.com/sbditto85/compiler/symbol_table"
)

func (s *SemanticManager) IPush(value, scope string, curTok *tok.Token) {
	s.sas.push(&Id_Sar{value:value, scope:scope, token:curTok})
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
	tok := sar.GetToken()
	return fmt.Errorf("%s does not exist on line %d\n",tok.Lexeme,tok.Linenum)
}

func (s *SemanticManager) OPush(value string, precIn, precOn int) (err error) {
	topOp := s.ops.topElement()
	for s.ops.len() != 0 && topOp.precOn >= precIn {
		op := s.ops.pop()
		err = op.Perform(s)
		if err != nil {
			return
		}
		topOp = s.ops.topElement()
	}
	s.ops.push(&Operator{value:value, precIn:precIn, precOn:precOn})
	return
}

func (s *SemanticManager) EoE() (err error) {
	for i := s.ops.len(); i > 0; i-- {
		op := s.ops.pop()
		s.debugMessage(fmt.Sprintf("Testing operation %s ...",op.value))
		err = op.Perform(s)
		if err != nil {
			return err
		}
		s.debugMessage(fmt.Sprintf("... finished operation %s",op.value))
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
	if op1Typ == "" {
		return fmt.Errorf("Operator doesn't have type %#v\n",op1)
	}
	if op2Typ == "" {
		return fmt.Errorf("Operator doesn't have type %#v\n",op2)
	}
	if op1Typ != op2Typ {
		return fmt.Errorf("Cann't assign operand %s(%s) to %s(%s) types mismatch",op1.GetValue(), op1Typ, op2.GetValue(), op2Typ)
	}
	return nil
}
