package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// initial cell of args is name of the program first cell is argument itself
	// check if an argument is given
	if len(os.Args) <= 1 {
		log.Fatal("usage: gohex  <filename>")
	}
	name := os.Args[2]
	greeting := CreateGreet(name)
	fmt.Printf("%s \n", greeting)
}

func CreateGreet(name string) string {
	greeting := "Selam " + name + " :) "
	return greeting
}
