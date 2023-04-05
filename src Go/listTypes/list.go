package listTypes

// Interface which will define the common methods
type List interface {
	Add(size int)
	AddOnIndex(value int, index int)
	RmFromIndex(index int)
	Get(index int) (int error)
	Set(value int, index int)
	Size() int
}

// Interface which will define additional methods that require extra functionality
type Auxilary interface {
	ChangeSize()
	RmDiff()
	PrintLinked()
}

// Struct which will contain an array to hold elements and a size variable to track the current size
type ArrayList struct {
	value []int
	size  int
}

// Struct which will contain a head and tail pointer to track the first and last nodes, and a size variable to track the current size
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// Struct which will contain a value variable to hold the element value and a next pointer to point to the next node
type Node struct {
	value int
	next  *Node
}
