package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	//convertendo int para float64 em cálculos (poderia fazer ao contrario também)
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado!!
	fmt.Println("Teste " + string(97)) //Nesse caso, ele não transforma o 97 em string, ele pega o valor 97 da tabela Unicode (a)

	// convertendo int para string corretamente
	fmt.Println("Teste " + strconv.Itoa(123))

	// convertendo string para int
	num, _ := strconv.Atoi("123")
	fmt.Println("A string convertida é", num)

	// convertendo string para boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}

