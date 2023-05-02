package sortingTypes

// Sorts an integer slice in ascending order using the selection sort algorithm
func SelectionSort(v []int) {
	for i := 0; i < len(v)-1; i++ { // Loop through the array up to the second to last element
		minor := i
		for j := i + 1; j < len(v); j++ { // Loop through the remaining unsorted elements and find the index of the smallest element
			if v[j] < v[minor] {
				minor = j
			}
		}

		v[i], v[minor] = v[minor], v[i] // Swap the smallest element with the current element, putting it in its sorted position
	}
}
