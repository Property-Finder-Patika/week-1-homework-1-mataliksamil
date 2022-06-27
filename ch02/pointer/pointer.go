package main

import (
	"fmt"
)

func main() {
	i := 5  // i as an integer variable
	p := &i // p as pointer that points the address of i
	fmt.Printf("%d %d \n", i, *p)
	fmt.Printf("%d %d \n\n", p, &i)

	i = 6                         // this is an assignment
	fmt.Printf("%d %d \n", i, *p) // assignment affects the value that pointer points
	fmt.Printf("%d %d \n\n", p, &i)

	*p = 7 // pointer value updated which means i variables value also updated
	fmt.Printf("%d %d \n", i, *p)
	fmt.Printf("%d %d \n\n", p, &i)

	fmt.Println("After creating one more pointer.")
	q := p // pointer declaration and assignment together
	fmt.Printf("%d %d %d \n", i, *p, *q)
	fmt.Printf("%d %d %d\n\n", p, q, &i) // nowp and q pointers points the addres of i variable so their addresses are equal

	p = f(i) // now the p pointer updated with the new memory location where the i++ value is kept
	fmt.Printf("%d %d \n", i, *p)
	fmt.Printf("%d %d \n\n", p, &i)

	fmt.Println("After setting to nil.")
	p = nil // that is how to free a pointer so that gc can detect it
	fmt.Printf("%d %d \n", i, p)
	fmt.Printf("%d %d \n", p, &i)

}

func f(i int) *int { // takes an int var and returns pointer
	fmt.Println("in f()")
	i++       // raise the i val
	return &i // returns the address of i however it is not the same i in the func because it is generated in f functions scope
}
