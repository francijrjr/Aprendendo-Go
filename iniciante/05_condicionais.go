// ============================================================
// AULA 5: Condicionais (if, else, switch)
// "Se acontecer X, faça Y. Senão, faça Z."
// ============================================================
package main

import "fmt"

func main() {
	// --------------------------------------------------------
	// IF simples: "SE a condição for verdadeira, execute"
	// --------------------------------------------------------
	idade := 20

	if idade >= 18 {
		fmt.Println("Você é maior de idade!")
	}

	// --------------------------------------------------------
	// IF + ELSE: "SE... SENÃO..."
	// --------------------------------------------------------
	temperatura := 35.0

	if temperatura > 30 {
		fmt.Println("Está muito quente! Beba água.")
	} else {
		fmt.Println("A temperatura está agradável.")
	}

	// --------------------------------------------------------
	// IF + ELSE IF + ELSE: "SE... SENÃO SE... SENÃO..."
	// --------------------------------------------------------
	nota := 7.5

	if nota >= 9.0 {
		fmt.Println("Conceito A - Excelente!")
	} else if nota >= 7.0 {
		fmt.Println("Conceito B - Bom!")
	} else if nota >= 5.0 {
		fmt.Println("Conceito C - Regular.")
	} else {
		fmt.Println("Reprovado. Tente novamente!")
	}

	// --------------------------------------------------------
	// IF com inicialização: cria variável só dentro do if
	// --------------------------------------------------------
	if saldo := 150.0; saldo > 100.0 {
		fmt.Println("Saldo suficiente:", saldo)
	} else {
		fmt.Println("Saldo insuficiente:", saldo)
	}
	// "saldo" não existe aqui fora!

	// --------------------------------------------------------
	// SWITCH: quando você tem muitas condições para verificar
	// É mais limpo que vários else if
	// --------------------------------------------------------
	diaDaSemana := 3

	switch diaDaSemana {
	case 1:
		fmt.Println("Segunda-feira")
	case 2:
		fmt.Println("Terça-feira")
	case 3:
		fmt.Println("Quarta-feira")
	case 4:
		fmt.Println("Quinta-feira")
	case 5:
		fmt.Println("Sexta-feira")
	case 6, 7: // múltiplos valores no mesmo case!
		fmt.Println("Final de semana! 🎉")
	default: // se nenhum case bater
		fmt.Println("Dia inválido")
	}

	// --------------------------------------------------------
	// SWITCH com strings
	// --------------------------------------------------------
	animal := "cachorro"

	switch animal {
	case "gato":
		fmt.Println("Miau!")
	case "cachorro":
		fmt.Println("Au au!")
	case "pássaro":
		fmt.Println("Piu piu!")
	default:
		fmt.Println("Animal desconhecido")
	}

	// --------------------------------------------------------
	// SWITCH sem variável (funciona como if/else if)
	// --------------------------------------------------------
	hora := 14

	switch {
	case hora < 12:
		fmt.Println("Bom dia!")
	case hora < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
