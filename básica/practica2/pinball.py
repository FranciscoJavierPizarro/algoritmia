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
#           Los parametros 1 y 2 son nombres de ficheros, el 3º parametros es  #
#           un número entre 0 y 2 para especificar el algoritmo a emplear      #
#           0 - Haskell simulado                                               #
#           1 - Haskell directo                                                #
#           2 - Python simulado                                                #
#                                                                              #
################################################################################
import time
import subprocess
import sys
from binaryTree import *
CMD = "./simulation "

def ejecutarPrueba(prueba,algoritmo,debug):
    if len(prueba.split()) <= 1:
        raise ValueError
    [profundidad, n] = prueba.split()
    if type(int(profundidad)) is not int or type(int(n)) is not int or int(profundidad) < 1 or int(n) < 1:
        raise ValueError
    if algoritmo == 0:
        resultado = subprocess.run(CMD + str(profundidad) + " " + str(n) + " 0", stdout=subprocess.PIPE, shell=True, universal_newlines=True).stdout
    elif algoritmo == 1:
        resultado = subprocess.run(CMD + str(profundidad) + " " + str(n), stdout=subprocess.PIPE, shell=True, universal_newlines=True).stdout
    elif algoritmo == 2:
        start = time.time()
        resultado = str(lanzarSimulacion(int(profundidad),int(n))) +"\n"
        end = time.time()
        resultado = "{:.9f}".format((end - start) * 1000) + " " + resultado
    if debug:
        resultado = resultado.split()[1]
    return resultado

def ejecutar(args):
    algoritmo = 0
    if len(args) != 3 and len(args) != 2:
        raise TypeError
    if (len(args) == 3):
        algoritmo = int(args[2])
    # print(args)
    if (len(args) == 2 or len(args) == 3):
        ficheroPruebas = args[0]
        ficheroSoluciones = args[1]
        try:
            with open(ficheroPruebas) as fPruebas, open(ficheroSoluciones, "w") as fSoluciones: 
                for prueba in fPruebas:
                    resultado = ejecutarPrueba(prueba,algoritmo,False)
                    fSoluciones.write(str(resultado))
                    
        except FileNotFoundError:
            raise FileNotFoundError
        except PermissionError:
            raise PermissionError



if __name__ == "__main__":
    ejecutar(sys.argv[1:])
    