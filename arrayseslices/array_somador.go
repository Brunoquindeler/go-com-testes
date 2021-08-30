package arrayseslices

func ArraySomador(numeros [5]int) int {
	total := 0

	for i := 0; i < 5; i++ {
		total += numeros[i]
	}
	return total
}
