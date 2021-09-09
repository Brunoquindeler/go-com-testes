package ponteiroseerros

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(valor Bitcoin) {
	c.saldo += valor
}

var ErroSaldoInsuficiente = errors.New("Não é possível sacar: Saldo insuficiente")

func (c *Carteira) Sacar(valor Bitcoin) error {
	if valor > c.saldo {
		return ErroSaldoInsuficiente
	}

	c.saldo -= valor
	return nil
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}
