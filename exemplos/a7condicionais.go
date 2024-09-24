package exemplos

import "fmt"

// Condicionais no GO
func Condicionais() {

	// if, else if e else
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Printf("%v == 5\n", i)
		} else if i == 6 {
			fmt.Printf("%v == 6\n", i)
		} else {
			fmt.Printf("%v != 5 && %v != 6\n", i, i)
		}
	}

	// switch
	for i := 0; i < 10; i++ {
		switch i {
		case 5:
			fmt.Printf("%v == 5\n", i)
		case 6:
			fmt.Printf("%v == 6\n", i)
		default:
			fmt.Printf("%v != 5 && %v != 6\n", i, i)
		}
	}
}
