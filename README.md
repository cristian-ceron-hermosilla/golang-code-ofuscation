# golang-code-ofuscation
Golang code ofuscation

If you are encoding an application and want to extreme the measures to avoid that in one way or another "someone" steals your code, perhaps this may interest you.

This program is motivated by the impossibility of obfuscar your source code for Golang. While it is feasible to "obfuscar" the binary when you compile, no one can assure you 100% that the compiler will not include your source code "without obfuscar" in the binary.

Of course, the great corporations that develop compilers, languages and platforms could gather a team of engineers who build the same as you developed but, it cannot be denied that it would be of great help if they can look at your code and the "how" you resolved things They took a lot of analysis and development.

This code that I release in this repository can be used and modified at your entire freedom and is motivated in the difficulty I had to find some utility that obfuscara the variables, identifiers and functions of the Galang source code.

Use:

ofus-source ./source-code ./source-ofus
cd ./source-ofus

-- Normal build
go build .

-- Build ofuscated binary. Require "go install mvdan.cc/garble@latest" 
garble build .


If you compile the offered source code found in the folder. " So, if your source code is included in the binary without your consent, the binary will have a much more difficult code to read and interpret.

100% guaranteed security does not exist but, this application leaves me a little quieter, I hope you too and I hope it serves you.


