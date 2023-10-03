package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 //ou exclusivo

	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv30, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t\nTv32: %t\nSorvete: %t\nSaudável: %t", tv50, tv30, sorvete, !sorvete)
}