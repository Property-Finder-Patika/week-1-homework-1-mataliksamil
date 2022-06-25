package main

import "fmt"

type Person struct {
	name string
}

// this declaration means when I call greet func p instance will be available

func (p Person) greet(q Person) string {
	return "Selam from " + q.name + " to " + p.name + " :) "
}

func main() {
	var greeter Person = Person{name: "Nihal"}
	var myFriend Person = Person{name: "Mehmet"}
	// now we call greet funciton with greeter object and passing myFriend
	greeting := greeter.greet(myFriend)
	fmt.Printf("%s \n", greeting)
}
