package analyzer

import (
	//"testing"
	"fmt"
	lex "github.com/sbditto85/compiler/lexer"
	tok "github.com/sbditto85/compiler/token"
)

func ExampleICodeAssignment() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "icode/tests/assignment.kxi"
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

	a.PrintQuadTable()

	//Output:
	//Num Rows: 3 curRow: 2
	//Lables:
	//Rows:
	//MOV Lv2, Li5 ;     int r = 7;
	//MOV Lv3, Li6 ;     int s = 8;
	//MOV Lv2, Lv3 ;     r = s;

}

func ExampleICodeAddition() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "icode/tests/addition.kxi"
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

	a.PrintQuadTable()

	//Output:
	//Num Rows: 4 curRow: 3
	//Lables:
	//Rows:
	//MOV Lv2, Li5 ;     int r = 7;
	//MOV Lv3, Li6 ;     int s = 8;
	//ADD Tv7, Lv3, Lv2 ;     r = r + s;
	//MOV Lv2, Tv7 ;     r = r + s;

}

func ExampleICodeSubtraction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "icode/tests/subtraction.kxi"
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

	a.PrintQuadTable()

	//Output:
	//Num Rows: 4 curRow: 3
	//Lables:
	//Rows:
	//MOV Lv2, Li5 ;     int r = 7;
	//MOV Lv3, Li6 ;     int s = 8;
	//SUB Tv7, Lv3, Lv2 ;     r = r - s;
	//MOV Lv2, Tv7 ;     r = r - s;

}

func ExampleICodeMultiply() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "icode/tests/multiply.kxi"
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

	a.PrintQuadTable()

	//Output:
	//Num Rows: 4 curRow: 3
	//Lables:
	//Rows:
	//MOV Lv2, Li5 ;     int r = 7;
	//MOV Lv3, Li6 ;     int s = 8;
	//MUL Tv7, Lv3, Lv2 ;     r = r * s;
	//MOV Lv2, Tv7 ;     r = r * s;
}

func ExampleICodeDivide() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "icode/tests/divide.kxi"
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

	a.PrintQuadTable()

	//Output:
	//Num Rows: 4 curRow: 3
	//Lables:
	//Rows:
	//MOV Lv2, Li5 ;     int r = 7;
	//MOV Lv3, Li6 ;     int s = 8;
	//DIV Tv7, Lv3, Lv2 ;     r = r / s;
	//MOV Lv2, Tv7 ;     r = r / s;
}

func ExampleICodeArithmetic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "icode/tests/arithmetic.kxi"
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

	a.PrintQuadTable()

	//Output:
	//Num Rows: 10 curRow: 9
	//Lables:
	//Rows:
	//MOV Lv2, Li7 ;     int r = -1;
	//MOV Lv3, Li8 ;     int s = 2;
	//MOV Lv4, Li9 ;     int t = 0;
	//MUL Tv10, Lv3, Lv2 ;     int z = r * s / t;
	//DIV Tv11, Lv4, Tv10 ;     int z = r * s / t;
	//MOV Lv5, Tv11 ;     int z = r * s / t;
	//DIV Tv12, Lv5, Lv4 ;     r = r + t / z - 1;
	//ADD Tv13, Tv12, Lv2 ;     r = r + t / z - 1;
	//SUB Tv15, Li14, Tv13 ;     r = r + t / z - 1;
	//MOV Lv2, Tv15 ;     r = r + t / z - 1;

}
