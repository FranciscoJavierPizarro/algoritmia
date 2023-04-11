#!/usr/bin/bash
# Este shell ejecuta el código encargado de lanzar main.go

echo "Compilando el código golang empleando $(go version)"
go build main.go 
mkdir ./tmp
rm ./rendimientos/*.txt
echo "Para generar los ficheros de entrada se emplea $(python3 --version)"
#10^5 son demasiadas palabras para el lector de ficheros de golang
for i in {1..3}; do
    n=$((10**i))
    echo "Ejecutando para ${n} palabras con los ficheros sin modificaciones"
    python3 generarEntrada.py ${n}
    echo "${n} $(./main ./tmp/f${n}.txt 1 | cut -d ' ' -f 2)" >> ./rendimientos/precalc.txt
    echo "${n} $(./main ./tmp/f${n}.txt 0 | cut -d ' ' -f 2)" >> ./rendimientos/pdonfly.txt
done
# n=2000
# echo "Ejecutando para ${n} palabras con los ficheros sin modificaciones"
# python3 generarEntrada.py ${n}
# echo "${n} $(./main ./tmp/f${n}.txt 1 | cut -d ' ' -f 2)" >> ./rendimientos/precalc.txt
# echo "${n} $(./main ./tmp/f${n}.txt 0 | cut -d ' ' -f 2)" >> ./rendimientos/pdonfly.txt
for i in {1..3}; do
    n=$((10**i))
    echo "Ejecutando para ${n} palabras con los ficheros con modificaciones"
    echo "${n} $(./main ./tmp/fMod${n}.txt 1 | cut -d ' ' -f 2)" >> ./rendimientos/precalcMod.txt
    echo "${n} $(./main ./tmp/fMod${n}.txt 0 | cut -d ' ' -f 2)" >> ./rendimientos/pdonflyMod.txt
done
# n=2000
# echo "Ejecutando para ${n} palabras con los ficheros con modificaciones"
# echo "${n} $(./main ./tmp/fMod${n}.txt 1 | cut -d ' ' -f 2)" >> ./rendimientos/precalcMod.txt
# echo "${n} $(./main ./tmp/fMod${n}.txt 0 | cut -d ' ' -f 2)" >> ./rendimientos/pdonflyMod.txt