package main

import (
	"fmt"
)

func main() {
	var x int
	var y float64

	fmt.Scanln(&x)
	fmt.Scanln(&y)

	total := float64(x) / y

	fmt.Print(total, " km/l")

}
