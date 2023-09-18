//sorting_algoritms.go

package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

type IntVector []int

// Define a functor for a function that takes an IntVector and returns an int.
type IntVectorFunc func(IntVector, bool)

func RadixSort(ints IntVector, verbose bool) {
	max := getMax(ints)

	for exp := 1; max/exp > 0; exp *= 10 {
		ints = countingSort(ints, exp)
	}
	
	if (verbose) {
		fmt.Println(ints)
	}
	return
}

func QuickSort(ints IntVector, verbose bool) {
	result := auxQuickSort(ints)
	if (verbose) {
		fmt.Println(result)
	}
	return
}

func auxQuickSort(ints IntVector) IntVector {
	if ints != nil && len(ints) > 1 {
		pivote := ints[0]
		menoresIguales := filterLowerOrEqualThan(pivote, ints[1:])
		// fmt.Println(menoresIguales)
		mayores := filterGreaterThan(pivote, ints[1:])
		// fmt.Println(mayores)
		return append(append([]int(auxQuickSort(menoresIguales)), ints[:1]...), []int(auxQuickSort(mayores))...)
		// return concatMultipleSlices([]IntVector {auxQuickSort(menoresIguales), ints[:1], auxQuickSort(mayores)})
	} else {
		if len(ints) == 1 {
			return ints
		} else {
			return nil
		}
	}
}

func ConcurrentQuickSort(ints IntVector, verbose bool) {
	result := 1
	for _, v := range ints {
		result *= v
	}
	return
}

func bogoSortInstance(seguir, encontrado chan bool, ints IntVector, res chan IntVector) {
	sigo := true
	rand.Seed(time.Now().UnixNano())

	for sigo {
		newVec := shuffle(ints)
		sorted := isSorted(newVec)
		encontrado <- sorted
		sigo = <-seguir
		if (sorted) {
			res <- newVec
		}
	}
}

func ConcurrentBogoSort(ints IntVector, verbose bool) {
	keepSearching := true
	nWorkers := len(ints)
	seguir := make(chan bool)
	encontrado := make(chan bool)
	res := make(chan IntVector)
	for I := 0; I < nWorkers; I++ {
		go bogoSortInstance(seguir, encontrado, ints, res)
	}

	for keepSearching {
		keepSearching = !<-encontrado
		seguir <- keepSearching
	}
	resultado := <- res
	if (verbose) {
		fmt.Println(resultado)
	}
	nRestantes := nWorkers - 1
	for nRestantes > 0 {
		<-encontrado
		seguir <- keepSearching
		nRestantes--
	}
	return
}

func recMergeSort(ints IntVector) IntVector {
	N := len(ints)
	if N > 1 {
		firstHalf := recMergeSort(ints[:N/2])
		secondHalf := recMergeSort(ints[N/2:])
		mergedVec := merge(firstHalf, secondHalf)
		return mergedVec
	} else {
		return ints
	}
}

func MergeSort(ints IntVector, verbose bool) {
	resultado := recMergeSort(ints)
	if (verbose) {
		fmt.Println(resultado)
	}
}

func ConcurrentMergeSort(ints IntVector, verbose bool) {
	retChan := make(chan IntVector)
	go concurrentRecMergeSort(ints,retChan)
	resultado := <- retChan

	if (verbose) {
		fmt.Println(resultado)
	}
}

func concurrentRecMergeSort(ints IntVector, ret chan IntVector) {
	N := len(ints)
	if N > 1 {
		resultados := make(chan IntVector)
		go concurrentRecMergeSort(ints[:N/2], resultados)
		go concurrentRecMergeSort(ints[N/2:], resultados)
		a := <- resultados
		b := <- resultados
		mergedVec := merge(a, b)
		ret <- mergedVec
	} else {
		ret <- ints
	}
}

func BubbleSort(ints IntVector, verbose bool) {
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
	
	if (verbose) {
		fmt.Println(ints)
	}

	return
}

func HeapSort(ints IntVector, verbose bool) {
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, 3)
	for _, v := range ints {
		heap.Push(h, v)
	}
	// fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		if (verbose) {
			fmt.Printf("%d ", heap.Pop(h))
		} else {
			heap.Pop(h)
		}
	}
	return
}

func CubeSort(ints IntVector, verbose bool) {
	result := 1
	for _, v := range ints {
		result *= v
	}
	return
}

func TreeSort(ints IntVector, verbose bool) {
	var t Tree
	for _, v := range ints {
		t.insert(v)
	}
	if (verbose) {
		fmt.Print("[")
		printPostOrder(t.root)
		fmt.Print("]")
	} else {
		//añadir versión sin prints
	}
	return
}
