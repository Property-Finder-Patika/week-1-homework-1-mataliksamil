package main

import (
	"fmt"
)

var myInt = 100

// this declaration not allowed
// myInt2 := 200

/*
// standard variable declaration
func main() {
	// var "name of variable" type
	var name string = "Betül"
	var greeting = CreateGreet(name)
	fmt.Printf("%s", greeting)
}
*/

// inferred declaration exmple
func main() {
	// var "name of variable" type
	var name string = "Kerim"
	// type is inferred "INFERENCEIS THE KEY IN :="
	// type determined in compile time
	greeting := CreateGreet(name)
	fmt.Printf("%s", greeting)
}

func CreateGreet(name string) string {
	// ":="
	greeting := "Selam " + name + " :) "
	return greeting
}
