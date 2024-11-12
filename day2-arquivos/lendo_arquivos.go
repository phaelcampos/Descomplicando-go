package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// lendo byte a byte e "traduzindo" pra string
func main3() {
	file, err := os.Open("arquivo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				println(err.Error())
			}
			break
		}
		fmt.Print(string(b))
	}
}

// utilizando bufio para ler linha a linha do arquivo, e assim economizar memoria
func main2() {
	file, err := os.Open("arquivo.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// utilizando os.Open e OS.ReadFile
func main1() {
	// multiplos retornos
	//file, err := os.Open("arquivo.txt")
	//___________________________________________
	//Array = conteudo fixo, vc sabe o tamanho quando cria a variavel
	// Slice = conteudo variavel, voce pode aumentar / diminuir de tamanho

	file, err := os.ReadFile("arquivo.txt")

	//Tratamento de erro
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File contents:", string(file))
}
