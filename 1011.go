package main

import (
	"fmt"
)

func main() {
	pi := 3.14159
	var r float64

	fmt.Scanln(&r)

	raiz := (math.pow(r, 3))
	volume := (float64(4.0/3.0) * pi * (raiz))

	fmt.Printf("VOLUME = %.3f\n", volume)

}
