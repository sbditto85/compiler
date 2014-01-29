package analyzer

import (
	"fmt"
	lex "github.com/sbditto85/compiler/lexer"
	tok "github.com/sbditto85/compiler/token"
)

type ErrorType int

const (
	NONE ErrorType = iota
	COMPILER
	STATEMENT
	EXPRESSION
	TYPE
	CLASS_NAME
	EXPRESSIONZ
	FN_ARR_MEMBER
	MEMBER_REFZ
	ASSIGNMENT_EXPRESSION
	NEW_DECLARATION
	ARGUMENT_LIST
)

type Analyzer struct {
	lex *lex.Lexer
	pass int
	debug bool
}

func NewAnalyzer(l *lex.Lexer,debug bool) *Analyzer {
	a := &Analyzer{lex:l,debug:debug}
	return a
}

func (a *Analyzer) debugMessage(s string) {
	if a.debug {
		fmt.Println(s)
	}
}

func (a *Analyzer) GetNext() (*tok.Token,error) {
	curTok,err := a.lex.GetNextToken()
	a.debugMessage("Token: " + curTok.Lexeme)
	return curTok,err
}

func (a *Analyzer) GetCurr() (*tok.Token,error) {
	return a.lex.GetCurrentToken()
}

func (a *Analyzer) PerformNextPass(debug bool) error {
	return nil
}

func (a *Analyzer) IsClassName() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is classname with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Type {
	case tok.Identifier:
		a.GetNext()
	default:
		return fmt.Errorf("Not a classname"), CLASS_NAME
	}
	a.debugMessage("is classname!")
	return nil, NONE
}

func (a *Analyzer) IsType() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is type with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Lexeme {
	case "int","char","bool","void":
		a.GetNext()
	default:
		if err,_ := a.IsClassName(); err != nil {
			return fmt.Errorf("Not a type"), TYPE
		}
	}
	a.debugMessage("is type!")
	return nil, NONE
}

func (a *Analyzer) IsStatement() (error,ErrorType) {
	//TODO: put recover here
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is statement with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch {
	case curTok.Lexeme == "{":
		a.GetNext()
	case curTok.Lexeme == "if":
		a.GetNext()
	case curTok.Lexeme == "while":
		a.GetNext()
	case curTok.Lexeme == "return":
		a.GetNext()
	case curTok.Lexeme == "cout":
		a.GetNext()
	case curTok.Lexeme == "cin":
		a.GetNext()
	default:
		if err,_ :=  a.IsExpression(); err == nil {
			curTok,err = a.GetCurr()
			if curTok.Lexeme != ";" {
				return fmt.Errorf("Expected ';' at new line received %s on line number: %d.",curTok.Lexeme,curTok.Linenum), STATEMENT
			}
		} else {
			return fmt.Errorf("Expected statement line number: %d.",curTok.Linenum), STATEMENT
		}
	}
	a.debugMessage("is a statement!")
	return nil, NONE
}

func (a *Analyzer) IsExpression() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is expression with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch {
	case curTok.Lexeme == "(":
		a.GetNext()
		if e,t := a.IsExpression(); e != nil && t != EXPRESSIONZ {
			return fmt.Errorf("Not an expression"), EXPRESSION
		}
		curTok,err = a.GetCurr() //now at next token after expression
		if curTok.Lexeme == ")" {
			a.GetNext()
		} else {
			return fmt.Errorf("Not an expression"), EXPRESSION
		}
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			return fmt.Errorf("Not an expression"), EXPRESSION
		}
	case curTok.Lexeme == "true":
		a.GetNext()
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			return fmt.Errorf("Not an expression"), EXPRESSION
		}
	case curTok.Lexeme == "false":
		a.GetNext()
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			return fmt.Errorf("Not an expression"), EXPRESSION
		}
	case curTok.Lexeme == "null":
		a.GetNext()
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			return fmt.Errorf("Not an expression"), EXPRESSION
		}
	case curTok.Type == tok.Number:
		a.GetNext()
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			return fmt.Errorf("Not an expression"), EXPRESSION
		}
	case curTok.Type == tok.Character:
		a.GetNext()
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			return fmt.Errorf("Not an expression"), EXPRESSION
		}
	case curTok.Type == tok.Identifier:
		a.GetNext()
		if e,t := a.IsFnArrMember(); e != nil && t != EXPRESSIONZ && t != FN_ARR_MEMBER {
			return fmt.Errorf("Not an expression"), EXPRESSION
		}
		if e,t := a.IsMemberRefz(); e != nil && t != EXPRESSIONZ && t != MEMBER_REFZ {
			return fmt.Errorf("Not an expression"), EXPRESSION
		}
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ  {
			return fmt.Errorf("Not an expression"), EXPRESSION
		}
	default:
		return fmt.Errorf("Not an expression"), EXPRESSION
	}
	a.debugMessage("is expression!");
	return nil, NONE
}

func (a *Analyzer) IsFnArrMember() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is fn arr member with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Lexeme {
	case "(":
		a.GetNext()
		if err,_ := a.IsArgumentList(); err != nil {
			return err, FN_ARR_MEMBER
		}
		//should be pointing at ")"
		curTok,err = a.GetCurr()
		if curTok.Lexeme != ")" {
			return fmt.Errorf("Not a fn arr member"), FN_ARR_MEMBER
		}
		a.GetNext()
	case "[":
		a.GetNext()
		if err,_ := a.IsExpression(); err != nil {
			return fmt.Errorf("Not a fn arr member"), FN_ARR_MEMBER
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != "]" {
			return fmt.Errorf("Not a fn arr member"), FN_ARR_MEMBER
		}
		a.GetNext()
	default:
		return fmt.Errorf("Not a fn arr member"), FN_ARR_MEMBER
	}
	a.debugMessage("is fn arr member!")
	return nil, NONE
}

func (a *Analyzer) IsMemberRefz() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is member refz with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	if curTok.Lexeme != "." {
		return fmt.Errorf("Not a member refz"), MEMBER_REFZ
	}
	a.GetNext()
	curTok,err = a.GetCurr()
	if err != nil {
		return err, COMPILER
	}
	if curTok.Type != tok.Identifier {
		return fmt.Errorf("Not a member refz"), MEMBER_REFZ
	}
	if err,_ := a.IsFnArrMember(); err != nil && err.Error() != "Not a fn arr member" {
		return err, MEMBER_REFZ
	}
	if err,_ := a.IsMemberRefz(); err != nil && err.Error() != "Not a member refz" {
		return err, MEMBER_REFZ
	}
	a.debugMessage("is member refz!")
	return nil, NONE
}

func (a *Analyzer) IsExpressionZ() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is expressionz with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Lexeme {
	case "&&","||","==","!=","<=",">=",">","<","+","-","*","/":
		a.GetNext()
		if err,t := a.IsExpression(); err != nil {
			return err,t
		}
	case "=":
		a.GetNext()
		if err,t := a.IsAssignmentExpression(); err != nil {
			return err,t
		}
	default:
		return fmt.Errorf("Not an expressionz"), EXPRESSIONZ
	}
	a.debugMessage("is expressionz!");
	return nil, NONE
}

func (a *Analyzer) IsAssignmentExpression() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is assignment_expression with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch {
	case curTok.Lexeme == "this":
		a.GetNext()
	case curTok.Lexeme == "new":
		a.GetNext()
		if err,_ := a.IsType(); err != nil {
			return fmt.Errorf("Not an assignment expression"), ASSIGNMENT_EXPRESSION
		}
		if err,_ := a.IsNewDeclaration(); err != nil {
			return fmt.Errorf("Not an assignment expression"), ASSIGNMENT_EXPRESSION
		}
	case curTok.Lexeme == "atoi":
		curTok,err = a.GetNext()
		if curTok.Lexeme != "(" || err != nil {
			return fmt.Errorf("Not an assignment expression"), ASSIGNMENT_EXPRESSION
		}
		curTok,err = a.GetNext()
		if e,_ := a.IsExpression(); e != nil {
			return e, ASSIGNMENT_EXPRESSION
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != ")" || err != nil {
			return fmt.Errorf("Not an assignment expression"), ASSIGNMENT_EXPRESSION
		}
		a.GetNext()
	case curTok.Lexeme == "itoa":
		curTok,err = a.GetNext()
		if curTok.Lexeme != "(" || err != nil {
			return fmt.Errorf("Not an assignment expression"), ASSIGNMENT_EXPRESSION
		}
		curTok,err = a.GetNext()
		if e,_ := a.IsExpression(); e != nil {
			return e, ASSIGNMENT_EXPRESSION
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != ")" || err != nil {
			return fmt.Errorf("Not an assignment expression"), ASSIGNMENT_EXPRESSION
		}
		a.GetNext()
	default:
		if err,_ := a.IsExpression(); err != nil {
			return err, ASSIGNMENT_EXPRESSION
			//return fmt.Errorf("Not an assignment expression")
		}
	}
	a.debugMessage("is assignment_expression!");
	return nil, NONE
}

func (a *Analyzer) IsNewDeclaration() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is new declaration with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Lexeme {
	case "(":
		a.GetNext()
		a.IsArgumentList() //dont care if fails
		//should be pointing at ")"
		curTok,err = a.GetCurr()
		if curTok.Lexeme != ")" {
			return fmt.Errorf("Not a new declaration"), NEW_DECLARATION
		}
		a.GetNext()
	case "[":
		a.GetNext()
		if err,_ := a.IsExpression(); err != nil {
			return fmt.Errorf("Not a new declaration"), NEW_DECLARATION
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != "]" {
			return fmt.Errorf("Not a new declaration"), NEW_DECLARATION
		}
		a.GetNext()
	default:
		return fmt.Errorf("Not a new declaration"), NEW_DECLARATION
	}
	a.debugMessage("is new declaration!")
	return nil, NONE
}

func (a *Analyzer) IsArgumentList() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is argument list with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	if err,_ := a.IsExpression(); err != nil {
		return fmt.Errorf("Not a argument list"), ARGUMENT_LIST
	}
	for err == nil {
		curTok,err = a.GetCurr()
		if err != nil {
			return err, COMPILER
		}
		if curTok.Lexeme != "," {
			break
		}
		a.GetNext()
		err,_ = a.IsExpression()
	}
	a.debugMessage("is argument list!")
	return nil, NONE
}


