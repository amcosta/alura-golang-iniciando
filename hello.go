package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

const inputSites = "sites.txt"
const logFile = "log.txt"

func main() {
	nome := "Monitor de sites"
	versao := 1.1
	fmt.Println("Iniciando o programa", nome, "versão", versao)

	for {
		option := selectedOption()

		switch option {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Opção não reconhecida!!")
			os.Exit(-1)
		}
	}
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
	file, err := os.Open(inputSites)
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		site := reader.Text()

		response, err := http.Get(site)
		if err != nil {
			fmt.Println(err)
		}

		if response.StatusCode == 200 {
			writeLog(site, true)
		} else {
			writeLog(site, false)
		}
	}

	file.Close()
}

func writeLog(site string, online bool) {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	line := "[" + date + "] " + site + " - Online: " + strconv.FormatBool(online) + "\n"
	file.WriteString(line)

	file.Close()
}

func showLogs() {
	file, err := os.Open(logFile)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	file.Close()
}
