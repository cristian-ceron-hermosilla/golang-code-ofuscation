# golang-code-ofuscation
Ofuscación de codigo fuente Golang para ambiente Linux, Unix y Mac.

Proximamente, agregare las correciones del build para Windows. Sin embargo, si estas construyendo algún producto o solución de software que pretender distribuir y comercializar, te sugiero instalar una maquina Linux en VirtualBox y restringir al maximo la seguridad de tu VM para desarrollo.  

## source-ofus
Esta carpeta contiene una utilidad que permiten "ofuscar" el codigo fuente Golang.

Este proceso de ofuscación basicamente consiste en modificar nombres de variables y funciones en el codigo fuente con el fin de aumentar la dificultad de lectura y analisis de tu codigo, en caso de caer en manos "No autorizadas".

## go-source-dic
Permiten agregar al diccionario "reservedWords.txt" los comandos, funciones y/o metodos de las diversas librerias para que no sean ofuscadas ya que, esto generará errores de compilaciòn.

## En que consiste este proceso de "ofuscación"
Una vez compilado con "_build.sh" (linux/unix/mac) o "_build.bat" (windows), el programa "go-source-ofus" permite ofuscar todos los fuentes .go que se encuenten en una carpeta "origen" en una carpeta "destino".

## Mantención del diccionario
Al compilarse los binarios, idealmente con el batch "_build" que corresponda, el diccionario "reservedWords.txt" se moverá a la carpeta del usuario $HOME (linux/unix/mac) o %USERPROFILE% (windows).

Es posible que el programa "go-source-dic" o "go-source-dic.exe" segun corresponda, agregue al diccionario identificadores (variables o funciones) que no deban omitirse en el proceso de ofuscación; Puedes quitar del diccionario los identificadores que no desees omitir.

## ¿Porque DEBERIA ofuscar mi codigo golang sí, al compilarlo se genera un binario en codigo de maquina?

DEBERIAS OFUSCARLO sí estas construyendo o distrubuyendo un programa que pretendes comercializar ya que; (uno) Es TU PROPIEDAD INTELECTUAL y (dos), SI existe la posibilidad que el binario incorpore tu codigo fuente al momento de la compilaciòn.

## Si, esta bien. Mi codigo fuente podría incorporarse en el binario pero, solo si yo se lo indico "explicitamente".

¿Estas seguro?.



