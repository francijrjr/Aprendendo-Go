// ============================================================
// AULA INTERMEDIÁRIA 3: Tratamento de Erros
// Em Go, erros são VALORES normais — não são exceções!
// A regra: "se pode dar errado, retorne o erro junto"
// ============================================================
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// --------------------------------------------------------
// Criando erros simples
// --------------------------------------------------------

// Função que pode falhar — retorna (resultado, erro)
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("não é possível dividir por zero")
	}
	return a / b, nil // nil = "sem erro"
}

// --------------------------------------------------------
// Criando TIPOS de erro personalizados
// --------------------------------------------------------
type ErroDeSaque struct {
	Valor   float64
	Saldo   float64
	Mensagem string
}

// Para ser um error, precisa ter o método Error() string
func (e *ErroDeSaque) Error() string {
	return fmt.Sprintf("%s: tentou sacar %.2f mas o saldo é %.2f", 
		e.Mensagem, e.Valor, e.Saldo)
}

func sacar(saldo, valor float64) (float64, error) {
	if valor <= 0 {
		return saldo, errors.New("o valor do saque deve ser positivo")
	}
	if valor > saldo {
		return saldo, &ErroDeSaque{
			Valor:    valor,
			Saldo:    saldo,
			Mensagem: "saldo insuficiente",
		}
	}
	return saldo - valor, nil
}

// --------------------------------------------------------
// Verificar tipo específico de erro com errors.As
// --------------------------------------------------------
func processarSaque(saldo, valor float64) {
	novoSaldo, err := sacar(saldo, valor)
	
	if err != nil {
		// Verificar se é um ErroDeSaque específico
		var errSaque *ErroDeSaque
		if errors.As(err, &errSaque) {
			fmt.Printf("  ATENÇÃO: %s\n", err)
			fmt.Printf("  Faltam R$ %.2f para completar o saque\n", 
				errSaque.Valor - errSaque.Saldo)
		} else {
			fmt.Println("  Erro:", err)
		}
		return
	}
	
	fmt.Printf("  Saque de R$ %.2f realizado! Novo saldo: R$ %.2f\n", valor, novoSaldo)
}

// --------------------------------------------------------
// Encapsular erros com fmt.Errorf e %w (wrapping)
// --------------------------------------------------------
func converterParaInteiro(texto string) (int, error) {
	n, err := strconv.Atoi(texto)
	if err != nil {
		// %w = "empacota" o erro original dentro do nosso erro
		return 0, fmt.Errorf("falha ao converter '%s' para inteiro: %w", texto, err)
	}
	return n, nil
}

// --------------------------------------------------------
// MAIN
// --------------------------------------------------------
func main() {
	// --------------------------------------------------------
	// Tratando erros básicos
	// --------------------------------------------------------
	fmt.Println("=== Divisão ===")
	
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("10 / 2 =", resultado)
	}

	resultado2, err2 := dividir(10, 0)
	if err2 != nil {
		fmt.Println("Erro na divisão:", err2)
	} else {
		fmt.Println("10 / 0 =", resultado2)
	}

	// --------------------------------------------------------
	// Erros personalizados
	// --------------------------------------------------------
	fmt.Println("\n=== Saques ===")
	saldo := 500.0
	
	fmt.Printf("Saldo atual: R$ %.2f\n", saldo)
	
	processarSaque(saldo, 200.0) // deve funcionar
	processarSaque(saldo, 800.0) // deve falhar (saldo insuficiente)
	processarSaque(saldo, -50.0) // deve falhar (valor inválido)

	// --------------------------------------------------------
	// Erro encapsulado (wrapping)
	// --------------------------------------------------------
	fmt.Println("\n=== Conversão de texto ===")
	
	n1, err3 := converterParaInteiro("42")
	if err3 != nil {
		fmt.Println("Erro:", err3)
	} else {
		fmt.Println("Convertido:", n1)
	}
	
	n2, err4 := converterParaInteiro("abc")
	if err4 != nil {
		fmt.Println("Erro:", err4)
		
		// Desempacotar o erro original com errors.Unwrap
		erroOriginal := errors.Unwrap(err4)
		fmt.Println("Erro original:", erroOriginal)
	} else {
		fmt.Println("Convertido:", n2)
	}

	// --------------------------------------------------------
	// Padrão: múltiplos erros em sequência
	// --------------------------------------------------------
	fmt.Println("\n=== Processamento em cadeia ===")
	
	textos := []string{"10", "abc", "20", "xyz", "30"}
	total := 0
	
	for _, texto := range textos {
		n, err := converterParaInteiro(texto)
		if err != nil {
			fmt.Println("Pulando:", err)
			continue // pula para o próximo sem parar tudo
		}
		total += n
		fmt.Printf("  Adicionou %d ao total\n", n)
	}
	
	fmt.Println("Total dos válidos:", total)
}
