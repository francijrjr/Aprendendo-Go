// ============================================================
// AULA AVANÇADA 6: Concorrência Avançada
// Padrões avançados para trabalhar com goroutines de forma segura
// ============================================================
package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ============================================================
// PADRÃO 1: Worker Pool (Pool de Trabalhadores)
// "Crie N trabalhadores que processam uma fila de tarefas"
// Evita criar goroutines demais e sobrecarregar o sistema
// ============================================================

type Tarefa struct {
	ID    int
	Dados string
}

type Resultado struct {
	TarefaID  int
	Resultado string
	Duracao   time.Duration
}

func workerPool(numWorkers int, tarefas <-chan Tarefa, resultados chan<- Resultado) {
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for tarefa := range tarefas { // processa até o canal fechar
				inicio := time.Now()
				
				// Simula trabalho
				time.Sleep(100 * time.Millisecond)
				resultado := fmt.Sprintf("Worker%d processou: %s", id, tarefa.Dados)
				
				resultados <- Resultado{
					TarefaID:  tarefa.ID,
					Resultado: resultado,
					Duracao:   time.Since(inicio),
				}
			}
		}(i)
	}

	// Fechar o canal de resultados quando todos os workers terminarem
	go func() {
		wg.Wait()
		close(resultados)
	}()
}

// ============================================================
// PADRÃO 2: Pipeline (Cadeia de Processamento)
// "Cada etapa recebe dados, processa e passa para a próxima"
// ============================================================

func gerarNumeros(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func dobrar(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * 2:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func filtrarMaioresQue(ctx context.Context, in <-chan int, limite int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n > limite {
				select {
				case out <- n:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

// ============================================================
// PADRÃO 3: Atomic Operations (Operações Atômicas)
// Mais eficiente que Mutex para contadores simples
// ============================================================

type ContadorAtomico struct {
	valor int64
}

func (c *ContadorAtomico) Incrementar() {
	atomic.AddInt64(&c.valor, 1)
}

func (c *ContadorAtomico) Decrementar() {
	atomic.AddInt64(&c.valor, -1)
}

func (c *ContadorAtomico) Valor() int64 {
	return atomic.LoadInt64(&c.valor)
}

// ============================================================
// PADRÃO 4: ErrGroup — gerenciar erros de múltiplas goroutines
// ============================================================

type GrupoDeErros struct {
	mu     sync.Mutex
	erros  []error
	wg     sync.WaitGroup
}

func (g *GrupoDeErros) Go(f func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := f(); err != nil {
			g.mu.Lock()
			g.erros = append(g.erros, err)
			g.mu.Unlock()
		}
	}()
}

func (g *GrupoDeErros) Aguardar() []error {
	g.wg.Wait()
	return g.erros
}

// ============================================================
// MAIN
// ============================================================
func main() {
	// --------------------------------------------------------
	// Worker Pool
	// --------------------------------------------------------
	fmt.Println("=== Worker Pool ===")

	tarefas := make(chan Tarefa, 10)
	resultados := make(chan Resultado, 10)

	// Inicia 3 workers
	go workerPool(3, tarefas, resultados)

	// Envia 8 tarefas
	for i := 1; i <= 8; i++ {
		tarefas <- Tarefa{ID: i, Dados: fmt.Sprintf("dado-%d", i)}
	}
	close(tarefas) // sinaliza que não virão mais tarefas

	// Coleta resultados
	for r := range resultados {
		fmt.Printf("  Tarefa %d: %s (%.0fms)\n",
			r.TarefaID, r.Resultado, r.Duracao.Seconds()*1000)
	}

	// --------------------------------------------------------
	// Pipeline
	// --------------------------------------------------------
	fmt.Println("\n=== Pipeline ===")

	ctx := context.Background()

	// Cadeia: gerar → dobrar → filtrar > 10
	numeros := gerarNumeros(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	dobrados := dobrar(ctx, numeros)
	filtrados := filtrarMaioresQue(ctx, dobrados, 10)

	fmt.Print("Números dobrados e maiores que 10: ")
	for n := range filtrados {
		fmt.Printf("%d ", n)
	}
	fmt.Println()

	// --------------------------------------------------------
	// Atomic Counter
	// --------------------------------------------------------
	fmt.Println("\n=== Contador Atômico ===")

	contador := &ContadorAtomico{}
	var wg sync.WaitGroup

	// 1000 goroutines incrementando sem usar mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador.Incrementar()
		}()
	}
	wg.Wait()
	fmt.Println("Contador (esperado 1000):", contador.Valor())

	// --------------------------------------------------------
	// Grupo de Erros
	// --------------------------------------------------------
	fmt.Println("\n=== Grupo de Erros ===")

	grupo := &GrupoDeErros{}

	// Simulação de tarefas que podem falhar
	tarefasSimuladas := []struct {
		nome  string
		falha bool
	}{
		{"buscar dados", false},
		{"processar imagem", true},
		{"salvar no banco", false},
		{"enviar email", true},
		{"limpar cache", false},
	}

	for _, t := range tarefasSimuladas {
		t := t // captura para closure
		grupo.Go(func() error {
			time.Sleep(50 * time.Millisecond) // simula trabalho
			if t.falha {
				return fmt.Errorf("tarefa '%s' falhou", t.nome)
			}
			fmt.Printf("  ✅ '%s' concluída\n", t.nome)
			return nil
		})
	}

	erros := grupo.Aguardar()
	if len(erros) > 0 {
		fmt.Println("\nErros encontrados:")
		for _, err := range erros {
			fmt.Println("  ❌", err)
		}
	}
}
