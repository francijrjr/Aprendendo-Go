// ============================================================
// AULA INTERMEDIÁRIA 10: HTTP Server simples
// Go tem um servidor HTTP embutido — sem precisar de frameworks!
// ============================================================
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// --------------------------------------------------------
// Struct para as respostas em JSON
// --------------------------------------------------------
type Resposta struct {
	Mensagem string `json:"mensagem"`
	Horario  string `json:"horario"`
}

type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// --------------------------------------------------------
// "Banco de dados" em memória (lista de usuários)
// --------------------------------------------------------
var usuarios = []Usuario{
	{ID: 1, Nome: "João Silva", Email: "joao@email.com"},
	{ID: 2, Nome: "Maria Santos", Email: "maria@email.com"},
	{ID: 3, Nome: "Carlos Oliveira", Email: "carlos@email.com"},
}

// --------------------------------------------------------
// Handlers: funções que respondem às requisições HTTP
// --------------------------------------------------------

// GET / — Página inicial
func handleInicio(w http.ResponseWriter, r *http.Request) {
	resposta := Resposta{
		Mensagem: "Bem-vindo à API em Go! Acesse /usuarios para ver a lista.",
		Horario:  time.Now().Format("15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta)
}

// GET /usuarios — Retorna todos os usuários
func handleUsuarios(w http.ResponseWriter, r *http.Request) {
	// Só aceita GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// GET /saudacao?nome=João — Usa parâmetros da URL
func handleSaudacao(w http.ResponseWriter, r *http.Request) {
	// Pegar parâmetro da URL: /saudacao?nome=João
	nome := r.URL.Query().Get("nome")
	if nome == "" {
		nome = "Visitante"
	}

	resposta := Resposta{
		Mensagem: fmt.Sprintf("Olá, %s! Seja bem-vindo ao Go!", nome),
		Horario:  time.Now().Format("15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta)
}

// Middleware: adiciona log em todas as requisições
func comLog(proximo http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()
		fmt.Printf("[%s] %s %s\n", 
			inicio.Format("15:04:05"), r.Method, r.URL.Path)
		proximo(w, r)
		fmt.Printf("  -> Respondido em %v\n", time.Since(inicio))
	}
}

// --------------------------------------------------------
// MAIN: configure as rotas e inicie o servidor
// --------------------------------------------------------
func main() {
	// Registrar as rotas
	http.HandleFunc("/", comLog(handleInicio))
	http.HandleFunc("/usuarios", comLog(handleUsuarios))
	http.HandleFunc("/saudacao", comLog(handleSaudacao))

	porta := ":8080"
	fmt.Println("🚀 Servidor rodando em http://localhost" + porta)
	fmt.Println("")
	fmt.Println("Rotas disponíveis:")
	fmt.Println("  GET http://localhost:8080/")
	fmt.Println("  GET http://localhost:8080/usuarios")
	fmt.Println("  GET http://localhost:8080/saudacao?nome=SeuNome")
	fmt.Println("")
	fmt.Println("Pressione Ctrl+C para parar.")

	// Iniciar o servidor (fica bloqueado aqui)
	err := http.ListenAndServe(porta, nil)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}

// ============================================================
// Como testar:
// 1. Execute: go run 10_http_server.go
// 2. Abra o navegador ou use curl:
//    curl http://localhost:8080/
//    curl http://localhost:8080/usuarios
//    curl "http://localhost:8080/saudacao?nome=Maria"
// ============================================================
