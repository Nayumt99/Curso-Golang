package main

import "fmt"

const ebulicaoK float64 = 373.0

func main() {

	tempK := ebulicaoK
	tempC := (tempK - 273)

	fmt.Println("A temperatura de ebulição da água em graus °K é", tempK)
	fmt.Println("A temperatura de ebilução da água em graus °C é", tempC)
}
