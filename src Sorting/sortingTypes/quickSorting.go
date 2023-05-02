package sortingTypes

func partition(v []int, low, high int) int {
	// Select the last element as the pivot.
	pivot := v[high]

	// Initialize i to the index of the first element.
	i := low - 1

	// Loop through the slice elements from the first to the second-to-last.
	for j := low; j < high; j++ {
		if v[j] < pivot {
			// Increment i and swap the current element with the element at i.
			i++
			v[i], v[j] = v[j], v[i]
		}
	}

	// Swap the pivot element with the element at i+1 to put it in its final position.
	v[i+1], v[high] = v[high], v[i+1]

	// Return the index of the pivot element.
	return i + 1
}

// quickSort recursively sorts the slice using the quicksort algorithm.
// It takes a slice, a low index, and a high index as arguments.
func quickSort(v []int, low, high int) {
	if low < high {
		// Partition the slice around a pivot.
		p := partition(v, low, high)

		// Recursively sort the two partitions to the left and right of the pivot.
		quickSort(v, low, p-1)
		quickSort(v, p+1, high)
	}
}

// QuickSortStart sorts the given slice using the quicksort algorithm.
// It takes a slice as an argument and returns nothing.
func QuickSortStart(v []int) {
	// Call quickSort with the slice and its boundaries.
	quickSort(v, 0, len(v)-1)
}
