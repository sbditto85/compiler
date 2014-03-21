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
	//--------------
	//Scope: g.Cat, SymId: Iv2, Value: paw, Kind: Ivar
	//Extra Data:
	//Key: accessMod, Value: private
	//Key: type, Value: char
	//Key: isArray, Value: false
	//--------------
	//Scope: g.Cat, SymId: Iv3, Value: legsNum, Kind: Ivar
	//Extra Data:
	//Key: accessMod, Value: public
	//Key: type, Value: int
	//Key: isArray, Value: false
	//--------------
	//Scope: g.Cat, SymId: Iv4, Value: myArr, Kind: Ivar
	//Extra Data:
	//Key: accessMod, Value: private
	//Key: type, Value: bool
	//Key: isArray, Value: true
	//--------------
	//Scope: g.Cat.Cat, SymId: Pa5, Value: legsNum, Kind: Parameter
	//Extra Data:
	//Key: type, Value: int
	//Key: isArray, Value: false
	//--------------
	//Scope: g.Cat, SymId: Co6, Value: Cat, Kind: Constructor
	//Extra Data:
	//Key: class, Value: Cat
	//Key: type, Value: Cat
	//Key: parameters, Value: [{int legsNum false}]
	//Key: accessMod, Value: public
	//Key: paramSymIds, Value: [Pa5]
	//--------------
	//Scope: g.Cat, SymId: Iv7, Value: c, Kind: Ivar
	//Extra Data:
	//Key: accessMod, Value: private
	//Key: type, Value: char
	//Key: isArray, Value: true
	//--------------
	//Scope: g.Cat.Run, SymId: Pa8, Value: i, Kind: Parameter
	//Extra Data:
	//Key: type, Value: int
	//Key: isArray, Value: false
	//--------------
	//Scope: g.Cat.Run, SymId: Pa9, Value: c, Kind: Parameter
	//Extra Data:
	//Key: type, Value: char
	//Key: isArray, Value: false
	//--------------
	//Scope: g.Cat, SymId: Me10, Value: Run, Kind: Method
	//Extra Data:
	//Key: accessMod, Value: public
	//Key: type, Value: bool
	//Key: parameters, Value: [{int i false} {char c false}]
	//Key: paramSymIds, Value: [Pa8 Pa9]
	//--------------
	//Scope: g.Cat.Run, SymId: Lv11, Value: x, Kind: Lvar
	//Extra Data:
	//Key: isArray, Value: false
	//Key: type, Value: int
	//--------------
	//Scope: g, SymId: Ma12, Value: main, Kind: Main
	//Extra Data:
	//Key: type, Value: void
	//--------------
	//Scope: g.main, SymId: Lv13, Value: c, Kind: Lvar
	//Extra Data:
	//Key: isArray, Value: false
	//Key: type, Value: Cat
	//--------------
}
