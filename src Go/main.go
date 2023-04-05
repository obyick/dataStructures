package main

import (
	"fmt"

	"./listTypes"
)

func main() {
	// Create an ArrayList object and specify the size of the array
	arrayList := listTypes.ArrayList{}
	arraySize := 10

	// Initialize the array with zeros and print the result
	arrayList.Add(arraySize)
	fmt.Printf("Array with zeros: \t%v\n", arrayList)

	// Add integers to the array at specific indices and print the result
	for i := 0; i <= arraySize; i++ {
		arrayList.AddOnIndex(i, i)
	}
	fmt.Printf("Add integers: \t%v\n", arrayList)

	// Remove an integer from the array at a specific index and print the result
	arrayList.RmFromIndex(0)
	fmt.Printf("Remove an integer: \t%v\n", arrayList)

	// Get the value at a specific index in the array, and print the result or an error message
	value, err := arrayList.Get(6)
	if err == nil {
		fmt.Printf("Value at a specific index in the array: \t%v\n", value)
	} else {
		fmt.Printf("%v\t%v\n", value, err)
	}

	// Change the value at a specific index in the array and print the result or an error message
	arrayList.Set(666, 2)
	fmt.Printf("Array after changing value at specific index: \t%v\n", arrayList)

	// Print the size of the array
	fmt.Printf("Array size: \t%v\n", arrayList.Size())

	// Remove the empty space at the end of the array and print the result
	arrayList.RmDiff()
	fmt.Printf("Array without empty space: \t%v\n", arrayList)

	fmt.Printf("\n//////////////////////////////////////////////////////////////////////////////////////////////////////////\n\n")

	// Create a LinkedList object and specify the size of the linked list
	linkedList := listTypes.LinkedList{}
	linkedSize := 10

	// Add new nodes to the linked list at specific indices
	for i := 0; i < linkedSize; i++ {
		linkedList.Add(i, i)
	}

	// Print the contents of the linked list
	fmt.Printf("Contents of the linked list: ")
	linkedList.PrintLinked()
}
