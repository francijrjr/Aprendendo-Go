// ============================================================
// AULA 7: Funções
// Função = um bloco de código com nome que você pode reutilizar
// "Empacote um conjunto de instruções e chame quando precisar"
// ============================================================
package main

import "fmt"

// --------------------------------------------------------
// Função SIMPLES: não recebe nada, não devolve nada
// --------------------------------------------------------
func dizerOla() {
	fmt.Println("Olá! Eu sou uma função!")
}

// --------------------------------------------------------
// Função com PARÂMETROS: recebe informações
// --------------------------------------------------------
func saudar(nome string) {
	fmt.Println("Olá,", nome, "!")
}

// --------------------------------------------------------
// Função que RETORNA um valor
// --------------------------------------------------------
func somar(a int, b int) int { // o "int" depois dos parênteses = tipo do retorno
	return a + b
}

// Quando os parâmetros têm o mesmo tipo, pode abreviar:
func multiplicar(a, b int) int {
	return a * b
}

// --------------------------------------------------------
// Função que retorna MÚLTIPLOS valores (super útil em Go!)
// --------------------------------------------------------
func dividir(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Erro: não é possível dividir por zero!"
	}
	return a / b, "OK"
}

// --------------------------------------------------------
// Função com retorno NOMEADO (os nomes são declarados no começo)
// --------------------------------------------------------
func calcularArea(largura, altura float64) (area float64, perimetro float64) {
	area = largura * altura
	perimetro = 2 * (largura + altura)
	return // return sem valores = retorna "area" e "perimetro"
}

// --------------------------------------------------------
// Função com número VARIÁVEL de argumentos (variádica)
// --------------------------------------------------------
func somarTudo(numeros ...int) int {
	total := 0
	for _, n := range numeros {
		total += n
	}
	return total
}

// --------------------------------------------------------
// MAIN: onde tudo começa
// --------------------------------------------------------
func main() {
	// Chamando a função simples
	dizerOla()

	// Chamando com parâmetro
	saudar("Maria")
	saudar("João")

	// Usando o retorno da função
	resultado := somar(5, 3)
	fmt.Println("5 + 3 =", resultado)

	fmt.Println("4 x 6 =", multiplicar(4, 6))

	// Recebendo múltiplos retornos
	valor, status := dividir(10, 3)
	fmt.Println("10 / 3 =", valor, "-", status)

	valor2, status2 := dividir(10, 0)
	fmt.Println("10 / 0 =", valor2, "-", status2)

	// Retorno nomeado
	area, perimetro := calcularArea(5, 3)
	fmt.Printf("Área: %.2f | Perímetro: %.2f\n", area, perimetro)

	// Função variádica
	fmt.Println("Soma de 1+2+3+4+5 =", somarTudo(1, 2, 3, 4, 5))
	fmt.Println("Soma de 10+20 =", somarTudo(10, 20))

	// Passando uma slice para função variádica com ...
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println("Soma da slice:", somarTudo(numeros...)) // ... expande a slice

	// --------------------------------------------------------
	// FUNÇÃO ANÔNIMA: função sem nome, usada na hora
	// --------------------------------------------------------
	quadrado := func(n int) int {
		return n * n
	}
	fmt.Println("Quadrado de 7:", quadrado(7))

	// Função anônima chamada imediatamente
	resultado2 := func(a, b int) int {
		return a - b
	}(20, 5)
	fmt.Println("20 - 5 =", resultado2)
}
