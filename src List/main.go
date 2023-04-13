package main

import (
	"fmt"

	"./listTypes"
)

func main() {
	// Create an ArrayList object and specify the size of the array
	arrayList := listTypes.ArrayList{}
	arraySize := 10

	fmt.Printf("\n////////////////////////////////////// ARRAYLIST //////////////////////////////////////\n\n")

	// Initialize the ArrayList with zeros and print the result
	arrayList.Init()
	fmt.Printf("Initial ArrayList: %v\n", arrayList)

	// Adds values to the ArrayList at back and print the result
	for i := 0; i < arraySize; i++ {
		arrayList.AddToBack(i)
	}
	fmt.Printf("Adding integers at back: %v\n", arrayList)

	// Removes values from the back of the ArrayList and print the result
	for i := 0; i < arraySize/2; i++ {
		arrayList.RmFromBack()
	}
	fmt.Printf("Removes integers from back: %v\n", arrayList)

	// Add an integer to the ArrayList at a specific index and print the result
	arrayList.AddToIndex(88, 3)
	fmt.Printf("Adding integer at the index: %v\n", arrayList)

	// Remove an integer from the ArrayList at a specific index and print the result
	arrayList.RmFromIndex(3)
	fmt.Printf("Remove integer from index: %v\n", arrayList)

	// Get the value at a specific index in the ArrayList, and print the result
	fmt.Printf("Value at a specific index in the ArrayList: %v\n", arrayList.Get(0))

	// Set the value at a specific index in the ArrayList and print the result
	arrayList.Set(666, 0)
	fmt.Printf("ArrayList after changing value at specific index: %v\n", arrayList)

	// Print the size of the ArrayList
	fmt.Printf("ArrayList Length: %v\n", arrayList.Length())

	// Remove the empty space at the end of the ArrayList and print the result
	arrayList.RmDiff()
	fmt.Printf("ArrayList without empty space: %v\n", arrayList)

	fmt.Printf("\n////////////////////////////////////// LINKEDLIST /////////////////////////////////////\n\n")

	// Create a LinkedList object and specify the size
	linkedList := listTypes.LinkedList{}
	linkedSize := 10

	// Initializes an empty LinkedList and print the result
	linkedList.Init()
	linkedList.PrintLinked()

	// Add new nodes to the LinkedList at the back and print the result
	for i := 0; i < linkedSize; i++ {
		linkedList.AddToBack(i)
	}
	fmt.Printf("Adding Node to back: ")
	linkedList.PrintLinked()

	// Removes a node from the back of the LinkedList and print the result
	for i := 0; i < linkedSize/2; i++ {
		linkedList.RmFromBack()
	}
	fmt.Printf("Removes Node at back: ")
	linkedList.PrintLinked()

	// Add a new node to the LinkedList at a specific index and print the result
	linkedList.AddToIndex(88, 3)
	fmt.Printf("Adding node at specific index: ")
	linkedList.PrintLinked()

	// Remove a node from the LinkedList at a specific index and print the result
	linkedList.RmFromIndex(3)
	fmt.Printf("Remove node at specific index: ")
	linkedList.PrintLinked()

	// Get the value of the node at a specific index in the LinkedList and print the result
	fmt.Printf("Value at a specific index in the LinkedList: %v\n", linkedList.Get(3))

	// Set the value of the node at specific index in the LinkedList and print the result
	linkedList.Set(88, 3)
	fmt.Printf("LinkedList after change the value of the node at the specified index: ")
	linkedList.PrintLinked()

	// Print the size of the LinkedList
	fmt.Printf("LinkedList Length: %v\n", linkedList.Length())

	fmt.Printf("\n/////////////////////////////////// DOUBLYLINKEDLIST //////////////////////////////////\n\n")

	// Create a DoublyLinkedList object and specify the size
	doublyLinkedList := listTypes.DoublyLinkedList{}
	doublyLinkedSize := 10

	// Initializes an empty DoublyLinkedList and print the result
	doublyLinkedList.Init()
	doublyLinkedList.PrintDoublyLinked()

	// Add new nodes to the DoublyLinkedList at the back and print the result
	for i := 0; i < doublyLinkedSize; i++ {
		doublyLinkedList.AddToBack(i)
	}
	fmt.Printf("Adding Node to back: ")
	doublyLinkedList.PrintDoublyLinked()

	// Removes a node from the back of the DoublyLinkedList and print the result
	for i := 0; i < doublyLinkedSize/2; i++ {
		doublyLinkedList.RmFromBack()
	}
	fmt.Printf("Removes Node at back: ")
	doublyLinkedList.PrintDoublyLinked()

	// Add a new node to the DoublyLinkedList at a specific index and print the result
	doublyLinkedList.AddToIndex(88, 0)
	fmt.Printf("Adding node at specific index: ")
	doublyLinkedList.PrintDoublyLinked()

	// Remove a node from the DoublyLinkedList at a specific index and print the result
	doublyLinkedList.RmFromIndex(3)
	fmt.Printf("Remove node at specific index: ")
	doublyLinkedList.PrintDoublyLinked()

	// Get the value of the node at a specific index in the DoublyLinkedList and print the result
	fmt.Printf("Value at a specific index in the DoublyLinkedList: %v\n", doublyLinkedList.Get(0))

	// Set the value of the node at specific index in the DoublyLinkedList and print the result
	doublyLinkedList.Set(88, 3)
	fmt.Printf("DoublyLinkedList after change the value of the node at the specified index: ")
	doublyLinkedList.PrintDoublyLinked()

	// Print the size of the LinkedList
	fmt.Printf("DoublyLinkedList Length: %v\n", doublyLinkedList.Length())
}
