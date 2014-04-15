LDA     R9 FREE:
;; Call function "MAIN:"
;; Test for overflow
MOV     R10 RSP
ADI     R10 #-63          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke MAIN
MOV     R10 RFP
MOV     RFP RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
SUB     R1 R1           ; get this from where its at
STR     R1 (RSP)
ADI     RSP #-4
;; parameters on the stack
;; local varibales on the stack
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     MAIN:
EXIT:   TRP     #0
OVRFLW: LDB     R0 LTRCO:
TRP     #3
LDB     R0 NL:
TRP     #3
TRP     #0
HOVRFLW: LDB     R0 LTRCH:
TRP     #3
LDB     R0 LTRCO:
TRP     #3
LDB     R0 NL:
TRP     #3
TRP     #0
UDRFLW: LDB     R0 LTRCU:
TRP     #3
LDB     R0 NL:
TRP     #3
TRP     #0
;; global data
NL:     .BYT    '\n'
LTRCU:  .BYT    'U'
LTRCO:  .BYT    'O'
LTRCH:  .BYT    'H'
Li153:	.INT	4
Li163:	.INT	7
Li245:	.INT	27
Li174:	.INT	10
Li123:	.BYT	','
Li190:	.INT	14
Li156:	.INT	5
Li242:	.INT	26
Li251:	.BYT	'r'
Li146:	.BYT	'd'
Li236:	.INT	24
Li110:	.BYT	'\n'
Li170:	.INT	9
Li55:	.INT	0
Li105:	.BYT	0
Li143:	.BYT	'A'
Li162:	.BYT	'E'
Li167:	.INT	8
Li223:	.BYT	'a'
Li65:	.INT	2
Li177:	.INT	11
Li204:	.BYT	'p'
Li248:	.INT	28
Li59:	.INT	1
Li218:	.BYT	'c'
Li286:	.BYT	'g'
Li77:	.BYT	1
Li159:	.INT	6
Li194:	.BYT	'D'
Li152:	.BYT	'e'
Li277:	.INT	42
Li53:	.INT	0
Li181:	.INT	12
Li285:	.INT	37
Li121:	.BYT	32
Li185:	.INT	13
Li173:	.BYT	'm'
Li213:	.BYT	'i'
Li184:	.BYT	't'
Li199:	.BYT	'u'
Li180:	.BYT	'n'
Li188:	.BYT	':'
Li138:	.INT	100
Li149:	.INT	3
Li166:	.BYT	'l'
Li239:	.INT	25
;; functions
;; row: :	FUNC	Co4  ;     iTree() {
Co4:   ADI   R0 #0 ;    iTree() {
;; row: :	FRAME	this St47 ;     iTree() {
;; Call function "St47:        iTree() {"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-17          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke St47
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	St47  ;     iTree() {
;; local varibales on the stack    ;     iTree() {
;; Temp variables on the stack
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-17
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     St47:
;; row: :	REF	Tv52 Iv2 this; 	root = null;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv52 Li53 ; 	root = null;
	LDR	R3 Li53:	;	root = null;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	RETURN	this  ;     }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R0 (R10)	;
STR     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "    }"


;; row: :	FUNC	Me6  ;     private int fib(int root) {
Me6:   ADI   R0 #0 ;    private int fib(int root) {
;; row: :	EQ	Tv56 Pa5 Li55; 	if (root == 0) return 0;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	LDR	R4 Li55:	;
	CMP	R3 R4	;	if (root == 0) return 0;
	BRZ	R3 BT301:	
	SUB	R3 R3	; false branch
	JMP	BF302:	
BT301:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF302:	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STB	R3 (R10)	;
;; row: :	BF	Tv56 If57 ; 	if (root == 0) return 0;
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDB	R3 (R10)	;
	BRZ	R3 If57:	;	if (root == 0) return 0;
;; row: :	RETURN	Li55  ; 	if (root == 0) return 0;
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	LDR	R0 Li55:	;
STR     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "	if (root == 0) return 0;"


;; row: :	JMP	El62  ; 	else if (root == 1) return 1;
	JMP	El62:	;	else if (root == 1) return 1;
;; row: If57:	EQ	Tv60 Pa5 Li59; 	else if (root == 1) return 1;
If57:	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	LDR	R4 Li59:	;
	CMP	R3 R4	;	else if (root == 1) return 1;
	BRZ	R3 BT303:	
	SUB	R3 R3	; false branch
	JMP	BF304:	
BT303:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF304:	MOV	R10 RFP	;
	ADI	R10 #-17	;
	STB	R3 (R10)	;
;; row: :	BF	Tv60 If61 ; 	else if (root == 1) return 1;
	MOV	R10 RFP	;
	ADI	R10 #-17	;
	LDB	R3 (R10)	;
	BRZ	R3 If61:	;	else if (root == 1) return 1;
;; row: :	RETURN	Li59  ; 	else if (root == 1) return 1;
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	LDR	R0 Li59:	;
STR     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "	else if (root == 1) return 1;"


;; row: :	JMP	El62  ; 	else return (fib(root - 1) + fib(root - 2));
	JMP	El62:	;	else return (fib(root - 1) + fib(root - 2));
;; row: If61:	SUB	Tv63 Li59 Pa5; 	else return (fib(root - 1) + fib(root - 2));
If61:	LDR	R4 Li59:	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	SUB	R3 R4	;	else return (fib(root - 1) + fib(root - 2));
	MOV	R10 RFP	;
	ADI	R10 #-18	;
	STR	R3 (R10)	;
;; row: :	FRAME	this Me6 ; 	else return (fib(root - 1) + fib(root - 2));
;; Call function "Me6:    	else return (fib(root - 1) + fib(root - 2));"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-38          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me6
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv63  ; 	else return (fib(root - 1) + fib(root - 2));
;; parameters on the stack (Tv63)  ; 	else return (fib(root - 1) + fib(root - 2));
	MOV	R10 RFP	;
	ADI	R10 #-18	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me6  ; 	else return (fib(root - 1) + fib(root - 2));
;; local varibales on the stack    ; 	else return (fib(root - 1) + fib(root - 2));
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-38
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me6:
;; row: :	PEEK	Tv64  ; 	else return (fib(root - 1) + fib(root - 2));
	LDR	R11 (RSP)	;	else return (fib(root - 1) + fib(root - 2));
	MOV	R10 RFP	;
	ADI	R10 #-22	;
	STR	R11 (R10)	;
;; row: :	SUB	Tv66 Li65 Pa5; 	else return (fib(root - 1) + fib(root - 2));
	LDR	R4 Li65:	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	SUB	R3 R4	;	else return (fib(root - 1) + fib(root - 2));
	MOV	R10 RFP	;
	ADI	R10 #-26	;
	STR	R3 (R10)	;
;; row: :	FRAME	this Me6 ; 	else return (fib(root - 1) + fib(root - 2));
;; Call function "Me6:    	else return (fib(root - 1) + fib(root - 2));"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-38          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me6
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv66  ; 	else return (fib(root - 1) + fib(root - 2));
;; parameters on the stack (Tv66)  ; 	else return (fib(root - 1) + fib(root - 2));
	MOV	R10 RFP	;
	ADI	R10 #-26	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me6  ; 	else return (fib(root - 1) + fib(root - 2));
;; local varibales on the stack    ; 	else return (fib(root - 1) + fib(root - 2));
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-38
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me6:
;; row: :	PEEK	Tv67  ; 	else return (fib(root - 1) + fib(root - 2));
	LDR	R11 (RSP)	;	else return (fib(root - 1) + fib(root - 2));
	MOV	R10 RFP	;
	ADI	R10 #-30	;
	STR	R11 (R10)	;
;; row: :	ADD	Tv68 Tv67 Tv64; 	else return (fib(root - 1) + fib(root - 2));
	MOV	R10 RFP	;
	ADI	R10 #-30	;
	LDR	R4 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-22	;
	LDR	R3 (R10)	;
	ADD	R3 R4	;	else return (fib(root - 1) + fib(root - 2));
	MOV	R10 RFP	;
	ADI	R10 #-34	;
	STR	R3 (R10)	;
;; row: :	RETURN	Tv68  ; 	else return (fib(root - 1) + fib(root - 2));
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	MOV	R10 RFP	;
	ADI	R10 #-34	;
	LDR	R0 (R10)	;
STR     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "	else return (fib(root - 1) + fib(root - 2));"


;; row: El62:	RTN	  ;     }
;; return from function
;; test for underflow
El62:	MOV	RSP RFP	;     }
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "    }"


;; row: :	FUNC	Me8  ;     public bool add(int key) {
Me8:   ADI   R0 #0 ;    public bool add(int key) {
;; row: :	DIV	Tv69 Li65 Pa7;     key = key + fib(key/2);
	LDR	R4 Li65:	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	DIV	R3 R4	;    key = key + fib(key/2);
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STR	R3 (R10)	;
;; row: :	FRAME	this Me6 ;     key = key + fib(key/2);
;; Call function "Me6:        key = key + fib(key/2);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-38          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me6
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv69  ;     key = key + fib(key/2);
;; parameters on the stack (Tv69)  ;     key = key + fib(key/2);
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me6  ;     key = key + fib(key/2);
;; local varibales on the stack    ;     key = key + fib(key/2);
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-38
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me6:
;; row: :	PEEK	Tv70  ;     key = key + fib(key/2);
	LDR	R11 (RSP)	;    key = key + fib(key/2);
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	STR	R11 (R10)	;
;; row: :	ADD	Tv71 Tv70 Pa7;     key = key + fib(key/2);
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	LDR	R4 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	ADD	R3 R4	;    key = key + fib(key/2);
	MOV	R10 RFP	;
	ADI	R10 #-24	;
	STR	R3 (R10)	;
;; row: :	MOV	Pa7 Tv71 ;     key = key + fib(key/2);
	MOV	R10 RFP	;    key = key + fib(key/2);
	ADI	R10 #-24	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R3 (R10)	;
;; row: :	REF	Tv72 Iv2 this; 	if (root == null) {
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-28	;
	STR	R13 (R10)	;
;; row: :	EQ	Tv73 Tv72 Li53; 	if (root == null) {
	MOV	R10 RFP	;Load Address
	ADI	R10 #-28	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	LDR	R4 Li53:	;
	CMP	R3 R4	;	if (root == null) {
	BRZ	R3 BT305:	
	SUB	R3 R3	; false branch
	JMP	BF306:	
BT305:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF306:	MOV	R10 RFP	;
	ADI	R10 #-32	;
	STB	R3 (R10)	;
;; row: :	BF	Tv73 If74 ; 	if (root == null) {
	MOV	R10 RFP	;
	ADI	R10 #-32	;
	LDB	R3 (R10)	;
	BRZ	R3 If74:	;	if (root == null) {
;; row: :	REF	Tv75 Iv2 this; 	    root = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-33	;
	STR	R13 (R10)	;
;; row: :	NEWI	Cl17 Tv76 ; 	    root = new iNode(key);
;; Test for heap overflow
	MOV     R10 R9
ADI     R10 #12
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADI     R9 #12
	MOV	R10 RFP	;
	ADI	R10 #-37	;
	STR	R11 (R10)	;
;; row: :	FRAME	Tv76 Co22 ; 	    root = new iNode(key);
;; Call function "Co22:    	    root = new iNode(key);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-28          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Co22
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-37	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Pa7  ; 	    root = new iNode(key);
;; parameters on the stack (Pa7)  ; 	    root = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Co22  ; 	    root = new iNode(key);
;; local varibales on the stack    ; 	    root = new iNode(key);
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-28
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Co22:
;; row: :	PEEK	Tv76  ; 	    root = new iNode(key);
	LDR	R11 (RSP)	;	    root = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-37	;
	STR	R11 (R10)	;
;; row: :	MOV	Tv75 Tv76 ; 	    root = new iNode(key);
	MOV	R10 RFP	;	    root = new iNode(key);
	ADI	R10 #-37	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-33	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	RETURN	Li77  ; 	    return true;
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	LDB	R0 Li77:	;
STB     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "	    return true;"


;; row: :	JMP	El78  ; 	else
	JMP	El78:	;	else
;; row: If74:	REF	Tv79 Iv2 this; 	    return insert(key, root);
If74:	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-41	;
	STR	R13 (R10)	;
;; row: :	FRAME	this Me11 ; 	    return insert(key, root);
;; Call function "Me11:    	    return insert(key, root);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-66          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me11
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Pa7  ; 	    return insert(key, root);
;; parameters on the stack (Pa7)  ; 	    return insert(key, root);
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv79  ; 	    return insert(key, root);
;; parameters on the stack (Tv79)  ; 	    return insert(key, root);
	MOV	R10 RFP	;Load Address
	ADI	R10 #-41	;
	LDR	R13 (R10)	;
	LDR	R1 (R13)	;Load to register
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me11  ; 	    return insert(key, root);
;; local varibales on the stack    ; 	    return insert(key, root);
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-66
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me11:
;; row: :	PEEK	Tv80  ; 	    return insert(key, root);
	LDB	R11 (RSP)	;	    return insert(key, root);
	MOV	R10 RFP	;
	ADI	R10 #-45	;
	STB	R11 (R10)	;
;; row: :	RETURN	Tv80  ; 	    return insert(key, root);
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	MOV	R10 RFP	;
	ADI	R10 #-45	;
	LDB	R0 (R10)	;
STB     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "	    return insert(key, root);"


;; row: El78:	RTN	  ;     }
;; return from function
;; test for underflow
El78:	MOV	RSP RFP	;     }
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "    }"


;; row: :	FUNC	Me11  ;     private bool insert(int key, iNode node) {
Me11:   ADI   R0 #0 ;    private bool insert(int key, iNode node) {
;; row: :	REF	Tv81 Iv18 Pa10; 	if (key < node.root)
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-20	;
	STR	R13 (R10)	;
;; row: :	LT	Tv82 Pa9 Tv81; 	if (key < node.root)
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-20	;
	LDR	R13 (R10)	;
	LDR	R4 (R13)	;Load to register
	CMP	R3 R4	;	if (key < node.root)
	BLT	R3 BT307:	
	SUB	R3 R3	; false branch
	JMP	BF308:	
BT307:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF308:	MOV	R10 RFP	;
	ADI	R10 #-24	;
	STB	R3 (R10)	;
;; row: :	BF	Tv82 If83 ; 	if (key < node.root)
	MOV	R10 RFP	;
	ADI	R10 #-24	;
	LDB	R3 (R10)	;
	BRZ	R3 If83:	;	if (key < node.root)
;; row: :	REF	Tv84 Iv19 Pa10; 	    if (node.left == null) {
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-25	;
	STR	R13 (R10)	;
;; row: :	EQ	Tv85 Tv84 Li53; 	    if (node.left == null) {
	MOV	R10 RFP	;Load Address
	ADI	R10 #-25	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	LDR	R4 Li53:	;
	CMP	R3 R4	;	    if (node.left == null) {
	BRZ	R3 BT309:	
	SUB	R3 R3	; false branch
	JMP	BF310:	
BT309:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF310:	MOV	R10 RFP	;
	ADI	R10 #-29	;
	STB	R3 (R10)	;
;; row: :	BF	Tv85 If86 ; 	    if (node.left == null) {
	MOV	R10 RFP	;
	ADI	R10 #-29	;
	LDB	R3 (R10)	;
	BRZ	R3 If86:	;	    if (node.left == null) {
;; row: :	REF	Tv87 Iv19 Pa10; 		node.left = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-30	;
	STR	R13 (R10)	;
;; row: :	NEWI	Cl17 Tv88 ; 		node.left = new iNode(key);
;; Test for heap overflow
	MOV     R10 R9
ADI     R10 #12
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADI     R9 #12
	MOV	R10 RFP	;
	ADI	R10 #-34	;
	STR	R11 (R10)	;
;; row: :	FRAME	Tv88 Co22 ; 		node.left = new iNode(key);
;; Call function "Co22:    		node.left = new iNode(key);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-28          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Co22
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-34	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Pa9  ; 		node.left = new iNode(key);
;; parameters on the stack (Pa9)  ; 		node.left = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Co22  ; 		node.left = new iNode(key);
;; local varibales on the stack    ; 		node.left = new iNode(key);
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-28
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Co22:
;; row: :	PEEK	Tv88  ; 		node.left = new iNode(key);
	LDR	R11 (RSP)	;		node.left = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-34	;
	STR	R11 (R10)	;
;; row: :	MOV	Tv87 Tv88 ; 		node.left = new iNode(key);
	MOV	R10 RFP	;		node.left = new iNode(key);
	ADI	R10 #-34	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-30	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	RETURN	Li77  ; 		return true;
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	LDB	R0 Li77:	;
STB     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "		return true;"


;; row: :	JMP	El89  ; 	    else 
	JMP	El89:	;	    else 
;; row: If86:	REF	Tv90 Iv19 Pa10; 		return insert(key, node.left);
If86:	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-38	;
	STR	R13 (R10)	;
;; row: :	FRAME	this Me11 ; 		return insert(key, node.left);
;; Call function "Me11:    		return insert(key, node.left);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-66          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me11
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Pa9  ; 		return insert(key, node.left);
;; parameters on the stack (Pa9)  ; 		return insert(key, node.left);
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv90  ; 		return insert(key, node.left);
;; parameters on the stack (Tv90)  ; 		return insert(key, node.left);
	MOV	R10 RFP	;Load Address
	ADI	R10 #-38	;
	LDR	R13 (R10)	;
	LDR	R1 (R13)	;Load to register
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me11  ; 		return insert(key, node.left);
;; local varibales on the stack    ; 		return insert(key, node.left);
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-66
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me11:
;; row: :	PEEK	Tv91  ; 		return insert(key, node.left);
	LDB	R11 (RSP)	;		return insert(key, node.left);
	MOV	R10 RFP	;
	ADI	R10 #-42	;
	STB	R11 (R10)	;
;; row: :	RETURN	Tv91  ; 		return insert(key, node.left);
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	MOV	R10 RFP	;
	ADI	R10 #-42	;
	LDB	R0 (R10)	;
STB     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "		return insert(key, node.left);"


;; row: El89:	JMP	El104  ; 	else if (key > node.root)
El89:	JMP	El104:	;	else if (key > node.root)
;; row: If83:	REF	Tv93 Iv18 Pa10; 	else if (key > node.root)
If83:	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-43	;
	STR	R13 (R10)	;
;; row: :	GT	Tv94 Pa9 Tv93; 	else if (key > node.root)
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-43	;
	LDR	R13 (R10)	;
	LDR	R4 (R13)	;Load to register
	CMP	R3 R4	;	else if (key > node.root)
	BGT	R3 BT311:	
	SUB	R3 R3	; false branch
	JMP	BF312:	
BT311:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF312:	MOV	R10 RFP	;
	ADI	R10 #-47	;
	STB	R3 (R10)	;
;; row: :	BF	Tv94 If95 ; 	else if (key > node.root)
	MOV	R10 RFP	;
	ADI	R10 #-47	;
	LDB	R3 (R10)	;
	BRZ	R3 If95:	;	else if (key > node.root)
;; row: :	REF	Tv96 Iv20 Pa10; 	    if (node.right == null) {
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-48	;
	STR	R13 (R10)	;
;; row: :	EQ	Tv97 Tv96 Li53; 	    if (node.right == null) {
	MOV	R10 RFP	;Load Address
	ADI	R10 #-48	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	LDR	R4 Li53:	;
	CMP	R3 R4	;	    if (node.right == null) {
	BRZ	R3 BT313:	
	SUB	R3 R3	; false branch
	JMP	BF314:	
BT313:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF314:	MOV	R10 RFP	;
	ADI	R10 #-52	;
	STB	R3 (R10)	;
;; row: :	BF	Tv97 If98 ; 	    if (node.right == null) {
	MOV	R10 RFP	;
	ADI	R10 #-52	;
	LDB	R3 (R10)	;
	BRZ	R3 If98:	;	    if (node.right == null) {
;; row: :	REF	Tv99 Iv20 Pa10; 		node.right = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-53	;
	STR	R13 (R10)	;
;; row: :	NEWI	Cl17 Tv100 ; 		node.right = new iNode(key);
;; Test for heap overflow
	MOV     R10 R9
ADI     R10 #12
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADI     R9 #12
	MOV	R10 RFP	;
	ADI	R10 #-57	;
	STR	R11 (R10)	;
;; row: :	FRAME	Tv100 Co22 ; 		node.right = new iNode(key);
;; Call function "Co22:    		node.right = new iNode(key);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-28          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Co22
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-57	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Pa9  ; 		node.right = new iNode(key);
;; parameters on the stack (Pa9)  ; 		node.right = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Co22  ; 		node.right = new iNode(key);
;; local varibales on the stack    ; 		node.right = new iNode(key);
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-28
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Co22:
;; row: :	PEEK	Tv100  ; 		node.right = new iNode(key);
	LDR	R11 (RSP)	;		node.right = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-57	;
	STR	R11 (R10)	;
;; row: :	MOV	Tv99 Tv100 ; 		node.right = new iNode(key);
	MOV	R10 RFP	;		node.right = new iNode(key);
	ADI	R10 #-57	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-53	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	RETURN	Li77  ; 		return true;
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	LDB	R0 Li77:	;
STB     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "		return true;"


;; row: :	JMP	El101  ; 	    else
	JMP	El101:	;	    else
;; row: If98:	REF	Tv102 Iv20 Pa10; 		return insert(key, node.right);
If98:	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-61	;
	STR	R13 (R10)	;
;; row: :	FRAME	this Me11 ; 		return insert(key, node.right);
;; Call function "Me11:    		return insert(key, node.right);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-66          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me11
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Pa9  ; 		return insert(key, node.right);
;; parameters on the stack (Pa9)  ; 		return insert(key, node.right);
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv102  ; 		return insert(key, node.right);
;; parameters on the stack (Tv102)  ; 		return insert(key, node.right);
	MOV	R10 RFP	;Load Address
	ADI	R10 #-61	;
	LDR	R13 (R10)	;
	LDR	R1 (R13)	;Load to register
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me11  ; 		return insert(key, node.right);
;; local varibales on the stack    ; 		return insert(key, node.right);
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-66
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me11:
;; row: :	PEEK	Tv103  ; 		return insert(key, node.right);
	LDB	R11 (RSP)	;		return insert(key, node.right);
	MOV	R10 RFP	;
	ADI	R10 #-65	;
	STB	R11 (R10)	;
;; row: :	RETURN	Tv103  ; 		return insert(key, node.right);
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	MOV	R10 RFP	;
	ADI	R10 #-65	;
	LDB	R0 (R10)	;
STB     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "		return insert(key, node.right);"


;; row: El101:	JMP	El104  ; 	else
El101:	JMP	El104:	;	else
;; row: If95:	RETURN	Li105  ; 	    return false;
;; return from function
;; test for underflow
If95:	MOV	RSP RFP	; 	    return false;
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	LDB	R0 Li105:	;
STB     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "	    return false;"


;; row: El104:	RTN	  ;     }
;; return from function
;; test for underflow
El104:	MOV	RSP RFP	;     }
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "    }"


;; row: :	FUNC	Me12  ;     public void print() {
Me12:   ADI   R0 #0 ;    public void print() {
;; row: :	REF	Tv107 Iv3 this; 	first = true;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv107 Li77 ; 	first = true;
	LDB	R3 Li77:	;	first = true;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv108 Iv2 this; 	inorder(root);
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-13	;
	STR	R13 (R10)	;
;; row: :	FRAME	this Me14 ; 	inorder(root);
;; Call function "Me14:    	inorder(root);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-37          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me14
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv108  ; 	inorder(root);
;; parameters on the stack (Tv108)  ; 	inorder(root);
	MOV	R10 RFP	;Load Address
	ADI	R10 #-13	;
	LDR	R13 (R10)	;
	LDR	R1 (R13)	;Load to register
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me14  ; 	inorder(root);
;; local varibales on the stack    ; 	inorder(root);
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-37
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me14:
;; row: :	PEEK	Tv109  ; 	inorder(root);
	LDR	R11 (RSP)	;	inorder(root);
	MOV	R10 RFP	;
	ADI	R10 #-17	;
	STR	R11 (R10)	;
;; row: :	WRITE	Li110  ; 	cout << '\n';
	LDB	R0 Li110:	;
	TRP	#3	;	cout << '\n';
;; row: :	RTN	  ;     }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "    }"


;; row: :	FUNC	Me14  ;     private void inorder(iNode node) {
Me14:   ADI   R0 #0 ;    private void inorder(iNode node) {
;; row: :	EQ	Tv111 Pa13 Li53; 	if (node == null) return;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	LDR	R4 Li53:	;
	CMP	R3 R4	;	if (node == null) return;
	BRZ	R3 BT315:	
	SUB	R3 R3	; false branch
	JMP	BF316:	
BT315:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF316:	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STB	R3 (R10)	;
;; row: :	BF	Tv111 If112 ; 	if (node == null) return;
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDB	R3 (R10)	;
	BRZ	R3 If112:	;	if (node == null) return;
;; row: :	RTN	  ; 	if (node == null) return;
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "	if (node == null) return;"


;; row: If112:	REF	Tv113 Iv19 Pa13; 	inorder(node.left);
If112:	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-17	;
	STR	R13 (R10)	;
;; row: :	FRAME	this Me14 ; 	inorder(node.left);
;; Call function "Me14:    	inorder(node.left);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-37          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me14
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv113  ; 	inorder(node.left);
;; parameters on the stack (Tv113)  ; 	inorder(node.left);
	MOV	R10 RFP	;Load Address
	ADI	R10 #-17	;
	LDR	R13 (R10)	;
	LDR	R1 (R13)	;Load to register
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me14  ; 	inorder(node.left);
;; local varibales on the stack    ; 	inorder(node.left);
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-37
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me14:
;; row: :	PEEK	Tv114  ; 	inorder(node.left);
	LDR	R11 (RSP)	;	inorder(node.left);
	MOV	R10 RFP	;
	ADI	R10 #-21	;
	STR	R11 (R10)	;
;; row: :	FRAME	this Me16 ; 	visit(node);
;; Call function "Me16:    	visit(node);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-22          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me16
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Pa13  ; 	visit(node);
;; parameters on the stack (Pa13)  ; 	visit(node);
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me16  ; 	visit(node);
;; local varibales on the stack    ; 	visit(node);
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-22
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me16:
;; row: :	PEEK	Tv115  ; 	visit(node);
	LDR	R11 (RSP)	;	visit(node);
	MOV	R10 RFP	;
	ADI	R10 #-25	;
	STR	R11 (R10)	;
;; row: :	REF	Tv116 Iv20 Pa13; 	inorder(node.right);
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-29	;
	STR	R13 (R10)	;
;; row: :	FRAME	this Me14 ; 	inorder(node.right);
;; Call function "Me14:    	inorder(node.right);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-37          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me14
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv116  ; 	inorder(node.right);
;; parameters on the stack (Tv116)  ; 	inorder(node.right);
	MOV	R10 RFP	;Load Address
	ADI	R10 #-29	;
	LDR	R13 (R10)	;
	LDR	R1 (R13)	;Load to register
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me14  ; 	inorder(node.right);
;; local varibales on the stack    ; 	inorder(node.right);
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-37
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me14:
;; row: :	PEEK	Tv117  ; 	inorder(node.right);
	LDR	R11 (RSP)	;	inorder(node.right);
	MOV	R10 RFP	;
	ADI	R10 #-33	;
	STR	R11 (R10)	;
;; row: :	RTN	  ;     }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "    }"


;; row: :	FUNC	Me16  ;     private void visit(iNode node) {
Me16:   ADI   R0 #0 ;    private void visit(iNode node) {
;; row: :	REF	Tv118 Iv3 this; 	if (first) {
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	BF	Tv118 If119 ; 	if (first) {
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	LDB	R3 (R13)	;Load to register
	BRZ	R3 If119:	;	if (first) {
;; row: :	REF	Tv120 Iv3 this; 	    first = false;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-17	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv120 Li105 ; 	    first = false;
	LDB	R3 Li105:	;	    first = false;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-17	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	WRITE	Li121  ; 	    cout << ' ';
	LDB	R0 Li121:	;
	TRP	#3	;	    cout << ' ';
;; row: :	JMP	El122  ; 	else cout << ',';
	JMP	El122:	;	else cout << ',';
;; row: If119:	WRITE	Li123  ; 	else cout << ',';
If119:	LDB	R0 Li123:	;
	TRP	#3	;	else cout << ',';
;; row: El122:	REF	Tv124 Iv18 Pa15; 	cout << node.root;
El122:	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-18	;
	STR	R13 (R10)	;
;; row: :	WRITE	Tv124  ; 	cout << node.root;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-18	;
	LDR	R13 (R10)	;
	LDR	R0 (R13)	;Load to register
	TRP	#1	;	cout << node.root;
;; row: :	RTN	  ;     }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "    }"


;; row: :	FUNC	St47  ; }
St47:   ADI   R0 #0 ;}
;; row: :	REF	Tv49 Iv2 this;     private iNode root;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	REF	Tv51 Iv3 this;     private bool first;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	RTN	  ; }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "}"


;; row: :	FUNC	Co22  ;     iNode(int key) {
Co22:   ADI   R0 #0 ;    iNode(int key) {
;; row: :	FRAME	this St125 ;     iNode(int key) {
;; Call function "St125:        iNode(int key) {"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-24          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke St125
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	St125  ;     iNode(int key) {
;; local varibales on the stack    ;     iNode(int key) {
;; Temp variables on the stack
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-24
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     St125:
;; row: :	REF	Tv129 Iv18 this; 	root = key;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv129 Pa21 ; 	root = key;
	MOV	R10 RFP	;	root = key;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv130 Iv19 this; 	left = null;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-20	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv130 Li53 ; 	left = null;
	LDR	R3 Li53:	;	left = null;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-20	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv131 Iv20 this; 	right = null;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-24	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv131 Li53 ; 	right = null;
	LDR	R3 Li53:	;	right = null;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-24	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	RETURN	this  ;     }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R0 (R10)	;
STR     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "    }"


;; row: :	FUNC	St125  ; }
St125:   ADI   R0 #0 ;}
;; row: :	REF	Tv126 Iv18 this;     public int root;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	REF	Tv127 Iv19 this;     public iNode left;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	REF	Tv128 Iv20 this;     public iNode right;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-20	;
	STR	R13 (R10)	;
;; row: :	RTN	  ; }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "}"


;; row: :	FUNC	Co27  ;     Message() {
Co27:   ADI   R0 #0 ;    Message() {
;; row: :	FRAME	this St132 ;     Message() {
;; Call function "St132:        Message() {"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-21          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke St132
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	St132  ;     Message() {
;; local varibales on the stack    ;     Message() {
;; Temp variables on the stack
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-21
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     St132:
;; row: :	REF	Tv137 Iv24 this;     	msg = new char[100];
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	MUL	Tv140 1 Li138;     	msg = new char[100];
	SUB	R4 R4	;    	msg = new char[100];
	ADI	R4 #1	;    	msg = new char[100];
	LDR	R3 Li138:	;
	MUL	R3 R4	;    	msg = new char[100];
	MOV	R10 RFP	;
	ADI	R10 #-17	;
	STR	R3 (R10)	;
;; row: :	NEW	Tv140 Tv139 ;     	msg = new char[100];
	MOV	R10 RFP	;
	ADI	R10 #-17	;
	LDR	R3 (R10)	;
;; Test for heap overflow
MOV     R10 R9
ADD     R10 R3
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADD     R9 R3
	MOV	R10 RFP	;
	ADI	R10 #-13	;
	STR	R11 (R10)	;
;; row: :	MOV	Tv137 Tv139 ;     	msg = new char[100];
	MOV	R10 RFP	;    	msg = new char[100];
	ADI	R10 #-13	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv141 Iv24 this; 	msg[0] = 'A';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-21	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv142 Li55 Tv141; 	msg[0] = 'A';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-21	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li55:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-22	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv142 Li143 ; 	msg[0] = 'A';
	LDB	R3 Li143:	;	msg[0] = 'A';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-22	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv144 Iv24 this; 	msg[1] = 'd';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-23	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv145 Li59 Tv144; 	msg[1] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-23	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li59:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-24	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv145 Li146 ; 	msg[1] = 'd';
	LDB	R3 Li146:	;	msg[1] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-24	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv147 Iv24 this; 	msg[2] = 'd';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-25	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv148 Li65 Tv147; 	msg[2] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-25	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li65:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-26	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv148 Li146 ; 	msg[2] = 'd';
	LDB	R3 Li146:	;	msg[2] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-26	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv150 Iv24 this; 	msg[3] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-27	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv151 Li149 Tv150; 	msg[3] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-27	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li149:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-28	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv151 Li152 ; 	msg[3] = 'e';
	LDB	R3 Li152:	;	msg[3] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-28	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv154 Iv24 this; 	msg[4] = 'd';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-29	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv155 Li153 Tv154; 	msg[4] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-29	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li153:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-30	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv155 Li146 ; 	msg[4] = 'd';
	LDB	R3 Li146:	;	msg[4] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-30	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv157 Iv24 this; 	msg[5] = ' ';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-31	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv158 Li156 Tv157; 	msg[5] = ' ';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-31	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li156:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-32	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv158 Li121 ; 	msg[5] = ' ';
	LDB	R3 Li121:	;	msg[5] = ' ';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-32	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv160 Iv24 this; 	msg[6] = 'E';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-33	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv161 Li159 Tv160; 	msg[6] = 'E';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-33	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li159:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-34	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv161 Li162 ; 	msg[6] = 'E';
	LDB	R3 Li162:	;	msg[6] = 'E';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-34	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv164 Iv24 this; 	msg[7] = 'l';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-35	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv165 Li163 Tv164; 	msg[7] = 'l';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-35	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li163:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-36	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv165 Li166 ; 	msg[7] = 'l';
	LDB	R3 Li166:	;	msg[7] = 'l';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-36	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv168 Iv24 this; 	msg[8] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-37	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv169 Li167 Tv168; 	msg[8] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-37	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li167:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-38	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv169 Li152 ; 	msg[8] = 'e';
	LDB	R3 Li152:	;	msg[8] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-38	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv171 Iv24 this; 	msg[9] = 'm';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-39	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv172 Li170 Tv171; 	msg[9] = 'm';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-39	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li170:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-40	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv172 Li173 ; 	msg[9] = 'm';
	LDB	R3 Li173:	;	msg[9] = 'm';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-40	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv175 Iv24 this; 	msg[10] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-41	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv176 Li174 Tv175; 	msg[10] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-41	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li174:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-42	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv176 Li152 ; 	msg[10] = 'e';
	LDB	R3 Li152:	;	msg[10] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-42	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv178 Iv24 this; 	msg[11] = 'n';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-43	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv179 Li177 Tv178; 	msg[11] = 'n';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-43	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li177:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-44	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv179 Li180 ; 	msg[11] = 'n';
	LDB	R3 Li180:	;	msg[11] = 'n';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-44	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv182 Iv24 this; 	msg[12] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-45	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv183 Li181 Tv182; 	msg[12] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-45	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li181:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-46	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv183 Li184 ; 	msg[12] = 't';
	LDB	R3 Li184:	;	msg[12] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-46	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv186 Iv24 this; 	msg[13] = ':';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-47	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv187 Li185 Tv186; 	msg[13] = ':';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-47	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li185:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-48	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv187 Li188 ; 	msg[13] = ':';
	LDB	R3 Li188:	;	msg[13] = ':';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-48	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv189 Iv25 this; 	i = 14;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-49	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv189 Li190 ; 	i = 14;
	LDR	R3 Li190:	;	i = 14;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-49	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv191 Iv25 this; 	msg[i] = 'D';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-53	;
	STR	R13 (R10)	;
;; row: :	REF	Tv192 Iv24 this; 	msg[i] = 'D';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-57	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv193 Tv191 Tv192; 	msg[i] = 'D';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-57	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;Load Address
	ADI	R10 #-53	;
	LDR	R13 (R10)	;
	LDR	R14 (R13)	;Load to register
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-58	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv193 Li194 ; 	msg[i] = 'D';
	LDB	R3 Li194:	;	msg[i] = 'D';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-58	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv195 Iv25 this; 	msg[i+1] = 'u';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-59	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv196 Li59 Tv195; 	msg[i+1] = 'u';
	LDR	R4 Li59:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-59	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+1] = 'u';
	MOV	R10 RFP	;
	ADI	R10 #-63	;
	STR	R3 (R10)	;
;; row: :	REF	Tv197 Iv24 this; 	msg[i+1] = 'u';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-67	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv198 Tv196 Tv197; 	msg[i+1] = 'u';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-67	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-63	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-68	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv198 Li199 ; 	msg[i+1] = 'u';
	LDB	R3 Li199:	;	msg[i+1] = 'u';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-68	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv200 Iv25 this; 	msg[i+2] = 'p';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-69	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv201 Li65 Tv200; 	msg[i+2] = 'p';
	LDR	R4 Li65:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-69	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+2] = 'p';
	MOV	R10 RFP	;
	ADI	R10 #-73	;
	STR	R3 (R10)	;
;; row: :	REF	Tv202 Iv24 this; 	msg[i+2] = 'p';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-77	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv203 Tv201 Tv202; 	msg[i+2] = 'p';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-77	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-73	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-78	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv203 Li204 ; 	msg[i+2] = 'p';
	LDB	R3 Li204:	;	msg[i+2] = 'p';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-78	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv205 Iv25 this; 	msg[i+3] = 'l';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-79	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv206 Li149 Tv205; 	msg[i+3] = 'l';
	LDR	R4 Li149:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-79	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+3] = 'l';
	MOV	R10 RFP	;
	ADI	R10 #-83	;
	STR	R3 (R10)	;
;; row: :	REF	Tv207 Iv24 this; 	msg[i+3] = 'l';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-87	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv208 Tv206 Tv207; 	msg[i+3] = 'l';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-87	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-83	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-88	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv208 Li166 ; 	msg[i+3] = 'l';
	LDB	R3 Li166:	;	msg[i+3] = 'l';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-88	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv209 Iv25 this; 	msg[i+4] = 'i';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-89	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv210 Li153 Tv209; 	msg[i+4] = 'i';
	LDR	R4 Li153:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-89	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+4] = 'i';
	MOV	R10 RFP	;
	ADI	R10 #-93	;
	STR	R3 (R10)	;
;; row: :	REF	Tv211 Iv24 this; 	msg[i+4] = 'i';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-97	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv212 Tv210 Tv211; 	msg[i+4] = 'i';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-97	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-93	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-98	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv212 Li213 ; 	msg[i+4] = 'i';
	LDB	R3 Li213:	;	msg[i+4] = 'i';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-98	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv214 Iv25 this; 	msg[i+5] = 'c';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-99	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv215 Li156 Tv214; 	msg[i+5] = 'c';
	LDR	R4 Li156:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-99	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+5] = 'c';
	MOV	R10 RFP	;
	ADI	R10 #-103	;
	STR	R3 (R10)	;
;; row: :	REF	Tv216 Iv24 this; 	msg[i+5] = 'c';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-107	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv217 Tv215 Tv216; 	msg[i+5] = 'c';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-107	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-103	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-108	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv217 Li218 ; 	msg[i+5] = 'c';
	LDB	R3 Li218:	;	msg[i+5] = 'c';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-108	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv219 Iv25 this; 	msg[i+6] = 'a';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-109	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv220 Li159 Tv219; 	msg[i+6] = 'a';
	LDR	R4 Li159:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-109	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+6] = 'a';
	MOV	R10 RFP	;
	ADI	R10 #-113	;
	STR	R3 (R10)	;
;; row: :	REF	Tv221 Iv24 this; 	msg[i+6] = 'a';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-117	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv222 Tv220 Tv221; 	msg[i+6] = 'a';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-117	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-113	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-118	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv222 Li223 ; 	msg[i+6] = 'a';
	LDB	R3 Li223:	;	msg[i+6] = 'a';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-118	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv224 Iv25 this; 	msg[i+7] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-119	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv225 Li163 Tv224; 	msg[i+7] = 't';
	LDR	R4 Li163:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-119	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+7] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-123	;
	STR	R3 (R10)	;
;; row: :	REF	Tv226 Iv24 this; 	msg[i+7] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-127	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv227 Tv225 Tv226; 	msg[i+7] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-127	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-123	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-128	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv227 Li184 ; 	msg[i+7] = 't';
	LDB	R3 Li184:	;	msg[i+7] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-128	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv228 Iv25 this; 	msg[i+8] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-129	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv229 Li167 Tv228; 	msg[i+8] = 'e';
	LDR	R4 Li167:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-129	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+8] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-133	;
	STR	R3 (R10)	;
;; row: :	REF	Tv230 Iv24 this; 	msg[i+8] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-137	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv231 Tv229 Tv230; 	msg[i+8] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-137	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-133	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-138	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv231 Li152 ; 	msg[i+8] = 'e';
	LDB	R3 Li152:	;	msg[i+8] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-138	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv232 Iv25 this; 	msg[i+9] = 'm';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-139	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv233 Li170 Tv232; 	msg[i+9] = 'm';
	LDR	R4 Li170:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-139	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+9] = 'm';
	MOV	R10 RFP	;
	ADI	R10 #-143	;
	STR	R3 (R10)	;
;; row: :	REF	Tv234 Iv24 this; 	msg[i+9] = 'm';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-147	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv235 Tv233 Tv234; 	msg[i+9] = 'm';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-147	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-143	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-148	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv235 Li173 ; 	msg[i+9] = 'm';
	LDB	R3 Li173:	;	msg[i+9] = 'm';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-148	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv237 Iv24 this; 	msg[24] = 'E';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-149	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv238 Li236 Tv237; 	msg[24] = 'E';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-149	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li236:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-150	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv238 Li162 ; 	msg[24] = 'E';
	LDB	R3 Li162:	;	msg[24] = 'E';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-150	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv240 Iv24 this; 	msg[25] = 'n';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-151	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv241 Li239 Tv240; 	msg[25] = 'n';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-151	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li239:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-152	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv241 Li180 ; 	msg[25] = 'n';
	LDB	R3 Li180:	;	msg[25] = 'n';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-152	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv243 Iv24 this; 	msg[26] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-153	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv244 Li242 Tv243; 	msg[26] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-153	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li242:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-154	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv244 Li184 ; 	msg[26] = 't';
	LDB	R3 Li184:	;	msg[26] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-154	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv246 Iv24 this; 	msg[27] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-155	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv247 Li245 Tv246; 	msg[27] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-155	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li245:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-156	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv247 Li152 ; 	msg[27] = 'e';
	LDB	R3 Li152:	;	msg[27] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-156	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv249 Iv24 this; 	msg[28] = 'r';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-157	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv250 Li248 Tv249; 	msg[28] = 'r';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-157	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li248:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-158	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv250 Li251 ; 	msg[28] = 'r';
	LDB	R3 Li251:	;	msg[28] = 'r';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-158	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	RETURN	this  ;     }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R0 (R10)	;
STR     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "    }"


;; row: :	FUNC	Me30  ;     private void print(int i, int end) {
Me30:   ADI   R0 #0 ;    private void print(int i, int end) {
;; row: Wh252:	LTE	Tv253 Pa28 Pa29; 	while (i <= end) {
Wh252:	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R4 (R10)	;
	CMP	R3 R4	;	while (i <= end) {
	BLT	R3 BT317:	
	BRZ	R3 BT317:	
	SUB	R3 R3	; false branch
	JMP	BF318:	
BT317:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF318:	MOV	R10 RFP	;
	ADI	R10 #-20	;
	STB	R3 (R10)	;
;; row: :	BF	Tv253 En254 ; 	while (i <= end) {
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	LDB	R3 (R10)	;
	BRZ	R3 En254:	;	while (i <= end) {
;; row: :	REF	Tv255 Iv24 this; 	    cout << msg[i];
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-21	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv256 Pa28 Tv255; 	    cout << msg[i];
	MOV	R10 RFP	;Load Address
	ADI	R10 #-21	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-22	;
	STR	R13 (R10)	;
;; row: :	WRITE	Tv256  ; 	    cout << msg[i];
	MOV	R10 RFP	;Load Address
	ADI	R10 #-22	;
	LDR	R13 (R10)	;
	LDB	R0 (R13)	;Load to register
	TRP	#3	;	    cout << msg[i];
;; row: :	ADD	Tv257 Li59 Pa28; 	    i = i + 1;
	LDR	R4 Li59:	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	ADD	R3 R4	;	    i = i + 1;
	MOV	R10 RFP	;
	ADI	R10 #-23	;
	STR	R3 (R10)	;
;; row: :	MOV	Pa28 Tv257 ; 	    i = i + 1;
	MOV	R10 RFP	;	    i = i + 1;
	ADI	R10 #-23	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R3 (R10)	;
;; row: :	JMP	Wh252  ;     }	
	JMP	Wh252:	;    }	
;; row: En254:	RTN	  ;     }	
;; return from function
;; test for underflow
En254:	MOV	RSP RFP	;     }	
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "    }	"


;; row: :	FUNC	Me32  ;     public void msg1(int elm) {
Me32:   ADI   R0 #0 ;    public void msg1(int elm) {
;; row: :	FRAME	this Me30 ; 	print(0, 13);
;; Call function "Me30:    	print(0, 13);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-27          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me30
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li55  ; 	print(0, 13);
;; parameters on the stack (Li55)  ; 	print(0, 13);
	LDR	R1 Li55:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li185  ; 	print(0, 13);
;; parameters on the stack (Li185)  ; 	print(0, 13);
	LDR	R1 Li185:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me30  ; 	print(0, 13);
;; local varibales on the stack    ; 	print(0, 13);
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-27
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me30:
;; row: :	PEEK	Tv258  ; 	print(0, 13);
	LDR	R11 (RSP)	;	print(0, 13);
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STR	R11 (R10)	;
;; row: :	WRITE	Pa31  ; 	cout << elm;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R0 (R10)	;
	TRP	#1	;	cout << elm;
;; row: :	WRITE	Li110  ; 	cout << '\n';
	LDB	R0 Li110:	;
	TRP	#3	;	cout << '\n';
;; row: :	RTN	  ;     }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "    }"


;; row: :	FUNC	Me34  ;     public bool msg2(int elm) {
Me34:   ADI   R0 #0 ;    public bool msg2(int elm) {
;; row: :	REF	Tv259 Iv25 this; 	i = 14;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv259 Li190 ; 	i = 14;
	LDR	R3 Li190:	;	i = 14;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv260 Iv26 this; 	end = (i + 8);
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-20	;
	STR	R13 (R10)	;
;; row: :	REF	Tv261 Iv25 this; 	end = (i + 8);
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-24	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv262 Li167 Tv261; 	end = (i + 8);
	LDR	R4 Li167:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-24	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	end = (i + 8);
	MOV	R10 RFP	;
	ADI	R10 #-28	;
	STR	R3 (R10)	;
;; row: :	MOV	Tv260 Tv262 ; 	end = (i + 8);
	MOV	R10 RFP	;	end = (i + 8);
	ADI	R10 #-28	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-20	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv263 Iv25 this; 	print(i, end);
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-32	;
	STR	R13 (R10)	;
;; row: :	REF	Tv264 Iv26 this; 	print(i, end);
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-36	;
	STR	R13 (R10)	;
;; row: :	FRAME	this Me30 ; 	print(i, end);
;; Call function "Me30:    	print(i, end);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-27          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me30
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv263  ; 	print(i, end);
;; parameters on the stack (Tv263)  ; 	print(i, end);
	MOV	R10 RFP	;Load Address
	ADI	R10 #-32	;
	LDR	R13 (R10)	;
	LDR	R1 (R13)	;Load to register
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv264  ; 	print(i, end);
;; parameters on the stack (Tv264)  ; 	print(i, end);
	MOV	R10 RFP	;Load Address
	ADI	R10 #-36	;
	LDR	R13 (R10)	;
	LDR	R1 (R13)	;Load to register
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me30  ; 	print(i, end);
;; local varibales on the stack    ; 	print(i, end);
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-27
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me30:
;; row: :	PEEK	Tv265  ; 	print(i, end);
	LDR	R11 (RSP)	;	print(i, end);
	MOV	R10 RFP	;
	ADI	R10 #-40	;
	STR	R11 (R10)	;
;; row: :	REF	Tv266 Iv24 this; 	cout << msg[5];
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-44	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv267 Li156 Tv266; 	cout << msg[5];
	MOV	R10 RFP	;Load Address
	ADI	R10 #-44	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li156:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-45	;
	STR	R13 (R10)	;
;; row: :	WRITE	Tv267  ; 	cout << msg[5];
	MOV	R10 RFP	;Load Address
	ADI	R10 #-45	;
	LDR	R13 (R10)	;
	LDB	R0 (R13)	;Load to register
	TRP	#3	;	cout << msg[5];
;; row: :	FRAME	this Me30 ; 	print(6, 13);
;; Call function "Me30:    	print(6, 13);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-27          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me30
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li159  ; 	print(6, 13);
;; parameters on the stack (Li159)  ; 	print(6, 13);
	LDR	R1 Li159:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li185  ; 	print(6, 13);
;; parameters on the stack (Li185)  ; 	print(6, 13);
	LDR	R1 Li185:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me30  ; 	print(6, 13);
;; local varibales on the stack    ; 	print(6, 13);
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-27
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me30:
;; row: :	PEEK	Tv268  ; 	print(6, 13);
	LDR	R11 (RSP)	;	print(6, 13);
	MOV	R10 RFP	;
	ADI	R10 #-46	;
	STR	R11 (R10)	;
;; row: :	WRITE	Pa33  ; 	cout << elm;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R0 (R10)	;
	TRP	#1	;	cout << elm;
;; row: :	WRITE	Li110  ; 	cout << '\n';
	LDB	R0 Li110:	;
	TRP	#3	;	cout << '\n';
;; row: :	RTN	  ;     }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "    }"


;; row: :	FUNC	Me35  ;     public int msg3() {
Me35:   ADI   R0 #0 ;    public int msg3() {
;; row: :	FRAME	this Me30 ; 	print(24, 28);
;; Call function "Me30:    	print(24, 28);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-27          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me30
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li236  ; 	print(24, 28);
;; parameters on the stack (Li236)  ; 	print(24, 28);
	LDR	R1 Li236:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li248  ; 	print(24, 28);
;; parameters on the stack (Li248)  ; 	print(24, 28);
	LDR	R1 Li248:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me30  ; 	print(24, 28);
;; local varibales on the stack    ; 	print(24, 28);
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-27
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me30:
;; row: :	PEEK	Tv269  ; 	print(24, 28);
	LDR	R11 (RSP)	;	print(24, 28);
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R11 (R10)	;
;; row: :	REF	Tv270 Iv25 this; 	i = 5;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv270 Li156 ; 	i = 5;
	LDR	R3 Li156:	;	i = 5;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv271 Iv25 this; 	print(i, i);
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-20	;
	STR	R13 (R10)	;
;; row: :	REF	Tv272 Iv25 this; 	print(i, i);
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-24	;
	STR	R13 (R10)	;
;; row: :	FRAME	this Me30 ; 	print(i, i);
;; Call function "Me30:    	print(i, i);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-27          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me30
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv271  ; 	print(i, i);
;; parameters on the stack (Tv271)  ; 	print(i, i);
	MOV	R10 RFP	;Load Address
	ADI	R10 #-20	;
	LDR	R13 (R10)	;
	LDR	R1 (R13)	;Load to register
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv272  ; 	print(i, i);
;; parameters on the stack (Tv272)  ; 	print(i, i);
	MOV	R10 RFP	;Load Address
	ADI	R10 #-24	;
	LDR	R13 (R10)	;
	LDR	R1 (R13)	;Load to register
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me30  ; 	print(i, i);
;; local varibales on the stack    ; 	print(i, i);
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-27
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me30:
;; row: :	PEEK	Tv273  ; 	print(i, i);
	LDR	R11 (RSP)	;	print(i, i);
	MOV	R10 RFP	;
	ADI	R10 #-28	;
	STR	R11 (R10)	;
;; row: :	FRAME	this Me30 ; 	print(6, 13);
;; Call function "Me30:    	print(6, 13);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-27          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me30
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li159  ; 	print(6, 13);
;; parameters on the stack (Li159)  ; 	print(6, 13);
	LDR	R1 Li159:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li185  ; 	print(6, 13);
;; parameters on the stack (Li185)  ; 	print(6, 13);
	LDR	R1 Li185:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me30  ; 	print(6, 13);
;; local varibales on the stack    ; 	print(6, 13);
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-27
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me30:
;; row: :	PEEK	Tv274  ; 	print(6, 13);
	LDR	R11 (RSP)	;	print(6, 13);
	MOV	R10 RFP	;
	ADI	R10 #-32	;
	STR	R11 (R10)	;
;; row: :	RTN	  ;     }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "    }"


;; row: :	FUNC	St132  ; }
St132:   ADI   R0 #0 ;}
;; row: :	REF	Tv134 Iv24 this;     private char msg[];
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	REF	Tv135 Iv25 this;     private int i;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-13	;
	STR	R13 (R10)	;
;; row: :	REF	Tv136 Iv26 this;     private int end;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-17	;
	STR	R13 (R10)	;
;; row: :	RTN	  ; }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "}"


;; row: :	FUNC	Co41  ;       Butterfly(int age, char type) {
Co41:   ADI   R0 #0 ;      Butterfly(int age, char type) {
;; row: :	FRAME	this St275 ;       Butterfly(int age, char type) {
;; Call function "St275:          Butterfly(int age, char type) {"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-17          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke St275
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	St275  ;       Butterfly(int age, char type) {
;; local varibales on the stack    ;       Butterfly(int age, char type) {
;; Temp variables on the stack
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-17
;; set the frame pointer
MOV     RFP R15
        TRP #99
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     St275:
        TRP     #99
;; row: :	WRITE	Pa39  ;           cout << age;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R0 (R10)	;
	TRP	#1	;          cout << age;
;; row: :	WRITE	Li110  ;           cout << '\n';
	LDB	R0 Li110:	;
	TRP	#3	;          cout << '\n';
;; row: :	WRITE	Pa40  ; 	  cout << type;
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDB	R0 (R10)	;
	TRP	#3	;	  cout << type;
;; row: :	WRITE	Li110  ;           cout << '\n';
	LDB	R0 Li110:	;
	TRP	#3	;          cout << '\n';
;; row: :	RETURN	this  ;       }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R0 (R10)	;
STR     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "      }"


;; row: :	FUNC	Me42  ;       public void nest() {
Me42:   ADI   R0 #0 ;      public void nest() {
;; row: :	REF	Tv279 Iv37 this;           cout << age;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	WRITE	Tv279  ;           cout << age;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	LDR	R0 (R13)	;Load to register
	TRP	#1	;          cout << age;
;; row: :	WRITE	Li110  ;           cout << '\n';
	LDB	R0 Li110:	;
	TRP	#3	;          cout << '\n';
;; row: :	REF	Tv280 Iv38 this; 	  cout << type;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	WRITE	Tv280  ; 	  cout << type;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	LDB	R0 (R13)	;Load to register
	TRP	#3	;	  cout << type;
;; row: :	WRITE	Li110  ;           cout << '\n';
	LDB	R0 Li110:	;
	TRP	#3	;          cout << '\n';
;; row: :	RTN	  ;       }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "      }"


;; row: :	FUNC	St275  ; }
St275:   ADI   R0 #0 ;}
;; row: :	REF	Tv276 Iv37 this;       private int age = 42;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv276 Li277 ;       private int age = 42;
	LDR	R3 Li277:	;      private int age = 42;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv278 Iv38 this;       private char type = 'm';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv278 Li173 ;       private char type = 'm';
	LDB	R3 Li173:	;      private char type = 'm';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	RTN	  ; }
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "}"


;; row: :	FUNC	MAIN  ; void main() {
MAIN:   ADI   R0 #0 ;void main() {
;; row: :	NEWI	Cl23 Tv283 ;     Message msg = new Message();
;; Test for heap overflow
	MOV     R10 R9
ADI     R10 #12
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADI     R9 #12
	MOV	R10 RFP	;
	ADI	R10 #-28	;
	STR	R11 (R10)	;
;; row: :	FRAME	Tv283 Co27 ;     Message msg = new Message();
;; Call function "Co27:        Message msg = new Message();"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-159          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Co27
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-28	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Co27  ;     Message msg = new Message();
;; local varibales on the stack    ;     Message msg = new Message();
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-1
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-159
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Co27:
;; row: :	PEEK	Tv283  ;     Message msg = new Message();
	LDR	R11 (RSP)	;    Message msg = new Message();
	MOV	R10 RFP	;
	ADI	R10 #-28	;
	STR	R11 (R10)	;
;; row: :	MOV	Lv45 Tv283 ;     Message msg = new Message();
	MOV	R10 RFP	;    Message msg = new Message();
	ADI	R10 #-28	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	STR	R3 (R10)	;
;; row: :	NEWI	Cl36 Tv287 ;     Butterfly bff = new Butterfly(37, 'g');
;; Test for heap overflow
	MOV     R10 R9
ADI     R10 #5
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADI     R9 #5
	MOV	R10 RFP	;
	ADI	R10 #-32	;
	STR	R11 (R10)	;
;; row: :	FRAME	Tv287 Co41 ;     Butterfly bff = new Butterfly(37, 'g');
;; Call function "Co41:        Butterfly bff = new Butterfly(37, 'g');"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-17          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Co41
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-32	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li285  ;     Butterfly bff = new Butterfly(37, 'g');
;; parameters on the stack (Li285)  ;     Butterfly bff = new Butterfly(37, 'g');
	LDR	R1 Li285:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li286  ;     Butterfly bff = new Butterfly(37, 'g');
;; parameters on the stack (Li286)  ;     Butterfly bff = new Butterfly(37, 'g');
	LDB	R1 Li286:	;
STB     R1 (RSP)
ADI     RSP #-1
;; row: :	CALL	Co41  ;     Butterfly bff = new Butterfly(37, 'g');
;; local varibales on the stack    ;     Butterfly bff = new Butterfly(37, 'g');
;; Temp variables on the stack
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-17
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Co41:
;; row: :	PEEK	Tv287  ;     Butterfly bff = new Butterfly(37, 'g');
	LDR	R11 (RSP)	;    Butterfly bff = new Butterfly(37, 'g');
	MOV	R10 RFP	;
	ADI	R10 #-32	;
	STR	R11 (R10)	;
;; row: :	MOV	Lv46 Tv287 ;     Butterfly bff = new Butterfly(37, 'g');
	MOV	R10 RFP	;    Butterfly bff = new Butterfly(37, 'g');
	ADI	R10 #-32	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-24	;
	STR	R3 (R10)	;
;; row: :	FRAME	Lv46 Me42 ;     bff.nest();
;; Call function "Me42:        bff.nest();"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-17          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me42
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-24	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me42  ;     bff.nest();
;; local varibales on the stack    ;     bff.nest();
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-1
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-17
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me42:
;; row: :	PEEK	Tv288  ;     bff.nest();
	LDR	R11 (RSP)	;    bff.nest();
	MOV	R10 RFP	;
	ADI	R10 #-36	;
	STR	R11 (R10)	;
;; row: :	NEWI	Cl1 Tv289 ;     tree = new iTree();
;; Test for heap overflow
	MOV     R10 R9
ADI     R10 #5
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADI     R9 #5
	MOV	R10 RFP	;
	ADI	R10 #-40	;
	STR	R11 (R10)	;
;; row: :	FRAME	Tv289 Co4 ;     tree = new iTree();
;; Call function "Co4:        tree = new iTree();"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-16          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Co4
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-40	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Co4  ;     tree = new iTree();
;; local varibales on the stack    ;     tree = new iTree();
;; Temp variables on the stack
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-16
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Co4:
;; row: :	PEEK	Tv289  ;     tree = new iTree();
	LDR	R11 (RSP)	;    tree = new iTree();
	MOV	R10 RFP	;
	ADI	R10 #-40	;
	STR	R11 (R10)	;
;; row: :	MOV	Lv44 Tv289 ;     tree = new iTree();
	MOV	R10 RFP	;    tree = new iTree();
	ADI	R10 #-40	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STR	R3 (R10)	;
;; row: :	FRAME	Lv45 Me35 ;     msg.msg3();
;; Call function "Me35:        msg.msg3();"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-36          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me35
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me35  ;     msg.msg3();
;; local varibales on the stack    ;     msg.msg3();
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-36
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me35:
;; row: :	PEEK	Tv290  ;     msg.msg3();
	LDR	R11 (RSP)	;    msg.msg3();
	MOV	R10 RFP	;
	ADI	R10 #-44	;
	STR	R11 (R10)	;
;; row: :	READ	Lv43  ;     cin >> key;
	TRP	#2	;    cin >> key;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R0 (R10)	;
;; row: :	WRITE	Li110  ;     cout << '\n';
	LDB	R0 Li110:	;
	TRP	#3	;    cout << '\n';
;; row: Wh291:	NEQ	Tv292 Lv43 Li55;     while (key != 0) {
Wh291:	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	LDR	R4 Li55:	;
	CMP	R3 R4	;    while (key != 0) {
	BNZ	R3 BT319:	
	SUB	R3 R3	; false branch
	JMP	BF320:	
BT319:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF320:	MOV	R10 RFP	;
	ADI	R10 #-48	;
	STB	R3 (R10)	;
;; row: :	BF	Tv292 En293 ;     while (key != 0) {
	MOV	R10 RFP	;
	ADI	R10 #-48	;
	LDB	R3 (R10)	;
	BRZ	R3 En293:	;    while (key != 0) {
;; row: :	FRAME	Lv44 Me8 ; 	if (tree.add(key)) {
;; Call function "Me8:    	if (tree.add(key)) {"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-46          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me8
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Lv43  ; 	if (tree.add(key)) {
;; parameters on the stack (Lv43)  ; 	if (tree.add(key)) {
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me8  ; 	if (tree.add(key)) {
;; local varibales on the stack    ; 	if (tree.add(key)) {
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-46
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me8:
;; row: :	PEEK	Tv294  ; 	if (tree.add(key)) {
	LDB	R11 (RSP)	;	if (tree.add(key)) {
	MOV	R10 RFP	;
	ADI	R10 #-49	;
	STB	R11 (R10)	;
;; row: :	BF	Tv294 If295 ; 	if (tree.add(key)) {
	MOV	R10 RFP	;
	ADI	R10 #-49	;
	LDB	R3 (R10)	;
	BRZ	R3 If295:	;	if (tree.add(key)) {
;; row: :	FRAME	Lv45 Me32 ; 	    msg.msg1(key);
;; Call function "Me32:    	    msg.msg1(key);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-20          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me32
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Lv43  ; 	    msg.msg1(key);
;; parameters on the stack (Lv43)  ; 	    msg.msg1(key);
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me32  ; 	    msg.msg1(key);
;; local varibales on the stack    ; 	    msg.msg1(key);
;; Temp variables on the stack
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-20
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me32:
;; row: :	PEEK	Tv296  ; 	    msg.msg1(key);
	LDR	R11 (RSP)	;	    msg.msg1(key);
	MOV	R10 RFP	;
	ADI	R10 #-50	;
	STR	R11 (R10)	;
;; row: :	FRAME	Lv44 Me12 ; 	    tree.print();
;; Call function "Me12:    	    tree.print();"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-21          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me12
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me12  ; 	    tree.print();
;; local varibales on the stack    ; 	    tree.print();
;; Temp variables on the stack
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-21
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me12:
;; row: :	PEEK	Tv297  ; 	    tree.print();
	LDR	R11 (RSP)	;	    tree.print();
	MOV	R10 RFP	;
	ADI	R10 #-54	;
	STR	R11 (R10)	;
;; row: :	JMP	El298  ; 	else 
	JMP	El298:	;	else 
;; row: If295:	FRAME	Lv45 Me34 ; 	    msg.msg2(key);
;; Call function "Me34:    	    msg.msg2(key);"
;; Test for overflow
If295:   MOV     R10 RSP
ADI     R10 #-50          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me34
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Lv43  ; 	    msg.msg2(key);
;; parameters on the stack (Lv43)  ; 	    msg.msg2(key);
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me34  ; 	    msg.msg2(key);
;; local varibales on the stack    ; 	    msg.msg2(key);
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-1
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-50
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me34:
;; row: :	PEEK	Tv299  ; 	    msg.msg2(key);
	LDB	R11 (RSP)	;	    msg.msg2(key);
	MOV	R10 RFP	;
	ADI	R10 #-58	;
	STB	R11 (R10)	;
;; row: El298:	FRAME	Lv45 Me35 ; 	msg.msg3();
;; Call function "Me35:    	msg.msg3();"
;; Test for overflow
El298:   MOV     R10 RSP
ADI     R10 #-36          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Me35
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me35  ; 	msg.msg3();
;; local varibales on the stack    ; 	msg.msg3();
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-36
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me35:
;; row: :	PEEK	Tv300  ; 	msg.msg3();
	LDR	R11 (RSP)	;	msg.msg3();
	MOV	R10 RFP	;
	ADI	R10 #-59	;
	STR	R11 (R10)	;
;; row: :	READ	Lv43  ; 	cin >> key;
	TRP	#2	;	cin >> key;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R0 (R10)	;
;; row: :	WRITE	Li110  ; 	cout << '\n';
	LDB	R0 Li110:	;
	TRP	#3	;	cout << '\n';
;; row: :	JMP	Wh291  ; }
	JMP	Wh291:	;}
;; row: En293:	RTN	  ; }
;; return from function
;; test for underflow
En293:	MOV	RSP RFP	; }
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "}"


;; Heap starts here
FREE:    .INT 0
