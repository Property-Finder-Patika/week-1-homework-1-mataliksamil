package main

import (
	"fmt"
	"math"
)

func main() {
	fp1 := 3.14
	var i int
	i = int(fp1)
	fmt.Println(i)

	//i = int(6.28) // Cannot convert an expression of the type 'float64' to the type 'int'
	i = int(math.Floor(6.28))
	fmt.Println(i)

	type yazi string
	var y yazi = "I like Go!"
	type text string
	var t text = "I love Go!"
	var s1 string = string(y)
	var s2 string = string(t)
	fmt.Println(s1, s2)

}
