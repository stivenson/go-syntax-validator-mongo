package main

import (
	"fmt"
)

func main() {

	fmt.Println("------ MongoDB ---------")
	res := Parse("expr")
	fmt.Println(res)
}
