package main

import "fmt"

func main() {
	nome := "Monitor de sites"
	versao := 1.1

	fmt.Println("Iniciando o programa", nome, "versão", versao)
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	var opcao int
	fmt.Scan(&opcao)

	fmt.Println("A opção escolhida foi", opcao)
}
