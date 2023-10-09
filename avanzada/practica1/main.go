////////////////////////////////////////////////////////////////////////////////
//                                                                            //
//     Archivo: main.go                                                       //
//     Fecha de última revisión: 08/10/2023                                   //
//     Autores: Francisco Javier Pizarro 821259                               //
//              Jorge Solán Morote   	816259                                //
//     Comms:                                                                 //
//           Este archivo contiene el core de la práctica 1 de algoritmia     //
//           avanzada											  			  //
//     Use:  																  //
//			go build main.go auxiliar_funcs.go  sorting_algoritms.go		  //
// 			go run *.go									  					  //
//			Lanzar el script ejecutar.sh                                      //
//                                                                            //
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Parseo de flags de ejecucción
	verb := flag.Bool("v", false, "Enable verbose mode")
	dataset := flag.String("dataset", "medio", "Dataset to use (simple, medioParcialmenteOrdenado ,medio, grande)")

	flag.Parse()
	verbose := *verb
	if verbose {
		fmt.Println("Vectores cargados.")
	}
	input := ""
	functions := []IntVectorFunc{}
	switch *dataset {
	case "real":
		input = "./datasets/real.tsv"
		functions = []IntVectorFunc{HeapSort, TreeSort, RadixSort, MergeSort, QuickSort, ConcurrentMergeSort, ConcurrentQuickSort}
	case "simple":
		input = "./datasets/small.txt"
		functions = []IntVectorFunc{RadixSort, ConcurrentBogoSort}
	case "medio-small":
		input = "./datasets/mediumsmall.txt"
		functions = []IntVectorFunc{HeapSort, TreeSort, RadixSort, MergeSort, QuickSort, BubbleSort, PancakeSort}
	case "medio-big":
		input = "./datasets/mediumbig.txt"
		functions = []IntVectorFunc{HeapSort, TreeSort, RadixSort, MergeSort, QuickSort}
	case "big":
		input = "./datasets/big.txt"
		functions = []IntVectorFunc{RadixSort, ConcurrentMergeSort, ConcurrentQuickSort}
	default:
		input = "./datasets/mediumbig.txt"
		functions = []IntVectorFunc{HeapSort, TreeSort, RadixSort, MergeSort, QuickSort}
	}

	
	// Gestión del fichero de salida
	outputFilePath := "medidas.txt"
	file, _ := os.Create(outputFilePath)
	defer file.Close()
	// Carga de los vectores de entrada
	vectors := ReadVectorsFromFile(input)

	// Escritura de la cabezera del fichero de salida
	header := "Size"
	for _, function := range functions {
		header += " " + strings.Split(FunctionName(function), ".")[1]
	}
	header += "\n"
	file.Write([]byte(header))
	timeMeasure := ""

	// Para cada vector del dataset, se ordena empleando cada una de las funciones asociadas al dataset
	// se mide el tiempo de cada función y se escribe en el fichero de salida.
	for _, vector := range vectors {
		timeMeasure = fmt.Sprintf("%d", len(vector))
		for _, function := range functions {
			// Measure execution time.
			if verbose {
				fmt.Printf("Function: %s\n", FunctionName(function))
				fmt.Println(len(vector))
			}

			start := time.Now()
			function(vector, verbose)
			duration := time.Since(start).Milliseconds()
			timeMeasure += " " + fmt.Sprintf("%d", duration)

			if verbose {
				fmt.Printf("Execution Time: %d\n", duration)
				fmt.Println()
			}
		}
		timeMeasure += "\n"
		file.Write([]byte(timeMeasure))
	}
}
