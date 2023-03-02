#!/usr/bin/python3
import sys
from binaryTree import *
if __name__ == "__main__":
    # print(sys.argv)
    if (len(sys.argv) == 3):
        ficheroPruebas = sys.argv[1]
        ficheroSoluciones = sys.argv[2]
        with open(ficheroPruebas) as fPruebas, open(ficheroSoluciones, "w") as fSoluciones: 
            for prueba in fPruebas:
                [profundidad, n] = prueba.split()
                raiz = generateTree(int(profundidad))
                fSoluciones.write(str(simularNBolas(raiz,int(n))))


    