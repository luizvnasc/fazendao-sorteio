package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func carregaPoteComTimes(path string) (pote []string, err error) {
	arquivo, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	pote = make([]string, 0)
	reader := bufio.NewReader(arquivo)
	for {
		time, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		pote = append(pote, time)
	}
	return pote, nil
}

func chacoalhaOPote(pote []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pote), func(i, j int) { pote[i], pote[j] = pote[j], pote[i] })
	return pote
}

func pegaOTimeDoPote(pote []string) (string, []string) {
	return pote[0], pote[1:]
}

func main() {
	pote1, err := carregaPoteComTimes("pote1.txt")
	if err != nil {
		log.Fatalf("Erro ao carregar pote 1: %s", err)
	}
	pote2, err := carregaPoteComTimes("pote2.txt")
	if err != nil {
		log.Fatalf("Erro ao carregar pote 2: %s", err)
	}

	if len(pote1) != len(pote2) {
		log.Fatalf("Tamanho dos potes divergem.")
	}

	confronto := 1
	for len(pote1) > 0 {
		pote1 = chacoalhaOPote(pote1)
		pote2 = chacoalhaOPote(pote2)

		var time1, time2 string
		time1, pote1 = pegaOTimeDoPote(pote1)
		time2, pote2 = pegaOTimeDoPote(pote2)
		fmt.Printf("Confronto #%d: %s X %s\n", confronto, time1, time2)
		confronto++
	}

}
