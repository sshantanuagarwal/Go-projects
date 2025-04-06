package main

import (
	"fmt"
)

var name = "Sshantanu"
var name1 string = "Agarwal"

const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func main() {
	fmt.Println("Hello World!")

	fmt.Print("const: ", A, B, C, "\n")
	var (
		a = "sshaagar"
		b = 1
		c = true
	)

	fmt.Printf("string  has value: %s and type: %T\n", a, a)
	fmt.Printf("int has value: %b and type: %T\n", b, b)
	fmt.Printf("bool has value: %b and type: %T\n ", c, c)

	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)

	arr2 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(arr2))
}
