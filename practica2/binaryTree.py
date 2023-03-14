################################################################################
#                                                                              #
#     Archivo: binaryTree.py                                                   #
#     Fecha de última revisión: 14/03/2023                                     #
#     Autores: Francisco Javier Pizarro 821259                                 #
#              Jorge Solán Morote   	816259                                 #
#     Comms:                                                                   #
#           Este archivo contiene el TAD nodo sobre el cual implementa el      #
#           árbol binario de la práctica 2 de algoritmia básica.               #
#           El módulo contiene también las funciones necesarias para crear un  #
#           árbol binario de profundidad P, así como el código necesario para  #
#           simular el lanzamiento de N bolas a lo largo de un arbol dado      #
#                                                                              #
################################################################################



maxNode = 1
class Node:
    def __init__(self):
      global maxNode
      self.left = None
      self.right = None
      self.data = False
      self.nodeId = maxNode
      maxNode += 1

    def insertLeft(self):
        self.left = Node()
        
    def insertRight(self):
        self.right = Node()

def _generateTree(profundidad):
    raiz = Node()
    profundidad -= 1
    pendientesDeGenerarSubnodos = []
    
    if profundidad > 0:
        raiz.insertLeft()
        raiz.insertRight()
        pendientesDeGenerarSubnodos.append(raiz.left)
        pendientesDeGenerarSubnodos.append(raiz.right)
        profundidad -= 1
    while profundidad > 0:
        siguientesSubnodos = []
        for nodo in pendientesDeGenerarSubnodos:
            nodo.insertLeft()
            nodo.insertRight()
            siguientesSubnodos.append(nodo.left)
            siguientesSubnodos.append(nodo.right)
        profundidad -= 1
        pendientesDeGenerarSubnodos = siguientesSubnodos
    
    return raiz

def _simularBola(raiz):
    nodo = raiz
    while nodo.left != None or nodo.right != None:
        nodo.data = not nodo.data
        if nodo.data:
            nodo = nodo.left
        else:
            nodo = nodo.right
    return nodo.nodeId

def _simularNBolas(raiz,n):
    for i in range(n):
        valorUltimaHojaRecorrida = _simularBola(raiz)
    return valorUltimaHojaRecorrida

def lanzarSimulacion(P,N):
    return _simularNBolas(_generateTree(P),N)

if __name__ == "__main__":
    lanzarSimulacion(7,1)