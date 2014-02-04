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
	CLASS_DECLARATION
	CLASS_MEMBER_DECLARATION
)
func GetErrorToStringMap() map[ErrorType]string {
	ErrorToString := make(map[ErrorType]string)
	ErrorToString[NONE] = "None"
	ErrorToString[COMPILER] = "Compiler"
	ErrorToString[STATEMENT] = "Statement"
	ErrorToString[EXPRESSION] = "Expression"
	ErrorToString[TYPE] = "Type"
	ErrorToString[CLASS_NAME] = "Class Name"
	ErrorToString[EXPRESSIONZ] = "Expressionz"
	ErrorToString[FN_ARR_MEMBER] = "Function Arg Member"
	ErrorToString[MEMBER_REFZ] = "Member Refz"
	ErrorToString[ASSIGNMENT_EXPRESSION] = "Assignment Expression"
	ErrorToString[NEW_DECLARATION] = "New Declaration"
	ErrorToString[ARGUMENT_LIST] = "Argument List"
	ErrorToString[CLASS_DECLARATION] = "Class Declaration"
	ErrorToString[CLASS_MEMBER_DECLARATION] = "Class Member Declaration"
	return ErrorToString
}
func BuildErrFromTokErrType(t *tok.Token, e ErrorType) error {
	var str string
	trans := GetErrorToStringMap()
	if tran,ok := trans[e]; ok {
		str = tran
	}
	return fmt.Errorf("Expected %s, received '%s' on line %d", str, t.Lexeme, t.Linenum + 1)
}
func BuildErrFromTok(t *tok.Token, expected string) error {
	return fmt.Errorf(BuildErrMessFromTok(t,expected))
}
func BuildErrMessFromTok(t *tok.Token, expected string) string {
	return fmt.Sprintf("Expected '%s', received '%s' on line %d", expected, t.Lexeme, t.Linenum + 1)
}
func BuildTtErrMessFromTok(t *tok.Token, tt tok.TokenType) string {
	var str string
	trans := tok.GetTokToStringMap()
	if tran,ok := trans[tt]; ok {
		str = tran
	}
	return fmt.Sprintf("Expected %s, received '%s' on line %d", str, t.Lexeme, t.Linenum + 1)
}

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
	if curTok.Lexeme == "" {
		a.debugMessage("Token: ''")
	} else {
		a.debugMessage("Token: " + curTok.Lexeme)
	}
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
		return BuildErrFromTokErrType(curTok, CLASS_NAME), CLASS_NAME
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
			return BuildErrFromTokErrType(curTok, TYPE), TYPE
		}
	}
	a.debugMessage("is type!")
	return nil, NONE
}

func (a *Analyzer) IsClassDeclaration() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is class declaration with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	if curTok.Lexeme != "class" {
		panic(BuildErrMessFromTok(curTok,"class"))
	}
	a.GetNext()
	if err,_ := a.IsClassName(); err != nil {
		panic(BuildErrFromTokErrType(curTok, CLASS_DECLARATION))
	}
	curTok,err = a.GetCurr()
	if curTok.Lexeme != "{" {
		panic(BuildErrMessFromTok(curTok,"{"))
	}
	a.GetNext()
	
	for err == nil {
		err,_ = a.IsClassMemberDeclaration()
	}
	curTok,err = a.GetCurr()
	if curTok.Lexeme != "}" {
		panic(BuildErrMessFromTok(curTok,"}"))
	}
	a.GetNext()
	a.debugMessage("is a class declaration!")
	return nil, NONE
}
 
func (a *Analyzer) IsClassMemberDeclaration() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is class member declaration with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}

	

	a.debugMessage("is a class member declaration!")
	return fmt.Errorf("i is a fake error"), NONE
}

func (a *Analyzer) IsStatement() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessage(fmt.Sprintf("Testing is statement with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch {
	case curTok.Lexeme == "{":
		a.GetNext()
		for err == nil {
			err,_ =  a.IsStatement();
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != "}" {
			panic(BuildErrMessFromTok(curTok,"}"))
		}
		a.GetNext()
	case curTok.Lexeme == "if":
		a.GetNext()
		curTok,err = a.GetCurr() //now at next token after expression
		if curTok.Lexeme != "(" {
			panic(BuildErrMessFromTok(curTok,"("))
		}
		a.GetNext()
		if err,_ :=  a.IsExpression(); err == nil {
			curTok,err = a.GetCurr()
			if curTok.Lexeme != ")" {
				panic(BuildErrMessFromTok(curTok, ")"))
			}
			a.GetNext()
			if err,_ :=  a.IsStatement(); err != nil {
				panic(BuildErrFromTokErrType(curTok, STATEMENT))
			}
			curTok,err = a.GetCurr()
			if curTok.Lexeme == "else" {
				a.GetNext()
				if err,_ :=  a.IsStatement(); err != nil {
					panic(BuildErrFromTokErrType(curTok, STATEMENT))
				}
			}
		} else {
			return BuildErrFromTokErrType(curTok, STATEMENT), STATEMENT
		}
	case curTok.Lexeme == "while":
		a.GetNext()
		curTok,err = a.GetCurr() //now at next token after expression
		if curTok.Lexeme != "(" {
			panic(BuildErrMessFromTok(curTok,"("))
		}
		a.GetNext()
		if err,_ :=  a.IsExpression(); err == nil {
			curTok,err = a.GetCurr()
			if curTok.Lexeme != ")" {
				panic(BuildErrMessFromTok(curTok, ")"))
			}
			a.GetNext()
			if err,_ :=  a.IsStatement(); err != nil {
				panic(BuildErrFromTokErrType(curTok, STATEMENT))
			}
		} else {
			return BuildErrFromTokErrType(curTok, STATEMENT), STATEMENT
		}
	case curTok.Lexeme == "return":
		a.GetNext()
		if err,_ :=  a.IsExpression(); err == nil {
			//TODO: do anything about this?
		}
		curTok,err = a.GetCurr()
		if curTok.Lexeme != ";" {
			panic(BuildErrMessFromTok(curTok, ";"))
		}
		a.GetNext()
	case curTok.Lexeme == "cout":
		a.GetNext()
		curTok,err = a.GetCurr() //now at next token after expression
		if curTok.Lexeme != "<<" {
			panic(BuildErrMessFromTok(curTok,"<<"))
		}
		a.GetNext()
		if err,_ :=  a.IsExpression(); err == nil {
			curTok,err = a.GetCurr()
			if curTok.Lexeme != ";" {
				panic(BuildErrMessFromTok(curTok, ";"))
			}
			a.GetNext()
		} else {
			return BuildErrFromTokErrType(curTok, STATEMENT), STATEMENT
		}
	case curTok.Lexeme == "cin":
		a.GetNext()
		curTok,err = a.GetCurr() //now at next token after expression
		if curTok.Lexeme != ">>" {
			panic(BuildErrMessFromTok(curTok,">>"))
		}
		a.GetNext()
		if err,_ :=  a.IsExpression(); err == nil {
			curTok,err = a.GetCurr()
			if curTok.Lexeme != ";" {
				panic(BuildErrMessFromTok(curTok, ";"))
			}
			a.GetNext()
		} else {
			return BuildErrFromTokErrType(curTok, STATEMENT), STATEMENT
		}
	default:
		if err,_ :=  a.IsExpression(); err == nil {
			curTok,err = a.GetCurr()
			if curTok.Lexeme != ";" {
				panic(BuildErrMessFromTok(curTok, ";"))
			}
			a.GetNext()
		} else {
			return BuildErrFromTokErrType(curTok, STATEMENT), STATEMENT
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
		if e,_ := a.IsExpression(); e != nil {
			panic(e.Error())
		}
		curTok,err = a.GetCurr() //now at next token after expression
		if curTok.Lexeme == ")" {
			a.GetNext()
		} else {
			panic(BuildErrMessFromTok(curTok,")"))
		}
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			panic(e.Error())
		}
	case curTok.Lexeme == "true":
		a.GetNext()
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			panic(e.Error())
		}
	case curTok.Lexeme == "false":
		a.GetNext()
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			panic(e.Error())
		}
	case curTok.Lexeme == "null":
		a.GetNext()
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			panic(e.Error())
		}
	case curTok.Type == tok.Number:
		a.GetNext()
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			panic(e.Error())
		}
	case curTok.Type == tok.Character:
		a.GetNext()
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			panic(e.Error())
		}
	case curTok.Type == tok.Identifier:
		a.GetNext()
		if e,t := a.IsFnArrMember(); e != nil && t != FN_ARR_MEMBER {
			panic(e.Error())
		}
		if e,t := a.IsMemberRefz(); e != nil && t != MEMBER_REFZ {
			panic(e.Error())
		}
		if e,t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ  {
			panic(e.Error())
		}
	default:
		return BuildErrFromTokErrType(curTok, EXPRESSION), EXPRESSION
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
			panic(BuildErrMessFromTok(curTok, ")"))
		}
		a.GetNext()
	case "[":
		a.GetNext()
		if e,_ := a.IsExpression(); err != nil {
			panic(e.Error())
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != "]" {
			panic(BuildErrMessFromTok(curTok, "]"))
		}
		a.GetNext()
	default:
		return BuildErrFromTokErrType(curTok, FN_ARR_MEMBER), FN_ARR_MEMBER
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
		return BuildErrFromTokErrType(curTok, MEMBER_REFZ), MEMBER_REFZ
	}
	a.GetNext()
	curTok,err = a.GetCurr()
	if err != nil {
		return err, COMPILER
	}
	if curTok.Type != tok.Identifier {
		panic( BuildTtErrMessFromTok(curTok, tok.Identifier))
	}
	if e,t := a.IsFnArrMember(); e != nil && t != FN_ARR_MEMBER {
		panic(e.Error())
	}
	if e,t := a.IsMemberRefz(); e != nil && t != MEMBER_REFZ {
		panic(e.Error())
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
		if err,_ := a.IsExpression(); err != nil {
			panic(err.Error())
		}
	case "=":
		a.GetNext()
		if err,_ := a.IsAssignmentExpression(); err != nil {
			panic(err.Error())
		}
	default:
		return BuildErrFromTokErrType(curTok, EXPRESSIONZ), EXPRESSIONZ
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
			panic(err.Error())
		}
		if err,_ := a.IsNewDeclaration(); err != nil {
			panic(err.Error())
		}
	case curTok.Lexeme == "atoi":
		curTok,err = a.GetNext()
		if curTok.Lexeme != "(" || err != nil {
			panic(BuildErrMessFromTok(curTok,"("))
		}
		curTok,err = a.GetNext()
		if e,_ := a.IsExpression(); e != nil {
			panic(e.Error())
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != ")" || err != nil {
			panic(BuildErrMessFromTok(curTok,")"))
		}
		a.GetNext()
	case curTok.Lexeme == "itoa":
		curTok,err = a.GetNext()
		if curTok.Lexeme != "(" || err != nil {
			panic(BuildErrMessFromTok(curTok,"("))
		}
		curTok,err = a.GetNext()
		if e,_ := a.IsExpression(); e != nil {
			panic(e.Error())
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != ")" || err != nil {
			panic(BuildErrMessFromTok(curTok,")"))
		}
		a.GetNext()
	default:
		if err,_ := a.IsExpression(); err != nil {
			return err, ASSIGNMENT_EXPRESSION
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
			panic(BuildErrMessFromTok(curTok,")"))
		}
		a.GetNext()
	case "[":
		a.GetNext()
		if err,_ := a.IsExpression(); err != nil {
			panic(err.Error())
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != "]" {
			panic(BuildErrMessFromTok(curTok,"]"))
		}
		a.GetNext()
	default:
		return BuildErrFromTokErrType(curTok, NEW_DECLARATION), NEW_DECLARATION
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
	if e,t := a.IsExpression(); e != nil {
		return e,t
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
		if e,_ := a.IsExpression(); e != nil {
			panic(e.Error())
		}
	}
	a.debugMessage("is argument list!")
	return nil, NONE
}


