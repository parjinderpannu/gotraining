// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import (
	"fmt"
)

// Add imports.

// Add user type and provide comment.
type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	u1 := user{
		name:  "tom",
		email: "tom@gmail.com",
		age:   5,
	}
	// Display the field values.
	fmt.Println("u1.name : ", u1.name)
	fmt.Println("u1.email : ", u1.email)
	fmt.Println("u1.age : ", u1.age)
	// Declare a variable using an anonymous struct.
	anU := struct {
		name  string
		email string
		age   int
	}{
		name:  "john",
		email: "john@gmail.com",
		age:   45,
	}

	// Display the field values.
	fmt.Println("anU.name : ", anU.name)
	fmt.Println("anU.email : ", anU.email)
	fmt.Println("anU.age : ", anU.age)
}
