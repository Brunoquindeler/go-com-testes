package ponteiroseerros

import "testing"

func TestCarteira(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(10)

	obtido := carteira.Saldo()
	esperado := 10

	if esperado != obtido {
		t.Errorf("\nDado: %v \nEsperado: %v \nObtido: %v", carteira, esperado, obtido)
	}

}
