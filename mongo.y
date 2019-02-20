%{

package main

import (
	"fmt"
)

%}

%union{
	ex string
}

%token <ex> DB MONGO

%% 
query:  mongo
        |
        db
        ;

db:     DB
        {
            fmt.Printf("\tDB %s\n", $1 )
        }
        ;

mongo:  MONGO
        {
            fmt.Printf("\tMONGO %s\n", $1 )
        }
        ;
%%

type mlex struct {
	expr   string
	result int
}

func (f *mlex) Lex(lval *yySymType) int {
	yyErrorVerbose = true
    fmt.Println("lval")
    fmt.Println(lval)
	return MONGO
}

func (f *mlex) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func Parse(expr string) int {
	m := &mlex{expr: expr}
	yyParse(m)
	return m.result
}
