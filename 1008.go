package main

import (
	"fmt"
)

func main() {
	var NF, HT int
	var VH float64
	fmt.Scanln(&NF)
	fmt.Scanln(&HT)
	fmt.Scanln(&VH)

	S := (float64(HT) * VH)

	fmt.Printf("NUMBER = %d\n", NF)
	fmt.Printf("SALARY = U$ %.2f\n", S)
}
