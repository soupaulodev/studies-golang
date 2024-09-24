package exemplos

import "fmt"

// Loops no GO
func Loops() {

	// for normal
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for iterando sobre um array
	array := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	// for while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// for iterando com range
	for i, v := range array {
		fmt.Println(i, v)
	}

	// loop infinito
	for {
		fmt.Println("loop infinito")
		break
	}
}
