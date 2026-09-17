// ============================================================
// AULA 9: Maps (Mapas / Dicionários)
// Map = lista de CHAVE + VALOR
// "Dado o nome (chave), me dê o valor"
// É como um dicionário: palavra → definição
// ============================================================
package main

import "fmt"

func main() {
	// --------------------------------------------------------
	// Criando um Map
	// Sintaxe: map[tipo_da_chave]tipo_do_valor
	// --------------------------------------------------------
	
	// Jeito 1: declarar e depois adicionar
	capitais := map[string]string{}
	capitais["Brasil"] = "Brasília"
	capitais["Argentina"] = "Buenos Aires"
	capitais["França"] = "Paris"
	capitais["Japão"] = "Tóquio"
	
	fmt.Println("Capitais:", capitais)
	fmt.Println("Capital do Brasil:", capitais["Brasil"])

	// Jeito 2: já com valores
	idades := map[string]int{
		"Ana":   25,
		"Carlos": 30,
		"Bia":   22,
	}
	fmt.Println("\nIdades:", idades)

	// Jeito 3: com make
	precos := make(map[string]float64)
	precos["Maçã"] = 3.50
	precos["Banana"] = 2.00
	precos["Uva"] = 8.90

	// --------------------------------------------------------
	// Acessar valores
	// --------------------------------------------------------
	fmt.Println("\n--- Acessando valores ---")
	fmt.Println("Preço da Maçã:", precos["Maçã"])

	// Verificar se uma chave existe!
	// Retorna o valor E um bool dizendo se achou
	preco, existe := precos["Manga"]
	if existe {
		fmt.Println("Preço da Manga:", preco)
	} else {
		fmt.Println("Manga não está no mapa. Preço retornado:", preco) // retorna 0
	}

	// --------------------------------------------------------
	// Adicionar e Atualizar
	// --------------------------------------------------------
	fmt.Println("\n--- Adicionar / Atualizar ---")
	idades["Diego"] = 28          // adicionar novo
	idades["Ana"] = 26            // atualizar existente
	fmt.Println("Idades atualizadas:", idades)

	// --------------------------------------------------------
	// Remover uma chave
	// --------------------------------------------------------
	delete(idades, "Bia") // remove "Bia" do map
	fmt.Println("Após remover Bia:", idades)

	// --------------------------------------------------------
	// Tamanho do Map
	// --------------------------------------------------------
	fmt.Println("Quantidade de pessoas:", len(idades))

	// --------------------------------------------------------
	// Percorrer um Map com for range
	// --------------------------------------------------------
	fmt.Println("\n--- Percorrendo capitais ---")
	for pais, capital := range capitais {
		fmt.Printf("Capital de %s = %s\n", pais, capital)
	}

	// --------------------------------------------------------
	// Map de Map (estrutura mais complexa)
	// --------------------------------------------------------
	fmt.Println("\n--- Ficha de Usuários ---")
	usuarios := map[string]map[string]string{
		"joao": {
			"email": "joao@email.com",
			"cidade": "São Paulo",
		},
		"maria": {
			"email": "maria@email.com",
			"cidade": "Rio de Janeiro",
		},
	}

	for usuario, dados := range usuarios {
		fmt.Printf("Usuário: %s | Email: %s | Cidade: %s\n",
			usuario, dados["email"], dados["cidade"])
	}
}
