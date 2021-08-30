package arrayseslices

// ArraySomador recebe um slice de inteiros
// e retorna a soma total de todos os números contidos nele
func SliceSomador(numeros []int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}
