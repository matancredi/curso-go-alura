package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaDaMariana := ContaCorrente{
		titular:       "Mariana",
		numeroAgencia: 532,
		numeroConta:   8963,
		saldo:         563.3,
	}

	contaDaBruna := ContaCorrente{
		"Bruna", 222, 333, 100.3}

	fmt.Println(contaDaMariana, contaDaBruna)

}
