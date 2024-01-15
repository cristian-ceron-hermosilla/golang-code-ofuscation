# golang-code-ofuscation
Ofuscación de codigo fuente Golang

Si estás codificando una aplicación y deseas extremar las medidas para evitar que tu código fuente (y propiedad intelectual) se vea "expuesto" de alguna forma o derechamente, exista alguna posibilidad que personas no autorizadas por ti, lo pudiesen ver, copiar o vulnerar, tal vez esto te interese.

Si bien, es factible ofuscar el archivo binario (ejecutable) en el proceso de compilación, nada impide que en su "origen" es decir, en el proceso de compilación, los archivos fuentes de tu codigo sean incorporados al binario sin tu autorización.

Esa era mi preocupación y este programa viene de alguna forma a subsanar este problema, incorporando una capa de protección en ausencia de herramientas que permitan obfusca tu código fuente creado para Golang.

No es 100% infalible ya que, para generar el archivo binario ejecutable se requieren los fuentes y con esta herramienta solo, intentamos hacer un poco mas complicada la tarea de aquellos "ojos" indeseados que no tenemos como impedir que vean nuestro codigo si este llegase a ser el caso.

Claro está que otros lenguajes y plataformas permiten minimizar el codigo para evitar lo mismo que tratamos acá pero, esto se trata de soluciones creadas para ser compiladas con Golang y distribuidas como binarios ejecutables.

El código de este repositorio puede usarse y modificarse libremente y está motivado en la dificultad que tuve para encontrar herramientas que ofusquen las variables, identificadores y funciones.

Esta utilidad está disponible en este repositorio en su version fuente, legible y documentada para que la analices, la uses, la modifiques y la distribuyas sin ninguna restricción y espero realmente, si este fuese el caso, te sea de real utilidad.

Una vez que descargues o clones el repositorio, debes compilar el codigo para utilizarlo. A continuación te muestro como hacerlo si no tienes experiencia con Golang.  Por favor, revisa y trata de comprender el codigo antes de compilarlo y ejecutarlo y, si no estas seguro o no tienes confianza en utilizarlo, obviamente tienes toda la libertad para no utilizarlo y si este fuese el caso, te suguiero no utilizarlo. Si te da desconfianza no lo uses. De lo contrario, te reirero que, espero que sea de gran utilidad y que sea un real aporte en la proteccion de tu codigo y tu esfuerzo.

Requisitos:

- Disponer de una instalación de Golang en su ultima versión.

Compilación de la herramienta:




La ofuscación y compilación se realiza en dos fases:

1.- Se ofuscan la carpeta con los fuentes originales en una carpeta con los fuentes ofuscados que seran utilizados en la compilación y,

2.- La compilación de los fuentes desde la carpeta de archivos fuentes ofuscados; Para ello, se utilizará la libreria "garble" que permite ofuscar el binario generado.

Use:

ofus-source ./source-code ./source-ofus
cd ./source-ofus

-- Normal build
go build .

-- Build ofuscated binary. Require "go install mvdan.cc/garble@latest" 
garble build .


If you compile the offered source code found in the folder. " So, if your source code is included in the binary without your consent, the binary will have a much more difficult code to read and interpret.

100% guaranteed security does not exist but, this application leaves me a little quieter, I hope you too and I hope it serves you.


