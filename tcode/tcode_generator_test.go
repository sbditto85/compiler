package tcode

import (
	//"testing"
	"fmt"
	an "github.com/sbditto85/compiler/analyzer"
	lex "github.com/sbditto85/compiler/lexer"
	tok "github.com/sbditto85/compiler/token"
	amb "github.com/sbditto85/virtualmachine/assembler"
	vm "github.com/sbditto85/virtualmachine/virtualmachine"
)

func ExampleTCodeMain() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "tests/main.kxi"
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

	table, symbolTable := a.GetICodeInfo()

	asm := GenerateASM(table, symbolTable)

	/* 
	fmt.Printf("ASM:\n")

	for i, line := range asm {
		fmt.Printf("%d : %s\n", i+1, line)
	}
	*/

	assembler := amb.NewAssembler()
	assembler.ReadStrings(asm)

	fperr := assembler.FirstPass()
	if fperr == nil {
		sperr := assembler.SecondPass()
		if sperr == nil {
			sperr = sperr
		} else {
			fmt.Println(sperr)
		}
	} else {
		fmt.Println(fperr)
	}

	v := vm.NewVirtualMachine(assembler.GetBytes())
	verr := v.Run()
	if verr != nil {
		fmt.Printf("%s\n", verr.Error())
	}

	//Output:
	//2
	//x
	//y

}

func ExampleTCodeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "tests/function.kxi"
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

	table, symbolTable := a.GetICodeInfo()

	asm := GenerateASM(table, symbolTable)
 
	fmt.Printf("ASM:\n")

	for i, line := range asm {
		fmt.Printf("%d : %s\n", i+1, line)
	}

	assembler := amb.NewAssembler()
	assembler.ReadStrings(asm)

	fperr := assembler.FirstPass()
	if fperr == nil {
		sperr := assembler.SecondPass()
		if sperr == nil {
			sperr = sperr
		} else {
			fmt.Println(sperr)
		}
	} else {
		fmt.Println(fperr)
	}

	v := vm.NewVirtualMachine(assembler.GetBytes())
	verr := v.Run()
	if verr != nil {
		fmt.Printf("%s\n", verr.Error())
	}

	//Output:
	//
}
