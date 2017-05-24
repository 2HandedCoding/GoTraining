// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"
import "time"

// main is the entry point for the application.
func main() {
	// Declare variables that are set to their zero value.
	var foo string
	var largeInteger int64
	var smallInteger int8

	// Display the value of those variables.
	fmt.Printf("\nThe three variables declared above are: \n%T: %v\n%T: %v\n%T: %v\n", foo, foo, largeInteger, largeInteger, smallInteger, smallInteger)

	// Declare variables and initialize.
	var name = "Yair"
	var age = 38
	var dob, error = time.Parse("2006-1-2", "1979-4-16")
	if error != nil {
		fmt.Println(error)
	}

	// Using the short variable declaration operator.
	name2 := "Stu"
	age2 := 25
	var dob2, error2 = time.Parse("2006-1-2", "1992-4-16")
	if error2 != nil {
		fmt.Println(error2)

	}
	// Display the value of those variables.
	fmt.Printf("\nThe three variables declared with regular initialization are: \n%T: %v\n%T: %v\n%T: %v\n", name, name, age, age, dob, dob)
	fmt.Printf("\nThe three variables declared with shorthand initialization are: \n%T: %v\n%T: %v\n%T: %v\n", name2, name2, age2, age2, dob2, dob2)

	// Perform a type conversion.
	ageFloat := float32(age)

	// Display the value of that variable.
	fmt.Printf("The converted variable is:\n%T: %v\n", ageFloat, ageFloat)
}
