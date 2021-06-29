package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	start()

	for {
		option := selectedOption()

		switch option {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Logs...")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Opção não reconhecida!!")
			os.Exit(-1)
		}
	}
}

func start() {
	nome := "Monitor de sites"
	versao := 1.1
	fmt.Println("Iniciando o programa", nome, "versão", versao)
}

func selectedOption() int {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	var opcao int
	fmt.Scan(&opcao)

	return opcao
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	url := "https://www.alura.com.br"
	response, _ := http.Get(url)

	if response.StatusCode == 200 {
		fmt.Println("Site:", url, "Ok!")
	} else {
		fmt.Println("Site:", url, "Falhou!")
	}
}
