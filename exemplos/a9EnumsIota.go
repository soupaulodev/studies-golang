package exemplos

import "fmt"

type Cor int

const (
	Verde Cor = iota
	Azul
	Vermelho
)

var cor1 Cor

// Enums - Iota no GO
func EnumsIota() {
	cor1 = Azul
	fmt.Println(cor1)
}
