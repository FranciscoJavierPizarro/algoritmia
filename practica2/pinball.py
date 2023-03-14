#!/usr/bin/python3

################################################################################
#                                                                              #
#     Archivo: pinball.py                                                      #
#     Fecha de última revisión: 14/03/2023                                     #
#     Autores: Francisco Javier Pizarro 821259                                 #
#              Jorge Solán Morote   	816259                                 #
#     Comms:                                                                   #
#           Este archivo es el core del programa de simulación de la           #
#           práctica 2 de algoritmia básica.                                   #
#                                                                              #
################################################################################


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
                fSoluciones.write(str(lanzarSimulacion(int(profundidad),int(n))))


    