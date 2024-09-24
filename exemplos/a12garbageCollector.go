package exemplos

import (
	"fmt"
	"runtime"
)

// GarbageCollector no GO
func GarbageCollector() {

	// Forcando garbage collector
	// Idealmente utilizado em casos de pre-benchmarks,liberação de memória previsível
	// E teste verificando se o GC está liberando a memória corretamente
	runtime.GC()
	fmt.Println("Exemplo garbage collector")
}
