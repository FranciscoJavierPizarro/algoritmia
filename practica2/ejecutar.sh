#!/usr/bin/bash
# Este shell ejecuta el código encargado de testear cifrado.py

echo "Compilando el código haskell"
ghc -O2 simulation.hs 
mkdir ./tmp
mv ./*.{hi,o} ./tmp
echo "Lanzando test automatizados para comprobar el funcionamiento de cifrado.py"
echo "Empleando version de python:"
python3 --version
python3 testsPinball.py -v