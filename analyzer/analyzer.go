package analyzer

import (
	"fmt"
	lex "github.com/sbditto85/compiler/lexer"
	tok "github.com/sbditto85/compiler/token"
	sym "github.com/sbditto85/compiler/symbol_table"
	sem "github.com/sbditto85/compiler/analyzer/semantics"
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
	MODIFIER
	PARAMETER
	PARAMETER_LIST
	FIELD_DECLARATION
	METHOD_BODY
	VARIABLE_DECLARATION
	CONSTRUCTOR_DECLARATION
	SEMANTICS
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
	ErrorToString[MODIFIER] = "Modifier"
	ErrorToString[PARAMETER] = "Parameter"
	ErrorToString[PARAMETER_LIST] = "Parameter Lis"
	ErrorToString[FIELD_DECLARATION] = "Field Declaration"
	ErrorToString[METHOD_BODY] = "Method Body"
	ErrorToString[VARIABLE_DECLARATION] = "Variable Declaration"
	ErrorToString[CONSTRUCTOR_DECLARATION] = "Constructor Declaration"
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
func BuildErrMessFromTokErrType(t *tok.Token, e ErrorType) string {
	var str string
	trans := GetErrorToStringMap()
	if tran,ok := trans[e]; ok {
		str = tran
	}
	return fmt.Sprintf("Expected %s, received '%s' on line %d", str, t.Lexeme, t.Linenum + 1)
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
	st *sym.SymbolTable
	sm *sem.SemanticManager
}

func NewAnalyzer(l *lex.Lexer,debug bool) *Analyzer {
	st := sym.NewSymbolTable()
	sm := sem.NewSemanticManager(st, debug)
	a := &Analyzer{lex:l, debug:debug, st:st, sm:sm, pass:1}
	return a
}

func (a *Analyzer) SetLexer(l *lex.Lexer) {
	a.lex = l
}

func (a *Analyzer) debugMessagePassOne(s string) {
	if a.debug && a.pass == 1 {
		fmt.Println(s)
	}
}
func (a *Analyzer) debugMessagePassTwo(s string) {
	if a.debug && a.pass == 2 {
		fmt.Println(s)
	}
}

func (a *Analyzer) AddSymbol(value string, kind string, data map[string]interface{},addTable bool) string {
	//.debugMessagePassOne(fmt.Sprintf("ST: added %s (%s)",value,kind))
	return a.st.AddElement(value,kind,data,addTable)
}

func (a *Analyzer) DropScope() error {
	return a.st.DownScope()
}

func (a *Analyzer) PrintSymbolTable() {
	a.st.PrintTable()
}

func (a *Analyzer) PrintTableInAddOrder() {
	a.st.PrintTableInAddOrder()
}

func (a *Analyzer) GetNext() (*tok.Token,error) {
	curTok,err := a.lex.GetNextToken()
	if curTok.Lexeme == "" {
		a.debugMessagePassOne("Token: ''")
	} else {
		a.debugMessagePassOne("Token: " + curTok.Lexeme)
	}
	return curTok,err
}

func (a *Analyzer) GetCurr() (*tok.Token,error) {
	return a.lex.GetCurrentToken()
}


func (a *Analyzer) Peek() (*tok.Token,error) {
	return a.lex.PeekNextToken()
}

func (a *Analyzer) PerformPass() (err error) {
	defer func(){
		if r:= recover(); r != nil {
			switch r.(type) {
			case string:
				err = fmt.Errorf(r.(string))
			case error:
				err = r.(error)
			}
		}
	}()

	//fmt.Printf("%#v\n",a.st)

	err,_ = a.IsCompilationUnit()

	return err
}

func (a *Analyzer) PerformNextPass(debug bool) error {
	a.pass += 1
	a.debug = debug
	a.sm.SetDebug(debug)
	err := a.PerformPass()
	return err
}

func (a *Analyzer) IsModifier() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is modifier with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Lexeme {
	case "public","private":
		a.GetNext()
	default:
		return BuildErrFromTokErrType(curTok, MODIFIER), MODIFIER
	}
	a.debugMessagePassOne("is modifier!")
	return nil, NONE
}

func (a *Analyzer) IsClassName() (error,ErrorType,string) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is classname with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER, ""
	}
	switch curTok.Type {
	case tok.Identifier:		
		a.GetNext()
	default:
		return BuildErrFromTokErrType(curTok, CLASS_NAME), CLASS_NAME, ""
	}
	a.debugMessagePassOne("is classname!")
	return nil, NONE, curTok.Lexeme
}

func (a *Analyzer) IsType() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is type with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Lexeme {
	case "int","char","bool","void":
		a.GetNext()
	default:
		if err,_,_ := a.IsClassName(); err != nil {
			return BuildErrFromTokErrType(curTok, TYPE), TYPE
		}
	}
	a.debugMessagePassOne("is type!")
	return nil, NONE
}

func (a *Analyzer) IsCompilationUnit() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is compilation unit with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	
	for err == nil {
		curTok,_ = a.GetCurr()
		if curTok.Lexeme == "class"{
			if e,_ := a.IsClassDeclaration(); e != nil {
				panic(BuildErrMessFromTokErrType(curTok, CLASS_DECLARATION))
			}
		} else {
			err = fmt.Errorf("Move along")
		}
	}

	curTok,_ = a.GetCurr()
	if curTok.Lexeme != "void" {
		panic(BuildErrMessFromTok(curTok,"void"))
	}
	curTok,err = a.GetNext()
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}
	
	if curTok.Lexeme != "main" {
		panic(BuildErrMessFromTok(curTok,"main"))
	}
	curTok,err = a.GetNext()
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}
	
	if curTok.Lexeme != "(" {
		panic(BuildErrMessFromTok(curTok,"("))
	}
	curTok,err = a.GetNext()
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}
	
	if curTok.Lexeme != ")" {
		panic(BuildErrMessFromTok(curTok,")"))
	}
	
	//symbol table opperation
	symdata := make(map[string]interface{})
	symdata["type"] = "void"
	a.AddSymbol("main", "Main", symdata, a.pass==1)
	
	curTok,err = a.GetNext()
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}
	
	if e,_ := a.IsMethodBody(); e != nil {
		panic(BuildErrFromTokErrType(curTok, METHOD_BODY))
	}

	a.debugMessagePassOne("is a compliation unit!")
	return nil, NONE
}

func (a *Analyzer) IsClassDeclaration() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is class declaration with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	if curTok.Lexeme != "class" {
		panic(BuildErrMessFromTok(curTok,"class"))
	}
	a.GetNext()
	err,_,className := a.IsClassName()
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, CLASS_DECLARATION))
	}

	//symbol table opperation
	symdata := make(map[string]interface{})
	a.AddSymbol(className, "Class", symdata,a.pass==1)

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


	//symbol table opperation
	//fmt.Println("in class declaration")
	a.DropScope()


	a.GetNext()
	a.debugMessagePassOne("is a class declaration!")
	return nil, NONE
}
 
func (a *Analyzer) IsClassMemberDeclaration() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is class member declaration with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}

	switch curTok.Lexeme {
	case "public", "private":
		modifier := curTok.Lexeme
		if e,t := a.IsModifier(); e != nil {
			return e,t //definition of modifier changed
		}
		curTok,_ = a.GetCurr()
		typ := curTok.Lexeme
		if e,_ := a.IsType(); e != nil {
			curTok,_ = a.GetCurr()
			panic(BuildErrFromTokErrType(curTok, CLASS_MEMBER_DECLARATION))
		}
		curTok,err = a.GetCurr()
		identifier := curTok.Lexeme
		if curTok.Type != tok.Identifier {
			panic(BuildErrFromTokErrType(curTok, CLASS_MEMBER_DECLARATION))
		}
		curTok,err = a.GetNext()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		
		e,_ := a.IsFieldDeclaration(modifier,typ,identifier)
		if e != nil {
			curTok,_ = a.GetCurr()
			panic(BuildErrFromTokErrType(curTok, CLASS_MEMBER_DECLARATION))
		}
	default:
		if e,_ := a.IsConstructorDeclaration(); e != nil {
			curTok,_ = a.GetCurr()
			return BuildErrFromTokErrType(curTok, CLASS_MEMBER_DECLARATION), CLASS_MEMBER_DECLARATION
		}
	}

	a.debugMessagePassOne("is a class member declaration!")
	return nil, NONE
}

func (a *Analyzer) IsFieldDeclaration(modifier string, typ string, identifier string) (e error,et ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is field declaration with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}


	//symbol table operation
	symdata := make(map[string]interface{})
	symdata["accessMod"] = modifier
	symdata["type"] = typ


	switch curTok.Lexeme {
	case "[","=",";":
		isArr := false
		if curTok.Lexeme == "[" {
			curTok, err = a.GetNext()
			if err != nil {
				panic(BuildErrFromTokErrType(curTok, COMPILER))
			}
			if curTok.Lexeme != "]" {
				panic(BuildErrMessFromTok(curTok, "]"))
			}

			isArr = true

			curTok, err = a.GetNext()
			if err != nil {
				panic(BuildErrFromTokErrType(curTok, COMPILER))
			}
		}
		
		//symbol table operation
		symdata["isArray"] = isArr
		a.AddSymbol(identifier, "Ivar", symdata,a.pass==1)

		if curTok.Lexeme == "=" {
			curTok, err = a.GetNext()
			if err != nil {
				panic(BuildErrFromTokErrType(curTok, COMPILER))
			}
			if e,_ := a.IsAssignmentExpression(); e != nil {
				panic(BuildErrFromTokErrType(curTok, ASSIGNMENT_EXPRESSION))
			}
		}

		curTok, _ = a.GetCurr()
		if curTok.Lexeme != ";" {
			panic(BuildErrMessFromTok(curTok, ";"))
		}
		curTok, err = a.GetNext()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}

		a.debugMessagePassOne("is a field declaration!")
		return nil, NONE
	}
	if curTok.Lexeme == "(" {
		curTok, err = a.GetNext()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		_,_,paramList := a.IsParameterList();
		curTok,_ = a.GetCurr()
		if curTok.Lexeme != ")" {
			panic(BuildErrMessFromTok(curTok, ")"))
		}
		curTok, err = a.GetNext()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}


		//symbol table operation
		symdata["parameters"] = paramList
		a.AddSymbol(identifier, "Method", symdata,a.pass==1)


		if e,_ := a.IsMethodBody(); e != nil {
			panic(BuildErrFromTokErrType(curTok, METHOD_BODY))
		}

		a.debugMessagePassOne("is a field declaration!")
		return nil, NONE
	}

	a.debugMessagePassOne("is a field declaration!")
	return BuildErrFromTokErrType(curTok, FIELD_DECLARATION), FIELD_DECLARATION
}

func (a *Analyzer) IsConstructorDeclaration() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is constructor declaration with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}

	e,_,className := a.IsClassName()
	if e != nil {
		curTok,_ = a.GetCurr()
		return BuildErrFromTokErrType(curTok, CONSTRUCTOR_DECLARATION), CONSTRUCTOR_DECLARATION
	}
	curTok,_ = a.GetCurr()
	if curTok.Lexeme != "(" {
		return BuildErrFromTok(curTok, "("), CONSTRUCTOR_DECLARATION
	}
	curTok, err = a.GetNext()
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}

	_,_,paramsList := a.IsParameterList()
	curTok,_ = a.GetCurr()
	if curTok.Lexeme != ")" {
		panic(BuildErrMessFromTok(curTok, ")"))
	}
	curTok, err = a.GetNext()
	if err != nil {
		panic(BuildErrMessFromTokErrType(curTok, COMPILER))
	}

	//symbol table opperation
	symdata := make(map[string]interface{})
	symdata["class"] = className
	symdata["parameters"] = paramsList
	a.AddSymbol(className, "Constructor", symdata, a.pass==1)
	
	if e,t := a.IsMethodBody(); e != nil {
		panic(BuildErrMessFromTokErrType(curTok, t))
	}
	a.debugMessagePassOne("is a constructor declaration!")
	return nil, NONE
}

func (a *Analyzer) IsMethodBody() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is method body with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	
	if curTok.Lexeme != "{" {
		return BuildErrFromTok(curTok, "{"), METHOD_BODY
	}
	curTok, err = a.GetNext()
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}

	for err == nil {
		curTok,err = a.GetCurr()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		err,_ = a.IsVariableDeclaration();
	}
	err = nil
	
	for err == nil {
		curTok,err = a.GetCurr()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		err,_ = a.IsStatement();
	}

	if curTok.Lexeme != "}" {
		panic(BuildErrMessFromTok(curTok, "}"))
	}

	//symbol table opperation
	//fmt.Println("in method body")
	a.DropScope()

	curTok,err = a.GetNext()
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}

	a.debugMessagePassOne("is a method body!")
	return nil, NONE
}

func (a *Analyzer) IsVariableDeclaration() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is variable declaration with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}

	peekTok,_ := a.Peek()
	if peekTok.Type != tok.Identifier {
		return BuildErrFromTokErrType(curTok, VARIABLE_DECLARATION), VARIABLE_DECLARATION
	}

	typ := curTok.Lexeme
	if e,_ := a.IsType(); e != nil {
		curTok,_ = a.GetCurr()
		return BuildErrFromTokErrType(curTok, VARIABLE_DECLARATION), VARIABLE_DECLARATION
		//panic(BuildErrFromTokErrType(curTok, VARIABLE_DECLARATION))
	}
	curTok,err = a.GetCurr()

	identifier := curTok.Lexeme
	if curTok.Type != tok.Identifier {
		return BuildErrFromTokErrType(curTok, VARIABLE_DECLARATION), VARIABLE_DECLARATION
		//panic(BuildErrFromTokErrType(curTok, VARIABLE_DECLARATION))
	}
	curTok,err = a.GetNext()
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}

	isArr := false
	if curTok.Lexeme == "[" {
		curTok, err = a.GetNext()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		if curTok.Lexeme != "]" {
			panic(BuildErrMessFromTok(curTok,"{"))
		}
		isArr = true

		curTok,err = a.GetNext()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
	}

	//symbol table opperation
	symdata := make(map[string]interface{})
	symdata["isArray"] = isArr
	symdata["type"] = typ
	a.AddSymbol(identifier, "Lvar", symdata, a.pass==1)

	curTok,_ = a.GetCurr()
	if curTok.Lexeme == "=" {
		curTok,err = a.GetNext()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		if e,_ := a.IsAssignmentExpression(); e != nil {
			panic(BuildErrFromTokErrType(curTok, VARIABLE_DECLARATION))
		}
	}
	curTok,_ = a.GetCurr()
	if curTok.Lexeme != ";" {
		panic(BuildErrMessFromTok(curTok,";"))
	}
	curTok,err = a.GetNext()
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}
	a.debugMessagePassOne("is a variable declaration!")
	return nil, NONE
}

func (a *Analyzer) IsParameterList() (error,ErrorType,[]sym.Parameter) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is parameter list with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER,[]sym.Parameter{}
	}
	params := make([]sym.Parameter,0)

	e,_,param := a.IsParameter()
	if e != nil {
		return BuildErrFromTokErrType(curTok, PARAMETER_LIST), PARAMETER_LIST, params
	}
	params = append(params,param)
	for err == nil {
		curTok,err = a.GetCurr()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		if curTok.Lexeme == "," {
			a.GetNext()
			e,_,param := a.IsParameter();
			if e != nil {
				panic(BuildErrFromTokErrType(curTok, PARAMETER_LIST))
			}
			params = append(params,param)
		} else {
			err = BuildErrFromTokErrType(curTok, PARAMETER_LIST)
		}
	}

	a.debugMessagePassOne("is a parameter list!")
	return nil, NONE, params
}

func (a *Analyzer) IsParameter() (error,ErrorType,sym.Parameter) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is parameter with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER, sym.Parameter{}
	}
	
	typ := curTok.Lexeme
	if e,_ := a.IsType(); e != nil {
		return BuildErrFromTokErrType(curTok, PARAMETER), PARAMETER, sym.Parameter{}
		//panic(BuildErrFromTokErrType(curTok, PARAMETER))
	}
	curTok,_ = a.GetCurr() //what are the odds an err would occur? like none right? ... guys?
	
	identifier := curTok.Lexeme
	if curTok.Type != tok.Identifier {
		panic(BuildTtErrMessFromTok(curTok, tok.Identifier))
	}
	curTok,err = a.GetNext()
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}
	
	isArr := false
	if curTok.Lexeme == "[" {
		curTok, err = a.GetNext()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		if curTok.Lexeme != "]" {
			panic(BuildErrMessFromTok(curTok,"{"))
		}
		a.GetNext()

		isArr = true
	}
	a.debugMessagePassOne("is a parameter!")
	return nil, NONE, sym.Parameter{Typ:typ,Identifier:identifier,IsArr:isArr}
}

func (a *Analyzer) IsStatement() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is statement with token %s...",curTok.Lexeme))
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
			//Semantic Action
			if a.pass == 2 {
				if err := a.sm.EoE(); err != nil {
					panic(fmt.Sprintf("%s on line %d",err.Error(),curTok.Linenum + 1))
				}
				a.debugMessagePassTwo("EOE")
			}
			a.GetNext()
		} else {
			return BuildErrFromTokErrType(curTok, STATEMENT), STATEMENT
		}
	}
	a.debugMessagePassOne("is a statement!")
	return nil, NONE
}

func (a *Analyzer) IsExpression() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is expression with token %s...",curTok.Lexeme))
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
		//Semantic Action
		if a.pass == 2 {
			a.sm.IPush(curTok.Lexeme, a.st.GetScope())
			a.debugMessagePassTwo(fmt.Sprintf("IPush: %s from scope %s",curTok.Lexeme,a.st.GetScope()))
		}

		a.GetNext()
		if e,t := a.IsFnArrMember(); e != nil && t != FN_ARR_MEMBER {
			panic(e.Error())
		}

		//Semantic Action
		if a.pass == 2 {
			if e := a.sm.IExist(a.st); e != nil {
				panic(fmt.Sprintf("%s on line %d",e.Error(),curTok.Linenum + 1))
			}
			a.debugMessagePassTwo("IExists!");
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
	a.debugMessagePassOne("is expression!");
	return nil, NONE
}

func (a *Analyzer) IsFnArrMember() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is fn arr member with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Lexeme {
	case "(":
		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush(curTok.Lexeme); err != nil {
				panic(fmt.Sprintf("%s on line %d",err.Error(),curTok.Linenum + 1))
			}
			a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s",curTok.Lexeme))

			a.sm.BAL(a.st.GetScope())
			a.debugMessagePassTwo("BAL")
		}

		curTok, err = a.GetNext()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		if curTok.Lexeme != ")" {
			if err,_ := a.IsArgumentList(); err != nil {
				return err, FN_ARR_MEMBER
			}
		}
		//should be pointing at ")"
		curTok,err = a.GetCurr()
		if curTok.Lexeme != ")" {
			panic(BuildErrMessFromTok(curTok, ")"))
		}

		//Semantic Action EAL
		if a.pass == 2 {
			a.sm.EAL(a.st.GetScope())
			a.debugMessagePassTwo("EAL")

			a.sm.Func(a.st.GetScope())
			a.debugMessagePassTwo("func")
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
	a.debugMessagePassOne("is fn arr member!")
	return nil, NONE
}

func (a *Analyzer) IsMemberRefz() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is member refz with token %s...",curTok.Lexeme))
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

	//Semantic Action
	if a.pass == 2 {
		a.sm.IPush(curTok.Lexeme, a.st.GetScope())
		a.debugMessagePassTwo(fmt.Sprintf("IPush: %s from scope %s",curTok.Lexeme,a.st.GetScope()))
	}

	a.GetNext()
	if e,t := a.IsFnArrMember(); e != nil && t != FN_ARR_MEMBER {
		panic(e.Error())
	}

	//Semantic Action
	if a.pass == 2 {
		if err := a.sm.RExist(a.st); err != nil {
			panic(fmt.Sprintf("%s on line %d",err.Error(), curTok.Linenum + 1))
		}
		a.debugMessagePassTwo("RExists!")
	}

	if e,t := a.IsMemberRefz(); e != nil && t != MEMBER_REFZ {
		panic(e.Error())
	}
	a.debugMessagePassOne("is member refz!")
	return nil, NONE
}

func (a *Analyzer) IsExpressionZ() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is expressionz with token %s...",curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Lexeme {
	case "&&","||","==","!=","<=",">=",">","<","+","-","*","/":
		a.GetNext()
		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush(curTok.Lexeme); err != nil {
				panic(fmt.Sprintf("%s on line %d",err.Error(),curTok.Linenum + 1))
			}
			a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s",curTok.Lexeme))
		}
		if err,_ := a.IsExpression(); err != nil {
			panic(err.Error())
		}
	case "=":
		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush("="); err != nil {
				panic(fmt.Sprintf("%s on line %d",err.Error(),curTok.Linenum + 1))
			}
			a.debugMessagePassTwo("Pushed operator =")
		}

		a.GetNext()
		if err,_ := a.IsAssignmentExpression(); err != nil {
			panic(err.Error())
		}
	default:
		return BuildErrFromTokErrType(curTok, EXPRESSIONZ), EXPRESSIONZ
	}
	a.debugMessagePassOne("is expressionz!");
	return nil, NONE
}

func (a *Analyzer) IsAssignmentExpression() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is assignment_expression with token %s...",curTok.Lexeme))
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
	a.debugMessagePassOne("is assignment_expression!");
	return nil, NONE
}

func (a *Analyzer) IsNewDeclaration() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is new declaration with token %s...",curTok.Lexeme))
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
	a.debugMessagePassOne("is new declaration!")
	return nil, NONE
}

func (a *Analyzer) IsArgumentList() (error,ErrorType) {
	curTok,err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is argument list with token %s...",curTok.Lexeme))
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
	a.debugMessagePassOne("is argument list!")
	return nil, NONE
}


