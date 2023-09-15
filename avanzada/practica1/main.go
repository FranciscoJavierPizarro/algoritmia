package main

// go build main.go auxiliar_funcs.go  sorting_algoritms.go
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
	// functions := []IntVectorFunc{RadixSort, QuickSort, ConcurrentQuickSort, ConcurrentBogoSort, MergeSort, ConcurrentMergeSort, BubbleSort, HeapSort, CubeSort, TreeSort}

	// vectors := []IntVector{{43, 29, 51, 21, 74}}
	functions := []IntVectorFunc{TreeSort, RadixSort}
	for _, vector := range vectors {
		for _, function := range functions {
			// Measure execution time.
			fmt.Printf("Function: %s\n", FunctionName(function))
			start := time.Now()
			function(vector)
			duration := time.Since(start)

			// Print the result and execution time.
			fmt.Printf("Execution Time: %s\n", duration)
			fmt.Println()
		}
	}
}
