// ============================================================
// AULA AVANÇADA 1: Generics (Tipos Genéricos)
// Genéricos = escreva código que funciona para QUALQUER tipo
// "Em vez de escrever a mesma função 10 vezes para int, float,
//  string... escreva UMA vez que funciona pra tudo!"
// Disponível desde Go 1.18
// ============================================================
package main

import (
	"fmt"
)

// --------------------------------------------------------
// SEM genérico: precisaríamos de uma função para cada tipo
// --------------------------------------------------------
func somarInts(a, b int) int       { return a + b }
func somarFloat64(a, b float64) float64 { return a + b }
// ... e repetiria para string, int64, etc. Que chato!

// --------------------------------------------------------
// COM genérico: UMA função para todos os tipos numéricos
// --------------------------------------------------------
// [T int | float64 | int64] = T pode ser qualquer um desses tipos
func somar[T int | float64 | int64](a, b T) T {
	return a + b
}

// --------------------------------------------------------
// Usando "constraints" para ser mais flexível
// --------------------------------------------------------
// Number agrupa todos os tipos numéricos
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func maximo[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func minimo[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// --------------------------------------------------------
// Slice genérica: funções que operam em listas de qualquer tipo
// --------------------------------------------------------
func filtrar[T any](lista []T, teste func(T) bool) []T {
	resultado := []T{}
	for _, item := range lista {
		if teste(item) {
			resultado = append(resultado, item)
		}
	}
	return resultado
}

func mapear[T any, R any](lista []T, transformar func(T) R) []R {
	resultado := make([]R, len(lista))
	for i, item := range lista {
		resultado[i] = transformar(item)
	}
	return resultado
}

func reduzir[T any, R any](lista []T, inicial R, acumular func(R, T) R) R {
	resultado := inicial
	for _, item := range lista {
		resultado = acumular(resultado, item)
	}
	return resultado
}

// --------------------------------------------------------
// Struct genérica: Par de valores de qualquer tipo
// --------------------------------------------------------
type Par[A, B any] struct {
	Primeiro A
	Segundo  B
}

func (p Par[A, B]) String() string {
	return fmt.Sprintf("(%v, %v)", p.Primeiro, p.Segundo)
}

// --------------------------------------------------------
// Pilha genérica (Stack)
// --------------------------------------------------------
type Pilha[T any] struct {
	itens []T
}

func (p *Pilha[T]) Empilhar(item T) {
	p.itens = append(p.itens, item)
}

func (p *Pilha[T]) Desempilhar() (T, bool) {
	var zero T
	if len(p.itens) == 0 {
		return zero, false
	}
	topo := p.itens[len(p.itens)-1]
	p.itens = p.itens[:len(p.itens)-1]
	return topo, true
}

func (p *Pilha[T]) Tamanho() int {
	return len(p.itens)
}

// --------------------------------------------------------
// MAIN
// --------------------------------------------------------
func main() {
	// --------------------------------------------------------
	// Funções genéricas
	// --------------------------------------------------------
	fmt.Println("=== Funções Genéricas ===")
	fmt.Println("Soma ints:", somar(5, 3))
	fmt.Println("Soma float64:", somar(3.14, 2.86))
	fmt.Println("Soma int64:", somar[int64](100, 200))

	fmt.Println("\nMáximo:", maximo(10, 20))
	fmt.Println("Máximo float:", maximo(3.5, 7.2))
	fmt.Println("Mínimo:", minimo(10, 20))

	// --------------------------------------------------------
	// Funções de slice genéricas (estilo funcional)
	// --------------------------------------------------------
	fmt.Println("\n=== Funções de Slice Genéricas ===")

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	pares := filtrar(numeros, func(n int) bool { return n%2 == 0 })
	fmt.Println("Pares:", pares)

	maioresQue5 := filtrar(numeros, func(n int) bool { return n > 5 })
	fmt.Println("Maiores que 5:", maioresQue5)

	quadrados := mapear(numeros, func(n int) int { return n * n })
	fmt.Println("Quadrados:", quadrados)

	textos := mapear(numeros, func(n int) string {
		return fmt.Sprintf("num%d", n)
	})
	fmt.Println("Como texto:", textos)

	soma := reduzir(numeros, 0, func(acc, n int) int { return acc + n })
	fmt.Println("Soma total:", soma)

	// --------------------------------------------------------
	// Struct genérica
	// --------------------------------------------------------
	fmt.Println("\n=== Struct Genérica (Par) ===")

	par1 := Par[string, int]{Primeiro: "idade", Segundo: 30}
	par2 := Par[float64, bool]{Primeiro: 3.14, Segundo: true}
	par3 := Par[string, string]{Primeiro: "Brasil", Segundo: "Brasília"}

	fmt.Println("Par1:", par1)
	fmt.Println("Par2:", par2)
	fmt.Println("Par3:", par3)

	// --------------------------------------------------------
	// Pilha genérica
	// --------------------------------------------------------
	fmt.Println("\n=== Pilha Genérica ===")

	// Pilha de inteiros
	pilhaInt := &Pilha[int]{}
	pilhaInt.Empilhar(10)
	pilhaInt.Empilhar(20)
	pilhaInt.Empilhar(30)

	fmt.Println("Tamanho da pilha:", pilhaInt.Tamanho())
	for pilhaInt.Tamanho() > 0 {
		valor, _ := pilhaInt.Desempilhar()
		fmt.Println("Desempilhou:", valor)
	}

	// Pilha de strings
	pilhaStr := &Pilha[string]{}
	pilhaStr.Empilhar("primeiro")
	pilhaStr.Empilhar("segundo")
	pilhaStr.Empilhar("terceiro")

	v, ok := pilhaStr.Desempilhar()
	fmt.Println("\nDesempilhou string:", v, "| OK:", ok)
}
