package main

import (
	"fmt"
)

// general syntax for declaring a function in go is:
// func functionName(parameter type) returnType {...}

// if consecutive parameters are one of the same type - can write at end
func calculateBill(price, no int) int {
	totalPrice := price * no
	return totalPrice
}

// possible to return multiple types
// multiple return types are surrounded by parentheses
func rectProps(length, width float64) (float64, float64) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}

// can used name return values rather than returning the variables
func sqrProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // no explicit return value
}

func demoFunctions() {
	totalPrice := calculateBill(10, 8)
	fmt.Println("total bill:", totalPrice)
	area, perimeter := rectProps(10.8, 4.63)
	fmt.Printf("area: %f perimeter: %f", area, perimeter)
	// blank identifiers are used to discard an extra value
	a, _ := sqrProps(10.4, 3.2)
	fmt.Printf("\narea: %f", a)
}
