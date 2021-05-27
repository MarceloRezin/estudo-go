package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	// Scan le até o espaço ou a nova linha, o que vier primeiro
	// Scanln le até a nova linha, porém também filtra espaços, isso é próprio da implementação do go
	
	// Para ler strings completas, se usa inputReader
	
	inputReader := bufio.NewReader(os.Stdin)
	
	fmt.Printf("Informe seu nome: ")
	//fmt.Scanln(&nome)
	nome, _ := inputReader.ReadString('\n')
	
	fmt.Printf("Informe seu sobrenome: ")
	sobrenome, _ := inputReader.ReadString('\n')
	
	fmt.Println("Seu nome completo é:", strings.TrimSuffix(nome, "\n"), strings.TrimSuffix(sobrenome, "\n"))
	
	fmt.Printf("\nInforme o número de notas: ")
	
	var numeroNotas int
	fmt.Scanln(&numeroNotas)
	
	fmt.Println("Número de notas: ", numeroNotas)
	
	//Tamanho de array não pode usar valores de run time
	// Matriz tem tamanho fixo, fatia pode mudar, por isso se usa o make para criar um slice
	//var notas [numeroNotas]int -> Gera o erro: non-constant array bound
	
	notas	:= make([]int, numeroNotas)
	soma	:= 0	
	
	for i:=0; i<numeroNotas; i++ {
		fmt.Println("Informe a nota", i+1)
		fmt.Scanln(&notas[i])
		
		soma += notas[i]
	}
	
	fmt.Println("Notas informadas: ", notas)
	fmt.Println("Sua média:", soma/numeroNotas)
		
}
