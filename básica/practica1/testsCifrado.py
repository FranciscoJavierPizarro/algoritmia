################################################################################
#                                                                              #
#     Archivo: testsCifrado.py                                                 #
#     Fecha de última revisión: 27/02/2023                                     #
#     Autores: Francisco Javier Pizarro 821259                                 #
#              Jorge Solán Morote   	816259                                 #
#     Comms:                                                                   #
#           Este archivo contiene los test del programa de cifrado y descifrado#
#           de la práctica 1 de algoritmia básica. El programa cuenta con tres #
#           clases que realizan diferentes tipos de test                       #
#           La primera realiza tests de casos de uso correctos y normales      #
#           La segunda realiza tests contra casos de uso que emplean valores   #
#           inválidos para algún tipo de dato(siendo los tipos correctos)      #
#           El tercero realiza tests empleando tipos de datos incorrectos      #
#           Para ejecutar este  programa ejecutar python3 testsCifrado.py -v   #
#                                                                              #
################################################################################

from cifrado import *
import unittest


################################################################################
#                                                                              #
#    CASOS DE USO CORRECTO Y NORMAL                                            #
#                                                                              #
################################################################################
class TestCifradoYDescifrado(unittest.TestCase):

    def test_mensajeSencilloDatosValidos(self):
        msj = "mensajeSimple"
        datos = generarDatosInicio(8)
        self.assertEqual(cifradoYDescifrado(msj,datos[0],datos[1],datos[2]),msj,"Error descifrando el mensaje simple")

    def test_mensajeConEspaciosDatosValidos(self):
        msj = "mensaje con espacios"
        datos = generarDatosInicio(8)
        self.assertEqual(cifradoYDescifrado(msj,datos[0],datos[1],datos[2]),msj,"Error descifrando el mensaje con espacios")

    def test_mensajeConCaracteresDatosValidos(self):
        msj = "mensaje con .``'\ $ "
        datos = generarDatosInicio(8)
        self.assertEqual(cifradoYDescifrado(msj,datos[0],datos[1],datos[2]),msj,"Error descifrando el mensaje con carácteres especiales")

################################################################################
#                                                                              #
#    CASOS DE USO INCORRECTO POR VALORES INVÁLIDOS                             #
#                                                                              #
################################################################################    
class WrongParametersValuesTest(unittest.TestCase):

    def test_mochilaNoOptima(self):
        msj = "mensaje"
        try:
            self.assertNotEqual(cifradoYDescifrado(msj,[8, 87, 555, 731, 4234, 39685, 2], 1479689, 1422497),msj,"Error funciona con una mochila inválida")
        except ValueError:
            pass
        else:
            self.fail('Error mochila inválida')

    def test_mochilaPequeña(self):
        msj = "mensaje"
        try:
            self.assertNotEqual(cifradoYDescifrado(msj,[1,2,3],7,5),msj,"Error funciona con una mochila muy pequeña")
        except ValueError:
            pass
        else:
            self.fail('Error mochila pequeña')

    def test_wNegativa(self):
        msj = "mensaje"
        try:
            self.assertNotEqual(cifradoYDescifrado(msj,[8, 87, 555, 731, 4234, 39685, 370240], 1479689, -1),msj,"Error funciona con una w negativa")
        except ValueError:
            pass
        else:
            self.fail('Error w negativa')
    
    def test_wMayorQueN(self):
        msj = "mensaje"
        try:
            self.assertNotEqual(cifradoYDescifrado(msj,[8, 87, 555, 731, 4234, 39685, 370240], 1422497, 1479689),msj,"Error funciona con una w mayor que N")
        except ValueError:
            pass
        else:
            self.fail('Error w mayor que N')
    
    def test_wNoPrimoConN(self):
        msj = "mensaje"
        try:
            self.assertNotEqual(cifradoYDescifrado(msj,[8, 87, 555, 731, 4234, 39685, 370240], 1479689, 1479689),msj,"Error funciona con una w que no es prima respecto a N")
        except ValueError:
            pass
        else:
            self.fail('Error w no prima respecto a N')

    def test_NnoEsElMayor(self):
        msj = "mensaje"
        try:
            self.assertNotEqual(cifradoYDescifrado(msj,[8, 87, 555, 731, 4234, 39685, 370240], 2, 1479689),msj,"Error funciona con una N que no es mas grande que la mochila")
        except ValueError:
            pass
        else:
            self.fail('Error N menor que la mochila')

################################################################################
#                                                                              #
#    CASOS DE USO INCORRECTO POR TIPOS DE VALOR INVÁLIDO                       #
#                                                                              #
################################################################################   
class WrongParametersTypesTest(unittest.TestCase):

    def test_mensajeNoString(self):
        msj = "mensaje"
        try:
            self.assertNotEqual(cifradoYDescifrado(0,[8, 87, 555, 731, 4234, 39685, 370240], 1479689, 1422497),msj,"Error funciona con un mensaje que no es un string")
        except ValueError:
            pass
        else:
            self.fail('Error funciona con un mensaje que no es un string')
    
    def test_mochilaNoLista(self):
        msj = "mensaje"
        try:
            self.assertNotEqual(cifradoYDescifrado(msj,'a', 1479689, 1422497),msj,"Error funciona con una mochila que no es una lista")
        except ValueError:
            pass
        else:
            self.fail('Error funciona con una mochila que no es una lista')
        
    def test_wNoInt(self):
        msj = "mensaje"
        try:
            self.assertNotEqual(cifradoYDescifrado(msj,[8, 87, 555, 731, 4234, 39685, 370240],'a',1422497),msj,"Error funciona con un w que no es un int")
        except ValueError:
            pass
        else:
            self.fail('Error funciona con un w que no es un int')

    def test_NNoInt(self):
        msj = "mensaje"
        try:
            self.assertNotEqual(cifradoYDescifrado(msj,[8, 87, 555, 731, 4234, 39685, 370240], 1479689,'a'),msj,"Error funciona con un N que no es un int")
        except ValueError:
            pass
        else:
            self.fail('Error funciona con un N que no es un int')
       
    
if __name__ == "__main__":
    unittest.main()