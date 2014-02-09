package analyzer

import (
	//"testing"
	lex "github.com/sbditto85/compiler/lexer"
	//tok "github.com/sbditto85/compiler/token"
	"fmt"
) 

func ExampleSymbolTableIntergrationBasicFile() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	l := lex.NewLexer()
	l.ReadFile("basic_class.kxi")

	a := NewAnalyzer(l,false)
	a.IsCompilationUnit()

	fmt.Println("Valid")

	//a.PrintSymbolTable()
	a.PrintTableInAddOrder()

	//Output:
	//Valid
	//Current Scope: g
	//=================
	//Elements:
	//Scope: g, SymId: C1, Value: Cat, Kind: Class
	//Extra Data:
	//--------------
	//Scope: g.Cat, SymId: I2, Value: paw, Kind: Ivar
	//Extra Data:
	//Key: accessMod, Value: private
	//Key: type, Value: char
	//Key: isArray, Value: false
	//--------------
	//Scope: g.Cat, SymId: I3, Value: legsNum, Kind: Ivar
	//Extra Data:
	//Key: accessMod, Value: public
	//Key: type, Value: int
	//Key: isArray, Value: false
	//--------------
	//Scope: g.Cat, SymId: I4, Value: myArr, Kind: Ivar
	//Extra Data:
	//Key: accessMod, Value: private
	//Key: type, Value: bool
	//Key: isArray, Value: true
	//--------------
	//Scope: g.Cat, SymId: C5, Value: Cat, Kind: Constructor
	//Extra Data:
	//Key: class, Value: Cat
	//Key: parameters, Value: []
	//Key: paramSymIds, Value: []
	//--------------
	//Scope: g.Cat, SymId: I6, Value: c, Kind: Ivar
	//Extra Data:
	//Key: accessMod, Value: private
	//Key: type, Value: char
	//Key: isArray, Value: true
	//--------------
	//Scope: g.Cat.Run, SymId: P7, Value: i, Kind: Parameter
	//Extra Data:
	//Key: type, Value: int
	//Key: isArray, Value: false
	//--------------
	//Scope: g.Cat.Run, SymId: P8, Value: c, Kind: Parameter
	//Extra Data:
	//Key: type, Value: char
	//Key: isArray, Value: false
	//--------------
	//Scope: g.Cat, SymId: M9, Value: Run, Kind: Method
	//Extra Data:
	//Key: accessMod, Value: public
	//Key: type, Value: bool
	//Key: parameters, Value: [{int i false} {char c false}]
	//Key: paramSymIds, Value: [P7 P8]
	//--------------
	//Scope: g.Cat.Run, SymId: L10, Value: x, Kind: Lvar
	//Extra Data:
	//Key: isArray, Value: false
	//Key: type, Value: int
	//--------------
	//Scope: g, SymId: M11, Value: main, Kind: Main
	//Extra Data:
	//Key: type, Value: void
	//--------------
	//Scope: g.main, SymId: L12, Value: c, Kind: Lvar
	//Extra Data:
	//Key: isArray, Value: false
	//Key: type, Value: Cat
	//--------------

}


