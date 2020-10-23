package main

import (
	"fmt"
)

func demoConstants() {
	// consts cannot be reassigned again by any other value
	const a = 50
	fmt.Println("a:", a)

	const (
		name    = "Joe"
		age     = 22
		country = "England"
	)
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("country:", country)

	// value of a constant should be known at compile time
	// var a = math.sqrt(4) - allowed
	// const a = math.sqrt(4) - not allowed

	demoConstantTypes()
}

func demoConstantTypes() {
	var defaultName = "Sam"
	type myString string
	var customName myString = "Sam"
	fmt.Printf("defaultName: %T customName: %T", defaultName, customName)
	// customName = defaultName - not allowed
	// even though both types are string, Go thinks they are different types
}
