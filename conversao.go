package main

import "fmt"

var a int = 30
var b string = "Nome"

func main() {

	var c float64 = float64(a) //colocar o TIPO(variavel)
	fmt.Printf("O valor de C é: %g  e o valor de B é: %s ", c, b)
}
