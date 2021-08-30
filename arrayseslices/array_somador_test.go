package arrayseslices

import (
	"testing"
)

func TestArraySomador(t *testing.T) {
	numeros := [5]int{1, 2, 3, 4, 5}

	obtido := ArraySomador(numeros)
	esperado := 15

	t.Errorf("\nDado: %v \nEsperado: %d \nObtido: %d", numeros, esperado, obtido)
}
