// ============================================================
// AULA 8: Arrays e Slices
// Array = lista de tamanho FIXO
// Slice = lista de tamanho FLEXÍVEL (muito mais comum!)
// ============================================================
package main

import "fmt"

func main() {
	// --------------------------------------------------------
	// ARRAY: tamanho fixo, definido na criação
	// --------------------------------------------------------
	fmt.Println("=== ARRAYS ===")

	var notas [5]float64 // array de 5 float64
	notas[0] = 9.5
	notas[1] = 8.0
	notas[2] = 7.5
	notas[3] = 6.0
	notas[4] = 8.5
	fmt.Println("Notas:", notas)

	// Array com valores já declarados
	cores := [3]string{"Vermelho", "Verde", "Azul"}
	fmt.Println("Cores:", cores)
	fmt.Println("Primeira cor:", cores[0])
	fmt.Println("Total de cores:", len(cores))

	// Deixar o Go contar automaticamente com [...]
	paises := [...]string{"Brasil", "Argentina", "Chile", "Uruguai"}
	fmt.Println("Países:", paises)
	fmt.Println("Qtd países:", len(paises))

	// --------------------------------------------------------
	// SLICE: tamanho flexível (o que você vai usar na prática!)
	// --------------------------------------------------------
	fmt.Println("\n=== SLICES ===")

	// Criando slice vazia e adicionando elementos com append
	var frutas []string
	frutas = append(frutas, "Maçã")
	frutas = append(frutas, "Banana")
	frutas = append(frutas, "Laranja")
	fmt.Println("Frutas:", frutas)

	// Slice já com valores
	numeros := []int{10, 20, 30, 40, 50}
	fmt.Println("Números:", numeros)
	fmt.Println("Comprimento:", len(numeros))
	fmt.Println("Capacidade:", cap(numeros)) // cap = espaço interno reservado

	// Acessando elementos
	fmt.Println("Primeiro:", numeros[0])
	fmt.Println("Último:", numeros[len(numeros)-1])

	// --------------------------------------------------------
	// FATIAMENTO (Slicing): pegar pedaços de uma slice
	// --------------------------------------------------------
	fmt.Println("\n--- Fatiamento ---")
	// Sintaxe: slice[início:fim]  (fim não incluso!)
	fmt.Println("numeros[1:3] =", numeros[1:3]) // pega índice 1 e 2
	fmt.Println("numeros[:3]  =", numeros[:3])  // do início até 2
	fmt.Println("numeros[2:]  =", numeros[2:])  // do 2 até o fim
	fmt.Println("numeros[:]   =", numeros[:])   // tudo

	// --------------------------------------------------------
	// Adicionar múltiplos elementos de uma vez
	// --------------------------------------------------------
	extras := []int{60, 70, 80}
	numeros = append(numeros, extras...) // ... expande a slice
	fmt.Println("\nNúmeros com extras:", numeros)

	// --------------------------------------------------------
	// MAKE: criar slice com tamanho e capacidade definidos
	// --------------------------------------------------------
	listaVazia := make([]int, 5)    // slice com 5 zeros
	fmt.Println("Lista vazia:", listaVazia)
	
	listaCapacidade := make([]int, 3, 10) // 3 elementos, capacidade para 10
	fmt.Println("Lista com capacidade:", listaCapacidade, "| cap:", cap(listaCapacidade))

	// --------------------------------------------------------
	// Copiar uma slice
	// --------------------------------------------------------
	original := []int{1, 2, 3, 4, 5}
	copia := make([]int, len(original))
	copy(copia, original) // copia original para copia

	copia[0] = 999 // modificar copia não afeta original!
	fmt.Println("\nOriginal:", original)
	fmt.Println("Cópia (modificada):", copia)

	// --------------------------------------------------------
	// Percorrer slice com for range
	// --------------------------------------------------------
	fmt.Println("\n--- Percorrendo frutas ---")
	for i, fruta := range frutas {
		fmt.Printf("Posição %d: %s\n", i, fruta)
	}
}
