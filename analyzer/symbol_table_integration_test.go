package analyzer

import (
	//"testing"
	lex "github.com/sbditto85/compiler/lexer"
	//tok "github.com/sbditto85/compiler/token"
	"fmt"
)

func ExampleSymbolTableIntergrationBasicFile() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	l := lex.NewLexer()
	l.ReadFile("basic_class.kxi")

	a := NewAnalyzer(l, false)
	a.IsCompilationUnit()

	fmt.Println("Valid")

	//a.PrintSymbolTable()
	a.PrintTableInAddOrder()

	//Output:
	//Valid
	//Current Scope: g
	//=================
	//Elements:
	//Scope: g, SymId: Cl1, Value: Cat, Kind: Class
	//Extra Data:
	//Key: size, Value: 13
	//--------------
	//Scope: g.Cat, SymId: Iv2, Value: paw, Kind: Ivar
	//Extra Data:
	//Key: accessMod, Value: private
	//Key: type, Value: char
	//Key: isArray, Value: false
	//Key: this_class, Value: Cat
	//Key: offset, Value: 0
	//--------------
	//Scope: g.Cat, SymId: Iv3, Value: legsNum, Kind: Ivar
	//Extra Data:
	//Key: accessMod, Value: public
	//Key: type, Value: int
	//Key: isArray, Value: false
	//Key: this_class, Value: Cat
	//Key: offset, Value: 1
	//--------------
	//Scope: g.Cat, SymId: Iv4, Value: myArr, Kind: Ivar
	//Extra Data:
	//Key: accessMod, Value: private
	//Key: type, Value: bool
	//Key: isArray, Value: true
	//Key: this_class, Value: Cat
	//Key: offset, Value: 5
	//--------------
	//Scope: g.Cat.Cat, SymId: Pa5, Value: legsNum, Kind: Parameter
	//Extra Data:
	//Key: type, Value: int
	//Key: isArray, Value: false
	//Key: offset, Value: 0
	//--------------
	//Scope: g.Cat, SymId: Co6, Value: Cat, Kind: Constructor
	//Extra Data:
	//Key: class, Value: Cat
	//Key: type, Value: Cat
	//Key: parameters, Value: [{int legsNum false}]
	//Key: accessMod, Value: public
	//Key: paramSymIds, Value: [Pa5]
	//Key: size, Value: 4
	//--------------
	//Scope: g.Cat, SymId: Iv7, Value: c, Kind: Ivar
	//Extra Data:
	//Key: accessMod, Value: private
	//Key: type, Value: char
	//Key: isArray, Value: true
	//Key: this_class, Value: Cat
	//Key: offset, Value: 9
	//--------------
	//Scope: g.Cat.Run, SymId: Pa8, Value: i, Kind: Parameter
	//Extra Data:
	//Key: type, Value: int
	//Key: isArray, Value: false
	//Key: offset, Value: 0
	//--------------
	//Scope: g.Cat.Run, SymId: Pa9, Value: c, Kind: Parameter
	//Extra Data:
	//Key: type, Value: char
	//Key: isArray, Value: false
	//Key: offset, Value: 4
	//--------------
	//Scope: g.Cat, SymId: Me10, Value: Run, Kind: Method
	//Extra Data:
	//Key: accessMod, Value: public
	//Key: type, Value: bool
	//Key: parameters, Value: [{int i false} {char c false}]
	//Key: paramSymIds, Value: [Pa8 Pa9]
	//Key: size, Value: 9
	//--------------
	//Scope: g.Cat.Run, SymId: Lv11, Value: x, Kind: Lvar
	//Extra Data:
	//Key: isArray, Value: false
	//Key: type, Value: int
	//Key: offset, Value: 5
	//--------------
	//Scope: g, SymId: MAIN, Value: main, Kind: Main
	//Extra Data:
	//Key: type, Value: void
	//Key: size, Value: 4
	//--------------
	//Scope: g.main, SymId: Lv12, Value: c, Kind: Lvar
	//Extra Data:
	//Key: isArray, Value: false
	//Key: type, Value: Cat
	//Key: offset, Value: 0
	//--------------
}
