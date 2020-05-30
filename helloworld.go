// Para compilar: go build helloworld.go - Gera .exe
// Para compilar e executar: go run helloworld.go

// Convenções:
// Não é necessário colocar ponto e vírgula
// Chave de abertura da função fica na mesma linha da função

package main

import "fmt"

func main() {
	//Funções começando com letra maiúscula são de outros pacotes
	fmt.Println("Hello world")

}
