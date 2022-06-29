package main

import "fmt"

//s1 := "An example declaration." // := can only be used for local variables
var s1 string = "An example declaration."
var s2 = "Another declaration."
var s3 string

func main() {

	// The most verbose way
	var anInteger0 int = 32_768   // int is an alias for int32 in 32-bit system or int64 in 64-bit systems
	var anInteger1 int32 = 32_768 // manually declared the type
	var anInteger2 int64 = 32_768
	fmt.Printf("1) anInteger0: %d anInteger1: %d anInteger2: %d\n", anInteger0, anInteger1, anInteger2)

	// Shorter way, type is omitted, the type is inferred from the value
	var anotherInteger = 15
	fmt.Println("2) Another integer: ", anotherInteger)

	// But of course you can't do this due to missing type info!
	// var anotherInteger2
	// fmt.Print("Another integer 2: ", anotherInteger2)

	//The most concise way
	yetAnotherInteger := 15 //int or int64 in 64-bit machines
	fmt.Println("3) Yet another integer: ", yetAnotherInteger)

	// It takes a default value as zero
	var yetAnotherInteger2 int
	fmt.Println("4) Yet another integer 2: ", yetAnotherInteger2)

	// More than one variable can be created using their default values
	var a, b, c int    // 0 default
	var u1, u2 float32 // 0 default
	var s0 string      // default as ""
	var bool0 bool     // false default
	fmt.Printf("5) a: %d b: %d c: %d u1: %v u2: %v s0: %s bool0: %t \n", a, b, c, u1, u2, s0, bool0)

	// More than one variable can be created using their initial values
	var x, y, z int = 10, 20, 30
	fmt.Printf("6) %d %d %d\n", x, y, z)

	// More than one variable of different types can be created using their initial values
	var d, e, s1 = 100, 3.14, "Goooo" // types are inferred from values
	fmt.Printf("7) %d %f %s\n", d, e, s1)

	// Or in a declaration list more than one variable of different types can be created using their initial values
	var (
		s2         = "Hi, how are you?"
		i      int = 5
		j          = 12.8 //float64
		s3, s4     = "Love Go", "very much!"
		k          = 6.02e-23 //float64
	)
	fmt.Printf("8) %s %d %2.2f %s %s %v\n", s2, i, j, s3, s4, k)

	var byte1 byte = 255 // uint8 in range 0 - 255
	fmt.Printf("9) In binary %b and in decimal %d\n", byte1, byte1)

	other()

}

func other() {

	i, j := 3, 5
	fmt.Printf("10) %d %d\n", i, j)

	i, j, k := 3, 5, 7.5
	fmt.Printf("11) %d %d %d\n", i, j, k)

	//i, j, k := 13, 15, 17 // Compile time error due to no new declaration
	fmt.Printf("12) %d %d %d\n", i, j, k)

	i, j = j, i // swap values of i and j
	fmt.Printf("13) %d %d\n", i, j)

	var s1 string
	s1, s2 := "Heyy, what's up?", "I'm fine, thanks." // Can't only say: s1 := "Heyy, what's up?"
	fmt.Printf("14) %s %s\n", s1, s2)                 // these variables is valid in the scope of f function below
	f()
}

func f() {
	fmt.Printf("15) %s %s\n", s1, s2)
}
