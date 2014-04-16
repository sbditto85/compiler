LDA     R9 FREE:
;; Call function "MAIN:"
;; Test for overflow
MOV     R10 RSP
ADI     R10 #-28          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
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
	ADI	RSP #-32
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
Li64:	.BYT	'n'
Li72:	.BYT	':'
Li120:	.INT	24
Li46:	.BYT	'E'
Li135:	.BYT	'r'
Li97:	.BYT	'i'
Li54:	.INT	9
Li102:	.BYT	'c'
Li58:	.INT	10
Li61:	.INT	11
Li47:	.INT	7
Li69:	.INT	13
Li21:	.INT	0
Li36:	.INT	4
Li43:	.INT	6
Li42:	.BYT	32
Li57:	.BYT	'm'
Li68:	.BYT	't'
Li18:	.INT	100
Li123:	.INT	25
Li24:	.BYT	'A'
Li29:	.INT	2
Li83:	.BYT	'u'
Li132:	.INT	28
Li39:	.INT	5
Li50:	.BYT	'l'
Li78:	.BYT	'D'
Li141:	.INT	22
Li35:	.BYT	'e'
Li32:	.INT	3
Li74:	.INT	14
Li126:	.INT	26
Li129:	.INT	27
Li51:	.INT	8
Li65:	.INT	12
Li25:	.INT	1
Li107:	.BYT	'a'
Li28:	.BYT	'd'
Li88:	.BYT	'p'
;; functions
;; row: :	FUNC	Co5  ;     Message() {
Co5:   ADI   R0 #0 ;    Message() {
;; row: :	FRAME	this St11 ;     Message() {
;; Call function "St11:        Message() {"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-24          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke St11
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
;; row: :	CALL	St11  ;     Message() {
;; local varibales on the stack    ;     Message() {
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-28
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     St11:
;; row: :	REF	Tv17 Iv2 this;     	msg = new char[100];
        TRP     #99
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	MUL	Tv20 1 Li18;     	msg = new char[100];
	SUB	R4 R4	;    	msg = new char[100];
	ADI	R4 #1	;    	msg = new char[100];
	LDR	R3 Li18:	;
	MUL	R3 R4	;    	msg = new char[100];
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	STR	R3 (R10)	;
;; row: :	NEW	Tv20 Tv19 ;     	msg = new char[100];
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	LDR	R3 (R10)	;
;; Test for heap overflow
MOV     R10 R9
ADD     R10 R3
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADD     R9 R3
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STR	R11 (R10)	;
;; row: :	MOV	Tv17 Tv19 ;     	msg = new char[100];
	MOV	R10 RFP	;    	msg = new char[100];
	ADI	R10 #-16	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-12	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv22 Iv2 this; 	msg[0] = 'A';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-24	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv23 Li21 Tv22; 	msg[0] = 'A';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-24	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li21:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-28	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv23 Li24 ; 	msg[0] = 'A';
	LDB	R3 Li24:	;	msg[0] = 'A';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-28	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv26 Iv2 this; 	msg[1] = 'd';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-32	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv27 Li25 Tv26; 	msg[1] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-32	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li25:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-36	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv27 Li28 ; 	msg[1] = 'd';
	LDB	R3 Li28:	;	msg[1] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-36	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv30 Iv2 this; 	msg[2] = 'd';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-40	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv31 Li29 Tv30; 	msg[2] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-40	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li29:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-44	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv31 Li28 ; 	msg[2] = 'd';
	LDB	R3 Li28:	;	msg[2] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-44	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv33 Iv2 this; 	msg[3] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-48	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv34 Li32 Tv33; 	msg[3] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-48	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li32:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-52	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv34 Li35 ; 	msg[3] = 'e';
	LDB	R3 Li35:	;	msg[3] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-52	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv37 Iv2 this; 	msg[4] = 'd';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-56	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv38 Li36 Tv37; 	msg[4] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-56	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li36:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-60	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv38 Li28 ; 	msg[4] = 'd';
	LDB	R3 Li28:	;	msg[4] = 'd';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-60	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv40 Iv2 this; 	msg[5] = ' ';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-64	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv41 Li39 Tv40; 	msg[5] = ' ';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-64	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li39:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-68	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv41 Li42 ; 	msg[5] = ' ';
	LDB	R3 Li42:	;	msg[5] = ' ';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-68	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv44 Iv2 this; 	msg[6] = 'E';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-72	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv45 Li43 Tv44; 	msg[6] = 'E';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-72	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li43:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-76	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv45 Li46 ; 	msg[6] = 'E';
	LDB	R3 Li46:	;	msg[6] = 'E';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-76	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv48 Iv2 this; 	msg[7] = 'l';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-80	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv49 Li47 Tv48; 	msg[7] = 'l';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-80	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li47:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-84	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv49 Li50 ; 	msg[7] = 'l';
	LDB	R3 Li50:	;	msg[7] = 'l';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-84	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv52 Iv2 this; 	msg[8] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-88	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv53 Li51 Tv52; 	msg[8] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-88	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li51:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-92	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv53 Li35 ; 	msg[8] = 'e';
	LDB	R3 Li35:	;	msg[8] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-92	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv55 Iv2 this; 	msg[9] = 'm';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-96	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv56 Li54 Tv55; 	msg[9] = 'm';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-96	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li54:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-100	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv56 Li57 ; 	msg[9] = 'm';
	LDB	R3 Li57:	;	msg[9] = 'm';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-100	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv59 Iv2 this; 	msg[10] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-104	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv60 Li58 Tv59; 	msg[10] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-104	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li58:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-108	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv60 Li35 ; 	msg[10] = 'e';
	LDB	R3 Li35:	;	msg[10] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-108	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv62 Iv2 this; 	msg[11] = 'n';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-112	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv63 Li61 Tv62; 	msg[11] = 'n';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-112	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li61:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-116	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv63 Li64 ; 	msg[11] = 'n';
	LDB	R3 Li64:	;	msg[11] = 'n';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-116	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv66 Iv2 this; 	msg[12] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-120	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv67 Li65 Tv66; 	msg[12] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-120	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li65:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-124	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv67 Li68 ; 	msg[12] = 't';
	LDB	R3 Li68:	;	msg[12] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-124	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv70 Iv2 this; 	msg[13] = ':';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-128	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv71 Li69 Tv70; 	msg[13] = ':';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-128	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li69:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-132	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv71 Li72 ; 	msg[13] = ':';
	LDB	R3 Li72:	;	msg[13] = ':';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-132	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv73 Iv3 this; 	i = 14;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-136	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv73 Li74 ; 	i = 14;
	LDR	R3 Li74:	;	i = 14;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-136	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv75 Iv3 this; 	msg[i] = 'D';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-140	;
	STR	R13 (R10)	;
;; row: :	REF	Tv76 Iv2 this; 	msg[i] = 'D';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-144	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv77 Tv75 Tv76; 	msg[i] = 'D';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-144	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;Load Address
	ADI	R10 #-140	;
	LDR	R13 (R10)	;
	LDR	R14 (R13)	;Load to register
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-148	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv77 Li78 ; 	msg[i] = 'D';
	LDB	R3 Li78:	;	msg[i] = 'D';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-148	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv79 Iv3 this; 	msg[i+1] = 'u';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-152	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv80 Li25 Tv79; 	msg[i+1] = 'u';
	LDR	R4 Li25:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-152	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+1] = 'u';
	MOV	R10 RFP	;
	ADI	R10 #-156	;
	STR	R3 (R10)	;
;; row: :	REF	Tv81 Iv2 this; 	msg[i+1] = 'u';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-160	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv82 Tv80 Tv81; 	msg[i+1] = 'u';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-160	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-156	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-164	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv82 Li83 ; 	msg[i+1] = 'u';
	LDB	R3 Li83:	;	msg[i+1] = 'u';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-164	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv84 Iv3 this; 	msg[i+2] = 'p';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-168	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv85 Li29 Tv84; 	msg[i+2] = 'p';
	LDR	R4 Li29:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-168	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+2] = 'p';
	MOV	R10 RFP	;
	ADI	R10 #-172	;
	STR	R3 (R10)	;
;; row: :	REF	Tv86 Iv2 this; 	msg[i+2] = 'p';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-176	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv87 Tv85 Tv86; 	msg[i+2] = 'p';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-176	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-172	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-180	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv87 Li88 ; 	msg[i+2] = 'p';
	LDB	R3 Li88:	;	msg[i+2] = 'p';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-180	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv89 Iv3 this; 	msg[i+3] = 'l';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-184	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv90 Li32 Tv89; 	msg[i+3] = 'l';
	LDR	R4 Li32:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-184	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+3] = 'l';
	MOV	R10 RFP	;
	ADI	R10 #-188	;
	STR	R3 (R10)	;
;; row: :	REF	Tv91 Iv2 this; 	msg[i+3] = 'l';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-192	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv92 Tv90 Tv91; 	msg[i+3] = 'l';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-192	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-188	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-196	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv92 Li50 ; 	msg[i+3] = 'l';
	LDB	R3 Li50:	;	msg[i+3] = 'l';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-196	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv93 Iv3 this; 	msg[i+4] = 'i';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-200	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv94 Li36 Tv93; 	msg[i+4] = 'i';
	LDR	R4 Li36:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-200	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+4] = 'i';
	MOV	R10 RFP	;
	ADI	R10 #-204	;
	STR	R3 (R10)	;
;; row: :	REF	Tv95 Iv2 this; 	msg[i+4] = 'i';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-208	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv96 Tv94 Tv95; 	msg[i+4] = 'i';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-208	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-204	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-212	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv96 Li97 ; 	msg[i+4] = 'i';
	LDB	R3 Li97:	;	msg[i+4] = 'i';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-212	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv98 Iv3 this; 	msg[i+5] = 'c';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-216	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv99 Li39 Tv98; 	msg[i+5] = 'c';
	LDR	R4 Li39:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-216	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+5] = 'c';
	MOV	R10 RFP	;
	ADI	R10 #-220	;
	STR	R3 (R10)	;
;; row: :	REF	Tv100 Iv2 this; 	msg[i+5] = 'c';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-224	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv101 Tv99 Tv100; 	msg[i+5] = 'c';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-224	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-220	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-228	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv101 Li102 ; 	msg[i+5] = 'c';
	LDB	R3 Li102:	;	msg[i+5] = 'c';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-228	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv103 Iv3 this; 	msg[i+6] = 'a';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-232	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv104 Li43 Tv103; 	msg[i+6] = 'a';
	LDR	R4 Li43:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-232	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+6] = 'a';
	MOV	R10 RFP	;
	ADI	R10 #-236	;
	STR	R3 (R10)	;
;; row: :	REF	Tv105 Iv2 this; 	msg[i+6] = 'a';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-240	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv106 Tv104 Tv105; 	msg[i+6] = 'a';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-240	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-236	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-244	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv106 Li107 ; 	msg[i+6] = 'a';
	LDB	R3 Li107:	;	msg[i+6] = 'a';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-244	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv108 Iv3 this; 	msg[i+7] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-248	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv109 Li47 Tv108; 	msg[i+7] = 't';
	LDR	R4 Li47:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-248	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+7] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-252	;
	STR	R3 (R10)	;
;; row: :	REF	Tv110 Iv2 this; 	msg[i+7] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-256	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv111 Tv109 Tv110; 	msg[i+7] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-256	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-252	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-260	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv111 Li68 ; 	msg[i+7] = 't';
	LDB	R3 Li68:	;	msg[i+7] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-260	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv112 Iv3 this; 	msg[i+8] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-264	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv113 Li51 Tv112; 	msg[i+8] = 'e';
	LDR	R4 Li51:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-264	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+8] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-268	;
	STR	R3 (R10)	;
;; row: :	REF	Tv114 Iv2 this; 	msg[i+8] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-272	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv115 Tv113 Tv114; 	msg[i+8] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-272	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-268	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-276	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv115 Li35 ; 	msg[i+8] = 'e';
	LDB	R3 Li35:	;	msg[i+8] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-276	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv116 Iv3 this; 	msg[i+9] = 'm';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-280	;
	STR	R13 (R10)	;
;; row: :	ADD	Tv117 Li54 Tv116; 	msg[i+9] = 'm';
	LDR	R4 Li54:	;
	MOV	R10 RFP	;Load Address
	ADI	R10 #-280	;
	LDR	R13 (R10)	;
	LDR	R3 (R13)	;Load to register
	ADD	R3 R4	;	msg[i+9] = 'm';
	MOV	R10 RFP	;
	ADI	R10 #-284	;
	STR	R3 (R10)	;
;; row: :	REF	Tv118 Iv2 this; 	msg[i+9] = 'm';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-288	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv119 Tv117 Tv118; 	msg[i+9] = 'm';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-288	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	MOV	R10 RFP	;
	ADI	R10 #-284	;
	LDR	R14 (R10)	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-292	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv119 Li57 ; 	msg[i+9] = 'm';
	LDB	R3 Li57:	;	msg[i+9] = 'm';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-292	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv121 Iv2 this; 	msg[24] = 'E';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-296	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv122 Li120 Tv121; 	msg[24] = 'E';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-296	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li120:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-300	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv122 Li46 ; 	msg[24] = 'E';
	LDB	R3 Li46:	;	msg[24] = 'E';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-300	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv124 Iv2 this; 	msg[25] = 'n';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-304	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv125 Li123 Tv124; 	msg[25] = 'n';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-304	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li123:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-308	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv125 Li64 ; 	msg[25] = 'n';
	LDB	R3 Li64:	;	msg[25] = 'n';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-308	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv127 Iv2 this; 	msg[26] = 't';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-312	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv128 Li126 Tv127; 	msg[26] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-312	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li126:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-316	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv128 Li68 ; 	msg[26] = 't';
	LDB	R3 Li68:	;	msg[26] = 't';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-316	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv130 Iv2 this; 	msg[27] = 'e';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-320	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv131 Li129 Tv130; 	msg[27] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-320	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li129:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-324	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv131 Li35 ; 	msg[27] = 'e';
	LDB	R3 Li35:	;	msg[27] = 'e';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-324	;
	LDR	R13 (R10)	;
	STR	R3 (R13)	;Save from Register
;; row: :	REF	Tv133 Iv2 this; 	msg[28] = 'r';
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-328	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv134 Li132 Tv133; 	msg[28] = 'r';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-328	;
	LDR	R13 (R10)	;
	LDR	R13 (R13)
	LDR	R14 Li132:	;
	SUB	R12 R12
	ADI	R12 #1
	MUL	R14 R12
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-332	;
	STR	R13 (R10)	;
;; row: :	MOV	Tv134 Li135 ; 	msg[28] = 'r';
	LDB	R3 Li135:	;	msg[28] = 'r';
	MOV	R10 RFP	;Load Address
	ADI	R10 #-332	;
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


;; row: :	FUNC	Me8  ;     public void print(int i, int end) {
Me8:   ADI   R0 #0 ;    public void print(int i, int end) {
;; row: :	REF	Tv137 Iv2 this;     	    cout << msg[i];
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-20	;
	STR	R13 (R10)	;
;; row: :	AEF	Tv138 Pa6 Tv137;     	    cout << msg[i];
	MOV	R10 RFP	;Load Address
	ADI	R10 #-20	;
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
	ADI	R10 #-24	;
	STR	R13 (R10)	;
;; row: :	WRITE	Tv138  ;     	    cout << msg[i];
	MOV	R10 RFP	;Load Address
	ADI	R10 #-24	;
	LDR	R13 (R10)	;
	LDB	R0 (R13)	;Load to register
	TRP	#3	;    	    cout << msg[i];
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
JMR     R15             ; go back "    }	"


;; row: :	FUNC	St11  ; }
St11:   ADI   R0 #0 ;}
;; row: :	REF	Tv13 Iv2 this;     private char msg[];
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #0
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-12	;
	STR	R13 (R10)	;
;; row: :	REF	Tv15 Iv3 this;     private int i;
	MOV	R10 RFP	;
	ADI	R10 #-8	;
	LDR	R13 (R10)	;
	SUB	R14 R14
	ADI	R14 #4
	ADD	R13 R14
	MOV	R10 RFP	;Save Address
	ADI	R10 #-16	;
	STR	R13 (R10)	;
;; row: :	REF	Tv16 Iv4 this;     private int end;
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


;; row: :	FUNC	MAIN  ; void main() {
MAIN:   ADI   R0 #0 ;void main() {
;; row: :	MOV	Lv9 Li25 ;     int key = 1;
	LDR	R3 Li25:	;    int key = 1;
	MOV	R10 RFP	;
	ADI	R10 #-12	;
	STR	R3 (R10)	;
;; row: :	NEWI	Cl1 Tv140 ;     Message msg = new Message();
;; Test for heap overflow
	MOV     R10 R9
ADI     R10 #12
CMP     R10 RSL
BGT     R10 HOVRFLW:
MOV     R11 R9
ADI     R9 #12
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	STR	R11 (R10)	;
;; row: :	FRAME	Tv140 Co5 ;     Message msg = new Message();
;; Call function "Co5:        Message msg = new Message();"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-336          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
CMP     R10 RSL
BLT     R10 OVRFLW:
;; Create Activation Record and invoke Co5
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
;; row: :	CALL	Co5  ;     Message msg = new Message();
;; local varibales on the stack    ;     Message msg = new Message();
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-340
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Co5:
;; row: :	PEEK	Tv140  ;     Message msg = new Message();
	LDR	R11 (RSP)	;    Message msg = new Message();
	MOV	R10 RFP	;
	ADI	R10 #-20	;
	STR	R11 (R10)	;
;; row: :	MOV	Lv10 Tv140 ;     Message msg = new Message();
	MOV	R10 RFP	;    Message msg = new Message();
	ADI	R10 #-20	;
	LDR	R3 (R10)	;
	MOV	R10 RFP	;
	ADI	R10 #-16	;
	STR	R3 (R10)	;
;; row: :	FRAME	Lv10 Me8 ;     msg.print(14,22);
;; Call function "Me8:        msg.print(14,22);"
;; Test for overflow
:   MOV     R10 RSP
ADI     R10 #-28          ; 4 bytes for Return address & 4 bytes for Previous Frame Pointer 4 bytes for this (+ params) (+ local variables) (+ temp variables)
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
;; row: :	PUSH	Li74  ;     msg.print(14,22);
;; parameters on the stack (Li74)  ;     msg.print(14,22);
	LDR	R1 Li74:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	PUSH	Li141  ;     msg.print(14,22);
;; parameters on the stack (Li141)  ;     msg.print(14,22);
	LDR	R1 Li141:	;
STR     R1 (RSP)
ADI     RSP #-4
;; row: :	CALL	Me8  ;     msg.print(14,22);
;; local varibales on the stack    ;     msg.print(14,22);
;; set the stack pointer
	MOV	RSP R15
	ADI	RSP #-32
;; set the frame pointer
MOV     RFP R15
;; set the return address and jump
MOV     R10 RPC         ; PC already at next instruction
ADI     R10 #12
STR     R10 (RFP)
JMP     Me8:
;; row: :	PEEK	Tv142  ;     msg.print(14,22);
	LDR	R11 (RSP)	;    msg.print(14,22);
	MOV	R10 RFP	;
	ADI	R10 #-24	;
	STR	R11 (R10)	;
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


;; Heap starts here
FREE:    .INT 0
