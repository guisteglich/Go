package main

import "fmt"

func main() {
	var A, B, C float64
	pi := 3.14159

	fmt.Scanln(&A, &B, &C)

	triangulo := (A * C) / 2
	circulo := pi * (C * *2)
	trapezio := ((A + B) * C) / 2
	quadrado := B * *2
	retangulo := A * B

	fmt.Printf("TRIANGULO: %.3f\n", triangulo)
	fmt.Printf("CIRCULO: %.3f\n", circulo)
	fmt.Printf("TRAPEZIO: %.3f\n", trapezio)
	fmt.Printf("QUADRADO: %.3f\n", quadrado)
	fmt.Printf("RETANGULO: %.3f\n", retangulo)

}
