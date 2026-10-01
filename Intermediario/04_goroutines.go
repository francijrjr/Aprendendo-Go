// ============================================================
// AULA INTERMEDIÁRIA 4: Goroutines (Paralelismo)
// Goroutine = uma tarefa leve que roda "ao mesmo tempo" que outras
// "Faça várias coisas ao mesmo tempo sem bloquear o programa!"
// ============================================================
package main

import (
	"fmt"
	"sync"
	"time"
)

// --------------------------------------------------------
// Função simples que simula um trabalho demorado
// --------------------------------------------------------
func trabalhar(nome string, duracao time.Duration) {
	fmt.Printf("  %s começou a trabalhar\n", nome)
	time.Sleep(duracao) // simula trabalho demorado
	fmt.Printf("  %s terminou!\n", nome)
}

// --------------------------------------------------------
// MAIN
// --------------------------------------------------------
func main() {
	// --------------------------------------------------------
	// SEM goroutine: um de cada vez (sequencial)
	// --------------------------------------------------------
	fmt.Println("=== Sem Goroutine (sequencial) ===")
	inicio := time.Now()
	
	trabalhar("Tarefa A", 2*time.Second)
	trabalhar("Tarefa B", 2*time.Second)
	trabalhar("Tarefa C", 2*time.Second)
	
	fmt.Printf("Tempo total: %.1f segundos\n\n", time.Since(inicio).Seconds())

	// --------------------------------------------------------
	// COM goroutine: ao mesmo tempo (paralelo)
	// go + função = roda em segundo plano
	// --------------------------------------------------------
	fmt.Println("=== Com Goroutine (paralelo) ===")
	inicio = time.Now()

	// WaitGroup = "espere estas goroutines terminarem antes de continuar"
	var wg sync.WaitGroup

	wg.Add(1) // diz "vou criar 1 goroutine"
	go func() {
		defer wg.Done() // quando esta função terminar, avisa o WaitGroup
		trabalhar("Tarefa A", 2*time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		trabalhar("Tarefa B", 2*time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		trabalhar("Tarefa C", 2*time.Second)
	}()

	wg.Wait() // espera todas as goroutines terminarem
	fmt.Printf("Tempo total: %.1f segundos\n\n", time.Since(inicio).Seconds())
	// Note que agora leva ~2 segundos em vez de ~6!

	// --------------------------------------------------------
	// Goroutines com resultados usando Channel
	// Channel = um "cano" para passar dados entre goroutines
	// --------------------------------------------------------
	fmt.Println("=== Goroutines com Channel ===")

	resultados := make(chan string, 3) // canal que aguenta 3 mensagens

	go func() {
		time.Sleep(1 * time.Second)
		resultados <- "Resultado da Goroutine 1" // envia no canal
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		resultados <- "Resultado da Goroutine 2"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		resultados <- "Resultado da Goroutine 3"
	}()

	// Receber os 3 resultados
	for i := 0; i < 3; i++ {
		msg := <-resultados // recebe do canal (espera se necessário)
		fmt.Println(" Recebido:", msg)
	}

	// --------------------------------------------------------
	// Mutex: proteger dados compartilhados entre goroutines
	// --------------------------------------------------------
	fmt.Println("\n=== Mutex (proteção de dados) ===")

	var mu sync.Mutex
	contador := 0
	var wg2 sync.WaitGroup

	// 100 goroutines tentando incrementar o contador ao mesmo tempo
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()   // trava o acesso — só uma goroutine de cada vez!
			contador++
			mu.Unlock() // libera para a próxima
		}()
	}

	wg2.Wait()
	fmt.Println("Contador final (esperado 100):", contador)
}
