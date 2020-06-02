package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

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
			imprimeLogs()
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
	// site := "https://www.alura.com.br"
	// site := "https://random-status-code.herokuapp.com"
	// sites := []string{"https://random-status-code.herokuapp.com",
	// 	"https://www.alura.com.br",
	// 	"https://www.caelum.com.br",
	// }

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu erro: ", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	// Retornará vetor (isso parece java)
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	// sem converter para string, vem array de bytes
	// fmt.Println(string(arquivo))
	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05 - ") + site + " - online:" + strconv.FormatBool(status) + "\n")
	arquivo.Close()

}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

}

// func exibeNomes() {
// 	// Slices
// 	nomes := []string{"Douglas", "Daniel", "Bernardo"}
// 	fmt.Println(nomes)
// 	nomes = append(nomes, "Mariana")
// 	fmt.Println(len(nomes))
// }
