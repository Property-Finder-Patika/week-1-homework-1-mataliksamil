// That's the main package.
package main

import (
	"errors"
	"fmt"
	"log"
)

//const lowestC = -273.0
//const lowestF = -460.0

// These are the constants representing the lowest possible degrees.
const (
	lowestC = -273.0
	lowestF = -460.0
)

// C2F Converts from celsius to fahrenheit
func C2F(cDegree float64) (float64, error) {
	if cDegree < lowestC {
		return 0, errors.New("Invalid degree passed: " + fmt.Sprintf("%f", cDegree))
	} else {
		return (9 * cDegree / 5) + 32, nil // nil in here represents no error message
	}
}

// Converts from fahrenheit to celsius.
func f2C(fDegree float64) (float64, error) {
	if fDegree < lowestF { // constant values can be compared with float64 values without an error in the scope of package
		return 0, errors.New("invalid degree passed: " + fmt.Sprintf("%f", fDegree))
	}
	return (fDegree - 32) * 5 / 9, nil
}

/*
That's a longer comment.
Using this format makes writing longer comments easier.
*/
func main() {
	fmt.Println(desc)
	cDegree := -40.0             // Try -340
	fDegree, err := C2F(cDegree) // this function call returns one value and an error message
	if err != nil {              // this case represent error case
		log.Fatal(err)
	} else {
		fmt.Printf("%f celcius is %f fahrenheit\n", cDegree, fDegree)
	}

	fDegree = 98.0 // Try -498
	cDegree, _ = f2C(fDegree)
	fmt.Printf("%f fahrenheit is %f celcius", fDegree, cDegree)
}

// this variable can be used above in the code even if it declared here below
var desc string = "This is a Celsius-Fahrenheit converter."
