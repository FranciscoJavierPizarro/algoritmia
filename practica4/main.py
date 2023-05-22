from mip import Model, xsum, maximize, BINARY
import os
import sys
import time

def cargar_datos(filename):
    enunciados = []
    pedidosEnunciados = []

    with open(filename, 'r') as f:
        for linea in f:
            linea = linea.strip()
            if linea == '0 0 0':
                break

            enunciados.append(linea)
            datos = linea.split(' ')
            p = int(datos[2])
            pedidos = []
            for _ in range(p):
                pedido = f.readline().strip()
                pedidos.append(pedido)
            pedidosEnunciados.append(pedidos)

    return enunciados, pedidosEnunciados

def solve_problem(capacidad, m, p, pedidos):
    # Creación modelo, variables booleanas que representan sin un pedido se ha recogido o no
    # y variables enteras que representan el estado de la capacidad en cada una de las estaciones
    model = Model()
    y = [model.add_var(var_type=BINARY) for _ in range(p)]
    x = [model.add_var(lb=0) for _ in range(m + 1)]

    # Función a maximizar, en este caso el beneficio de los pedidos recogidos
    model.objective = maximize(xsum(pedidos[i][2] * (pedidos[i][1] - pedidos[i][0]) * y[i] for i in range(p)))

    # Restricciones
    # El valor de X en cada estación debe ser exactamente igual a la suma de los pasajeros de los pedidos recogidos que transiten en esa estación
    # (excluimos los que finalizan dado que los pasajeros bajan al llegar al destino)
    for j in range(m + 1):
        model.add_constr(x[j] == xsum(pedidos[i][2] * y[i] for i in range(p) if pedidos[i][0] <= j < pedidos[i][1]))
    # En ningún punto del recorrido el valor local de la X(capacidad) puede superar la capacidad absoluta máxima del tren
    for j in range(m + 1):
        model.add_constr(x[j] <= capacidad)

    # Ejecutamos el modelo midiendo el tiempo de ejecución
    start_time = time.time()
    model.optimize()
    end_time = time.time()

    # Extraemos resultados
    max_income = model.objective_value
    runtime = (end_time - start_time) * 1000

    # Para visualizar mas información
    # x_values = [v.x for v in x]
    # y_values = [v.x for v in y]
    # print(x_values,y_values)

    return max_income, runtime

def solve_problem_instance(enunciado, pedidosEnunciado):

    # Parseamos los datos leidos del fichero correspondientes a una instancia del problema
    capacidad = int(enunciado.split(' ')[0])
    m = int(enunciado.split(' ')[1])
    pedidos = []
    for pedido in pedidosEnunciado:
        data = pedido.split(' ')
        startpos = int(data[0])
        endpos = int(data[1])
        passenger = int(data[2])
        # ticketvalue = passenger * (endpos - startpos)
        pedidos.append((startpos, endpos, passenger))

    # Resolvemos el problema
    max_income, runtime = solve_problem(capacidad, m, len(pedidos), pedidos)
    
    # Para visualizar mas información
    # Print the solution
    # print(f"Max Income: {max_income}")
    # print(f"Runtime: {runtime} ms")
   
    return max_income, runtime


enunciados, pedidosEnunciados = cargar_datos(sys.argv[1])
with open("outputLinear.txt", 'w') as f:
    for i in range(len(enunciados)):
        max_income, runtime = solve_problem_instance(enunciados[i],pedidosEnunciados[i])
        f.write(f"{max_income} {runtime}\n")