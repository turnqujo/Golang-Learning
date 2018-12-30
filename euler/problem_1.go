package eulerproblems

func problemOne(limit int) int {
	return sumSlice(buildMultiplesSlice(limit))
}

func buildMultiplesSlice(limit int) []int {
	result := []int{}
	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			result = append(result, i)
		}
	}

	return result
}

func sumSlice(slice []int) int {
	result := 0
	for i := 0; i < len(slice); i++ {
		result += slice[i]
	}

	return result
}
