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
	//Li5: []int{1}
	//Li6: []int{2}
	//Lv2: []int{1, 3}
	//Lv3: []int{2, 3}
	//Ma1: []int{0}
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
	//Li5: []int{1}
	//Li6: []int{2}
	//Lv2: []int{1, 3, 4}
	//Lv3: []int{2, 3}
	//Ma1: []int{0}
	//Tv7: []int{3, 4}
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
	//Li5: []int{1}
	//Li6: []int{2}
	//Lv2: []int{1, 3, 4}
	//Lv3: []int{2, 3}
	//Ma1: []int{0}
	//Tv7: []int{3, 4}
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
	//Li5: []int{1}
	//Li6: []int{2}
	//Lv2: []int{1, 3, 4}
	//Lv3: []int{2, 3}
	//Ma1: []int{0}
	//Tv7: []int{3, 4}
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
	//Li5: []int{1}
	//Li6: []int{2}
	//Lv2: []int{1, 3, 4}
	//Lv3: []int{2, 3}
	//Ma1: []int{0}
	//Tv7: []int{3, 4}
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
	//Li14: []int{9}
	//Li7: []int{1}
	//Li8: []int{2}
	//Li9: []int{3}
	//Lv2: []int{1, 4, 8, 10}
	//Lv3: []int{2, 4}
	//Lv4: []int{3, 5, 7}
	//Lv5: []int{6, 7}
	//Ma1: []int{0}
	//Tv10: []int{4, 5}
	//Tv11: []int{5, 6}
	//Tv12: []int{7, 8}
	//Tv13: []int{8, 9}
	//Tv15: []int{9, 10}
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
	//Cl1: []int{16}
	//Co3: []int{0, 17, 18}
	//Iv2: []int{5, 6, 7, 12}
	//Li10: []int{13}
	//Lv7: []int{20, 21, 22}
	//Ma6: []int{15}
	//Me5: []int{4, 21, 23}
	//Pa4: []int{7}
	//St11: []int{1, 2, 11}
	//Tv14: []int{5, 9}
	//Tv15: []int{6, 8}
	//Tv16: []int{7, 8}
	//Tv17: []int{8, 9}
	//Tv18: []int{16, 17, 19, 20}
	//Tv19: []int{24}
	//Tv9: []int{12, 13}
	//this: []int{1, 5, 6, 12}
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
	//Num Rows: 22 curRow: 21
	//Lables:
	//El18: []int{13}
	//El9: []int{5}
	//En23: []int{17, 21}
	//If16: []int{11, 14}
	//If7: []int{3, 6}
	//Li10: []int{6, 7, 8, 16}
	//Li17: []int{12}
	//Li19: []int{14, 18}
	//Li4: []int{1}
	//Li5: []int{2}
	//Li8: []int{4}
	//Lv2: []int{1, 2, 4, 6, 7, 8, 12, 14, 15, 16, 18, 19}
	//Ma1: []int{0}
	//Tv11: []int{6, 10}
	//Tv12: []int{7, 9}
	//Tv13: []int{8, 9}
	//Tv14: []int{9, 10}
	//Tv15: []int{10, 11}
	//Tv20: []int{14, 15}
	//Tv22: []int{16, 17}
	//Tv24: []int{18, 19}
	//Tv6: []int{2, 3}
	//Wh21: []int{13, 5, 16, 20}
	//Rows:
	//FUNC Ma1 ;void main() {
	//MOV Lv2, Li4 ;    int x = 7;
	//LT Tv6, Lv2, Li5 ;    if ( x < 4) {
	//BF Tv6, If7 ;    if ( x < 4) {
	//MOV Lv2, Li8 ;	x = 8;
	//JMP Wh21 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//If7: LTE Tv11, Lv2, Li10 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//GTE Tv12, Lv2, Li10 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//NEQ Tv13, Lv2, Li10 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//AND Tv14, Tv12, Tv13 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//OR Tv15, Tv11, Tv14 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//BF Tv15, If16 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//MOV Lv2, Li17 ;	x = 2;
	//JMP Wh21 ;    } else {
	//If16: ADD Tv20, Li19, Lv2 ;	x = x + 1;
	//MOV Lv2, Tv20 ;	x = x + 1;
	//Wh21: GT Tv22, Lv2, Li10 ;    while(x > 3) {
	//BF Tv22, En23 ;    while(x > 3) {
	//SUB Tv24, Li19, Lv2 ;	x = x - 1;
	//MOV Lv2, Tv24 ;	x = x - 1;
	//JMP Wh21 ;}
	//En23: RTN  ;}
	//Num Rows: 0 curRow: -1
	//Lables:
	//Rows:

}

func ExampleICodeFunctionChain() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "icode/tests/functionChain.kxi"
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
	//Num Rows: 47 curRow: 46
	//Lables:
	//Cl1: []int{22}
	//Co3: []int{0, 23, 24}
	//Iv2: []int{8, 9, 14, 18, 30, 38}
	//Li12: []int{19}
	//Li18: []int{10, 31}
	//Li28: []int{43}
	//Li29: []int{43}
	//Lv8: []int{26, 27, 32, 40}
	//Lv9: []int{39, 45}
	//Ma7: []int{21}
	//Me4: []int{4, 27, 28, 32, 33, 35, 36}
	//Me5: []int{7}
	//Me6: []int{13, 40, 41}
	//St13: []int{1, 2, 17}
	//Tv11: []int{18, 19}
	//Tv16: []int{8, 11}
	//Tv17: []int{9, 10}
	//Tv19: []int{10, 11}
	//Tv20: []int{14, 15}
	//Tv21: []int{22, 23, 25, 26}
	//Tv22: []int{29, 30}
	//Tv23: []int{30, 31}
	//Tv24: []int{34, 35}
	//Tv25: []int{37, 38}
	//Tv26: []int{38, 39}
	//Tv27: []int{42, 44}
	//Tv30: []int{43, 44}
	//Tv31: []int{44, 45}
	//this: []int{1, 5, 8, 9, 14, 18}
	//Rows:
	//FUNC Co3 ;      Dog() {
	//FRAME this, St13 ;      Dog() {
	//CALL St13 ;      Dog() {
	//RTN  ;      }
	//FUNC Me4 ;      public Dog GetDog() {
	//RETURN this ;             return this;
	//RTN  ;      }
	//FUNC Me5 ;      public void Bite() {
	//REF Tv16, Iv2, this ;             bites = bites + 1;
	//REF Tv17, Iv2, this ;             bites = bites + 1;
	//ADD Tv19, Li18, Tv17 ;             bites = bites + 1;
	//MOV Tv16, Tv19 ;             bites = bites + 1;
	//RTN  ;      }
	//FUNC Me6 ;      public int GetBites() {
	//REF Tv20, Iv2, this ;             return bites;
	//RETURN Tv20 ;             return bites;
	//RTN  ;      }
	//FUNC St13 ;}
	//REF Tv11, Iv2, this ;      public int bites = 0;
	//MOV Tv11, Li12 ;      public int bites = 0;
	//RTN  ;}
	//FUNC Ma7 ;void main() {
	//NEWI Cl1, Tv21 ;     Dog d = new Dog();
	//FRAME Tv21, Co3 ;     Dog d = new Dog();
	//CALL Co3 ;     Dog d = new Dog();
	//PEEK Tv21 ;     Dog d = new Dog();
	//MOV Lv8, Tv21 ;     Dog d = new Dog();
	//FRAME Lv8, Me4 ;     d.GetDog().bites = 1;
	//CALL Me4 ;     d.GetDog().bites = 1;
	//PEEK Tv22 ;     d.GetDog().bites = 1;
	//REF Tv23, Iv2, Tv22 ;     d.GetDog().bites = 1;
	//MOV Tv23, Li18 ;     d.GetDog().bites = 1;
	//FRAME Lv8, Me4 ;     b = d.GetDog().GetDog().bites;
	//CALL Me4 ;     b = d.GetDog().GetDog().bites;
	//PEEK Tv24 ;     b = d.GetDog().GetDog().bites;
	//FRAME Tv24, Me4 ;     b = d.GetDog().GetDog().bites;
	//CALL Me4 ;     b = d.GetDog().GetDog().bites;
	//PEEK Tv25 ;     b = d.GetDog().GetDog().bites;
	//REF Tv26, Iv2, Tv25 ;     b = d.GetDog().GetDog().bites;
	//MOV Lv9, Tv26 ;     b = d.GetDog().GetDog().bites;
	//FRAME Lv8, Me6 ;     b = d.GetBites() + 7 * 3;
	//CALL Me6 ;     b = d.GetBites() + 7 * 3;
	//PEEK Tv27 ;     b = d.GetBites() + 7 * 3;
	//MUL Tv30, Li29, Li28 ;     b = d.GetBites() + 7 * 3;
	//ADD Tv31, Tv30, Tv27 ;     b = d.GetBites() + 7 * 3;
	//MOV Lv9, Tv31 ;     b = d.GetBites() + 7 * 3;
	//RTN  ;}
	//Num Rows: 0 curRow: -1
	//Lables:
	//Rows:

}

func ExampleICodeFunctionCall() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "icode/tests/functionCall.kxi"
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
	//Num Rows: 17 curRow: 16
	//Lables:
	//Co2: []int{0}
	//Li13: []int{9}
	//Li14: []int{9}
	//Li16: []int{10}
	//Li17: []int{10}
	//Lv7: []int{11}
	//Ma6: []int{8}
	//Me5: []int{4, 11, 14}
	//St8: []int{1, 2, 6}
	//Tv15: []int{9, 12}
	//Tv18: []int{10, 13}
	//Tv19: []int{15}
	//this: []int{1}
	//Rows:
	//FUNC Co2 ;      Apple() {
	//FRAME this, St8 ;      Apple() {
	//CALL St8 ;      Apple() {
	//RTN  ;      }
	//FUNC Me5 ;      public void MyFunc(int i, bool j) {
	//RTN  ;      }
	//FUNC St8 ;}
	//RTN  ;}
	//FUNC Ma6 ;void main() {
	//ADD Tv15, Li14, Li13 ;     a.MyFunc(1 + 3, 4 < 7);
	//LT Tv18, Li16, Li17 ;     a.MyFunc(1 + 3, 4 < 7);
	//FRAME Lv7, Me5 ;     a.MyFunc(1 + 3, 4 < 7);
	//PUSH Tv15 ;     a.MyFunc(1 + 3, 4 < 7);
	//PUSH Tv18 ;     a.MyFunc(1 + 3, 4 < 7);
	//CALL Me5 ;     a.MyFunc(1 + 3, 4 < 7);
	//PEEK Tv19 ;     a.MyFunc(1 + 3, 4 < 7);
	//RTN  ;}
	//Num Rows: 0 curRow: -1
	//Lables:
	//Rows:

}

func ExampleICodeArraysEverywhere() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "icode/tests/arraysEverywhere.kxi"
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
