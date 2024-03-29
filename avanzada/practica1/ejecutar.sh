#!/usr/bin/bash

echo "generando datasets aleatorios"
/usr/bin/ruby ./generator.rb

echo "datasets generados"
echo "generando ejecutable del programa principal"

/usr/local/go/bin/go build main.go auxiliar_funcs.go  sorting_algoritms.go

flags=("real" "medio-small" "medio-big")

for flag in "${flags[@]}"
do
  echo "Ejecutando dataset: $flag"
  ./main -dataset ${flag}
  echo "Tiempos de ejecuccion del dataset:"
  cat ./medidas.txt
done