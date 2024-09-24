package exemplos

import (
	"fmt"
)

// Declaracao explicita de variaveis
var vareA int = 10
var vareB string = "Ola"

// Declaracao de constante
const PI float32 = 3.14

// DeclaracaoDeVariaveis no GO
func DeclaracaoDeVariaveis() {

	// Declaracao de constante
	const PERCENT float32 = 100.0

	// Declaração implicita de variáveis
	variA := 10
	variB := "Ola"

	fmt.Printf("Var: variA, Tipo: %T, Valor: %v\n", variA, variA)
	fmt.Printf("Var: variB, Tipo: %T, Valor: %v\n", variB, variB)
	fmt.Printf("Var: vareA, Tipo: %T, Valor: %v\n", vareA, vareA)
	fmt.Printf("Var: vareB, Tipo: %T, Valor: %v\n", vareB, vareB)

	fmt.Printf("Var: PI, Tipo: %T, Valor: %v\n", PI, PI)
	fmt.Printf("Var: PERCENT, Tipo: %T, Valor: %v\n", PERCENT, PERCENT)
}
