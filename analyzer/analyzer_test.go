package analyzer

import (
	//"testing"
	lex "github.com/sbditto85/compiler/lexer"
	tok "github.com/sbditto85/compiler/token"
	"fmt"
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
	//Token: ''
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
	//Token: ''
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
	//Token: ''
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
	//Token: ''
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
	//Token: ''
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
	//Token: ''
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
	//Token: ''
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
	//Token: ''
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
	//Token: ''
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
	//Token: ''
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
	//Token: ''
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
	//Token: ''
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
	//Token: ''
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
	//Token: ''
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
	//Token: ''
	//is a statement!
}

//x=;
func ExampleFailXEqualSemi() {
defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`x=;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		fmt.Println(err.Error())
	}
	
	//Output:
	//Token: x
	//Testing is statement with token x...
	//Testing is expression with token x...
	//Token: =
	//Testing is fn arr member with token =...
	//Testing is member refz with token =...
	//Testing is expressionz with token =...
	//Token: ;
	//Testing is assignment_expression with token ;...
	//Testing is expression with token ;...
	//Expected Expression, received ';' on line 1
}
//=y; 
func ExampleFailEqualYSemi() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`=y;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		fmt.Println(err.Error())
	}
	
	//Output:
	//Token: =
	//Testing is statement with token =...
	//Testing is expression with token =...
	//Expected Statement, received '=' on line 1
}
//f(g,((g+g)>(g*g)))
func ExampleFunctionCallWithExpressionParams() {
	var str []string
	str = append(str,`f(g,((g+g)>(g*g)));`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		fmt.Println(err.Error())
	}
	
	//Output:
	//Token: f
	//Testing is statement with token f...
	//Testing is expression with token f...
	//Token: (
	//Testing is fn arr member with token (...
	//Token: g
	//Testing is argument list with token g...
	//Testing is expression with token g...
	//Token: ,
	//Testing is fn arr member with token ,...
	//Testing is member refz with token ,...
	//Testing is expressionz with token ,...
	//is expression!
	//Token: (
	//Testing is expression with token (...
	//Token: (
	//Testing is expression with token (...
	//Token: g
	//Testing is expression with token g...
	//Token: +
	//Testing is fn arr member with token +...
	//Testing is member refz with token +...
	//Testing is expressionz with token +...
	//Token: g
	//Testing is expression with token g...
	//Token: )
	//Testing is fn arr member with token )...
	//Testing is member refz with token )...
	//Testing is expressionz with token )...
	//is expression!
	//is expressionz!
	//is expression!
	//Token: >
	//Testing is expressionz with token >...
	//Token: (
	//Testing is expression with token (...
	//Token: g
	//Testing is expression with token g...
	//Token: *
	//Testing is fn arr member with token *...
	//Testing is member refz with token *...
	//Testing is expressionz with token *...
	//Token: g
	//Testing is expression with token g...
	//Token: )
	//Testing is fn arr member with token )...
	//Testing is member refz with token )...
	//Testing is expressionz with token )...
	//is expression!
	//is expressionz!
	//is expression!
	//Token: )
	//Testing is expressionz with token )...
	//is expression!
	//is expressionz!
	//is expression!
	//Token: )
	//Testing is expressionz with token )...
	//is expression!
	//is argument list!
	//Token: ;
	//is fn arr member!
	//Testing is member refz with token ;...
	//Testing is expressionz with token ;...
	//is expression!
	//Token: ''
	//is a statement!
}

func ExampleCinPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`cin >> a;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		fmt.Println(err.Error())
	}
	
	//Output:
	//Token: cin
	//Testing is statement with token cin...
	//Token: >>
	//Token: a
	//Testing is expression with token a...
	//Token: ;
	//Testing is fn arr member with token ;...
	//Testing is member refz with token ;...
	//Testing is expressionz with token ;...
	//is expression!
	//Token: ''
	//is a statement!
}

func ExampleCoutPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`cout << a;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: cout
	//Testing is statement with token cout...
	//Token: <<
	//Token: a
	//Testing is expression with token a...
	//Token: ;
	//Testing is fn arr member with token ;...
	//Testing is member refz with token ;...
	//Testing is expressionz with token ;...
	//is expression!
	//Token: ''
	//is a statement!
	
}

func ExampleReturnPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`return;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: return
	//Testing is statement with token return...
	//Token: ;
	//Testing is expression with token ;...
	//Token: ''
	//is a statement!
}

func ExampleReturnParamPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`return a;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: return
	//Testing is statement with token return...
	//Token: a
	//Testing is expression with token a...
	//Token: ;
	//Testing is fn arr member with token ;...
	//Testing is member refz with token ;...
	//Testing is expressionz with token ;...
	//is expression!
	//Token: ''
	//is a statement!
}

func ExampleWhilePass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`while(true) a;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: while
	//Testing is statement with token while...
	//Token: (
	//Token: true
	//Testing is expression with token true...
	//Token: )
	//Testing is expressionz with token )...
	//is expression!
	//Token: a
	//Testing is statement with token a...
	//Testing is expression with token a...
	//Token: ;
	//Testing is fn arr member with token ;...
	//Testing is member refz with token ;...
	//Testing is expressionz with token ;...
	//is expression!
	//Token: ''
	//is a statement!
	//is a statement!
}

func ExampleIfPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`if(true)a;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: if
	//Testing is statement with token if...
	//Token: (
	//Token: true
	//Testing is expression with token true...
	//Token: )
	//Testing is expressionz with token )...
	//is expression!
	//Token: a
	//Testing is statement with token a...
	//Testing is expression with token a...
	//Token: ;
	//Testing is fn arr member with token ;...
	//Testing is member refz with token ;...
	//Testing is expressionz with token ;...
	//is expression!
	//Token: ''
	//is a statement!
	//is a statement!
}

func ExampleIfElsePass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`if(true) a; else b;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: if
	//Testing is statement with token if...
	//Token: (
	//Token: true
	//Testing is expression with token true...
	//Token: )
	//Testing is expressionz with token )...
	//is expression!
	//Token: a
	//Testing is statement with token a...
	//Testing is expression with token a...
	//Token: ;
	//Testing is fn arr member with token ;...
	//Testing is member refz with token ;...
	//Testing is expressionz with token ;...
	//is expression!
	//Token: else
	//is a statement!
	//Token: b
	//Testing is statement with token b...
	//Testing is expression with token b...
	//Token: ;
	//Testing is fn arr member with token ;...
	//Testing is member refz with token ;...
	//Testing is expressionz with token ;...
	//is expression!
	//Token: ''
	//is a statement!
	//is a statement!
}

func ExampleBlockStatementPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`{a;b;}`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: {
	//Testing is statement with token {...
	//Token: a
	//Testing is statement with token a...
	//Testing is expression with token a...
	//Token: ;
	//Testing is fn arr member with token ;...
	//Testing is member refz with token ;...
	//Testing is expressionz with token ;...
	//is expression!
	//Token: b
	//is a statement!
	//Testing is statement with token b...
	//Testing is expression with token b...
	//Token: ;
	//Testing is fn arr member with token ;...
	//Testing is member refz with token ;...
	//Testing is expressionz with token ;...
	//is expression!
	//Token: }
	//is a statement!
	//Testing is statement with token }...
	//Testing is expression with token }...
	//Token: ''
	//is a statement!
}

func ExampleEmptyBlockPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`{{}}`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsStatement()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: {
	//Testing is statement with token {...
	//Token: {
	//Testing is statement with token {...
	//Token: }
	//Testing is statement with token }...
	//Testing is expression with token }...
	//Token: }
	//is a statement!
	//Testing is statement with token }...
	//Testing is expression with token }...
	//Token: ''
	//is a statement!
}

func ExampleParameterPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`int apple`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_,_ := a.IsParameter()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: int
	//Testing is parameter with token int...
	//Testing is type with token int...
	//Token: apple
	//is type!
	//Token: ''
	//is a parameter!
}

func ExampleParameterArrPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`int apple[]`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_,_ := a.IsParameter()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: int
	//Testing is parameter with token int...
	//Testing is type with token int...
	//Token: apple
	//is type!
	//Token: [
	//Token: ]
	//Token: ''
	//is a parameter!
}

func ExampleParameterAsParameterListPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`int apple[]`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_,_ := a.IsParameterList()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: int
	//Testing is parameter list with token int...
	//Testing is parameter with token int...
	//Testing is type with token int...
	//Token: apple
	//is type!
	//Token: [
	//Token: ]
	//Token: ''
	//is a parameter!
	//is a parameter list!
}

func ExampleParameterListPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`int a, cat c[]`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_,_ := a.IsParameterList()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: int
	//Testing is parameter list with token int...
	//Testing is parameter with token int...
	//Testing is type with token int...
	//Token: a
	//is type!
	//Token: ,
	//is a parameter!
	//Token: cat
	//Testing is parameter with token cat...
	//Testing is type with token cat...
	//Testing is classname with token cat...
	//Token: c
	//is classname!
	//is type!
	//Token: [
	//Token: ]
	//Token: ''
	//is a parameter!
	//is a parameter list!
}

func ExampleVariableDeclarationPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`bool a;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsVariableDeclaration()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: bool
	//Testing is variable declaration with token bool...
	//Testing is type with token bool...
	//Token: a
	//is type!
	//Token: ;
	//Token: ''
	//is a variable declaration!
}

func ExampleVariableDeclarationAssignPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`bool a = true;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsVariableDeclaration()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: bool
	//Testing is variable declaration with token bool...
	//Testing is type with token bool...
	//Token: a
	//is type!
	//Token: =
	//Token: true
	//Testing is assignment_expression with token true...
	//Testing is expression with token true...
	//Token: ;
	//Testing is expressionz with token ;...
	//is expression!
	//is assignment_expression!
	//Token: ''
	//is a variable declaration!
}

func ExampleMethodBodyPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`{}`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsMethodBody()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: {
	//Testing is method body with token {...
	//Token: }
	//Testing is variable declaration with token }...
	//Testing is statement with token }...
	//Testing is expression with token }...
	//Token: ''
	//is a method body!
}

func ExampleMethodBodyFullPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`{ bool a; a = true; }`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsMethodBody()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: {
	//Testing is method body with token {...
	//Token: bool
	//Testing is variable declaration with token bool...
	//Testing is type with token bool...
	//Token: a
	//is type!
	//Token: ;
	//Token: a
	//is a variable declaration!
	//Testing is variable declaration with token a...
	//Testing is statement with token a...
	//Testing is expression with token a...
	//Token: =
	//Testing is fn arr member with token =...
	//Testing is member refz with token =...
	//Testing is expressionz with token =...
	//Token: true
	//Testing is assignment_expression with token true...
	//Testing is expression with token true...
	//Token: ;
	//Testing is expressionz with token ;...
	//is expression!
	//is assignment_expression!
	//is expressionz!
	//is expression!
	//Token: }
	//is a statement!
	//Testing is statement with token }...
	//Testing is expression with token }...
	//Token: ''
	//is a method body!
}

func ExampleConstructorPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`Cat() {}`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsConstructorDeclaration()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: Cat
	//Testing is constructor declaration with token Cat...
	//Testing is classname with token Cat...
	//Token: (
	//is classname!
	//Token: )
	//Testing is parameter list with token )...
	//Testing is parameter with token )...
	//Testing is type with token )...
	//Testing is classname with token )...
	//Token: {
	//Testing is method body with token {...
	//Token: }
	//Testing is variable declaration with token }...
	//Testing is statement with token }...
	//Testing is expression with token }...
	//Token: ''
	//is a method body!
	//is a constructor declaration!
}

func ExampleConstructorParamsPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`Cat(bool a, void c) {}`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsConstructorDeclaration()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: Cat
	//Testing is constructor declaration with token Cat...
	//Testing is classname with token Cat...
	//Token: (
	//is classname!
	//Token: bool
	//Testing is parameter list with token bool...
	//Testing is parameter with token bool...
	//Testing is type with token bool...
	//Token: a
	//is type!
	//Token: ,
	//is a parameter!
	//Token: void
	//Testing is parameter with token void...
	//Testing is type with token void...
	//Token: c
	//is type!
	//Token: )
	//is a parameter!
	//is a parameter list!
	//Token: {
	//Testing is method body with token {...
	//Token: }
	//Testing is variable declaration with token }...
	//Testing is statement with token }...
	//Testing is expression with token }...
	//Token: ''
	//is a method body!
	//is a constructor declaration!
}

func ExampleFieldDeclarationPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_,_,_,_ := a.IsFieldDeclaration()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: ;
	//Testing is field declaration with token ;...
	//Token: ''
	//is a field declaration!
}

func ExampleFieldDeclarationArrPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`[];`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_,_,_,_ := a.IsFieldDeclaration()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: [
	//Testing is field declaration with token [...
	//Token: ]
	//Token: ;
	//Token: ''
	//is a field declaration!
}

func ExampleFieldDeclarationAssignPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`= x + y;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_,_,_,_ := a.IsFieldDeclaration()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: =
	//Testing is field declaration with token =...
	//Token: x
	//Testing is assignment_expression with token x...
	//Testing is expression with token x...
	//Token: +
	//Testing is fn arr member with token +...
	//Testing is member refz with token +...
	//Testing is expressionz with token +...
	//Token: y
	//Testing is expression with token y...
	//Token: ;
	//Testing is fn arr member with token ;...
	//Testing is member refz with token ;...
	//Testing is expressionz with token ;...
	//is expression!
	//is expressionz!
	//is expression!
	//is assignment_expression!
	//Token: ''
	//is a field declaration!
}

func ExampleFieldDeclarationArrAssignPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`[] = true;`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_,_,_,_ := a.IsFieldDeclaration()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: [
	//Testing is field declaration with token [...
	//Token: ]
	//Token: =
	//Token: true
	//Testing is assignment_expression with token true...
	//Testing is expression with token true...
	//Token: ;
	//Testing is expressionz with token ;...
	//is expression!
	//is assignment_expression!
	//Token: ''
	//is a field declaration!
}

func ExampleBasicClassDeclarationPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`class myclass {}`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsClassDeclaration()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: class
	//Testing is class declaration with token class...
	//Token: myclass
	//Testing is classname with token myclass...
	//Token: {
	//is classname!
	//Token: }
	//Testing is class member declaration with token }...
	//Testing is constructor declaration with token }...
	//Testing is classname with token }...
	//Token: ''
	//is a class declaration!
}

func ExampleBasicClassDeclarationComplexPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`class myclass {`)
	str = append(str,`private int x;`)
	str = append(str,`myclass(){}`)
	str = append(str,`public int f(){}`)
	str = append(str,`}`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsClassDeclaration()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: class
	//Testing is class declaration with token class...
	//Token: myclass
	//Testing is classname with token myclass...
	//Token: {
	//is classname!
	//Token: private
	//Testing is class member declaration with token private...
	//Testing is modifier with token private...
	//Token: int
	//is modifier!
	//Testing is type with token int...
	//Token: x
	//is type!
	//Token: ;
	//Testing is field declaration with token ;...
	//Token: myclass
	//is a field declaration!
	//is a class member declaration!
	//Testing is class member declaration with token myclass...
	//Testing is constructor declaration with token myclass...
	//Testing is classname with token myclass...
	//Token: (
	//is classname!
	//Token: )
	//Testing is parameter list with token )...
	//Testing is parameter with token )...
	//Testing is type with token )...
	//Testing is classname with token )...
	//Token: {
	//Testing is method body with token {...
	//Token: }
	//Testing is variable declaration with token }...
	//Testing is statement with token }...
	//Testing is expression with token }...
	//Token: public
	//is a method body!
	//is a constructor declaration!
	//is a class member declaration!
	//Testing is class member declaration with token public...
	//Testing is modifier with token public...
	//Token: int
	//is modifier!
	//Testing is type with token int...
	//Token: f
	//is type!
	//Token: (
	//Testing is field declaration with token (...
	//Token: )
	//Testing is parameter list with token )...
	//Testing is parameter with token )...
	//Testing is type with token )...
	//Testing is classname with token )...
	//Token: {
	//Testing is method body with token {...
	//Token: }
	//Testing is variable declaration with token }...
	//Testing is statement with token }...
	//Testing is expression with token }...
	//Token: }
	//is a method body!
	//is a field declaration!
	//is a class member declaration!
	//Testing is class member declaration with token }...
	//Testing is constructor declaration with token }...
	//Testing is classname with token }...
	//Token: ''
	//is a class declaration!
}

func ExampleCompilationUnitPass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,`void main () {}`)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsCompilationUnit()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Token: void
	//Testing is compilation unit with token void...
	//Token: main
	//Token: (
	//Token: )
	//Token: {
	//Testing is method body with token {...
	//Token: }
	//Testing is variable declaration with token }...
	//Testing is statement with token }...
	//Testing is expression with token }...
	//Token: ''
	//is a method body!
	//is a compliation unit!
}

//f(r*e, g < k) ;
//c [ r + 3 ] = c [ r - 5 ] ;
// a = b;
// a = b[i];
// a[i] = b[i];
// a[i] = b;
// a[] = b[];



/*
func ExamplePass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	var str []string
	str = append(str,``)
	l := lex.NewLexer()
	l.LoadStrings(str)

	a := NewAnalyzer(l,true)
	a.GetNext()
	err,_ := a.IsParameter()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:

}
*/
