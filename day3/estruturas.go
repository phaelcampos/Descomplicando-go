package main

import (
	"fmt"
)

type Carro1 struct {
	Marca string //atributos
	Ano   int
}

func estruturas() {
	caracteres := []string{"a", "b", "c", "d", "e"} // slice
	caracteres = append(caracteres, "f")
	fmt.Println(caracteres)

	alfabeto := make(map[string]string)
	alfabeto["vogais"] = "aeiou"
	alfabeto["consoantes"] = "bcdfg..."
	fmt.Println(alfabeto)

	alfabeto = map[string]string{
		"vogais":     "aeiou",
		"consoantes": "bcdfg...",
	}
	fmt.Println(alfabeto)

	carro := Carro{
		Marca: "aaaaaa",
		Ano:   1982,
	}

	//Struct Anonima
	// carro := struct {
	// 	Marca string //atributos
	// 	Ano   int
	// }{
	// 	Marca: "Dmc DeLorean",
	// 	Ano:   1982,
	// }
	fmt.Println(carro)
	fmt.Println(carro.Marca)
}
