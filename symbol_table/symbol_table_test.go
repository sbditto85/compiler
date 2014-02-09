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

	sym.AddElement("myclass","Class",d)
	sym.PrintTable()
	//Output:
	//Current Scope: g.myclass
	//=================
	//Elements:
	//Scope: g, SymId: C1, Value: myclass, Kind: Class
	//Extra Data:
	//--------------
}

func ExampleAddSymbolClassFuncThenRemove() {
	sym := NewSymbolTable()

	dClass := make(map[string]interface{})
	dMethod := make(map[string]interface{})

	sym.AddElement("myclass","Class",dClass)
	dMethod["additional"] = "more"
	sym.AddElement("myfun","Method",dMethod)
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

	sym.AddElement("othermethod","Method",otherMethod)
	sym.PrintTable()
	fmt.Println()
	sym.DownScope()
	sym.PrintTable()
	fmt.Println()



	//Output:
	//Current Scope: g.myclass.myfun
	//=================
	//Elements:
	//Scope: g, SymId: C1, Value: myclass, Kind: Class
	//Extra Data:
	//--------------
	//Scope: g.myclass, SymId: M2, Value: myfun, Kind: Method
	//Extra Data:
	//Key: additional, Value: more
	//--------------
	//
	//Current Scope: g.myclass
	//=================
	//Elements:
	//Scope: g, SymId: C1, Value: myclass, Kind: Class
	//Extra Data:
	//--------------
	//Scope: g.myclass, SymId: M2, Value: myfun, Kind: Method
	//Extra Data:
	//Key: additional, Value: more
	//--------------
	//
	//Current Scope: g
	//=================
	//Elements:
	//Scope: g, SymId: C1, Value: myclass, Kind: Class
	//Extra Data:
	//--------------
	//Scope: g.myclass, SymId: M2, Value: myfun, Kind: Method
	//Extra Data:
	//Key: additional, Value: more
	//--------------
	//
	//Current Scope: g.othermethod
	//=================
	//Elements:
	//Scope: g, SymId: C1, Value: myclass, Kind: Class
	//Extra Data:
	//--------------
	//Scope: g.myclass, SymId: M2, Value: myfun, Kind: Method
	//Extra Data:
	//Key: additional, Value: more
	//--------------
	//Scope: g, SymId: M3, Value: othermethod, Kind: Method
	//Extra Data:
	//Key: testing, Value: testing
	//--------------
	//
	//Current Scope: g
	//=================
	//Elements:
	//Scope: g, SymId: C1, Value: myclass, Kind: Class
	//Extra Data:
	//--------------
	//Scope: g.myclass, SymId: M2, Value: myfun, Kind: Method
	//Extra Data:
	//Key: additional, Value: more
	//--------------
	//Scope: g, SymId: M3, Value: othermethod, Kind: Method
	//Extra Data:
	//Key: testing, Value: testing
	//--------------
}
