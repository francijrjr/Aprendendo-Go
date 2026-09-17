// ============================================================
// AULA 4: Operadores
// Operadores são os símbolos que fazem cálculos e comparações
// ============================================================
package main

import "fmt"

func main() {
	// --------------------------------------------------------
	// OPERADORES MATEMÁTICOS
	// --------------------------------------------------------
	a := 10
	b := 3

	fmt.Println("--- Matemática ---")
	fmt.Println("Soma:       ", a + b)  // 13
	fmt.Println("Subtração:  ", a - b)  // 7
	fmt.Println("Multiplicação:", a * b) // 30
	fmt.Println("Divisão:    ", a / b)  // 3 (divisão inteira, sem decimal!)
	fmt.Println("Resto:      ", a % b)  // 1 (sobra da divisão)

	// Divisão com decimal: use float64
	x := 10.0
	y := 3.0
	fmt.Println("Divisão decimal:", x / y) // 3.333...

	// --------------------------------------------------------
	// ATALHOS MATEMÁTICOS
	// --------------------------------------------------------
	fmt.Println("\n--- Atalhos ---")
	numero := 10
	
	numero += 5  // mesmo que: numero = numero + 5
	fmt.Println("numero += 5:", numero) // 15
	
	numero -= 3  // mesmo que: numero = numero - 3
	fmt.Println("numero -= 3:", numero) // 12
	
	numero *= 2  // mesmo que: numero = numero * 2
	fmt.Println("numero *= 2:", numero) // 24
	
	numero /= 4  // mesmo que: numero = numero / 4
	fmt.Println("numero /= 4:", numero) // 6
	
	numero++     // soma 1 (mesmo que numero = numero + 1)
	fmt.Println("numero++:", numero)    // 7
	
	numero--     // subtrai 1 (mesmo que numero = numero - 1)
	fmt.Println("numero--:", numero)    // 6

	// --------------------------------------------------------
	// OPERADORES DE COMPARAÇÃO (o resultado é true ou false)
	// --------------------------------------------------------
	fmt.Println("\n--- Comparações ---")
	c := 10
	d := 20
	
	fmt.Println("c == d (igual?):", c == d)         // false
	fmt.Println("c != d (diferente?):", c != d)     // true
	fmt.Println("c > d (maior?):", c > d)            // false
	fmt.Println("c < d (menor?):", c < d)            // true
	fmt.Println("c >= d (maior ou igual?):", c >= d) // false
	fmt.Println("c <= d (menor ou igual?):", c <= d) // true

	// --------------------------------------------------------
	// OPERADORES LÓGICOS (combinar condições)
	// --------------------------------------------------------
	fmt.Println("\n--- Lógicos ---")
	chovendo := true
	fimDeSemana := false

	// && = E (AND): os DOIS precisam ser verdadeiros
	fmt.Println("Chovendo E fim de semana:", chovendo && fimDeSemana) // false

	// || = OU (OR): PELO MENOS UM precisa ser verdadeiro
	fmt.Println("Chovendo OU fim de semana:", chovendo || fimDeSemana) // true

	// ! = NÃO (NOT): inverte o valor
	fmt.Println("NÃO está chovendo:", !chovendo) // false
}
