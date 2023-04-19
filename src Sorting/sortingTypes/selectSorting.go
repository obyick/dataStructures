package sortingTypes

func SelectSort(v []int) {
	// Loop through the array
	for i := 0; i < len(v)-1; i++ {
		for j := 0 + i; j < len(v); j++ {
			if v[i] > v[j] {
				v[i], v[j] = v[j], v[i] // Swap the elements at index i and j
			}
		}
	}
}
