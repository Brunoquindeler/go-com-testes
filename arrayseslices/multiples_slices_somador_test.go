package arrayseslices

import "testing"

func TestMultiplesSlicesSomador(t *testing.T) {
	slice1 := []int{1, 2}
	slice2 := []int{0, 9}

	obtido := MultiplesSlicesSomador(slice1, slice2)
	esperado := []int{3, 9}

	if esperado != obtido {
		t.Errorf("\nDado: %v %v \nEsperado: %v \nObtido: %v", slice1, slice2, esperado, obtido)
	}
}
