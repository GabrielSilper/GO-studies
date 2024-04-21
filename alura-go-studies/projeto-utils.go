package alura_projeto_1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func ReadFiles() (sites []string, err error) {
	arquivo, err := os.Open("./alura-go-studies/sites.txt")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	reader := bufio.NewReader(arquivo)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return
}

func registerLog(site string, status bool) {
	arquivo, err := os.OpenFile("./alura-go-studies/logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {
		fmt.Println(err)
	}

	line := fmt.Sprintf("%s - online: %t\n", site, status)
	arquivo.WriteString(line)

	arquivo.Close()
}

func addTimeToLog() {
	arquivo, err := os.OpenFile("./alura-go-studies/logs.txt", os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString("time: " + time.Now().Format("02/01/2006 - 15:04:05 - Monday"))
	arquivo.WriteString("\n--------------------\n")

	arquivo.Close()
}

func showLog() {
	arquivo, err := os.ReadFile("./alura-go-studies/logs.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
