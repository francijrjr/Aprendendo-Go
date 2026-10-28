// ============================================================
// AULA AVANÇADA 2: Context (Contexto)
// Context = uma forma de passar informações e cancelar operações
// "Diga para todas as goroutines filhas: prazo máximo é X,
//  e se cancelar, todas param juntas"
// ============================================================
package main

import (
	"context"
	"fmt"
	"time"
)

// --------------------------------------------------------
// Simulação de operação demorada que respeita o contexto
// --------------------------------------------------------
func buscarDados(ctx context.Context, fonte string) (string, error) {
	// Simula uma busca que demora entre 1-3 segundos
	tempos := map[string]time.Duration{
		"banco":  2 * time.Second,
		"api":    3 * time.Second,
		"cache":  500 * time.Millisecond,
	}

	duracao := tempos[fonte]
	if duracao == 0 {
		duracao = 1 * time.Second
	}

	select {
	case <-time.After(duracao): // simulação do trabalho
		return fmt.Sprintf("dados de %s", fonte), nil

	case <-ctx.Done(): // contexto cancelado ou expirado!
		return "", fmt.Errorf("busca em '%s' cancelada: %v", fonte, ctx.Err())
	}
}

// --------------------------------------------------------
// Função com múltiplas buscas paralelas
// --------------------------------------------------------
func buscarTudo(ctx context.Context) {
	type resultado struct {
		fonte string
		dados string
		erro  error
	}

	fontes := []string{"cache", "banco", "api"}
	canal := make(chan resultado, len(fontes))

	// Disparar todas as buscas ao mesmo tempo
	for _, fonte := range fontes {
		go func(f string) {
			dados, err := buscarDados(ctx, f)
			canal <- resultado{fonte: f, dados: dados, erro: err}
		}(fonte)
	}

	// Coletar os resultados
	for range fontes {
		r := <-canal
		if r.erro != nil {
			fmt.Printf("  ❌ %s: %v\n", r.fonte, r.erro)
		} else {
			fmt.Printf("  ✅ %s: %s\n", r.fonte, r.dados)
		}
	}
}

// --------------------------------------------------------
// Passando valores no contexto
// --------------------------------------------------------
type chaveTipo string

const (
	chaveUsuario  chaveTipo = "usuario"
	chaveRequisicao chaveTipo = "requisicao_id"
)

func processarRequisicao(ctx context.Context) {
	// Ler valores do contexto
	usuario, ok := ctx.Value(chaveUsuario).(string)
	if !ok {
		usuario = "anônimo"
	}

	reqID, _ := ctx.Value(chaveRequisicao).(string)

	fmt.Printf("  Processando requisição %s para usuário: %s\n", reqID, usuario)
	
	// Simular trabalho
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("  Requisição %s concluída!\n", reqID)
}

// --------------------------------------------------------
// MAIN
// --------------------------------------------------------
func main() {
	// --------------------------------------------------------
	// 1. context.Background() = contexto raiz (sempre disponível)
	// --------------------------------------------------------
	fmt.Println("=== Context Básico ===")
	ctx := context.Background()
	
	dados, err := buscarDados(ctx, "cache")
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", dados)
	}

	// --------------------------------------------------------
	// 2. WithTimeout: cancela após X tempo
	// --------------------------------------------------------
	fmt.Println("\n=== WithTimeout (1 segundo) ===")
	
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelTimeout() // sempre faça defer do cancel!

	fmt.Println("Buscando com timeout de 1 segundo:")
	buscarTudo(ctxTimeout)

	// --------------------------------------------------------
	// 3. WithDeadline: cancela em um horário específico
	// --------------------------------------------------------
	fmt.Println("\n=== WithDeadline ===")
	
	prazo := time.Now().Add(800 * time.Millisecond)
	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), prazo)
	defer cancelDeadline()

	dados2, err2 := buscarDados(ctxDeadline, "cache") // cache leva 500ms, vai funcionar!
	if err2 != nil {
		fmt.Println("Cache falhou:", err2)
	} else {
		fmt.Println("Cache OK:", dados2)
	}

	dados3, err3 := buscarDados(ctxDeadline, "banco") // banco leva 2s, vai falhar!
	if err3 != nil {
		fmt.Println("Banco falhou:", err3)
	} else {
		fmt.Println("Banco OK:", dados3)
	}

	// --------------------------------------------------------
	// 4. WithCancel: cancela manualmente
	// --------------------------------------------------------
	fmt.Println("\n=== WithCancel ===")
	
	ctxCancel, cancel := context.WithCancel(context.Background())
	
	// Goroutine que cancela depois de 1 segundo
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("  (cancelando o contexto agora!)")
		cancel()
	}()

	dados4, err4 := buscarDados(ctxCancel, "api") // api leva 3s, vai ser cancelada!
	if err4 != nil {
		fmt.Println("API cancelada:", err4)
	} else {
		fmt.Println("API OK:", dados4)
	}

	// --------------------------------------------------------
	// 5. WithValue: passar dados pelo contexto
	// --------------------------------------------------------
	fmt.Println("\n=== WithValue ===")
	
	ctxBase := context.Background()
	ctxComUsuario := context.WithValue(ctxBase, chaveUsuario, "joao@email.com")
	ctxComTudo := context.WithValue(ctxComUsuario, chaveRequisicao, "REQ-001")

	processarRequisicao(ctxComTudo)
}
