package analyzer

import (
	"fmt"
	ic "github.com/sbditto85/compiler/analyzer/icode"
	sem "github.com/sbditto85/compiler/analyzer/semantics"
	lex "github.com/sbditto85/compiler/lexer"
	sym "github.com/sbditto85/compiler/symbol_table"
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
	MODIFIER
	PARAMETER
	PARAMETER_LIST
	FIELD_DECLARATION
	METHOD_BODY
	VARIABLE_DECLARATION
	CONSTRUCTOR_DECLARATION
	SEMANTICS
)

var ErrorToString map[ErrorType]string = nil

func GetErrorToStringMap() map[ErrorType]string {
	if ErrorToString == nil {
		ErrorToString = make(map[ErrorType]string)
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
	}
	return ErrorToString
}
func BuildErrFromTokErrType(t *tok.Token, e ErrorType) error {
	var str string
	trans := GetErrorToStringMap()
	if tran, ok := trans[e]; ok {
		str = tran
	}
	return fmt.Errorf("Expected %s, received '%s' on line %d", str, t.Lexeme, t.Linenum+1)
}
func BuildErrMessFromTokErrType(t *tok.Token, e ErrorType) string {
	var str string
	trans := GetErrorToStringMap()
	if tran, ok := trans[e]; ok {
		str = tran
	}
	return fmt.Sprintf("Expected %s, received '%s' on line %d", str, t.Lexeme, t.Linenum+1)
}
func BuildErrFromTok(t *tok.Token, expected string) error {
	return fmt.Errorf(BuildErrMessFromTok(t, expected))
}
func BuildErrMessFromTok(t *tok.Token, expected string) string {
	return fmt.Sprintf("Expected '%s', received '%s' on line %d", expected, t.Lexeme, t.Linenum+1)
}
func BuildTtErrMessFromTok(t *tok.Token, tt tok.TokenType) string {
	var str string
	trans := tok.GetTokToStringMap()
	if tran, ok := trans[tt]; ok {
		str = tran
	}
	return fmt.Sprintf("Expected %s, received '%s' on line %d", str, t.Lexeme, t.Linenum+1)
}

type Analyzer struct {
	lex   *lex.Lexer
	pass  int
	debug bool
	st    *sym.SymbolTable
	sm    *sem.SemanticManager
	gen   *ic.Generator
}

func NewAnalyzer(l *lex.Lexer, debug bool) *Analyzer {
	st := sym.NewSymbolTable()
	gen := ic.NewGenerator(st)
	sm := sem.NewSemanticManager(st, l, gen, debug)
	a := &Analyzer{lex: l, debug: debug, st: st, sm: sm, pass: 1, gen: gen}
	return a
}

func (a *Analyzer) GetICodeInfo() (*ic.Quad, *sym.SymbolTable) {
	return a.gen.GetQuad(), a.st
}

func (a *Analyzer) PrintQuadTable() {
	a.gen.PrintQuadTable()
}

func (a *Analyzer) PrintQuadStatic() {
	a.gen.PrintQuadStatic()
}

func (a *Analyzer) SetLexer(l *lex.Lexer) {
	a.lex = l
	a.sm.SetLexer(l)
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

func (a *Analyzer) AddSymbol(value string, kind string, data map[string]interface{}, addTable bool) string {
	//.debugMessagePassOne(fmt.Sprintf("ST: added %s (%s)",value,kind))
	return a.st.AddElement(value, kind, data, addTable)
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

func (a *Analyzer) GetNext(expectSymbol bool) (*tok.Token, error) {
	curTok, err := a.lex.GetNextToken(expectSymbol)
	if curTok.Lexeme == "" {
		a.debugMessagePassOne("Token: ''")
	} else {
		a.debugMessagePassOne("Token: " + curTok.Lexeme)
	}
	return curTok, err
}

func (a *Analyzer) GetCurr() (*tok.Token, error) {
	return a.lex.GetCurrentToken()
}

func (a *Analyzer) Peek() (*tok.Token, error) {
	return a.lex.PeekNextToken()
}

func (a *Analyzer) PerformPass() (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case string:
				err = fmt.Errorf(r.(string))
			case error:
				err = r.(error)
			}
		}
	}()

	//fmt.Printf("%#v\n",a.st)

	err, _ = a.IsCompilationUnit()

	return err
}

func (a *Analyzer) PerformNextPass(debug bool) error {
	a.pass += 1
	a.debug = debug
	a.sm.SetDebug(debug)
	err := a.PerformPass()
	return err
}

func (a *Analyzer) IsModifier() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is modifier with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Lexeme {
	case "public", "private":
		a.GetNext(false)
	default:
		return BuildErrFromTokErrType(curTok, MODIFIER), MODIFIER
	}
	a.debugMessagePassOne("is modifier!")
	return nil, NONE
}

func (a *Analyzer) IsClassName() (error, ErrorType, string) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is classname with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER, ""
	}
	switch curTok.Type {
	case tok.Identifier:
		a.GetNext(false)
	default:
		return BuildErrFromTokErrType(curTok, CLASS_NAME), CLASS_NAME, ""
	}
	a.debugMessagePassOne("is classname!")
	return nil, NONE, curTok.Lexeme
}

func (a *Analyzer) IsType() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is type with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Lexeme {
	case "int", "char", "bool", "void":
		a.GetNext(false)
	default:
		if err, _, _ := a.IsClassName(); err != nil {
			return BuildErrFromTokErrType(curTok, TYPE), TYPE
		}
	}

	//Semantic Action EAL
	if a.pass == 2 {
		a.sm.TPush(curTok.Lexeme, a.st.GetScope())
		a.debugMessagePassTwo(fmt.Sprintf("Type %s pushed", curTok.Lexeme))
	}

	a.debugMessagePassOne("is type!")
	return nil, NONE
}

func (a *Analyzer) IsCompilationUnit() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is compilation unit with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}

	for err == nil {
		curTok, _ = a.GetCurr()
		if curTok.Lexeme == "class" {
			if e, _ := a.IsClassDeclaration(); e != nil {
				panic(BuildErrMessFromTokErrType(curTok, CLASS_DECLARATION))
			}
		} else {
			err = fmt.Errorf("Move along")
		}
	}

	curTok, _ = a.GetCurr()
	if curTok.Lexeme != "void" {
		panic(BuildErrMessFromTok(curTok, "void"))
	}
	curTok, err = a.GetNext(false)
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}

	if curTok.Lexeme != "main" {
		panic(BuildErrMessFromTok(curTok, "main"))
	}
	curTok, err = a.GetNext(false)
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}

	if curTok.Lexeme != "(" {
		panic(BuildErrMessFromTok(curTok, "("))
	}
	curTok, err = a.GetNext(false)
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}

	if curTok.Lexeme != ")" {
		panic(BuildErrMessFromTok(curTok, ")"))
	}

	//symbol table opperation
	symdata := make(map[string]interface{})
	symdata["type"] = "void"
	a.AddSymbol("main", "Main", symdata, a.pass == 1)

	curTok, err = a.GetNext(false)
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}

	if e, _ := a.IsMethodBody(false); e != nil {
		panic(BuildErrFromTokErrType(curTok, METHOD_BODY))
	}

	a.debugMessagePassOne("is a compliation unit!")
	return nil, NONE
}

func (a *Analyzer) IsClassDeclaration() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is class declaration with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	if curTok.Lexeme != "class" {
		panic(BuildErrMessFromTok(curTok, "class"))
	}
	a.GetNext(false)
	err, _, className := a.IsClassName()
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, CLASS_DECLARATION))
	}

	//symbol table opperation
	symdata := make(map[string]interface{})
	a.AddSymbol(className, "Class", symdata, a.pass == 1)

	//symbol table action (for icode)
	data := make(map[string]interface{})
	data["type"] = className
	data["accessMod"] = "private"
	data["scope"] = "g." + className
	a.AddSymbol(className+"StaticInit", "StaticInit", data, a.pass == 2)

	curTok, err = a.GetCurr()
	if curTok.Lexeme != "{" {
		panic(BuildErrMessFromTok(curTok, "{"))
	}
	a.GetNext(false)

	for err == nil {
		err, _ = a.IsClassMemberDeclaration()
	}

	curTok, err = a.GetCurr()
	if curTok.Lexeme != "}" {
		panic(BuildErrMessFromTok(curTok, "}"))
	}

	//Semantic Action (icode)
	if a.pass == 2 {
		a.sm.StaticInit(className, a.st)
	}

	//symbol table opperation
	//fmt.Println("in class declaration")
	a.DropScope()

	a.GetNext(false)
	a.debugMessagePassOne("is a class declaration!")
	return nil, NONE
}

func (a *Analyzer) IsClassMemberDeclaration() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is class member declaration with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}

	switch curTok.Lexeme {
	case "public", "private":
		modifier := curTok.Lexeme
		if e, t := a.IsModifier(); e != nil {
			return e, t //definition of modifier changed
		}
		curTok, _ = a.GetCurr()
		typ := curTok.Lexeme
		if e, _ := a.IsType(); e != nil {
			curTok, _ = a.GetCurr()
			panic(BuildErrFromTokErrType(curTok, CLASS_MEMBER_DECLARATION))
		}

		//Semantic Action
		if a.pass == 2 {
			if err := a.sm.TExists(a.st); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("TExists!")
		}

		curTok, err = a.GetCurr()
		identifier := curTok.Lexeme
		if curTok.Type != tok.Identifier {
			panic(BuildErrFromTokErrType(curTok, CLASS_MEMBER_DECLARATION))
		}
		curTok, err = a.GetNext(false)
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}

		e, _ := a.IsFieldDeclaration(modifier, typ, identifier)
		if e != nil {
			curTok, _ = a.GetCurr()
			panic(BuildErrFromTokErrType(curTok, CLASS_MEMBER_DECLARATION))
		}
	default:
		if e, _ := a.IsConstructorDeclaration(); e != nil {
			curTok, _ = a.GetCurr()
			return BuildErrFromTokErrType(curTok, CLASS_MEMBER_DECLARATION), CLASS_MEMBER_DECLARATION
		}
	}

	a.debugMessagePassOne("is a class member declaration!")
	return nil, NONE
}

func (a *Analyzer) IsFieldDeclaration(modifier string, typ string, identifier string) (e error, et ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is field declaration with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}

	//symbol table operation
	symdata := make(map[string]interface{})
	symdata["accessMod"] = modifier
	symdata["type"] = typ

	switch curTok.Lexeme {
	case "[", "=", ";":
		isArr := false
		if curTok.Lexeme == "[" {
			curTok, err = a.GetNext(false)
			if err != nil {
				panic(BuildErrFromTokErrType(curTok, COMPILER))
			}
			if curTok.Lexeme != "]" {
				panic(BuildErrMessFromTok(curTok, "]"))
			}

			isArr = true

			curTok, err = a.GetNext(false)
			if err != nil {
				panic(BuildErrFromTokErrType(curTok, COMPILER))
			}
		}

		//Semantic Action
		if a.pass == 2 {
			a.sm.VPush(identifier, a.st.GetScope(), typ)
			a.debugMessagePassTwo(fmt.Sprintf("vPush %s (%s)", identifier, typ))
		}

		//symbol table operation
		symdata["isArray"] = isArr
		a.AddSymbol(identifier, "Ivar", symdata, a.pass == 1)

		if curTok.Lexeme == "=" {

			//Semantic Action OPush
			if a.pass == 2 {
				if err := a.sm.OPush(curTok.Lexeme); err != nil {
					panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
				}
				a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s", curTok.Lexeme))
			}

			curTok, err = a.GetNext(false)
			if err != nil {
				panic(BuildErrFromTokErrType(curTok, COMPILER))
			}
			if e, _ := a.IsAssignmentExpression(); e != nil {
				panic(BuildErrFromTokErrType(curTok, ASSIGNMENT_EXPRESSION))
			}
		}

		curTok, _ = a.GetCurr()
		if curTok.Lexeme != ";" {
			panic(BuildErrMessFromTok(curTok, ";"))
		}

		//Semantic Action
		if a.pass == 2 {
			if err := a.sm.EoE(); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("EOE")
		}

		curTok, err = a.GetNext(false)
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}

		a.debugMessagePassOne("is a field declaration!")
		return nil, NONE
	}
	if curTok.Lexeme == "(" {
		curTok, err = a.GetNext(false)
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		_, _, paramList := a.IsParameterList()

		curTok, _ = a.GetCurr()
		if curTok.Lexeme != ")" {
			panic(BuildErrMessFromTok(curTok, ")"))
		}
		curTok, err = a.GetNext(false)
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}

		//symbol table operation
		symdata["parameters"] = paramList
		a.AddSymbol(identifier, "Method", symdata, a.pass == 1)

		if e, _ := a.IsMethodBody(false); e != nil {
			panic(BuildErrFromTokErrType(curTok, METHOD_BODY))
		}

		a.debugMessagePassOne("is a field declaration!")
		return nil, NONE
	}

	a.debugMessagePassOne("is a field declaration!")
	return BuildErrFromTokErrType(curTok, FIELD_DECLARATION), FIELD_DECLARATION
}

func (a *Analyzer) IsConstructorDeclaration() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is constructor declaration with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}

	e, _, className := a.IsClassName()
	if e != nil {
		curTok, _ = a.GetCurr()
		return BuildErrFromTokErrType(curTok, CONSTRUCTOR_DECLARATION), CONSTRUCTOR_DECLARATION
	}

	//Semantic Action
	if a.pass == 2 {
		if err := a.sm.Cd(a.st, className); err != nil {
			panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
		}
		a.debugMessagePassTwo(fmt.Sprintf("Cd %s", className))
	}

	curTok, _ = a.GetCurr()
	if curTok.Lexeme != "(" {
		return BuildErrFromTok(curTok, "("), CONSTRUCTOR_DECLARATION
	}
	curTok, err = a.GetNext(false)
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}

	_, _, paramsList := a.IsParameterList()
	curTok, _ = a.GetCurr()
	if curTok.Lexeme != ")" {
		panic(BuildErrMessFromTok(curTok, ")"))
	}
	curTok, err = a.GetNext(false)
	if err != nil {
		panic(BuildErrMessFromTokErrType(curTok, COMPILER))
	}

	//symbol table opperation
	symdata := make(map[string]interface{})
	symdata["class"] = className
	symdata["type"] = className
	symdata["parameters"] = paramsList
	symdata["accessMod"] = "public"
	a.AddSymbol(className, "Constructor", symdata, a.pass == 1)

	if e, t := a.IsMethodBody(true); e != nil {
		panic(BuildErrMessFromTokErrType(curTok, t))
	}

	a.debugMessagePassOne("is a constructor declaration!")
	return nil, NONE
}

func (a *Analyzer) IsMethodBody(isConstructor bool) (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is method body with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}

	if curTok.Lexeme != "{" {
		return BuildErrFromTok(curTok, "{"), METHOD_BODY
	}

	//Semantic Action (icode)
	if a.pass == 2 {
		a.sm.SetupFunc(a.st)
	}

	curTok, err = a.GetNext(false)
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}

	//Semantic Action (icode)
	endFunc := false
	if curTok.Lexeme == "}" && a.pass == 2 {
		if isConstructor {
			a.sm.ReturnThisFunc(a.st)
		} else {
			a.sm.ReturnFunc(a.st)
		}
		endFunc = true
	}

	for err == nil {
		curTok, err = a.GetCurr()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		err, _ = a.IsVariableDeclaration()
	}
	err = nil

	for err == nil {
		curTok, err = a.GetCurr()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		err, _ = a.IsStatement()
	}

	if curTok.Lexeme != "}" {
		panic(BuildErrMessFromTok(curTok, "}"))
	}

	//Semantic Action (icode)
	if a.pass == 2 && !endFunc {
		if isConstructor {
			a.sm.ReturnThisFunc(a.st)
		} else {
			a.sm.ReturnFunc(a.st)
		}
	}

	//symbol table opperation
	//fmt.Println("in method body")
	a.DropScope()

	curTok, err = a.GetNext(false)
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}

	a.debugMessagePassOne("is a method body!")
	return nil, NONE
}

func (a *Analyzer) IsVariableDeclaration() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is variable declaration with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}

	peekTok, _ := a.Peek()
	if peekTok.Type != tok.Identifier {
		return BuildErrFromTokErrType(curTok, VARIABLE_DECLARATION), VARIABLE_DECLARATION
	}

	typ := curTok.Lexeme
	if e, _ := a.IsType(); e != nil {
		curTok, _ = a.GetCurr()
		return BuildErrFromTokErrType(curTok, VARIABLE_DECLARATION), VARIABLE_DECLARATION
		//panic(BuildErrFromTokErrType(curTok, VARIABLE_DECLARATION))
	}

	//Semantic Action
	if a.pass == 2 {
		if err := a.sm.TExists(a.st); err != nil {
			panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
		}
		a.debugMessagePassTwo("TExists!")
	}

	curTok, err = a.GetCurr()

	identifier := curTok.Lexeme
	if curTok.Type != tok.Identifier {
		return BuildErrFromTokErrType(curTok, VARIABLE_DECLARATION), VARIABLE_DECLARATION
		//panic(BuildErrFromTokErrType(curTok, VARIABLE_DECLARATION))
	}
	curTok, err = a.GetNext(false)
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}

	isArr := false
	if curTok.Lexeme == "[" {
		curTok, err = a.GetNext(false)
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		if curTok.Lexeme != "]" {
			panic(BuildErrMessFromTok(curTok, "{"))
		}
		isArr = true

		curTok, err = a.GetNext(false)
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
	}

	//Semantic Action
	if a.pass == 2 {
		a.sm.VPush(identifier, a.st.GetScope(), typ)
		a.debugMessagePassTwo(fmt.Sprintf("vPush %s (%s)", identifier, typ))
	}

	//symbol table opperation
	symdata := make(map[string]interface{})
	symdata["isArray"] = isArr
	symdata["type"] = typ
	a.AddSymbol(identifier, "Lvar", symdata, a.pass == 1)

	curTok, _ = a.GetCurr()
	if curTok.Lexeme == "=" {

		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush(curTok.Lexeme); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s", curTok.Lexeme))
		}

		curTok, err = a.GetNext(false)
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		if e, _ := a.IsAssignmentExpression(); e != nil {
			panic(BuildErrFromTokErrType(curTok, VARIABLE_DECLARATION))
		}
	}
	curTok, _ = a.GetCurr()
	if curTok.Lexeme != ";" {
		panic(BuildErrMessFromTok(curTok, ";"))
	}

	//Semantic Action
	if a.pass == 2 {
		if err := a.sm.EoE(); err != nil {
			panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
		}
		a.debugMessagePassTwo("EOE")
	}

	curTok, err = a.GetNext(false)
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}
	a.debugMessagePassOne("is a variable declaration!")
	return nil, NONE
}

func (a *Analyzer) IsParameterList() (error, ErrorType, []sym.Parameter) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is parameter list with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER, []sym.Parameter{}
	}
	params := make([]sym.Parameter, 0)

	e, _, param := a.IsParameter()
	if e != nil {
		return BuildErrFromTokErrType(curTok, PARAMETER_LIST), PARAMETER_LIST, params
	}

	params = append(params, param)

	for err == nil {
		curTok, err = a.GetCurr()
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		if curTok.Lexeme == "," {
			a.GetNext(false)
			e, _, param := a.IsParameter()
			if e != nil {
				panic(BuildErrFromTokErrType(curTok, PARAMETER_LIST))
			}
			params = append(params, param)
		} else {
			err = BuildErrFromTokErrType(curTok, PARAMETER_LIST)
		}
	}

	a.debugMessagePassOne("is a parameter list!")
	return nil, NONE, params
}

func (a *Analyzer) IsParameter() (error, ErrorType, sym.Parameter) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is parameter with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER, sym.Parameter{}
	}

	typ := curTok.Lexeme
	if e, _ := a.IsType(); e != nil {
		return BuildErrFromTokErrType(curTok, PARAMETER), PARAMETER, sym.Parameter{}
		//panic(BuildErrFromTokErrType(curTok, PARAMETER))
	}

	//Semantic Action
	if a.pass == 2 {
		if err := a.sm.TExists(a.st); err != nil {
			panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
		}
		a.debugMessagePassTwo("TExists!")
	}

	curTok, _ = a.GetCurr() //what are the odds an err would occur? like none right? ... guys?

	identifier := curTok.Lexeme
	if curTok.Type != tok.Identifier {
		panic(BuildTtErrMessFromTok(curTok, tok.Identifier))
	}
	curTok, err = a.GetNext(false)
	if err != nil {
		panic(BuildErrFromTokErrType(curTok, COMPILER))
	}

	isArr := false
	if curTok.Lexeme == "[" {
		curTok, err = a.GetNext(false)
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		if curTok.Lexeme != "]" {
			panic(BuildErrMessFromTok(curTok, "{"))
		}
		a.GetNext(false)

		isArr = true
	}
	a.debugMessagePassOne("is a parameter!")
	return nil, NONE, sym.Parameter{Typ: typ, Identifier: identifier, IsArr: isArr}
}

func (a *Analyzer) IsStatement() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is statement with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch {
	case curTok.Lexeme == "{":
		a.GetNext(false)
		for err == nil {
			err, _ = a.IsStatement()
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != "}" {
			panic(BuildErrMessFromTok(curTok, "}"))
		}
		a.GetNext(false)
	case curTok.Lexeme == "if":
		a.GetNext(false)
		curTok, err = a.GetCurr() //now at next token after expression
		if curTok.Lexeme != "(" {
			panic(BuildErrMessFromTok(curTok, "("))
		}

		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush(curTok.Lexeme); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s", curTok.Lexeme))
		}

		a.GetNext(false)
		if err, _ := a.IsExpression(); err == nil {
			curTok, err = a.GetCurr()
			if curTok.Lexeme != ")" {
				panic(BuildErrMessFromTok(curTok, ")"))
			}

			//Semantic Action EAL
			if a.pass == 2 {
				if err := a.sm.CloseParen(); err != nil {
					panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
				}
				a.debugMessagePassTwo("Close Paren")

				if err := a.sm.If(); err != nil {
					panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
				}
				a.debugMessagePassTwo("If is bool")
			}

			a.GetNext(false)
			if err, _ := a.IsStatement(); err != nil {
				panic(BuildErrFromTokErrType(curTok, STATEMENT))
			}
			curTok, err = a.GetCurr()
			if curTok.Lexeme == "else" {

				//iCode
				if a.pass == 2 {
					a.sm.Else()
				}

				a.GetNext(false)
				if err, _ := a.IsStatement(); err != nil {
					panic(BuildErrFromTokErrType(curTok, STATEMENT))
				}

				//iCode
				if a.pass == 2 {
					a.sm.EndElse()
				}
			} else {
				//iCode
				if a.pass == 2 {
					a.sm.EndIf()
				}
			}
		} else {
			return BuildErrFromTokErrType(curTok, STATEMENT), STATEMENT
		}
	case curTok.Lexeme == "while":
		//Semantic Action Push (iCode)
		var initLabel string
		if a.pass == 2 {
			initLabel = a.sm.InitWhile()
		}

		a.GetNext(false)
		curTok, err = a.GetCurr() //now at next token after expression
		if curTok.Lexeme != "(" {
			panic(BuildErrMessFromTok(curTok, "("))
		}

		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush(curTok.Lexeme); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s", curTok.Lexeme))
		}

		a.GetNext(false)
		if err, _ := a.IsExpression(); err == nil {
			curTok, err = a.GetCurr()
			if curTok.Lexeme != ")" {
				panic(BuildErrMessFromTok(curTok, ")"))
			}

			//Semantic Action EAL
			if a.pass == 2 {
				if err := a.sm.CloseParen(); err != nil {
					panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
				}
				a.debugMessagePassTwo("Close Paren")

				if err := a.sm.While(); err != nil {
					panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
				}
				a.debugMessagePassTwo("While is bool")
			}

			a.GetNext(false)
			if err, _ := a.IsStatement(); err != nil {
				panic(BuildErrFromTokErrType(curTok, STATEMENT))
			}

			//Semantic Action (ICode)
			if a.pass == 2 {
				a.sm.EndWhile(initLabel)
			}
		} else {
			return BuildErrFromTokErrType(curTok, STATEMENT), STATEMENT
		}
	case curTok.Lexeme == "return":
		a.GetNext(false)

		//Semantic Action Meta
		var isVoid bool
		curTok, err = a.GetCurr()
		if curTok.Lexeme == ";" {
			isVoid = true
		} else {
			if err, _ := a.IsExpression(); err != nil {
				panic(err)
			}
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != ";" {
			panic(BuildErrMessFromTok(curTok, ";"))
		}
		//Semantic Action EAL
		if a.pass == 2 {
			if err := a.sm.Return(a.st, isVoid); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("Return")
		}
		a.GetNext(false)
	case curTok.Lexeme == "cout":
		a.GetNext(false)
		curTok, err = a.GetCurr() //now at next token after expression
		if curTok.Lexeme != "<<" {
			panic(BuildErrMessFromTok(curTok, "<<"))
		}
		a.GetNext(false)
		if err, _ := a.IsExpression(); err == nil {
			curTok, err = a.GetCurr()
			if curTok.Lexeme != ";" {
				panic(BuildErrMessFromTok(curTok, ";"))
			}

			//Semantic Action
			if a.pass == 2 {
				if err := a.sm.Cout(); err != nil {
					panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
				}
				a.debugMessagePassTwo("Cout")
			}

			a.GetNext(false)
		} else {
			return BuildErrFromTokErrType(curTok, STATEMENT), STATEMENT
		}
	case curTok.Lexeme == "cin":
		a.GetNext(false)
		curTok, err = a.GetCurr() //now at next token after expression
		if curTok.Lexeme != ">>" {
			panic(BuildErrMessFromTok(curTok, ">>"))
		}
		a.GetNext(false)
		if err, _ := a.IsExpression(); err == nil {
			curTok, err = a.GetCurr()
			if curTok.Lexeme != ";" {
				panic(BuildErrMessFromTok(curTok, ";"))
			}

			//Semantic Action
			if a.pass == 2 {
				if err := a.sm.Cin(); err != nil {
					panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
				}
				a.debugMessagePassTwo("Cout")
			}

			a.GetNext(false)
		} else {
			return BuildErrFromTokErrType(curTok, STATEMENT), STATEMENT
		}
	default:
		if err, _ := a.IsExpression(); err == nil {
			curTok, err = a.GetCurr()
			if curTok.Lexeme != ";" {
				panic(BuildErrMessFromTok(curTok, ";"))
			}
			//Semantic Action
			if a.pass == 2 {
				if err := a.sm.EoE(); err != nil {
					panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
				}
				a.debugMessagePassTwo("EOE")
			}
			a.GetNext(false)
		} else {
			return BuildErrFromTokErrType(curTok, STATEMENT), STATEMENT
		}
	}
	a.debugMessagePassOne("is a statement!")
	return nil, NONE
}

func (a *Analyzer) IsExpression() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is expression with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch {
	case curTok.Lexeme == "(":
		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush(curTok.Lexeme); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s", curTok.Lexeme))
		}

		a.GetNext(false)
		if e, _ := a.IsExpression(); e != nil {
			panic(e.Error())
		}
		curTok, err = a.GetCurr() //now at next token after expression
		if curTok.Lexeme == ")" {
			a.GetNext(false)
		} else {
			panic(BuildErrMessFromTok(curTok, ")"))
		}
		//Semantic Action EAL
		if a.pass == 2 {
			if err := a.sm.CloseParen(); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("Close Paren")
		}
		if e, t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			panic(e.Error())
		}
	case curTok.Lexeme == "true":

		//Semantic Action
		if a.pass == 2 {
			a.sm.LPush(curTok.Lexeme, a.st.GetScope(), "bool")
			a.debugMessagePassTwo(fmt.Sprintf("LPush: %s from scope %s", curTok.Lexeme, a.st.GetScope()))
		}

		a.GetNext(false)
		if e, t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			panic(e.Error())
		}
	case curTok.Lexeme == "false":

		//Semantic Action
		if a.pass == 2 {
			a.sm.LPush(curTok.Lexeme, a.st.GetScope(), "bool")
			a.debugMessagePassTwo(fmt.Sprintf("LPush: %s from scope %s", curTok.Lexeme, a.st.GetScope()))
		}

		a.GetNext(false)
		if e, t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			panic(e.Error())
		}
	case curTok.Lexeme == "null":

		//Semantic Action
		if a.pass == 2 {
			a.sm.LPush(curTok.Lexeme, a.st.GetScope(), "null")
			a.debugMessagePassTwo(fmt.Sprintf("LPush: %s from scope %s", curTok.Lexeme, a.st.GetScope()))
		}

		a.GetNext(false)
		if e, t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			panic(e.Error())
		}
	case curTok.Type == tok.Number:

		//Semantic Action
		if a.pass == 2 {
			a.sm.LPush(curTok.Lexeme, a.st.GetScope(), "int")
			a.debugMessagePassTwo(fmt.Sprintf("LPush: %s from scope %s", curTok.Lexeme, a.st.GetScope()))
		}

		a.GetNext(true)
		if e, t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			panic(e.Error())
		}
	case curTok.Type == tok.Character:

		//Semantic Action
		if a.pass == 2 {
			a.sm.LPush(curTok.Lexeme, a.st.GetScope(), "char")
			a.debugMessagePassTwo(fmt.Sprintf("LPush: %s from scope %s", curTok.Lexeme, a.st.GetScope()))
		}

		a.GetNext(false)
		if e, t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			panic(e.Error())
		}
	case curTok.Lexeme == "this":
		fallthrough
	case curTok.Type == tok.Identifier:
		//Semantic Action
		if a.pass == 2 {
			a.sm.IPush(curTok.Lexeme, a.st.GetScope())
			a.debugMessagePassTwo(fmt.Sprintf("IPush: %s from scope %s", curTok.Lexeme, a.st.GetScope()))
		}

		a.GetNext(true)
		if e, t := a.IsFnArrMember(false); e != nil && t != FN_ARR_MEMBER {
			panic(e.Error())
		}

		//Semantic Action
		if a.pass == 2 {
			if e := a.sm.IExist(a.st); e != nil {
				panic(fmt.Sprintf("%s on line %d", e.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("IExists!")
		}

		if e, t := a.IsMemberRefz(); e != nil && t != MEMBER_REFZ {
			panic(e.Error())
		}
		if e, t := a.IsExpressionZ(); e != nil && t != EXPRESSIONZ {
			panic(e.Error())
		}
	default:
		return BuildErrFromTokErrType(curTok, EXPRESSION), EXPRESSION
	}
	a.debugMessagePassOne("is expression!")
	return nil, NONE
}

func (a *Analyzer) IsFnArrMember(hasRef bool) (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is fn arr member with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Lexeme {
	case "(":
		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush(curTok.Lexeme); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s", curTok.Lexeme))

			a.sm.BAL(a.st.GetScope())
			a.debugMessagePassTwo("BAL")
		}

		curTok, err = a.GetNext(false)
		if err != nil {
			panic(BuildErrFromTokErrType(curTok, COMPILER))
		}
		if curTok.Lexeme != ")" {
			if err, _ := a.IsArgumentList(); err != nil {
				return err, FN_ARR_MEMBER
			}
		}
		//should be pointing at ")"
		curTok, err = a.GetCurr()
		if curTok.Lexeme != ")" {
			panic(BuildErrMessFromTok(curTok, ")"))
		}

		//Semantic Action EAL
		if a.pass == 2 {
			if err := a.sm.CloseParen(); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("Close Paren")

			a.sm.EAL(a.st.GetScope())
			a.debugMessagePassTwo("EAL")

			if err := a.sm.Func(a.st.GetScope(), a.st, hasRef); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("func")
		}

		a.GetNext(false)
	case "[":

		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush(curTok.Lexeme); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s", curTok.Lexeme))
		}

		a.GetNext(false)
		if e, _ := a.IsExpression(); err != nil {
			panic(e.Error())
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != "]" {
			panic(BuildErrMessFromTok(curTok, "]"))
		}

		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.CloseAngleBracket(); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("Close AngleBracket")

			if err := a.sm.Arr(a.st); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("Arr")

		}

		a.GetNext(false)
	default:
		return BuildErrFromTokErrType(curTok, FN_ARR_MEMBER), FN_ARR_MEMBER
	}
	a.debugMessagePassOne("is fn arr member!")
	return nil, NONE
}

func (a *Analyzer) IsMemberRefz() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is member refz with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	if curTok.Lexeme != "." {
		return BuildErrFromTokErrType(curTok, MEMBER_REFZ), MEMBER_REFZ
	}
	a.GetNext(false)
	curTok, err = a.GetCurr()
	if err != nil {
		return err, COMPILER
	}
	if curTok.Type != tok.Identifier {
		panic(BuildTtErrMessFromTok(curTok, tok.Identifier))
	}

	//Semantic Action
	if a.pass == 2 {
		a.sm.IPush(curTok.Lexeme, a.st.GetScope())
		a.debugMessagePassTwo(fmt.Sprintf("IPush: %s from scope %s", curTok.Lexeme, a.st.GetScope()))
	}

	a.GetNext(false)
	if e, t := a.IsFnArrMember(true); e != nil && t != FN_ARR_MEMBER {
		panic(e.Error())
	}

	//Semantic Action
	if a.pass == 2 {
		if err := a.sm.RExist(a.st); err != nil {
			panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
		}
		a.debugMessagePassTwo("RExists!")
	}

	if e, t := a.IsMemberRefz(); e != nil && t != MEMBER_REFZ {
		panic(e.Error())
	}
	a.debugMessagePassOne("is member refz!")
	return nil, NONE
}

func (a *Analyzer) IsExpressionZ() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is expressionz with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Lexeme {
	case "&&", "||", "==", "!=", "<=", ">=", ">", "<", "+", "-", "*", "/":
		a.GetNext(false)
		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush(curTok.Lexeme); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s", curTok.Lexeme))
		}
		if err, _ := a.IsExpression(); err != nil {
			panic(err.Error())
		}
	case "=":
		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush("="); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("Pushed operator =")
		}

		a.GetNext(false)
		if err, _ := a.IsAssignmentExpression(); err != nil {
			panic(err.Error())
		}
	default:
		return BuildErrFromTokErrType(curTok, EXPRESSIONZ), EXPRESSIONZ
	}
	a.debugMessagePassOne("is expressionz!")
	return nil, NONE
}

func (a *Analyzer) IsAssignmentExpression() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is assignment_expression with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch {
	case curTok.Lexeme == "this":
		a.GetNext(false)
	case curTok.Lexeme == "new":
		a.GetNext(false)
		if err, _ := a.IsType(); err != nil {
			panic(err.Error())
		}
		if err, _ := a.IsNewDeclaration(); err != nil {
			panic(err.Error())
		}
	case curTok.Lexeme == "atoi":
		curTok, err = a.GetNext(false)
		if curTok.Lexeme != "(" || err != nil {
			panic(BuildErrMessFromTok(curTok, "("))
		}

		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush(curTok.Lexeme); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s", curTok.Lexeme))
		}

		curTok, err = a.GetNext(false)
		if e, _ := a.IsExpression(); e != nil {
			panic(e.Error())
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != ")" || err != nil {
			panic(BuildErrMessFromTok(curTok, ")"))
		}

		//Semantic Action EAL
		if a.pass == 2 {
			if err := a.sm.CloseParen(); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("Close Paren")

			a.sm.Atoi(a.st.GetScope())
			a.debugMessagePassTwo("atoi")
		}

		a.GetNext(false)
	case curTok.Lexeme == "itoa":
		curTok, err = a.GetNext(false)
		if curTok.Lexeme != "(" || err != nil {
			panic(BuildErrMessFromTok(curTok, "("))
		}

		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush(curTok.Lexeme); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s", curTok.Lexeme))
		}

		curTok, err = a.GetNext(false)
		if e, _ := a.IsExpression(); e != nil {
			panic(e.Error())
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != ")" || err != nil {
			panic(BuildErrMessFromTok(curTok, ")"))
		}

		//Semantic Action EAL
		if a.pass == 2 {
			if err := a.sm.CloseParen(); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("Close Paren")

			a.sm.Itoa(a.st.GetScope())
			a.debugMessagePassTwo("itoa")
		}

		a.GetNext(false)
	default:
		if err, _ := a.IsExpression(); err != nil {
			return err, ASSIGNMENT_EXPRESSION
		}
	}
	a.debugMessagePassOne("is assignment_expression!")
	return nil, NONE
}

func (a *Analyzer) IsNewDeclaration() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is new declaration with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	switch curTok.Lexeme {
	case "(":

		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush(curTok.Lexeme); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s", curTok.Lexeme))

			a.sm.BAL(a.st.GetScope())
			a.debugMessagePassTwo("BAL")
		}

		a.GetNext(false)
		a.IsArgumentList() //dont care if fails
		//should be pointing at ")"
		curTok, err = a.GetCurr()
		if curTok.Lexeme != ")" {
			panic(BuildErrMessFromTok(curTok, ")"))
		}

		//Semantic Action EAL
		if a.pass == 2 {
			if err := a.sm.CloseParen(); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("Close Paren")

			a.sm.EAL(a.st.GetScope())
			a.debugMessagePassTwo("EAL")

			if err := a.sm.NewObj(a.st); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("newObj")
		}

		a.GetNext(false)
	case "[":

		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.OPush(curTok.Lexeme); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo(fmt.Sprintf("Pushed operator %s", curTok.Lexeme))
		}

		a.GetNext(false)
		if err, _ := a.IsExpression(); err != nil {
			panic(err.Error())
		}
		curTok, err = a.GetCurr()
		if curTok.Lexeme != "]" {
			panic(BuildErrMessFromTok(curTok, "]"))
		}

		//Semantic Action OPush
		if a.pass == 2 {
			if err := a.sm.CloseAngleBracket(); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("Close AngleBracket")

			if err := a.sm.NewArray(a.st); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("New Array")

		}

		a.GetNext(false)
	default:
		return BuildErrFromTokErrType(curTok, NEW_DECLARATION), NEW_DECLARATION
	}
	a.debugMessagePassOne("is new declaration!")
	return nil, NONE
}

func (a *Analyzer) IsArgumentList() (error, ErrorType) {
	curTok, err := a.GetCurr()
	a.debugMessagePassOne(fmt.Sprintf("Testing is argument list with token %s...", curTok.Lexeme))
	if err != nil {
		return err, COMPILER
	}
	if e, t := a.IsExpression(); e != nil {
		return e, t
	}
	for err == nil {
		curTok, err = a.GetCurr()
		if err != nil {
			return err, COMPILER
		}
		if curTok.Lexeme != "," {
			break
		}

		//Semantic Action
		if a.pass == 2 {
			if err := a.sm.Comma(); err != nil {
				panic(fmt.Sprintf("%s on line %d", err.Error(), curTok.Linenum+1))
			}
			a.debugMessagePassTwo("Comma")
		}

		a.GetNext(false)
		if e, _ := a.IsExpression(); e != nil {
			panic(e.Error())
		}
	}
	a.debugMessagePassOne("is argument list!")
	return nil, NONE
}
