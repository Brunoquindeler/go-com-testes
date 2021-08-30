package arrayseslices

func MultiplesSlicesSomador(slices ...[]int) []int {
	var total []int

	for _, slice := range slices {
		total = append(total, SliceSomador(slice))
	}
	return total
}
