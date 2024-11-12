package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func dividir(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func subtrair(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func main() {
	var n1 int
	var n2 int
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	resposta := soma(n1, n2)
	fmt.Println("Soma:", resposta)

	resposta1 := dividir(n1, n2)
	fmt.Println("Soma:", resposta1)

	resposta2 := subtrair(n1, n2)
	fmt.Println("Subtração:", resposta2)

	resposta3 := multiplicar(n1, n2)
	fmt.Println("Multiplicação:", resposta3)

	// for i := 0; i < 100; i++ {
	// 	for j := 0; j < 100; j++ {
	// 		fmt.Println(i)
	// 		fmt.Println(j)

	// 		resposta := soma(i, j)
	// 		fmt.Println("Soma:", resposta)
	// 		resposta2 := dividir(i, j)
	// 		fmt.Println("Divisão:", resposta2)
	// 		resposta3 := subtrair(i, j)
	// 		fmt.Println("Subtração:", resposta3)
	// 		resposta4 := multiplicar(i, j)
	// 		fmt.Println("Multiplicação:", resposta4)
	// 	}
	// }

}
