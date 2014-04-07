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
	//10
	//a
	//n
}

func ExampleTCodeMath() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "tests/math.kxi"
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
	//-> 0
	//-> 2
	//-> 1
	//-> -1
	//6
	//1
	//-3
	//0

}

func ExampleTCodeLoop() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "tests/loop.kxi"
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
	//yyyyy
	//fffftttt

}

func ExampleTCodeNewObj() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "tests/newobj.kxi"
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
	//Emmy = 8
	//Emmy = 8
	//Emmy is 8!
	//Emmy is 8!
	//Emmy is 9!
	//Emmy is 8!
	//Emmy is 10!
	//Emmy is 8!
	//------
	//Emmy is 10!
	//Emmy is 9!
	//Emmy is 10!
	//Emmy is 10!

}

func ExampleTCodeNewComplex() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "tests/newcomplex.kxi"
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
	//a = 8
	//Emmy = 9
	//a = 8
	//Emmy = 7
	//a = 8
	//Emmy = 3
	//Emmy is 9!
	//Emmy is 7!
	//Emmy is 11!
	//Emmy is 7!
	//Emmy is 12!
	//Emmy is 7!
	//------
	//Emmy is 12!
	//Emmy is 11!
	//Emmy is 12!
	//Emmy is 13!
	//------
	//Emmy is 3!
	//Emmy is 13!

}

func ExampleTCodeBasicArray() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	file := "tests/basicarray.kxi"
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
	//5

}
