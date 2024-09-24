package exemplos

import "fmt"

// alterarValor função que altera o valor utilizando um ponteiro
func alterarValor(v *int) {
	*v = 20 // Modificando o valor na memória para onde o ponteiro aponta
}

// TestePonteiros demonstra o uso de ponteiros
func TestePonteiros() {
	valor := 10

	fmt.Println("Valor original:", valor)

	// Alterando o valor utilizando um ponteiro passando o endereço da variável como argumento
	alterarValor(&valor)

	// Imprimindo o valor após a alteração
	fmt.Println("Valor alterado:", valor)
}
