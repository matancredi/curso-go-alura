package main

import (
	"fmt"
)

func main() {
	// var nome = "Mariana"
	// var versao = 1.1
	// var idade = 23

	nome := "Mariana"
	versao := 1.1
	//idade := 23

	// Se não atribuísse valor na variável, ela inicializaria com 0 (para int)
	// var idade int

	//fmt.Println("Olá,", nome, "sua idade é", idade)
	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão", versao)

	// fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
	// fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
	// fmt.Println("O tipo da variável versao é", reflect.TypeOf(versao))

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int

	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)
	fmt.Println("O endereço da variável comando é", &comando)

	fmt.Println("O comando escolhido foi", comando)

	// If não tem parênteses
	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo logs")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa")
	// } else {
	// 	fmt.Println("Não conheço esse comando")
	// }

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs")
	case 0:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Desconheço esse comando")

	}

}
