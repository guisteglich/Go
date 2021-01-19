package main

import "fmt"

func main() {
	var cod1, qtd1, cod2, qtd2 int
	var valor1, valor2 float64

	fmt.Scanln(&cod1, &qtd1, &valor1)
	fmt.Scanln(&cod2, &qtd2, &valor2)

	total := (float64(qtd1) * valor1) + (float64(qtd2) * valor2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
