#!/usr/bin/bash
# Compilamos el programa de go
go build main.go

echo "Ejecutando pruebas estandar"
./main input.txt 
python3 main.py input.txt &> /dev/null

for size in {10,25,50,75,100}; do
    echo "Ejecutando pruebas custom de tamaño ${size}"
    ./main tests/input${size}.txt ${size}
    python3 main.py tests/input${size}.txt ${size} &> /dev/null
    
    c1_f1=$(cat "outputLinear${size}.txt" | cut -d ' ' -f 1)
    c1_f2=$(cat "outputRamif${size}.txt" | cut -d ' ' -f 1)

    if [ "$c1_f1" == "$c1_f2" ]; then
        echo "La solución en ambos ficheros para tamaño ${size} es ${c1_f1}"
    else
        echo "Las soluciones de ambos algoritmos son dispares entre si siendo estas: ${c1_f1} y ${c1_f2}"
    fi
done