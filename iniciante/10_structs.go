// ============================================================
// AULA 10: Structs
// Struct = uma "ficha" que agrupa informações relacionadas
// "Em vez de ter nome, idade e email separados,
//  coloque tudo numa ficha chamada Pessoa"
// ============================================================
package main

import "fmt"

// --------------------------------------------------------
// Definindo uma Struct (fora do main, é global)
// --------------------------------------------------------
type Pessoa struct {
	Nome   string
	Idade  int
	Email  string
	Ativo  bool
}

type Endereco struct {
	Rua    string
	Cidade string
	CEP    string
}

// Struct que usa outra Struct (composição)
type Cliente struct {
	Pessoa   // "herda" todos os campos de Pessoa
	Endereco // "herda" todos os campos de Endereco
	Credito  float64
}

type Produto struct {
	Nome     string
	Preco    float64
	Estoque  int
}

// --------------------------------------------------------
// Função que recebe uma Struct como parâmetro
// --------------------------------------------------------
func apresentar(p Pessoa) {
	fmt.Printf("Nome: %s | Idade: %d | Email: %s\n", p.Nome, p.Idade, p.Email)
}

// --------------------------------------------------------
// Método: função que "pertence" a uma Struct
// --------------------------------------------------------
func (p Produto) descricao() string {
	return fmt.Sprintf("%s - R$ %.2f (Estoque: %d)", p.Nome, p.Preco, p.Estoque)
}

func (p *Produto) aplicarDesconto(percentual float64) {
	p.Preco = p.Preco * (1 - percentual/100)
}

// --------------------------------------------------------
// MAIN
// --------------------------------------------------------
func main() {
	// Criando uma Struct - Jeito 1: campo a campo
	var pessoa1 Pessoa
	pessoa1.Nome = "João Silva"
	pessoa1.Idade = 30
	pessoa1.Email = "joao@email.com"
	pessoa1.Ativo = true
	
	fmt.Println("=== Pessoa 1 ===")
	fmt.Println(pessoa1)
	fmt.Println("Nome:", pessoa1.Nome)
	fmt.Println("Idade:", pessoa1.Idade)

	// Criando uma Struct - Jeito 2: tudo de uma vez (preferido!)
	pessoa2 := Pessoa{
		Nome:  "Maria Santos",
		Idade: 25,
		Email: "maria@email.com",
		Ativo: true,
	}
	
	fmt.Println("\n=== Pessoa 2 ===")
	apresentar(pessoa2)

	// --------------------------------------------------------
	// Modificar campos
	// --------------------------------------------------------
	pessoa2.Idade = 26
	fmt.Println("Idade atualizada:", pessoa2.Idade)

	// --------------------------------------------------------
	// Struct com composição
	// --------------------------------------------------------
	fmt.Println("\n=== Cliente ===")
	cliente := Cliente{
		Pessoa: Pessoa{
			Nome:  "Carlos Oliveira",
			Idade: 35,
			Email: "carlos@email.com",
		},
		Endereco: Endereco{
			Rua:    "Rua das Flores, 123",
			Cidade: "São Paulo",
			CEP:    "01000-000",
		},
		Credito: 5000.00,
	}

	// Acesso direto aos campos aninhados
	fmt.Println("Nome:", cliente.Nome)      // vai até Pessoa.Nome
	fmt.Println("Cidade:", cliente.Cidade)  // vai até Endereco.Cidade
	fmt.Println("Crédito:", cliente.Credito)

	// --------------------------------------------------------
	// Usando métodos da Struct
	// --------------------------------------------------------
	fmt.Println("\n=== Produto ===")
	produto := Produto{
		Nome:    "Camiseta Go",
		Preco:   79.90,
		Estoque: 50,
	}
	
	fmt.Println(produto.descricao())
	produto.aplicarDesconto(10) // desconto de 10%
	fmt.Println("Após 10% de desconto:")
	fmt.Println(produto.descricao())

	// --------------------------------------------------------
	// Slice de Structs (lista de fichas)
	// --------------------------------------------------------
	fmt.Println("\n=== Lista de Pessoas ===")
	pessoas := []Pessoa{
		{Nome: "Ana", Idade: 22, Email: "ana@email.com"},
		{Nome: "Bob", Idade: 28, Email: "bob@email.com"},
		{Nome: "Cia", Idade: 31, Email: "cia@email.com"},
	}

	for _, p := range pessoas {
		apresentar(p)
	}
}
