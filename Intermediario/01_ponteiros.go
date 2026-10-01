// ============================================================
// AULA INTERMEDIÁRIA 1: Ponteiros
// Ponteiro = o ENDEREÇO de onde uma variável mora na memória
// "Em vez de copiar o valor, aponte para onde ele está"
// ============================================================
package main

import "fmt"

func main() {
	// --------------------------------------------------------
	// O problema SEM ponteiro: a função recebe uma CÓPIA
	// --------------------------------------------------------
	numero := 10
	fmt.Println("Antes de dobrarSemPonteiro:", numero) // 10
	dobrarSemPonteiro(numero)
	fmt.Println("Depois de dobrarSemPonteiro:", numero) // 10 (não mudou!)

	// --------------------------------------------------------
	// A solução COM ponteiro: a função recebe o ENDEREÇO
	// --------------------------------------------------------
	fmt.Println("\nAntes de dobrarComPonteiro:", numero) // 10
	dobrarComPonteiro(&numero) // & = "pegue o endereço de numero"
	fmt.Println("Depois de dobrarComPonteiro:", numero)  // 20 (mudou!)

	// --------------------------------------------------------
	// Entendendo ponteiros na prática
	// --------------------------------------------------------
	fmt.Println("\n--- Entendendo ponteiros ---")
	
	x := 42
	var p *int = &x // p é um ponteiro para int. & pega o endereço de x.
	
	fmt.Println("Valor de x:", x)       // 42
	fmt.Println("Endereço de x:", &x)   // algo como: 0xc0000b4010
	fmt.Println("Valor de p (endereço):", p) // mesmo endereço acima
	fmt.Println("Valor apontado por p:", *p) // 42  (* = "vá até o endereço e pegue o valor")

	// Modificar o valor ATRAVÉS do ponteiro
	*p = 100
	fmt.Println("x após *p = 100:", x) // 100

	// --------------------------------------------------------
	// new(): cria um ponteiro para um tipo, já inicializado
	// --------------------------------------------------------
	fmt.Println("\n--- Usando new() ---")
	ponteiroParaInt := new(int) // cria um int com valor 0, retorna o ponteiro
	fmt.Println("Valor inicial:", *ponteiroParaInt) // 0
	*ponteiroParaInt = 55
	fmt.Println("Após atribuição:", *ponteiroParaInt) // 55

	// --------------------------------------------------------
	// Ponteiros com Structs
	// --------------------------------------------------------
	fmt.Println("\n--- Ponteiro com Struct ---")
	type Cachorro struct {
		Nome  string
		Raca  string
		Idade int
	}

	dog := Cachorro{Nome: "Rex", Raca: "Labrador", Idade: 3}
	fmt.Println("Antes:", dog.Nome)

	renomear(&dog, "Buddy") // passa o endereço da struct
	fmt.Println("Depois:", dog.Nome) // Rex virou Buddy!
	
	// Atalho: Go já sabe que você quer o campo do ponteiro
	ponteiroDog := &dog
	ponteiroDog.Idade = 4        // Go entende isso como (*ponteiroDog).Idade = 4
	fmt.Println("Idade:", dog.Idade) // 4
}

// Recebe cópia - NÃO modifica o original
func dobrarSemPonteiro(n int) {
	n = n * 2
	fmt.Println("  (dentro da função):", n)
}

// Recebe ponteiro - MODIFICA o original
func dobrarComPonteiro(n *int) { // *int = ponteiro para int
	*n = *n * 2 // * = "vá até o endereço e modifique o valor lá"
}

type Cachorro struct {
	Nome  string
	Raca  string
	Idade int
}

func renomear(c *Cachorro, novoNome string) {
	c.Nome = novoNome // Go entende como (*c).Nome = novoNome
}
