package main

import "fmt"

var CurrentVersion = "dev"

func main() {
	fmt.Println("hello world")
	fmt.Println("Current version :", CurrentVersion)
}

// go build -ldflags="-X 'main.CurrentVersion=v1.0.0'" . para definir as flags de compilação
