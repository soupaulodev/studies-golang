package exemplos

import (
	"fmt"
	"sync"
	"time"
)

// MutexTeste - mutual exclusion no GO
func MutexTeste() {

	/*
		É uma estrutura que permite a sincronização de acesso a dados compartilhados em
		um ambiente com concorrência. Ele é utilizado para garantir que apenas uma goroutine
		tenha acesso a uma seção crítica de código ou um dado compartilhado por vez,
		prevenindo condições de corrida (race conditions), que podem ocorrer quando várias goroutines
		tentam modificar uma variável ou estrutura de dados ao mesmo tempo.
	*/
	/*
		Implementação de mutex no GO é fornecido no pacote sync com o tipo Mutex.
		Contem dois métodos: Lock e Unlock
		- Lock(): bloqueia o mutex
		- Unlock(): desbloqueia o mutex
	*/

	casoMutex()

	/*
		Além do sync.Mutex, Go também fornece o sync.RWMutex, que permite múltiplas leituras simultâneas,
		mas escrita exclusiva. Isso é útil quando você tem muitos leitores e poucos escritores, e quer
		permitir o máximo de paralelismo possível para operações de leitura.

		- RLock(): Trava o mutex para leitura (várias goroutines podem fazer isso ao mesmo tempo).
		- RUnlock(): Destrava o mutex para leitura.
		- Lock(): Trava o mutex para escrita (somente uma goroutine pode acessar a seção crítica para escrita).
		- Unlock(): Destrava o mutex para escrita.
	*/

	casoRWMutex()
}

var (
	contador int        // Variável compartilhada
	mu       sync.Mutex // Declaração do mutex
)

func incrementar(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock() // Trava o mutex para acessar a variável

	// Evitando Deadlock
	// Um deadlock ocorre quando uma goroutine trava o mutex e nunca o destrava,
	// defer mu.Unlock() // Destrava o mutex

	contador++
	mu.Unlock() // Destrava o mutex

}

func casoMutex() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementar(&wg)
	}

	wg.Wait()
	fmt.Println("Contador final:", contador) // Saída esperada: 1000
}

var (
	contador1 int
	rwMu      sync.RWMutex
)

func leitor(wg *sync.WaitGroup) {
	defer wg.Done()

	rwMu.RLock()         // Trava para leitura (várias goroutines podem fazer isso ao mesmo tempo)
	defer rwMu.RUnlock() // Destrava após a leitura

	fmt.Println("Valor atual do contador:", contador1)
}

func escritor(wg *sync.WaitGroup) {
	defer wg.Done()

	rwMu.Lock()         // Trava para escrita (somente uma goroutine pode fazer isso)
	defer rwMu.Unlock() // Destrava após a escrita

	contador1++
	fmt.Println("Incrementando contador:", contador1)
}

func casoRWMutex() {
	var wg sync.WaitGroup

	// Iniciando 5 leitores
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go leitor(&wg)
	}

	time.Sleep(time.Millisecond)

	// Iniciando 2 escritores
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go escritor(&wg)
	}

	wg.Wait()
}
