package tcode

import (
	"fmt"
	ic "github.com/sbditto85/compiler/analyzer/icode"
	sym "github.com/sbditto85/compiler/symbol_table"
	"strconv"
)

func GenerateASM(table *ic.Quad, st *sym.SymbolTable) (asm []string) {
	asm = make([]string, 0, table.Size())

	//table.Print()
	//st.PrintTable()

	mainElem := st.GetElementInScope("g", "main")
	mainSize, _ := sym.IntFromData(mainElem.Data, "size")

	//debug?
	//asm = append(asm, `TRP     #99`)

	//setup heap
	asm = append(asm, `LDA     R9 FREE:`)

	//Go to main
	asm = append(asm, `;; Call function "MAIN:"`)
	asm = append(asm, `;; Test for overflow`)
	asm = append(asm, `MOV     R10 RSP`)
	asm = append(asm, fmt.Sprintf(`ADI     R10 #%d          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)`, -12+(mainSize*-1)))
	asm = append(asm, `CMP     R10 RSL`)
	asm = append(asm, `BLT     R10 OVRFLW:`)
	asm = append(asm, `;; Create Activation Record and invoke MAIN`)
	asm = append(asm, `MOV     R10 RFP`)
	asm = append(asm, `MOV     RFP RSP`)
	asm = append(asm, `ADI     RSP #-4`)
	asm = append(asm, `STR     R10 (RSP)`)
	asm = append(asm, `ADI     RSP #-4`)
	asm = append(asm, `;; this`)
	asm = append(asm, `SUB     R1 R1           ; get this from where its at`)
	asm = append(asm, `STR     R1 (RSP)`)
	asm = append(asm, `ADI     RSP #-4`)
	asm = append(asm, `;; parameters on the stack`)
	asm = append(asm, `;; local varibales on the stack`)
	//add main's local variables
	for _, e := range getLocalVars(st, mainElem) {
		typ, _ := sym.StringFromData(e.Data, "type")
		isArray, _ := sym.BoolFromData(e.Data, "isArray")
		asm = append(asm, fmt.Sprintf(`ADI     RSP #%d`, (sym.SizeOfType(typ, isArray)*-1)))
	}

	asm = append(asm, `;; Temp variables on the stack`)
	//add main's temp variables
	for _, e := range getTempVars(st, mainElem) {
		typ, _ := sym.StringFromData(e.Data, "type")
		isArray, _ := sym.BoolFromData(e.Data, "isArray")
		asm = append(asm, fmt.Sprintf(`ADI     RSP #%d`, (sym.SizeOfType(typ, isArray)*-1)))
	}

	asm = append(asm, `;; set the return address and jump`)
	asm = append(asm, `MOV     R10 RPC         ; PC already at next instruction`)
	asm = append(asm, `ADI     R10 #12`)
	asm = append(asm, `STR     R10 (RFP)`)
	asm = append(asm, `JMP     MAIN:`)

	//Exits
	asm = append(asm, `EXIT:   TRP     #0`)
	asm = append(asm, `OVRFLW: LDB     R0 LTRCO:`)
	asm = append(asm, `TRP     #3`)
	asm = append(asm, `LDB     R0 NL:`)
	asm = append(asm, `TRP     #3`)
	asm = append(asm, `TRP     #0`)
	asm = append(asm, `HOVRFLW: LDB     R0 LTRCH:`)
	asm = append(asm, `TRP     #3`)
	asm = append(asm, `LDB     R0 LTRCO:`)
	asm = append(asm, `TRP     #3`)
	asm = append(asm, `LDB     R0 NL:`)
	asm = append(asm, `TRP     #3`)
	asm = append(asm, `TRP     #0`)
	asm = append(asm, `UDRFLW: LDB     R0 LTRCU:`)
	asm = append(asm, `TRP     #3`)
	asm = append(asm, `LDB     R0 NL:`)
	asm = append(asm, `TRP     #3`)
	asm = append(asm, `TRP     #0`)

	//global data
	asm = append(asm, `;; global data`)
	asm = append(asm, `NL:     .BYT    '\n'`)
	asm = append(asm, `LTRCU:  .BYT    'U'`)
	asm = append(asm, `LTRCO:  .BYT    'O'`)
	asm = append(asm, `LTRCH:  .BYT    'H'`)
	for _, e := range st.GetAllOfKind("LitVar") {
		typ, _ := sym.StringFromData(e.Data, "type")
		switch typ {
		case "int":
			asm = append(asm, fmt.Sprintf("%s:\t.INT\t%s", e.SymId, e.Value))
		case "char":
			if e.Value == `' '` {
				asm = append(asm, fmt.Sprintf("%s:\t.BYT\t32", e.SymId))
			} else {
				asm = append(asm, fmt.Sprintf("%s:\t.BYT\t%s", e.SymId, e.Value))
			}
		}
	}

	// HERE IS THE MEAT OF EVERYTHING
	asm = append(asm, `;; functions`)

	for _, row := range table.GetRows() {
		//fmt.Printf("row: %#v\n", row)
		switch row.GetCommand() {
		case "ADD", "SUB", "MUL", "DIV":
			label := row.GetLabel()
			for i, r := range loadToRegister(st, row.GetOp2(), "R4") {
				beg := ""
				if label != "" && i == 0 {
					beg = label + ":"
				}
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s %s\t;%s", beg, r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s\t;%s", beg, r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("%s\t%s\t;%s", beg, r.GetCommand(), r.GetComment()))
				}

			}

			for _, r := range loadToRegister(st, row.GetOp3(), "R3") {
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s %s\t;%s", r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s\t;%s", r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("\t%s\t;%s", r.GetCommand(), r.GetComment()))
				}
			}

			//compare for greater then
			//asm = append(asm, "TRP #99")
			switch row.GetCommand() {
			case "ADD":
				asm = append(asm, fmt.Sprintf("\tADD\tR3 R4\t;%s", row.GetComment()))
			case "SUB":
				asm = append(asm, fmt.Sprintf("\tSUB\tR3 R4\t;%s", row.GetComment()))
			case "MUL":
				asm = append(asm, fmt.Sprintf("\tMUL\tR3 R4\t;%s", row.GetComment()))
			case "DIV":
				asm = append(asm, fmt.Sprintf("\tDIV\tR3 R4\t;%s", row.GetComment()))
			}
			//asm = append(asm, "TRP #99")

			//save it to where its supposed to go
			for _, r := range saveFromRegister(st, row.GetOp1(), "R3") {
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s %s\t;%s", r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s\t;%s", r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("\t%s\t;%s", r.GetCommand(), r.GetComment()))
				}
			}
		case "NEWI":
			//get size of obj
			elem, err := st.GetElement(row.GetOp1())
			if err != nil {
				panic(fmt.Sprintf("Could not get elem for symId %s", row.GetOp1()))
			}

			objSize, err := sym.IntFromData(elem.Data, "size") //if err then assume 0

			//check for overflow (RSL)
			asm = append(asm, `;; Test for heap overflow`)
			asm = append(asm, `MOV     R10 R9`)                          //copy free pointer to tmp
			asm = append(asm, fmt.Sprintf(`ADI     R10 #%d`, objSize*1)) //add size of obj
			asm = append(asm, `CMP     R10 RSL`)                         //comp with the stack limit
			asm = append(asm, `BGT     R10 HOVRFLW:`)                    //if it would put it over the stack limit then branch to overflow

			//FREE: (R9) reg value moved for storage (tmp)
			asm = append(asm, `MOV     R10 R9`) //copy free pointer to tmp
			//FREE: (R9) updated by size of Op1
			asm = append(asm, fmt.Sprintf(`ADI     R9 #%d`, objSize*1)) //add size of obj
			//Store tmp in Op2
			for _, r := range saveFromRegister(st, row.GetOp2(), "R9") {
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s %s\t;%s", r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s\t;%s", r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("\t%s\t;%s", r.GetCommand(), r.GetComment()))
				}
			}
		case "FUNC":
			asm = append(asm, fmt.Sprintf("%s:   ADI   R0 #0 ;%s", row.GetOp1(), row.GetComment()))
		case "FRAME":
			funcElem, err := st.GetElement(row.GetOp2())
			if err != nil {
				panic(fmt.Sprintf("Could not find elem for symbol %s", row.GetOp2()))
			}
			funcSize, _ := sym.IntFromData(funcElem.Data, "size")

			//Go to main
			asm = append(asm, fmt.Sprintf(`;; Call function "%s:    %s"`, row.GetOp2(), row.GetComment()))
			asm = append(asm, `;; Test for overflow`)
			asm = append(asm, fmt.Sprintf(`%s:   MOV     R10 RSP`, row.GetLabel()))
			asm = append(asm, fmt.Sprintf(`ADI     R10 #%d          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)`, -12+(funcSize*-1)))
			asm = append(asm, `CMP     R10 RSL`)
			asm = append(asm, `BLT     R10 OVRFLW:`)
			asm = append(asm, fmt.Sprintf(`;; Create Activation Record and invoke %s`, row.GetOp2()))
			asm = append(asm, `MOV     R10 RFP`)
			asm = append(asm, `MOV     RFP RSP`)
			asm = append(asm, `ADI     RSP #-4`)
			asm = append(asm, `STR     R10 (RSP)`)
			asm = append(asm, `ADI     RSP #-4`)

			//get this and then store it in the frame
			asm = append(asm, `;; this`)
			for _, r := range loadToRegister(st, row.GetOp1(), "R1") {
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s %s\t;%s", r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s\t;%s", r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("\t%s\t;%s", r.GetCommand(), r.GetComment()))
				}
			}
			asm = append(asm, `STR     R1 (RSP)`)
			asm = append(asm, `ADI     RSP #-4`)

			//parameters (PUSH)
		case "PUSH":
			asm = append(asm, fmt.Sprintf(`;; parameters on the stack (%s)  ; %s`, row.GetOp1(), row.GetComment()))
			//add the functions parameters to the stack
			e, err := st.GetElement(row.GetOp1())
			if err != nil {
				panic(fmt.Sprintf("Could not find elem for symbol %s", row.GetOp1()))
			}

			typ, _ := sym.StringFromData(e.Data, "type")
			isArray, _ := sym.BoolFromData(e.Data, "isArray")
			for _, r := range loadToRegister(st, e.SymId, "R1") {
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s %s\t;%s", r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s\t;%s", r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("\t%s\t;%s", r.GetCommand(), r.GetComment()))
				}
			}
			eSize := sym.SizeOfType(typ, isArray) * -1
			switch eSize {
			case -1:
				asm = append(asm, `STB     R1 (RSP)`)
			case -4:
				asm = append(asm, `STR     R1 (RSP)`)
			}
			asm = append(asm, fmt.Sprintf(`ADI     RSP #%d`, eSize))

		case "CALL":
			funSymId := row.GetOp1()
			funcElem, err := st.GetElement(funSymId)
			if err != nil {
				panic(fmt.Sprintf("Could not find elem for symbol %s", funSymId))
			}

			asm = append(asm, fmt.Sprintf(`;; local varibales on the stack    ; %s`, row.GetComment()))
			//add main's local variables
			for _, e := range getLocalVars(st, funcElem) {
				typ, _ := sym.StringFromData(e.Data, "type")
				isArray, _ := sym.BoolFromData(e.Data, "isArray")
				asm = append(asm, fmt.Sprintf(`ADI     RSP #%d`, (sym.SizeOfType(typ, isArray)*-1)))
			}

			asm = append(asm, `;; Temp variables on the stack`)
			//add main's temp variables
			for _, e := range getTempVars(st, funcElem) {
				typ, _ := sym.StringFromData(e.Data, "type")
				isArray, _ := sym.BoolFromData(e.Data, "isArray")
				asm = append(asm, fmt.Sprintf(`ADI     RSP #%d`, (sym.SizeOfType(typ, isArray)*-1)))
			}

			asm = append(asm, `;; set the return address and jump`)
			asm = append(asm, `MOV     R10 RPC         ; PC already at next instruction`)
			asm = append(asm, `ADI     R10 #12`)
			asm = append(asm, `STR     R10 (RFP)`)
			asm = append(asm, fmt.Sprintf(`JMP     %s:`, funSymId))

		case "MOV":
			//get op2 into a register
			label := row.GetLabel()
			for i, r := range loadToRegister(st, row.GetOp2(), "R3") {
				beg := ""
				comment := r.GetComment()
				if label != "" && i == 0 {
					beg = label + ":"
				}
				if c := row.GetComment(); c != "" && i == 0 {
					comment = c
				}
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s %s\t;%s", beg, r.GetCommand(), r.GetOp1(), r.GetOp2(), comment))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s\t;%s", beg, r.GetCommand(), r.GetOp1(), comment))
				default:
					asm = append(asm, fmt.Sprintf("%s\t%s\t;%s", beg, r.GetCommand(), comment))
				}

			}

			//store it into op1
			for _, r := range saveFromRegister(st, row.GetOp1(), "R3") {
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s %s\t;%s", r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s\t;%s", r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("\t%s\t;%s", r.GetCommand(), r.GetComment()))
				}
			}
		case "WRITE":
			elem, err := st.GetElement(row.GetOp1())
			if err != nil {
				panic(fmt.Sprintf("Could not find the symbol table element to write %s", row.GetOp1()))
			}
			//get the value to R0 for writing
			label := row.GetLabel()
			for i, r := range loadToRegister(st, row.GetOp1(), "R0") {
				beg := ""
				if label != "" && i == 0 {
					beg = label + ":"
				}
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s %s\t;%s", beg, r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s\t;%s", beg, r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("%s\t%s\t;%s", beg, r.GetCommand(), r.GetComment()))
				}

			}

			size, err := sym.IntFromData(elem.Data, "size")
			if err != nil {
				typ, err := sym.StringFromData(elem.Data, "type")
				if err != nil {
					panic(fmt.Sprintf("could not find size for symId: %s", elem.SymId))
				}

				isArr, _ := sym.BoolFromData(elem.Data, "isArray")

				size = sym.SizeOfType(typ, isArr)
			}
			switch size {
			case 1:
				asm = append(asm, fmt.Sprintf("\tTRP\t#3\t;%s", row.GetComment()))
			default:
				asm = append(asm, fmt.Sprintf("\tTRP\t#1\t;%s", row.GetComment()))
			}

		case "GT", "GTE":
			label := row.GetLabel()
			for i, r := range loadToRegister(st, row.GetOp2(), "R3") {
				beg := ""
				if label != "" && i == 0 {
					beg = label + ":"
				}
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s %s\t;%s", beg, r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s\t;%s", beg, r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("%s\t%s\t;%s", beg, r.GetCommand(), r.GetComment()))
				}

			}

			for _, r := range loadToRegister(st, row.GetOp3(), "R4") {
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s %s\t;%s", r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s\t;%s", r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("\t%s\t;%s", r.GetCommand(), r.GetComment()))
				}
			}

			//compare for greater then
			//asm = append(asm, "TRP #99")
			asm = append(asm, fmt.Sprintf("\tCMP\tR3 R4\t;%s", row.GetComment()))
			trueBranch := st.GenSymId("BTrue")
			falseBranch := st.GenSymId("BFalse")
			asm = append(asm, fmt.Sprintf("\tBGT\tR3 %s:\t", trueBranch))
			if row.GetCommand() == "GTE" {
				asm = append(asm, fmt.Sprintf("\tBRZ\tR3 %s:\t", trueBranch))
			}
			asm = append(asm, fmt.Sprintf("\tSUB\tR3 R3\t; false branch"))
			asm = append(asm, fmt.Sprintf("\tJMP\t%s:\t", falseBranch))
			asm = append(asm, fmt.Sprintf("%s:\tSUB\tR3 R3\t;True Branch", trueBranch))
			asm = append(asm, fmt.Sprintf("\tADI\tR3 #1\t;True Branch"))
			//asm = append(asm, "TRP #99")

			//save it to where its supposed to go
			for i, r := range saveFromRegister(st, row.GetOp1(), "R3") {
				label := ""
				if i == 0 {
					label = falseBranch + ":"
				}
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s %s\t;%s", label, r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s\t;%s", label, r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("%s\t%s\t;%s", label, r.GetCommand(), r.GetComment()))
				}
			}
		case "LT", "LTE":
			label := row.GetLabel()
			for i, r := range loadToRegister(st, row.GetOp2(), "R3") {
				beg := ""
				if label != "" && i == 0 {
					beg = label + ":"
				}
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s %s\t;%s", beg, r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s\t;%s", beg, r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("%s\t%s\t;%s", beg, r.GetCommand(), r.GetComment()))
				}

			}

			for _, r := range loadToRegister(st, row.GetOp3(), "R4") {
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s %s\t;%s", r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s\t;%s", r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("\t%s\t;%s", r.GetCommand(), r.GetComment()))
				}
			}

			//compare for less then
			//asm = append(asm, "TRP #99")
			asm = append(asm, fmt.Sprintf("\tCMP\tR3 R4\t;%s", row.GetComment()))
			trueBranch := st.GenSymId("BTrue")
			falseBranch := st.GenSymId("BFalse")
			asm = append(asm, fmt.Sprintf("\tBLT\tR3 %s:\t", trueBranch))
			if row.GetCommand() == "LTE" {
				asm = append(asm, fmt.Sprintf("\tBRZ\tR3 %s:\t", trueBranch))
			}
			asm = append(asm, fmt.Sprintf("\tSUB\tR3 R3\t; false branch"))
			asm = append(asm, fmt.Sprintf("\tJMP\t%s:\t", falseBranch))
			asm = append(asm, fmt.Sprintf("%s:\tSUB\tR3 R3\t;True Branch", trueBranch))
			asm = append(asm, fmt.Sprintf("\tADI\tR3 #1\t;True Branch"))
			//asm = append(asm, "TRP #99")

			//save it to where its supposed to go
			for i, r := range saveFromRegister(st, row.GetOp1(), "R3") {
				label := ""
				if i == 0 {
					label = falseBranch + ":"
				}
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s %s\t;%s", label, r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s\t;%s", label, r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("%s\t%s\t;%s", label, r.GetCommand(), r.GetComment()))
				}
			}
		case "EQ", "NEQ":
			label := row.GetLabel()
			for i, r := range loadToRegister(st, row.GetOp2(), "R3") {
				beg := ""
				if label != "" && i == 0 {
					beg = label + ":"
				}
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s %s\t;%s", beg, r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s\t;%s", beg, r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("%s\t%s\t;%s", beg, r.GetCommand(), r.GetComment()))
				}

			}

			for _, r := range loadToRegister(st, row.GetOp3(), "R4") {
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s %s\t;%s", r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s\t;%s", r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("\t%s\t;%s", r.GetCommand(), r.GetComment()))
				}
			}

			//compare for greater then
			//asm = append(asm, "TRP #99")
			asm = append(asm, fmt.Sprintf("\tCMP\tR3 R4\t;%s", row.GetComment()))
			trueBranch := st.GenSymId("BTrue")
			falseBranch := st.GenSymId("BFalse")
			//asm = append(asm, fmt.Sprintf("\tCBGT\tR3 %s:\t", trueBranch))
			if row.GetCommand() == "EQ" {
				asm = append(asm, fmt.Sprintf("\tBRZ\tR3 %s:\t", trueBranch))
			} else {
				asm = append(asm, fmt.Sprintf("\tBNZ\tR3 %s:\t", trueBranch))
			}
			asm = append(asm, fmt.Sprintf("\tSUB\tR3 R3\t; false branch"))
			asm = append(asm, fmt.Sprintf("\tJMP\t%s\t", falseBranch))
			asm = append(asm, fmt.Sprintf("%s:\tSUB\tR3 R3\t;True Branch", trueBranch))
			asm = append(asm, fmt.Sprintf("\tADI\tR3 #1\t;True Branch"))
			//asm = append(asm, "TRP #99")

			//save it to where its supposed to go
			for i, r := range saveFromRegister(st, row.GetOp1(), "R3") {
				label := ""
				if i == 0 {
					label = falseBranch + ":"
				}
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s %s\t;%s", label, r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s\t;%s", label, r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("%s\t%s\t;%s", label, r.GetCommand(), r.GetComment()))
				}
			}
		case "BF":
			//Get what we need
			label := row.GetLabel()
			for i, r := range loadToRegister(st, row.GetOp1(), "R3") {
				beg := ""
				if label != "" && i == 0 {
					beg = label + ":"
				}
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s %s\t;%s", beg, r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s\t;%s", beg, r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("%s\t%s\t;%s", beg, r.GetCommand(), r.GetComment()))
				}

			}

			//break if false
			//asm = append(asm, "TRP #99")
			asm = append(asm, fmt.Sprintf("\tBRZ\tR3 %s:\t;%s", row.GetOp2(), row.GetComment()))
			//asm = append(asm, "TRP #99")

		case "JMP":
			if label := row.GetLabel(); label != "" {
				asm = append(asm, fmt.Sprintf("%s:\tJMP\t%s:\t;%s", label, row.GetOp1(), row.GetComment()))
			} else {
				asm = append(asm, fmt.Sprintf("\tJMP\t%s:\t;%s", row.GetOp1(), row.GetComment()))
			}
		case "PEEK":
			//get size of obj
			elem, err := st.GetElement(row.GetOp1())
			if err != nil {
				panic(fmt.Sprintf("Could not get elem for symId %s", row.GetOp1()))
			}

			varSize, err := sym.IntFromData(elem.Data, "size") //if err then assume 0
			if err != nil {
				typ, err := sym.StringFromData(elem.Data, "type")
				if err != nil {
					panic(fmt.Sprintf("Could not get type of elem with symId %s", elem.SymId))
				}

				isArray, _ := sym.BoolFromData(elem.Data, "isArray")

				varSize = sym.SizeOfType(typ, isArray)
			}
			label := row.GetLabel()
			if label != "" {
				label += ":"
			}
			switch varSize {
			case 1:
				asm = append(asm, fmt.Sprintf("%s\tLDB\tR11 (RSP)\t;%s", label, row.GetComment()))
			default:
				asm = append(asm, fmt.Sprintf("%s\tLDR\tR11 (RSP)\t;%s", label, row.GetComment()))
			}

			//save it to the desired location
			for _, r := range saveFromRegister(st, row.GetOp1(), "R11") {
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s %s\t;%s", r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s\t;%s", r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("\t%s\t;%s", r.GetCommand(), r.GetComment()))
				}
			}
		case "RTN", "RETURN":
			asm = append(asm, `;; return from function`)
			asm = append(asm, `;; test for underflow`)

			if label := row.GetLabel(); label != "" {
				asm = append(asm, fmt.Sprintf("%s:\tMOV\tRSP RFP\t; %s", label, row.GetComment()))
			} else {
				asm = append(asm, `MOV     RSP RFP`)
			}
			asm = append(asm, `MOV     R10 RSP`)
			asm = append(asm, `CMP     R10 RSB`)
			asm = append(asm, `BGT     R10 UDRFLW:     ; oopsy underflow problem`)
			asm = append(asm, `;; set previous frame to current frame and return`)
			asm = append(asm, `LDR     R10 (RFP)`)
			asm = append(asm, `MOV     R11 RFP`)
			asm = append(asm, `ADI     R11 #-4         ; now pointing at PFP`)
			asm = append(asm, `LDR     RFP (R11)       ; make FP = PFP`)
			if row.GetCommand() == "RETURN" {
				asm = append(asm, `;; store the return value`)
				for _, r := range loadToRegister(st, row.GetOp1(), "R0") {
					switch {
					case r.GetOp2() != "":
						asm = append(asm, fmt.Sprintf("\t%s\t%s %s\t;%s", r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
					case r.GetOp1() != "":
						asm = append(asm, fmt.Sprintf("\t%s\t%s\t;%s", r.GetCommand(), r.GetOp1(), r.GetComment()))
					default:
						asm = append(asm, fmt.Sprintf("\t%s\t;%s", r.GetCommand(), r.GetComment()))
					}
				}
				//get size of obj
				elem, err := st.GetElement(row.GetOp1())
				if err != nil {
					panic(fmt.Sprintf("Could not get elem for symId %s", row.GetOp1()))
				}

				varSize, err := sym.IntFromData(elem.Data, "size") //if err then assume 0
				if err != nil {
					typ, err := sym.StringFromData(elem.Data, "type")
					if err != nil {
						panic(fmt.Sprintf("Could not get type of elem with symId %s", elem.SymId))
					}

					isArray, _ := sym.BoolFromData(elem.Data, "isArray")

					varSize = sym.SizeOfType(typ, isArray)
				}
				switch varSize {
				case 1:
					asm = append(asm, `STB     R0 (RSP)        ; R0 is whatever the value is for return`)
				default:
					asm = append(asm, `STR     R0 (RSP)        ; R0 is whatever the value is for return`)
				}
			}
			asm = append(asm, fmt.Sprintf(`JMR     R10             ; go back "%s"`, row.GetComment()))
			asm = append(asm, "\n")
		default:
			panic(fmt.Sprintf("Dont have translation for %#v\n", row))
		}
	}

	//the heap starts here
	asm = append(asm, `;; Heap starts here`)
	asm = append(asm, `FREE:    .INT 0`)

	return
}

func getLocation(st *sym.SymbolTable, symId string) (loc string, offset int, size int) {
	if symId == "this" {
		return "stack", -8, 4
	}
	elem, err := st.GetElement(symId)
	if err != nil {
		panic("trying to access non element in symbol table")
	}

	size, err = sym.IntFromData(elem.Data, "size")
	if err != nil {
		typ, err := sym.StringFromData(elem.Data, "type")
		if err != nil {
			panic(fmt.Sprintf("no type for element %s", elem.SymId))
		}

		isArray, err := sym.BoolFromData(elem.Data, "isArray")
		//if err != nil { panic(fmt.Sprintf("no isArray for element %s",elem.SymId)) }

		size = sym.SizeOfType(typ, isArray)
	}

	switch elem.Kind {
	case "LitVar":
		loc = "memory"
	default:
		loc = "stack"
		offset, err = sym.IntFromData(elem.Data, "offset")
		if err != nil {
			panic(fmt.Sprintf("no offset for symId", symId))
		}
		offset += 12 //for pfp,ret,this
		offset *= -1
	}
	return
}

func loadToRegister(st *sym.SymbolTable, symId, reg string) (rows []*ic.QuadRow) {
	loc, offset, size := getLocation(st, symId)
	rows = make([]*ic.QuadRow, 0)
	switch loc {
	case "memory":
		switch size {
		case 1:
			rows = append(rows, ic.NewQuadRow("", "LDB", reg, symId+":", "", ""))
		default:
			rows = append(rows, ic.NewQuadRow("", "LDR", reg, symId+":", "", ""))
		}
	case "stack":
		//rows = append(rows, ic.NewQuadRow("","TRP", "#99", "", "", ""))
		rows = append(rows, ic.NewQuadRow("", "MOV", "R10", "RFP", "", ""))
		rows = append(rows, ic.NewQuadRow("", "ADI", "R10", "#"+strconv.Itoa(offset), "", ""))
		switch size {
		case 1:
			rows = append(rows, ic.NewQuadRow("", "LDB", reg, "(R10)", "", ""))
		default:
			rows = append(rows, ic.NewQuadRow("", "LDR", reg, "(R10)", "", ""))
		}
		//rows = append(rows, ic.NewQuadRow("","TRP", "#99", "", "", ""))
	default:
		panic(fmt.Sprintf("loading from location %s unknown", loc))
	}
	return
}

func saveFromRegister(st *sym.SymbolTable, symId, reg string) (rows []*ic.QuadRow) {
	loc, offset, size := getLocation(st, symId)
	rows = make([]*ic.QuadRow, 0)
	switch loc {
	case "stack":
		//rows = append(rows, ic.NewQuadRow("","TRP", "#99", "", "", ""))
		rows = append(rows, ic.NewQuadRow("", "MOV", "R10", "RFP", "", ""))
		rows = append(rows, ic.NewQuadRow("", "ADI", "R10", "#"+strconv.Itoa(offset), "", ""))
		switch size {
		case 1:
			rows = append(rows, ic.NewQuadRow("", "STB", reg, "(R10)", "", ""))
		default:
			rows = append(rows, ic.NewQuadRow("", "STR", reg, "(R10)", "", ""))
		}
		//rows = append(rows, ic.NewQuadRow("","TRP", "#99", "", "", ""))
	default:
		panic(fmt.Sprintf("saving to location %s unknown", loc))
	}
	return
}

func getParamVars(st *sym.SymbolTable, fun sym.SymbolTableElement) (elems []sym.SymbolTableElement) {
	elems = make([]sym.SymbolTableElement, 0)
	es := st.GetScopeElements(fun.Scope + "." + fun.Value)
	for _, e := range es {
		switch e.Kind {
		case "Parameter":
			elems = append(elems, e)
		}
	}
	return
}

func getLocalVars(st *sym.SymbolTable, fun sym.SymbolTableElement) (elems []sym.SymbolTableElement) {
	elems = make([]sym.SymbolTableElement, 0)
	es := st.GetScopeElements(fun.Scope + "." + fun.Value)
	for _, e := range es {
		switch e.Kind {
		case "Lvar":
			elems = append(elems, e)
		}
	}
	return
}

func getTempVars(st *sym.SymbolTable, fun sym.SymbolTableElement) (elems []sym.SymbolTableElement) {
	elems = make([]sym.SymbolTableElement, 0)
	es := st.GetScopeElements(fun.Scope + "." + fun.Value)
	for _, e := range es {
		switch e.Kind {
		case "Tvar":
			elems = append(elems, e)
		}
	}
	return
}
