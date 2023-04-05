package listTypes

import (
	"errors"
	"fmt"
)

// Methods included in Auxilary Interface (ChangeSize, RmDiff)

// Resizes the underlying array if the size is equal to the capacity
func (arrayList *ArrayList) ChangeSize() {

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
func (arrayList *ArrayList) Add(size int) {
	arrayList.value = make([]int, size)
}

// Adds a specific value to the array at a given index
func (arrayList *ArrayList) AddOnIndex(value int, index int) {

	if index >= arrayList.size { // If the index is greater than or equal to the size of the array, the underlying array is resized
		arrayList.ChangeSize()
	}

	for i := arrayList.size; i > index; i-- {
		arrayList.value[i] = arrayList.value[i-1]
	}
	arrayList.value[index] = value
	arrayList.size++
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
func (arrayList *ArrayList) Get(index int) (int, error) {

	if index >= 0 && index <= arrayList.size { // If the index is out of range, an error is returned
		return arrayList.value[index], nil
	} else {
		return index, errors.New("it's a index out of range")
	}
}

// Sets the value at a specified index in the array to the given value
func (arrayList *ArrayList) Set(value int, index int) error {

	if index >= 0 && index <= arrayList.size { // If the index is out of range, an error is returned
		arrayList.value[index] = value
		return nil
	} else {
		return fmt.Errorf("%d it's a index out of range", index)
	}
}

// Returns the size of the array
func (arrayList *ArrayList) Size() int {
	return arrayList.size
}
