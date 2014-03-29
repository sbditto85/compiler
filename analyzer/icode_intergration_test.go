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
	//Num Rows: 5 curRow: 4
	//Lables:
	//Rows:
	//FUNC Ma1 ;
	//MOV Lv2, Li5 ;     int r = 7;
	//MOV Lv3, Li6 ;     int s = 8;
	//MOV Lv2, Lv3 ;     r = s;
	//RTN  ;}

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
	//Num Rows: 6 curRow: 5
	//Lables:
	//Rows:
	//FUNC Ma1 ;
	//MOV Lv2, Li5 ;     int r = 7;
	//MOV Lv3, Li6 ;     int s = 8;
	//ADD Tv7, Lv3, Lv2 ;     r = r + s;
	//MOV Lv2, Tv7 ;     r = r + s;
	//RTN  ;}

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
	//Num Rows: 6 curRow: 5
	//Lables:
	//Rows:
	//FUNC Ma1 ;
	//MOV Lv2, Li5 ;     int r = 7;
	//MOV Lv3, Li6 ;     int s = 8;
	//SUB Tv7, Lv3, Lv2 ;     r = r - s;
	//MOV Lv2, Tv7 ;     r = r - s;
	//RTN  ;}

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
	//Num Rows: 6 curRow: 5
	//Lables:
	//Rows:
	//FUNC Ma1 ;
	//MOV Lv2, Li5 ;     int r = 7;
	//MOV Lv3, Li6 ;     int s = 8;
	//MUL Tv7, Lv3, Lv2 ;     r = r * s;
	//MOV Lv2, Tv7 ;     r = r * s;
	//RTN  ;}
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
	//Num Rows: 6 curRow: 5
	//Lables:
	//Rows:
	//FUNC Ma1 ;
	//MOV Lv2, Li5 ;     int r = 7;
	//MOV Lv3, Li6 ;     int s = 8;
	//DIV Tv7, Lv3, Lv2 ;     r = r / s;
	//MOV Lv2, Tv7 ;     r = r / s;
	//RTN  ;}
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
	a.PrintQuadStatic()

	//Output:
	//Num Rows: 12 curRow: 11
	//Lables:
	//Rows:
	//FUNC Ma1 ;
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
	//RTN  ;}
	//Num Rows: 0 curRow: -1
	//Lables:
	//Rows:

}

func ExampleICodeReference() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "icode/tests/reference.kxi"
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
	//Num Rows: 26 curRow: 25
	//Lables:
	//Rows:
	//FUNC Co3 ;      Dog(){}
	//FRAME this, St11 ;      Dog(){}
	//CALL St11 ;      Dog(){}
	//RTN  ;      Dog(){}
	//FUNC Me5 ;      public void AddX(Dog d) {
	//REF Tv14, Iv2, this ;             x = x + d.x;
	//REF Tv15, Iv2, this ;             x = x + d.x;
	//REF Tv16, Iv2, Pa4 ;             x = x + d.x;
	//ADD Tv17, Tv16, Tv15 ;             x = x + d.x;
	//MOV Tv14, Tv17 ;             x = x + d.x;
	//RTN  ;      }
	//FUNC St11 ;}
	//REF Tv9, Iv2, this ;      public int x = 7;
	//MOV Tv9, Li10 ;      public int x = 7;
	//RTN  ;}
	//FUNC Ma6 ;void main() {
	//NEWI Cl1, Tv18 ;     Dog d = new Dog();
	//FRAME Tv18, Co3 ;     Dog d = new Dog();
	//CALL Co3 ;     Dog d = new Dog();
	//PEEK Tv18 ;     Dog d = new Dog();
	//MOV Lv7, Tv18 ;     Dog d = new Dog();
	//FRAME Lv7, Me5 ;     d.AddX(d);
	//PUSH Lv7 ;     d.AddX(d);
	//CALL Me5 ;     d.AddX(d);
	//PEEK Tv19 ;     d.AddX(d);
	//RTN  ;}

}


func ExampleICodeFlowControlBasic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "icode/tests/flowControlBasic.kxi"
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
	a.PrintQuadStatic()

	//Output:

}
