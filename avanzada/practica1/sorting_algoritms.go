//sorting_algoritms.go

package main

import "fmt"

type IntVector []int

// Define a functor for a function that takes an IntVector and returns an int.
type IntVectorFunc func(IntVector)

func RadixSort(ints IntVector) {
	result := 0
	for _, v := range ints {
		result += v
	}
	return
}

func QuickSort(ints IntVector) {
	if len(ints) > 1 {
		// pivote := ints[0]

	}

	result := 1
	for _, v := range ints {
		result *= v
	}
	return
}

func ConcurrentQuickSort(ints IntVector) {
	result := 1
	for _, v := range ints {
		result *= v
	}
	return
}

func ConcurrentBogoSort(ints IntVector) {
	result := 1
	for _, v := range ints {
		result *= v
	}
	return
}

func MergeSort(ints IntVector) {
	result := 1
	for _, v := range ints {
		result *= v
	}
	return
}

func ConcurrentMergeSort(ints IntVector) {
	result := 1
	for _, v := range ints {
		result *= v
	}
	return
}

func BubbleSort(ints IntVector) {
	N := len(ints)
	aux := 0
	for i := range ints {
		for j := i; j < N; j++ {
			if ints[i] > ints[j] {
				aux = ints[i]
				ints[i] = ints[j]
				ints[j] = aux
			}
		}
	}
	fmt.Println(ints)
	return
}

func HeapSort(ints IntVector) {
	result := 1
	for _, v := range ints {
		result *= v
	}
	return
}

func CubeSort(ints IntVector) {
	result := 1
	for _, v := range ints {
		result *= v
	}
	return
}

func TreeSort(ints IntVector) {
	result := 1
	for _, v := range ints {
		result *= v
	}
	return
}
