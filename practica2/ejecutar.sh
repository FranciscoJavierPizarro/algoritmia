#!/usr/bin/bash
# Este shell ejecuta el código encargado de testear pinball.py

echo "Compilando el código haskell"
ghc -O2 simulation.hs 
mkdir ./tmp
mkdir ./outputs
mv ./*.{hi,o} ./tmp
echo "Lanzando test automatizados para comprobar el funcionamiento de pinball.py"
echo "Empleando version de python: $(python3 --version)"

./test.sh
# python3 testsPinball.py -v