package exemplos

import "fmt"

// Malloc no GO - Alocando memoria
func Malloc() {

	mallocNew()
	mallocMake()
	mallocPonteiro()
}

func mallocNew() {
	fmt.Println("mallocNew:")

	// Alocando memória para um inteiro e retornando um ponteiro
	p := new(int) // 'p' é um ponteiro para um int

	fmt.Println(*p) // Saída: 0 (valor zero do tipo int)

	// Alterando o valor do ponteiro
	*p = 42
	fmt.Println(*p) // Saída: 42
}

func mallocMake() {
	fmt.Println("mallocMake:")

	// Alocando memória e inicializando um slice de inteiros
	slice := make([]int, 5) // Slice com tamanho 5, todos elementos iniciados em 0

	fmt.Println(slice) // Saída: [0 0 0 0 0]

	// Alterando valores no slice
	slice[0] = 42
	fmt.Println(slice) // Saída: [42 0 0 0 0]
}

func mallocPonteiro() {
	fmt.Println("mallocPonteiro:")

	var x int = 10
	var p *int = &x // 'p' é um ponteiro para o endereço de 'x'

	fmt.Println("Valor de x:", x)    // Saída: 10
	fmt.Println("Endereço de x:", p) // Mostra o endereço de memória de 'x'

	// Modificando o valor de x via ponteiro
	*p = 20
	fmt.Println("Novo valor de x:", x) // Saída: 20
}
