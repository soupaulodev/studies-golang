package exemplos

import (
	"context"
	"fmt"
	"time"
)

func ContextTeste() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // Define um timeout
	defer cancel()                                                          // Garante que o cancelamento é chamado ao final

	go longaOperacao(ctx) // Inicia a operação longa

	time.Sleep(2 * time.Second) // Simula outro trabalho
	fmt.Println("Main: esperando 2 segundos antes do cancelamento")

	// O timeout ocorre após 3 segundos, então a operação longa será cancelada
	time.Sleep(4 * time.Second)
	fmt.Println("Main: concluído")
}

func longaOperacao(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): // Simula uma operação longa
		fmt.Println("Operação concluída")
	case <-ctx.Done(): // Aguarda o cancelamento do contexto
		fmt.Println("Operação cancelada:", ctx.Err())
	}
}

/*
// O pacote context em Go é uma ferramenta essencial para gerenciar deadlines, cancelamentos
e valores de request scopes em aplicações concorrentes. Ele é especialmente útil em cenários onde você
precisa passar informações que controlam o comportamento de goroutines, como tempo de espera e cancelamento de operações.

// Principais Conceitos do Pacote context
- Contexto (Context): Um Context carrega informações sobre um contexto específico,
como cancelamento e deadlines, e pode ser passado através de funções e goroutines.

- Cancelamento: O Context pode ser usado para cancelar operações que estão em andamento,
liberando recursos e evitando que goroutines continuem a execução desnecessariamente.

- Deadlines: Você pode definir um prazo (deadline) para operações que devem ser concluídas até um certo momento.

- Valores: O Context pode ser usado para passar valores entre funções, embora você deva evitar usar
o contexto para transportar dados em vez de valores.


// Tipos de Context - Existem três tipos principais de contexto no pacote context:

- context.Background(): Um contexto vazio, geralmente usado como o ponto de partida para outros contextos.
É um contexto base que não tem deadline, cancelamento ou valores.

- context.TOD0(): Usado quando você não tem certeza de qual contexto usar.
É um marcador que indica que você deve decidir posteriormente qual contexto será mais apropriado.

- context.WithCancel(parent Context): Cria um novo contexto que pode ser cancelado.
O cancelamento do contexto pai também cancela este novo contexto.

- context.WithDeadline(parent Context, deadline time.Time): Cria um novo contexto com um prazo específico.
Se o prazo expirar, as operações que dependem desse contexto podem ser canceladas.

- context.WithTimeout(parent Context, timeout time.Duration): Similar ao WithDeadline, mas define um prazo
relativo ao invés de um específico.

- context.WithValue(parent Context, key, val): Cria um novo contexto que carrega um valor
associado a uma chave específica.
*/
