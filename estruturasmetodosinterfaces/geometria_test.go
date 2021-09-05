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
	t.Run("Retângulos", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		esperado := 72.0
		obtido := Area(retangulo)

		if obtido != esperado {
			t.Errorf("\nEsperado: %.2f \nObtido: %.2f", esperado, obtido)
		}
	})

	t.Run("Círculos", func(t *testing.T) {
		circulo := Circulo{10}
		esperado := 314.1592653589793
		obtido := Area(circulo)

		if obtido != esperado {
			t.Errorf("\nEsperado: %.2f \nObtido: %.2f", esperado, obtido)
		}
	})

}
