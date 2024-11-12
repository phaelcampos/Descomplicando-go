package main

import (
	"encoding/json"
	"fmt"
)

// Usando structs tags para definir o nome do campo quando você fizer o Marshal para o seu json!

// Chamamos de struct tags um campinho especial que podemos adicionar nos atributos de uma struct ao lado do tipo do atributo.
// Se você alterar a nossa struct Carro para adicionar uma struct tag json, o pacote encoding/json da standard library do Go entende isso e usa o valor da tag como chave quando você fizer o Marshal pro seu JSON!

// Exemplo:

// type Carro struct {
// 	Marca string `json:"minha tag"`
// 	Ano   intv   `json:"outra tag"`
// }

// Vira, com os valores da nossa variavel carro do vídeo anterior:

// { "minha tag" : "DMC DeLorean", "outra tag": "1982" }

// Um guia mais completo pode ser visto aqui: https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go-pt

type Veiculo1 interface {
	Buzinar1()
}

type Carro2 struct {
	Marca string `json:"marca"` // use tags para mudar o nome do campo
	Ano   int
}

func (c *Carro) Buzinar1() {
	fmt.Println("beep")
}

type Moto1 struct {
	Marca string
	Ano   int
}

func (m *Moto) Buzinar1() {
	fmt.Println("boop")
}

func buzina1(v Veiculo1) {
	v.Buzinar1()
}

func main() {
	carro := Carro{
		Marca: "DMC DeLorean",
		Ano:   1982,
	}
	data, err := json.Marshal(carro)
	if err != nil {
		fmt.Println("Error ::", err.Error())
	}
	fmt.Printf("%s\n", data)
	var moto Moto1
	err = json.Unmarshal([]byte("{\"Marca\": \"Cinquentinha\"}"), &moto)
	if err != nil {
		fmt.Println("Error ::", err.Error())
	}
	fmt.Printf("%s\n", moto.Marca)
}
