package main

import "strings"

func main() {
	var a int = 1
	var b bool = true
	var c string = "Ola"
	var d float32 = 3.14
	var e byte = 'a'
	var f rune = 'b'
	var g complex64 = 2 + 3i
	var h uint = 10
	var i uint8 = 20
	var j uint16 = 30
	var k uint32 = 40
	var l uint64 = 50

	println(a, b, c, d, e, f, g, h, i, j, k, l)

	// Exemplo de variaveis final
	finalA := 10
	finalB := "Ola"
	print(finalA, finalB)

	// Exemplo de constante
	const PI float32 = 3.14
	print(PI)

	// Exemplo de estruturas de dados vetores
	var lista []int
	lista = append(lista, 1)
	print(lista)

	// Exemplo de maps
	var m map[string]int
	m = make(map[string]int)
	m["a"] = 1
	print(m["a"])

	// Exemplos de encapsulamento
	private := "Privada"
	print(private)

	// Exemplos de interfaces
	var i1 interface{}
	i1 = 1
	print(i1)

	// Exemplos de ponteiros
	var p *int
	p = new(int)
	*p = 1
	print(*p)

	// Exemplos de condicionais
	var x int
	if x == 1 {
		print("x == 1")
	} else if x == 2 {
		print("x == 2")
	} else {
		print("x != 1 && x != 2")
	}

	// Exemplo de loops
	for i := 0; i < 10; i++ {
		print(i)
	}

	// Exemplo de loops iterando sobre um array
	array := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(array); i++ {
		print(array[i])
	}

	// Exemplo de loops infinitos
	for {
		print("loop infinito")
		break
	}

	// Exemplo loop while
	i = 0
	for i < 10 {
		print(i)
		i++
	}

	// Exemplo de loop range
	for i, v := range array {
		print(i, v)
	}

	// Exemplo de structs
	type Person struct {
		name string
		age  int
	}
	var p1 Person
	p1.name = "John"
	p1.age = 30
	print(p1.name, p1.age)

	// Exemplo de enums
	type Color int
	const (
		Red Color = iota
		Green
		Blue
	)
	var cor Color
	cor = Red
	print(cor)

	// Exemplo de arrays multidimensionais
	var matrix [3][3]int
	matrix[0][0] = 1
	print(matrix[0][0])

	// Exemplo Strings Concatenadas
	print("Ola" + " " + "Mundo")

	// StringBuilder
	var sb strings.Builder
	sb.WriteString("Ola")
	sb.WriteString(" ")
	sb.WriteString("Mundo")
	print(sb.String())

	// Strings.Split
	s := "Ola Mundo"
	split := strings.Split(s, " ")
	print(split[0], split[1])

	// Strings.Join
	s2 := []string{"Ola", "Mundo"}
	join := strings.Join(s2, " ")
	print(join)

	// Exemplo de Funções anonimas
	func() {
		print("Ola Mundo")
	}()

	// Exemplo funcoes anonimas com argumentos
	func(s string, x int) {
		print(s, x)
	}("Ola Mundo", 1)

	// Exemplo de defer
	// Defer é utilizado para finalizar um bloco de código
	defer print("Ola Mundo")

	// Exemplo de chamadas de funções
	functionDeclarada("Ola Mundo")

	// Exemplo switch-case
	switch cor {
	case Red:
		print("Red")
	case Green:
		print("Green")
	case Blue:
		print("Blue")
	}
}

// Exemplo de metodo com argumento e retorno
func functionDeclarada(s string) string {
	return s
}