package analyzer

import (
	//"testing"
	lex "github.com/sbditto85/compiler/lexer"
	tok "github.com/sbditto85/compiler/token"
	"fmt"
) 

func ExampleXEqualYPass2() {
	defer func(){
		if r:= recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/xequaly.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := NewAnalyzer(l,false)
	a.GetNext()
	err := a.PerformPass()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	l = lex.NewLexer()
	l.ReadFile(file)
	a.SetLexer(l)

	err = a.PerformNextPass(true)
	if err != nil {
		fmt.Println(err.Error())
	}

	curTok,err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s",curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}
	
	//Output:
	//IPush: x from scope g.main
	//IExists!
	//Pushed operator =
	//IPush: y from scope g.main
	//IExists!
	//SM: Testing operation = ...
	//SM: ... finished operation =
	//EOE
}
