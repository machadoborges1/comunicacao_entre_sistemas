package main

import (
	"fmt"
	"time"
)

func processData(data string) string {
	// Simula o processamento de dados
	time.Sleep(2 * time.Second) // Espera 2 segundos para simular um trabalho pesado
	return "Processado: " + data
}

func main() {
	fmt.Println("Iniciando processamento...")

	result := processData("dados importantes") // Chamada síncrona

	fmt.Println(result) // A execução do programa é bloqueada até que o processData retorne
	fmt.Println("Processamento concluído.")
}
