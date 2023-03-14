#!/usr/bin/bash
# Este shell borra los archivos generados al compilar los fuentes de Haskell o al ejecutar Python
rm -r ./tmp
rm ./*{hi,o} simulation
rm ./rendimientos/rendimiento*.txt
rm -r __pycache__