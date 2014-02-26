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
	//Type int pushed
	//TExists!
	//vPush x (int)
	//EOE
	//Type int pushed
	//TExists!
	//vPush y (int)
	//EOE
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
	//Type int pushed
	//TExists!
	//vPush x (int)
	//EOE
	//Type int pushed
	//TExists!
	//vPush y (int)
	//EOE
	//Type int pushed
	//TExists!
	//vPush g (int)
	//EOE
	//Type int pushed
	//TExists!
	//vPush f (int)
	//EOE
	//Type int pushed
	//TExists!
	//vPush k (int)
	//EOE
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
	//Type int pushed
	//TExists!
	//Type int pushed
	//TExists!
	//Type Cat pushed
	//TExists!
	//vPush a (Cat)
	//EOE
	//Type Dog pushed
	//TExists!
	//vPush b (Dog)
	//EOE
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
	//Type char pushed
	//TExists!
	//Type int pushed
	//TExists!
	//Type bool pushed
	//TExists!
	//Type int pushed
	//TExists!
	//Type bool pushed
	//TExists!
	//vPush k (bool)
	//Pushed operator =
	//LPush: true from scope g.main
	//SM: Testing operation = ...
	//SM: Comparing true(bool) to k(bool)
	//SM: ... finished operation =
	//EOE
	//Type int pushed
	//TExists!
	//vPush g (int)
	//Pushed operator =
	//LPush: 1 from scope g.main
	//SM: Testing operation = ...
	//SM: Comparing 1(int) to g(int)
	//SM: ... finished operation =
	//EOE
	//Type int pushed
	//TExists!
	//vPush r (int)
	//Pushed operator =
	//LPush: 2 from scope g.main
	//SM: Testing operation = ...
	//SM: Comparing 2(int) to r(int)
	//SM: ... finished operation =
	//EOE
	//Type Foo pushed
	//TExists!
	//vPush y (Foo)
	//EOE
	//Type int pushed
	//TExists!
	//vPush x (int)
	//EOE
	//IPush: x from scope g.main
	//IExists!
	//Pushed operator =
	//IPush: y from scope g.main
	//IExists!
	//IPush: f from scope g.main
	//Pushed operator (
	//BAL
	//IPush: k from scope g.main
	//IExists!
	//SM: Finished ,
	//Comma
	//IPush: g from scope g.main
	//IExists!
	//SM: Finished )
	//Close Paren
	//EAL
	//SM: Identifer: f, with 2 Arguments
	//func
	//RExists!
	//Pushed operator +
	//IPush: g from scope g.main
	//IExists!
	//Pushed operator *
	//IPush: r from scope g.main
	//IExists!
	//SM: Testing operation * ...
	//SM: Comparing r(int) to g(int) for *
	//SM: ... finished operation *
	//SM: Testing operation + ...
	//SM: Comparing g * r(int) to y.f(bool, int)(int) for +
	//SM: ... finished operation +
	//SM: Testing operation = ...
	//SM: Comparing y.f(bool, int) + g * r(int) to x(int)
	//SM: ... finished operation =
	//EOE
}

func ExampleAssignMemberElemntReturnByFunc() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/assignMemberElemntReturnByFunc.kxi"
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
	//Type char pushed
	//TExists!
	//Type Bar pushed
	//TExists!
	//Type Baz pushed
	//TExists!
	//Type int pushed
	//TExists!
	//Type int pushed
	//TExists!
	//Type int pushed
	//TExists!
	//vPush r (int)
	//Pushed operator =
	//LPush: 2 from scope g.main
	//SM: Testing operation = ...
	//SM: Comparing 2(int) to r(int)
	//SM: ... finished operation =
	//EOE
	//Type Foo pushed
	//TExists!
	//vPush x (Foo)
	//EOE
	//IPush: x from scope g.main
	//IExists!
	//IPush: f from scope g.main
	//Pushed operator (
	//BAL
	//SM: Finished )
	//Close Paren
	//EAL
	//SM: Identifer: f, with 0 Arguments
	//func
	//RExists!
	//IPush: g from scope g.main
	//Pushed operator (
	//BAL
	//SM: Finished )
	//Close Paren
	//EAL
	//SM: Identifer: g, with 0 Arguments
	//func
	//RExists!
	//IPush: y from scope g.main
	//RExists!
	//Pushed operator =
	//IPush: r from scope g.main
	//IExists!
	//SM: Testing operation = ...
	//SM: Comparing r(int) to x.f().g().y(int)
	//SM: ... finished operation =
	//EOE
}

func ExampleCreateInstanceOfClass() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/createInstanceOfClass.kxi"
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
	//Type int pushed
	//TExists!
	//Type int pushed
	//TExists!
	//Type Baz pushed
	//TExists!
	//vPush b (Baz)
	//Pushed operator =
	//Type Baz pushed
	//Pushed operator (
	//BAL
	//SM: Finished )
	//Close Paren
	//EAL
	//SM: Type: Baz, with 0 Arguments
	//newObj
	//SM: Testing operation = ...
	//SM: Comparing Baz()(Baz) to b(Baz)
	//SM: ... finished operation =
	//EOE
}
