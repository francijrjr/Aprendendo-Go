// ============================================================
// AULA AVANÇADA 3: Testes em Go
// Go tem suporte nativo para testes — sem framework externo!
// Este arquivo demonstra como ESCREVER testes.
//
// Como rodar os testes:
//   go test ./...            (roda todos)
//   go test -v ./...         (verbose: mostra cada teste)
//   go test -run TestSomar   (roda só o TestSomar)
//   go test -bench=.         (roda benchmarks)
//   go test -cover           (mostra cobertura de código)
// ============================================================
package main

import (
	"fmt"
	"strings"
	"testing"
)

// ============================================================
// FUNÇÕES A SEREM TESTADAS (normalmente ficam em outro arquivo)
// ============================================================

func Somar(a, b int) int {
	return a + b
}

func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("não pode dividir por zero")
	}
	return a / b, nil
}

func PalavrasEmMaiusculo(s string) string {
	return strings.ToUpper(s)
}

func EhPar(n int) bool {
	return n%2 == 0
}

func FatorialRecursivo(n int) int {
	if n <= 1 {
		return 1
	}
	return n * FatorialRecursivo(n-1)
}

// ============================================================
// TESTES UNITÁRIOS
// Regra: função deve começar com "Test" e receber *testing.T
// ============================================================

func TestSomar(t *testing.T) {
	resultado := Somar(2, 3)
	esperado := 5

	if resultado != esperado {
		t.Errorf("Somar(2, 3) = %d; esperado %d", resultado, esperado)
	}
}

func TestSomarNegativo(t *testing.T) {
	if Somar(-1, -1) != -2 {
		t.Error("Soma de negativos falhou")
	}
}

func TestDividir(t *testing.T) {
	resultado, err := Dividir(10, 2)
	if err != nil {
		t.Fatalf("Dividir(10,2) retornou erro inesperado: %v", err)
	}
	if resultado != 5.0 {
		t.Errorf("Dividir(10, 2) = %.2f; esperado 5.0", resultado)
	}
}

func TestDividirPorZero(t *testing.T) {
	_, err := Dividir(10, 0)
	if err == nil {
		t.Error("Esperava erro ao dividir por zero, mas não veio!")
	}
}

// ============================================================
// TABLE-DRIVEN TESTS (Testes em tabela) - padrão Go
// Em vez de escrever 10 funções, escreva uma com uma tabela!
// ============================================================

func TestEhPar(t *testing.T) {
	// Tabela de casos de teste
	casos := []struct {
		numero   int
		esperado bool
	}{
		{numero: 2, esperado: true},
		{numero: 3, esperado: false},
		{numero: 0, esperado: true},
		{numero: -4, esperado: true},
		{numero: -5, esperado: false},
		{numero: 100, esperado: true},
	}

	for _, caso := range casos {
		t.Run(fmt.Sprintf("EhPar(%d)", caso.numero), func(t *testing.T) {
			resultado := EhPar(caso.numero)
			if resultado != caso.esperado {
				t.Errorf("EhPar(%d) = %v; esperado %v",
					caso.numero, resultado, caso.esperado)
			}
		})
	}
}

func TestFatorial(t *testing.T) {
	casos := []struct {
		entrada  int
		esperado int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	}

	for _, c := range casos {
		t.Run(fmt.Sprintf("Fatorial(%d)", c.entrada), func(t *testing.T) {
			resultado := FatorialRecursivo(c.entrada)
			if resultado != c.esperado {
				t.Errorf("Fatorial(%d) = %d; esperado %d", c.entrada, resultado, c.esperado)
			}
		})
	}
}

// ============================================================
// BENCHMARKS: mede a velocidade de uma função
// Regra: começa com "Benchmark" e recebe *testing.B
// ============================================================

func BenchmarkSomar(b *testing.B) {
	// O loop b.N é automaticamente ajustado pelo Go
	for i := 0; i < b.N; i++ {
		Somar(100, 200)
	}
}

func BenchmarkFatorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FatorialRecursivo(10)
	}
}

// ============================================================
// SETUP e TEARDOWN: preparar e limpar antes/depois dos testes
// ============================================================

func TestMain(m *testing.M) {
	// Roda ANTES de todos os testes
	fmt.Println("[Setup] Preparando testes...")

	// Executa os testes
	m.Run()

	// Roda DEPOIS de todos os testes
	fmt.Println("[Teardown] Limpando após testes...")
}

// main() não é necessário aqui — o go test cuida disso
func main() {
	fmt.Println("Execute os testes com: go test -v 03_testes_test.go")
	fmt.Println("")
	fmt.Println("Comandos úteis:")
	fmt.Println("  go test -v                     (todos os testes com detalhes)")
	fmt.Println("  go test -run TestSomar         (só o TestSomar)")
	fmt.Println("  go test -bench=. -benchmem     (benchmark com memória)")
	fmt.Println("  go test -cover                 (cobertura de código)")
}
