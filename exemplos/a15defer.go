package exemplos

import "fmt"

func DeferTeste() {
	fmt.Println("Início")

	// Adia a execução até o final da função
	defer fmt.Println("Defer: Executado no final")

	fmt.Println("Fim")
}
