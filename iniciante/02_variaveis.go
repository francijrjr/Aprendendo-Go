// ============================================================
// AULA 2: Variáveis
// Variável = uma caixinha que guarda um valor
// ============================================================
package main

import "fmt"

func main() {
	// --------------------------------------------------------
	// JEITO 1: Declarar e depois atribuir
	// --------------------------------------------------------
	var nome string       // declarei uma caixinha chamada "nome" que guarda texto
	nome = "Maria"        // coloquei o valor "Maria" dentro dela
	fmt.Println(nome)

	// --------------------------------------------------------
	// JEITO 2: Declarar e já colocar o valor
	// --------------------------------------------------------
	var idade int = 25    // caixinha "idade" que guarda número inteiro
	fmt.Println(idade)

	// --------------------------------------------------------
	// JEITO 3: O mais comum! Go descobre o tipo sozinho
	// O := significa "crie a variável e coloque este valor"
	// --------------------------------------------------------
	cidade := "São Paulo"  // Go vê que é texto e cria como string
	peso := 75.5           // Go vê o ponto decimal e cria como float64
	casado := true         // Go vê true/false e cria como bool
	
	fmt.Println(cidade)
	fmt.Println(peso)
	fmt.Println(casado)

	// --------------------------------------------------------
	// Declarar várias variáveis de uma vez
	// --------------------------------------------------------
	var (
		pais    string = "Brasil"
		estado  string = "SP"
		capital string = "Brasília"
	)
	fmt.Println(pais, estado, capital)

	// --------------------------------------------------------
	// Variável sem valor = valor padrão (zero value)
	// --------------------------------------------------------
	var numero int     // padrão para int = 0
	var texto string   // padrão para string = "" (texto vazio)
	var ativo bool     // padrão para bool = false
	
	fmt.Println("Número padrão:", numero)
	fmt.Println("Texto padrão:", texto)
	fmt.Println("Bool padrão:", ativo)

	// --------------------------------------------------------
	// Trocar o valor de uma variável
	// --------------------------------------------------------
	cor := "azul"
	fmt.Println("Cor antes:", cor)
	cor = "vermelho"   // Note: só um = para trocar o valor
	fmt.Println("Cor depois:", cor)

	// --------------------------------------------------------
	// Constante = valor que NUNCA muda
	// --------------------------------------------------------
	const PI = 3.14159
	const PAIS_NATAL = "Brasil"
	fmt.Println("PI =", PI)
	fmt.Println("País natal:", PAIS_NATAL)
}
