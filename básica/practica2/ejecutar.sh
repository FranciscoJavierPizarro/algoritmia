#!/usr/bin/bash
# Este shell ejecuta el código encargado de testear pinball.py

echo "Compilando el código haskell"
ghc -O2 simulation.hs 
mkdir ./tmp
mkdir ./outputs
mv ./*.{hi,o} ./tmp
echo "Empleando version de python: $(python3 --version)"
echo "Lanzando tests automatizados para comprobar el funcionamiento de pinball"
./test.sh
echo "Lanzando tests de rendimiento"
python3 rendimiento.py
echo "Lanzando tests de excepciones"
python3 testsPinball.py -v