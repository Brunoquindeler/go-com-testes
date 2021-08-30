package arrayseslices

func MultiplesSlicesSomador(slices ...[]int) (total []int) {
	quantidadesDeSlices := len(slices)
	total = make([]int, quantidadesDeSlices)

	for i, numeros := range slices {
		total[i] = SliceSomador(numeros)
	}
	return
}
