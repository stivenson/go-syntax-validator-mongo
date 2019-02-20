%{
package mongoparser

import (
	"errors"
	"unicode"
	"bytes"
  "strings"
  "strconv"
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

%}

%union{
  obj map[string]interface{}
  list []interface{}
  pair pair
  val interface{}
}

%token LexError
%token <val> String Number Literal

%type <obj> object members
%type <pair> pair
%type <val> array
%type <list> elements
%type <val> value


%start object

%%

object: '{' members '}'
  {
    $$ = $2
    setResult(yylex, $$)
  }

members:
  {
    $$ = map[string]interface{}{}
  }
| pair
  {
    $$ = map[string]interface{}{
      $1.key: $1.val,
    }
  }
| members ',' pair
  {
    $1[$3.key] = $3.val
    $$ = $1
  }

pair: String ':' value
  {
    $$ = pair{key: $1.(string), val: $3}
  }

array: '[' elements ']'
  {
    $$ = $2
  }

elements:
  {
    $$ = []interface{}{}
  }
| value
  {
    $$ = []interface{}{$1}
  }
| elements ',' value
  {
    $$ = append($1, $3)
  }

value:
  String
| Number
| Literal
| object
  {
    $$ = $1
  }
| array
