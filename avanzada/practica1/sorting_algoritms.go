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
type IntVectorFunc func(IntVector)

func RadixSort(ints IntVector) {
	max := getMax(ints)

	for exp := 1; max/exp > 0; exp *= 10 {
		ints = countingSort(ints, exp)
	}
	fmt.Println(ints)
	return
}

func QuickSort(ints IntVector) {
	fmt.Println(auxQuickSort(ints))
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

func ConcurrentQuickSort(ints IntVector) {
	result := 1
	for _, v := range ints {
		result *= v
	}
	return
}

func bogoSortInstance(seguir,encontrado chan bool, ints IntVector) {
	sigo := true
	rand.Seed(time.Now().UnixNano())

	for sigo {
		encontrado <- isSorted(shuffle(ints))
		sigo = <- seguir
	}
}

func ConcurrentBogoSort(ints IntVector) {
	keepSearching := true
	nWorkers := len(ints)
	seguir := make(chan bool)
	encontrado := make(chan bool)
	for I:= 0; I < nWorkers; I++ {
		go bogoSortInstance(seguir,encontrado,ints)
	}

	for keepSearching {
		keepSearching = !<- encontrado
		seguir <- keepSearching
	}

	nRestantes := nWorkers - 1
	for nRestantes > 0 {
		<- encontrado
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

func MergeSort(ints IntVector) {
	fmt.Println(recMergeSort(ints))
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
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, 3)
	for _, v := range ints {
		heap.Push(h, v)
	}
	// fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
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
	var root *Node

	root = insert(root, ints[0])
	for i := 1; i < len(ints); i++ {
		root = insert(root, ints[i])
	}

	i := 0
	storeSorted(root, ints, &i)
	fmt.Println(ints)
	return
}
