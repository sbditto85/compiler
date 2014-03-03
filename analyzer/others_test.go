package analyzer

import (
	"fmt"
	lex "github.com/sbditto85/compiler/lexer"
	"path/filepath"
)

func ExampleOthersTests() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	files, err := filepath.Glob("./otherstests/*.kxi")

	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Printf("%d\n",len(files))

	for _, f := range files {
		fmt.Println(f)

		l := lex.NewLexer()
		l.ReadFile(f)

		a := NewAnalyzer(l, false)
		a.IsCompilationUnit()

		fmt.Println("Valid")

	}

	//Output:
	//otherstests/DemoA.kxi
	//Valid
	//otherstests/DemoB.kxi
	//Valid
	//otherstests/DemoBV2.kxi
	//Valid
	//otherstests/DemoC.kxi
	//Valid
	//otherstests/DemoCMinus.kxi
	//Valid
	//otherstests/Returns.kxi
	//Valid
	//otherstests/class.kxi
	//Valid
	//otherstests/dcSimple.kxi
	//Valid
	//otherstests/fib.kxi
	//Valid
	//otherstests/icode.kxi
	//Valid
	//otherstests/main.kxi
	//Valid
	//otherstests/main2.kxi
	//Valid
	//otherstests/main3.kxi
	//Valid
	//otherstests/target.kxi
	//Valid
	//otherstests/whiletest.kxi
	//Valid

}
