################################################################################
#                                                                              #
#     Archivo: cifrado.py                                                      #
#     Fecha de última revisión: 27/02/2023                                     #
#     Autores: Francisco Javier Pizarro 821259                                 #
#              Jorge Solán Morote   	816259                                 #
#     Comms:                                                                   #
#           Este archivo contiene el core del programa de cifrado y descifrado #
#           de la práctica 1 de algoritmia básica. El programa cuenta con dos  #
#           funciones interfaz empleables desde el exterior                    #
#           GenerarDatosInicio(n)                                              #
#           cifradoYDescifrado(mensaje, mochila,N,w)                           #
#           El resto son funciones internas para la comprobación de datos y    #
#           el uso de cifrado y descifrado                                     #
#           Se incluye un ejemplo de ejecución en caso de ejecutar este        #
#           programa con python3 cifrado.py al final del fichero               #
#                                                                              #
################################################################################

import random

################################################################################
#                                                                              #
#    FUNCIONES MATEMÁTICAS AUXILIARES                                          #
#                                                                              #
################################################################################

def _egcd(w: int, N: int) -> tuple:
    "Algoritmo extendido de Euclides"  
    if w == 0:
        return (N, 0, 1)
    else:
        g, y, x = _egcd(N % w, w)
        return (g, x - (N // w) * y, y)

def _calcularInverso(w: int, N: int) -> int:
    """Calcula el inverso de 'w' mod 'N'
    
    Parameters:
    w(int):Valor intermedio primo respecto a N
    N(int):Valor máximo calculado previamente
    
    Returns:
    int:El inverso de los dos valores previos
    """      
    g, x, y = _egcd(w, N)
    return x % N

def _mcd(a: int,b: int) -> int:
    """Devuelve el mínimo común múltiplo de dos números
    
    Parameters:
    a(int)
    b(int)

    Returns:
    int:Su valor es el mcd de los valores introducidos
    """
    if a > b:
        pequeño = b
    else:
        pequeño = a
    for i in range(1, pequeño+1):
        if((a % i == 0) and (b % i == 0)):
            mcd = i
    return mcd


################################################################################
#                                                                              #
#    COMPROBACIONES PREVIAS                                                    #
#                                                                              #
################################################################################
def _checkMochila(mochila):
    if type(mochila) is not list:
        raise ValueError
    if len(mochila) < 7:#caso especial en el que no hay bits suficientes para cifrar/descifrar la tabla ascii
        raise ValueError
    resul = 0
    for M in mochila:
        if( M > resul ):
            resul+=M
        else:
            raise ValueError
    return

def _checkN(mochila,N):
    resul = 0
    if type(N) is not int:
        raise ValueError
    for M in mochila:
        resul += M
    if(N <= resul):
        raise ValueError

def _checkW(w,N):
    if type(w) is not int:
        raise ValueError
    if((w < 1) or (_mcd(w,N) != 1)  or (w >= N)):
        raise ValueError

def _checkMsj(msj):
    if type(msj) is not str:
        raise ValueError

################################################################################
#                                                                              #
#    GENERAR DATOS INICIO VÁLIDOS                                              #
#                                                                              #
################################################################################

def generarDatosInicio(n: int) -> list:
    """Genera los parámetros empleados a lo largo del programa para ahorrar al usuario el tener que introducirlos a mano

    Parameters:
    n(int): Número de elementos deseados en la mochila

    Returns: [
        mochila,
        N,
        w,
        w-1
    ]
    """

    mochila = []
    i = 0
    maxValue = 0
    while i < n:
        auxNum = random.randint(maxValue + 1, (maxValue+1)*10)
        mochila.append(auxNum)
        maxValue += auxNum
        i += 1
    N = random.randint(maxValue + 1, (maxValue+1)*10)

    posibleW = random.randint(0,N)
    while _mcd(posibleW,N) != 1:
        posibleW = random.randint(0,N)
    wInv = _calcularInverso(posibleW,N)
    return [mochila,N,posibleW,wInv]

def _clavePublica(mochila: list,N: int,w: int) -> list:
    """Calcula la clave pública de la mochila
    
    Pameters:
    mochila(list(int)):Lista de enteros
    N(int):N máximo
    w(int):valor intermedio primo respecto a N
    
    Returns:
    list(int):Tupla que es la clave pública de la mochila introducida
    """    
    tupla = []
    for e in mochila:
        tupla.append((w*e)%N)
    return tupla


################################################################################
#                                                                              #
#    FUNCIONES CONVERSIONES BINARIOS                                           #
#                                                                              #
################################################################################

def _msjToBin(mensaje: str,n: int) -> list:
    """
    A partir de un string de caracteres devuelve su array de binarios correspondiente

    Parameters:
    mensaje(str)
    n(int)

    Returns:
    list:Equivalente en binario al mensaje introducido, cada elemento de la lista es una letra
    """   
    binary = []
    getBinary = lambda x, n: format(x, 'b').zfill(n) # Devuelve un binario de 'x' con 'n' elementos
    for letra in mensaje:
        binary.append(getBinary(ord(letra),n))
    return binary

def _binToMsj(binMsj: str) -> str:
    """
    A partir de un string de binarios devuelve su string de caracteres

    Parameters:
    binMsj(str)
    
    Returns:
    str:Mensaje obtenido al traducir de binario a texto plano
    """
    MensajeFinal = ""
    for binValue in binMsj:
        MensajeFinal += chr(int(binValue,2))
    return MensajeFinal

def _intToBin(number: int, mochila: list) -> str:
    """
    Función auxiliar de descifrar donde devuelve un binario a partir de un 
    entero obtenido de la mochila original

    Parameters:
    number(int)
    mochila(list)

    Returns:
    str:Su valor es el resultado de descifrar el numero a binario empleando la mochila
    """

    binValue = ''
    for valorMochila in reversed(mochila):
        if (number >= valorMochila):
            number -= valorMochila
            binValue = '1' + binValue
        else:
            binValue = '0' + binValue
    return binValue


def _calculateC(binario: str,kpub: list) -> int:
    """
    Función auxiliar de cifrar donde a partir de un número binario lo transforma 
    en natural según kpub

    Parameters:
    binario(str)
    kpub(list)

    Returns:
    int:Su valor es el número natural obtenido de aplicar kpub sobre el binario
    """
    C = 0
    for i, bit in enumerate(binario):
        C += int(bit) * kpub[i] 
    return C

################################################################################
#                                                                              #
#    CORE DEL PROGRAMA                                                         #
#                                                                              #
################################################################################

def _cifrar(mensajeEnBinario: list,kpub: list) -> list:
    """
    Cifra 'mensajeEnBinario'

    Parameters:
    mensajeEnBinario(list)
    kpub(list)

    Returns:
    list:Su valor es el resultado de cifrar cada elemento en binario empleando kpub, es una lista de enteros
    """
    mensajeCifrado = []
    for binario in mensajeEnBinario:
        mensajeCifrado.append(_calculateC(binario,kpub))
    return mensajeCifrado
    
def _descifrar(mensajeCifrado: list, inverso: int, N: int, mochila: list) -> list:
    """
    Descifra 'mensajeCifrado' con 'w' inverso mod 'N'

    Parameters:
    mensajeCifrado(list)
    inverso(int)
    N(int)
    mochila(list)

    Returns:
    list:Su valor es una lista de binarios del mensaje descifrado
    """
    binariosResueltos = []
    for M in mensajeCifrado:
        binariosResueltos.append(_intToBin((M * inverso)%N,mochila))
    
    return binariosResueltos

def cifradoYDescifrado(mensaje: str, mochila: list,N: int,w: int) -> str:
    """
    Dado 'mensaje', 'mochila', 'N' y 'w' te recrea el procedimiento completo de, 
    a partir de una mochila fácil dada(mochila), codificarte el mensaje dado y,
    posteriormente, descodificarlo con sus respectivas claves públicas y privadas
    """  
    #checks previos de los valores de los datos de entrada
    _checkMsj(mensaje)
    _checkMochila(mochila)
    _checkN(mochila,N)
    _checkW(w,N)
    n = len(mochila)
    
    #calculos previos
    kpub = _clavePublica(mochila,N,w)
    kpriv = [w,N]
    inverso = _calcularInverso(w,N)

    #ciframos el mensaje en binario
    mensajeEnBinario = _msjToBin(mensaje,n)
    mensajeCifrado = _cifrar(mensajeEnBinario,kpub)
    
    #desciframos el binario a mensaje
    binarioDescifrado = _descifrar(mensajeCifrado,inverso,kpriv[1],mochila)
    mensajeDecodificado = _binToMsj(binarioDescifrado)
    return mensajeDecodificado

if __name__ == "__main__":
    msj = "mensajeSimple"
    datos = generarDatosInicio(7)
    print(datos)
    # print(cifradoYDescifrado(msj,datos[0],datos[1],datos[2]))
