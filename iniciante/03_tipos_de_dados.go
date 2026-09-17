// ============================================================
// AULA 3: Tipos de Dados
// Cada variável tem um "tipo" que diz o que ela pode guardar
// ============================================================
package main

import "fmt"

func main() {
	// --------------------------------------------------------
	// NÚMEROS INTEIROS (sem casas decimais)
	// --------------------------------------------------------
	var a int = 42          // int  = número inteiro (tamanho automático)
	var b int8 = 127        // int8 = de -128 até 127
	var c int16 = 32767     // int16 = de -32768 até 32767
	var d int32 = 2147483647 // int32 = número grande
	var e int64 = 9223372036854775807 // int64 = número enorme!
	
	fmt.Println("int:", a)
	fmt.Println("int8:", b)
	fmt.Println("int16:", c)
	fmt.Println("int32:", d)
	fmt.Println("int64:", e)

	// --------------------------------------------------------
	// NÚMEROS INTEIROS POSITIVOS (sem sinal negativo)
	// --------------------------------------------------------
	var positivo uint = 100  // uint = só números 0, 1, 2, 3...
	fmt.Println("uint:", positivo)

	// --------------------------------------------------------
	// NÚMEROS DECIMAIS (com casas decimais)
	// --------------------------------------------------------
	var preco float32 = 9.99    // float32 = decimal com menos precisão
	var salario float64 = 5000.50 // float64 = decimal mais preciso (preferido)
	
	fmt.Println("float32:", preco)
	fmt.Println("float64:", salario)

	// --------------------------------------------------------
	// TEXTO (string)
	// --------------------------------------------------------
	var saudacao string = "Olá, Mundo!"
	fmt.Println("string:", saudacao)
	fmt.Println("Tamanho do texto:", len(saudacao)) // len() conta os caracteres

	// Acessar letra por posição (começa em 0!)
	fmt.Println("Primeira letra:", string(saudacao[0]))

	// --------------------------------------------------------
	// VERDADEIRO ou FALSO (bool)
	// --------------------------------------------------------
	var ligado bool = true
	var desligado bool = false
	
	fmt.Println("ligado:", ligado)
	fmt.Println("desligado:", desligado)

	// --------------------------------------------------------
	// BYTE - representa um caractere ASCII (mesmo que uint8)
	// --------------------------------------------------------
	var letra byte = 'A' // Note: aspas simples para byte/rune
	fmt.Println("byte A =", letra) // vai mostrar 65 (código ASCII de 'A')

	// --------------------------------------------------------
	// RUNE - representa um caractere Unicode (para ç, ã, 中, etc.)
	// --------------------------------------------------------
	var acento rune = 'ç'
	fmt.Println("rune ç =", acento)

	// --------------------------------------------------------
	// Converter tipos (type conversion)
	// --------------------------------------------------------
	var inteiro int = 10
	var decimal float64 = float64(inteiro) // converte int para float64
	var textoNumero string = fmt.Sprintf("%d", inteiro) // int para string
	
	fmt.Println("int para float64:", decimal)
	fmt.Println("int para string:", textoNumero)
}
