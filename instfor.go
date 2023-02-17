package main

import "fmt"

func main() {
	//	se i%2 ==0
	//		imprimir: número par
	//	se não
	//		imprimir: número impar
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Ímpar")
		}
		fmt.Println(i)
	}
}

//SE o número é PAR ou ÍMPAR
