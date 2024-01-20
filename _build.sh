#!/bin/bash

current_dir=$(pwd)
echo "Chequeando si es la primera ejecución."
if [ ! -e $HOME/reservedWords.txt ]; then
    if [ -e reservedWords.txt ]; then
        echo "Es la primera ejecución."
        echo "Movemos la lista de Palabras Reservadas a la carpeta del usuario."
        mv reservedWords.txt $HOME/reservedWords.txt
    else
        echo "El archivo reservedWords.txt no se encuentra en el directorio actual."
        exit 1
    fi
else
    echo "No es la primera ejecución."
fi

echo "Generando binarios ejecutables..."

# Verifica si la carpeta execs existe, si no, la crea
if [ ! -d "execs" ]; then
    echo "Creamos una carpeta para los binarios generados."
    mkdir execs
    echo "Para setear la carpeta de binarios en el PATH necesitamos permisos administrativos."
    echo "export PATH=\"\$PATH:$(pwd)/execs\"" | sudo tee -a $HOME/.bashrc
fi

# Compila source-ofus y mueve el binario si la compilación es exitosa
cd source-ofus
if go build; then
    chmod a+x go-source-ofus
    mv go-source-ofus ../execs/.
    echo "go-source-ofus compilado y movido a execs."
else
    echo "Error al compilar go-source-ofus."
    exit 1
fi
cd ..

# Compila source-dic y mueve el binario si la compilación es exitosa
cd source-dic
if go build; then
    chmod a+x go-source-dic
    mv go-source-dic ../execs/.
    echo "go-source-ofus compilado y movido a execs."
else
    echo "Error al compilar go-source-dic."
    exit 1
fi
cd ..

echo "Binarios generados y movidos a carpeta [./exec]."
