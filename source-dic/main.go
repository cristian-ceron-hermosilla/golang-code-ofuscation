package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var reservedWords []string

func main() {

	if len(os.Args) != 1 {
		println("Usage: go run main.go <source_folder>")
		return
	}

	sourceFolder := os.Args[1]

	err := filepath.Walk(sourceFolder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".go") {
			genDictionary(path)
			return nil
		}
		return nil
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func genDictionary(sourceFile string) {

	println("Inspeccionando fuente ...", sourceFile)

	// Abrir el archivo fuente en formato texto
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, sourceFile, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Leer la lista de palabras reservadas y nombres de librerías desde un archivo
	reservedWords, err = readReservedWordsFromFile("../reservedWords.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Agregar los nombres de los métodos y funciones de las librerías importadas a la lista de palabras reservadas
	for _, importSpec := range file.Imports {
		importPath, _ := strconv.Unquote(importSpec.Path.Value)
		reservedWords = append(reservedWords, importPath)
	}

	// Recorrer el AST del archivo y agregar los identificadores a la lista de palabras reservadas
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:
			// Agregar el identificador a la lista de palabras reservadas
			if !isReservedWord(x.Name) {
				reservedWords = append(reservedWords, x.Name)
			}
		}
		return true
	})

	// Escribir la lista de palabras reservadas en el archivo
	err = writeReservedWordsToFile("../reservedWords.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func readReservedWordsFromFile(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	return lines, nil
}

func writeReservedWordsToFile(filename string) error {
	content := strings.Join(reservedWords, "\n")
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func isReservedWord(word string) bool {
	for _, reservedWord := range reservedWords {
		if word == reservedWord {
			return true
		}
	}
	return false
}
