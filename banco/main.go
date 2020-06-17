package main

import (
	"curso-go-alura/banco/contas"
	"fmt"
)

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

	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDaSilvia.Tranferir(200, &contaDoGustavo)

	fmt.Println(status)

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
