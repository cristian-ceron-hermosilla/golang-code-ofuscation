package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var names = make(map[string]string)
var reservedWords []string

func main() {
	// Abrir el archivo fuente en formato texto
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Leer la lista de palabras reservadas y nombres de librerías desde un archivo
	reservedWords, err = readReservedWordsFromFile("reservedWords.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Agregar los nombres de los métodos y funciones de las librerías importadas a la lista de palabras reservadas
	for _, importSpec := range file.Imports {
		importPath, _ := strconv.Unquote(importSpec.Path.Value)
		reservedWords = append(reservedWords, importPath)
	}

	// Cambiar los nombres de las variables y funciones
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:
			// Cambiar el nombre de la variable o función
			if !isReservedWord(x.Name) {
				if _, ok := names[x.Name]; !ok {
					names[x.Name] = randomName()
				}
				x.Name = names[x.Name]
			}
		}
		return true
	})

	// Escribir el archivo fuente modificado en un nuevo archivo
	f, err := os.Create("newMain.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	err = printer.Fprint(f, fset, file)
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

func isReservedWord(word string) bool {
	for _, reservedWord := range reservedWords {
		if word == reservedWord {
			return true
		}
	}
	return false
}

func randomName() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
