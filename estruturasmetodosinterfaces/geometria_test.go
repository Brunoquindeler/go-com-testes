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
		nome      string
		geometria Geometria
		esperado  float64
	}{
		{nome: "Retângulo", geometria: Retangulo{Largura: 12, Altura: 6}, esperado: 72.0},
		{nome: "Círculo", geometria: Circulo{Raio: 10}, esperado: 314.1592653589793},
		{nome: "Triângulo", geometria: Triangulo{Base: 12, Altura: 6}, esperado: 36.0},
	}

	for _, tt := range cases {
		t.Run(tt.nome, func(t *testing.T) {
			obtido := tt.geometria.Area()
			if obtido != tt.esperado {
				t.Errorf("\nGeometria: %#v \nEsperado: %.2f \nObtido: %.2f", tt.geometria, tt.esperado, obtido)
			}
		})
	}
}
