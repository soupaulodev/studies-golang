package modulos

import "fmt"

// Estruturas de Dados em GO
func EstruturasDeDados() {

	// Arrays
	array := [5]int{1, 2, 3, 4, 5}
	array[0] = 0

	// Slice
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)

	// Map
	mapa := make(map[string]int)
	mapa["key"] = 1

	fmt.Printf("Array: %v\n", array)
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Mapa: %v\n", mapa)
}
