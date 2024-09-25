package exemplos

import (
	"fmt"
	"strings"
)

func Functions() {

	// Funcao declarada
	funcaoDeclarada()
	funcaoComParametro("Função com parametro")
	funcaoComVariosParametros("Função com vários parametros", "Função com vários parametros", 1)
	funcaoComParametroOpcional(1, "Função com parametro opcional")
	funcaoComRetorno("Função com retorno")
	funcaoComVariosRetornos("Função com vários retornos", "Função com vários retornos", 1)
	funcaoComRetornoNomeado()
	funcaoRetornandoOutraFuncao()()

	// Função anônima
	func() {
		println("Função anoônima")
	}()

	// Função lambda
	lambda := func() {
		println("Função lambda")
	}
	lambda()
}

// Função sem parametro ou retorno
func funcaoDeclarada() {
	fmt.Println("Função declarada")
}

// Função anônima
func _() {
	fmt.Println("Função anônima")
}

// Funcao com parametro
func funcaoComParametro(s string) {
	fmt.Printf("Função com parametro: %s\n", s)
}

// Funcao com vários parametros
func funcaoComVariosParametros(s1, s2 string, i int) {
	fmt.Printf("Função com vários parametros: %s, %s, %d\n", s1, s2, i)
}

// Função com parametro opcional 0 ou muitos
func funcaoComParametroOpcional(i int, s ...string) {
	fmt.Printf("Função com parametro opcional: %d, %s\n", i, strings.Join(s, ", "))
}

// Função com retorno
func funcaoComRetorno(s string) string {
	fmt.Printf("Função com retorno: %s\n", s)
	return s
}

// Função com vários retornos
func funcaoComVariosRetornos(s1, s2 string, i int) (string, string, int) {
	fmt.Printf("Função com vários retornos: %s, %s, %d\n", s1, s2, i)
	return s1, s2, i
}

// Funcao com retorno nomeado
func funcaoComRetornoNomeado() (s string) {
	s = "Função com retorno nomeado"
	fmt.Printf("Função com retorno nomeado: %s\n", s)
	return
}

// Funcao retornando outra função
func funcaoRetornandoOutraFuncao() func() {
	return func() {
		fmt.Println("Função retornando outra função")
	}
}
