package main

import (
	"fmt"
	"math"
)

func main() {
	pi := 3.14159
	var r float64

	fmt.Scanln(&r)

	raiz := math.Pow(r, 3)

	volume := (4.0/3) * pi * raiz

	fmt.Printf("VOLUME = %.3f\n", volume)
		
}
