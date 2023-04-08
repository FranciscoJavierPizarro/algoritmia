################################################################################
#                                                                              #
#     Archivo: testsPinball.py                                                 #
#     Fecha de última revisión: 08/04/2023                                     #
#     Autores: Francisco Javier Pizarro 821259                                 #
#              Jorge Solán Morote   	816259                                 #
#     Comms:                                                                   #
#           Este archivo contiene los test del programa pinball                #
#           de la práctica 2 de algoritmia básica. El programa cuenta con dos  #
#           clases que realizan diferentes tipos de test                       #
#           La primera realiza tests contra casos de uso que emplean valores   #
#           inválidos para algún tipo de dato(en los ficheros)                 #
#           El segundo realiza tests empleando parametros de ejecución         #
#           incorrectos                                                        #
#           Para ejecutar este  programa ejecutar python3 testsPinball.py -v   #
#                                                                              #
################################################################################

import unittest
from pinball import *
################################################################################
#                                                                              #
#    CASOS DE USO INCORRECTO POR VALORES INVÁLIDOS EN LOS FICHEROS             #
#                                                                              #
################################################################################    
class WrongValuesTest(unittest.TestCase):

    def test_valoresNegativos(self):
        try:
            self.assertNotEqual(ejecutarPrueba("-1 -2",0,True),"1","Error funciona con valores negativos")
        except ValueError:
            pass
        else:
            self.fail('Error funciona con valores negativos')
    
    def test_valoresNoNumericos(self):
        try:
            self.assertNotEqual(ejecutarPrueba("a Z",0,True),"1","Error funciona con valores no númericos")
        except ValueError:
            pass
        else:
            self.fail('Error funciona con valores no númericos')
    
    def test_valoresNoParseables(self):
        try:
            self.assertNotEqual(ejecutarPrueba("2",0,True),"1","Error funciona con valores no parseables")
        except ValueError:
            pass
        else:
            self.fail('Error funciona con valores no parseables')

    
    

################################################################################
#                                                                              #
#    CASOS DE USO INCORRECTO POR TIPOS DE VALOR INVÁLIDO                       #
#                                                                              #
################################################################################   
class WrongParametersTypesTest(unittest.TestCase):

    def test_numeroInadecuadoDeParametros(self):
        try:
            self.assertNotEqual(ejecutar([]),"1","Error funciona con un N de parámetros inadecuado")
        except TypeError:
            pass
        else:
            self.fail('Error funciona con un N de parámetros inadecuado')

    def test_ficherosInaccesibles(self):
        try:
            self.assertNotEqual(ejecutar(["tests/noexiste.txt","tests/noexiste.txt",1]),"1","Error funciona con un fichero que no existe")
        except FileNotFoundError:
            pass
        else:
            self.fail('Error funciona con un fichero que no existe')

    def test_ficherosIlegible(self):
        try:
            self.assertNotEqual(ejecutar(["tests/wrongPerm.txt","tests/wrongPerm.txt",1]),"1","Error funciona con un fichero que es ilegible")
        except PermissionError:
            pass
        else:
            self.fail('Error funciona con un fichero que es ilegible')
       
    
if __name__ == "__main__":
    unittest.main()