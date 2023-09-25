package main

// go build main.go auxiliar_funcs.go  sorting_algoritms.go
// go run *.go
import (
	"fmt"
	"os"
	"time"
	"strings"
	"flag"
)

func main() {
	verb := flag.Bool("v", false, "Enable verbose mode")
	dataset := flag.String("dataset", "medio", "Dataset to use (simple, medioParcialmenteOrdenado ,medio, grande)")

	flag.Parse()
	verbose := *verb
	if (verbose) {
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
		functions = []IntVectorFunc{RadixSort, ConcurrentBogoSort }
	case "medio-small":
		input = "./datasets/mediumsmall.txt"
		functions = []IntVectorFunc{HeapSort, TreeSort, RadixSort, MergeSort, QuickSort, BubbleSort,PancakeSort}
	case "medioParcialmenteOrdenado":
		input = "./datasets/semiSortedmedium.txt"
		functions = []IntVectorFunc{HeapSort, TreeSort, RadixSort, MergeSort, QuickSort}
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

	outputFilePath := "medidas.txt"
	
	// Open the file for writing. Create it if it doesn't exist, truncate it if it does.
	file, _ := os.Create(outputFilePath)
	defer file.Close()
	
	vectors := ReadVectorsFromFile(input)
	// vectors := []IntVector{{43, 29, 51, 21, 74}}
	// functions := []IntVectorFunc{}

	header := "Size"
	for _, function := range functions {
		header += " " + strings.Split(FunctionName(function), ".")[1]
	}
	header += "\n"
	file.Write([]byte(header))
	timeMeasure := ""
	for _, vector := range vectors {
		timeMeasure = fmt.Sprintf("%d", len(vector))
		for _, function := range functions {
			// Measure execution time.
			if (verbose) {
				fmt.Printf("Function: %s\n", FunctionName(function))
				fmt.Println(len(vector))
			}

			start := time.Now()
			function(vector,verbose)
			duration := time.Since(start).Milliseconds()
			timeMeasure += " " + fmt.Sprintf("%d", duration)

			if (verbose) {
				fmt.Printf("Execution Time: %d\n", duration)
				fmt.Println()
			}
		}
		timeMeasure += "\n"
		file.Write([]byte(timeMeasure))
	}
}
