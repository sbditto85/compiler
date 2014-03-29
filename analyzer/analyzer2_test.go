package analyzer

import (
	//"testing"
	"fmt"
	lex "github.com/sbditto85/compiler/lexer"
	tok "github.com/sbditto85/compiler/token"
)

func ExampleXEqualYPass2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/xequaly.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l, false)
	a.GetNext()
	err := a.PerformPass()

	if err != nil {
		fmt.Println(err.Error())
	}

	curTok, err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s", curTok.Lexeme)
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

	curTok, err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s", curTok.Lexeme)
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/assigninfixtopostfix.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l, false)
	a.GetNext()
	err := a.PerformPass()

	if err != nil {
		fmt.Println(err.Error())
	}

	curTok, err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s", curTok.Lexeme)
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

	curTok, err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s", curTok.Lexeme)
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/assignclassref.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l, false)
	a.GetNext()
	err := a.PerformPass()

	if err != nil {
		fmt.Println(err.Error())
	}

	curTok, err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
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

	curTok, err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/assignMemberFunctionWInfixToPostfix.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l, false)
	a.GetNext()
	err := a.PerformPass()

	if err != nil {
		fmt.Println(err.Error())
	}

	curTok, err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
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

	curTok, err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Type char pushed
	//TExists!
	//vPush f (char)
	//EOE
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/assignMemberElemntReturnByFunc.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l, false)
	a.GetNext()
	err := a.PerformPass()

	if err != nil {
		fmt.Println(err.Error())
	}

	curTok, err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
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

	curTok, err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Type char pushed
	//TExists!
	//vPush f (char)
	//EOE
	//Type Bar pushed
	//TExists!
	//Type Baz pushed
	//TExists!
	//Type int pushed
	//TExists!
	//vPush y (int)
	//EOE
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/createInstanceOfClass.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l, false)
	a.GetNext()
	err := a.PerformPass()

	if err != nil {
		fmt.Println(err.Error())
	}

	curTok, err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
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

	curTok, err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Cd Baz
	//Type int pushed
	//TExists!
	//Type int pushed
	//TExists!
	//vPush y (int)
	//EOE
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/createInstanceOfArray.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l, false)
	a.GetNext()
	err := a.PerformPass()

	if err != nil {
		fmt.Println(err.Error())
	}

	curTok, err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
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

	curTok, err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Cd Baz
	//Type int pushed
	//TExists!
	//Type int pushed
	//TExists!
	//vPush y (int)
	//EOE
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/funcInfixToPostfixArg.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l, false)
	a.GetNext()
	err := a.PerformPass()

	if err != nil {
		fmt.Println(err.Error())
	}

	curTok, err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
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

	curTok, err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/arrayArrayInfixPostfix.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l, false)
	a.GetNext()
	err := a.PerformPass()

	if err != nil {
		fmt.Println(err.Error())
	}

	curTok, err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
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

	curTok, err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Cd Baz
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/ifWhileReturnCoutCinAtoiItoaCdBoolops.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l, false)
	a.GetNext()
	err := a.PerformPass()

	if err != nil {
		fmt.Println(err.Error())
	}

	curTok, err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
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

	curTok, err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	//Output:
	//Type int pushed
	//TExists!
	//vPush notused (int)
	//EOE
	//Cd Baz
	//Type int pushed
	//TExists!
	//LPush: 1 from scope g.Baz.FUNC
	//SM: Expression return type (int) expected (int)
	//Return
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
	//SM: Returning from a void function
	//Return

}

func ExampleIfWhileReturnCoutCinAtoiItoaCdBoolopsSymTable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/ifWhileReturnCoutCinAtoiItoaCdBoolops.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l, false)
	a.GetNext()
	err := a.PerformPass()

	if err != nil {
		fmt.Println(err.Error())
	}

	curTok, err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	l = lex.NewLexer()
	l.ReadFile(file)
	a.SetLexer(l)

	err = a.PerformNextPass(false)
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok, err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s\n", curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	a.PrintTableInAddOrder()

	//Output:
	//Current Scope: g
	//=================
	//Elements:
	//Scope: g, SymId: Cl1, Value: Baz, Kind: Class
	//Extra Data:
	//Key: size, Value: 4
	//Key: StaticInit, Value: St12
	//--------------
	//Scope: g.Baz, SymId: Iv2, Value: notused, Kind: Ivar
	//Extra Data:
	//Key: accessMod, Value: private
	//Key: type, Value: int
	//Key: isArray, Value: false
	//Key: this_class, Value: Baz
	//--------------
	//Scope: g.Baz, SymId: Co3, Value: Baz, Kind: Constructor
	//Extra Data:
	//Key: class, Value: Baz
	//Key: type, Value: Baz
	//Key: parameters, Value: []
	//Key: accessMod, Value: public
	//Key: paramSymIds, Value: []
	//--------------
	//Scope: g.Baz, SymId: Me4, Value: FUNC, Kind: Method
	//Extra Data:
	//Key: accessMod, Value: public
	//Key: type, Value: int
	//Key: parameters, Value: []
	//Key: paramSymIds, Value: []
	//--------------
	//Scope: g, SymId: Ma5, Value: main, Kind: Main
	//Extra Data:
	//Key: type, Value: void
	//--------------
	//Scope: g.main, SymId: Lv6, Value: i, Kind: Lvar
	//Extra Data:
	//Key: isArray, Value: false
	//Key: type, Value: int
	//--------------
	//Scope: g.main, SymId: Lv7, Value: j, Kind: Lvar
	//Extra Data:
	//Key: isArray, Value: false
	//Key: type, Value: int
	//--------------
	//Scope: g.main, SymId: Lv8, Value: c, Kind: Lvar
	//Extra Data:
	//Key: isArray, Value: false
	//Key: type, Value: char
	//--------------
	//Scope: g.main, SymId: Lv9, Value: b, Kind: Lvar
	//Extra Data:
	//Key: isArray, Value: false
	//Key: type, Value: bool
	//--------------
	//Scope: g.main, SymId: Lv10, Value: b2, Kind: Lvar
	//Extra Data:
	//Key: isArray, Value: false
	//Key: type, Value: bool
	//--------------
	//Scope: g, SymId: Ty11, Value: int, Kind: Type
	//Extra Data:
	//Key: type, Value: int
	//Key: scope, Value: g
	//--------------
	//Scope: g.Baz, SymId: St12, Value: BazStaticInit, Kind: StaticInit
	//Extra Data:
	//Key: type, Value: Baz
	//Key: accessMod, Value: private
	//Key: scope, Value: g.Baz
	//--------------
	//Scope: g, SymId: Li13, Value: 1, Kind: LitVar
	//Extra Data:
	//Key: type, Value: int
	//Key: scope, Value: g
	//--------------
	//Scope: g, SymId: Ty14, Value: char, Kind: Type
	//Extra Data:
	//Key: type, Value: char
	//Key: scope, Value: g
	//--------------
	//Scope: g, SymId: Ty15, Value: bool, Kind: Type
	//Extra Data:
	//Key: type, Value: bool
	//Key: scope, Value: g
	//--------------
	//Scope: g.main, SymId: Tv16, Value: i < j, Kind: Tvar
	//Extra Data:
	//Key: type, Value: bool
	//--------------
	//Scope: g.main, SymId: Tv19, Value: b && b2, Kind: Tvar
	//Extra Data:
	//Key: type, Value: bool
	//--------------
	//Scope: g.main, SymId: Tv20, Value: b && b2 || b, Kind: Tvar
	//Extra Data:
	//Key: type, Value: bool
	//--------------
	//Scope: g.main, SymId: Tv22, Value: atoi(c), Kind: Tvar
	//Extra Data:
	//Key: type, Value: char
	//--------------
	//Scope: g.main, SymId: Tv23, Value: itoa(i), Kind: Tvar
	//Extra Data:
	//Key: type, Value: int
	//--------------

}
