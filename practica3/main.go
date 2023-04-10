package main

import "fmt"
import (
    "bufio"
    "os"
	"time"
)

func searchWords(cadenaBuscada string, dic map[string]bool, precalcMatrix [][]string) {
    ch := make(chan []string)
    for i := 0; i < len(cadenaBuscada); i++ {
        go func(start int) {
            var encontradas []string
            for j := start; j < len(cadenaBuscada); j++ {
                prefix := cadenaBuscada[start : j+1]
                if dic[prefix] {
                    encontradas = append(encontradas, prefix)
                }
            }
            // Send the results to the channel
            ch <- encontradas
        }(i)
    }

    // Collect the results from the channel and store them in the matrix
    for i := 0; i < len(cadenaBuscada); i++ {
        precalcMatrix[i] = <-ch
    }
}

func printCombinations(cadenaBuscada string, precalcMatrix [][]string) bool {
    var totalCombinations int = 1
    for i := 0; i < len(precalcMatrix); i++ {
        totalCombinations *= len(precalcMatrix[i])
    }
    hasCombination := false
    for i := 0; i < totalCombinations; i++ {
        var combination string
        var index int = i
        for j := 0; j < len(precalcMatrix); j++ {
            wordIndex := index % len(precalcMatrix[j])
            combination += precalcMatrix[j][wordIndex] + " "
            index /= len(precalcMatrix[j])
        }
        combination = combination[:len(combination)-1]
        if combination == cadenaBuscada {
			fmt.Println(combination)
            hasCombination = true
        }
    }

    return hasCombination
}

func buscarCadenaPD(dic map[string]bool, cadenaBuscada string, n int, encontrados []string, pseudomatrix [][]string, alreadyCalculated int) bool {
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
	// esta asginación a entero en cuenta de hacerlo con booleanos directamente 
	// es porque en golang el || esta cortocircuitado y el | que usa enteros no.
	if (exists && buscarCadenaPD(dic, cadenaBuscada[n:], 1, encontradosNuevos, pseudomatrix, alreadyCalculated + 1)) {
		encontradaLocal = 1
		pseudomatrix[alreadyCalculated + 1 - n] = append(pseudomatrix[alreadyCalculated + 1 - n], cadenaBuscada[:n]) 
		
	}
	if (buscarCadenaPD(dic, cadenaBuscada, n+1, encontrados, pseudomatrix, alreadyCalculated + 1)) {encontradaGlobal = 1}
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
	precalcMatrix := make([][]string, len(cadenaObjetivo))
    searchWords(cadenaObjetivo, dic, precalcMatrix)
	if printCombinations(cadenaObjetivo,precalcMatrix) {
		fmt.Print("Sí. ")
	} else {
		fmt.Print("No. ")
	}
	elapsed := time.Since(start)
	fmt.Printf("Tiempo de ejecución de buscarCadenaPDPrecalc: %s\n", elapsed)

	start = time.Now()
	if buscarCadenaPD(dic, cadenaObjetivo, 1, make([]string, 0), pseudomatrix, 0) {
		fmt.Print("Sí. ")
	} else {
		fmt.Print("No. ")
	}
	elapsed = time.Since(start)
	fmt.Printf("Tiempo de ejecución de buscarCadenaPD: %s\n", elapsed)
}