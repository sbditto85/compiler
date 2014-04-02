package main

import (
	"fmt"
	an "github.com/sbditto85/compiler/analyzer"
	lex "github.com/sbditto85/compiler/lexer"
	tc "github.com/sbditto85/compiler/tcode"
	tok "github.com/sbditto85/compiler/token"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "passtwo/xequaly.kxi"
	l := lex.NewLexer()
	l.ReadFile(file)

	a := an.NewAnalyzer(l, false)
	a.GetNext()
	err := a.PerformPass()

	if err != nil {
		fmt.Println(err.Error())
	}

	curTok, err := l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s", curTok.Lexeme)
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

	curTok, err = l.GetCurrentToken()
	if curTok.Type != tok.EOT {
		fmt.Printf("Last token not EOT it is %s", curTok.Lexeme)
	}
	if err != nil {
		fmt.Println("Error getting last token!")
	}

	table, symbolTable := a.GetICodeInfo()

	asm := tc.GenerateASM(table, symbolTable)

	fmt.Printf("ASM: %#v\n", asm)
}
