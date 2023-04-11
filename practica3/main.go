////////////////////////////////////////////////////////////////////////////////
//                                                                            //
//     Archivo: main.go                                                       //
//     Fecha de última revisión: 11/05/2023                                   //
//     Autores: Francisco Javier Pizarro 821259                               //
//              Jorge Solán Morote   	816259                                //
//     Comms:                                                                 //
//           Este archivo contiene el core de la práctica 3 de algoritmia     //
//           básica, se encarga de usar programación dinámica para determinar //
//           si una frase es el resultado de al menos una combinación de      //
//           palabras de un diccionario dado.                                 //
//     Use:  Lanzar el script ejecutar.sh                                     //
//                                                                            //
////////////////////////////////////////////////////////////////////////////////

package main

////////////////////////////////////////////////////////////////////////////////
//                                                                            //
//    IMPORTS Y PARÁMETROS                                                    //
//                                                                            //
////////////////////////////////////////////////////////////////////////////////

import "fmt"
import (
    "bufio"
    "os"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
//                                                                            //
//    FUNCIONES DE BÚSQUEDA CON PROGRAMACIÓN DINÁMICA                         //
//                                                                            //
////////////////////////////////////////////////////////////////////////////////

func searchWords(cadenaBuscada string, dic map[string]bool, precalcMatrix [][]string) {
    // Función de precalculo de la matriz
    // Para cada letra de la frase original lanza una gorutina, dicha gorutina busca
    // todas las palabras existentes en el diccionario que comiencen a partir de dicha letra
    // Devuelven los resultados por un canal y se almacena para el indice de dicha letra en la matriz

    // Crear e iniciar canales
    chs := make([]chan []string, len(cadenaBuscada))
    for i := range chs {
        chs[i] = make(chan []string)
    }

    // Crear y lanzar gorutinas
    for i := 0; i < len(cadenaBuscada); i++ {
        go func(start int,ch chan<- []string) {
            var encontradas []string
            for j := start; j < len(cadenaBuscada); j++ {
                prefix := cadenaBuscada[start : j+1]
                if dic[prefix] {
                    encontradas = append(encontradas, prefix)
                }
            }
            ch <- encontradas
        }(i, chs[i])
    }

    // Guardar resultados
    for i, ch := range chs {
        precalcMatrix[i] = <-ch
    }
}

func buscarPDPrecalc(cadenaBuscada string, precalcMatrix [][]string,alreadyCalculated int, combination string) bool {
    // Empleando la matriz precalculada previamente busca empezando por el principio
    // de la matriz todas las combinaciones de las palabras existentes en la matriz de forma recursiva
    // Si llega al final de forma exitosa devuelve true

    if cadenaBuscada == "" {
        return true}
    hasCombination := false
    for _, element := range precalcMatrix[alreadyCalculated] {
        hasCombination = hasCombination || buscarPDPrecalc(cadenaBuscada[len(element):], precalcMatrix ,alreadyCalculated+(len(element)),combination + " " + element)
    }
    return hasCombination
}

func buscarCadenaPD(dic map[string]bool, cadenaBuscada string, n int, encontrados []string, pseudomatrix [][]string, alreadyCalculated int) bool {
	// Realiza una busqueda del tipo fuerza bruta de forma descendente, en el proceso
    // cada palabra encontrada es almacenada al vuelo en la pseudomatriz para agilizar
    // las siguientes recursiones que la necesiten
    // En caso de que la frase sea una combinación de palabras del diccionario devuelve true
    
    if cadenaBuscada == "" {
		// fmt.Println(encontrados);
		 return true}
	if n > len(cadenaBuscada) {return false}
	if n == 1 && len(pseudomatrix[alreadyCalculated]) != 0 {
		for _, element := range pseudomatrix[alreadyCalculated] {
			buscarCadenaPD(dic, cadenaBuscada[len(element)-1:] , 1, append(encontrados,element),pseudomatrix,alreadyCalculated+(len(element)-1))
		}
	}
	_, exists := dic[cadenaBuscada[:n]]
	encontradaGlobal := 0
	encontradaLocal := 0
	encontradosNuevos := encontrados
	if exists {
		encontradosNuevos = append(encontrados,cadenaBuscada[:n])
	}
	// esta asignación a entero en cuenta de hacerlo con booleanos directamente 
	// es porque en golang el || esta cortocircuitado y el | que usa enteros no.
	if (exists && buscarCadenaPD(dic, cadenaBuscada[n:], 1, encontradosNuevos, pseudomatrix, alreadyCalculated + 1)) {
		encontradaLocal = 1
		pseudomatrix[alreadyCalculated + 1 - n] = append(pseudomatrix[alreadyCalculated + 1 - n], cadenaBuscada[:n]) 
		
	}
	if (buscarCadenaPD(dic, cadenaBuscada, n+1, encontrados, pseudomatrix, alreadyCalculated + 1)) {encontradaGlobal = 1}
	return  ((encontradaLocal) | (encontradaGlobal)) == 1
}


////////////////////////////////////////////////////////////////////////////////
//                                                                            //
//    CORE DEL PROGRAMA                                                       //
//                                                                            //
////////////////////////////////////////////////////////////////////////////////

func main() {
    //Apertura del fichero de datos
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()

    //Lectura de la frase del fichero
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    cadenaObjetivo := scanner.Text()

    //Lectura del diccionario de palabras
    dic := make(map[string]bool)
    for scanner.Scan() {
        dic[scanner.Text()] = true
    }
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        os.Exit(1)
    }
	
    //Generación de la pseudomatriz empleada en el algoritmo de al vuelo
	pseudomatrix := make([][]string, len(cadenaObjetivo) + 1)
	for i := range pseudomatrix {
		pseudomatrix[i] = make([]string, 0)
	}

    //Elección del algoritmo a emplear y ejecución cronometrada del mismo
    if ("0" != os.Args[2]) {
        // fmt.Fprintln(os.Stderr, "Algoritmo de precalc")
        start := time.Now()
        precalcMatrix := make([][]string, len(cadenaObjetivo))
        searchWords(cadenaObjetivo, dic, precalcMatrix)
        if buscarPDPrecalc(cadenaObjetivo,precalcMatrix,0,"") {
            fmt.Print("Sí. ")
        } else {
            fmt.Print("No. ")
        }
        elapsed := time.Since(start)
        fmt.Printf("%d\n", elapsed.Microseconds())
    } else {
        // fmt.Fprintln(os.Stderr, "Algoritmo de onfly")
        start := time.Now()
        if buscarCadenaPD(dic, cadenaObjetivo, 1, make([]string, 0), pseudomatrix, 0) {
            fmt.Print("Sí. ")
        } else {
            fmt.Print("No. ")
        }
        elapsed := time.Since(start)
        fmt.Printf("%d\n", elapsed.Microseconds())
    }
}