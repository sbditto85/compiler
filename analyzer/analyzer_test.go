package analyzer

import (
	//"testing"
	lex "github.com/sbditto85/compiler/lexer"
)

func ExampleParseXEqualsY() {
	var str []string
	str = append(str,"x = y;")
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: x
	//Testing is statement with token x...
	//Testing is expression with token x...
	//Token: =
	//Testing is fn arr member with token =...
	//Testing is member refz with token =...
	//Testing is expressionz with token =...
	//Token: y
	//Testing is assignment_expression with token y...
	//Testing is expression with token y...
	//Token: ;
	//Testing is fn arr member with token ;...
	//Testing is member refz with token ;...
	//Testing is expressionz with token ;...
	//is expression!
	//is assignment_expression!
	//is expressionz!
	//is expression!
	//is a statement!
}

func ExampleParseTrueEqualsFalse() {
	var str []string
	str = append(str,"true == false;")
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: true
	//Testing is statement with token true...
	//Testing is expression with token true...
	//Token: ==
	//Testing is expressionz with token ==...
	//Token: false
	//Testing is expression with token false...
	//Token: ;
	//Testing is expressionz with token ;...
	//is expression!
	//is expressionz!
	//is expression!
	//is a statement!
}

func ExampleParseTrue() {
	var str []string
	str = append(str,"true;")
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: true
	//Testing is statement with token true...
	//Testing is expression with token true...
	//Token: ;
	//Testing is expressionz with token ;...
	//is expression!
	//is a statement!
}

func ExampleParse123() {
	var str []string
	str = append(str,"123;")
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: 123
	//Testing is statement with token 123...
	//Testing is expression with token 123...
	//Token: ;
	//Testing is expressionz with token ;...
	//is expression!
	//is a statement!
}

func ExampleParseCharacter() {
	var str []string
	str = append(str,"'a';")
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: 'a'
	//Testing is statement with token 'a'...
	//Testing is expression with token 'a'...
	//Token: ;
	//Testing is expressionz with token ;...
	//is expression!
	//is a statement!
}

func ExampleParseCharacterNewLine() {
	var str []string
	str = append(str,`'\n';`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: '\n'
	//Testing is statement with token '\n'...
	//Testing is expression with token '\n'...
	//Token: ;
	//Testing is expressionz with token ;...
	//is expression!
	//is a statement!
}


func ExampleParseXGreaterThanYPlus3() {
	var str []string
	str = append(str,`x>y+3;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: x
	//Testing is statement with token x...
	//Testing is expression with token x...
	//Token: >
	//Testing is fn arr member with token >...
	//Testing is member refz with token >...
	//Testing is expressionz with token >...
	//Token: y
	//Testing is expression with token y...
	//Token: +
	//Testing is fn arr member with token +...
	//Testing is member refz with token +...
	//Testing is expressionz with token +...
	//Token: 3
	//Testing is expression with token 3...
	//Token: ;
	//Testing is expressionz with token ;...
	//is expression!
	//is expressionz!
	//is expression!
	//is expressionz!
	//is expression!
	//is a statement!
}

func ExampleParseXAssignThis() {
	var str []string
	str = append(str,`x=this;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: x
	//Testing is statement with token x...
	//Testing is expression with token x...
	//Token: =
	//Testing is fn arr member with token =...
	//Testing is member refz with token =...
	//Testing is expressionz with token =...
	//Token: this
	//Testing is assignment_expression with token this...
	//Token: ;
	//is assignment_expression!
	//is expressionz!
	//is expression!
	//is a statement!
}

func ExampleParseXAssignIToA() {
	var str []string
	str = append(str,`x=itoa(123);`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: x
	//Testing is statement with token x...
	//Testing is expression with token x...
	//Token: =
	//Testing is fn arr member with token =...
	//Testing is member refz with token =...
	//Testing is expressionz with token =...
	//Token: itoa
	//Testing is assignment_expression with token itoa...
	//Token: (
	//Token: 123
	//Testing is expression with token 123...
	//Token: )
	//Testing is expressionz with token )...
	//is expression!
	//Token: ;
	//is assignment_expression!
	//is expressionz!
	//is expression!
	//is a statement!
}

func ExampleParseXAssignAToI() {
	var str []string
	str = append(str,`x=atoi('a');`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: x
	//Testing is statement with token x...
	//Testing is expression with token x...
	//Token: =
	//Testing is fn arr member with token =...
	//Testing is member refz with token =...
	//Testing is expressionz with token =...
	//Token: atoi
	//Testing is assignment_expression with token atoi...
	//Token: (
	//Token: 'a'
	//Testing is expression with token 'a'...
	//Token: )
	//Testing is expressionz with token )...
	//is expression!
	//Token: ;
	//is assignment_expression!
	//is expressionz!
	//is expression!
	//is a statement!
}

func ExampleParseXAssignNewYNoArgs() {
	var str []string
	str = append(str,`x=new Y();`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: x
	//Testing is statement with token x...
	//Testing is expression with token x...
	//Token: =
	//Testing is fn arr member with token =...
	//Testing is member refz with token =...
	//Testing is expressionz with token =...
	//Token: new
	//Testing is assignment_expression with token new...
	//Token: Y
	//Testing is type with token Y...
	//Testing is classname with token Y...
	//Token: (
	//is classname!
	//is type!
	//Testing is new declaration with token (...
	//Token: )
	//Testing is argument list with token )...
	//Testing is expression with token )...
	//Token: ;
	//is new declaration!
	//is assignment_expression!
	//is expressionz!
	//is expression!
	//is a statement!
}

func ExampleParseXAssignNewYArg() {
	var str []string
	str = append(str,`x=new Y(true);`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: x
	//Testing is statement with token x...
	//Testing is expression with token x...
	//Token: =
	//Testing is fn arr member with token =...
	//Testing is member refz with token =...
	//Testing is expressionz with token =...
	//Token: new
	//Testing is assignment_expression with token new...
	//Token: Y
	//Testing is type with token Y...
	//Testing is classname with token Y...
	//Token: (
	//is classname!
	//is type!
	//Testing is new declaration with token (...
	//Token: true
	//Testing is argument list with token true...
	//Testing is expression with token true...
	//Token: )
	//Testing is expressionz with token )...
	//is expression!
	//is argument list!
	//Token: ;
	//is new declaration!
	//is assignment_expression!
	//is expressionz!
	//is expression!
	//is a statement!
}

func ExampleParseXAssignNewYArgs() {
	var str []string
	str = append(str,`x=new Y(true,false);`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: x
	//Testing is statement with token x...
	//Testing is expression with token x...
	//Token: =
	//Testing is fn arr member with token =...
	//Testing is member refz with token =...
	//Testing is expressionz with token =...
	//Token: new
	//Testing is assignment_expression with token new...
	//Token: Y
	//Testing is type with token Y...
	//Testing is classname with token Y...
	//Token: (
	//is classname!
	//is type!
	//Testing is new declaration with token (...
	//Token: true
	//Testing is argument list with token true...
	//Testing is expression with token true...
	//Token: ,
	//Testing is expressionz with token ,...
	//is expression!
	//Token: false
	//Testing is expression with token false...
	//Token: )
	//Testing is expressionz with token )...
	//is expression!
	//is argument list!
	//Token: ;
	//is new declaration!
	//is assignment_expression!
	//is expressionz!
	//is expression!
	//is a statement!
}

func ExampleParseXAssignNewYArgs2() {
	var str []string
	str = append(str,`x=new Y(true,false,x<y);`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: x
	//Testing is statement with token x...
	//Testing is expression with token x...
	//Token: =
	//Testing is fn arr member with token =...
	//Testing is member refz with token =...
	//Testing is expressionz with token =...
	//Token: new
	//Testing is assignment_expression with token new...
	//Token: Y
	//Testing is type with token Y...
	//Testing is classname with token Y...
	//Token: (
	//is classname!
	//is type!
	//Testing is new declaration with token (...
	//Token: true
	//Testing is argument list with token true...
	//Testing is expression with token true...
	//Token: ,
	//Testing is expressionz with token ,...
	//is expression!
	//Token: false
	//Testing is expression with token false...
	//Token: ,
	//Testing is expressionz with token ,...
	//is expression!
	//Token: x
	//Testing is expression with token x...
	//Token: <
	//Testing is fn arr member with token <...
	//Testing is member refz with token <...
	//Testing is expressionz with token <...
	//Token: y
	//Testing is expression with token y...
	//Token: )
	//Testing is fn arr member with token )...
	//Testing is member refz with token )...
	//Testing is expressionz with token )...
	//is expression!
	//is expressionz!
	//is expression!
	//is argument list!
	//Token: ;
	//is new declaration!
	//is assignment_expression!
	//is expressionz!
	//is expression!
	//is a statement!
}

func ExampleParseXAssignNewYArr() {
	var str []string
	str = append(str,`x=new Y[true];`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	a.IsStatement()

	//Output:
	//Token: x
	//Testing is statement with token x...
	//Testing is expression with token x...
	//Token: =
	//Testing is fn arr member with token =...
	//Testing is member refz with token =...
	//Testing is expressionz with token =...
	//Token: new
	//Testing is assignment_expression with token new...
	//Token: Y
	//Testing is type with token Y...
	//Testing is classname with token Y...
	//Token: [
	//is classname!
	//is type!
	//Testing is new declaration with token [...
	//Token: true
	//Testing is expression with token true...
	//Token: ]
	//Testing is expressionz with token ]...
	//is expression!
	//Token: ;
	//is new declaration!
	//is assignment_expression!
	//is expressionz!
	//is expression!
	//is a statement!
}

//x=;
func ExampleFailXEqualSemi() {
	var str []string
	str = append(str,`x=;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,false)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		println(err.Error())
	}
	
	//Output:
	//
}
//=y;
func ExampleFailEqualYSemi() {
	var str []string
	str = append(str,`=y;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,false)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		println(err.Error())
	}
	
	//Output:
	//
}
//f(g,((g+g)>(g*g)))
