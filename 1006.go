package main

import (
	"fmt"
)

func main() {
	var A, B, C float64
	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Scanln(&C)

	MEDIA := ((A * 2) + (B * 3) + (C * 5)) / 10

	fmt.Printf("MEDIA = %.1f\n", MEDIA)
}
