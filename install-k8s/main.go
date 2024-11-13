package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	// Instalar o kubectl
	//suportar linux e macOS
	// Exemplo!!

	// Linux:

	// MacOS: bew install kubectl
	fmt.Println("Instalando kubectl...")

	command, args := GetKubeCTLInstallCommand()
	cmd := exec.Command(command, args...)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
