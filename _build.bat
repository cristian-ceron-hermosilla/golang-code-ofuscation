@echo off

set current_dir=%cd%
echo Chequeando si es la primera ejecución.
if not exist %USERPROFILE%\reservedWords.txt (
    if exist reservedWords.txt (
        echo Es la primera ejecución.
        echo Movemos la lista de Palabras Reservadas a la carpeta del usuario.
        move reservedWords.txt %USERPROFILE%\reservedWords.txt
    ) else (
        echo El archivo reservedWords.txt no se encuentra en el directorio actual.
        exit /b 1
    )
) else (
    echo No es la primera ejecución.
)

echo Generando binarios ejecutables...

:: Verificar si la carpeta execs existe, si no, crearla
if not exist "execs" (
    echo Creamos una carpeta para los binarios generados.
    mkdir execs
    echo Para setear la carpeta de binarios en el PATH necesitamos permisos administrativos.
    echo export PATH="$PATH:%cd%\execs" >> %USERPROFILE%\.bashrc
)

:: Compilar source-ofus y mover el binario si la compilación es exitosa
cd source-ofus
go build
if %errorlevel% equ 0 (
    echo go-source-ofus compilado y movido a execs.
    move go-source-ofus ..\execs\.
) else (
    echo Error al compilar go-source-ofus.
    exit /b 1
)
cd ..

:: Compilar source-dic y mover el binario si la compilación es exitosa
cd source-dic
go build
if %errorlevel% equ 0 (
    echo go-source-dic compilado y movido a execs.
    move go-source-dic ..\execs\.
) else (
    echo Error al compilar go-source-dic.
    exit /b 1
)
cd ..

echo Binarios generados y movidos a carpeta [./exec].