package main

import (
	"curso-go-alura/banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	// contaDaMariana := ContaCorrente{
	// 	titular:       "Mariana",
	// 	numeroAgencia: 532,
	// 	numeroConta:   8963,
	// 	saldo:         563.3,
	// }

	// contaDaBruna := ContaCorrente{
	// 	"Bruna", 222, 333, 100.3}

	// fmt.Println(contaDaMariana, contaDaBruna)

	// var contadaMagali *ContaCorrente
	// contadaMagali = new(ContaCorrente)
	// contadaMagali.titular = "Magali"

	// // Conteúdo da conta da Magali
	// fmt.Println(*contadaMagali)

	// // Endereço da conta da Magali
	// fmt.Println(&contadaMagali)

	// contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	// contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	// status := contaDaSilvia.Tranferir(200, &contaDoGustavo)

	// fmt.Println(status)

	// fmt.Println(contaDaSilvia)
	// fmt.Println(contaDoGustavo)

	// contaDoBruno := contas.ContaCorrente{
	// 	Titular: clientes.Titular{
	// 		Nome:      "Bruno",
	// 		CPF:       "012.345.678-90",
	// 		Profissao: "Analista de Sistemas",
	// 	},
	// 	NumeroAgencia: 123,
	// 	NumeroConta:   456,
	// 	Saldo:         100,
	// }

	// clienteBruno := clientes.Titular{
	// 	Nome:      "Bruno",
	// 	CPF:       "012.345.678-90",
	// 	Profissao: "Analista de Sistemas",
	// }

	// contaDoBruno := contas.ContaCorrente{
	// 	clienteBruno,
	// 	123,
	// 	456,
	// 	100,
	// }

	// fmt.Println(contaDoBruno)

	// contaExemplo := contas.ContaCorrente{}
	// contaExemplo.Depositar(100)

	// fmt.Println(contaExemplo.ObterSaldo())

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	contaDaPaty := contas.ContaCorrente{}
	contaDaPaty.Depositar(500)
	PagarBoleto(&contaDaPaty, 400)

	fmt.Println(contaDoDenis)
	fmt.Println(contaDaPaty)
}
