package main

import (
	"fmt"
	"os"
)

var s string = "I love Go!" // in the scope of whole package

func main() {
	fmt.Printf("%s\n", s)                 //
	fmt.Println("______________________") //
	var s string = "I really love Go!"    // assignment (update) on s variable
	fmt.Printf("%s\n", s)                 // assignment works on func main scope

	//var s string = "I said you, I really love Go!"
	//fmt.Printf("%s\n", s)
	//x := "hellooww"
	for i, s := range s { // That's so strange!
		s := s // parse s string indexes and use all of the instances seperately but IN THE SCOPE OF FOR LOOP
		fmt.Printf("%d:  %c\n", i, s)
	}
	fmt.Printf("%s\n", s)

	fmt.Println("*****************")

	length := len(s)
	fmt.Printf("%d\n", length)
	if length := f(); /*This assignment is valid in if statemen scope*/ length > length {
		fmt.Printf("In if %d\n", length)
	} else {
		fmt.Printf("In else %d\n", length)
	}
	fmt.Printf("%d\n", length) // assignment in the if statement scope does not effect outside

	fmt.Println("*****************")
}

func f() int {
	return 20
}

// GOPL p. 48
func g1() error {
	//if f, err := os.Open("scope.go"); err != nil { // compile error: unused: f
	//	return err
	//}
	//f.Stat() // compile error: undefined f
	//f.Close()    // compile error: undefined f
	return nil
}

func g2() error {
	f, err := os.Open("/Users/akin/Development/Go/Projects/Fundamentals of Go/ch02/scope/scope.go")
	if err != nil {
		return err
	}
	f.Stat()  // compile error: undefined f
	f.Close() // compile error: undefined f
	return nil
}
