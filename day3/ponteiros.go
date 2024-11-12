//Valores e ponteiros, e um exemplo de como isso afeta arrays e slices

// Um ponteiro é uma variável que guarda o local da memória que um valor está guardado.

// Cada variável é guardada em um ou mais espaços subsequentes de memória, chamados endereços.

package main

import "fmt"

func mainPonteiro1() {

	// Conteudo: 0 0 0 0 0 0 1 0 | 0 | 0 0  0  0  | 0  0   0  8 |
	// Endereço: 0 1 2 3 4 5 6 7 | 8 | 9 10 11 12 | 13 14 15 16 |
	// Variavel: ---- x -------- | b | -- px --   | --- pb -----|

	// px = 0
	// pb = 8
	// x = 10
	// b = 0 / false

	var x int64
	x = 10 // 8 bytes
	var b bool
	b = false // 1 byte

	pointerOfX := &x
	pointerOfB := &b

	fmt.Println(x)
	fmt.Println(b)
	fmt.Println(pointerOfX)
	fmt.Println(pointerOfB)
}

// Quando usar ponteiros

//     Especialmente quando você precisar passar valores pra funções que recebem uma interface =D

// Como slices, arrays e ponteiros estão interligados

func mainPonteiro2() {
	alfabeto := []string{"a", "b", "c", "d"}
	fmt.Println("alfabeto:", alfabeto)

	var primeirasLetras = alfabeto[0:1]
	primeirasLetras[0] = "z"

	fmt.Println("primeirasLetras:", primeirasLetras)
	fmt.Println("alfabeto:", alfabeto)

	alfabeto = append(alfabeto, "e")
	primeirasLetras[0] = "a"
	fmt.Println("primeirasLetras:", primeirasLetras)
	fmt.Println("alfabeto:", alfabeto)
}
