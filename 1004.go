package main

import (
	"fmt"
)

func main() {
	var v1, v2 int
	fmt.Scanln(&v1)
	fmt.Scanln(&v2)

	PROD := v1 * v2

	fmt.Printf("PROD = %d\n", PROD)
}
