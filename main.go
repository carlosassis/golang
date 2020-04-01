package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

type verificarContaDeSaque interface {
	Sacar(valor float64) string
}

type verficarContasdeTransferencia interface {
	Sacar(valordeSaque float64) string
	Depositar(valorDeDeposito float64) string
	Obtersaldo() float64
}

func PagarBoleto(conta verificarContaDeSaque, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func TransferenciaHibrida(contaOrigem verficarContasdeTransferencia,
	contaDestino verficarContasdeTransferencia, valor float64) bool {

	if contaOrigem.Obtersaldo() >= valor && valor > 0 {
		contaOrigem.Sacar(valor)
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}

}

func main() {
	contaCarlos := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Carlos",
		CPF:       "12312312323",
		Profissao: "Desenvolvedor"},
		Agencia: "1233",
		Conta:   "23123233"}

	contaIsadora := contas.ContaPoupanca{Titular: clientes.Titular{
		Nome:      "Isadora",
		CPF:       "33333333222",
		Profissao: "Desenvolvedora"},
		Agencia:  "1244",
		Conta:    "23177233",
		Operacao: "1"}

	contaCarlos.Depositar(1000)
	contaIsadora.Depositar(2000)
	PagarBoleto(&contaCarlos, 900)
	PagarBoleto(&contaIsadora, 800)

	fmt.Println("Saldo de", contaCarlos.Titular.Nome, "é de", contaCarlos.Obtersaldo())
	fmt.Println("Saldo de", contaIsadora.Titular.Nome, "é de", contaIsadora.Obtersaldo())
	status := TransferenciaHibrida(&contaCarlos, &contaIsadora, 100)

	if status {
		fmt.Println("Transferencia realizada com sucesso")
		fmt.Println("Saldo de", contaCarlos.Titular.Nome, "é de", contaCarlos.Obtersaldo())
		fmt.Println("Saldo de", contaIsadora.Titular.Nome, "é de", contaIsadora.Obtersaldo())
	} else {
		fmt.Println("Erro ao realizar transferencia")
		fmt.Println("Saldo de", contaCarlos.Titular.Nome, "é de", contaCarlos.Obtersaldo())
		fmt.Println("Saldo de", contaIsadora.Titular.Nome, "é de", contaIsadora.Obtersaldo())
	}

}
