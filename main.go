package main

import "fmt"

// Função privada
func dummyFunciton(s string) string {
	return fmt.Sprintf("Run Dummy funtion. %s \n", s)
}

// Função publica
func DummyPublicFunction(s string) string {
	return dummyFunciton(s)
}

func main() {
	//Atribuição curta de váriavel (declaração + atribuição)
	//Short assignment
	s := dummyFunciton("teste")
	fmt.Println(s)

	var u string
	u = dummyFunciton("hello world again!")
	fmt.Println(u)

	//Controle de fluxo
	//if, else, switch
	var input int
	fmt.Scan(&input)
	if input == 1 {
		fmt.Println("é 1")
	} else {
		fmt.Println(("Não é 1"))
	}

	switch input {
	case 1:
		fmt.Println("é 1")
	case 2:
		fmt.Println("é 2")
	default:
		fmt.Println("Não é 1 nem 2")
	}

	//Estruturas de repetição
	//for
	//loop
	for i := 0; i < 10; i++ {

	}

	//WHILE
	running := true
	for running {
		if input == 3 {
			running = false
		}
	}
	fmt.Println(input)
}
