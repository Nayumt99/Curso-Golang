// declarar meu pacote principal
package main

//importar função fmt
import "fmt"

//declaração da variável CONST do ponto de ebulição da água em F.
const ebulicaoF float64 = 212.0

//função principal
func main() {

	var tempF float64 = ebulicaoF            //variável do VALOR da temperatura em graus F
	var tempC float64 = (tempF - 32) * 5 / 9 //conversão de F para C
	//apareça na tela o resultado
	fmt.Println("A temperatura de ebulição da água em graus °F é", tempF)
	fmt.Println("A temperatura de ebilução da água em graus °C é", tempC)
	// A gente espera que apareça F = 212 e C = 100

}
