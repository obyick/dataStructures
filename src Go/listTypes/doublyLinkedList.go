package listTypes

import "fmt"

// Methods included in Auxilary Interface (PrintLinked)

// This method prints the contents of the Doubly Linked List
func (doublyLinkedList *DoublyLinkedList) PrintDoublyLinked() {

	if doublyLinkedList.head == nil { // If the Doubly Linked List is empty, it will print "Empty list"
		fmt.Println("Empty list")
		return
	}

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

// Methods included in the List Interface.

// This method adds a new node to the Doubly Linked List at the specified index
func (doublyLinkedList *DoublyLinkedList) Add(value int, index int) error {

	if index < 0 || index > doublyLinkedList.size { // If the index is out of range it will return an error.
		return fmt.Errorf("%d it's a index out of range", index)
	}

	// Create a new node with the given value and nil next pointer
	newDoublyNode := &DoublyNode{value, nil, nil}

	if doublyLinkedList.head == nil { // If the Doubly Linked List is empty, set the new node as both head and tail
		doublyLinkedList.head = newDoublyNode
		doublyLinkedList.tail = newDoublyNode
	} else if index == 0 { // If the index is 0, insert the new node at the head of the list
		newDoublyNode.next = doublyLinkedList.head
		doublyLinkedList.head.prev = newDoublyNode
		doublyLinkedList.head = newDoublyNode
	} else if index == doublyLinkedList.size { // If the index is the size of the list, insert the new node at the tail
		newDoublyNode.prev = doublyLinkedList.tail
		doublyLinkedList.tail.next = newDoublyNode
		doublyLinkedList.tail = newDoublyNode
	} else { // Otherwise, iterate through the list to find the node that comes before the specified index
		prev := doublyLinkedList.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}

		// Insert the new node between the previous node and the next node
		newDoublyNode.prev = prev
		newDoublyNode.next = prev.next
		prev.next.prev = newDoublyNode
		prev.next = newDoublyNode
	}

	// Increment the size and return nil to indicate that no errors
	doublyLinkedList.size++
	return nil
}

// Removes a node from the Doubly Linked List at the specified index
func (doublyLinkedList *DoublyLinkedList) Rm(index int) error {

	if index < 0 || index > doublyLinkedList.size { // If the index is out of range it will return an error.
		return fmt.Errorf("%d it's a index out of range", index)
	}

	if doublyLinkedList.head == nil { // If the Doubly Linked List is empty, return an error
		return fmt.Errorf("doubly linked list is empty")
	}

	if index == 0 { // If the index is 0, remove the head node
		doublyLinkedList.head = doublyLinkedList.head.next
		if doublyLinkedList.head == nil {
			doublyLinkedList.tail = nil
		} else {
			doublyLinkedList.head.prev = nil
		}
	} else if index == doublyLinkedList.size {
		doublyLinkedList.tail = doublyLinkedList.tail.prev
		doublyLinkedList.tail.next = nil
	} else { // Otherwise, iterate through the list to find the node that comes before the specified index
		prev := doublyLinkedList.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}

		// Remove the node by updating the previous node's next pointer to skip over the node to be removed
		prev.next = prev.next.next
		prev.next.prev = prev

		if prev.next == nil { // If the node to be removed was the tail, update the tail to the previous node
			doublyLinkedList.tail = prev
		}
	}

	// Decrement the size and return nil to indicate that no errors
	doublyLinkedList.size--
	return nil
}

// Retrieves the value of the node at the specified index
func (doublyLinkedList *DoublyLinkedList) Get(index int) (int, error) {

	if index < 0 || index >= doublyLinkedList.size { // If the index is out of range it will return an error.
		return 0, fmt.Errorf("%d it's a index out of range", index)
	}

	// Start at the head of the list
	prev := doublyLinkedList.head
	for i := 0; i < index; i++ { // Iterate through the list until the specified index is reached
		prev = prev.next
	}

	return prev.value, nil // Return the value of the node at the specified index
}

// Sets the value of the node at the specified index
func (doublyLinkedList *DoublyLinkedList) Set(value int, index int) error {

	if index < 0 || index >= doublyLinkedList.size { // If the index is out of range it will return an error.
		return fmt.Errorf("%d it's a index out of range", index)
	}

	// Start at the head of the list
	prev := doublyLinkedList.head
	for i := 0; i < index; i++ { // Iterate through the list to find the node at the specified index
		prev = prev.next
	}

	// Set the value of the node to the specified value
	prev.value = value

	return nil
}
