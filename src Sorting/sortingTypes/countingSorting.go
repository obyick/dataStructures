package sortingTypes

func CountingSort(v []int) []int {
	minor := v[0]
	major := v[0]

	for i := 1; i < len(v); i++ {
		if v[i] > major {
			major = v[i]
		}
		if v[i] < minor {
			minor = v[i]
		}
	}

	c := make([]int, (major - minor + 1))

	for i := 0; i < len(v); i++ {
		c[v[i]-minor]++
	}
	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	o := make([]int, len(v))

	for i := 0; i < len(o); i++ {
		o[c[v[i]-minor]-1] = v[i]
	}

	return o
}
