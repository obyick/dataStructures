package listTypes

import "fmt"

// Methods included in Auxilary Interface (PrintLinked)

// This method prints the contents of the linked list
func (linkedList *LinkedList) PrintLinked() {

	if linkedList.head == nil { // If the linked list is empty, it will print "Empty list"
		fmt.Println("Empty list")
		return
	}

	aux := linkedList.head
	fmt.Print("{[")
	for aux != nil {
		fmt.Print(aux.value)
		if aux.next != nil {
			fmt.Print(", ")
		}
		aux = aux.next
	}
	fmt.Printf("] %d}\n", linkedList.size)
}

// Methods included in the List Interface.

// This method adds a new node to the linked list at the specified index
func (linkedList *LinkedList) Add(value int, index int) error {

	if index < 0 || index > linkedList.size { // If the index is out of range it will return an error.
		return fmt.Errorf("%d it's a index out of range", index)
	}

	// Create a new node with the given value and nil next pointer
	newNode := &Node{value, nil}

	if linkedList.head == nil { // If the linked list is empty, set the new node as both head and tail
		linkedList.head = newNode
		linkedList.tail = newNode
	} else if index == 0 { // If the index is 0, insert the new node at the head of the list
		newNode.next = linkedList.head
		linkedList.head = newNode
	} else if index == linkedList.size { // If the index is the size of the list, insert the new node at the tail
		linkedList.tail.next = newNode
		linkedList.tail = newNode
	} else { // Otherwise, iterate through the list to find the node that comes before the specified index
		prev := linkedList.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}

		// Insert the new node between the previous node and the next node
		newNode.next = prev.next
		prev.next = newNode
	}

	// Increment the size and return nil to indicate that no errors
	linkedList.size++
	return nil
}

// Removes a node from the linked list at the specified index
func (linkedList *LinkedList) Rm(index int) error {

	if index < 0 || index > linkedList.size { // If the index is out of range it will return an error.
		return fmt.Errorf("%d it's a index out of range", index)
	}

	if linkedList.head == nil { // If the linked list is empty, return an error
		return fmt.Errorf("linked list is empty")
	}

	if index == 0 { // If the index is 0, remove the head node
		linkedList.head = linkedList.head.next
		if linkedList.head == nil {
			linkedList.tail = nil
		}
	} else { // Otherwise, iterate through the list to find the node that comes before the specified index
		prev := linkedList.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}

		// Remove the node by updating the previous node's next pointer to skip over the node to be removed
		prev.next = prev.next.next

		if prev.next == nil { // If the node to be removed was the tail, update the tail to the previous node
			linkedList.tail = prev
		}
	}

	// Decrement the size and return nil to indicate that no errors
	linkedList.size--
	return nil
}

// Retrieves the value of the node at the specified index
func (linkedList *LinkedList) Get(index int) (int, error) {

	if index < 0 || index >= linkedList.size { // If the index is out of range it will return an error.
		return 0, fmt.Errorf("%d it's a index out of range", index)
	}

	// Start at the head of the list
	prev := linkedList.head
	for i := 0; i < index; i++ { // Iterate through the list until the specified index is reached
		prev = prev.next
	}

	return prev.value, nil // Return the value of the node at the specified index
}

// Sets the value of the node at the specified index
func (linkedList *LinkedList) Set(value int, index int) error {

	if index < 0 || index >= linkedList.size { // If the index is out of range it will return an error.
		return fmt.Errorf("%d it's a index out of range", index)
	}

	// Start at the head of the list
	prev := linkedList.head
	for i := 0; i < index; i++ { // Iterate through the list to find the node at the specified index
		prev = prev.next
	}

	// Set the value of the node to the specified value
	prev.value = value

	return nil
}
