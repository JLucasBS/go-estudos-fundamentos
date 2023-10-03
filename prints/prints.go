package main

import "fmt"

func main() {

	// Função print utiliza a mesma linha, mesmo se for duas sentenças diferentes
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// Função println quebra a linha (começa na linha do print anterior)
	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.141516

	// fmt.Println("O valor de x é " + x)  isso da erro devido a tentar concatenar uma string e um float64

	// retorna o que eu passei como uma string
	xs := fmt.Sprint(x)
	// com isso, consigo concatenar duas strings
	fmt.Println("O valor de x é " + xs)

	// utilizando a virgula, da pra concatenar uma string com outro valor, sem precisar converter
	fmt.Println("O valor de x é", x, "!!!")

	/* 
		print f eu consigo formatar a string
		é utilizado o f pois significa que vai utilizar uma variável do tipo float
		utilizando .2f, eu trago somente duas casas decimais na string(fazendo isso ele arredonda)
	*/
	fmt.Printf("O valor de x é %.2f.", x)

	a, b, c, d := 1, 1.9999, false, "opa!"
	/*
		%d = formata valores do tipo inteiro
		%f = formata valores do tipo float
		$.1f = formata valores o tipo float e retorna 1 cas decimal
		%t = formata valores do tipo boolean
		%s = formata valores do tipo string 
	*/
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

	// utilizando %v ele formata quase todos os tipos
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}