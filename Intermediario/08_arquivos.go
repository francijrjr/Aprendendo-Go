// ============================================================
// AULA INTERMEDIÁRIA 8: Leitura e Escrita de Arquivos
// Como ler e gravar arquivos em Go
// ============================================================
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// --------------------------------------------------------
	// ESCREVER em um arquivo
	// --------------------------------------------------------
	fmt.Println("=== Escrevendo arquivo ===")

	// os.WriteFile: forma mais simples para escrever texto
	conteudo := "Linha 1: Olá, mundo!\nLinha 2: Aprendendo Go!\nLinha 3: Arquivos são fáceis!\n"
	
	err := os.WriteFile("meu_arquivo.txt", []byte(conteudo), 0644)
	if err != nil {
		fmt.Println("Erro ao escrever:", err)
		return
	}
	fmt.Println("Arquivo criado com sucesso!")

	// --------------------------------------------------------
	// LER um arquivo inteiro
	// --------------------------------------------------------
	fmt.Println("\n=== Lendo arquivo inteiro ===")

	dados, err := os.ReadFile("meu_arquivo.txt")
	if err != nil {
		fmt.Println("Erro ao ler:", err)
		return
	}
	fmt.Println("Conteúdo:")
	fmt.Println(string(dados)) // converte []byte para string

	// --------------------------------------------------------
	// LER linha por linha (eficiente para arquivos grandes)
	// --------------------------------------------------------
	fmt.Println("=== Lendo linha por linha ===")

	arquivo, err := os.Open("meu_arquivo.txt") // abre só para leitura
	if err != nil {
		fmt.Println("Erro ao abrir:", err)
		return
	}
	defer arquivo.Close() // garante que fecha o arquivo ao final!

	scanner := bufio.NewScanner(arquivo)
	numero := 1
	for scanner.Scan() { // lê uma linha por vez
		linha := scanner.Text()
		fmt.Printf("Linha %d: %s\n", numero, linha)
		numero++
	}

	// --------------------------------------------------------
	// ADICIONAR conteúdo ao final (append)
	// --------------------------------------------------------
	fmt.Println("\n=== Adicionando ao arquivo ===")

	// os.O_APPEND = adicionar ao final, os.O_WRONLY = só escrita
	arquivoAppend, err := os.OpenFile("meu_arquivo.txt",
		os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir para append:", err)
		return
	}
	defer arquivoAppend.Close()

	_, err = arquivoAppend.WriteString("Linha 4: Adicionada depois!\n")
	if err != nil {
		fmt.Println("Erro ao adicionar:", err)
		return
	}
	fmt.Println("Linha adicionada!")

	// --------------------------------------------------------
	// VERIFICAR se arquivo existe
	// --------------------------------------------------------
	fmt.Println("\n=== Verificando arquivos ===")

	verificarArquivo("meu_arquivo.txt")
	verificarArquivo("arquivo_que_nao_existe.txt")

	// --------------------------------------------------------
	// LISTAR arquivos de uma pasta
	// --------------------------------------------------------
	fmt.Println("\n=== Listando arquivos da pasta atual ===")

	entradas, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Erro ao listar pasta:", err)
		return
	}

	for _, entrada := range entradas {
		if entrada.IsDir() {
			fmt.Println("  📁", entrada.Name())
		} else {
			fmt.Println("  📄", entrada.Name())
		}
	}

	// --------------------------------------------------------
	// DELETAR um arquivo
	// --------------------------------------------------------
	fmt.Println("\n=== Deletando arquivo ===")

	err = os.Remove("meu_arquivo.txt")
	if err != nil {
		fmt.Println("Erro ao deletar:", err)
		return
	}
	fmt.Println("Arquivo deletado!")

	// --------------------------------------------------------
	// Trabalhar com strings multilinhas
	// --------------------------------------------------------
	fmt.Println("\n=== Processando texto ===")
	
	textoGrande := `Nome: João
Idade: 30
Cidade: São Paulo
Profissão: Desenvolvedor`

	linhas := strings.Split(textoGrande, "\n")
	for i, linha := range linhas {
		partes := strings.Split(linha, ": ")
		if len(partes) == 2 {
			fmt.Printf("Campo %d: chave='%s', valor='%s'\n", i+1, partes[0], partes[1])
		}
	}
}

func verificarArquivo(caminho string) {
	_, err := os.Stat(caminho)
	if os.IsNotExist(err) {
		fmt.Printf("  '%s' NÃO existe\n", caminho)
	} else if err != nil {
		fmt.Printf("  Erro ao verificar '%s': %v\n", caminho, err)
	} else {
		fmt.Printf("  '%s' EXISTE!\n", caminho)
	}
}
