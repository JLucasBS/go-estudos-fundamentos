package main

import (
	"fmt"
	// Referenciando o pacote math com a letra m
	m "math"
)

func main () {
	// constant variável - tipo = valor
	const PI float64 = 3.1415

	// Se atribuído valor pra variável, compilador consegue atribuir um tipo por padrão
	var raio = 3.2 //tipo (float64)

	/*
	 forma reduzida de criar uma var (melhor forma)

	 := inicializa uma variável e ja atribui um valor a ela

	 variável nao utilizada aponta erro!
	*/
	area :=  PI * m.Pow(raio, 2)

	// Concatenando utilizando virgula, se utiliza o +, por ser tipo diferente, da erro
	fmt.Println("A Area da circunferência é", area)


	// Iniciando mais de uma variável em mais de uma linha
	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	// iniciando mais de uma variável em uma única linha
	var e, f bool = true, false
	fmt.Println(e, f)

	// Utilizando inicialização simplificada para criar mais de uma variável
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}