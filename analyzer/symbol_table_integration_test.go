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

	a.PrintSymbolTable()

	//Output:
	//Valid

}


