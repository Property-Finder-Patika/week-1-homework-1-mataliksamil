package main

import "fmt"

// i can be used even if it is declared below parts of the code
var j int = i + 1

// function calls  by declaration out of the scope of the main function
func createB() bool {
	return true
}

var i int = createI()

func main() {
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(b)
}

func createI() int {
	fmt.Println("In createI")
	//fmt.Printf("%d %d \n", i, j)
	return 5
}

func init() {
	fmt.Println("In init I")
	fmt.Printf("%d %d \n", i, j)
}

func init() {
	fmt.Println("In init II")
	fmt.Printf("%d %d \n", i, j)
}

var b bool = createB()
