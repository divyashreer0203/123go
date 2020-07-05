// Sample code to explain all the arithmetic operations
package main

import (
	"fmt"
)

func main() {
	var x int	// method 1: declaring & assigning a variable
	var y int	//	y is a variable of type int

	x = 1	// assigning a variable
	y = 2
	
	// it can also be assigned as x, y = 1, 2
	
	a:=5.0	// method 2: := takes type of the value assigned
	b:=2.5

	sum := x+y	// arithmetic + operator
	sub := y-x	// arithmetic - operator
	mul := x*y	// arithmetic * operator
	quotient := a / b		// arithmetic / oprator
	
	// All the variables used in one expression must be of same type
	
	fmt.Printf("Sum: %v, type of %T\n", sum, sum)	// %v--> print the go object  &  %T--> prints the type of the variable
	fmt.Printf("Difference: %v, type of %T\n", sub, sub)
	fmt.Printf("Product: %v, type of %T\n", mul, mul)
	fmt.Printf("Quotient: %v, type of %T\n", quotient, quotient)
}
