package main

import (
	"fmt"
	"os"
)

// var perm fs.FileMode
// perm = 0666
// os.WriteFile("arquivo-novo.txt", txt, perm)

// file, err = os.Create("Arquivo-novo2.txt")
//
//	txt := []byte("Esse eh o conteudo do meu arquivo novo")
func main5() {
	file, err := os.OpenFile("medicamento.csv", os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("\n--------- \n")
	file.WriteString("Minha String")
}
