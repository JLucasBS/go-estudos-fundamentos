package main

// Não existe operador ternário no Go!
func obterRestultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	} 
	return "Reprovado"
}

func main() {
	obterRestultado(6.2)
}