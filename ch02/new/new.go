package main

import "fmt"

func main() {
	p := new(int)
	fmt.Printf("%d %d\n", p, *p)

	q := new(int)
	fmt.Printf("%d %d\n", q, *q)

	fmt.Printf("%t\n", p == q)
	fmt.Printf("%t\n", *p == *q)
}
