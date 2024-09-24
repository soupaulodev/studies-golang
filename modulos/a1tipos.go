package modulos

import (
	"fmt"
)

// Tipos primitivos no GO
func Tipos() {

	// Tipos
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

	fmt.Printf("Var: a, Tipo: %T, Valor: %v\n", a, a)
	fmt.Printf("Var: b, Tipo: %T, Valor: %v\n", b, b)
	fmt.Printf("Var: c, Tipo: %T, Valor: %v\n", c, c)
	fmt.Printf("Var: d, Tipo: %T, Valor: %v\n", d, d)
	fmt.Printf("Var: e, Tipo: %T, Valor: %v\n", e, e)
	fmt.Printf("Var: f, Tipo: %T, Valor: %v\n", f, f)
	fmt.Printf("Var: g, Tipo: %T, Valor: %v\n", g, g)
	fmt.Printf("Var: h, Tipo: %T, Valor: %v\n", h, h)
	fmt.Printf("Var: i, Tipo: %T, Valor: %v\n", i, i)
	fmt.Printf("Var: j, Tipo: %T, Valor: %v\n", j, j)
	fmt.Printf("Var: k, Tipo: %T, Valor: %v\n", k, k)
	fmt.Printf("Var: l, Tipo: %T, Valor: %v\n", l, l)
}
