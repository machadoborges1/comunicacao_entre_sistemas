package main

import (
	"fmt"
	"time"
)

func processDataAsync(data string, resultChan chan string) {
	time.Sleep(2 * time.Second)
	resultChan <- "Processado: " + data
}

func main() {
	resultChan := make(chan string) // Canal para comunicação assíncrona

	fmt.Println("Iniciando processamento...")

	go processDataAsync("dados importantes", resultChan) // Chamada assíncrona

	fmt.Println("Fazendo outras tarefas enquanto os dados são processados...")

	result := <-resultChan // Aguardando o resultado da operação assíncrona

	fmt.Println(result)
	fmt.Println("Processamento concluído.")
}
