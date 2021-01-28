package main

import "fmt"

func main() {

	forzin(0, 6)
	forsera()

}

func forzin(i, num int) {
	for i <= num {
		fmt.Print("\nqual eras: ", i)
		i++
	}
}
func forsera() {
	for {
		fmt.Println("\n\n\nloop")
		break
	}
}
