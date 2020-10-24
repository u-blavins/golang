package main

// every Go file must start with the `package name` statement
// packages are used to provide code reusability. The main function should
// always reside in the main package

import (
	"fmt"
	"math"
)

// import statement used to import other packages
// `fmt` package is imported and used in main function to print

func main() {
	// fmt.Println("Hello World")
	// demoVariables()
	// shortHandDec()
	// demoConstants()
	demoFunctions()
}

func demoVariables() {
	// `var name type` is the syntax to declare a single variable
	var age int = 22 // variable declaration
	// age is initialised with a zero value '0'
	fmt.Println("My age is", age)
	//  if a variable has an initial value - go can infer the type
	var anotherAge = 22
	fmt.Println("Bar's age is", anotherAge)
	var height, width int = 50, 100
	fmt.Println("Width is", width, "and height is", height)
	var (
		name   = "Ruby"
		animal = "Cat"
		colour string
	)
	fmt.Println("My", animal, "is called", name, "and is", colour)
}

func shortHandDec() {
	// Go provides another concise way to declare variables using ':='
	// `name := initialvalue`
	// count := 10
	// fmt.Println("Count = ", count)
	name, age := "Foo", 21
	fmt.Println("my name is", name, "age is", age)
	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("Minimum value is", c)
	// Go is strongly typed, `name` = 29 cannot happen
}
