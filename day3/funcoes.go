package main

import (
	"fmt"
)

func main1() {

	func() {
		fmt.Println("Preciso de uma função separada, pequena e que não vou usar muito, nem fora desse escopo")
	}()
	func(a int) {
		fmt.Println(a)
	}(5)

	f := func(a int) {
		fmt.Println(a)
	}

	f(5)
}

// Funções com parametricos variadicos
func soma(num int, valores ...int) int {
	for _, valor := range valores {
		num += valor
	}
	return num
}
func main2() {
	fmt.Println(soma(10))
	fmt.Println(soma(10, 5, 5))
	fmt.Println(soma(10, 1, 1, 1, 1))
	fmt.Println(soma(10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1))
}

// Funções com multiplos retornos
func soma2(num int, valores ...int) (int, int, error) {
	if num == 0 {
		return 0, 0, fmt.Errorf("Num não pode ser zero")
	}
	for _, valor := range valores {
		num += valor
	}
	return num, len(valores), nil
}

func main3() {
	n, l, err := soma2(10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(n)
		fmt.Println(l)
	}

	n, l, err = soma2(10)
}
