package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
)

func main() {
	var linhas [2]string
	
	linhas[0] = "Seu nome é:"
	linhas[1] = "Marcelos"
	
	const dir = "meu-arquivo.txt"
	arquivo, err := os.Create(dir);
	if err != nil {
		log.Fatalf("Erro: ", err);
	}
	
	//Garante que o arquivo vai ser fechado no final
	defer arquivo.Close();
	
	escritor := bufio.NewWriter(arquivo);
	for _, linha := range linhas {
		fmt.Fprintln(escritor, linha);
	}
	
	escritor.Flush();
	
	arquivoLido, err := os.Open(dir);
	
	if err != nil {
		log.Fatalf("Erro: ", err);
	}
	
	defer arquivoLido.Close();
	
	var linhasLidas []string
	scanner := bufio.NewScanner(arquivoLido)
	for scanner.Scan() {
		linhasLidas = append(linhasLidas, scanner.Text())
	}
	
	fmt.Println(linhasLidas)
}
