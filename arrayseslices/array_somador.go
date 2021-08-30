package arrayseslices

func ArraySomador(numeros [5]int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}
