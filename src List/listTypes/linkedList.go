package listTypes

import "fmt"

// Auxilary Interface (PrintLinked)

// PrintLinked method prints the contents of the LinkedList
func (linkedList *LinkedList) PrintLinked() {

	if linkedList.head == nil { // If the LinkedList is empty, it will print "Empty list"
		fmt.Println("Empty list")
		return
	}

	// Otherwise, traverse the list and print each node's value
	aux := linkedList.head
	fmt.Print("{[")
	for aux != nil {
		fmt.Print(aux.value)
		if aux.next != nil {
			fmt.Print(", ")
		}
		aux = aux.next
	}

	// Output
	fmt.Printf("] %d}\n", linkedList.size)
}

// List Interface

// Init method initializes the LinkedList
func (linkedList *LinkedList) Init() {
	linkedList.head = nil // Set head to nil
	linkedList.size = 0   // Set size to 0
}

// AddToBack method adds a new node at back to the LinkedList
func (linkedList *LinkedList) AddToBack(value int) {

	if linkedList.size == 0 { // If the LinkedList is empty, create a new node, set it as the head and increment the size
		linkedList.head = &Node{value, nil}
		linkedList.size++
	} else {
		// Look for the last node in the LinkedList
		current := linkedList.head
		for i := 0; i < linkedList.size-1; i++ {
			current = current.next
		}

		// Add a new node at back to the LinkedList and increment the size
		current.next = &Node{value, nil}
		linkedList.size++
	}
}

// RmFromBack method removes a node at back in the LinkedList
func (linkedList *LinkedList) RmFromBack() {

	if linkedList.size == 1 { // If the LinkedList has only one node, remove it and decrement the size
		linkedList.head = nil
		linkedList.size--
	} else {
		// Look for the last node in the LinkedList
		current := linkedList.head
		for i := 0; i < linkedList.size-1; i++ {
			current = current.next
		}

		// Set last node to 0 and decrement the size
		current.value = 0
		linkedList.size--
	}
}

// AddToIndex method adds a new node to the LinkedList at the specified index
func (linkedList *LinkedList) AddToIndex(value int, index int) {
	if index > linkedList.size {
		return
	}

	current := linkedList.head

	if index == 0 { // If the index is 0, set the new node as the head of the LinkedList and increment the size
		linkedList.head = &Node{value, current}
		linkedList.size++
		return
	}

	// Traverse the LinkedList until the node at the specified 'index - 1' is reached
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	// Insert the new node after the node at the specified index - 1 and increment the size
	current.next = &Node{value, current.next}
	linkedList.size++
}

// RmFromIndex method removes a node from the LinkedList at the specified index
func (linkedList *LinkedList) RmFromIndex(index int) {
	current := linkedList.head

	if index == 0 { // If the index is 0, remove the head of the LinkedList and decrement the size
		linkedList.head = current.next
		linkedList.size--
		return
	}
	// Traverse the LinkedList until the node at the specified 'index - 1' is reached
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	// Remove the node at the specified index by updating the previous node's next pointer and decrement the size
	if index < linkedList.size-1 {
		current.next = current.next.next
	}
	linkedList.size--
}

// Get method retrieves the value of the node at the specified index
func (linkedList *LinkedList) Get(index int) int {
	// Traverse the LinkedList until the node at the specified index is reached
	current := linkedList.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	// Return the value of the node at the specified index
	return current.value
}

// Set method sets the value of the node at the specified index
func (linkedList *LinkedList) Set(value int, index int) {
	// Traverse the LinkedList until the node at the specified index is reached
	current := linkedList.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	// Update the value of the node at the specified index
	current.value = value
}

// Length method returns the number of nodes in the LinkedList
func (linkedList *LinkedList) Length() int {
	// Return the size of the LinkedList
	return linkedList.size
}
