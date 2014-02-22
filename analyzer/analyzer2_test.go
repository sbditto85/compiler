package analyzer

import (
	//"testing"
	lex "github.com/sbditto85/compiler/lexer"
	tok "github.com/sbditto85/compiler/token"
	"fmt"
) 

func ExampleXEqualYPass2() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/xequaly.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l,false)
	a.GetNext()
	err := a.PerformPass()
	
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

	l = lex.NewLexer()
	l.ReadFile(file)
	a.SetLexer(l)

	err = a.PerformNextPass(true)
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}
	
	//Output:
	//IPush: x from scope g.main
	//IExists!
	//Pushed operator =
	//IPush: y from scope g.main
	//IExists!
	//SM: Testing operation = ...
	//SM: Comparing y(int) to x(int)
	//SM: ... finished operation =
	//EOE
}

func ExampleAssignInfixToPostfix() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/assigninfixtopostfix.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l,false)
	a.GetNext()
	err := a.PerformPass()
	
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

	l = lex.NewLexer()
	l.ReadFile(file)
	a.SetLexer(l)

	err = a.PerformNextPass(true)
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}
	
	//Output:
	//IPush: x from scope g.main
	//IExists!
	//Pushed operator =
	//IPush: y from scope g.main
	//IExists!
	//Pushed operator *
	//IPush: g from scope g.main
	//IExists!
	//SM: Comparing g(int) to y(int) for *
	//Pushed operator +
	//IPush: f from scope g.main
	//IExists!
	//Pushed operator /
	//IPush: k from scope g.main
	//IExists!
	//SM: Testing operation / ...
	//SM: Comparing k(int) to f(int) for /
	//SM: ... finished operation /
	//SM: Testing operation + ...
	//SM: Comparing f / k(int) to y * g(int) for +
	//SM: ... finished operation +
	//SM: Testing operation = ...
	//SM: Comparing y * g + f / k(int) to x(int)
	//SM: ... finished operation =
	//EOE
}

func ExampleAssignClassRef() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/assignclassref.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l,false)
	a.GetNext()
	err := a.PerformPass()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	l = lex.NewLexer()
	l.ReadFile(file)
	a.SetLexer(l)

	err = a.PerformNextPass(true)
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}
	
	//Output:
	//IPush: a from scope g.main
	//IExists!
	//IPush: x from scope g.main
	//RExists!
	//Pushed operator =
	//IPush: b from scope g.main
	//IExists!
	//IPush: y from scope g.main
	//RExists!
	//SM: Testing operation = ...
	//SM: Comparing b.y(int) to a.x(int)
	//SM: ... finished operation =
	//EOE
}

func ExampleAssignMemberFunctionWInfixToPostfix() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/assignMemberFunctionWInfixToPostfix.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l,false)
	a.GetNext()
	err := a.PerformPass()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	l = lex.NewLexer()
	l.ReadFile(file)
	a.SetLexer(l)

	err = a.PerformNextPass(true)
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}
	
	//Output:
	//
}
