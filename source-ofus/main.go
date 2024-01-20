package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var names = make(map[string]string)
var reservedWords []string

func main() {
	if len(os.Args) != 3 {
		println("Usage: go run main.go <source_folder> <destination_folder>")
		return
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error al obtener el directorio de inicio del usuario: %v\n", err)
		return
	}

	sourceFolder := os.Args[1]
	destinationFolder := os.Args[2]

	err = filepath.Walk(sourceFolder, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(sourceFolder, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destinationFolder, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		if strings.HasSuffix(path, ".go") {
			ofusca(homeDir, path, destPath)
			return nil
		}

		// Si no es un archivo .go, simplemente se copia
		return copyFile(path, destPath)
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func copyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return err
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func ofusca(homeDir string, sourceFile string, destFile string) {
	// Abrir el archivo fuente en formato texto
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, sourceFile, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Eliminar comentarios del archivo
	file.Comments = []*ast.CommentGroup{}

	// Leer la lista de palabras reservadas y nombres de librerías desde un archivo
	reservedWords, err = readReservedWordsFromFile(homeDir + "/reservedWords.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	println("Ofuscando ...", sourceFile, "-->", destFile)

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
	f, err := os.Create(destFile)
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
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")
	b := make([]rune, 28)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	firstLetter := string(letters[rand.Intn(50)])
	return firstLetter + string(b)
}
