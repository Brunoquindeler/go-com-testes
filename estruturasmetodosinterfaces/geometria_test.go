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

	cases := []struct {
		geometria Geometria
		esperado  float64
	}{
		{Retangulo{12, 6}, 72.0},
		{Circulo{10}, 314.1592653589793},
		{Triangulo{12, 6}, 36.0},
	}

	for _, tt := range cases {
		obtido := tt.geometria.Area()
		if obtido != tt.esperado {
			t.Errorf("\nEsperado: %.2f \nObtido: %.2f", tt.esperado, obtido)
		}
	}
}
