package contas

import (
	"banco/clientes"
)

type ContaCorrente struct {
	Titular        clientes.Titular
	Agencia, Conta string
	saldo          float64
}

func (c *ContaCorrente) Obtersaldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) Sacar(saque float64) string {
	if c.saldo >= saque && saque > 0 {
		c.saldo -= saque
		return "Saque efetuado com sucesso"
	} else {
		return "Erro ao efetuar saque"
	}

}

func (c *ContaCorrente) Depositar(valor float64) string {
	if valor > 0 {
		c.saldo += valor
		return "Deposito efetuado com sucesso"
	} else {
		return "Erro ao efetuar Deposito"
	}
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if c.saldo >= valor && valor > 0 {
		c.Sacar(valor)
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}
