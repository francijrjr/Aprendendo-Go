// ============================================================
// AULA AVANÇADA 4: Design Patterns (Padrões de Design)
// Os padrões mais úteis do dia a dia em Go
// ============================================================
package main

import (
	"fmt"
	"sync"
)

// ============================================================
// PADRÃO 1: SINGLETON
// "Garante que só existe UMA instância de um objeto no programa"
// Exemplo: conexão com banco de dados, configurações globais
// ============================================================

type Configuracao struct {
	Host     string
	Porta    int
	Ambiente string
}

var (
	instanciaConfig *Configuracao
	onceConfig      sync.Once
)

func GetConfiguracao() *Configuracao {
	onceConfig.Do(func() { // garante que executa UMA vez só
		instanciaConfig = &Configuracao{
			Host:     "localhost",
			Porta:    5432,
			Ambiente: "producao",
		}
		fmt.Println("  [Singleton] Configuração criada pela primeira vez!")
	})
	return instanciaConfig
}

// ============================================================
// PADRÃO 2: FACTORY (Fábrica)
// "Cria objetos sem precisar saber qual tipo exato vai criar"
// ============================================================

type Animal interface {
	Falar() string
	Nome() string
}

type Cachorro struct{ nome string }
type Gato struct{ nome string }
type Passaro struct{ nome string }

func (c *Cachorro) Falar() string { return "Au au!" }
func (g *Gato) Falar() string    { return "Miau!" }
func (p *Passaro) Falar() string { return "Piu piu!" }

func (c *Cachorro) Nome() string { return c.nome }
func (g *Gato) Nome() string    { return g.nome }
func (p *Passaro) Nome() string { return p.nome }

// Fábrica: devolve o animal certo baseado no tipo
func CriarAnimal(tipo, nome string) (Animal, error) {
	switch tipo {
	case "cachorro":
		return &Cachorro{nome: nome}, nil
	case "gato":
		return &Gato{nome: nome}, nil
	case "passaro":
		return &Passaro{nome: nome}, nil
	default:
		return nil, fmt.Errorf("tipo de animal desconhecido: %s", tipo)
	}
}

// ============================================================
// PADRÃO 3: OBSERVER (Observador)
// "Quando algo muda, notifique automaticamente os interessados"
// Exemplo: eventos, notificações, pub/sub
// ============================================================

type Observador interface {
	Notificar(evento string, dados interface{})
}

type GerenciadorDeEventos struct {
	observadores map[string][]Observador
	mu           sync.RWMutex
}

func NovoGerenciador() *GerenciadorDeEventos {
	return &GerenciadorDeEventos{
		observadores: make(map[string][]Observador),
	}
}

func (g *GerenciadorDeEventos) Inscrever(evento string, obs Observador) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.observadores[evento] = append(g.observadores[evento], obs)
}

func (g *GerenciadorDeEventos) Publicar(evento string, dados interface{}) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	for _, obs := range g.observadores[evento] {
		obs.Notificar(evento, dados)
	}
}

// Observadores concretos
type LogObservador struct{ nome string }
type EmailObservador struct{ email string }

func (l *LogObservador) Notificar(evento string, dados interface{}) {
	fmt.Printf("  [LOG - %s] Evento: %s | Dados: %v\n", l.nome, evento, dados)
}

func (e *EmailObservador) Notificar(evento string, dados interface{}) {
	fmt.Printf("  [EMAIL para %s] Evento '%s' ocorreu: %v\n", e.email, evento, dados)
}

// ============================================================
// PADRÃO 4: OPTION PATTERN (Opções Funcionais)
// "Configure um objeto de forma flexível e legível"
// Muito usado em bibliotecas Go modernas
// ============================================================

type ServidorHTTP struct {
	host    string
	porta   int
	timeout int
	maxConn int
}

type OpcaoServidor func(*ServidorHTTP)

func ComHost(host string) OpcaoServidor {
	return func(s *ServidorHTTP) { s.host = host }
}

func ComPorta(porta int) OpcaoServidor {
	return func(s *ServidorHTTP) { s.porta = porta }
}

func ComTimeout(segundos int) OpcaoServidor {
	return func(s *ServidorHTTP) { s.timeout = segundos }
}

func ComMaxConexoes(max int) OpcaoServidor {
	return func(s *ServidorHTTP) { s.maxConn = max }
}

func NovoServidor(opcoes ...OpcaoServidor) *ServidorHTTP {
	// Valores padrão
	s := &ServidorHTTP{
		host:    "localhost",
		porta:   8080,
		timeout: 30,
		maxConn: 100,
	}
	// Aplica cada opção
	for _, opcao := range opcoes {
		opcao(s)
	}
	return s
}

// ============================================================
// MAIN
// ============================================================
func main() {
	// --------------------------------------------------------
	// Singleton
	// --------------------------------------------------------
	fmt.Println("=== SINGLETON ===")
	c1 := GetConfiguracao()
	c2 := GetConfiguracao() // não cria de novo!
	c3 := GetConfiguracao() // ainda a mesma instância!
	
	fmt.Printf("c1 == c2? %v\n", c1 == c2) // true
	fmt.Printf("c2 == c3? %v\n", c2 == c3) // true
	fmt.Printf("Host: %s | Porta: %d\n", c1.Host, c1.Porta)

	// --------------------------------------------------------
	// Factory
	// --------------------------------------------------------
	fmt.Println("\n=== FACTORY ===")
	tipos := []struct{ tipo, nome string }{
		{"cachorro", "Rex"},
		{"gato", "Whiskers"},
		{"passaro", "Tweety"},
		{"peixe", "Nemo"},
	}

	for _, t := range tipos {
		animal, err := CriarAnimal(t.tipo, t.nome)
		if err != nil {
			fmt.Println("  Erro:", err)
			continue
		}
		fmt.Printf("  %s diz: %s\n", animal.Nome(), animal.Falar())
	}

	// --------------------------------------------------------
	// Observer
	// --------------------------------------------------------
	fmt.Println("\n=== OBSERVER ===")
	gerenciador := NovoGerenciador()

	// Registrar observadores
	gerenciador.Inscrever("compra", &LogObservador{nome: "sistema"})
	gerenciador.Inscrever("compra", &EmailObservador{email: "vendas@loja.com"})
	gerenciador.Inscrever("erro", &LogObservador{nome: "monitoramento"})

	// Publicar eventos
	gerenciador.Publicar("compra", map[string]interface{}{"produto": "Notebook", "valor": 3499.90})
	gerenciador.Publicar("erro", "Falha na conexão com banco de dados")

	// --------------------------------------------------------
	// Option Pattern
	// --------------------------------------------------------
	fmt.Println("\n=== OPTION PATTERN ===")

	// Servidor com configuração mínima (usa padrões)
	s1 := NovoServidor()
	fmt.Printf("Servidor 1: %s:%d (timeout: %ds, max: %d)\n",
		s1.host, s1.porta, s1.timeout, s1.maxConn)

	// Servidor totalmente configurado
	s2 := NovoServidor(
		ComHost("0.0.0.0"),
		ComPorta(9090),
		ComTimeout(60),
		ComMaxConexoes(1000),
	)
	fmt.Printf("Servidor 2: %s:%d (timeout: %ds, max: %d)\n",
		s2.host, s2.porta, s2.timeout, s2.maxConn)
}
