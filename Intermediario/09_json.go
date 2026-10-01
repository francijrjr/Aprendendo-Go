// ============================================================
// AULA INTERMEDIÁRIA 9: JSON
// JSON = formato popular para trocar dados entre sistemas
// Go tem suporte nativo para converter structs <-> JSON
// ============================================================
package main

import (
	"encoding/json"
	"fmt"
)

// --------------------------------------------------------
// Structs para trabalhar com JSON
// As tags `json:"nome"` definem os nomes no JSON
// --------------------------------------------------------
type Pessoa struct {
	Nome      string   `json:"nome"`
	Idade     int      `json:"idade"`
	Email     string   `json:"email"`
	Ativo     bool     `json:"ativo"`
	Hobbies   []string `json:"hobbies"`
}

type Produto struct {
	ID       int     `json:"id"`
	Nome     string  `json:"nome"`
	Preco    float64 `json:"preco"`
	Estoque  int     `json:"estoque,omitempty"` // omitempty = ignora se for zero
	Interno  string  `json:"-"`                  // - = nunca aparece no JSON
}

type Resposta struct {
	Status  string      `json:"status"`
	Dados   interface{} `json:"dados"`
	Total   int         `json:"total,omitempty"`
}

func main() {
	// --------------------------------------------------------
	// MARSHAL: Struct → JSON (objeto Go vira texto JSON)
	// --------------------------------------------------------
	fmt.Println("=== Struct para JSON ===")

	pessoa := Pessoa{
		Nome:    "João Silva",
		Idade:   30,
		Email:   "joao@email.com",
		Ativo:   true,
		Hobbies: []string{"Programação", "Leitura", "Ciclismo"},
	}

	// json.Marshal converte para JSON
	jsonBytes, err := json.Marshal(pessoa)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("JSON compacto:")
	fmt.Println(string(jsonBytes))

	// MarshalIndent = JSON formatado (legível)
	jsonFormatado, _ := json.MarshalIndent(pessoa, "", "  ")
	fmt.Println("\nJSON formatado:")
	fmt.Println(string(jsonFormatado))

	// --------------------------------------------------------
	// Testando omitempty e -
	// --------------------------------------------------------
	fmt.Println("\n=== Tags especiais ===")
	
	produto1 := Produto{ID: 1, Nome: "Teclado", Preco: 299.90, Estoque: 5, Interno: "segredo"}
	produto2 := Produto{ID: 2, Nome: "Mouse", Preco: 89.90} // Estoque = 0 (omitido)

	j1, _ := json.MarshalIndent(produto1, "", "  ")
	j2, _ := json.MarshalIndent(produto2, "", "  ")

	fmt.Println("Produto com estoque:")
	fmt.Println(string(j1))
	fmt.Println("\nProduto sem estoque (omitempty):")
	fmt.Println(string(j2)) // "estoque" não aparece, "interno" nunca aparece

	// --------------------------------------------------------
	// UNMARSHAL: JSON → Struct (texto JSON vira objeto Go)
	// --------------------------------------------------------
	fmt.Println("\n=== JSON para Struct ===")

	textoJSON := `{
		"nome": "Maria Santos",
		"idade": 25,
		"email": "maria@email.com",
		"ativo": true,
		"hobbies": ["Dança", "Culinária", "Viagem"]
	}`

	var pessoaLida Pessoa
	err = json.Unmarshal([]byte(textoJSON), &pessoaLida)
	if err != nil {
		fmt.Println("Erro ao ler JSON:", err)
		return
	}

	fmt.Printf("Nome: %s\n", pessoaLida.Nome)
	fmt.Printf("Idade: %d\n", pessoaLida.Idade)
	fmt.Printf("Hobbies: %v\n", pessoaLida.Hobbies)

	// --------------------------------------------------------
	// JSON com lista (slice)
	// --------------------------------------------------------
	fmt.Println("\n=== JSON com Lista ===")

	jsonLista := `[
		{"id": 1, "nome": "Notebook", "preco": 3499.90},
		{"id": 2, "nome": "Celular", "preco": 1299.90},
		{"id": 3, "nome": "Tablet", "preco": 899.90}
	]`

	var produtos []Produto
	json.Unmarshal([]byte(jsonLista), &produtos)

	for _, p := range produtos {
		fmt.Printf("ID: %d | %s - R$ %.2f\n", p.ID, p.Nome, p.Preco)
	}

	// --------------------------------------------------------
	// JSON genérico com map[string]interface{}
	// Quando você não sabe a estrutura do JSON antecipadamente
	// --------------------------------------------------------
	fmt.Println("\n=== JSON Genérico ===")

	jsonGenerico := `{"status": "ok", "versao": 2, "dados": {"total": 100}}`

	var dados map[string]interface{}
	json.Unmarshal([]byte(jsonGenerico), &dados)

	for chave, valor := range dados {
		fmt.Printf("  %s: %v (tipo: %T)\n", chave, valor, valor)
	}

	// --------------------------------------------------------
	// Construindo Resposta de API
	// --------------------------------------------------------
	fmt.Println("\n=== Resposta de API ===")

	resposta := Resposta{
		Status: "success",
		Dados:  produtos,
		Total:  len(produtos),
	}

	jsonResposta, _ := json.MarshalIndent(resposta, "", "  ")
	fmt.Println(string(jsonResposta))
}
