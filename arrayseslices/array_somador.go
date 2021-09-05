package arrayseslices

// ArraySomador recebe um array de inteiros
// e retorna a soma total de todos os números contidos nele
func ArraySomador(numeros [5]int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}
