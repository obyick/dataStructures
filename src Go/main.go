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
	fmt.Printf("Initial Array: %v\n", arrayList)

	// Add integers to the array at specific indices and print the result
	for i := 0; i <= arraySize; i++ {
		arrayList.AddOnIndex(i, i)
	}
	fmt.Printf("Add integers: %v\n", arrayList)

	// Remove an integer from the array at a specific index and print the result
	arrayList.RmFromIndex(0)
	fmt.Printf("Remove an integer: %v\n", arrayList)

	// Get the value at a specific index in the array, and print the result or an error message
	value, err := arrayList.Get(6)
	if err == nil {
		fmt.Printf("Value at a specific index in the array: %v\n", value)
	} else {
		fmt.Printf("%v\t%v\n", value, err)
	}

	// Change the value at a specific index in the array and print the result or an error message
	arrayList.Set(666, 2)
	fmt.Printf("Array after changing value at specific index: %v\n", arrayList)

	// Print the size of the array
	fmt.Printf("Array size: %v\n", arrayList.Size())

	// Remove the empty space at the end of the array and print the result
	arrayList.RmDiff()
	fmt.Printf("Array without empty space: %v\n", arrayList)

	fmt.Printf("\n//////////////////////////////////////////////////////////////////////////////////////////////////////////\n\n")

	// Create a LinkedList object and specify the size of the linked list
	linkedList := listTypes.LinkedList{}
	linkedSize := 10

	// Add new node to the linked list at specific indices and print the result
	for i := 0; i < linkedSize; i++ {
		linkedList.Add(i, i)
	}
	fmt.Printf("Initial Linked List: ")
	linkedList.PrintLinked()

	// Removes a node to the linked list at specific indices and print the result
	linkedList.Rm(6)
	linkedList.Rm(5)
	linkedList.Rm(4)
	linkedList.Rm(0)
	fmt.Printf("Linked List after remove nodes: ")
	linkedList.PrintLinked()

	// Gets the value of the node at specific index in the linked list and print the result
	value, err = linkedList.Get(2)
	if value == 0 {
		fmt.Println(err)
	} else {
		fmt.Printf("Value at a specific index in the Linked List: %v\n", value)
	}

	// Sets the value of the node at specific index in the linked list and print the result
	linkedList.Set(666, 3)
	fmt.Printf("Linked List after change the value of the node at the specified index: ")
	linkedList.PrintLinked()
}
