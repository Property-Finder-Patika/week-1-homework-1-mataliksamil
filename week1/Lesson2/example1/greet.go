package main

import (
	"fmt"
)

// main func takes no argument and return nothing
func main() {

	greet("Zeynep")

}

// ask about the "#" case error
func greet(name string) {
	//Printf formatted print as in C language
	fmt.Printf("Selam %s :)\n", name)
}
