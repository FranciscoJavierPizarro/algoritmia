#!/usr/bin/bash
# Este shell borra los archivos generados al compilar los fuentes de Haskell

rm ./*{hi,o} simulation
rm ./rendimiento*.txt
rm -r __pycache__