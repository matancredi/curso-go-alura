package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	// Loop infinito
	for {
		exibeMenu()
		comando := leComando()
		// Não é necessário colocar o break:
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Desconheço esse comando")
			os.Exit(-1)
		}
	}
}

// var nome = "Mariana"
// var versao = 1.1
// var idade = 23

// nome := "Mariana"
// versao := 1.1
//idade := 23

// Se não atribuísse valor na variável, ela inicializaria com 0 (para int)
// var idade int

//fmt.Println("Olá,", nome, "sua idade é", idade)
// fmt.Println("Olá,", nome)
// fmt.Println("Este programa está na versão", versao)

// fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
// fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
// fmt.Println("O tipo da variável versao é", reflect.TypeOf(versao))

//var comando int

//fmt.Scanf("%d", &comando)
//fmt.Scan(&comando)
//fmt.Println("O endereço da variável comando é", &comando)

//fmt.Println("O comando escolhido foi", comando)

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

func exibeIntroducao() {
	nome := "Mariana"
	versao := 1.1
	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão", versao)
}

// Função retornará int
func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//site := "https://www.alura.com.br"
	site := "https://random-status-code.herokuapp.com"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
	}
}
