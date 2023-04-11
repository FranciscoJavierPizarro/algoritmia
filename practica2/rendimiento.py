################################################################################
#                                                                              #
#     Archivo: rendimiento.py                                                  #
#     Fecha de última revisión: 14/03/2023                                     #
#     Autores: Francisco Javier Pizarro 821259                                 #
#              Jorge Solán Morote   	816259                                 #
#     Comms:                                                                   #
#           Este archivo contiene la medida de rendimiento del programa de     #
#           simulación de la práctica 2 de algoritmia básica.                  #
#           El programa realiza varios test de rendimiento                     #
#           1.Compara coste temporal de los 3 algoritmos para profundiades de  #
#           [1-25] empleando 100 bolas                                         #
#           2.Compara coste temporal de los algoritmos haskell para N =1000000 #
#           y profundidades € [1-10000000]                                     #
#           3.Compara coste temporal de los algoritmos haskell para P =1000000 #
#           y BOLAS € [1-100000000]                                            #
#                                                                              #
################################################################################

from binaryTree import *
import time
import subprocess

################################################################################
#                                                                              #
#    FUNCIONES PARA MEDIR RENDIMIENTOS                                         #
#                                                                              #
################################################################################

def _performanceTest3Algs():
    """Compara el coste temporal para N = 100 y P € [1-25] para los tres algoritmos""" 
    N = 20
    with open("rendimientos/rendimientoPyth.txt", "w") as fpy,open("rendimientos/rendimientoHaskSimulado.txt", "w") as fhsS,open("rendimientos/rendimientoHaskDirecto.txt", "w") as fhsD: 
        for P in range(2,15):
            comando = "./simulation " + str(P) + " " + str(N)
            
            start = time.time()
            lanzarSimulacion(P,N)
            end = time.time()
            fpy.write(str (P) + " " + str((end - start) * 1000)+ "\n")# iteracion tiempoEnMs
            
            hstime = subprocess.run(comando + " 0", stdout=subprocess.PIPE, shell=True, text=True).stdout.split()[0] 
            fhsS.write(str (P) + " " + hstime+ "\n")# iteracion tiempoEnMs

            hstime = subprocess.run(comando, stdout=subprocess.PIPE, shell=True, text=True).stdout.split()[0]
            fhsD.write(str (P) + " " + hstime+ "\n")# iteracion tiempoEnMs

def _performanceTest2AlgsProf():
    """Compara el coste temporal para N = 1000000 y P € [10^1-10^7] para los algoritmos hechos en Haskell""" 
    N = 10000
    with open("rendimientos/rendimientoProfHaskSimulado.txt", "w") as fhsS,open("rendimientos/rendimientoProfHaskDirecto.txt", "w") as fhsD: 
        for P in range(1,6):
            comando = "./simulation " + str(10**P) + " " + str(N)
            
            hstime = subprocess.run(comando + " 0", stdout=subprocess.PIPE, shell=True, text=True).stdout.split()[0] 
            fhsS.write("10^"+str (P) + " " + hstime+ "\n")# iteracion tiempoEnMs

            hstime = subprocess.run(comando, stdout=subprocess.PIPE, shell=True, text=True).stdout.split()[0]
            fhsD.write("10^"+str (P) + " " + hstime+ "\n")# iteracion tiempoEnMs

def _performanceTest2AlgsBalls():
    """Compara el coste temporal para P = 1000000 y N € [10^1-10^8] para los algoritmos hechos en Haskell""" 
    P = 100000
    with open("rendimientos/rendimientoBallsHaskSimulado.txt", "w") as fhsS,open("rendimientos/rendimientoBallsHaskDirecto.txt", "w") as fhsD: 
        for N in range(1,6):
            comando = "./simulation " + str(P) + " " + str(10**N)
            
            hstime = subprocess.run(comando + " 0", stdout=subprocess.PIPE, shell=True, text=True).stdout.split()[0] 
            fhsS.write("10^"+str (N) + " " + hstime+ "\n")# iteracion tiempoEnMs

            hstime = subprocess.run(comando, stdout=subprocess.PIPE, shell=True, text=True).stdout.split()[0]
            fhsD.write("10^"+str (N) + " " + hstime+ "\n")# iteracion tiempoEnMs

if __name__ == "__main__":
    print("Ejecutando métricas rendimiento de los 3 algoritmos para N = 100 y P € [1-25]")
    _performanceTest3Algs()
    print("Ejecutando métricas rendimiento para N = 1000000 y P € [10^1-10^7] para los algoritmos hechos en Haskell")
    _performanceTest2AlgsProf()
    print("Ejecutando métricas rendimiento para P = 1000000 y N € [10^1-10^8] para los algoritmos hechos en Haskell")
    _performanceTest2AlgsBalls()