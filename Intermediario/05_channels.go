// ============================================================
// AULA INTERMEDIÁRIA 5: Channels (Canais)
// Channel = um cano para comunicação segura entre goroutines
// "Goroutine A manda dados → cano → Goroutine B recebe"
// ============================================================
package main

import (
	"fmt"
	"time"
)

func main() {
	// --------------------------------------------------------
	// Channel básico: enviar e receber
	// --------------------------------------------------------
	fmt.Println("=== Channel Básico ===")

	// make(chan tipo) = cria um canal para o tipo especificado
	mensagens := make(chan string)

	// Goroutine que ENVIA uma mensagem
	go func() {
		time.Sleep(1 * time.Second)
		mensagens <- "Olá do outro lado!" // <- envia para o canal
	}()

	// Main espera e RECEBE a mensagem
	msg := <-mensagens // <- recebe do canal (fica bloqueado até chegar algo)
	fmt.Println("Recebido:", msg)

	// --------------------------------------------------------
	// Channel com buffer: não bloqueia imediatamente
	// --------------------------------------------------------
	fmt.Println("\n=== Channel com Buffer ===")

	// Canal com capacidade para 3 mensagens sem bloquear
	bufferizado := make(chan int, 3)

	// Pode enviar sem goroutine (não bloqueia porque tem espaço)
	bufferizado <- 10
	bufferizado <- 20
	bufferizado <- 30

	fmt.Println("Recebido:", <-bufferizado) // 10
	fmt.Println("Recebido:", <-bufferizado) // 20
	fmt.Println("Recebido:", <-bufferizado) // 30

	// --------------------------------------------------------
	// Fechar um canal: diz "não vai mais chegar nada"
	// --------------------------------------------------------
	fmt.Println("\n=== Fechar Canal ===")

	numeros := make(chan int, 5)

	// Produz números e fecha o canal
	go func() {
		for i := 1; i <= 5; i++ {
			numeros <- i
		}
		close(numeros) // fecha o canal após enviar tudo
	}()

	// range no canal: lê até o canal ser fechado
	for n := range numeros {
		fmt.Println("Número:", n)
	}

	// --------------------------------------------------------
	// SELECT: espera por múltiplos canais ao mesmo tempo
	// --------------------------------------------------------
	fmt.Println("\n=== SELECT ===")

	canal1 := make(chan string, 1)
	canal2 := make(chan string, 1)

	go func() {
		time.Sleep(200 * time.Millisecond)
		canal1 <- "mensagem do canal 1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		canal2 <- "mensagem do canal 2"
	}()

	// Espera o que vier primeiro
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-canal1:
			fmt.Println("Canal 1:", msg1)
		case msg2 := <-canal2:
			fmt.Println("Canal 2:", msg2)
		}
	}

	// --------------------------------------------------------
	// SELECT com timeout: "espera até X segundos"
	// --------------------------------------------------------
	fmt.Println("\n=== SELECT com Timeout ===")

	lento := make(chan string)

	go func() {
		time.Sleep(3 * time.Second) // muito lento!
		lento <- "resposta lenta"
	}()

	select {
	case resposta := <-lento:
		fmt.Println("Resposta recebida:", resposta)
	case <-time.After(1 * time.Second): // espera no máximo 1 segundo
		fmt.Println("Timeout! A operação demorou demais.")
	}

	// --------------------------------------------------------
	// Padrão Fan-out: distribuir trabalho para várias goroutines
	// --------------------------------------------------------
	fmt.Println("\n=== Fan-out (distribuir trabalho) ===")

	tarefas := make(chan int, 10)
	resultados := make(chan int, 10)

	// 3 "trabalhadores" pegando tarefas do mesmo canal
	for w := 1; w <= 3; w++ {
		go func(id int) {
			for tarefa := range tarefas {
				resultado := tarefa * tarefa // calcula o quadrado
				fmt.Printf("  Trabalhador %d: %d² = %d\n", id, tarefa, resultado)
				resultados <- resultado
			}
		}(w)
	}

	// Envia 9 tarefas
	for i := 1; i <= 9; i++ {
		tarefas <- i
	}
	close(tarefas)

	// Coleta os 9 resultados
	total := 0
	for i := 0; i < 9; i++ {
		total += <-resultados
	}
	fmt.Println("Soma de todos os quadrados:", total)
}
