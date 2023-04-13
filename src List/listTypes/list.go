package listTypes

// Interface which will define the common methods
type List interface {
	Init()
	AddToBack(value int)
	RmFromBack()
	AddToIndex(value int, index int)
	RmFromIndex(index int)
	Get(index int) int
	Set(value int, index int)
	Length() int
}

// Interface which will define additional methods that require extra functionality
type Auxilary interface {
	Double()
	RmDiff()
	PrintLinked()
	PrintDoublyLinked()
}

// Struct which will contain an array to hold elements and a size variable to track the current size
type ArrayList struct {
	value []int
	size  int
}

// Struct which will contain a head and tail pointer to track the first and last nodes, and a size variable to track the current size
type LinkedList struct {
	size int
	head *Node
}

// Struct which will contain a value variable to hold the element value and a next pointer to point to the next node
type Node struct {
	value int
	next  *Node
}

// Struct which will contain a head and tail pointer to track the first and last nodes, and a size variable to track the current size
type DoublyLinkedList struct {
	size int
	head *DoublyNode
	tail *DoublyNode
}

// Struct which will contain a value variable to hold the element value, a next pointer to point to the next node, and a prev pointer to point to the previous node.
type DoublyNode struct {
	value int
	next  *DoublyNode
	prev  *DoublyNode
}
