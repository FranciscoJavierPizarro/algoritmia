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
	// Define a boolean flag for verbose mode
	verb := flag.Bool("v", false, "Enable verbose mode")

	// Parse command line arguments
	flag.Parse()
	verbose := *verb
	if (verbose) {
		fmt.Println("Vectores cargados.")
	}
	// Sample integer vector.
	// vector := IntVector{1, 2, 3, 4, 5}
	
	// Define an array of functors.
	// functions := []IntVectorFunc{RadixSort, QuickSort, ConcurrentQuickSort, ConcurrentBogoSort, MergeSort, ConcurrentMergeSort, BubbleSort, HeapSort, CubeSort, TreeSort}
	
	filePath := "medidas.txt"
	
	// Open the file for writing. Create it if it doesn't exist, truncate it if it does.
	file, _ := os.Create(filePath)
	defer file.Close()
	
	vectors := ReadVectorsFromFile("./random_arrays.txt")
	// vectors := []IntVector{{43, 29, 51, 21, 74}}
	functions := []IntVectorFunc{HeapSort, TreeSort, RadixSort, MergeSort, QuickSort, ConcurrentMergeSort}
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
