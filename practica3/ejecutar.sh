#!/usr/bin/bash
# Este shell ejecuta el código encargado de lanzar main.go

echo "Compilando el código golang empleando $(go version)"
go build main.go 
mkdir ./tmp
echo "Para generar los ficheros de entrada se emplea $(python3 --version)"
#10^5 son demasiadas palabras para el lector de ficheros de golang
for i in {1..3}; do
    n=$((10**i))
    echo "Ejecutando para ${n} palabras"
    python3 generarEntrada.py ${n}
    ./main ./tmp/f${n}.txt
    ./main ./tmp/fMod${n}.txt
done