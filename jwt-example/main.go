package main

import (
	"fmt"

	"github.com/phaelcampos/jwt-example/tokens"
)

func main() {
	fmt.Println("Gerando tokens")
	token := tokens.Generate()
	fmt.Println(token)
}
