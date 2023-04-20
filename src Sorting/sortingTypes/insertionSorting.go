package sortingTypes

func InsertionSort(v []int) {
	for i := 1; i < len(v); i++ { // Loop through the array
		for j := i; j > 0 && v[j] < v[j-1]; j-- {
			v[j], v[j-1] = v[j-1], v[j] // Swap the elements
		}
	}
}
