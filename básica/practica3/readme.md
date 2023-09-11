# Práctica 3    Programación dinámica

En esta práctica se emplea la programación dinámica para analizar si una cadena es el resultado o no de juntar de cualquier forma un conjunto de palabras.

Este problema al tener un espacio de soluciones tan grande resulta extremadamente costoso de resolver empleando fuerza bruta por ello se emplea programación dinámica, esta almacena información ya calculada para no tener que recalcularla la próxima ves que esta sea necesaria.

Se han empleado 2 enfoques distintos de programación dinámica: el primero de ellos consiste en realizar la búsqueda en formato DFS e ir almacenando sobre la marcha los resultados obtenidos para no tener que recalcularlos cuando se den de nuevo, el segundo de ellos consiste en realizar el precálculo de todos los posibles resultados para que la solución simplemente tenga que mirar el mejor resultado obtenido. Realmente ambas soluciones son igual de buenas si las comparamos en su forma base dado que mas o menos tardan lo mismo y realizan el mismo número de cálculos, no obstante esto a la vez no es cierto dado que si mejoramos la solución, la precalculada es infinitamente mejor.

La gran ventaja de la solución precalculada respecto a la búsqueda es que esta primera es completamente **paralelizable** y dado que empleamos Golang como lenguaje, dicha parelización resulta casi inmediata.

El código principal se encuentra en main.go. Adicionalmente para hacerlo mas interesante se genera de forma aleatoria una entrada parametrizada con generarEntrada.py