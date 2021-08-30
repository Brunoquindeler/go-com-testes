package arrayseslices

import (
	"reflect"
	"testing"
)

func TestSomaTodoResto(t *testing.T) {
	slice1 := []int{1, 2}
	slice2 := []int{0, 9}

	esperado := []int{2, 9}

	obtido := SomaTodoResto(slice1, slice2)

	if !reflect.DeepEqual(esperado, obtido) {
		t.Errorf("\nDado: %v %v \nEsperado: %v \nObtido: %v", slice1, slice2, esperado, obtido)
	}

}
