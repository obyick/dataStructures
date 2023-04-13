package listTypes

import "fmt"

// Auxilary Interface (PrintDoublyLinked)

// This method prints the contents of the DoublyLinkedList
func (doublyLinkedList *DoublyLinkedList) PrintDoublyLinked() {

	if doublyLinkedList.head == nil { // If the DoublyLinkedList is empty, it will print "Empty list"
		fmt.Println("Empty list")
		return
	}

	// Otherwise, traverse the list and print each node's value
	aux := doublyLinkedList.head
	fmt.Print("{[")
	for aux != nil {
		fmt.Print(aux.value)
		if aux.next != nil {
			fmt.Print(", ")
		}
		aux = aux.next
	}
	fmt.Printf("] %d}\n", doublyLinkedList.size)
}

// List Interface

// Init method initializes the DoublyLinkedList
func (doublyLinkedList *DoublyLinkedList) Init() {
	doublyLinkedList.head = nil                   // Set head to nil
	doublyLinkedList.tail = doublyLinkedList.head // Set tail equal head
	doublyLinkedList.size = 0                     // Set size to 0
}

// AddToBack method adds a new node at back to the DoublyLinkedList
func (doublyLinkedList *DoublyLinkedList) AddToBack(value int) {

	if doublyLinkedList.size == 0 { // If the DoublyLinkedList is empty, create a new node, set it as the head and tail and increment the size
		doublyLinkedList.head = &DoublyNode{value, nil, nil}
		doublyLinkedList.tail = doublyLinkedList.head
		doublyLinkedList.size++
	} else { // Otherwise add a new node at back to the DoublyLinkedList and increment the size
		doublyLinkedList.tail.next = &DoublyNode{value, nil, doublyLinkedList.tail}
		doublyLinkedList.tail = doublyLinkedList.tail.next
		doublyLinkedList.size++
	}
}

// RmFromBack method removes a node at back in the DoublyLinkedList
func (doublyLinkedList *DoublyLinkedList) RmFromBack() {

	if doublyLinkedList.size == 1 { // If the DoublyLinkedList has only one node, remove it and decrement the size
		doublyLinkedList.head = nil
		doublyLinkedList.tail = nil
		doublyLinkedList.size--
	} else { // Otherwise, set last node to nil and decrement the size
		doublyLinkedList.tail = doublyLinkedList.tail.prev
		doublyLinkedList.tail.next = nil
		doublyLinkedList.size--
	}
}

// AddToIndex method adds a new node to the DoublyLinkedList at the specified index
func (doublyLinkedList *DoublyLinkedList) AddToIndex(value int, index int) {
	if index < 0 || index > doublyLinkedList.size {
		return
	}

	current := doublyLinkedList.head

	if index == 0 { // If the index is 0, set the new node as the head of the DoublyLinkedList, tail equal nil and increment the size
		doublyLinkedList.head = &DoublyNode{value, current, nil}
		current.prev = doublyLinkedList.head
		doublyLinkedList.size++
		return
	}

	// Traverse the DoublyLinkedList until the node at the specified 'index - 1' is reached
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	// Insert the new node after the node at the specified index - 1 and increment the size
	current.next = &DoublyNode{value, current.next, current}
	doublyLinkedList.size++
}

// RmFromIndex method removes a node from the DoublyLinkedList at the specified index
func (doublyLinkedList *DoublyLinkedList) RmFromIndex(index int) {
	if index < 0 || index > doublyLinkedList.size {
		return
	}

	current := doublyLinkedList.head

	if index == 0 { // If the index is 0, remove the head of the DoublyLinkedList and decrement the size
		doublyLinkedList.head = current.next
		doublyLinkedList.head.prev = nil
		doublyLinkedList.size--
		return
	} else if doublyLinkedList.size == 1 {
		doublyLinkedList.head = nil
		doublyLinkedList.size--
		return
	}

	// Traverse the LinkedList until the node at the specified 'index - 1' is reached
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	// Remove the node at the specified index by updating the previous node's next pointer and decrement the size
	if index < doublyLinkedList.size-1 {
		current.next = current.next.next
		current.next.prev = current
	} else {
		current.next = nil
		doublyLinkedList.tail = current
	}

	doublyLinkedList.size--
}

// Get method retrieves the value of the node at the specified index
func (doublyLinkedList *DoublyLinkedList) Get(index int) int {
	if index < 0 || index > doublyLinkedList.size {
		return 0
	}

	current := doublyLinkedList.head
	if index == 0 {
		return current.value
	}

	// Traverse the DoublyLinkedList until the node at the specified index is reached
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	// Return the value of the node at the specified index
	return current.next.value
}

// Set method sets the value of the node at the specified index
func (doublyLinkedList *DoublyLinkedList) Set(value int, index int) {
	if index < 0 || index > doublyLinkedList.size {
		return
	}

	// Traverse the DoublyLinkedList until the node at the specified index is reached
	current := doublyLinkedList.head
	for i := 0; i < index-1; i++ {
		current = current.next
	}

	// Update the value of the node at the specified index
	current.next.value = value
}

// Length method returns the number of nodes in the DoublyLinkedList
func (doublyLinkedList *DoublyLinkedList) Length() int {
	// Return the size of the DoublyLinkedList
	return doublyLinkedList.size
}
