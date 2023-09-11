package main

// go build main.go auxiliar_funcs.go  sorting_algoritms.goC
// go run *.go
import (
	"fmt"
	"time"
)

func main() {
	vectors := ReadVectorsFromFile("./random_arrays.txt")
	// Sample integer vector.
	// vector := IntVector{1, 2, 3, 4, 5}

	// Define an array of functors.
	functions := []IntVectorFunc{RadixSort, QuickSort, ConcurrentQuickSort, ConcurrentBogoSort, MergeSort, ConcurrentMergeSort, BubbleSort, HeapSort, CubeSort, TreeSort}
	for _, vector := range vectors {
		for _, function := range functions {
			// Measure execution time.
			start := time.Now()
			function(vector)
			duration := time.Since(start)

			// Print the result and execution time.
			fmt.Printf("Function: %s\n", FunctionName(function))
			fmt.Printf("Execution Time: %s\n", duration)
			fmt.Println()
		}
	}
}
