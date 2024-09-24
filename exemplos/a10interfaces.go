package exemplos

import "fmt"

// Aviao interface
type Aviao interface {
	Voar()
	Decolar(velocidade int) bool
	Pousar()
}

// AviaoPadrao struct que implementa a interface Aviao
type AviaoPadrao struct {
	Aviao
}

// Decolar implementa a interface Aviao para a struct AviaoPadrao
func (a AviaoPadrao) Decolar(velocidade int) bool {
	return true
}

// Airbus struct
type Airbus struct {
	modelo string
}

// Implementação dos métodos da interface Aviao para a struct Airbus
// Voar implementa a interface Aviao para a struct Airbus
func (a Airbus) Voar() {
	fmt.Println(a.modelo, "está voando.")
}

// Decolar implementa a interface Aviao para a struct Airbus
func (a Airbus) Decolar(velocidade int) bool {
	if velocidade > 250 {
		fmt.Println(a.modelo, "decolou com sucesso!")
		return true
	}
	fmt.Println(a.modelo, "não conseguiu decolar.")
	return false
}

// Pousar implementa a interface Aviao para a struct Airbus
func (a Airbus) Pousar() {
	fmt.Println(a.modelo, "pousou com segurança.")
}

// TesteAviao Função que aceita qualquer objeto que implemente a interface Aviao
func TesteAviao(a Aviao) {
	a.Decolar(300)
	a.Voar()
	a.Pousar()
}

func TestarAviao() {
	aviao := AviaoPadrao{Aviao: Airbus{modelo: "Airbus"}}
	TesteAviao(aviao)
}
