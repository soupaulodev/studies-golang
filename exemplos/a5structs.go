package exemplos

import (
	"fmt"
)

type Pessoa struct {
	name string
}

// FuncaoOnStruct
func (p *Pessoa) FuncaoOnStruct() {
	fmt.Printf("Função on Struct: %s\n", p.name)
}

func UseStructs() {
	InstanciandoStructs()
	StructsAnonimos()
}

// Pessoa Explicita
var pessoaE Pessoa = Pessoa{
	name: "Pessoa Exp",
}

// Instanciando Structs
func InstanciandoStructs() {
	pessoaI := Pessoa{
		name: "Pessoa Impl",
	}

	pessoaI.name = "Pessoa Impl Edit"

	fmt.Printf("Struct Pessoa Explicita: %v\n", pessoaE)
	fmt.Printf("Struct Pessoa Implicita: %v\n", pessoaI)
	pessoaI.FuncaoOnStruct()
}

// Structs Anônimos
func StructsAnonimos() {
	pessoaA := struct {
		name string
	}{
		name: "Pessoa Anônima",
	}

	fmt.Printf("Struct Pessoa Anônima: %v\n", pessoaA)
}
