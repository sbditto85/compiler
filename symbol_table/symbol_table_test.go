package symbol_table

import (
	//"testing"
	//lex "github.com/sbditto85/compiler/lexer"
	//tok "github.com/sbditto85/compiler/token"
	"fmt"
)

func ExampleDefaultSymbolTable() {
	sym := NewSymbolTable()
	sym.PrintTable()
	//Output:
	//Current Scope: g
	//=================
	//Elements:
}

func ExampleAddSymbol() {
	sym := NewSymbolTable()

	d := make(map[string]interface{})

	sym.AddElement("myclass", "Class", d, true)
	sym.PrintTable()
	//Output:
	//Current Scope: g.myclass
	//=================
	//Elements:
	//Scope: g, SymId: Cl1, Value: myclass, Kind: Class
	//Extra Data:
	//--------------
}

func ExampleAddSymbolClassFuncThenRemove() {
	sym := NewSymbolTable()

	dClass := make(map[string]interface{})
	dMethod := make(map[string]interface{})

	sym.AddElement("myclass", "Class", dClass, true)
	dMethod["additional"] = "more"
	sym.AddElement("myfun", "Method", dMethod, true)
	sym.PrintTable()
	fmt.Println()
	sym.DownScope()
	sym.PrintTable()
	fmt.Println()
	sym.DownScope()
	sym.PrintTable()
	fmt.Println()

	otherMethod := make(map[string]interface{})
	otherMethod["testing"] = "testing"

	sym.AddElement("othermethod", "Method", otherMethod, true)
	sym.PrintTable()
	fmt.Println()
	sym.DownScope()
	sym.PrintTable()
	fmt.Println()

	//Output:
	//Current Scope: g.myclass.myfun
	//=================
	//Elements:
	//Scope: g, SymId: Cl1, Value: myclass, Kind: Class
	//Extra Data:
	//--------------
	//Scope: g.myclass, SymId: Me2, Value: myfun, Kind: Method
	//Extra Data:
	//Key: additional, Value: more
	//Key: this_class, Value: myclass
	//Key: size, Value: 0
	//--------------
	//
	//Current Scope: g.myclass
	//=================
	//Elements:
	//Scope: g, SymId: Cl1, Value: myclass, Kind: Class
	//Extra Data:
	//--------------
	//Scope: g.myclass, SymId: Me2, Value: myfun, Kind: Method
	//Extra Data:
	//Key: additional, Value: more
	//Key: this_class, Value: myclass
	//Key: size, Value: 0
	//--------------
	//
	//Current Scope: g
	//=================
	//Elements:
	//Scope: g, SymId: Cl1, Value: myclass, Kind: Class
	//Extra Data:
	//--------------
	//Scope: g.myclass, SymId: Me2, Value: myfun, Kind: Method
	//Extra Data:
	//Key: additional, Value: more
	//Key: this_class, Value: myclass
	//Key: size, Value: 0
	//--------------
	//
	//Current Scope: g.othermethod
	//=================
	//Elements:
	//Scope: g, SymId: Cl1, Value: myclass, Kind: Class
	//Extra Data:
	//--------------
	//Scope: g.myclass, SymId: Me2, Value: myfun, Kind: Method
	//Extra Data:
	//Key: additional, Value: more
	//Key: this_class, Value: myclass
	//Key: size, Value: 0
	//--------------
	//Scope: g, SymId: Me3, Value: othermethod, Kind: Method
	//Extra Data:
	//Key: testing, Value: testing
	//Key: this_class, Value: g
	//Key: size, Value: 0
	//--------------
	//
	//Current Scope: g
	//=================
	//Elements:
	//Scope: g, SymId: Cl1, Value: myclass, Kind: Class
	//Extra Data:
	//--------------
	//Scope: g.myclass, SymId: Me2, Value: myfun, Kind: Method
	//Extra Data:
	//Key: additional, Value: more
	//Key: this_class, Value: myclass
	//Key: size, Value: 0
	//--------------
	//Scope: g, SymId: Me3, Value: othermethod, Kind: Method
	//Extra Data:
	//Key: testing, Value: testing
	//Key: this_class, Value: g
	//Key: size, Value: 0
	//--------------
}
