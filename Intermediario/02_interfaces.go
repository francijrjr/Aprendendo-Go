// ============================================================
// AULA INTERMEDIÁRIA 2: Interfaces
// Interface = um CONTRATO: "qualquer um que tiver estes métodos
// pode ser chamado desta forma"
// ============================================================
package main

import (
	"fmt"
	"math"
)

// --------------------------------------------------------
// Definindo uma Interface
// --------------------------------------------------------

// A interface "Forma" diz: qualquer coisa que quiser ser uma Forma
// precisa ter o método Area() e Perimetro() que retornam float64.
type Forma interface {
	Area() float64
	Perimetro() float64
}

// --------------------------------------------------------
// Tipos concretos que IMPLEMENTAM a interface
// (não precisa declarar explicitamente, o Go descobre sozinho!)
// --------------------------------------------------------

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}

type Triangulo struct {
	Base   float64
	Altura float64
	Lado1  float64
	Lado2  float64
	Lado3  float64
}

// Retangulo implementa Forma
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}
func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Largura + r.Altura)
}

// Circulo implementa Forma
func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}
func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Raio
}

// Triangulo implementa Forma
func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}
func (t Triangulo) Perimetro() float64 {
	return t.Lado1 + t.Lado2 + t.Lado3
}

// --------------------------------------------------------
// Função que aceita QUALQUER Forma (polimorfismo!)
// --------------------------------------------------------
func imprimirInfo(f Forma) {
	fmt.Printf("  Área: %.2f | Perímetro: %.2f\n", f.Area(), f.Perimetro())
}

// --------------------------------------------------------
// Interface vazia: aceita QUALQUER tipo
// --------------------------------------------------------
func imprimirQualquer(valor interface{}) {
	fmt.Println("  Valor:", valor)
}

// --------------------------------------------------------
// MAIN
// --------------------------------------------------------
func main() {
	// --------------------------------------------------------
	// Usando a interface
	// --------------------------------------------------------
	fmt.Println("=== Formas Geométricas ===")
	
	ret := Retangulo{Largura: 5, Altura: 3}
	cir := Circulo{Raio: 4}
	tri := Triangulo{Base: 6, Altura: 4, Lado1: 5, Lado2: 5, Lado3: 6}

	fmt.Println("Retângulo:")
	imprimirInfo(ret)

	fmt.Println("Círculo:")
	imprimirInfo(cir)

	fmt.Println("Triângulo:")
	imprimirInfo(tri)

	// --------------------------------------------------------
	// Slice de interface: guardar tipos diferentes na mesma lista!
	// --------------------------------------------------------
	fmt.Println("\n=== Lista de Formas ===")
	formas := []Forma{
		Retangulo{Largura: 10, Altura: 2},
		Circulo{Raio: 3},
		Triangulo{Base: 4, Altura: 5, Lado1: 4, Lado2: 5, Lado3: 6},
	}

	for _, forma := range formas {
		imprimirInfo(forma)
	}

	// --------------------------------------------------------
	// Interface vazia (aceita qualquer coisa)
	// --------------------------------------------------------
	fmt.Println("\n=== Interface Vazia ===")
	imprimirQualquer(42)
	imprimirQualquer("texto")
	imprimirQualquer(3.14)
	imprimirQualquer([]int{1, 2, 3})

	// --------------------------------------------------------
	// Type Assertion: descobrir o tipo real
	// --------------------------------------------------------
	fmt.Println("\n=== Type Assertion ===")
	var f Forma = Circulo{Raio: 5}

	// Verificar se f é um Circulo
	if circulo, ok := f.(Circulo); ok {
		fmt.Println("É um círculo com raio:", circulo.Raio)
	}

	// --------------------------------------------------------
	// Type Switch: checar o tipo de interface{}
	// --------------------------------------------------------
	fmt.Println("\n=== Type Switch ===")
	valores := []interface{}{42, "olá", 3.14, true, Circulo{Raio: 2}}
	
	for _, v := range valores {
		switch t := v.(type) {
		case int:
			fmt.Println("int:", t)
		case string:
			fmt.Println("string:", t)
		case float64:
			fmt.Println("float64:", t)
		case bool:
			fmt.Println("bool:", t)
		default:
			fmt.Printf("tipo desconhecido: %T\n", t)
		}
	}
}
