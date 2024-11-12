package main

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
)

func percorreDiretorio(path string, d fs.DirEntry, err error) error {
	fmt.Println("é diretório?", d.IsDir())
	info, _ := d.Info()
	fmt.Println("Name::", info.Name())
	fmt.Println("---------------\n")
	return nil
}

func main() {
	path.Join("junta", "o caminho")
	filepath.WalkDir(".", percorreDiretorio)
}
