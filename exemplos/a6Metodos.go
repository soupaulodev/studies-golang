package exemplos

import "fmt"

type Carro struct {
	Nome       string
	Velocidade int
}

var carro1 = Carro{
	Nome:       "Fusca",
	Velocidade: 0,
}

func UseMetodos() {
	carro1.acelerarInstancia()
	carro1.acelerarPonteiro()
}

func (c Carro) acelerarInstancia() {
	c.Velocidade = 20
	fmt.Println("Aceletando Instancia de Carro")
	fmt.Println(c)
}

func (c *Carro) acelerarPonteiro() {
	c.Velocidade = 30
	fmt.Println("Aceletando Ponteiro de Carro")
	fmt.Println(*c)
}
