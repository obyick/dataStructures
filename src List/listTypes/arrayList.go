package listTypes

import "fmt"

// Methods included in Auxilary Interface (Double, RmDiff)

// Doubles the underlying array if the size is equal to the capacity
func (arrayList *ArrayList) Double() {
	if cap(arrayList.value) == arrayList.size {
		array := make([]int, cap(arrayList.value)*2)
		copy(array, arrayList.value)
		arrayList.value = array
	}
}

// Removes the unused capacity of the underlying array
func (arrayList *ArrayList) RmDiff() {
	arrayList.value = arrayList.value[:arrayList.size]
}

// Methods included in the List Interface.

// Initializes the array with a specific size
func (arrayList *ArrayList) Init() {
	arrayList.value = make([]int, 10)
	arrayList.size = 0
}

// Adds a specific value to the array at back
func (arrayList *ArrayList) AddToBack(value int) {
	arrayList.Double()
	arrayList.value[arrayList.size] = value
	arrayList.size++
}

// Removes the value at back in the array
func (arrayList *ArrayList) RmFromBack() {
	arrayList.value[arrayList.size-1] = 0
	arrayList.size--
}

// Adds a specific value to the array at a given index
func (arrayList *ArrayList) AddToIndex(value int, index int) {
	if index >= 0 && index <= arrayList.size {
		// Doubles the underlying array if the size is equal to the capacity
		arrayList.Double()

		for i := arrayList.size; i > index; i-- { // Moves all values after the index one position to back
			arrayList.value[i] = arrayList.value[i-1]
		}

		// Adds a specific value to the array at a given index and increment size
		arrayList.value[index] = value
		arrayList.size++
	} else {
		// Output if out of range
		fmt.Printf("%v it's out of the range\n", index)
	}
}

// Removes the value at a specific index in the array
func (arrayList *ArrayList) RmFromIndex(index int) {

	for i := index; i < arrayList.size-1; i++ {
		arrayList.value[i] = arrayList.value[i+1]
	}
	arrayList.value[arrayList.size-1] = 0
	arrayList.size--
}

// Returns the value at a specified index in the array, provided that the index meets the specified criteria
func (arrayList *ArrayList) Get(index int) int {

	if index >= 0 && index < arrayList.size { // If the index is out of range, an error is returned
		return arrayList.value[index]
	} else {
		return 0
	}
}

// Sets the value at a specified index in the array to the given value
func (arrayList *ArrayList) Set(value int, index int) {

	if index >= 0 && index <= arrayList.size { // If the index is out of range, an error is returned
		arrayList.value[index] = value
	}
}

// Returns the length of the array
func (arrayList *ArrayList) Length() int {
	return arrayList.size
}
