package exemplos

import (
	"fmt"
	"sync"
	"time"
)

func GoroutineTeste() {
	goroutineTesteSimple()
	goutineTesteGerenciada()
}

func tarefaSimples() {
	for i := 0; i < 5; i++ {
		fmt.Println("Tarefa:", i)
		time.Sleep(time.Second) // Simula trabalho com um atraso
	}
}

func goroutineTesteSimple() {
	go tarefaSimples() // Inicia a goroutine

	// A função main continua executando
	for i := 0; i < 5; i++ {
		fmt.Println("Main:", i)
		time.Sleep(500 * time.Millisecond) // Simula trabalho com um atraso
	}

	// Aguardar um tempo para a goroutine concluir
	time.Sleep(6 * time.Second)
	fmt.Println("Programa concluído")
}

func tarefaGerenciada(wg *sync.WaitGroup, id int) {
	defer wg.Done() // Marca a tarefa como concluída quando retornar
	fmt.Printf("Iniciando tarefa %d\n", id)
	time.Sleep(2 * time.Second) // Simula trabalho
	fmt.Printf("Tarefa %d concluída\n", id)
}

func goutineTesteGerenciada() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)                   // Incrementa o contador do WaitGroup
		go tarefaGerenciada(&wg, i) // Inicia a goroutine
	}

	wg.Wait() // Aguarda todas as goroutines serem concluídas
	fmt.Println("Todas as tarefas concluídas.")
}
