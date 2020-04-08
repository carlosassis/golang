package contas

import (
	"github.com/carlos/git/golang/banco/clientes"
)

type ContaPoupanca struct {
	Titular                  clientes.Titular
	Agencia, Conta, Operacao string
	saldo                    float64
}

func (c *ContaPoupanca) Obtersaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) Sacar(saque float64) string {
	if c.saldo >= saque && saque > 0 {
		c.saldo -= saque
		return "Saque efetuado com sucesso"
	} else {
		return "Erro ao efetuar saque"
	}

}

func (c *ContaPoupanca) Depositar(valor float64) string {
	if valor > 0 {
		c.saldo += valor
		return "Deposito efetuado com sucesso"
	} else {
		return "Erro ao efetuar Deposito"
	}
}

func (c *ContaPoupanca) Transferir(valor float64, contaDestino *ContaPoupanca) bool {
	if c.saldo >= valor && valor > 0 {
		c.Sacar(valor)
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}
