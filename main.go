package main

import (
	"fmt"
	an "github.com/sbditto85/compiler/analyzer"
	lex "github.com/sbditto85/compiler/lexer"
	tc "github.com/sbditto85/compiler/tcode"
	tok "github.com/sbditto85/compiler/token"
	amb "github.com/sbditto85/virtualmachine/assembler"
	vm "github.com/sbditto85/virtualmachine/virtualmachine"
	"os"
)

func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println(r)
	// 	}
	// }()
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: compiler file.kxi [file.kxi ...]")
		return
	}
	for _, file := range os.Args[1:] {
		l := lex.NewLexer()
		l.ReadFile(file)

		a := an.NewAnalyzer(l, false)
		a.GetNext(false)
		err := a.PerformPass()

		if err != nil {
			fmt.Println(err.Error())
		}

		curTok, err := l.GetCurrentToken()
		if curTok.Type != tok.EOT {
			panic(fmt.Sprintf("Last token parsed is %s\n", curTok.Lexeme))
		}
		if err != nil {
			panic("Error getting last token!")
		}

		l = lex.NewLexer()
		l.ReadFile(file)
		a.SetLexer(l)

		err = a.PerformNextPass(false)
		if err != nil {
			panic(err.Error())
		}

		curTok, err = l.GetCurrentToken()
		if curTok.Type != tok.EOT {
			panic(fmt.Sprintf("Last token not EOT it is %s\n", curTok.Lexeme))
		}
		if err != nil {
			panic("Error getting last token!")
		}

		table, symbolTable := a.GetICodeInfo()

		asm := tc.GenerateASM(table, symbolTable)

		fmt.Printf("Enter asm file name for %s: ", file)
		asmFileName := ""
		fmt.Scanf("%s", &asmFileName)

		if asmFileName != "" {
			asmFile, err := os.Create(asmFileName)
			if err != nil {
				panic(err)
			}
			defer func() {
				if err := asmFile.Close(); err != nil {
					panic(err)
				}
			}()
			for _, line := range asm {
				if _, err := asmFile.WriteString(line + "\n"); err != nil {
					panic(err)
				}
			}
		}

		fmt.Printf("Would you like to run the assembly? (y/n): ")
		yN := ""
		fmt.Scanf("%s", &yN)

		switch yN {
		case "y", "Y":
			assembler := amb.NewAssembler()
			assembler.ReadStrings(asm)

			fperr := assembler.FirstPass()
			if fperr == nil {
				sperr := assembler.SecondPass()
				if sperr != nil {
					panic(sperr.Error())
				}
			} else {
				panic(fperr.Error())
			}

			v := vm.NewVirtualMachine(assembler.GetBytes())
			verr := v.Run()
			if verr != nil {
				panic(verr.Error())
			}
		}
	}
}
