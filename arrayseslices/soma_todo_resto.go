package arrayseslices

func SomaTodoResto(slices ...[]int) []int {
	var total []int

	for _, slice := range slices {
		final := slice[1:]
		total = append(total, SliceSomador(final))
	}

	return total
}
