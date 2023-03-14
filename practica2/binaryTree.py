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