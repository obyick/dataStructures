package sortingTypes

func SelectionSort(v []int) {
	for i := 0; i < len(v)-1; i++ { // Loop through the array
		minor := i
		for j := i + 1; j < len(v); j++ {
			if v[j] < v[minor] {
				minor = j
			}
		}
		v[i], v[minor] = v[minor], v[i] // Swap the elements
	}
}
