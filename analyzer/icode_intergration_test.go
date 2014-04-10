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
	a.GetNext(false)
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
	//Li4: []int{1}
	//Li5: []int{2}
	//Lv1: []int{1, 3}
	//Lv2: []int{2, 3}
	//MAIN: []int{0}
	//Rows:
	//FUNC MAIN ;
	//MOV Lv1, Li4 ;     int r = 7;
	//MOV Lv2, Li5 ;     int s = 8;
	//MOV Lv1, Lv2 ;     r = s;
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
	a.GetNext(false)
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
	//Li4: []int{1}
	//Li5: []int{2}
	//Lv1: []int{1, 3, 4}
	//Lv2: []int{2, 3}
	//MAIN: []int{0}
	//Tv6: []int{3, 4}
	//Rows:
	//FUNC MAIN ;
	//MOV Lv1, Li4 ;     int r = 7;
	//MOV Lv2, Li5 ;     int s = 8;
	//ADD Tv6, Lv2, Lv1 ;     r = r + s;
	//MOV Lv1, Tv6 ;     r = r + s;
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
	a.GetNext(false)
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
	//Li4: []int{1}
	//Li5: []int{2}
	//Lv1: []int{1, 3, 4}
	//Lv2: []int{2, 3}
	//MAIN: []int{0}
	//Tv6: []int{3, 4}
	//Rows:
	//FUNC MAIN ;
	//MOV Lv1, Li4 ;     int r = 7;
	//MOV Lv2, Li5 ;     int s = 8;
	//SUB Tv6, Lv2, Lv1 ;     r = r - s;
	//MOV Lv1, Tv6 ;     r = r - s;
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
	a.GetNext(false)
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
	//Li4: []int{1}
	//Li5: []int{2}
	//Lv1: []int{1, 3, 4}
	//Lv2: []int{2, 3}
	//MAIN: []int{0}
	//Tv6: []int{3, 4}
	//Rows:
	//FUNC MAIN ;
	//MOV Lv1, Li4 ;     int r = 7;
	//MOV Lv2, Li5 ;     int s = 8;
	//MUL Tv6, Lv2, Lv1 ;     r = r * s;
	//MOV Lv1, Tv6 ;     r = r * s;
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
	a.GetNext(false)
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
	//Li4: []int{1}
	//Li5: []int{2}
	//Lv1: []int{1, 3, 4}
	//Lv2: []int{2, 3}
	//MAIN: []int{0}
	//Tv6: []int{3, 4}
	//Rows:
	//FUNC MAIN ;
	//MOV Lv1, Li4 ;     int r = 7;
	//MOV Lv2, Li5 ;     int s = 8;
	//DIV Tv6, Lv2, Lv1 ;     r = r / s;
	//MOV Lv1, Tv6 ;     r = r / s;
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
	a.GetNext(false)
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
	//Li13: []int{9}
	//Li6: []int{1}
	//Li7: []int{2}
	//Li8: []int{3}
	//Lv1: []int{1, 4, 8, 10}
	//Lv2: []int{2, 4}
	//Lv3: []int{3, 5, 7}
	//Lv4: []int{6, 7}
	//MAIN: []int{0}
	//Tv10: []int{5, 6}
	//Tv11: []int{7, 8}
	//Tv12: []int{8, 9}
	//Tv14: []int{9, 10}
	//Tv9: []int{4, 5}
	//Rows:
	//FUNC MAIN ;
	//MOV Lv1, Li6 ;     int r = -1;
	//MOV Lv2, Li7 ;     int s = 2;
	//MOV Lv3, Li8 ;     int t = 0;
	//MUL Tv9, Lv2, Lv1 ;     int z = r * s / t;
	//DIV Tv10, Lv3, Tv9 ;     int z = r * s / t;
	//MOV Lv4, Tv10 ;     int z = r * s / t;
	//DIV Tv11, Lv4, Lv3 ;     r = r + t / z - 1;
	//ADD Tv12, Tv11, Lv1 ;     r = r + t / z - 1;
	//SUB Tv14, Li13, Tv12 ;     r = r + t / z - 1;
	//MOV Lv1, Tv14 ;     r = r + t / z - 1;
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
	a.GetNext(false)
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
	//Lv6: []int{20, 21, 22}
	//MAIN: []int{15}
	//Me5: []int{4, 21, 23}
	//Pa4: []int{7}
	//St7: []int{1, 2, 11}
	//Tv13: []int{5, 9}
	//Tv14: []int{6, 8}
	//Tv15: []int{7, 8}
	//Tv16: []int{8, 9}
	//Tv17: []int{16, 17, 19, 20}
	//Tv18: []int{24}
	//Tv9: []int{12, 13}
	//this: []int{1, 3, 5, 6, 12}
	//Rows:
	//FUNC Co3 ;      Dog(){}
	//FRAME this, St7 ;      Dog(){}
	//CALL St7 ;      Dog(){}
	//RETURN this ;      Dog(){}
	//FUNC Me5 ;      public void AddX(Dog d) {
	//REF Tv13, Iv2, this ;             x = x + d.x;
	//REF Tv14, Iv2, this ;             x = x + d.x;
	//REF Tv15, Iv2, Pa4 ;             x = x + d.x;
	//ADD Tv16, Tv15, Tv14 ;             x = x + d.x;
	//MOV Tv13, Tv16 ;             x = x + d.x;
	//RTN  ;      }
	//FUNC St7 ;}
	//REF Tv9, Iv2, this ;      public int x = 7;
	//MOV Tv9, Li10 ;      public int x = 7;
	//RTN  ;}
	//FUNC MAIN ;void main() {
	//NEWI Cl1, Tv17 ;     Dog d = new Dog();
	//FRAME Tv17, Co3 ;     Dog d = new Dog();
	//CALL Co3 ;     Dog d = new Dog();
	//PEEK Tv17 ;     Dog d = new Dog();
	//MOV Lv6, Tv17 ;     Dog d = new Dog();
	//FRAME Lv6, Me5 ;     d.AddX(d);
	//PUSH Lv6 ;     d.AddX(d);
	//CALL Me5 ;     d.AddX(d);
	//PEEK Tv18 ;     d.AddX(d);
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
	a.GetNext(false)
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
	//El17: []int{13}
	//El8: []int{5}
	//En22: []int{17, 21}
	//If15: []int{11, 14}
	//If6: []int{3, 6}
	//Li16: []int{12}
	//Li18: []int{14, 18}
	//Li3: []int{1}
	//Li4: []int{2}
	//Li7: []int{4}
	//Li9: []int{6, 7, 8, 16}
	//Lv1: []int{1, 2, 4, 6, 7, 8, 12, 14, 15, 16, 18, 19}
	//MAIN: []int{0}
	//Tv10: []int{6, 10}
	//Tv11: []int{7, 9}
	//Tv12: []int{8, 9}
	//Tv13: []int{9, 10}
	//Tv14: []int{10, 11}
	//Tv19: []int{14, 15}
	//Tv21: []int{16, 17}
	//Tv23: []int{18, 19}
	//Tv5: []int{2, 3}
	//Wh20: []int{13, 5, 16, 20}
	//Rows:
	//FUNC MAIN ;void main() {
	//MOV Lv1, Li3 ;    int x = 7;
	//LT Tv5, Lv1, Li4 ;    if ( x < 4) {
	//BF Tv5, If6 ;    if ( x < 4) {
	//MOV Lv1, Li7 ;	x = 8;
	//JMP Wh20 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//If6: LTE Tv10, Lv1, Li9 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//GTE Tv11, Lv1, Li9 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//NEQ Tv12, Lv1, Li9 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//AND Tv13, Tv11, Tv12 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//OR Tv14, Tv10, Tv13 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//BF Tv14, If15 ;    } else if ( x <= 3 || x >= 3 && x != 3) {
	//MOV Lv1, Li16 ;	x = 2;
	//JMP Wh20 ;    } else {
	//If15: ADD Tv19, Li18, Lv1 ;	x = x + 1;
	//MOV Lv1, Tv19 ;	x = x + 1;
	//Wh20: GT Tv21, Lv1, Li9 ;    while(x > 3) {
	//BF Tv21, En22 ;    while(x > 3) {
	//SUB Tv23, Li18, Lv1 ;	x = x - 1;
	//MOV Lv1, Tv23 ;	x = x - 1;
	//JMP Wh20 ;}
	//En22: RTN  ;}
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
	a.GetNext(false)
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
	//Li17: []int{10, 31}
	//Li27: []int{43}
	//Li28: []int{43}
	//Lv7: []int{26, 27, 32, 40}
	//Lv8: []int{39, 45}
	//MAIN: []int{21}
	//Me4: []int{4, 27, 28, 32, 33, 35, 36}
	//Me5: []int{7}
	//Me6: []int{13, 40, 41}
	//St9: []int{1, 2, 17}
	//Tv11: []int{18, 19}
	//Tv15: []int{8, 11}
	//Tv16: []int{9, 10}
	//Tv18: []int{10, 11}
	//Tv19: []int{14, 15}
	//Tv20: []int{22, 23, 25, 26}
	//Tv21: []int{29, 30}
	//Tv22: []int{30, 31}
	//Tv23: []int{34, 35}
	//Tv24: []int{37, 38}
	//Tv25: []int{38, 39}
	//Tv26: []int{42, 44}
	//Tv29: []int{43, 44}
	//Tv30: []int{44, 45}
	//this: []int{1, 3, 5, 8, 9, 14, 18}
	//Rows:
	//FUNC Co3 ;      Dog() {
	//FRAME this, St9 ;      Dog() {
	//CALL St9 ;      Dog() {
	//RETURN this ;      }
	//FUNC Me4 ;      public Dog GetDog() {
	//RETURN this ;             return this;
	//RTN  ;      }
	//FUNC Me5 ;      public void Bite() {
	//REF Tv15, Iv2, this ;             bites = bites + 1;
	//REF Tv16, Iv2, this ;             bites = bites + 1;
	//ADD Tv18, Li17, Tv16 ;             bites = bites + 1;
	//MOV Tv15, Tv18 ;             bites = bites + 1;
	//RTN  ;      }
	//FUNC Me6 ;      public int GetBites() {
	//REF Tv19, Iv2, this ;             return bites;
	//RETURN Tv19 ;             return bites;
	//RTN  ;      }
	//FUNC St9 ;}
	//REF Tv11, Iv2, this ;      public int bites = 0;
	//MOV Tv11, Li12 ;      public int bites = 0;
	//RTN  ;}
	//FUNC MAIN ;void main() {
	//NEWI Cl1, Tv20 ;     Dog d = new Dog();
	//FRAME Tv20, Co3 ;     Dog d = new Dog();
	//CALL Co3 ;     Dog d = new Dog();
	//PEEK Tv20 ;     Dog d = new Dog();
	//MOV Lv7, Tv20 ;     Dog d = new Dog();
	//FRAME Lv7, Me4 ;     d.GetDog().bites = 1;
	//CALL Me4 ;     d.GetDog().bites = 1;
	//PEEK Tv21 ;     d.GetDog().bites = 1;
	//REF Tv22, Iv2, Tv21 ;     d.GetDog().bites = 1;
	//MOV Tv22, Li17 ;     d.GetDog().bites = 1;
	//FRAME Lv7, Me4 ;     b = d.GetDog().GetDog().bites;
	//CALL Me4 ;     b = d.GetDog().GetDog().bites;
	//PEEK Tv23 ;     b = d.GetDog().GetDog().bites;
	//FRAME Tv23, Me4 ;     b = d.GetDog().GetDog().bites;
	//CALL Me4 ;     b = d.GetDog().GetDog().bites;
	//PEEK Tv24 ;     b = d.GetDog().GetDog().bites;
	//REF Tv25, Iv2, Tv24 ;     b = d.GetDog().GetDog().bites;
	//MOV Lv8, Tv25 ;     b = d.GetDog().GetDog().bites;
	//FRAME Lv7, Me6 ;     b = d.GetBites() + 7 * 3;
	//CALL Me6 ;     b = d.GetBites() + 7 * 3;
	//PEEK Tv26 ;     b = d.GetBites() + 7 * 3;
	//MUL Tv29, Li28, Li27 ;     b = d.GetBites() + 7 * 3;
	//ADD Tv30, Tv29, Tv26 ;     b = d.GetBites() + 7 * 3;
	//MOV Lv8, Tv30 ;     b = d.GetBites() + 7 * 3;
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
	a.GetNext(false)
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
	//Li12: []int{9}
	//Li13: []int{9}
	//Li15: []int{10}
	//Li16: []int{10}
	//Lv6: []int{11}
	//MAIN: []int{8}
	//Me5: []int{4, 11, 14}
	//St7: []int{1, 2, 6}
	//Tv14: []int{9, 12}
	//Tv17: []int{10, 13}
	//Tv18: []int{15}
	//this: []int{1, 3}
	//Rows:
	//FUNC Co2 ;      Apple() {
	//FRAME this, St7 ;      Apple() {
	//CALL St7 ;      Apple() {
	//RETURN this ;      }
	//FUNC Me5 ;      public void MyFunc(int i, bool j) {
	//RTN  ;      }
	//FUNC St7 ;}
	//RTN  ;}
	//FUNC MAIN ;void main() {
	//ADD Tv14, Li13, Li12 ;     a.MyFunc(1 + 3, 4 < 7);
	//LT Tv17, Li15, Li16 ;     a.MyFunc(1 + 3, 4 < 7);
	//FRAME Lv6, Me5 ;     a.MyFunc(1 + 3, 4 < 7);
	//PUSH Tv14 ;     a.MyFunc(1 + 3, 4 < 7);
	//PUSH Tv17 ;     a.MyFunc(1 + 3, 4 < 7);
	//CALL Me5 ;     a.MyFunc(1 + 3, 4 < 7);
	//PEEK Tv18 ;     a.MyFunc(1 + 3, 4 < 7);
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
	a.GetNext(false)
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
	//Num Rows: 24 curRow: 23
	//Lables:
	//4: []int{10}
	//Cl1: []int{14}
	//Co4: []int{0, 15, 17}
	//Iv2: []int{3, 7}
	//Li11: []int{10}
	//Li14: []int{13, 21}
	//Li16: []int{16, 20}
	//Lv5: []int{12, 13, 20, 21}
	//MAIN: []int{9}
	//Pa3: []int{4}
	//St6: []int{1, 2, 6}
	//Tv12: []int{11, 12}
	//Tv13: []int{10, 11}
	//Tv15: []int{13, 19}
	//Tv17: []int{14, 15, 18, 19}
	//Tv18: []int{20, 22}
	//Tv19: []int{21, 22}
	//Tv8: []int{7}
	//Tv9: []int{3, 4}
	//this: []int{1, 3, 5, 7}
	//Rows:
	//FUNC Co4 ;    Frog(int i) {
	//FRAME this, St6 ;    Frog(int i) {
	//CALL St6 ;    Frog(int i) {
	//REF Tv9, Iv2, this ;	f = i;
	//MOV Tv9, Pa3 ;	f = i;
	//RETURN this ;    }
	//FUNC St6 ;}
	//REF Tv8, Iv2, this ;    private int f;
	//RTN  ;}
	//FUNC MAIN ;void main() {
	//MUL Tv13, 4, Li11 ;    Frog frogs[] = new Frog[10];
	//NEW Tv13, Tv12 ;    Frog frogs[] = new Frog[10];
	//MOV Lv5, Tv12 ;    Frog frogs[] = new Frog[10];
	//AEF Tv15, Li14, Lv5 ;    frogs[0] = new Frog(1);
	//NEWI Cl1, Tv17 ;    frogs[0] = new Frog(1);
	//FRAME Tv17, Co4 ;    frogs[0] = new Frog(1);
	//PUSH Li16 ;    frogs[0] = new Frog(1);
	//CALL Co4 ;    frogs[0] = new Frog(1);
	//PEEK Tv17 ;    frogs[0] = new Frog(1);
	//MOV Tv15, Tv17 ;    frogs[0] = new Frog(1);
	//AEF Tv18, Li16, Lv5 ;    frogs[1] = frogs[0];
	//AEF Tv19, Li14, Lv5 ;    frogs[1] = frogs[0];
	//MOV Tv18, Tv19 ;    frogs[1] = frogs[0];
	//RTN  ;}
	//Num Rows: 0 curRow: -1
	//Lables:
	//Rows:

}
