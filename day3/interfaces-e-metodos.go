// Interfaces e métodos

// Nós não vamos entrar em todo o detalhe de OO aqui ainda - a gente vai ver isso mais pra frente! Mas pra começar a entender algumas das coisas que estamos fazendo precisamos falar de interfaces e métodos aqui!

package main

import (
	"fmt"
)

type Veiculo interface {
	Buzinar()
}

type Carro struct {
	Marca string // atributos
	Ano   int
}

type Moto struct {
	Modelo string // atributos
	Ano    int
}

func (c *Carro) Buzinar() {
	fmt.Println("beep")
}

func (m *Moto) Buzinar() {
	fmt.Println("boop")
}

func buzina(v Veiculo) {
	v.Buzinar()
}

func mainInterface() {
	carro := Carro{
		Marca: "DMC DeLorean",
		Ano:   1982,
	}
	moto := Moto{
		Modelo: "Cinquentinha",
		Ano:    1992,
	}
	veiculos := make([]Veiculo, 2)
	veiculos[0] = &carro
	veiculos[1] = &moto
	for i := 0; i < len(veiculos); i++ {
		buzina(veiculos[i])
	}

}
