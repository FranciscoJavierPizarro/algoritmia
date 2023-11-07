import random

def generar_cuadrado_latino(n):
    cuadrado = [['*' for _ in range(n)] for _ in range(n)]

    for i in range(n):
        numeros_disponibles = set(range(1, n + 1))

        for j in range(n):
            numero = random.choice(list(numeros_disponibles))
            cuadrado[i][j] = str(numero)
            numeros_disponibles.remove(numero)

    for i in range(n):
        for j in range(n):
            numero = cuadrado[j][i]
            cuadrado[j][i] = cuadrado[i][j]
            cuadrado[i][j] = numero

    return cuadrado

def cambiar_a_asteriscos(cuadrado, num_asteriscos):
    for _ in range(num_asteriscos):
        i = random.randint(0, len(cuadrado) - 1)
        j = random.randint(0, len(cuadrado) - 1)
        cuadrado[i][j] = '*'

def escribir_cuadrado_en_archivo(cuadrado, nombre_archivo):
    with open(nombre_archivo, 'w') as archivo:
        for fila in cuadrado:
            archivo.write(' '.join(fila) + '\n')

if __name__ == "__main__":
    n = 9
    cuadrado_latino = generar_cuadrado_latino(n)
    
    num_asteriscos = int(input("Ingrese la cantidad de asteriscos deseada: "))
    cambiar_a_asteriscos(cuadrado_latino, num_asteriscos)
    
    nombre_archivo = "output.txt"
    
    escribir_cuadrado_en_archivo(cuadrado_latino, nombre_archivo)
    print(f"Cuadrado latino generado y guardado en '{nombre_archivo}'.")
