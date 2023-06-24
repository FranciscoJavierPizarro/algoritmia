# Práctica 2    Divide y vencerás

En esta práctica se emplea un de divide y vencerás para simular el movimiento de unas pelotas a través de un arbol binario de profundidad P.

En esta práctica pasamos de un rendimiento de O(2^P) a uno de O(P) y en última instancia **O(1)**. 

Concretamente en esta se propusieron 3 algoritmos distintos para resolver el problema:

El primero de ellos fue crear el arbol entero y simular pelota a pelota el recorrido, su código esta en binaryTree.py

El segundo de ellos fue con divide y venceras creando solo la parte del arbol a recorrer y simulando solo la última pelota, su código esta en simulation.hs

El tercero de ellos fue un razonamiento matemático puro para resolver el problema sin necesidad de simular la última pelota ni de crear algún nodo de árbol, su código esta en simulation.hs

El código principal se encuentra en pinball.py, los archivos rendimiento.py y testsPinball.py son para evaluar el rendimiento y la correcta ejecucción del código anterior. Adicionalmente en el fichero BinaryTree.hs se define el TAD árbol binario en Haskell.