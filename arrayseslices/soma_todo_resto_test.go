package arrayseslices

import (
	"reflect"
	"testing"
)

func TestSomaTodoResto(t *testing.T) {

	verificaTotal := func(t *testing.T, slice1, slice2, esperado, obtido []int) {
		t.Helper()
		if !reflect.DeepEqual(esperado, obtido) {
			t.Errorf("\nDado: %v %v \nEsperado: %v \nObtido: %v", slice1, slice2, esperado, obtido)
		}
	}

	t.Run("Soma alguns slices", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{0, 9}

		esperado := []int{2, 9}

		obtido := SomaTodoResto(slice1, slice2)

		verificaTotal(t, slice1, slice2, esperado, obtido)
	})

	t.Run("Soma alguns slices vazios de forma segura", func(t *testing.T) {
		slice1 := []int{}
		slice2 := []int{0, 9, 2}

		esperado := []int{0, 11}

		obtido := SomaTodoResto(slice1, slice2)

		verificaTotal(t, slice1, slice2, esperado, obtido)

	})

}
