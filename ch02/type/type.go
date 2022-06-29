package main

import "fmt"

type yazi string // type dec

func main() {

	var s yazi = "I love Go!"
	fmt.Printf("%s\n", s)

	type sayi int
	var i sayi = 3
	fmt.Printf("%d\n", i)

	type text string
	var t text = "I love Go!"
	// type mismatch below because yazi vs tex even if they are both strings
	//fmt.Println(s == t) // That's also related to assignability

	//var ss string = "I love Go!"
	//fmt.Println(s == ss)

	var sss string = string(t)
	fmt.Println(sss)
}
