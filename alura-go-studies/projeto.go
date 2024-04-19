package alura_projeto_1

import (
	"fmt"
	"net/http"
	"os"
)

func ProjetoExec() {
	showIntro()
	for {
		comando := readOption()

		switch comando {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse programa")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	name := "Gabriel"
	version := 1.1
	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func readOption() int {
	fmt.Println("\n1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	fmt.Scan(&comando)

	return comando
}

func startMonitoring() {
	fmt.Println("Monitorando...")
	site := "http://localhost:3001/customers"

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
