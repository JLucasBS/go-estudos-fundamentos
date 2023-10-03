package main

import "fmt"

func main() {
	i := 1

	
	var p *int = nil // * é utilizado para identificar ponteiro, nesse caso, estou atribuindo o tipo int para o ponteiro
	
	p = &i // pegando o endereço de memoria da variável i e atribuindo para p
	
	*p++ //desreferenciando o ponteiro (pegando o valor)
	i++ 

	// Go não tem como aritmética de ponteiros 
	// Exp: p++

	fmt.Println(p, *p, i, &i)
}
