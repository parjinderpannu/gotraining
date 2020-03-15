// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

import "fmt"

// Add imports.

func main() {

	// Declare an array of 5 strings set to its zero value.
	var arr1 [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	arr2 := [5]string{"a", "b", "c", "d", "e"}
	// Assign the populated array to the array of zero values.
	for i, v := range arr2 {
		arr1[i] = v
	}
	// Iterate over the first array declared.
	// Display the string value and address of each element.
	for i, v := range arr1 {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &arr1[i])

	}
}
