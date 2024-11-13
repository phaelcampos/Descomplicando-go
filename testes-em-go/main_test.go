package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	resultado := soma(10, 5)
	esperado := 15
	if resultado != esperado {
		t.Error("Resultado diferente do esperado")
	}
}
