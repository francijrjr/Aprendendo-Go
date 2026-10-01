// ============================================================
// AULA INTERMEDIÁRIA 7: Closures (Funções que "lembram")
// Closure = uma função que "captura" variáveis do escopo onde foi criada
// "A função lembra das variáveis ao redor dela, mesmo depois que o
//  escopo original terminou"
// ============================================================
package main

import "fmt"

// --------------------------------------------------------
// Exemplo 1: Contador que lembra o estado
// --------------------------------------------------------
func criarContador() func() int {
	contagem := 0 // esta variável fica "presa" na closure

	return func() int {
		contagem++ // lembra e modifica a variável de fora!
		return contagem
	}
}

// --------------------------------------------------------
// Exemplo 2: Closure com parâmetro inicial
// --------------------------------------------------------
func criarSomador(inicio int) func(int) int {
	return func(valor int) int {
		inicio += valor // inicio é "lembrado" entre chamadas
		return inicio
	}
}

// --------------------------------------------------------
// Exemplo 3: Multiplicador configurável
// --------------------------------------------------------
func criarMultiplicador(fator int) func(int) int {
	return func(n int) int {
		return n * fator // fator é "capturado" pela closure
	}
}

// --------------------------------------------------------
// Exemplo 4: Closure para gerar IDs únicos
// --------------------------------------------------------
func criarGeradorDeID() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

// --------------------------------------------------------
// Exemplo 5: Closure como middleware (padrão funcional)
// --------------------------------------------------------
type Operacao func(int) int

func comLog(nome string, op Operacao) Operacao {
	return func(n int) int {
		resultado := op(n)
		fmt.Printf("  [LOG] %s(%d) = %d\n", nome, n, resultado)
		return resultado
	}
}

// --------------------------------------------------------
// MAIN
// --------------------------------------------------------
func main() {
	// --------------------------------------------------------
	// Contador independente
	// --------------------------------------------------------
	fmt.Println("=== Contador com Closure ===")

	// Criamos DOIS contadores independentes
	contador1 := criarContador()
	contador2 := criarContador()

	fmt.Println("Contador1:", contador1()) // 1
	fmt.Println("Contador1:", contador1()) // 2
	fmt.Println("Contador1:", contador1()) // 3
	fmt.Println("Contador2:", contador2()) // 1 (independente do 1!)
	fmt.Println("Contador2:", contador2()) // 2
	fmt.Println("Contador1:", contador1()) // 4 (continua de onde parou)

	// --------------------------------------------------------
	// Somador configurável
	// --------------------------------------------------------
	fmt.Println("\n=== Somador com Closure ===")

	somadorDe10 := criarSomador(10)
	fmt.Println("Começa em 10, soma 5:", somadorDe10(5))   // 15
	fmt.Println("Agora soma mais 20:", somadorDe10(20))     // 35
	fmt.Println("Agora soma mais 15:", somadorDe10(15))     // 50

	// --------------------------------------------------------
	// Multiplicadores
	// --------------------------------------------------------
	fmt.Println("\n=== Multiplicadores ===")

	dobrar := criarMultiplicador(2)
	triplicar := criarMultiplicador(3)
	centuplicar := criarMultiplicador(100)

	fmt.Println("Dobrar 7:", dobrar(7))          // 14
	fmt.Println("Triplicar 7:", triplicar(7))    // 21
	fmt.Println("Centuplicar 7:", centuplicar(7)) // 700

	// --------------------------------------------------------
	// Gerador de ID
	// --------------------------------------------------------
	fmt.Println("\n=== Gerador de ID ===")

	gerarID := criarGeradorDeID()
	fmt.Println("ID:", gerarID()) // 1
	fmt.Println("ID:", gerarID()) // 2
	fmt.Println("ID:", gerarID()) // 3

	// --------------------------------------------------------
	// Closure com log
	// --------------------------------------------------------
	fmt.Println("\n=== Closure como Middleware ===")

	quadrado := func(n int) int { return n * n }
	quadradoComLog := comLog("quadrado", quadrado)

	quadradoComLog(5)  // [LOG] quadrado(5) = 25
	quadradoComLog(10) // [LOG] quadrado(10) = 100

	// --------------------------------------------------------
	// CUIDADO: closure em loop (erro comum!)
	// --------------------------------------------------------
	fmt.Println("\n=== Closure em Loop ===")

	// ERRADO: todas as closures capturam a mesma variável i
	funcoes := make([]func(), 3)
	for i := 0; i < 3; i++ {
		i := i // cria uma NOVA variável i a cada iteração (correto!)
		funcoes[i] = func() {
			fmt.Println("  Valor:", i)
		}
	}
	
	for _, f := range funcoes {
		f()
	}
}
