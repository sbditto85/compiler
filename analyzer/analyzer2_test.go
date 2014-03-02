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
	//Type int pushed
	//TExists!
	//Type int pushed
	//TExists!
	//vPush r (int)
	//EOE
	//Type Baz pushed
	//TExists!
	//vPush b (Baz)
	//Pushed operator =
	//Type Baz pushed
	//Pushed operator (
	//BAL
	//IPush: r from scope g.main
	//IExists!
	//SM: Finished )
	//Close Paren
	//EAL
	//SM: Type: Baz, with 1 Arguments
	//newObj
	//SM: Testing operation = ...
	//SM: Comparing Baz(r)(Baz) to b(Baz)
	//SM: ... finished operation =
	//EOE
}

func ExampleCreateInstanceOfArray() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/createInstanceOfArray.kxi"
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
	//Type int pushed
	//TExists!
	//Type int pushed
	//TExists!
	//vPush r (int)
	//EOE
	//Type Baz pushed
	//TExists!
	//vPush b (Baz)
	//Pushed operator =
	//Type Baz pushed
	//Pushed operator [
	//IPush: r from scope g.main
	//IExists!
	//SM: Finished ]
	//Close AngleBracket
	//SM: Type: Baz, with array size r
	//New Array
	//SM: Testing operation = ...
	//SM: Comparing Baz[r](Baz) to b(Baz)
	//SM: ... finished operation =
	//EOE
}

func ExampleFuncInfixToPostfixArg() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/funcInfixToPostfixArg.kxi"
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
	//Type bool pushed
	//TExists!
	//Type int pushed
	//TExists!
	//vPush g (int)
	//Pushed operator =
	//LPush: 1 from scope g.Baz.y
	//SM: Testing operation = ...
	//SM: Comparing 1(int) to g(int)
	//SM: ... finished operation =
	//EOE
	//Type int pushed
	//TExists!
	//vPush k (int)
	//Pushed operator =
	//LPush: 2 from scope g.Baz.y
	//SM: Testing operation = ...
	//SM: Comparing 2(int) to k(int)
	//SM: ... finished operation =
	//EOE
	//IPush: f from scope g.Baz.y
	//Pushed operator (
	//BAL
	//IPush: i from scope g.Baz.y
	//IExists!
	//Pushed operator *
	//LPush: 3 from scope g.Baz.y
	//SM: Testing operation * ...
	//SM: Comparing 3(int) to i(int) for *
	//SM: ... finished operation *
	//SM: Finished ,
	//Comma
	//IPush: g from scope g.Baz.y
	//IExists!
	//Pushed operator <
	//IPush: k from scope g.Baz.y
	//IExists!
	//SM: Testing operation < ...
	//SM: Comparing k(int) to g(int) for op <
	//SM: ... finished operation <
	//SM: Finished )
	//Close Paren
	//EAL
	//SM: Identifer: f, with 2 Arguments
	//func
	//IExists!
	//EOE
	//Type int pushed
	//TExists!
	//Type int pushed
	//TExists!
	//Type bool pushed
	//TExists!
}

func ExampleArrayArrayInfixPostfix() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/arrayArrayInfixPostfix.kxi"
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
	//Type Baz pushed
	//TExists!
	//vPush c (Baz)
	//Pushed operator =
	//Type Baz pushed
	//Pushed operator [
	//LPush: 10 from scope g.main
	//SM: Finished ]
	//Close AngleBracket
	//SM: Type: Baz, with array size 10
	//New Array
	//SM: Testing operation = ...
	//SM: Comparing Baz[10](Baz) to c(Baz)
	//SM: ... finished operation =
	//EOE
	//Type int pushed
	//TExists!
	//vPush r (int)
	//Pushed operator =
	//LPush: 7 from scope g.main
	//SM: Testing operation = ...
	//SM: Comparing 7(int) to r(int)
	//SM: ... finished operation =
	//EOE
	//IPush: c from scope g.main
	//Pushed operator [
	//IPush: r from scope g.main
	//IExists!
	//Pushed operator +
	//LPush: 3 from scope g.main
	//SM: Testing operation + ...
	//SM: Comparing 3(int) to r(int) for +
	//SM: ... finished operation +
	//SM: Finished ]
	//Close AngleBracket
	//SM: Type: c, with array size r + 3
	//Arr
	//IExists!
	//Pushed operator =
	//IPush: c from scope g.main
	//Pushed operator [
	//IPush: r from scope g.main
	//IExists!
	//Pushed operator -
	//LPush: 5 from scope g.main
	//SM: Testing operation - ...
	//SM: Comparing 5(int) to r(int) for -
	//SM: ... finished operation -
	//SM: Finished ]
	//Close AngleBracket
	//SM: Type: c, with array size r - 5
	//Arr
	//IExists!
	//SM: Testing operation = ...
	//SM: Comparing c[r - 5](Baz) to c[r + 3](Baz)
	//SM: ... finished operation =
	//EOE
}

func ExampleIfWhileReturnCoutCinAtoiItoaCdBoolops() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/ifWhileReturnCoutCinAtoiItoaCdBoolops.kxi"
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
	//vPush i (int)
	//EOE
	//Type int pushed
	//TExists!
	//vPush j (int)
	//EOE
	//Type char pushed
	//TExists!
	//vPush c (char)
	//EOE
	//Type bool pushed
	//TExists!
	//vPush b (bool)
	//EOE
	//Type bool pushed
	//TExists!
	//vPush b2 (bool)
	//EOE
	//Pushed operator (
	//IPush: i from scope g.main
	//IExists!
	//Pushed operator <
	//IPush: j from scope g.main
	//IExists!
	//SM: Testing operation < ...
	//SM: Comparing j(int) to i(int) for op <
	//SM: ... finished operation <
	//SM: Finished )
	//Close Paren
	//If is bool
	//Pushed operator (
	//IPush: b from scope g.main
	//IExists!
	//Pushed operator &&
	//IPush: b2 from scope g.main
	//IExists!
	//SM: Comparing b2 and b as bool for op &&
	//Pushed operator ||
	//IPush: b from scope g.main
	//IExists!
	//SM: Testing operation || ...
	//SM: Comparing b and b && b2 as bool for op ||
	//SM: ... finished operation ||
	//SM: Finished )
	//Close Paren
	//While is bool
	//IPush: i from scope g.main
	//IExists!
	//Pushed operator =
	//Pushed operator (
	//IPush: c from scope g.main
	//IExists!
	//SM: Finished )
	//Close Paren
	//atoi
	//SM: Testing operation = ...
	//SM: Comparing atoi(c)(int) to i(int)
	//SM: ... finished operation =
	//EOE
	//IPush: i from scope g.main
	//IExists!
	//Cout
	//IPush: i from scope g.main
	//IExists!
	//Cout
	//IPush: c from scope g.main
	//IExists!
	//Pushed operator =
	//Pushed operator (
	//IPush: i from scope g.main
	//IExists!
	//SM: Finished )
	//Close Paren
	//itoa
	//SM: Testing operation = ...
	//SM: Comparing itoa(i)(char) to c(char)
	//SM: ... finished operation =
	//EOE

}
