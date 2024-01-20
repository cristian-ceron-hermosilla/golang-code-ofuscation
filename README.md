# golang-code-ofuscation
Ofuscación de código fuente Golang para ambiente Linux, Unix y Mac.

Próximamente, agregaré las correciones del build para Windows. Sin embargo, si estás construyendo algún producto o solución de software que pretendas distribuir y comercializar, te sugiero instalar una maquina Linux en VirtualBox y restringir al máximo la seguridad de tu VM para desarrollo.  

## source-ofus
Esta carpeta contiene una utilidad que permite "ofuscar" el código fuente Golang.

Este proceso de ofuscación básicamente consiste en modificar los nombres de variables y funciones en el código fuente, con el fin de aumentar la dificultad de lectura y análisis de tu codigo, en caso de caer en manos "No autorizadas".

## go-source-dic
Esta carpeta contiene el código de un programa auxiliar, cuyo objetivo es agregar al diccionario "reservedWords.txt" los comandos, funciones y/o métodos de las diversas librerías utilitarias del lenguaje para que sean IGNORADAS ya que, esto generará errores de compilaciòn.

## Como trabaja el programa "go-source-ofus"
Una vez compilado con "_build.sh" (linux/unix/mac) o "_build.bat" (windows), el programa "go-source-ofus" modifica TODOS los fuentes .go que se encuenten en tu carpeta de trabajo u "origen" en una carpeta de compilación o "destino".

     go-source-ofus ./proyecto-de-trajajo <./proyecto-ofuscado>

Si no colocas el 2do. parametro (una carpeta destino) o, por error, colocas la misma carpeta en ambos parametros se generará una carpeta "proyecto-de-trajajo-ofus".

Como asumimos que, si tu proyecto es hiper-ultra secreto, como todo proyecto que pretende comercializarse; Imaginamos que no pretenderas subirlo a Github en un repo privado ni mucho menos publico (yo tampoco lo haría), "go-source-ofus" antes de hacer lo que debe hacer, realizará un respaldo de tus fuentes en la carpeta ./backup/<fecha>/proyecto-de-trabajo/<timestamp>.

Te sugiero mantener SIEMPRE tu VM desconectada de la red e internet, salvo en los casos en que requieras actualizar o agregar dependencias a tu proyecto, tomando todos los resguardos necesarios con tu codigo y te sugiero utilizar fuentes ofuscados para agregar dependencias.

## Mantención del diccionario
Al compilarse los binarios, idealmente con el batch "_build" que corresponda, el diccionario "reservedWords.txt" se moverá a la carpeta del usuario $HOME (linux/unix/mac) o %USERPROFILE% (windows).

Es posible que el programa "go-source-dic" o "go-source-dic.exe" segun corresponda, agregue al diccionario identificadores (variables o funciones) que no deban omitirse en el proceso de ofuscación; Puedes quitar del diccionario los identificadores que no desees omitir.

## ¿Porque DEBERIA ofuscar mi codigo golang sí, al compilarlo se genera un binario en codigo de maquina?

DEBERIAS OFUSCARLO sí estas construyendo o distrubuyendo un programa que pretendes comercializar ya que; (uno) Es TU PROPIEDAD INTELECTUAL y (dos), SI existe la posibilidad que el binario incorpore tu codigo fuente al momento de la compilaciòn.

## Si, esta bien. Mi codigo fuente podría incorporarse en el binario pero, solo si yo se lo indico "explicitamente".

¿Estas seguro?.



