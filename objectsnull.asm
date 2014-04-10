LDA     R9 FREE:
;; Call function "MAIN:"
;; Test for overflow
MOV     R10 RSP
ADI     R10 #-25          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
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
;; Temp variables on the stack
ADI     RSP #-4
ADI     RSP #-1
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
Li13:	.INT	0
Li19:	.BYT	'l'
Li14:	.BYT	'\n'
Li21:	.BYT	's'
Li11:	.INT	1
Li17:	.BYT	'n'
Li18:	.BYT	'u'
;; functions
Co3:   ADI   R0 #0 ;    Obj() {
;; Call function "St6:        Obj() {"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-16          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke St6
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
;; local varibales on the stack    ;     Obj() {
;; Temp variables on the stack
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-16
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     St6:
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
	MOV	R10 RFP	;	self = this;
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	MOV	R10 RFP	;Load Address
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
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


St6:   ADI   R0 #0 ;}
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
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


MAIN:   ADI   R0 #0 ;void main() {
	LDR	R4 Li11:	;
	LDR	R3 Li11:	;
	ADD	R3 R4	;    int i = 1+1;
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	STR	R3 (R10)	;
	MOV	R10 RFP	;    int i = 1+1;
	ADI	R10 #-20	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R3 (R10)	;
	LDR	R3 Li13:	;    Obj o = null;
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	LDR	R0 (R10)	;
	TRP	#1	;    cout << i;
	LDB	R0 Li14:	;
	TRP	#3	;    cout << '\n';
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	LDR	R3 (R10)	;
	LDR	R4 Li13:	;
	CMP	R3 R4	;    if(o == null) {
	BRZ	R3 BT22:	
	SUB	R3 R3	; false branch
	JMP	BF23:	
BT22:	SUB	R3 R3	;True Branch
	ADI	R3 #1	;True Branch
BF23:	MOV	R10 RFP	;
	ADI	R10 #-24	;
	STB	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-24	;
	LDB	R3 (R10)	;
	BRZ	R3 If16:	;    if(o == null) {
	LDB	R0 Li17:	;
	TRP	#3	;	cout << 'n';
	LDB	R0 Li18:	;
	TRP	#3	;	cout << 'u';
	LDB	R0 Li19:	;
	TRP	#3	;	cout << 'l';
	LDB	R0 Li19:	;
	TRP	#3	;	cout << 'l';
	LDB	R0 Li14:	;
	TRP	#3	;	cout << '\n';
	JMP	El20:	;    } else {
If16:	LDB	R0 Li21:	;
	TRP	#3	;	cout << 's';
	LDB	R0 Li14:	;
	TRP	#3	;	cout << '\n';
;; return from function
;; test for underflow
El20:	MOV	RSP RFP	; }
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
