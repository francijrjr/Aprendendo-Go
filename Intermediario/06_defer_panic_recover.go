// ============================================================
// AULA INTERMEDIÁRIA 6: Defer, Panic e Recover
// defer  = "execute isso quando a função terminar"
// panic  = "algo muito errado aconteceu, para tudo!"
// recover = "captura o panic antes de explodir o programa"
// ============================================================
package main

import "fmt"

func main() {
	// --------------------------------------------------------
	// DEFER: executa no final da função, sempre!
	// --------------------------------------------------------
	fmt.Println("=== Defer ===")
	demonstrarDefer()

	// --------------------------------------------------------
	// DEFER com múltiplos: executam em ordem inversa (LIFO)
	// --------------------------------------------------------
	fmt.Println("\n=== Múltiplos Defer ===")
	fmt.Println("Início")
	defer fmt.Println("Defer 1 - primeiro chamado, último executado")
	defer fmt.Println("Defer 2 - segundo chamado, segundo a sair")
	defer fmt.Println("Defer 3 - terceiro chamado, primeiro a sair")
	fmt.Println("Fim do main (os defers rodam agora)")

	// --------------------------------------------------------
	// PANIC e RECOVER
	// --------------------------------------------------------
	fmt.Println("\n=== Panic e Recover ===")

	// Chamamos uma função que vai dar panic, mas tratamos o erro
	err := funcaoArriscada()
	if err != nil {
		fmt.Println("Erro capturado no main:", err)
	}

	fmt.Println("O programa continua normalmente!")

	// --------------------------------------------------------
	// Caso real: abrir arquivo (simulado)
	// --------------------------------------------------------
	fmt.Println("\n=== Defer para limpeza de recursos ===")
	simularAberturaDeBancoDeDados()
}

func demonstrarDefer() {
	fmt.Println("  1. Início da função")
	defer fmt.Println("  3. Defer executado ao final")
	fmt.Println("  2. Meio da função")
	// Quando a função termina, o defer roda
}

// Função que pode dar panic, mas recover salva
func funcaoArriscada() (err error) {
	// defer com recover: "se der panic, pegue e retorne como erro"
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("  Recover pegou o panic:", r)
			err = fmt.Errorf("operação falhou: %v", r)
		}
	}()

	fmt.Println("  Fazendo algo arriscado...")
	// Simula uma divisão por zero que causaria panic
	a := 10
	b := 0
	if b == 0 {
		panic("divisão por zero detectada!") // lança um panic
	}
	fmt.Println("  Resultado:", a/b) // nunca chega aqui
	return nil
}

// Simula trabalhar com um recurso que precisa ser fechado
func simularAberturaDeBancoDeDados() {
	fmt.Println("  Abrindo conexão com o banco...")
	conexao := abrirConexao()
	defer conexao.fechar() // garante que vai fechar, mesmo se der erro!

	fmt.Println("  Fazendo queries no banco...")
	conexao.query("SELECT * FROM usuarios")
	conexao.query("UPDATE produtos SET estoque = 0")

	fmt.Println("  Trabalho concluído.")
	// Quando a função terminar, o defer vai chamar conexao.fechar()
}

type Conexao struct {
	nome string
}

func abrirConexao() *Conexao {
	c := &Conexao{nome: "MinhaConexao"}
	fmt.Println("  [Conexão aberta:", c.nome, "]")
	return c
}

func (c *Conexao) fechar() {
	fmt.Println("  [Conexão fechada:", c.nome, "] — chamado pelo defer!")
}

func (c *Conexao) query(sql string) {
	fmt.Println("  [Query executada]:", sql)
}
