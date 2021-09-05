package arrayseslices

func SomaTodoResto(slices ...[]int) []int {
	var total []int

	for _, slice := range slices {
		if len(slice) == 0 {
			total = append(total, 0)
		} else {
			final := slice[1:]
			total = append(total, SliceSomador(final))
		}

	}

	return total
}
