////////////////////////////////////////////////////////////////////////////////
//                                                                            //
//     Archivo: auxiliar_funcs.go                                             //
//     Fecha de última revisión: 08/10/2023                                   //
//     Autores: Francisco Javier Pizarro 821259                               //
//              Jorge Solán Morote   	816259                                //
//     Comms:                                                                 //
//           Este archivo contiene funciones auxiliares de la práctica 1 	  //
//			 de algoritmia avanzada								  			  //
//																			  //
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// Funcion auxiliar para obtener el nombre de una funcion
func FunctionName(f IntVectorFunc) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// Funcion auxiliar para cargar varios vectores desde un fichero de texto dado su nombre
func ReadVectorsFromFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	var vectors [][]int
	scanner := bufio.NewScanner(file)
	const maxBufferSize = 1024 * 1024 * 1024 * 2 // 1 GB buffer size
	buf := make([]byte, maxBufferSize)
	scanner.Buffer(buf, maxBufferSize)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		var vector []int

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil
			}
			vector = append(vector, num)
		}

		vectors = append(vectors, vector)
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return vectors
}

// Funcion auxiliar que dado un valor y un vector devuelve dos vectores cada uno contiene los menoresIguales y los mayores respectivamente
func divideInLowersAndGreaters(value int, vec []int) ([]int, []int) {
	var lowerEqual, greater []int
	for _, element := range vec {
		if element <= value {
			lowerEqual = append(lowerEqual, element)
		} else {
			greater = append(greater, element)
		}
	}
	return lowerEqual, greater
}

// Función de ordenación Counting Sort para números de 1 cifra
func countingSort(ints IntVector, exp int) IntVector {
	n := len(ints)
	output := make([]int, n)
	count := make([]int, 10)

	//Contamos la frecuencia de cada dígito en el vector
	for i := 0; i < n; i++ {
		digit := (ints[i] / exp) % 10
		count[digit]++
	}

	//Rellenamos el vector de conteo
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	//Realizamos la ordenación del vector
	for i := n - 1; i >= 0; i-- {
		digit := (ints[i] / exp) % 10
		output[count[digit]-1] = ints[i]
		count[digit]--
	}

	// Copiar el vector ordenado de vuelta al vector original.
	return output
}

// Funcion auxiliar para obtener el valor máximo de un vector
func getMax(ints IntVector) int {
	max := ints[0]
	// Encontrar el máximo valor del vector de Ints
	for _, num := range ints {
		if num > max {
			max = num
		}
	}
	return max
}

// Funcion auxiliar para mezclar dos vectores ordenados
func merge(A, B IntVector) IntVector {
	NA, NB := len(A), len(B)
	i, j := 0, 0
	mergedVec := IntVector{}

	for i < NA && j < NB {
		if A[i] < B[j] {
			mergedVec = append(mergedVec, A[i])
			i++
		} else {
			mergedVec = append(mergedVec, B[j])
			j++
		}
	}
	if i < NA {
		mergedVec = append(mergedVec, A[i:]...)
	} else if j < NB {
		mergedVec = append(mergedVec, B[j:]...)
	}
	return mergedVec
}

// Funcion auxiliar para comprobar si un vector esta ordenado
func isSorted(A IntVector) bool {
	N := len(A)
	for i := 1; i < N; i++ {
		if A[i-1] < A[i] {
			return false
		}
	}
	return true
}

// Funcion auxiliar para ordenar de forma aleatoria un vector
func shuffle(ints IntVector) IntVector {
	n := len(ints)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		ints[i], ints[j] = ints[j], ints[i]
	}
	return ints
}

// Funcion auxiliar para obtener el indice del maximo elemento de un vector
func findMaxIndex(ints []int) int {
	maxIndex := 0
	maxValue := ints[0]

	for i, v := range ints {
		if v > maxValue {
			maxIndex = i
			maxValue = v
		}
	}

	return maxIndex
}

// Funcion auxiliar para "flipear" los elementos de un vector ->
// 0 1 2 3 5 --> 5 3 2 1 0
func flip(ints []int, i int) {
	start := 0
	for start < i {
		temp := ints[start]
		ints[start] = ints[i]
		ints[i] = temp
		start++
		i--
	}
}

// //////////////////////////////////////////////////////////////////////////////
//
//	HEAP														  			   //
//
// https://pkg.go.dev/container/heap										   //
// //////////////////////////////////////////////////////////////////////////////

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push y Pop usan punteros porque modifican la longitud del slice,
	// no solo su contenido.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// //////////////////////////////////////////////////////////////////////////////
//
//	TREE														  			   //
//
//                                    										   //
// //////////////////////////////////////////////////////////////////////////////

type Tree struct {
	root *Node
}

type Node struct {
	key   int
	left  *Node
	right *Node
}

// Tree
func (t *Tree) insert(data int) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

// Node
func (n *Node) insert(data int) {
	//Si el nuevo valor es menor que el nodo donde está
	if data <= n.key {
		//Hijo izquierdo
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		//Hijo derecho
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func inOrder(n *Node, verbose bool) {
	//Recorre en Inorder el árbol
	if n == nil {
		return
	} else {
		inOrder(n.left, verbose)
		if verbose {
			fmt.Print(n.key, " ")
		}
		inOrder(n.right, verbose)
	}
}
