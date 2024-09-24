package exemplos

import "fmt"

// Condicionais no GO
func Condicionais() {

	// if, else if e else
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("i == 5")
		} else if i == 6 {
			fmt.Println("i == 6")
		} else {
			fmt.Println("i != 5 && i != 6")
		}
	}

	// switch
	for i := 0; i < 10; i++ {
		switch i {
		case 5:
			fmt.Println("i == 5")
		case 6:
			fmt.Println("i == 6")
		default:
			fmt.Println("i != 5 && i != 6")
		}
	}
}
