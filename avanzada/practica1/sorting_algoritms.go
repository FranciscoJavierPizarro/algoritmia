////////////////////////////////////////////////////////////////////////////////
//                                                                            //
//     Archivo: sorting_algoritms.go                                          //
//     Fecha de última revisión: 08/10/2023                                   //
//     Autores: Francisco Javier Pizarro 821259                               //
//              Jorge Solán Morote   	816259                                //
//     Comms:                                                                 //
//           Este archivo contiene todos los algortimos de ordenación		  //
//			que probaremos en la práctica							  		  //
//                                                                            //
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

type IntVector []int

type IntVectorFunc func(IntVector, bool)

func RadixSort(ints IntVector, verbose bool) {
	max := getMax(ints)
	//Para cada cifra, realizar la ordenación con el countingSort
	for exp := 1; max/exp > 0; exp *= 10 {
		ints = countingSort(ints, exp)
	}

	if verbose {
		fmt.Println(ints)
	}
	return
}

func QuickSort(ints IntVector, verbose bool) {
	result := recQuickSort(ints)
	if verbose {
		fmt.Println(result)
	}
	return
}

func recQuickSort(ints IntVector) IntVector {
	if ints != nil && len(ints) > 1 {
		pivote := ints[0]
		menoresIguales, mayores := divideInLowersAndGreaters(pivote, ints[1:])
		return append(append([]int(recQuickSort(menoresIguales)), ints[:1]...), []int(recQuickSort(mayores))...)
	} else {
		if len(ints) == 1 {
			return ints
		} else {
			return nil
		}
	}
}

func ConcurrentQuickSort(ints IntVector, verbose bool) {
	retChan := make(chan IntVector)
	go concurrentRecQuickSort(ints, retChan, 4)
	resultado := <-retChan

	if verbose {
		fmt.Println(resultado)
	}
}

func concurrentRecQuickSort(ints IntVector, ret chan IntVector, w int) {
	N := len(ints)
	var a, b IntVector
	if N > 1 {
		pivote := ints[0]
		menoresIguales, mayores := divideInLowersAndGreaters(pivote, ints[1:])
		if w > 0 {
			lowerRes, higherRes := make(chan IntVector), make(chan IntVector)
			go concurrentRecQuickSort(menoresIguales, lowerRes, w-1)
			go concurrentRecQuickSort(mayores, higherRes, w-1)
			a = <-lowerRes
			b = <-higherRes
		} else {
			a = recQuickSort(menoresIguales)
			b = recQuickSort(mayores)
		}
		finalVec := append(append(a, pivote), b...)
		ret <- finalVec
	} else {
		ret <- ints
	}
}

func bogoSortInstance(seguir, encontrado chan bool, ints IntVector, res chan IntVector) {
	sigo := true
	rand.Seed(time.Now().UnixNano())

	//Mientras no se encuentre el resultado
	for sigo {
		// Reordena el vector
		newVec := shuffle(ints)
		//Mira si está ordenado
		sorted := isSorted(newVec)
		encontrado <- sorted
		sigo = <-seguir
		if sorted {
			res <- newVec
		}
	}
}

func ConcurrentBogoSort(ints IntVector, verbose bool) {
	keepSearching := true
	nWorkers := 10
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
	resultado := <-res
	if verbose {
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
		// Mientras se pueda dividir el vector
		//Primera mitad
		firstHalf := recMergeSort(ints[:N/2])
		//Segunda mitad
		secondHalf := recMergeSort(ints[N/2:])
		//Junta y ordena resultado
		mergedVec := merge(firstHalf, secondHalf)
		return mergedVec
	} else {
		return ints
	}
}

func MergeSort(ints IntVector, verbose bool) {
	resultado := recMergeSort(ints)
	if verbose {
		fmt.Println(resultado)
	}
}

func ConcurrentMergeSort(ints IntVector, verbose bool) {
	retChan := make(chan IntVector)
	//Lanza la función con 4 workers
	go concurrentRecMergeSort(ints, retChan, 4)
	resultado := <-retChan

	if verbose {
		fmt.Println(resultado)
	}
}

func concurrentRecMergeSort(ints IntVector, ret chan IntVector, w int) {
	N := len(ints)
	var a, b IntVector
	if N > 1 {
		if w > 0 {
			//Si hay workers
			resultados := make(chan IntVector)
			//Recorrer primera parte del vector
			go concurrentRecMergeSort(ints[:N/2], resultados, w-1)
			//Recorrer segunda parte del vector
			go concurrentRecMergeSort(ints[N/2:], resultados, w-1)
			a = <-resultados
			b = <-resultados
		} else {
			//Si no quedan workers
			a = recMergeSort(ints[:N/2])
			b = recMergeSort(ints[N/2:])
		}
		//Junta el resultado
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

	if verbose {
		fmt.Println(ints)
	}

	return
}

func HeapSort(ints IntVector, verbose bool) {
	h := &IntHeap{}
	//Inicia el heap
	heap.Init(h)
	heap.Push(h, 3)
	//Rellena el heap con el contenido del vector
	for _, v := range ints {
		heap.Push(h, v)
	}
	for h.Len() > 0 {
		if verbose {
			fmt.Printf("%d ", heap.Pop(h))
		} else {
			heap.Pop(h)
		}
	}
	return
}

func PancakeSort(ints IntVector, verbose bool) {
	N := len(ints)
	currSize := N
	for currSize > 1 {
		maxIndex := findMaxIndex(ints[:currSize])
		if maxIndex != (currSize - 1) {
			flip(ints, maxIndex)
			flip(ints, currSize-1)
		}
		currSize--
	}

	if verbose {
		fmt.Println(ints)
	}
	return
}

func TreeSort(ints IntVector, verbose bool) {
	var t Tree
	//Crea el árbol
	for _, v := range ints {
		t.insert(v)
	}
	if verbose {
		fmt.Print("[")
	}
	//Recorre desde raiz
	inOrder(t.root, verbose)
	if verbose {
		fmt.Print("]")
	}
	return
}
