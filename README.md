# golang-code-ofuscation
Ofuscación de codigo fuente Golang

Si estás codificando una aplicación que quieres empaquetar en un binario (ejecutable) para comercializar y/o distribuir para hacerte millonario :-) y te da PAVOR que Micro$oft o Google te roben tu codigo, quizas esto te pueda interesar.

El objetivo de esta "utilidad" es ayudarte a extremar las medidas para ayudarte a evitar 
que tu código fuente (y propiedad intelectual) se vea "expuesto" de alguna forma o derechamente, que personas no autorizadas por ti, lo pudiesen ver, copiar, piratear y robar (suena fuerte pero, la historia está llena de millonarios que han robado el producto de otros).

Es MUY importante que sepas que, es posible ofuscar tu codigo golang PERO, aca tenemos dos sorpresillas:

1.- El codigo factible de ofuscar es el codigo binario (no el codigo fuente utilizado para generar el binarios) y, 

2.- (aca viene la otra "sorpresilla) Tu valioso codigo fuerte PODRIA SER INCLUIDO en el binario (como te quedó el ojo? :-)

Claro, claro, podría ser incluido en el binario utilizando "explicitamente" un parametro de la compilación PERO, ¿estas 100% seguro que el codigo fuente no se incluye en el binario, incluso si no lo indicas?.

Y claro, si el codigo fuente se podria INCLUIR en el binario, es obvio que se pueda recuperar del binario con algun otro comando, conocido por todos los usuarios y/o conocido SOLO por los desarrolladores de los compiladores (Micro$oft, Google, etc.).

Quizas a veces, los que intentamos crear alguna solucion nos diluimos creando mecanismos para evitar que nos copien el codigo y olvidamos de la solucion y con el codigo de este repositorio, pretendo aportar un software muy simple que te permitirá "ofuscar" u ocultar algunas caracteristicas de tu codigo fuente.

No es posible evitar que el proceso de compilación no lo incluya en el binario, queramos o no pero, podemos ocultar y enredar algunas partes de este para hacer algo mas complejo su analisis y entendimiento.

Si has llegado hasta aca, es muy posible que compartamos la misma preocupación y el programa que dejo a disposicion de la cominidad, viene de alguna forma a subsanar este problema, incorporando una capa de protección en ausencia de herramientas que permitan obfuscar tu código fuente creado para Golang.

Claro está que otros lenguajes y plataformas permiten minimizar el codigo para evitar lo mismo que tratamos acá pero, esto se trata de soluciones creadas para ser compiladas con Golang y distribuidas como binarios ejecutables.

Esta utilidad está disponible en este repositorio en su version fuente, legible y documentada para que la analices, la uses, la modifiques y la distribuyas sin ninguna restricción y espero realmente, si este fuese el caso, te sea de real utilidad.

Una vez que descargues o clones el repositorio, debes compilar el codigo para utilizarlo. A continuación te muestro como hacerlo si no tienes experiencia con Golang.  

Por favor, revisa el codigo antes de compilarlo y ejecutarlo. Si no estas seguro o no sientes confianza, te sugiero no utilizarlo hasta que comprendas bien lo que hace y como lo hace; Si comprendes el codigo fuente y te da confianza te reirero que, espero te sea de utilidad y que sea un real aporte para la proteccion de tu codigo y tu esfuerzo.

Requisitos:

- Disponer de una instalación de Golang en su ultima versión.

Compilación de la herramienta:




La ofuscación y compilación se realiza en dos fases:

1.- Se ofuscan la carpeta con los fuentes originales en una carpeta con los fuentes ofuscados que seran utilizados en la compilación y,

2.- La compilación de los fuentes desde la carpeta de archivos fuentes ofuscados; Para ello, se utilizará la libreria "garble" que permite ofuscar el binario generado.

Use:

go-cource-ofus ./source-code ./source-ofuscated
cd ./source-ofuscated

-- Build Normal
go build .

-- Build con ofuscación de binario. Require "go install mvdan.cc/garble@latest" 
garble build .


If you compile the offered source code found in the folder. " So, if your source code is included in the binary without your consent, the binary will have a much more difficult code to read and interpret.

100% guaranteed security does not exist but, this application leaves me a little quieter, I hope you too and I hope it serves you.


