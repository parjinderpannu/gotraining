// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare and initialize anonymous
// struct types.
package main

import "fmt"

type bill struct {
	flag    bool
	counter int16
	pi      float32
}

type nancy struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	// Declare a variable of an anonymous type set
	// to its zero value.
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of an anonymous type and init
	// using a struct literal.
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	var b bill
	var b1 bill
	var n nancy
	// b = n implicit conversion not allowed
	b = bill(n) // explicit conversion
	b1 = e2     // implicit conversion allowed for literals
	// functions are like literal and we pass functions like http
	// var b int
	// var n uint
	// above is real world senario if implicit conversion allowed
	// then value will be different
	fmt.Println(b, n)
	fmt.Println(b1)

	// Display the values.
	fmt.Printf("%+v\n", e2)
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
