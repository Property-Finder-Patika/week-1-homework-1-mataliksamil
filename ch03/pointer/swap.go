package main

import "fmt"

func main() {
	x := 5
	y := 7
	fmt.Printf("x = %d y = %d\n\n", x, y) // two var declared and printed

	x, y = swap1(x, y) // swap their values
	fmt.Printf("x = %d y = %d\n", x, y)

	x = 5
	y = 7
	swap2(x, y) // swaps but in the scope of function itself so swap never happens
	fmt.Printf("x = %d y = %d\n", x, y)

	x = 5
	y = 7
	swap3(&x, &y) // again no return but sends pointers so address change so swap successful
	fmt.Printf("x = %d y = %d\n", x, y)
}

func swap1(x, y int) (int, int) { // swaps x and y value
	temp := x
	x = y
	y = temp
	return x, y
}

func swap2(x, y int) {
	temp := x
	x = y
	y = temp
}

func swap3(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
