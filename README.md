# golang-code-ofuscation
Ofuscación de codigo fuente Golang

## go-source-ofus
Permiten ofuscar el codigo fuente Golang. 

## go-source-dic
Permiten agregar al diccionario "reservedWords.txt" los comandos, funciones y/o metodos de las diversas librerias para que no sean ofuscadas ya que, esto generará errores de compilaciòn.

## En que consiste este proceso de "ofuscación"
Una vez compilado con "_build.sh" (linux/unix/mac) o "_build.bat" (windows), el programa "go-source-ofus" permite ofuscar todos los fuentes .go que se encuenten en una carpeta "origen" en una carpeta "destino".

## Mantención del diccionario
Al compilarse los binarios, idealmente con el batch "_build" que corresponda, el diccionario "reservedWords.txt" se moverá a la carpeta del usuario $HOME (linux/unix/mac) o %USERPROFILE% (windows).

Es posible que el programa "go-source-dic" o "go-source-dic.exe" segun corresponda, agregue al diccionario identificadores (variables o funciones) que no deban omitirse en el proceso de ofuscación; Puedes quitar del diccionario los identificadores que no desees omitir.


