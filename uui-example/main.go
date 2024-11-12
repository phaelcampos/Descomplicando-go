package main

import (
	"fmt"

	"github.com/jakehl/goid"
)

func main() {
	v4UUID := goid.NewV4UUID()
	fmt.Println(v4UUID)
}
