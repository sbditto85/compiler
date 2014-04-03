package tcode

import (
	"fmt"
	"strconv"
	ic "github.com/sbditto85/compiler/analyzer/icode"
	sym "github.com/sbditto85/compiler/symbol_table"
)

func GenerateASM(table *ic.Quad, st *sym.SymbolTable) (asm []string) {
	asm = make([]string, 0, table.Size())

	//table.Print()
	//st.PrintTable()

	mainElem := st.GetElementInScope("g","main")
	mainSize, _ := sym.IntFromData(mainElem.Data,"size")

	//debug?
	//asm = append(asm, `TRP     #99`)

	//setup heap
	asm = append(asm, `LDA     R9 FREE:`)

	//Go to main
	asm = append(asm, `;; Call function "MAIN:"`)
	asm = append(asm, `;; Test for overflow`)
	asm = append(asm, `MOV     R10 RSP`)
	asm = append(asm, fmt.Sprintf(`ADI     R10 #%d          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)`, -12 + (mainSize * -1)))
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
	for _,e := range(getLocalVars(st,mainElem)) {
		typ, _ := sym.StringFromData(e.Data,"type")
		isArray, _ := sym.BoolFromData(e.Data,"isArray")
		asm = append(asm, fmt.Sprintf(`ADI     RSP #%d`,(sym.SizeOfType(typ,isArray) * -1)))
	}

	asm = append(asm, `;; Temp variables on the stack`)
	//add main's temp variables
	for _,e := range(getTempVars(st,mainElem)) {
		typ, _ := sym.StringFromData(e.Data,"type")
		isArray, _ := sym.BoolFromData(e.Data,"isArray")
		asm = append(asm, fmt.Sprintf(`ADI     RSP #%d`,(sym.SizeOfType(typ,isArray) * -1)))
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
	for _,e := range(st.GetAllOfKind("LitVar")) {
		typ, _ := sym.StringFromData(e.Data,"type")
		switch typ {
		case "int":
			asm = append(asm, fmt.Sprintf("%s:\t.INT\t%s",e.SymId,e.Value))
		case "char":
			asm = append(asm, fmt.Sprintf("%s:\t.BYT\t%s",e.SymId,e.Value))
		}
	}

	// HERE IS THE MEAT OF EVERYTHING
	asm = append(asm, `;; functions`)
	
	for _, row := range(table.GetRows()) {
		fmt.Printf("row: %#v\n",row)
		switch row.GetCommand() {
		case "FUNC":
			asm = append(asm, fmt.Sprintf("%s:   ADI   R0 #0 ;%s", row.GetOp1(), row.GetComment()))
		case "MOV":
		case "WRITE":
		case "GT":
			label := row.GetLabel()
			for i, r := range(loadToRegister(st, &asm, row.GetOp2(), "R3")) {
				beg := ""
				if label != "" && i == 0 {
					beg = label + ":"
				}
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s %s\t;%s",beg,r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s\t;%s",beg,r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("%s\t%s\t;%s",beg,r.GetCommand(), r.GetComment()))
				}
		
			}
			
			for _, r := range(loadToRegister(st, &asm, row.GetOp3(), "R4")) {
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s %s\t;%s",r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("\t%s\t%s\t;%s",r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("\t%s\t;%s",r.GetCommand(), r.GetComment()))
				}
			}
		case "BF":
			label := row.GetLabel()
			for i, r := range(loadToRegister(st, &asm, row.GetOp1(), "R3")) {
				beg := ""
				if label != "" && i == 0 {
					beg = label + ":"
				}
				switch {
				case r.GetOp2() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s %s\t;%s",beg,r.GetCommand(), r.GetOp1(), r.GetOp2(), r.GetComment()))
				case r.GetOp1() != "":
					asm = append(asm, fmt.Sprintf("%s\t%s\t%s\t;%s",beg,r.GetCommand(), r.GetOp1(), r.GetComment()))
				default:
					asm = append(asm, fmt.Sprintf("%s\t%s\t;%s",beg,r.GetCommand(), r.GetComment()))
				}
				
			}
		case "JMP":
			if label := row.GetLabel(); label != "" {
				asm = append(asm, fmt.Sprintf("%s:\tJMP\t%s:\t;%s",label, row.GetOp1(),row.GetComment()))
			} else {
				asm = append(asm, fmt.Sprintf("\tJMP\t%s:\t;%s", row.GetOp1(), row.GetComment()))
			}
		case "RTN":
			asm = append(asm, `;; return from function`)
			asm = append(asm, `;; test for underflow`)

			if label := row.GetLabel(); label != "" {
				asm = append(asm, fmt.Sprintf("%s:\tMOV\tRSP RFP\t; %s",label, row.GetComment()))
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
			asm = append(asm, `;; store the return value`)
			//asm = append(asm, `STR     R0 (RSP)        ; R0 is wherever the value is for return`)
			asm = append(asm, fmt.Sprintf(`JMR     R10             ; go back "%s"`,row.GetComment()))
		}
	}

	//the heap starts here
	asm = append(asm, `;; Heap starts here`)
	asm = append(asm, `FREE:    .INT 0`)

	return
}

func getLocation(st *sym.SymbolTable, symId string) (loc string, offset int) {
	if symId == "this" {
		return "stack", -8
	}
	elem, err := st.GetElement(symId)
	if err != nil {
		panic("trying to access non element in symbol table")
	}

	switch elem.Kind {
	case "LitVar":
		loc = "memory"
	default:
		loc = "stack"
		offset, err = sym.IntFromData(elem.Data,"offset")
		if err != nil {
			panic(fmt.Sprintf("no offset for symId",symId))
		}
		offset += 12 //for pfp,ret,this
		offset *= -1
	}
	return
}

func loadToRegister(st *sym.SymbolTable, asm *[]string, symId, reg string) (rows []*ic.QuadRow) {
	loc, offset := getLocation(st, symId)
	rows = make([]*ic.QuadRow,0)
	switch loc {
	case "memory":
		rows = append(rows, ic.NewQuadRow("","LDR", reg, symId+":", "", ""))
	case "stack":
		//rows = append(rows, ic.NewQuadRow("","TRP", "#99", "", "", ""))
		rows = append(rows, ic.NewQuadRow("","MOV", "R10", "RFP", "", ""))
		rows = append(rows, ic.NewQuadRow("","ADI", "R10", "#"+strconv.Itoa(offset), "", ""))
		rows = append(rows, ic.NewQuadRow("","LDR", reg, "(R10)", "", ""))
		//rows = append(rows, ic.NewQuadRow("","TRP", "#99", "", "", ""))
	}
	return
}

func getLocalVars(st *sym.SymbolTable, fun sym.SymbolTableElement) (elems []sym.SymbolTableElement) {
	elems = make([]sym.SymbolTableElement,0)
	es := st.GetScopeElements(fun.Scope + "." + fun.Value)
	for _, e := range(es) {
		switch e.Kind {
		case "Lvar":
			elems = append(elems,e)
		}
	}
	return
}

func getTempVars(st *sym.SymbolTable, fun sym.SymbolTableElement) (elems []sym.SymbolTableElement) {
	elems = make([]sym.SymbolTableElement,0)
	es := st.GetScopeElements(fun.Scope + "." + fun.Value)
	for _, e := range(es) {
		switch e.Kind {
		case "Tvar":
			elems = append(elems,e)
		}
	}
	return
}
