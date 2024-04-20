package alura_projeto_1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
	return
}
