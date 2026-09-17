// ============================================================
// AULA 6: Loops (Repetições)
// "Repita isso X vezes" ou "repita enquanto a condição for verdadeira"
// Em Go, só existe o "for" (mas ele faz tudo!)
// ============================================================
package main

import "fmt"

func main() {
	// --------------------------------------------------------
	// FOR clássico: conta de 1 até 5
	// Estrutura: for início; condição; incremento { }
	// --------------------------------------------------------
	fmt.Println("--- Contando de 1 até 5 ---")
	for i := 1; i <= 5; i++ {
		fmt.Println("Número:", i)
	}

	// --------------------------------------------------------
	// FOR como WHILE: repete enquanto condição for verdadeira
	// --------------------------------------------------------
	fmt.Println("\n--- Contagem regressiva ---")
	contador := 5
	for contador > 0 {
		fmt.Println(contador)
		contador--
	}
	fmt.Println("Fogo!")

	// --------------------------------------------------------
	// FOR infinito com BREAK: repete para sempre, mas tem saída
	// --------------------------------------------------------
	fmt.Println("\n--- Achando o número 7 ---")
	numero := 0
	for {
		numero++
		if numero == 7 {
			fmt.Println("Achei o 7!")
			break // BREAK = para o loop agora!
		}
	}

	// --------------------------------------------------------
	// CONTINUE: pula para a próxima repetição
	// --------------------------------------------------------
	fmt.Println("\n--- Números pares de 1 a 10 ---")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 { // se for ímpar...
			continue   // ...pula para o próximo i
		}
		fmt.Println(i) // só chega aqui se for par
	}

	// --------------------------------------------------------
	// FOR com RANGE: percorre listas, textos, maps, etc.
	// --------------------------------------------------------
	fmt.Println("\n--- Lista de frutas ---")
	frutas := []string{"Maçã", "Banana", "Laranja", "Uva"}
	
	for posicao, fruta := range frutas {
		fmt.Println(posicao, "-", fruta)
	}

	// Se não quiser a posição, use _ (underscore)
	fmt.Println("\n--- Só as frutas, sem posição ---")
	for _, fruta := range frutas {
		fmt.Println(fruta)
	}

	// --------------------------------------------------------
	// FOR dentro de FOR (loop aninhado)
	// Tabuada do 2 e do 3
	// --------------------------------------------------------
	fmt.Println("\n--- Tabuada ---")
	for tabela := 2; tabela <= 3; tabela++ {
		fmt.Printf("Tabuada do %d:\n", tabela)
		for i := 1; i <= 5; i++ {
			fmt.Printf("  %d x %d = %d\n", tabela, i, tabela*i)
		}
	}
}
