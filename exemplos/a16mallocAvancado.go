package exemplos

import (
	"fmt"
	"sync"
)

// MallocAvancadoTeste exemplo de uso
func MallocAvancadoTeste() {

	// Fornece mais controlesobre o comportamento de alocação de memória
	// Como otimizações de uso de heap e stack

	// Técnica Pooling de Objetos: Usando sync.Pool para reutilizar objetos e evitar alocações desnecessárias.
	poolingTeste()

	//		Técnica Escopo de Alocação: Variáveis alocadas dentro de uma função podem ser armazenadas no stack (rápido)
	//		se não forem retornadas como ponteiros, ou podem ser movidas para o heap (lento)
	//		se retornadas ou capturadas por goroutines.
	escopoDeAlocacaoTeste()

}

func poolingTeste() {
	var pool sync.Pool
	pool.New = func() interface{} {
		return 0
	}

	// Pegando um valor do pool
	value := pool.Get().(int)
	fmt.Println("Valor retirado do pool:", value) // Saída: 0

	// Devolvendo o valor ao pool
	pool.Put(10)

	// Pegando um valor do pool novamente
	value = pool.Get().(int)
	fmt.Println("Valor retirado do pool:", value) // Saída: 10
}

func escopoDeAlocacaoTeste() {
	a := pilha()                  // a recebe o valor 10
	fmt.Println("Valor de a:", a) // Saída: Valor de a: 10

	b := heap()                    // b recebe o ponteiro para o inteiro no heap
	fmt.Println("Valor de b:", *b) // Saída: Valor de b: 20

	// Para liberar a memória alocada no heap, o Garbage Collector cuidará disso automaticamente.
}

// Função que retorna um inteiro alocado na pilha
func pilha() int {
	x := 10 // x é alocado na pilha
	return x
}

// Função que retorna um ponteiro para um inteiro alocado no heap
func heap() *int {
	y := new(int) // y é alocado no heap
	*y = 20
	return y
}

// Anotações finais
// - Evitar Alocações Desnecessárias no Heap: Se uma variável não precisa sobreviver após a execução da função,
// evite alocá-la no heap, pois isso introduz a sobrecarga do Garbage Collector.

// Tipagem de Variáveis: O compilador do Go pode decidir automaticamente onde alocar a variável com base
// na forma como ela é utilizada. Se uma variável é retornada de uma função e não é usada como ponteiro,
// ela pode ser alocada na pilha.
