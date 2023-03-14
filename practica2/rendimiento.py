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
#           4.Analiza en un tiempo finito t=5s a que N y que P llega cada uno  #
#           de los distintos algoritmos.                                       #
#                                                                              #
################################################################################

from binaryTree import *
import time
import subprocess
def _performanceTest3Algs():
    N = 100
    with open("rendimientos/rendimientoPyth.txt", "w") as fpy,open("rendimientos/rendimientoHaskSimulado.txt", "w") as fhsS,open("rendimientos/rendimientoHaskDirecto.txt", "w") as fhsD: 
        for P in range(2,25):
            comando = "./simulation " + str(P) + " " + str(N)
            
            start = time.time()
            lanzarSimulacion(P,N)
            end = time.time()
            fpy.write(str (P) + " " + str((end - start) * 1000)+ "\n")# iteracion tiempoEnMs
            
            hstime = subprocess.run(comando + " 0", stdout=subprocess.PIPE, shell=True, text=True).stdout 
            fhsS.write(str (P) + " " + hstime+ "\n")# iteracion tiempoEnMs

            hstime = subprocess.run(comando, stdout=subprocess.PIPE, shell=True, text=True).stdout
            fhsD.write(str (P) + " " + hstime+ "\n")# iteracion tiempoEnMs

def _performanceTest2AlgsProf():
    N = 1000000
    with open("rendimientos/rendimientoProfHaskSimulado.txt", "w") as fhsS,open("rendimientos/rendimientoProfHaskDirecto.txt", "w") as fhsD: 
        for P in range(1,7):
            comando = "./simulation " + str(10**P) + " " + str(N)
            
            hstime = subprocess.run(comando + " 0", stdout=subprocess.PIPE, shell=True, text=True).stdout 
            fhsS.write("10^"+str (P) + " " + hstime+ "\n")# iteracion tiempoEnMs

            hstime = subprocess.run(comando, stdout=subprocess.PIPE, shell=True, text=True).stdout
            fhsD.write("10^"+str (P) + " " + hstime+ "\n")# iteracion tiempoEnMs

def _performanceTest2AlgsBalls():
    P = 1000000
    with open("rendimientos/rendimientoBallsHaskSimulado.txt", "w") as fhsS,open("rendimientos/rendimientoBallsHaskDirecto.txt", "w") as fhsD: 
        for N in range(1,8):
            comando = "./simulation " + str(P) + " " + str(10**N)
            
            hstime = subprocess.run(comando + " 0", stdout=subprocess.PIPE, shell=True, text=True).stdout 
            fhsS.write("10^"+str (N) + " " + hstime+ "\n")# iteracion tiempoEnMs

            hstime = subprocess.run(comando, stdout=subprocess.PIPE, shell=True, text=True).stdout
            fhsD.write("10^"+str (N) + " " + hstime+ "\n")# iteracion tiempoEnMs

if __name__ == "__main__":
    # _performanceTest3Algs()
    _performanceTest2AlgsProf()
    _performanceTest2AlgsBalls()