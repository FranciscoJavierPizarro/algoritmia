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
#     Use:  Llamar a la función lanzarSimulacion(P,N)                          #
#           Siendo N el número de bolas y P la profundidad del árbol           #
#                                                                              #
################################################################################


################################################################################
#                                                                              #
#    CLASE ÁRBOL BINARIO                                                       #
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

def _generateTree(P):
    """Devuelve un árbol binario de profundidad P
    
    Parameters:
    w(int):Profundidad deseada
        
    Returns:
    Nodo:Raíz del árbol binario generado
    """   
    raiz = Node()
    P -= 1
    pendientesDeGenerarSubnodos = []
    
    if P > 0:
        raiz.insertLeft()
        raiz.insertRight()
        pendientesDeGenerarSubnodos.append(raiz.left)
        pendientesDeGenerarSubnodos.append(raiz.right)
        P -= 1
    while P > 0:
        siguientesSubnodos = []
        for nodo in pendientesDeGenerarSubnodos:
            nodo.insertLeft()
            nodo.insertRight()
            siguientesSubnodos.append(nodo.left)
            siguientesSubnodos.append(nodo.right)
        P -= 1
        pendientesDeGenerarSubnodos = siguientesSubnodos
    
    return raiz


################################################################################
#                                                                              #
#    FUNCIONES DE SIMULACIÓN                                                   #
#                                                                              #
################################################################################

def _simularBola(raiz):
    """Dada la raíz de un árbol binario simula el recorrido de una bola a traves del mismo, modificando el árbol a su paso
    
    Parameters:
    raiz(Node):Raíz del árbol
    
    Returns:
    int:Valor del nodo por el que sale la bola
    """   
    nodo = raiz
    while nodo.left != None or nodo.right != None:
        nodo.data = not nodo.data
        if nodo.data:
            nodo = nodo.left
        else:
            nodo = nodo.right
    return nodo.nodeId

def _simularNBolas(raiz,n):
    """Dada la raíz de un árbol binario simula el recorrido de N bolas a traves del mismo, modificando el árbol a su paso
    
    Parameters:
    raiz(Node):Raíz del árbol
    n(int):Número de bolas a simular
    
    Returns:
    int:Valor del nodo por el que sale la última bola
    """ 
    # valorUltimaHojaRecorrida = 1
    for i in range(n):
        valorUltimaHojaRecorrida = _simularBola(raiz)

    return valorUltimaHojaRecorrida

def lanzarSimulacion(P,N):
    """Dada la profundidad de un árbol binario, lo crea y simula el recorrido de N bolas a traves del mismo, modificando el árbol a su paso
    
    Parameters:
    P(int):Profundidad del árbol
    N(int):Número de bolas a simular
    
    Returns:
    int:Valor del nodo por el que sale la última bola
    """ 
    return _simularNBolas(_generateTree(P),N)

if __name__ == "__main__":
    lanzarSimulacion(25,1)