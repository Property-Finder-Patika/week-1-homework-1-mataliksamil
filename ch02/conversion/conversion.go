package main

import (
	"fmt"
	"math"
)

func main() {
	fp1 := 3.14
	fmt.Printf("type of fp1 : %T \n", fp1)
	var i int
	i = int(fp1) // float64 to int conversion is valid but precision lost 3.14 -> 3
	fmt.Println(i)

	//i = int(6.28) // Cannot convert an expression of the type 'float64' to the type 'int'
	i = int(math.Floor(6.28))
	fmt.Println(i)

	type yazi string
	var y yazi = "I like Go!"
	type text string
	var t text = "I love Go!"
	var s1 string = string(y)
	var s2 string = string(t) // text(string) to string conversion is valid
	fmt.Println(s1, s2)

}
