LDA     R9 FREE:
;; Call function "MAIN:"
;; Test for overflow
MOV     R10 RSP
ADI     R10 #-57          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
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
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
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
Li143:	.INT	100
Li190:	.INT	13
Li228:	.BYT	'a'
Li168:	.INT	7
Li179:	.INT	10
Li199:	.BYT	'D'
Li250:	.INT	27
Li172:	.INT	8
Li175:	.INT	9
Li158:	.INT	4
Li82:	.BYT	1
Li223:	.BYT	'c'
Li293:	.INT	256
Li195:	.INT	14
Li128:	.BYT	','
Li244:	.INT	25
Li70:	.INT	2
Li241:	.INT	24
Li289:	.INT	512
Li303:	.INT	5000
Li115:	.BYT	'\n'
Li189:	.BYT	't'
Li167:	.BYT	'E'
Li247:	.INT	26
Li253:	.INT	28
Li218:	.BYT	'i'
Li110:	.BYT	0
Li60:	.INT	0
Li157:	.BYT	'e'
Li161:	.INT	5
Li209:	.BYT	'p'
Li182:	.INT	11
Li154:	.INT	3
Li256:	.BYT	'r'
Li64:	.INT	1
Li151:	.BYT	'd'
Li186:	.INT	12
Li126:	.BYT	32
Li171:	.BYT	'l'
Li286:	.INT	1000
Li164:	.INT	6
Li178:	.BYT	'm'
Li185:	.BYT	'n'
Li204:	.BYT	'u'
Li148:	.BYT	'A'
Li193:	.BYT	':'
Li58:	.INT	0
;; functions
;; row: :	FUNC	Co4  ;     iTree() {
Co4:   ADI   R0 #0 ;    iTree() {
;; row: :	FRAME	this St52 ;     iTree() {
;; Call function "St52:        iTree() {"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-17          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke St52
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
;; row: :	CALL	St52  ;     iTree() {
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
JMP     St52:
;; row: :	REF	Tv57 Iv2 this; 	root = null;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv57 Li58 ; 	root = null;
	LDR	R3 Li58:	;	root = null;
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
;; row: :	EQ	Tv61 Pa5 Li60; 	if (root == 0) return 0;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	LDR	R4 Li60:	;
	CMP	R3 R4	;	if (root == 0) return 0;
	BRZ	R3 BT339:	
	SUB	R3 R3	; false branch
	JMP	BF340:	
BT339:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF340:	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STB	R3 (R10)	;
;; row: :	BF	Tv61 If62 ; 	if (root == 0) return 0;
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDB	R3 (R10)	;
	BRZ	R3 If62:	;	if (root == 0) return 0;
;; row: :	RETURN	Li60  ; 	if (root == 0) return 0;
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	LDR	R0 Li60:	;
STR     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "	if (root == 0) return 0;"


;; row: :	JMP	El67  ; 	else if (root == 1) return 1;
	JMP	El67:	;	else if (root == 1) return 1;
;; row: If62:	EQ	Tv65 Pa5 Li64; 	else if (root == 1) return 1;
If62:	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	LDR	R4 Li64:	;
	CMP	R3 R4	;	else if (root == 1) return 1;
	BRZ	R3 BT341:	
	SUB	R3 R3	; false branch
	JMP	BF342:	
BT341:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF342:	MOV	R10 RFP	;
	ADI	R10 #-17	;
	STB	R3 (R10)	;
;; row: :	BF	Tv65 If66 ; 	else if (root == 1) return 1;
	MOV	R10 RFP	;
	ADI	R10 #-17	;
	LDB	R3 (R10)	;
	BRZ	R3 If66:	;	else if (root == 1) return 1;
;; row: :	RETURN	Li64  ; 	else if (root == 1) return 1;
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	LDR	R0 Li64:	;
STR     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "	else if (root == 1) return 1;"


;; row: :	JMP	El67  ; 	else return (fib(root - 1) + fib(root - 2));
	JMP	El67:	;	else return (fib(root - 1) + fib(root - 2));
;; row: If66:	SUB	Tv68 Li64 Pa5; 	else return (fib(root - 1) + fib(root - 2));
If66:	LDR	R4 Li64:	;
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
;; row: :	PUSH	Tv68  ; 	else return (fib(root - 1) + fib(root - 2));
;; parameters on the stack (Tv68)  ; 	else return (fib(root - 1) + fib(root - 2));
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
;; row: :	PEEK	Tv69  ; 	else return (fib(root - 1) + fib(root - 2));
	LDR	R11 (RSP)	;	else return (fib(root - 1) + fib(root - 2));
	MOV	R10 RFP	;
	ADI	R10 #-22	;
	STR	R11 (R10)	;
;; row: :	SUB	Tv71 Li70 Pa5; 	else return (fib(root - 1) + fib(root - 2));
	LDR	R4 Li70:	;
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
;; row: :	PUSH	Tv71  ; 	else return (fib(root - 1) + fib(root - 2));
;; parameters on the stack (Tv71)  ; 	else return (fib(root - 1) + fib(root - 2));
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
;; row: :	PEEK	Tv72  ; 	else return (fib(root - 1) + fib(root - 2));
	LDR	R11 (RSP)	;	else return (fib(root - 1) + fib(root - 2));
	MOV	R10 RFP	;
	ADI	R10 #-30	;
	STR	R11 (R10)	;
;; row: :	ADD	Tv73 Tv72 Tv69; 	else return (fib(root - 1) + fib(root - 2));
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
;; row: :	RETURN	Tv73  ; 	else return (fib(root - 1) + fib(root - 2));
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


;; row: El67:	RTN	  ;     }
;; return from function
;; test for underflow
El67:	MOV	RSP RFP	;     }
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
;; row: :	DIV	Tv74 Li70 Pa7; 	key = key + fib(key/2);
	LDR	R4 Li70:	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	DIV	R3 R4	;	key = key + fib(key/2);
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STR	R3 (R10)	;
;; row: :	FRAME	this Me6 ; 	key = key + fib(key/2);
;; Call function "Me6:    	key = key + fib(key/2);"
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
;; row: :	PUSH	Tv74  ; 	key = key + fib(key/2);
;; parameters on the stack (Tv74)  ; 	key = key + fib(key/2);
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me6  ; 	key = key + fib(key/2);
;; local varibales on the stack    ; 	key = key + fib(key/2);
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
;; row: :	PEEK	Tv75  ; 	key = key + fib(key/2);
	LDR	R11 (RSP)	;	key = key + fib(key/2);
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	STR	R11 (R10)	;
;; row: :	ADD	Tv76 Tv75 Pa7; 	key = key + fib(key/2);
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	LDR	R4 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	ADD	R3 R4	;	key = key + fib(key/2);
	MOV	R10 RFP	;
	ADI	R10 #-24	;
	STR	R3 (R10)	;
;; row: :	MOV	Pa7 Tv76 ; 	key = key + fib(key/2);
	MOV	R10 RFP	;	key = key + fib(key/2);
	ADI	R10 #-24	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R3 (R10)	;
;; row: :	REF	Tv77 Iv2 this; 	if (root == null) {
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-28	;
	STR	R13 (R10)	;
;; row: :	EQ	Tv78 Tv77 Li58; 	if (root == null) {
	MOV	R10 RFP	;Load Address
	ADI	R10 #-28	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	LDR	R4 Li58:	;
	CMP	R3 R4	;	if (root == null) {
	BRZ	R3 BT343:	
	SUB	R3 R3	; false branch
	JMP	BF344:	
BT343:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF344:	MOV	R10 RFP	;
	ADI	R10 #-32	;
	STB	R3 (R10)	;
;; row: :	BF	Tv78 If79 ; 	if (root == null) {
	MOV	R10 RFP	;
	ADI	R10 #-32	;
	LDB	R3 (R10)	;
	BRZ	R3 If79:	;	if (root == null) {
;; row: :	REF	Tv80 Iv2 this; 	    root = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-33	;
	STR	R13 (R10)	;
;; row: :	NEWI	Cl17 Tv81 ; 	    root = new iNode(key);
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
;; row: :	FRAME	Tv81 Co22 ; 	    root = new iNode(key);
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
;; row: :	PEEK	Tv81  ; 	    root = new iNode(key);
	LDR	R11 (RSP)	;	    root = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-37	;
	STR	R11 (R10)	;
;; row: :	MOV	Tv80 Tv81 ; 	    root = new iNode(key);
	MOV	R10 RFP	;	    root = new iNode(key);
	ADI	R10 #-37	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-33	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	RETURN	Li82  ; 	    return true;
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	LDB	R0 Li82:	;
STB     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "	    return true;"


;; row: :	JMP	El83  ; 	else
	JMP	El83:	;	else
;; row: If79:	REF	Tv84 Iv2 this; 	    return insert(key, root);
If79:	MOV	R10 RFP	;
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
ADI     R10 #-72          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
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
;; row: :	PUSH	Tv84  ; 	    return insert(key, root);
;; parameters on the stack (Tv84)  ; 	    return insert(key, root);
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
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-72
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me11:
;; row: :	PEEK	Tv85  ; 	    return insert(key, root);
	LDR	R11 (RSP)	;	    return insert(key, root);
	MOV	R10 RFP	;
	ADI	R10 #-45	;
	STR	R11 (R10)	;
;; row: :	RETURN	Tv85  ; 	    return insert(key, root);
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
	LDR	R0 (R10)	;
STR     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "	    return insert(key, root);"


;; row: El83:	RTN	  ;     }
;; return from function
;; test for underflow
El83:	MOV	RSP RFP	;     }
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
;; row: :	REF	Tv86 Iv18 Pa10; 	if (key < node.root)
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-20	;
	STR	R13 (R10)	;
;; row: :	LT	Tv87 Pa9 Tv86; 	if (key < node.root)
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-20	;
	LDR	R13 (R10)	;
	LDR	R4 (R13)	;Load to register
	CMP	R3 R4	;	if (key < node.root)
	BLT	R3 BT345:	
	SUB	R3 R3	; false branch
	JMP	BF346:	
BT345:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF346:	MOV	R10 RFP	;
	ADI	R10 #-24	;
	STB	R3 (R10)	;
;; row: :	BF	Tv87 If88 ; 	if (key < node.root)
	MOV	R10 RFP	;
	ADI	R10 #-24	;
	LDB	R3 (R10)	;
	BRZ	R3 If88:	;	if (key < node.root)
;; row: :	REF	Tv89 Iv19 Pa10; 	    if (node.left == null) {
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-25	;
	STR	R13 (R10)	;
;; row: :	EQ	Tv90 Tv89 Li58; 	    if (node.left == null) {
	MOV	R10 RFP	;Load Address
	ADI	R10 #-25	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	LDR	R4 Li58:	;
	CMP	R3 R4	;	    if (node.left == null) {
	BRZ	R3 BT347:	
	SUB	R3 R3	; false branch
	JMP	BF348:	
BT347:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF348:	MOV	R10 RFP	;
	ADI	R10 #-29	;
	STB	R3 (R10)	;
;; row: :	BF	Tv90 If91 ; 	    if (node.left == null) {
	MOV	R10 RFP	;
	ADI	R10 #-29	;
	LDB	R3 (R10)	;
	BRZ	R3 If91:	;	    if (node.left == null) {
;; row: :	REF	Tv92 Iv19 Pa10; 		node.left = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-30	;
	STR	R13 (R10)	;
;; row: :	NEWI	Cl17 Tv93 ; 		node.left = new iNode(key);
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
;; row: :	FRAME	Tv93 Co22 ; 		node.left = new iNode(key);
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
;; row: :	PEEK	Tv93  ; 		node.left = new iNode(key);
	LDR	R11 (RSP)	;		node.left = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-34	;
	STR	R11 (R10)	;
;; row: :	MOV	Tv92 Tv93 ; 		node.left = new iNode(key);
	MOV	R10 RFP	;		node.left = new iNode(key);
	ADI	R10 #-34	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-30	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	RETURN	Li82  ; 		return true;
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	LDB	R0 Li82:	;
STB     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "		return true;"


;; row: :	JMP	El94  ; 	    else 
	JMP	El94:	;	    else 
;; row: If91:	REF	Tv95 Iv19 Pa10; 		return insert(key, node.left);
If91:	MOV	R10 RFP	;
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
ADI     R10 #-72          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
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
;; row: :	PUSH	Tv95  ; 		return insert(key, node.left);
;; parameters on the stack (Tv95)  ; 		return insert(key, node.left);
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
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-72
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me11:
;; row: :	PEEK	Tv96  ; 		return insert(key, node.left);
	LDR	R11 (RSP)	;		return insert(key, node.left);
	MOV	R10 RFP	;
	ADI	R10 #-42	;
	STR	R11 (R10)	;
;; row: :	RETURN	Tv96  ; 		return insert(key, node.left);
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
	LDR	R0 (R10)	;
STR     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "		return insert(key, node.left);"


;; row: El94:	JMP	El109  ; 	else if (key > node.root)
El94:	JMP	El109:	;	else if (key > node.root)
;; row: If88:	REF	Tv98 Iv18 Pa10; 	else if (key > node.root)
If88:	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-46	;
	STR	R13 (R10)	;
;; row: :	GT	Tv99 Pa9 Tv98; 	else if (key > node.root)
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-46	;
	LDR	R13 (R10)	;
	LDR	R4 (R13)	;Load to register
	CMP	R3 R4	;	else if (key > node.root)
	BGT	R3 BT349:	
	SUB	R3 R3	; false branch
	JMP	BF350:	
BT349:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF350:	MOV	R10 RFP	;
	ADI	R10 #-50	;
	STB	R3 (R10)	;
;; row: :	BF	Tv99 If100 ; 	else if (key > node.root)
	MOV	R10 RFP	;
	ADI	R10 #-50	;
	LDB	R3 (R10)	;
	BRZ	R3 If100:	;	else if (key > node.root)
;; row: :	REF	Tv101 Iv20 Pa10; 	    if (node.right == null) {
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-51	;
	STR	R13 (R10)	;
;; row: :	EQ	Tv102 Tv101 Li58; 	    if (node.right == null) {
	MOV	R10 RFP	;Load Address
	ADI	R10 #-51	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	LDR	R4 Li58:	;
	CMP	R3 R4	;	    if (node.right == null) {
	BRZ	R3 BT351:	
	SUB	R3 R3	; false branch
	JMP	BF352:	
BT351:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF352:	MOV	R10 RFP	;
	ADI	R10 #-55	;
	STB	R3 (R10)	;
;; row: :	BF	Tv102 If103 ; 	    if (node.right == null) {
	MOV	R10 RFP	;
	ADI	R10 #-55	;
	LDB	R3 (R10)	;
	BRZ	R3 If103:	;	    if (node.right == null) {
;; row: :	REF	Tv104 Iv20 Pa10; 		node.right = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-56	;
	STR	R13 (R10)	;
;; row: :	NEWI	Cl17 Tv105 ; 		node.right = new iNode(key);
;; Test for heap overflow
	MOV     R10 R9
ADI     R10 #12
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADI     R9 #12
	MOV	R10 RFP	;
	ADI	R10 #-60	;
	STR	R11 (R10)	;
;; row: :	FRAME	Tv105 Co22 ; 		node.right = new iNode(key);
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
	ADI	R10 #-60	;
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
;; row: :	PEEK	Tv105  ; 		node.right = new iNode(key);
	LDR	R11 (RSP)	;		node.right = new iNode(key);
	MOV	R10 RFP	;
	ADI	R10 #-60	;
	STR	R11 (R10)	;
;; row: :	MOV	Tv104 Tv105 ; 		node.right = new iNode(key);
	MOV	R10 RFP	;		node.right = new iNode(key);
	ADI	R10 #-60	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-56	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	RETURN	Li82  ; 		return true;
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	LDB	R0 Li82:	;
STB     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "		return true;"


;; row: :	JMP	El106  ; 	    else
	JMP	El106:	;	    else
;; row: If103:	REF	Tv107 Iv20 Pa10; 		return insert(key, node.right);
If103:	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-64	;
	STR	R13 (R10)	;
;; row: :	FRAME	this Me11 ; 		return insert(key, node.right);
;; Call function "Me11:    		return insert(key, node.right);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-72          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
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
;; row: :	PUSH	Tv107  ; 		return insert(key, node.right);
;; parameters on the stack (Tv107)  ; 		return insert(key, node.right);
	MOV	R10 RFP	;Load Address
	ADI	R10 #-64	;
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
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-1
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-72
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me11:
;; row: :	PEEK	Tv108  ; 		return insert(key, node.right);
	LDR	R11 (RSP)	;		return insert(key, node.right);
	MOV	R10 RFP	;
	ADI	R10 #-68	;
	STR	R11 (R10)	;
;; row: :	RETURN	Tv108  ; 		return insert(key, node.right);
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	MOV	R10 RFP	;
	ADI	R10 #-68	;
	LDR	R0 (R10)	;
STR     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "		return insert(key, node.right);"


;; row: El106:	JMP	El109  ; 	else
El106:	JMP	El109:	;	else
;; row: If100:	RETURN	Li110  ; 	    return false;
;; return from function
;; test for underflow
If100:	MOV	RSP RFP	; 	    return false;
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	LDB	R0 Li110:	;
STB     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "	    return false;"


;; row: El109:	RTN	  ;     }
;; return from function
;; test for underflow
El109:	MOV	RSP RFP	;     }
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
;; row: :	REF	Tv112 Iv3 this; 	first = true;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv112 Li82 ; 	first = true;
	LDB	R3 Li82:	;	first = true;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv113 Iv2 this; 	inorder(root);
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
;; row: :	PUSH	Tv113  ; 	inorder(root);
;; parameters on the stack (Tv113)  ; 	inorder(root);
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
;; row: :	PEEK	Tv114  ; 	inorder(root);
	LDR	R11 (RSP)	;	inorder(root);
	MOV	R10 RFP	;
	ADI	R10 #-17	;
	STR	R11 (R10)	;
;; row: :	WRITE	Li115  ; 	cout << ('\n');
	LDB	R0 Li115:	;
	TRP	#3	;	cout << ('\n');
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
;; row: :	EQ	Tv116 Pa13 Li58; 	if (node == null) return;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	LDR	R4 Li58:	;
	CMP	R3 R4	;	if (node == null) return;
	BRZ	R3 BT353:	
	SUB	R3 R3	; false branch
	JMP	BF354:	
BT353:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF354:	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STB	R3 (R10)	;
;; row: :	BF	Tv116 If117 ; 	if (node == null) return;
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDB	R3 (R10)	;
	BRZ	R3 If117:	;	if (node == null) return;
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


;; row: If117:	REF	Tv118 Iv19 Pa13; 	inorder(node.left);
If117:	MOV	R10 RFP	;
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
;; row: :	PUSH	Tv118  ; 	inorder(node.left);
;; parameters on the stack (Tv118)  ; 	inorder(node.left);
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
;; row: :	PEEK	Tv119  ; 	inorder(node.left);
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
;; row: :	PEEK	Tv120  ; 	visit(node);
	LDR	R11 (RSP)	;	visit(node);
	MOV	R10 RFP	;
	ADI	R10 #-25	;
	STR	R11 (R10)	;
;; row: :	REF	Tv121 Iv20 Pa13; 	inorder(node.right);
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
;; row: :	PUSH	Tv121  ; 	inorder(node.right);
;; parameters on the stack (Tv121)  ; 	inorder(node.right);
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
;; row: :	PEEK	Tv122  ; 	inorder(node.right);
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
;; row: :	REF	Tv123 Iv3 this; 	if (first) {
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	BF	Tv123 If124 ; 	if (first) {
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	LDB	R3 (R13)	;Load to register
	BRZ	R3 If124:	;	if (first) {
;; row: :	REF	Tv125 Iv3 this; 	    first = false;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-17	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv125 Li110 ; 	    first = false;
	LDB	R3 Li110:	;	    first = false;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-17	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	WRITE	Li126  ; 	    cout << ' ';
	LDB	R0 Li126:	;
	TRP	#3	;	    cout << ' ';
;; row: :	JMP	El127  ; 	else cout << ',';
	JMP	El127:	;	else cout << ',';
;; row: If124:	WRITE	Li128  ; 	else cout << ',';
If124:	LDB	R0 Li128:	;
	TRP	#3	;	else cout << ',';
;; row: El127:	REF	Tv129 Iv18 Pa15; 	cout << node.root;
El127:	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-18	;
	STR	R13 (R10)	;
;; row: :	WRITE	Tv129  ; 	cout << node.root;
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


;; row: :	FUNC	St52  ; }
St52:   ADI   R0 #0 ;}
;; row: :	REF	Tv54 Iv2 this;     private iNode root;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	REF	Tv56 Iv3 this;     private bool first;
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
;; row: :	FRAME	this St130 ;     iNode(int key) {
;; Call function "St130:        iNode(int key) {"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-24          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke St130
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
;; row: :	CALL	St130  ;     iNode(int key) {
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
JMP     St130:
;; row: :	REF	Tv134 Iv18 this; 	root = key;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv134 Pa21 ; 	root = key;
	MOV	R10 RFP	;	root = key;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv135 Iv19 this; 	left = null;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-20	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv135 Li58 ; 	left = null;
	LDR	R3 Li58:	;	left = null;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-20	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv136 Iv20 this; 	right = null;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-24	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv136 Li58 ; 	right = null;
	LDR	R3 Li58:	;	right = null;
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


;; row: :	FUNC	St130  ; }
St130:   ADI   R0 #0 ;}
;; row: :	REF	Tv131 Iv18 this;     public int root;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	REF	Tv132 Iv19 this;     public iNode left;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	REF	Tv133 Iv20 this;     public iNode right;
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
;; row: :	FRAME	this St137 ;     Message() {
;; Call function "St137:        Message() {"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-21          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke St137
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
;; row: :	CALL	St137  ;     Message() {
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
JMP     St137:
;; row: :	REF	Tv142 Iv24 this;     	msg = new char[100];
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	MUL	Tv145 1 Li143;     	msg = new char[100];
	SUB	R4 R4	;    	msg = new char[100];
	ADI	R4 #1	;    	msg = new char[100];
	LDR	R3 Li143:	;
	MUL	R3 R4	;    	msg = new char[100];
	MOV	R10 RFP	;
	ADI	R10 #-17	;
	STR	R3 (R10)	;
;; row: :	NEW	Tv145 Tv144 ;     	msg = new char[100];
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
;; row: :	MOV	Tv142 Tv144 ;     	msg = new char[100];
	MOV	R10 RFP	;    	msg = new char[100];
	ADI	R10 #-13	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv146 Iv24 this; 	msg[0] = 'A';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-21	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv147 Li60 Tv146; 	msg[0] = 'A';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-21	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li60:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-22	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv147 Li148 ; 	msg[0] = 'A';
	LDB	R3 Li148:	;	msg[0] = 'A';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-22	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv149 Iv24 this; 	msg[1] = 'd';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-23	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv150 Li64 Tv149; 	msg[1] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-23	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li64:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-24	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv150 Li151 ; 	msg[1] = 'd';
	LDB	R3 Li151:	;	msg[1] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-24	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv152 Iv24 this; 	msg[2] = 'd';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-25	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv153 Li70 Tv152; 	msg[2] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-25	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li70:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-26	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv153 Li151 ; 	msg[2] = 'd';
	LDB	R3 Li151:	;	msg[2] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-26	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv155 Iv24 this; 	msg[3] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-27	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv156 Li154 Tv155; 	msg[3] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-27	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li154:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-28	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv156 Li157 ; 	msg[3] = 'e';
	LDB	R3 Li157:	;	msg[3] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-28	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv159 Iv24 this; 	msg[4] = 'd';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-29	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv160 Li158 Tv159; 	msg[4] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-29	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li158:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-30	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv160 Li151 ; 	msg[4] = 'd';
	LDB	R3 Li151:	;	msg[4] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-30	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv162 Iv24 this; 	msg[5] = ' ';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-31	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv163 Li161 Tv162; 	msg[5] = ' ';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-31	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li161:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-32	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv163 Li126 ; 	msg[5] = ' ';
	LDB	R3 Li126:	;	msg[5] = ' ';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-32	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv165 Iv24 this; 	msg[6] = 'E';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-33	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv166 Li164 Tv165; 	msg[6] = 'E';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-33	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li164:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-34	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv166 Li167 ; 	msg[6] = 'E';
	LDB	R3 Li167:	;	msg[6] = 'E';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-34	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv169 Iv24 this; 	msg[7] = 'l';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-35	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv170 Li168 Tv169; 	msg[7] = 'l';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-35	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li168:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-36	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv170 Li171 ; 	msg[7] = 'l';
	LDB	R3 Li171:	;	msg[7] = 'l';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-36	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv173 Iv24 this; 	msg[8] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-37	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv174 Li172 Tv173; 	msg[8] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-37	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li172:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-38	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv174 Li157 ; 	msg[8] = 'e';
	LDB	R3 Li157:	;	msg[8] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-38	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv176 Iv24 this; 	msg[9] = 'm';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-39	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv177 Li175 Tv176; 	msg[9] = 'm';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-39	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li175:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-40	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv177 Li178 ; 	msg[9] = 'm';
	LDB	R3 Li178:	;	msg[9] = 'm';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-40	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv180 Iv24 this; 	msg[10] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-41	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv181 Li179 Tv180; 	msg[10] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-41	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li179:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-42	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv181 Li157 ; 	msg[10] = 'e';
	LDB	R3 Li157:	;	msg[10] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-42	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv183 Iv24 this; 	msg[11] = 'n';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-43	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv184 Li182 Tv183; 	msg[11] = 'n';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-43	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li182:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-44	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv184 Li185 ; 	msg[11] = 'n';
	LDB	R3 Li185:	;	msg[11] = 'n';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-44	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv187 Iv24 this; 	msg[12] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-45	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv188 Li186 Tv187; 	msg[12] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-45	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li186:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-46	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv188 Li189 ; 	msg[12] = 't';
	LDB	R3 Li189:	;	msg[12] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-46	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv191 Iv24 this; 	msg[13] = ':';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-47	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv192 Li190 Tv191; 	msg[13] = ':';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-47	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li190:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-48	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv192 Li193 ; 	msg[13] = ':';
	LDB	R3 Li193:	;	msg[13] = ':';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-48	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv194 Iv25 this; 	i = 14;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-49	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv194 Li195 ; 	i = 14;
	LDR	R3 Li195:	;	i = 14;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-49	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv196 Iv25 this; 	msg[i] = 'D';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-53	;
	STR	R13 (R10)	;
;; row: :	REF	Tv197 Iv24 this; 	msg[i] = 'D';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-57	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv198 Tv196 Tv197; 	msg[i] = 'D';
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
;; row: :	MOV	Tv198 Li199 ; 	msg[i] = 'D';
	LDB	R3 Li199:	;	msg[i] = 'D';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-58	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv200 Iv25 this; 	msg[i+1] = 'u';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-59	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv201 Li64 Tv200; 	msg[i+1] = 'u';
	LDR	R4 Li64:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-59	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+1] = 'u';
	MOV	R10 RFP	;
	ADI	R10 #-63	;
	STR	R3 (R10)	;
;; row: :	REF	Tv202 Iv24 this; 	msg[i+1] = 'u';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-67	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv203 Tv201 Tv202; 	msg[i+1] = 'u';
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
;; row: :	MOV	Tv203 Li204 ; 	msg[i+1] = 'u';
	LDB	R3 Li204:	;	msg[i+1] = 'u';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-68	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv205 Iv25 this; 	msg[i+2] = 'p';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-69	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv206 Li70 Tv205; 	msg[i+2] = 'p';
	LDR	R4 Li70:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-69	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+2] = 'p';
	MOV	R10 RFP	;
	ADI	R10 #-73	;
	STR	R3 (R10)	;
;; row: :	REF	Tv207 Iv24 this; 	msg[i+2] = 'p';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-77	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv208 Tv206 Tv207; 	msg[i+2] = 'p';
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
;; row: :	MOV	Tv208 Li209 ; 	msg[i+2] = 'p';
	LDB	R3 Li209:	;	msg[i+2] = 'p';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-78	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv210 Iv25 this; 	msg[i+3] = 'l';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-79	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv211 Li154 Tv210; 	msg[i+3] = 'l';
	LDR	R4 Li154:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-79	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+3] = 'l';
	MOV	R10 RFP	;
	ADI	R10 #-83	;
	STR	R3 (R10)	;
;; row: :	REF	Tv212 Iv24 this; 	msg[i+3] = 'l';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-87	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv213 Tv211 Tv212; 	msg[i+3] = 'l';
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
;; row: :	MOV	Tv213 Li171 ; 	msg[i+3] = 'l';
	LDB	R3 Li171:	;	msg[i+3] = 'l';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-88	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv214 Iv25 this; 	msg[i+4] = 'i';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-89	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv215 Li158 Tv214; 	msg[i+4] = 'i';
	LDR	R4 Li158:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-89	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+4] = 'i';
	MOV	R10 RFP	;
	ADI	R10 #-93	;
	STR	R3 (R10)	;
;; row: :	REF	Tv216 Iv24 this; 	msg[i+4] = 'i';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-97	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv217 Tv215 Tv216; 	msg[i+4] = 'i';
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
;; row: :	MOV	Tv217 Li218 ; 	msg[i+4] = 'i';
	LDB	R3 Li218:	;	msg[i+4] = 'i';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-98	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv219 Iv25 this; 	msg[i+5] = 'c';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-99	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv220 Li161 Tv219; 	msg[i+5] = 'c';
	LDR	R4 Li161:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-99	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+5] = 'c';
	MOV	R10 RFP	;
	ADI	R10 #-103	;
	STR	R3 (R10)	;
;; row: :	REF	Tv221 Iv24 this; 	msg[i+5] = 'c';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-107	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv222 Tv220 Tv221; 	msg[i+5] = 'c';
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
;; row: :	MOV	Tv222 Li223 ; 	msg[i+5] = 'c';
	LDB	R3 Li223:	;	msg[i+5] = 'c';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-108	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv224 Iv25 this; 	msg[i+6] = 'a';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-109	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv225 Li164 Tv224; 	msg[i+6] = 'a';
	LDR	R4 Li164:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-109	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+6] = 'a';
	MOV	R10 RFP	;
	ADI	R10 #-113	;
	STR	R3 (R10)	;
;; row: :	REF	Tv226 Iv24 this; 	msg[i+6] = 'a';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-117	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv227 Tv225 Tv226; 	msg[i+6] = 'a';
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
;; row: :	MOV	Tv227 Li228 ; 	msg[i+6] = 'a';
	LDB	R3 Li228:	;	msg[i+6] = 'a';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-118	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv229 Iv25 this; 	msg[i+7] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-119	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv230 Li168 Tv229; 	msg[i+7] = 't';
	LDR	R4 Li168:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-119	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+7] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-123	;
	STR	R3 (R10)	;
;; row: :	REF	Tv231 Iv24 this; 	msg[i+7] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-127	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv232 Tv230 Tv231; 	msg[i+7] = 't';
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
;; row: :	MOV	Tv232 Li189 ; 	msg[i+7] = 't';
	LDB	R3 Li189:	;	msg[i+7] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-128	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv233 Iv25 this; 	msg[i+8] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-129	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv234 Li172 Tv233; 	msg[i+8] = 'e';
	LDR	R4 Li172:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-129	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+8] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-133	;
	STR	R3 (R10)	;
;; row: :	REF	Tv235 Iv24 this; 	msg[i+8] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-137	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv236 Tv234 Tv235; 	msg[i+8] = 'e';
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
;; row: :	MOV	Tv236 Li157 ; 	msg[i+8] = 'e';
	LDB	R3 Li157:	;	msg[i+8] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-138	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv237 Iv25 this; 	msg[i+9] = 'm';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-139	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv238 Li175 Tv237; 	msg[i+9] = 'm';
	LDR	R4 Li175:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-139	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+9] = 'm';
	MOV	R10 RFP	;
	ADI	R10 #-143	;
	STR	R3 (R10)	;
;; row: :	REF	Tv239 Iv24 this; 	msg[i+9] = 'm';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-147	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv240 Tv238 Tv239; 	msg[i+9] = 'm';
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
;; row: :	MOV	Tv240 Li178 ; 	msg[i+9] = 'm';
	LDB	R3 Li178:	;	msg[i+9] = 'm';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-148	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv242 Iv24 this; 	msg[24] = 'E';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-149	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv243 Li241 Tv242; 	msg[24] = 'E';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-149	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li241:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-150	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv243 Li167 ; 	msg[24] = 'E';
	LDB	R3 Li167:	;	msg[24] = 'E';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-150	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv245 Iv24 this; 	msg[25] = 'n';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-151	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv246 Li244 Tv245; 	msg[25] = 'n';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-151	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li244:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-152	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv246 Li185 ; 	msg[25] = 'n';
	LDB	R3 Li185:	;	msg[25] = 'n';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-152	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv248 Iv24 this; 	msg[26] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-153	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv249 Li247 Tv248; 	msg[26] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-153	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li247:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-154	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv249 Li189 ; 	msg[26] = 't';
	LDB	R3 Li189:	;	msg[26] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-154	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv251 Iv24 this; 	msg[27] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-155	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv252 Li250 Tv251; 	msg[27] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-155	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li250:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-156	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv252 Li157 ; 	msg[27] = 'e';
	LDB	R3 Li157:	;	msg[27] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-156	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv254 Iv24 this; 	msg[28] = 'r';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-157	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv255 Li253 Tv254; 	msg[28] = 'r';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-157	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li253:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-158	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv255 Li256 ; 	msg[28] = 'r';
	LDB	R3 Li256:	;	msg[28] = 'r';
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
;; row: Wh257:	LTE	Tv258 Pa28 Pa29; 	while (i <= end) {
Wh257:	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R4 (R10)	;
	CMP	R3 R4	;	while (i <= end) {
	BLT	R3 BT355:	
	BRZ	R3 BT355:	
	SUB	R3 R3	; false branch
	JMP	BF356:	
BT355:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF356:	MOV	R10 RFP	;
	ADI	R10 #-20	;
	STB	R3 (R10)	;
;; row: :	BF	Tv258 En259 ; 	while (i <= end) {
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	LDB	R3 (R10)	;
	BRZ	R3 En259:	;	while (i <= end) {
;; row: :	REF	Tv260 Iv24 this; 	    cout << msg[i];
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-21	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv261 Pa28 Tv260; 	    cout << msg[i];
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
;; row: :	WRITE	Tv261  ; 	    cout << msg[i];
	MOV	R10 RFP	;Load Address
	ADI	R10 #-22	;
	LDR	R13 (R10)	;
	LDB	R0 (R13)	;Load to register
	TRP	#3	;	    cout << msg[i];
;; row: :	ADD	Tv262 Li64 Pa28; 	    i = i + 1;
	LDR	R4 Li64:	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	ADD	R3 R4	;	    i = i + 1;
	MOV	R10 RFP	;
	ADI	R10 #-23	;
	STR	R3 (R10)	;
;; row: :	MOV	Pa28 Tv262 ; 	    i = i + 1;
	MOV	R10 RFP	;	    i = i + 1;
	ADI	R10 #-23	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R3 (R10)	;
;; row: :	JMP	Wh257  ;     }	
	JMP	Wh257:	;    }	
;; row: En259:	RTN	  ;     }	
;; return from function
;; test for underflow
En259:	MOV	RSP RFP	;     }	
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
;; row: :	PUSH	Li60  ; 	print(0, 13);
;; parameters on the stack (Li60)  ; 	print(0, 13);
	LDR	R1 Li60:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li190  ; 	print(0, 13);
;; parameters on the stack (Li190)  ; 	print(0, 13);
	LDR	R1 Li190:	;
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
;; row: :	PEEK	Tv263  ; 	print(0, 13);
	LDR	R11 (RSP)	;	print(0, 13);
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STR	R11 (R10)	;
;; row: :	WRITE	Pa31  ; 	cout << elm;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R0 (R10)	;
	TRP	#1	;	cout << elm;
;; row: :	WRITE	Li115  ; 	cout << '\n';
	LDB	R0 Li115:	;
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
;; row: :	REF	Tv264 Iv25 this; 	i = 14;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv264 Li195 ; 	i = 14;
	LDR	R3 Li195:	;	i = 14;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv265 Iv26 this; 	end = (i + 8);
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #8
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-20	;
	STR	R13 (R10)	;
;; row: :	REF	Tv266 Iv25 this; 	end = (i + 8);
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-24	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv267 Li172 Tv266; 	end = (i + 8);
	LDR	R4 Li172:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-24	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	end = (i + 8);
	MOV	R10 RFP	;
	ADI	R10 #-28	;
	STR	R3 (R10)	;
;; row: :	MOV	Tv265 Tv267 ; 	end = (i + 8);
	MOV	R10 RFP	;	end = (i + 8);
	ADI	R10 #-28	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-20	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv268 Iv25 this; 	print(i, end);
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-32	;
	STR	R13 (R10)	;
;; row: :	REF	Tv269 Iv26 this; 	print(i, end);
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
;; row: :	PUSH	Tv268  ; 	print(i, end);
;; parameters on the stack (Tv268)  ; 	print(i, end);
	MOV	R10 RFP	;Load Address
	ADI	R10 #-32	;
	LDR	R13 (R10)	;
	LDR	R1 (R13)	;Load to register
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv269  ; 	print(i, end);
;; parameters on the stack (Tv269)  ; 	print(i, end);
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
;; row: :	PEEK	Tv270  ; 	print(i, end);
	LDR	R11 (RSP)	;	print(i, end);
	MOV	R10 RFP	;
	ADI	R10 #-40	;
	STR	R11 (R10)	;
;; row: :	REF	Tv271 Iv24 this; 	cout << msg[5];
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-44	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv272 Li161 Tv271; 	cout << msg[5];
	MOV	R10 RFP	;Load Address
	ADI	R10 #-44	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li161:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-45	;
	STR	R13 (R10)	;
;; row: :	WRITE	Tv272  ; 	cout << msg[5];
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
;; row: :	PUSH	Li164  ; 	print(6, 13);
;; parameters on the stack (Li164)  ; 	print(6, 13);
	LDR	R1 Li164:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li190  ; 	print(6, 13);
;; parameters on the stack (Li190)  ; 	print(6, 13);
	LDR	R1 Li190:	;
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
;; row: :	PEEK	Tv273  ; 	print(6, 13);
	LDR	R11 (RSP)	;	print(6, 13);
	MOV	R10 RFP	;
	ADI	R10 #-46	;
	STR	R11 (R10)	;
;; row: :	WRITE	Pa33  ; 	cout << elm;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R0 (R10)	;
	TRP	#1	;	cout << elm;
;; row: :	WRITE	Li115  ; 	cout << '\n';
	LDB	R0 Li115:	;
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
;; row: :	PUSH	Li241  ; 	print(24, 28);
;; parameters on the stack (Li241)  ; 	print(24, 28);
	LDR	R1 Li241:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li253  ; 	print(24, 28);
;; parameters on the stack (Li253)  ; 	print(24, 28);
	LDR	R1 Li253:	;
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
;; row: :	PEEK	Tv274  ; 	print(24, 28);
	LDR	R11 (RSP)	;	print(24, 28);
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R11 (R10)	;
;; row: :	REF	Tv275 Iv25 this; 	i = 5;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv275 Li161 ; 	i = 5;
	LDR	R3 Li161:	;	i = 5;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv276 Iv25 this; 	print(i, i);
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-20	;
	STR	R13 (R10)	;
;; row: :	REF	Tv277 Iv25 this; 	print(i, i);
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
;; row: :	PUSH	Tv276  ; 	print(i, i);
;; parameters on the stack (Tv276)  ; 	print(i, i);
	MOV	R10 RFP	;Load Address
	ADI	R10 #-20	;
	LDR	R13 (R10)	;
	LDR	R1 (R13)	;Load to register
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv277  ; 	print(i, i);
;; parameters on the stack (Tv277)  ; 	print(i, i);
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
;; row: :	PEEK	Tv278  ; 	print(i, i);
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
;; row: :	PUSH	Li164  ; 	print(6, 13);
;; parameters on the stack (Li164)  ; 	print(6, 13);
	LDR	R1 Li164:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li190  ; 	print(6, 13);
;; parameters on the stack (Li190)  ; 	print(6, 13);
	LDR	R1 Li190:	;
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
;; row: :	PEEK	Tv279  ; 	print(6, 13);
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


;; row: :	FUNC	St137  ; }
St137:   ADI   R0 #0 ;}
;; row: :	REF	Tv139 Iv24 this;     private char msg[];
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	REF	Tv140 Iv25 this;     private int i;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-13	;
	STR	R13 (R10)	;
;; row: :	REF	Tv141 Iv26 this;     private int end;
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


;; row: :	FUNC	Co42  ;     Syntax(int j, char d) {
Co42:   ADI   R0 #0 ;    Syntax(int j, char d) {
;; row: :	FRAME	this St280 ;     Syntax(int j, char d) {
;; Call function "St280:        Syntax(int j, char d) {"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-18          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke St280
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
;; row: :	CALL	St280  ;     Syntax(int j, char d) {
;; local varibales on the stack    ;     Syntax(int j, char d) {
;; Temp variables on the stack
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-18
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     St280:
;; row: :	REF	Tv284 Iv37 this;          i = j;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-17	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv284 Pa40 ;          i = j;
	MOV	R10 RFP	;         i = j;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-17	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv285 Iv38 this;          c = d;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-21	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv285 Pa41 ;          c = d;
	MOV	R10 RFP	;         c = d;
	ADI	R10 #-16	;
	LDB	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-21	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
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


;; row: :	FUNC	Me43  ;       public void checkit() {
Me43:   ADI   R0 #0 ;      public void checkit() {
;; row: :	MUL	Tv288 1 Li286;          char cc[] = new char[1000];
	SUB	R4 R4	;         char cc[] = new char[1000];
	ADI	R4 #1	;         char cc[] = new char[1000];
	LDR	R3 Li286:	;
	MUL	R3 R4	;         char cc[] = new char[1000];
	MOV	R10 RFP	;
	ADI	R10 #-28	;
	STR	R3 (R10)	;
;; row: :	NEW	Tv288 Tv287 ;          char cc[] = new char[1000];
	MOV	R10 RFP	;
	ADI	R10 #-28	;
	LDR	R3 (R10)	;
;; Test for heap overflow
MOV     R10 R9
ADD     R10 R3
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADD     R9 R3
	MOV	R10 RFP	;
	ADI	R10 #-24	;
	STR	R11 (R10)	;
;; row: :	MOV	Lv44 Tv287 ;          char cc[] = new char[1000];
	MOV	R10 RFP	;         char cc[] = new char[1000];
	ADI	R10 #-24	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R3 (R10)	;
;; row: :	MUL	Tv291 4 Li289; 	 int  ii[] = new int[512];
	SUB	R4 R4	;	 int  ii[] = new int[512];
	ADI	R4 #4	;	 int  ii[] = new int[512];
	LDR	R3 Li289:	;
	MUL	R3 R4	;	 int  ii[] = new int[512];
	MOV	R10 RFP	;
	ADI	R10 #-36	;
	STR	R3 (R10)	;
;; row: :	NEW	Tv291 Tv290 ; 	 int  ii[] = new int[512];
	MOV	R10 RFP	;
	ADI	R10 #-36	;
	LDR	R3 (R10)	;
;; Test for heap overflow
MOV     R10 R9
ADD     R10 R3
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADD     R9 R3
	MOV	R10 RFP	;
	ADI	R10 #-32	;
	STR	R11 (R10)	;
;; row: :	MOV	Lv45 Tv290 ; 	 int  ii[] = new int[512];
	MOV	R10 RFP	;	 int  ii[] = new int[512];
	ADI	R10 #-32	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STR	R3 (R10)	;
;; row: :	MUL	Tv295 4 Li293; 	 Syntax ss[] = new Syntax[256];
	SUB	R4 R4	;	 Syntax ss[] = new Syntax[256];
	ADI	R4 #4	;	 Syntax ss[] = new Syntax[256];
	LDR	R3 Li293:	;
	MUL	R3 R4	;	 Syntax ss[] = new Syntax[256];
	MOV	R10 RFP	;
	ADI	R10 #-44	;
	STR	R3 (R10)	;
;; row: :	NEW	Tv295 Tv294 ; 	 Syntax ss[] = new Syntax[256];
	MOV	R10 RFP	;
	ADI	R10 #-44	;
	LDR	R3 (R10)	;
;; Test for heap overflow
MOV     R10 R9
ADD     R10 R3
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADD     R9 R3
	MOV	R10 RFP	;
	ADI	R10 #-40	;
	STR	R11 (R10)	;
;; row: :	MOV	Lv46 Tv294 ; 	 Syntax ss[] = new Syntax[256];
	MOV	R10 RFP	;	 Syntax ss[] = new Syntax[256];
	ADI	R10 #-40	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	STR	R3 (R10)	;
;; row: :	AEF	Tv296 Li64 Lv44; 	 cc[1] = cc[2]; // yes
	MOV	R10 RFP	;Load Address
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	LDR	R14 Li64:	;
	SUB	R12 R12
	ADI	R12 #4
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-48	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv297 Li70 Lv44; 	 cc[1] = cc[2]; // yes
	MOV	R10 RFP	;Load Address
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	LDR	R14 Li70:	;
	SUB	R12 R12
	ADI	R12 #4
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-49	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv296 Tv297 ; 	 cc[1] = cc[2]; // yes
	MOV	R10 RFP	;	 cc[1] = cc[2]; // yes
	ADI	R10 #-49	;
	LDR	R13 (R10)	;
	LDB	R3 (R13)	;Load to register
	MOV	R10 RFP	;Load Address
	ADI	R10 #-48	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	AEF	Tv298 Li179 Lv44; 	 cc[10] = c;     // yes
	MOV	R10 RFP	;Load Address
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	LDR	R14 Li179:	;
	SUB	R12 R12
	ADI	R12 #4
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-50	;
	STR	R13 (R10)	;
;; row: :	REF	Tv299 Iv38 this; 	 cc[10] = c;     // yes
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-51	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv298 Tv299 ; 	 cc[10] = c;     // yes
	MOV	R10 RFP	;	 cc[10] = c;     // yes
	ADI	R10 #-51	;
	LDR	R13 (R10)	;
	LDB	R3 (R13)	;Load to register
	MOV	R10 RFP	;Load Address
	ADI	R10 #-50	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	AEF	Tv300 Li158 Lv45; 	 ii[4] = 5 + i;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	LDR	R14 Li158:	;
	SUB	R12 R12
	ADI	R12 #4
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-52	;
	STR	R13 (R10)	;
;; row: :	REF	Tv301 Iv37 this; 	 ii[4] = 5 + i;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-56	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv302 Tv301 Li161; 	 ii[4] = 5 + i;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-56	;
	LDR	R13 (R10)	;
	LDR	R4 (R13)	;Load to register
	LDR	R3 Li161:	;
	ADD	R3 R4	;	 ii[4] = 5 + i;
	MOV	R10 RFP	;
	ADI	R10 #-60	;
	STR	R3 (R10)	;
;; row: :	MOV	Tv300 Tv302 ; 	 ii[4] = 5 + i;
	MOV	R10 RFP	;	 ii[4] = 5 + i;
	ADI	R10 #-60	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-52	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	AEF	Tv304 Li303 Lv45; 	 ii[5000] = 5 + i; // yes
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	LDR	R14 Li303:	;
	SUB	R12 R12
	ADI	R12 #4
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-64	;
	STR	R13 (R10)	;
;; row: :	REF	Tv305 Iv37 this; 	 ii[5000] = 5 + i; // yes
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-68	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv306 Tv305 Li161; 	 ii[5000] = 5 + i; // yes
	MOV	R10 RFP	;Load Address
	ADI	R10 #-68	;
	LDR	R13 (R10)	;
	LDR	R4 (R13)	;Load to register
	LDR	R3 Li161:	;
	ADD	R3 R4	;	 ii[5000] = 5 + i; // yes
	MOV	R10 RFP	;
	ADI	R10 #-72	;
	STR	R3 (R10)	;
;; row: :	MOV	Tv304 Tv306 ; 	 ii[5000] = 5 + i; // yes
	MOV	R10 RFP	;	 ii[5000] = 5 + i; // yes
	ADI	R10 #-72	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-64	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv307 Iv37 this; 	 i = ii[0];  // yes
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-76	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv308 Li60 Lv45; 	 i = ii[0];  // yes
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	LDR	R14 Li60:	;
	SUB	R12 R12
	ADI	R12 #4
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-80	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv307 Tv308 ; 	 i = ii[0];  // yes
	MOV	R10 RFP	;	 i = ii[0];  // yes
	ADI	R10 #-80	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	MOV	R10 RFP	;Load Address
	ADI	R10 #-76	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	AEF	Tv309 Li60 Lv46; 	 ss[0] = new Syntax(7, 'c');
	MOV	R10 RFP	;Load Address
	ADI	R10 #-20	;
	LDR	R13 (R10)	;
	LDR	R14 Li60:	;
	SUB	R12 R12
	ADI	R12 #4
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-84	;
	STR	R13 (R10)	;
;; row: :	NEWI	Cl36 Tv310 ; 	 ss[0] = new Syntax(7, 'c');
;; Test for heap overflow
	MOV     R10 R9
ADI     R10 #6
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADI     R9 #6
	MOV	R10 RFP	;
	ADI	R10 #-88	;
	STR	R11 (R10)	;
;; row: :	FRAME	Tv310 Co42 ; 	 ss[0] = new Syntax(7, 'c');
;; Call function "Co42:    	 ss[0] = new Syntax(7, 'c');"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-22          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Co42
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-88	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li168  ; 	 ss[0] = new Syntax(7, 'c');
;; parameters on the stack (Li168)  ; 	 ss[0] = new Syntax(7, 'c');
	LDR	R1 Li168:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li223  ; 	 ss[0] = new Syntax(7, 'c');
;; parameters on the stack (Li223)  ; 	 ss[0] = new Syntax(7, 'c');
	LDB	R1 Li223:	;
STB     R1 (RSP)
ADI     RSP #-1
;; row: :	CALL	Co42  ; 	 ss[0] = new Syntax(7, 'c');
;; local varibales on the stack    ; 	 ss[0] = new Syntax(7, 'c');
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-1
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-22
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Co42:
;; row: :	PEEK	Tv310  ; 	 ss[0] = new Syntax(7, 'c');
	LDR	R11 (RSP)	;	 ss[0] = new Syntax(7, 'c');
	MOV	R10 RFP	;
	ADI	R10 #-88	;
	STR	R11 (R10)	;
;; row: :	MOV	Tv309 Tv310 ; 	 ss[0] = new Syntax(7, 'c');
	MOV	R10 RFP	;	 ss[0] = new Syntax(7, 'c');
	ADI	R10 #-88	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-84	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv311 Iv37 this; 	 ss[i] = ss[i+1]; // yes
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-92	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv312 Tv311 Lv46; 	 ss[i] = ss[i+1]; // yes
	MOV	R10 RFP	;Load Address
	ADI	R10 #-20	;
	LDR	R13 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-92	;
	LDR	R13 (R10)	;
	LDR	R14 (R13)	;Load to register
	SUB	R12 R12
	ADI	R12 #4
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-96	;
	STR	R13 (R10)	;
;; row: :	REF	Tv313 Iv37 this; 	 ss[i] = ss[i+1]; // yes
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-100	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv314 Li64 Tv313; 	 ss[i] = ss[i+1]; // yes
	LDR	R4 Li64:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-100	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	 ss[i] = ss[i+1]; // yes
	MOV	R10 RFP	;
	ADI	R10 #-104	;
	STR	R3 (R10)	;
;; row: :	AEF	Tv315 Tv314 Lv46; 	 ss[i] = ss[i+1]; // yes
	MOV	R10 RFP	;Load Address
	ADI	R10 #-20	;
	LDR	R13 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-104	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #4
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-108	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv312 Tv315 ; 	 ss[i] = ss[i+1]; // yes
	MOV	R10 RFP	;	 ss[i] = ss[i+1]; // yes
	ADI	R10 #-108	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	MOV	R10 RFP	;Load Address
	ADI	R10 #-96	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv316 Iv37 this; 	 ss[i+7/3] = new Syntax(i, c); // yes
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-112	;
	STR	R13 (R10)	;
;; row: :	DIV	Tv317 Li154 Li168; 	 ss[i+7/3] = new Syntax(i, c); // yes
	LDR	R4 Li154:	;
	LDR	R3 Li168:	;
	DIV	R3 R4	;	 ss[i+7/3] = new Syntax(i, c); // yes
	MOV	R10 RFP	;
	ADI	R10 #-116	;
	STR	R3 (R10)	;
;; row: :	ADD	Tv318 Tv317 Tv316; 	 ss[i+7/3] = new Syntax(i, c); // yes
	MOV	R10 RFP	;
	ADI	R10 #-116	;
	LDR	R4 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-112	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	 ss[i+7/3] = new Syntax(i, c); // yes
	MOV	R10 RFP	;
	ADI	R10 #-120	;
	STR	R3 (R10)	;
;; row: :	AEF	Tv319 Tv318 Lv46; 	 ss[i+7/3] = new Syntax(i, c); // yes
	MOV	R10 RFP	;Load Address
	ADI	R10 #-20	;
	LDR	R13 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-120	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #4
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-124	;
	STR	R13 (R10)	;
;; row: :	REF	Tv320 Iv37 this; 	 ss[i+7/3] = new Syntax(i, c); // yes
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-128	;
	STR	R13 (R10)	;
;; row: :	REF	Tv321 Iv38 this; 	 ss[i+7/3] = new Syntax(i, c); // yes
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-132	;
	STR	R13 (R10)	;
;; row: :	NEWI	Cl36 Tv322 ; 	 ss[i+7/3] = new Syntax(i, c); // yes
;; Test for heap overflow
	MOV     R10 R9
ADI     R10 #6
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADI     R9 #6
	MOV	R10 RFP	;
	ADI	R10 #-133	;
	STR	R11 (R10)	;
;; row: :	FRAME	Tv322 Co42 ; 	 ss[i+7/3] = new Syntax(i, c); // yes
;; Call function "Co42:    	 ss[i+7/3] = new Syntax(i, c); // yes"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-22          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Co42
MOV     R10 RFP
MOV     R15 RSP
ADI     RSP #-4
STR     R10 (RSP)
ADI     RSP #-4
;; this
	MOV	R10 RFP	;
	ADI	R10 #-133	;
	LDR	R1 (R10)	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv320  ; 	 ss[i+7/3] = new Syntax(i, c); // yes
;; parameters on the stack (Tv320)  ; 	 ss[i+7/3] = new Syntax(i, c); // yes
	MOV	R10 RFP	;Load Address
	ADI	R10 #-128	;
	LDR	R13 (R10)	;
	LDR	R1 (R13)	;Load to register
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Tv321  ; 	 ss[i+7/3] = new Syntax(i, c); // yes
;; parameters on the stack (Tv321)  ; 	 ss[i+7/3] = new Syntax(i, c); // yes
	MOV	R10 RFP	;Load Address
	ADI	R10 #-132	;
	LDR	R13 (R10)	;
	LDB	R1 (R13)	;Load to register
STB     R1 (RSP)
ADI     RSP #-1
;; row: :	CALL	Co42  ; 	 ss[i+7/3] = new Syntax(i, c); // yes
;; local varibales on the stack    ; 	 ss[i+7/3] = new Syntax(i, c); // yes
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-1
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-22
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Co42:
;; row: :	PEEK	Tv322  ; 	 ss[i+7/3] = new Syntax(i, c); // yes
	LDR	R11 (RSP)	;	 ss[i+7/3] = new Syntax(i, c); // yes
	MOV	R10 RFP	;
	ADI	R10 #-133	;
	STR	R11 (R10)	;
;; row: :	MOV	Tv319 Tv322 ; 	 ss[i+7/3] = new Syntax(i, c); // yes
	MOV	R10 RFP	;	 ss[i+7/3] = new Syntax(i, c); // yes
	ADI	R10 #-133	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-124	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
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
JMR     R15             ; go back "      } "


;; row: :	FUNC	Me48  ;       public int which(int i) {
Me48:   ADI   R0 #0 ;      public int which(int i) {
;; row: :	MUL	Tv323 Pa47 Pa47;          i = i * i;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R4 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	MUL	R3 R4	;         i = i * i;
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STR	R3 (R10)	;
;; row: :	MOV	Pa47 Tv323 ;          i = i * i;
	MOV	R10 RFP	;         i = i * i;
	ADI	R10 #-16	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R3 (R10)	;
;; row: :	RETURN	Pa47  ; 	 return i;
;; return from function
;; test for underflow
MOV     RSP RFP
LDR     R15 (RSP)
MOV     R10 RSP
CMP     R10 RSB
BGT     R10 UDRFLW:     ; oopsy underflow problem
;; store the return value
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R0 (R10)	;
STR     R0 (RSP)        ; R0 is whatever the value is for return
;; set previous frame to current frame and return
MOV     R11 RFP
ADI     R11 #-4         ; now pointing at PFP
LDR     RFP (R11)       ; make FP = PFP
JMR     R15             ; go back "	 return i;"


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


;; row: :	FUNC	St280  ; }
St280:   ADI   R0 #0 ;}
;; row: :	REF	Tv281 Iv37 this;       private int i = 7;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv281 Li168 ;       private int i = 7;
	LDR	R3 Li168:	;      private int i = 7;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv282 Iv38 this;       private char c = 'a';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv282 Li228 ;       private char c = 'a';
	LDB	R3 Li228:	;      private char c = 'a';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-16	;
	LDR	R13 (R10)	;
	STB	R3 (R13)	;Save from Register
;; row: :	REF	Tv283 Iv39 this;       private bool b = false;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #5
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-17	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv283 Li110 ;       private bool b = false;
	LDB	R3 Li110:	;      private bool b = false;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-17	;
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
;; row: :	NEWI	Cl23 Tv326 ;     Message msg = new Message();
;; Test for heap overflow
	MOV     R10 R9
ADI     R10 #12
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADI     R9 #12
	MOV	R10 RFP	;
	ADI	R10 #-24	;
	STR	R11 (R10)	;
;; row: :	FRAME	Tv326 Co27 ;     Message msg = new Message();
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
	ADI	R10 #-24	;
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
;; row: :	PEEK	Tv326  ;     Message msg = new Message();
	LDR	R11 (RSP)	;    Message msg = new Message();
	MOV	R10 RFP	;
	ADI	R10 #-24	;
	STR	R11 (R10)	;
;; row: :	MOV	Lv51 Tv326 ;     Message msg = new Message();
	MOV	R10 RFP	;    Message msg = new Message();
	ADI	R10 #-24	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	STR	R3 (R10)	;
;; row: :	NEWI	Cl1 Tv327 ;     tree = new iTree();
;; Test for heap overflow
	MOV     R10 R9
ADI     R10 #5
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADI     R9 #5
	MOV	R10 RFP	;
	ADI	R10 #-28	;
	STR	R11 (R10)	;
;; row: :	FRAME	Tv327 Co4 ;     tree = new iTree();
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
	ADI	R10 #-28	;
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
;; row: :	PEEK	Tv327  ;     tree = new iTree();
	LDR	R11 (RSP)	;    tree = new iTree();
	MOV	R10 RFP	;
	ADI	R10 #-28	;
	STR	R11 (R10)	;
;; row: :	MOV	Lv50 Tv327 ;     tree = new iTree();
	MOV	R10 RFP	;    tree = new iTree();
	ADI	R10 #-28	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STR	R3 (R10)	;
;; row: :	FRAME	Lv51 Me35 ;     msg.msg3();
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
;; row: :	PEEK	Tv328  ;     msg.msg3();
	LDR	R11 (RSP)	;    msg.msg3();
	MOV	R10 RFP	;
	ADI	R10 #-32	;
	STR	R11 (R10)	;
;; row: :	READ	Lv49  ;     cin >> key;
	TRP	#2	;    cin >> key;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R0 (R10)	;
;; row: :	WRITE	Li115  ;     cout << '\n';
	LDB	R0 Li115:	;
	TRP	#3	;    cout << '\n';
;; row: Wh329:	NEQ	Tv330 Lv49 Li60;     while (key != 0) {
Wh329:	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R3 (R10)	;
	LDR	R4 Li60:	;
	CMP	R3 R4	;    while (key != 0) {
	BNZ	R3 BT357:	
	SUB	R3 R3	; false branch
	JMP	BF358:	
BT357:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF358:	MOV	R10 RFP	;
	ADI	R10 #-36	;
	STB	R3 (R10)	;
;; row: :	BF	Tv330 En331 ;     while (key != 0) {
	MOV	R10 RFP	;
	ADI	R10 #-36	;
	LDB	R3 (R10)	;
	BRZ	R3 En331:	;    while (key != 0) {
;; row: :	FRAME	Lv50 Me8 ; 	if (tree.add(key)) {
;; Call function "Me8:    	if (tree.add(key)) {"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-49          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
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
;; row: :	PUSH	Lv49  ; 	if (tree.add(key)) {
;; parameters on the stack (Lv49)  ; 	if (tree.add(key)) {
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
ADI     RSP #-4
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-49
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me8:
;; row: :	PEEK	Tv332  ; 	if (tree.add(key)) {
	LDR	R11 (RSP)	;	if (tree.add(key)) {
	MOV	R10 RFP	;
	ADI	R10 #-37	;
	STR	R11 (R10)	;
;; row: :	BF	Tv332 If333 ; 	if (tree.add(key)) {
	MOV	R10 RFP	;
	ADI	R10 #-37	;
	LDR	R3 (R10)	;
	BRZ	R3 If333:	;	if (tree.add(key)) {
;; row: :	FRAME	Lv51 Me32 ; 	    msg.msg1(key);
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
;; row: :	PUSH	Lv49  ; 	    msg.msg1(key);
;; parameters on the stack (Lv49)  ; 	    msg.msg1(key);
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
;; row: :	PEEK	Tv334  ; 	    msg.msg1(key);
	LDR	R11 (RSP)	;	    msg.msg1(key);
	MOV	R10 RFP	;
	ADI	R10 #-41	;
	STR	R11 (R10)	;
;; row: :	FRAME	Lv50 Me12 ; 	    tree.print();
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
;; row: :	PEEK	Tv335  ; 	    tree.print();
	LDR	R11 (RSP)	;	    tree.print();
	MOV	R10 RFP	;
	ADI	R10 #-45	;
	STR	R11 (R10)	;
;; row: :	JMP	El336  ; 	else 
	JMP	El336:	;	else 
;; row: If333:	FRAME	Lv51 Me34 ; 	    msg.msg2(key);
;; Call function "Me34:    	    msg.msg2(key);"
;; Test for overflow
If333:   MOV     R10 RSP
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
;; row: :	PUSH	Lv49  ; 	    msg.msg2(key);
;; parameters on the stack (Lv49)  ; 	    msg.msg2(key);
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
;; row: :	PEEK	Tv337  ; 	    msg.msg2(key);
	LDR	R11 (RSP)	;	    msg.msg2(key);
	MOV	R10 RFP	;
	ADI	R10 #-49	;
	STR	R11 (R10)	;
;; row: El336:	FRAME	Lv51 Me35 ; 	msg.msg3();
;; Call function "Me35:    	msg.msg3();"
;; Test for overflow
El336:   MOV     R10 RSP
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
;; row: :	PEEK	Tv338  ; 	msg.msg3();
	LDR	R11 (RSP)	;	msg.msg3();
	MOV	R10 RFP	;
	ADI	R10 #-53	;
	STR	R11 (R10)	;
;; row: :	READ	Lv49  ; 	cin >> key;
	TRP	#2	;	cin >> key;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R0 (R10)	;
;; row: :	WRITE	Li115  ; 	cout << '\n';
	LDB	R0 Li115:	;
	TRP	#3	;	cout << '\n';
;; row: :	JMP	Wh329  ; }
	JMP	Wh329:	;}
;; row: En331:	RTN	  ; }
;; return from function
;; test for underflow
En331:	MOV	RSP RFP	; }
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
