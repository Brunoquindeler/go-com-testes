package estruturasmetodosinterfaces

import "testing"

func TestPerimetro(t *testing.T) {
	esperado := 40.0
	obtido := Perimetro(10.0, 10.0)

	if obtido != esperado {
		t.Errorf("\nEsperado: %.2f \nObtido: %.2f", esperado, obtido)
	}
}

func TestArea(t *testing.T) {
	esperado := 72.0
	obtido := Area(12.0, 6.0)

	if obtido != esperado {
		t.Errorf("\nEsperado: %.2f \nObtido: %.2f", esperado, obtido)
	}
}
