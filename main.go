package main

import (
	"fmt"
)

func main() {

	fmt.Println("------ MongoDB ---------")
	res := Parse("db")
	fmt.Println(res)

}
