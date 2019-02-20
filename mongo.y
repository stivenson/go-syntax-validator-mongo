%{

package main

import (
	"fmt"
)

var expr string

%}

%union{
	val string
}

%type <val> expr

%token <val> DB MONGO

%token NUMBER

%%

query:
        expr 
        | 
        mongo
        |
        db
        ;

expr:   NUMBER
        {
            fmt.Printf("\nA number\n");
        }

db:     DB
        {
          fmt.Printf("\tDB\n");
        }
        ;

mongo: MONGO 
        {
          fmt.Printf("\tMONGO\n");
        }
        ;

%%  /*  start  of  programs  */
type mlex struct {
	expr   string
	result int
}

func (f *mlex) Lex(lval *yySymType) int {
	yyErrorVerbose = true
	return 0
}

func (f *mlex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func Parse(expr string) int {
	m := &mlex{expr, 0}
	yyParse(m)
	return m.result
}
