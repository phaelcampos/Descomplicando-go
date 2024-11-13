package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var CurrentVersion = "desenvolvimento"

var packages = flag.String("packages", "kubectl", "Pacotes a serem instalados, separados por virgula. Valores possiveis: Kubectl")

func main() {
	flag.Parse()
	fmt.Println("Packages ", *packages)
	fmt.Println("Versão: ", CurrentVersion)

	packageArr := strings.Split(*packages, ",")

	for _, p := range packageArr {
		switch p {
		case "kubectl":
			{
				fmt.Println("Instalando kubectl")
				command, args := GetKubeCTLInstallCommand()
				cmd := exec.Command(command, args...)

				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
			}
		default:
			fmt.Println("Pacote", p, "não suportado pelo nosso script")
		}

	}

}
