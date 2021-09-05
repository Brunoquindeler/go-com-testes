package estruturasmetodosinterfaces

import "testing"

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	esperado := 40.0
	obtido := Perimetro(retangulo)

	if obtido != esperado {
		t.Errorf("\nEsperado: %.2f \nObtido: %.2f", esperado, obtido)
	}
}

func TestArea(t *testing.T) {
	verificaArea := func(t *testing.T, geometria Geometria, esperado float64) {
		t.Helper()
		obtido := geometria.Area()

		if obtido != esperado {
			t.Errorf("\nEsperado: %.2f \nObtido: %.2f", esperado, obtido)
		}
	}

	t.Run("Retângulos", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		verificaArea(t, retangulo, 72.0)
	})

	t.Run("Círculos", func(t *testing.T) {
		circulo := Circulo{10}
		verificaArea(t, circulo, 314.1592653589793)
	})
}
