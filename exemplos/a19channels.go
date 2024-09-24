package exemplos

import (
	"fmt"
	"time"
)

/* Características dos Channels
- Tipados: Cada channel é associado a um tipo específico. Isso significa que
você pode enviar apenas valores desse tipo através do channel.

- Bloqueio: Por padrão, a operação de envio em um channel bloqueia a goroutine até que
outro goroutine esteja pronto para receber o valor. Da mesma forma, a operação de recebimento
bloqueia até que um valor esteja disponível no channel.

- Direcionalidade: Você pode declarar channels como bidirecionais (podem enviar e receber) ou
unidirecionais (podem apenas enviar ou receber), permitindo um controle mais rigoroso sobre
como os dados são passados.

- Buffer: Channels podem ser bufferizados ou não. Um channel não bufferizado
(ou seja, com buffer de tamanho zero) bloqueia o envio até que haja um receptor pronto.
Um channel bufferizado permite que você envie um número especificado de valores antes de bloquear.
*/

// ChannelsTeste demonstra o uso de channels
func ChannelsTeste() {
	// Channel nao bufferizado
	c := make(chan int) // Cria um channel

	go produtor(c) // Inicia o produtor
	consumidor(c)  // Inicia o consumidor

	// Channel bufferizado
	cb := make(chan int, 2) // Channel bufferizado

	cb <- 1 // Não bloqueia
	cb <- 2 // Não bloqueia

	fmt.Println(<-cb) // Recebe 1
	fmt.Println(<-cb) // Recebe 2
}

func produtor(c chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Produzindo:", i)
		c <- i                      // Enviando dados para o channel
		time.Sleep(1 * time.Second) // Simula trabalho
	}
	close(c) // Fecha o channel quando terminar
}

func consumidor(c <-chan int) {
	for valor := range c { // Recebendo dados do channel
		fmt.Println("Consumindo:", valor)
	}
}

/* Definindo a Direção de um Channel

- Channel apenas para envio: Para um channel que só pode enviar dados, você usa a sintaxe chan<- Tipo ao declarar a função.

- Channel apenas para recebimento: Para um channel que só pode receber dados, você usa a sintaxe <-chan Tipo.

- Channel bidirecional: Para um channel que pode enviar e receber dados, você simplesmente usa chan Tipo.
*/
