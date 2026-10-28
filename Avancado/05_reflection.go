// ============================================================
// AULA AVANÇADA 5: Reflection (Reflexão)
// Reflection = inspecionar tipos e valores em tempo de execução
// "Seu programa pode se olhar no espelho e saber o que é!"
// Use com moderação — é poderoso mas mais lento
// ============================================================
package main

import (
	"fmt"
	"reflect"
)

// --------------------------------------------------------
// Structs de exemplo
// --------------------------------------------------------
type Endereco struct {
	Rua    string `json:"rua" validate:"required"`
	Cidade string `json:"cidade" validate:"required"`
	CEP    string `json:"cep" validate:"len=9"`
}

type Usuario struct {
	ID     int     `json:"id"`
	Nome   string  `json:"nome" validate:"required,min=3"`
	Email  string  `json:"email" validate:"required,email"`
	Idade  int     `json:"idade" validate:"min=0,max=150"`
	Ativo  bool    `json:"ativo"`
	End    Endereco `json:"endereco"`
}

// --------------------------------------------------------
// Funções usando reflection
// --------------------------------------------------------

// Descobre o tipo de qualquer valor
func tipoDeValor(v interface{}) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	fmt.Printf("Tipo: %-15s | Kind: %-10s | Valor: %v\n", 
		t.Name(), t.Kind(), val)
}

// Imprime todos os campos de uma struct dinamicamente
func inspecionarStruct(obj interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	// Se for ponteiro, pegar o elemento apontado
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		fmt.Println("Não é uma struct!")
		return
	}

	fmt.Printf("Struct: %s (%d campos)\n", t.Name(), t.NumField())
	
	for i := 0; i < t.NumField(); i++ {
		campo := t.Field(i)
		valor := v.Field(i)

		// Ler tags
		tagJSON := campo.Tag.Get("json")
		tagValidate := campo.Tag.Get("validate")

		fmt.Printf("  Campo: %-10s | Tipo: %-10s | Valor: %-15v | json: %-12s | validate: %s\n",
			campo.Name, campo.Type, valor.Interface(),
			tagJSON, tagValidate)
	}
}

// Modificar campo de uma struct via reflection (precisa de ponteiro!)
func definirCampo(obj interface{}, campo string, valor interface{}) error {
	v := reflect.ValueOf(obj)
	
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("precisa de ponteiro, recebeu %s", v.Kind())
	}

	v = v.Elem()
	f := v.FieldByName(campo)

	if !f.IsValid() {
		return fmt.Errorf("campo '%s' não encontrado", campo)
	}

	if !f.CanSet() {
		return fmt.Errorf("campo '%s' não pode ser modificado", campo)
	}

	novoValor := reflect.ValueOf(valor)
	if f.Type() != novoValor.Type() {
		return fmt.Errorf("tipo incompatível: esperado %s, recebeu %s", f.Type(), novoValor.Type())
	}

	f.Set(novoValor)
	return nil
}

// Validador simples usando tags de reflect
func validarCamposObrigatorios(obj interface{}) []string {
	var erros []string

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		campo := t.Field(i)
		valor := v.Field(i)
		tag := campo.Tag.Get("validate")

		// Verificar se a tag tem "required"
		if tag == "" {
			continue
		}

		for _, regra := range splitRegras(tag) {
			if regra == "required" {
				// Verificar se o campo está vazio
				if valor.Kind() == reflect.String && valor.String() == "" {
					erros = append(erros, fmt.Sprintf("campo '%s' é obrigatório", campo.Name))
				}
			}
		}
	}

	return erros
}

func splitRegras(tag string) []string {
	resultado := []string{}
	atual := ""
	for _, c := range tag {
		if c == ',' {
			resultado = append(resultado, atual)
			atual = ""
		} else {
			atual += string(c)
		}
	}
	if atual != "" {
		resultado = append(resultado, atual)
	}
	return resultado
}

// --------------------------------------------------------
// MAIN
// --------------------------------------------------------
func main() {
	// --------------------------------------------------------
	// Inspecionar tipos básicos
	// --------------------------------------------------------
	fmt.Println("=== Tipos Básicos ===")
	tipoDeValor(42)
	tipoDeValor(3.14)
	tipoDeValor("texto")
	tipoDeValor(true)
	tipoDeValor([]int{1, 2, 3})
	tipoDeValor(map[string]int{"a": 1})

	// --------------------------------------------------------
	// Inspecionar Struct
	// --------------------------------------------------------
	fmt.Println("\n=== Inspecionando Struct ===")
	usuario := Usuario{
		ID:    1,
		Nome:  "João",
		Email: "joao@email.com",
		Idade: 30,
		Ativo: true,
		End: Endereco{
			Rua:    "Rua das Flores, 100",
			Cidade: "São Paulo",
			CEP:    "01000-000",
		},
	}
	inspecionarStruct(usuario)

	// --------------------------------------------------------
	// Modificar campos via reflection
	// --------------------------------------------------------
	fmt.Println("\n=== Modificar via Reflection ===")
	fmt.Println("Nome antes:", usuario.Nome)
	
	err := definirCampo(&usuario, "Nome", "Maria Silva")
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Nome depois:", usuario.Nome)
	}

	err2 := definirCampo(&usuario, "Idade", 25)
	if err2 != nil {
		fmt.Println("Erro:", err2)
	} else {
		fmt.Println("Idade depois:", usuario.Idade)
	}

	// Tentativa de erro
	err3 := definirCampo(&usuario, "CampoInexistente", "valor")
	if err3 != nil {
		fmt.Println("Erro esperado:", err3)
	}

	// --------------------------------------------------------
	// Validação via tags
	// --------------------------------------------------------
	fmt.Println("\n=== Validação via Tags ===")
	
	usuarioVazio := Usuario{ID: 2} // sem nome e email!
	erros := validarCamposObrigatorios(usuarioVazio)
	
	if len(erros) > 0 {
		fmt.Println("Erros de validação:")
		for _, e := range erros {
			fmt.Println(" -", e)
		}
	}
	
	usuarioCompleto := Usuario{ID: 3, Nome: "Carlos", Email: "carlos@email.com", Ativo: true}
	erros2 := validarCamposObrigatorios(usuarioCompleto)
	if len(erros2) == 0 {
		fmt.Println("Usuário completo: sem erros!")
	}
}
