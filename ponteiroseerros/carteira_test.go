package ponteiroseerros

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Sacar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		erro := carteira.Sacar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
		confirmaErroInexistente(t, erro)
	})

	t.Run("Sacar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Sacar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)
		confirmaErro(t, erro, ErroSaldoInsuficiente)
	})
}

func confirmaSaldo(t *testing.T, carteira Carteira, esperado Bitcoin) {
	t.Helper()
	obtido := carteira.Saldo()

	if esperado != obtido {
		t.Errorf("\nEsperado: %v \nObtido: %v", esperado, obtido)
	}
}

func confirmaErroInexistente(t *testing.T, resultado error) {
	t.Helper()
	if resultado != nil {
		t.Fatal("Erro inesperado recebido.")
	}
}

func confirmaErro(t *testing.T, erro error, esperado error) {
	t.Helper()
	if erro == nil {
		t.Fatal("Esperava um erro, mas nenhum ocorreu.")
	}

	if erro != esperado {
		t.Errorf("\nResultado: %s \nEsperado: %s", erro, esperado)
	}
}
