package main

import "fmt"
import (
    "bufio"
    "os"
	"time"
)

var ahorroPD int = 0

func buscarCadenaPD(dic map[string]bool, cadenaBuscada string, n int, encontrados []string, pseudomatrix [][]string, alreadyCalculated int) bool {
	if cadenaBuscada == "" {
		// fmt.Println(encontrados);
		 return true}
	if n > len(cadenaBuscada) {return false}
	if n == 1 && len(pseudomatrix[alreadyCalculated]) != 0 {
		for _, element := range pseudomatrix[alreadyCalculated] {
			buscarCadenaPD(dic, cadenaBuscada[len(element)-1:] , 1, append(encontrados,element),pseudomatrix,alreadyCalculated+(len(element)-1))
			ahorroPD += len(element)
			// fmt.Println("MATRIZ USADA")
		}
	}
	_, exists := dic[cadenaBuscada[:n]]
	encontradaGlobal := 0
	encontradaLocal := 0
	encontradosNuevos := encontrados
	if exists {
		encontradosNuevos = append(encontrados,cadenaBuscada[:n])
	}
	// esta asginación a entero en cuenta de hacerlo con booleanos directamente 
	// es porque en golang el || esta cortocircuitado y el | que usa enteros no.
	if (exists && buscarCadenaPD(dic, cadenaBuscada[n:], 1, encontradosNuevos, pseudomatrix, alreadyCalculated + 1)) {
		encontradaLocal = 1
		pseudomatrix[alreadyCalculated + 1 - n] = append(pseudomatrix[alreadyCalculated + 1 - n], cadenaBuscada[:n]) 
		
	}
	if (buscarCadenaPD(dic, cadenaBuscada, n+1, encontrados, pseudomatrix, alreadyCalculated + 1)) {encontradaGlobal = 1}
	return  ((encontradaLocal) | (encontradaGlobal)) == 1
}

func buscarCadenaBF(dic map[string]bool, cadenaBuscada string, n int, encontrados []string, alreadyCalculated int) bool {
	
	if cadenaBuscada == "" {
		// fmt.Println(encontrados); 
		return true}
	if n > len(cadenaBuscada) {return false}
	
	_, exists := dic[cadenaBuscada[:n]]
	encontradaGlobal := 0
	encontradaLocal := 0
	encontradosNuevos := encontrados
	if exists {
		encontradosNuevos = append(encontrados,cadenaBuscada[:n])
	}
	// esta asginación a entero en cuenta de hacerlo con booleanos directamente 
	// es porque en golang el || esta cortocircuitado y el | que usa enteros no.
	if (exists && buscarCadenaBF(dic, cadenaBuscada[n:], 1, encontradosNuevos,  alreadyCalculated + 1)) {
		encontradaLocal = 1}
	if (buscarCadenaBF(dic, cadenaBuscada, n+1, encontrados,  alreadyCalculated + 1)) {encontradaGlobal = 1}
	return  ((encontradaLocal) | (encontradaGlobal)) == 1
}


func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    cadenaObjetivo := scanner.Text()

    dic := make(map[string]bool)
    for scanner.Scan() {
        dic[scanner.Text()] = true
    }
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        os.Exit(1)
    }
	
	pseudomatrix := make([][]string, len(cadenaObjetivo) + 1)
	for i := range pseudomatrix {
		pseudomatrix[i] = make([]string, 0)
	}
	start := time.Now()
if buscarCadenaPD(dic, cadenaObjetivo, 1, make([]string, 0), pseudomatrix, 0) {
    fmt.Println("Sí.")
} else {
    fmt.Println("No.")
}
elapsed := time.Since(start)
fmt.Printf("Tiempo de ejecución de buscarCadenaPD: %s\n", elapsed)
fmt.Println(ahorroPD)
start = time.Now()
if buscarCadenaBF(dic, cadenaObjetivo, 1, make([]string, 0), 0) {
    fmt.Println("Sí.")
} else {
    fmt.Println("No.")
}
elapsed = time.Since(start)
fmt.Printf("Tiempo de ejecución de buscarCadenaBF: %s\n", elapsed)
}