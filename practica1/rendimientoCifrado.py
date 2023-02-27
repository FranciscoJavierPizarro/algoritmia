################################################################################
#                                                                              #
#     Archivo: rendimientoCifrado.py                                           #
#     Fecha de última revisión: 27/02/2023                                     #
#     Autores: Francisco Javier Pizarro 821259                                 #
#              Jorge Solán Morote   	816259                                 #
#     Comms:                                                                   #
#           Este archivo contiene la medida de rendimiento del programa de     #
#           cifrado y descifrado de la práctica 1 de algoritmia básica.        #
#           El programa genera una configuración de datos iniciales válidos    #
#           y la emplea para cifrar y descifrar un mensaje de longitud 10^n    #
#           Siendo n un valor que cambia en cada ejecución de la función       #
#           El programa muestra por salida estandar el tamaño del mensaje      #
#           y el coste en ms del tiempo de ejecución de ese cifrado/descifrado #
#           Para ejecutar este  programa ejecutar python3 rendimientpCifrado.py#
#                                                                              #
################################################################################

from cifrado import *
import time

def performanceTest():
    
    datos = generarDatosInicio(8)
    #anclamos los datos de inicio para que se emplee la misma mochila y no afecte
    #a las metricas obtenidas
    for exponent in range(1,8):
        start = time.time()
        msj = "a" * (10 ** exponent)
        cifradoYDescifrado(msj,datos[0],datos[1],datos[2])
        end = time.time()
        print((10 ** exponent), " ", ((end - start) * 1000))# iteracion tiempoEnMs
    
if __name__ == "__main__":
    performanceTest()