package analyzer

import (
	"fmt"
	lex "github.com/sbditto85/compiler/lexer"
	tok "github.com/sbditto85/compiler/token"
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

func (a *Analyzer) IsClassName() error {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is classname with token %s...",curTok.Lexeme))
	if err != nil {
		return err
	}
	switch curTok.Type {
	case tok.Identifier:
		a.GetNext()
	default:
		return fmt.Errorf("Not a classname")
	}
	a.debugMessage("is classname!")
	return nil
}

func (a *Analyzer) IsType() error {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is type with token %s...",curTok.Lexeme))
	if err != nil {
		return err
	}
	switch curTok.Lexeme {
	case "int","char","bool","void":
		a.GetNext()
	default:
		if err := a.IsClassName(); err != nil {
			return fmt.Errorf("Not a type")
		}
	}
	a.debugMessage("is type!")
	return nil
}

func (a *Analyzer) IsStatement() error {
	//TODO: put recover here
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is statement with token %s...",curTok.Lexeme))
	if err != nil {
		return err
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
		if a.IsExpression() == nil {
			curTok,err = a.GetCurr()
			if curTok.Lexeme != ";" {
				return fmt.Errorf("Expected ';' at new line received %s on line number: %d.",curTok.Lexeme,curTok.Linenum)
			}
		} else {
			return fmt.Errorf("Expected statement line number: %d.",curTok.Linenum)
		}
	}
	a.debugMessage("is a statement!")
	return nil
}

func (a *Analyzer) IsExpression() error {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is expression with token %s...",curTok.Lexeme))
	if err != nil {
		return err
	}
	switch {
	case curTok.Lexeme == "(":
		a.GetNext()
		if e := a.IsExpression(); e != nil && e.Error() != "Not an expression" {
			return fmt.Errorf("Not an expression")
		}
		curTok,err = a.GetCurr() //now at next token after expression
		if curTok.Lexeme == ")" {
			a.GetNext()
		} else {
			return fmt.Errorf("Not an expression")
		}
		if e := a.IsExpressionZ(); e != nil && e.Error() != "Not an expressionz" {
			return fmt.Errorf("Not an expression")
		}
	case curTok.Lexeme == "true":
		a.GetNext()
		if e := a.IsExpressionZ(); e != nil && e.Error() != "Not an expressionz" {
			return fmt.Errorf("Not an expression")
		}
	case curTok.Lexeme == "false":
		a.GetNext()
		if e := a.IsExpressionZ(); e != nil && e.Error() != "Not an expressionz" {
			return fmt.Errorf("Not an expression")
		}
	case curTok.Lexeme == "null":
		a.GetNext()
		if e := a.IsExpressionZ(); e != nil && e.Error() != "Not an expressionz" {
			return fmt.Errorf("Not an expression")
		}
	case curTok.Type == tok.Number:
		a.GetNext()
		if e := a.IsExpressionZ(); e != nil && e.Error() != "Not an expressionz" {
			return fmt.Errorf("Not an expression")
		}
	case curTok.Type == tok.Character:
		a.GetNext()
		if e := a.IsExpressionZ(); e != nil && e.Error() != "Not an expressionz" {
			return fmt.Errorf("Not an expression")
		}
	case curTok.Type == tok.Identifier:
		a.GetNext()
		if e := a.IsFnArrMember(); e != nil && e.Error() != "Not a fn arr member" {
			return fmt.Errorf("Not an expression")
		}
		if e := a.IsMemberRefz(); e != nil && e.Error() != "Not a member refz" {
			return fmt.Errorf("Not an expression")
		}
		if e := a.IsExpressionZ(); e != nil && e.Error() != "Not an expressionz" {
			return fmt.Errorf("Not an expression")
		}
	default:
		return fmt.Errorf("Not an expression")
	}
	a.debugMessage("is expression!");
	return nil
}

func (a *Analyzer) IsFnArrMember() error {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is fn arr member with token %s...",curTok.Lexeme))
	if err != nil {
		return err
	}
	switch curTok.Lexeme {
	case "(":
		a.GetNext()
		if err := a.IsArgumentList(); err != nil {
			return err
		}
		//should be pointing at ")"
		curTok,err = a.GetCurr()
		if curTok.Lexeme != ")" {
			return fmt.Errorf("Not a fn arr member")
		}
		a.GetNext()
	case "[":
		a.GetNext()
		if err := a.IsExpression(); err != nil {
			return fmt.Errorf("Not a fn arr member")
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != "]" {
			return fmt.Errorf("Not a fn arr member")
		}
		a.GetNext()
	default:
		return fmt.Errorf("Not a fn arr member")
	}
	a.debugMessage("is fn arr member!")
	return nil
}

func (a *Analyzer) IsMemberRefz() error {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is member refz with token %s...",curTok.Lexeme))
	if err != nil {
		return err
	}
	if curTok.Lexeme != "." {
		return fmt.Errorf("Not a member refz")
	}
	a.GetNext()
	curTok,err = a.GetCurr()
	if curTok.Type != tok.Identifier && err != nil {
		return fmt.Errorf("Not a member refz")
	}
	if err := a.IsFnArrMember(); err != nil && err.Error() != "Not a fn arr member" {
		return err
	}
	if err := a.IsMemberRefz(); err != nil && err.Error() != "Not a member refz" {
		return err
	}
	a.debugMessage("is member refz!")
	return nil
}

func (a *Analyzer) IsExpressionZ() error {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is expressionz with token %s...",curTok.Lexeme))
	if err != nil {
		return err
	}
	switch curTok.Lexeme {
	case "&&","||","==","!=","<=",">=",">","<","+","-","*","/":
		a.GetNext()
		if err = a.IsExpression(); err != nil {
			return err
		}
	case "=":
		a.GetNext()
		if err = a.IsAssignmentExpression(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Not an expressionz")
	}
	a.debugMessage("is expressionz!");
	return nil
}

func (a *Analyzer) IsAssignmentExpression() error {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is assignment_expression with token %s...",curTok.Lexeme))
	if err != nil {
		return err
	}
	switch {
	case curTok.Lexeme == "this":
		a.GetNext()
	case curTok.Lexeme == "new":
		a.GetNext()
		if err := a.IsType(); err != nil {
			return fmt.Errorf("Not an assignment expression")
		}
		if err := a.IsNewDeclaration(); err != nil {
			return fmt.Errorf("Not an assignment expression")
		}
	case curTok.Lexeme == "atoi":
		curTok,err = a.GetNext()
		if curTok.Lexeme != "(" || err != nil {
			return fmt.Errorf("Not an assignment expression")
		}
		curTok,err = a.GetNext()
		if e := a.IsExpression(); e != nil {
			return e;
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != ")" || err != nil {
			return fmt.Errorf("Not an assignment expression")
		}
		a.GetNext()
	case curTok.Lexeme == "itoa":
		curTok,err = a.GetNext()
		if curTok.Lexeme != "(" || err != nil {
			return fmt.Errorf("Not an assignment expression")
		}
		curTok,err = a.GetNext()
		if e := a.IsExpression(); e != nil {
			return e;
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != ")" || err != nil {
			return fmt.Errorf("Not an assignment expression")
		}
		a.GetNext()
	default:
		if err := a.IsExpression(); err != nil {
			return err
			//return fmt.Errorf("Not an assignment expression")
		}
	}
	a.debugMessage("is assignment_expression!");
	return nil
}

func (a *Analyzer) IsNewDeclaration() error {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is new declaration with token %s...",curTok.Lexeme))
	if err != nil {
		return err
	}
	switch curTok.Lexeme {
	case "(":
		a.GetNext()
		a.IsArgumentList() //dont care if fails
		//should be pointing at ")"
		curTok,err = a.GetCurr()
		if curTok.Lexeme != ")" {
			return fmt.Errorf("Not a new declaration")
		}
		a.GetNext()
	case "[":
		a.GetNext()
		if err := a.IsExpression(); err != nil {
			return fmt.Errorf("Not a new declaration")
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != "]" {
			return fmt.Errorf("Not a new declaration")
		}
		a.GetNext()
	default:
		return fmt.Errorf("Not a new declaration")
	}
	a.debugMessage("is new declaration!")
	return nil
}

func (a *Analyzer) IsArgumentList() error {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is argument list with token %s...",curTok.Lexeme))
	if err != nil {
		return err
	}
	if err := a.IsExpression(); err != nil {
		return fmt.Errorf("Not a argument list")
	}
	for err == nil {
		curTok,err = a.GetCurr()
		if curTok.Lexeme != "," {
			break
		}
		a.GetNext()
		err = a.IsExpression()
	}
	a.debugMessage("is argument list!")
	return nil
}


