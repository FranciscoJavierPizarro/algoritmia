#!/usr/bin/bash
# Este shell lanza los test de correcto funcionamiento de los distintos algoritmos
echo "Lanzando test con fichero pyth.txt"
for number in {0..2}
do
  ./pinball tests/pyth.txt outputs/pyth${number}.txt ${number}
done
echo "Lanzando test con fichero hs.txt"
for number in {0,1}
do
  ./pinball tests/hs.txt outputs/hs${number}.txt ${number}
done