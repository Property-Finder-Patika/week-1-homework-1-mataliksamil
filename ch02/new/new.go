package main

import "fmt"

func main() {
	p := new(int)
	fmt.Printf("%d %d \t - address of the pointer / value of the pointer initially zero\n", p, *p)

	q := new(int)
	fmt.Printf("%d %d\n", q, *q)

	fmt.Printf("%t \t - addresses does not match but values are \n", p == q)
	fmt.Printf("%t\n", *p == *q)
}
