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
	arrayList.Init(arraySize)
	fmt.Printf("Initial Array: %v\n", arrayList)

	// Adds values to the array at back
	for i := 0; i < arraySize; i++ {
		arrayList.AddToBack(i)
	}
	fmt.Printf("Adding integers at back: %v\n", arrayList)

	// Removes values to the array at back
	for i := 0; i < arraySize/2; i++ {
		arrayList.RmFromBack()
	}
	fmt.Printf("Removes integers at back: %v\n", arrayList)

	// Add integers to the array at specific indices and print the result
	arrayList.AddToIndex(88, 3)
	fmt.Printf("Adding integer at the index: %v\n", arrayList)

	// Remove an integer from the array at a specific index and print the result
	arrayList.RmFromIndex(3)
	fmt.Printf("Remove integer at the index: %v\n", arrayList)

	// Adds values to the array at front
	for i := 1; i < arraySize; i++ {
		arrayList.AddToFront(i)
	}
	fmt.Printf("Adding integer at front: %v\n", arrayList)

	// Removes values to the array at front
	for i := 1; i < arraySize/2; i++ {
		arrayList.RmFromFront()
	}
	fmt.Printf("Removes integer at front: %v\n", arrayList)

	// Get the value at a specific index in the array, and print the result or an error message
	fmt.Printf("Value at a specific index in the array: %v\n", arrayList.Get(0))

	// Change the value at a specific index in the array and print the result or an error message
	arrayList.Set(666, 0)
	fmt.Printf("Array after changing value at specific index: %v\n", arrayList)

	// Print the size of the array
	fmt.Printf("Array Length: %v\n", arrayList.Length())

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
	value, err := linkedList.Get(2)
	if value == 0 {
		fmt.Println(err)
	} else {
		fmt.Printf("Value at a specific index in the Linked List: %v\n", value)
	}

	// Sets the value of the node at specific index in the linked list and print the result
	linkedList.Set(666, 3)
	fmt.Printf("Linked List after change the value of the node at the specified index: ")
	linkedList.PrintLinked()

	fmt.Printf("\n//////////////////////////////////////////////////////////////////////////////////////////////////////////\n\n")

	// Create a LinkedList object and specify the size of the linked list
	doublyLinkedList := listTypes.DoublyLinkedList{}
	doublyLinkedSize := 10

	// Add new node to the linked list at specific indices and print the result
	for i := 0; i < doublyLinkedSize; i++ {
		doublyLinkedList.Add(i, i)
	}
	fmt.Printf("Initial Doubly Linked List: ")
	doublyLinkedList.PrintDoublyLinked()

	// Removes a node to the linked list at specific indices and print the result
	// Add new node to the linked list at specific indices and print the result
	for i := 0; i < 5; i++ {
		doublyLinkedList.Rm(i + 2)
	}
	fmt.Printf("Doubly Linked List after remove nodes: ")
	doublyLinkedList.PrintDoublyLinked()

	index := 0

	// Gets the value of the node at specific index in the linked list and print the result
	value, err = doublyLinkedList.Get(index)
	if value == 0 && index != 0 {
		fmt.Println(err)
	} else {
		fmt.Printf("Value at index %v in the Doubly Linked List: %v\n", index, value)
	}

	// Sets the value of the node at specific index in the linked list and print the result
	doublyLinkedList.Set(666, index)
	fmt.Printf("Doubly Linked List after change the value of the node at the specified index: ")
	doublyLinkedList.PrintDoublyLinked()
}
