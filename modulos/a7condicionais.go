package modulos

import "fmt"

// Condicionais no GO
func Condicionais() {

	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("i == 5")
		} else if i == 6 {
			fmt.Println("i == 6")
		} else {
			fmt.Println("i != 5 && i != 6")
		}
	}
}
