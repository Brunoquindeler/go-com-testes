package arrayseslices

import "testing"

func TestSliceSomador(t *testing.T) {
	numeros := []int{1, 2, 3}

	obtido := SliceSomador(numeros)
	esperado := 6

	if esperado != obtido {
		t.Errorf("\nDado: %v \nEsperado: %v \nObtido: %v", numeros, esperado, obtido)
	}
}
