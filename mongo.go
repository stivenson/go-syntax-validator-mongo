// Code generated by goyacc - DO NOT EDIT.

package mongoparser

import __yyfmt__ "fmt"

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type pair struct {
	key string
	val interface{}
}

type lex struct {
	input  []byte
	pos    int
	result map[string]interface{}
	err    error
}

func newLex(input []byte) *lex {
	return &lex{
		input: input,
	}
}

// Error satisfies yyLexer.
func (l *lex) Error(s string) {
	l.err = errors.New(s)
}

func (l *lex) backup() {
	if l.pos == -1 {
		return
	}
	l.pos--
}

func (l *lex) scanLiteral(lval *yySymType) int {
	buf := bytes.NewBuffer(nil)
	for {
		b := l.next()
		switch {
		case unicode.IsLetter(rune(b)):
			buf.WriteByte(b)
		default:
			l.backup()
			val, ok := literal[buf.String()]
			if !ok {
				return LexError
			}
			lval.val = val
			return Literal
		}
	}
}

func (l *lex) scanNormal(lval *yySymType) int {
	for b := l.next(); b != 0; b = l.next() {
		switch {
		case unicode.IsSpace(rune(b)):
			continue
		case b == '"':
			return l.scanString(lval)
		case unicode.IsDigit(rune(b)) || b == '+' || b == '-':
			l.backup()
			return l.scanNum(lval)
		case unicode.IsLetter(rune(b)):
			l.backup()
			return l.scanLiteral(lval)
		default:
			return int(b)
		}
	}
	return 0
}

func (l *lex) next() byte {
	if l.pos >= len(l.input) || l.pos == -1 {
		l.pos = -1
		return 0
	}
	l.pos++
	return l.input[l.pos-1]
}

func (l *lex) scanString(lval *yySymType) int {
	buf := bytes.NewBuffer(nil)
	for b := l.next(); b != 0; b = l.next() {
		switch b {
		case '\\':
			// TODO(sougou): handle \uxxxx construct.
			b2 := escape[l.next()]
			if b2 == 0 {
				return LexError
			}
			buf.WriteByte(b2)
		case '"':
			lval.val = buf.String()
			return String
		default:
			buf.WriteByte(b)
		}
	}
	return LexError
}

func (l *lex) scanNum(lval *yySymType) int {
	buf := bytes.NewBuffer(nil)
	for {
		b := l.next()
		switch {
		case unicode.IsDigit(rune(b)):
			buf.WriteByte(b)
		case strings.IndexByte(".+-eE", b) != -1:
			buf.WriteByte(b)
		default:
			l.backup()
			val, err := strconv.ParseFloat(buf.String(), 64)
			if err != nil {
				return LexError
			}
			lval.val = val
			return Number
		}
	}
}

// Lex satisfies yyLexer.
func (l *lex) Lex(lval *yySymType) int {
	return l.scanNormal(lval)
}

func setResult(l yyLexer, v map[string]interface{}) {
	l.(*lex).result = v
}

// Parse parses the input and returs the result.
func Parse(input []byte) (map[string]interface{}, error) {
	l := newLex(input)
	_ = yyParse(l)
	return l.result, l.err
}

var escape = map[byte]byte{
	'"':  '"',
	'\\': '\\',
	'/':  '/',
	'b':  '\b',
	'f':  '\f',
	'n':  '\n',
	'r':  '\r',
	't':  '\t',
}

var literal = map[string]interface{}{
	"true":  true,
	"false": false,
	"null":  nil,
}

type yySymType struct {
	yys  int
	obj  map[string]interface{}
	list []interface{}
	pair pair
	val  interface{}
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault = 57350
	yyEofCode = 57344
	LexError  = 57346
	Literal   = 57349
	Number    = 57348
	String    = 57347
	yyErrCode = 57345

	yyMaxDepth = 200
	yyTabOfs   = -15
)

var (
	yyPrec = map[int]int{}

	yyXLAT = map[int]int{
		44:    0,  // ',' (16x)
		125:   1,  // '}' (12x)
		93:    2,  // ']' (11x)
		57347: 3,  // String (5x)
		123:   4,  // '{' (4x)
		57354: 5,  // object (4x)
		91:    6,  // '[' (3x)
		57351: 7,  // array (3x)
		57349: 8,  // Literal (3x)
		57348: 9,  // Number (3x)
		57356: 10, // value (3x)
		57344: 11, // $end (2x)
		57355: 12, // pair (2x)
		58:    13, // ':' (1x)
		57352: 14, // elements (1x)
		57353: 15, // members (1x)
		57350: 16, // $default (0x)
		57345: 17, // error (0x)
		57346: 18, // LexError (0x)
	}

	yySymNames = []string{
		"','",
		"'}'",
		"']'",
		"String",
		"'{'",
		"object",
		"'['",
		"array",
		"Literal",
		"Number",
		"value",
		"$end",
		"pair",
		"':'",
		"elements",
		"members",
		"$default",
		"error",
		"LexError",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {5, 3},
		2:  {15, 0},
		3:  {15, 1},
		4:  {15, 3},
		5:  {12, 3},
		6:  {7, 3},
		7:  {14, 0},
		8:  {14, 1},
		9:  {14, 3},
		10: {10, 1},
		11: {10, 1},
		12: {10, 1},
		13: {10, 1},
		14: {10, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [22][]uint8{
		// 0
		{4: 17, 16},
		{11: 15},
		{13, 13, 3: 20, 12: 19, 15: 18},
		{35, 34},
		{12, 12},
		// 5
		{13: 21},
		{3: 24, 17, 27, 23, 28, 26, 25, 22},
		{10, 10},
		{8, 2: 8, 24, 17, 27, 23, 28, 26, 25, 30, 14: 29},
		{5, 5, 5},
		// 10
		{4, 4, 4},
		{3, 3, 3},
		{2, 2, 2},
		{1, 1, 1},
		{32, 2: 31},
		// 15
		{7, 2: 7},
		{9, 9, 9},
		{3: 24, 17, 27, 23, 28, 26, 25, 33},
		{6, 2: 6},
		{14, 14, 14, 11: 14},
		// 20
		{3: 20, 12: 36},
		{11, 11},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 17

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			yyVAL.obj = yyS[yypt-1].obj
			setResult(yylex, yyVAL.obj)
		}
	case 2:
		{
			yyVAL.obj = map[string]interface{}{}
		}
	case 3:
		{
			yyVAL.obj = map[string]interface{}{
				yyS[yypt-0].pair.key: yyS[yypt-0].pair.val,
			}
		}
	case 4:
		{
			yyS[yypt-2].obj[yyS[yypt-0].pair.key] = yyS[yypt-0].pair.val
			yyVAL.obj = yyS[yypt-2].obj
		}
	case 5:
		{
			yyVAL.pair = pair{key: yyS[yypt-2].val.(string), val: yyS[yypt-0].val}
		}
	case 6:
		{
			yyVAL.val = yyS[yypt-1].list
		}
	case 7:
		{
			yyVAL.list = []interface{}{}
		}
	case 8:
		{
			yyVAL.list = []interface{}{yyS[yypt-0].val}
		}
	case 9:
		{
			yyVAL.list = append(yyS[yypt-2].list, yyS[yypt-0].val)
		}
	case 13:
		{
			yyVAL.val = yyS[yypt-0].obj
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
