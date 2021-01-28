package main

import (
	"fmt"
)
 
func main() {

	var A, B, C float64
	
	fmt.Scanln(&A, &B, &C)
	
	if A > C; A > B {
		maior := A
	}else if A < B; B > C {
		maior := B
	}
	maior := C

	fmt.Print("%f eh o maior", maior)

}
