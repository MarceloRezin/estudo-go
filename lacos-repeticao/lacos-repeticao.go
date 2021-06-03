package main

import "fmt"

func main() {

	//Em go, tudo é for, inclusive while
	//Se não por nada no for, é um loop infinito
	//Pode ser usado continue e break
	
	//For 'comum'
	fmt.Println("For normal:");
	for i:=0; i<10; i++ {
		fmt.Println(i);
	}
	
	//While
	fmt.Println("\nWhile:");
	condicao	:= true
	count		:= 0
	for condicao {
		fmt.Println(count);
		count++
		if count == 5 {
			condicao = false;
		}
	}
		
	//For-each
	fmt.Println("\nFor-each:");
	strings	:= []string{"A", "B", "C", "D"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
