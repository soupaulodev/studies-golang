package exemplos

import (
	"fmt"
	"strings"
)

// StringTeste trabalhando com string em GO
func StringTeste() {

	// Strings Concatenadas
	fmt.Println("Ola" + " " + "Mundo")

	// StringBuilder
	var sb strings.Builder
	sb.WriteString("Ola")
	sb.WriteString(" ")
	sb.WriteString("Mundo")
	fmt.Println(sb.String())

	// Strings.Split
	s := "Ola Mundo"
	split := strings.Split(s, " ")
	fmt.Println(split[0], split[1])
}
