package main

import "fmt"

func main() {

	// Valores default para quando nao é inicializado com algum valor
	var a int
	var b float64
	var c bool
	var d string
	// ponteiro null é o nil
	var e *int

	// %q = vai mostrar minha string vazia
	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}