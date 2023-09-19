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

// Helper function to get the name of a function.
func FunctionName(f IntVectorFunc) string {
	// This is a simple way to get the function name as a string,
	// but it won't work if the function is an anonymous function.
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func ReadVectorsFromFile(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	var vectors [][]int
	scanner := bufio.NewScanner(file)
	const maxBufferSize = 1024 * 1024 * 1024 * 2// 1 GB buffer size
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

func filterGreaterThan(value int, vec []int) []int {
	var result []int
	for _, element := range vec {
		if element > value {
			result = append(result, element)
		}
	}
	return result
}

func filterLowerOrEqualThan(value int, vec []int) []int {
	var result []int
	for _, element := range vec {
		if element <= value {
			result = append(result, element)
		}
	}
	return result
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

func getMax(ints IntVector) int {
	max := ints[0]
	for _, num := range ints {
		if num > max {
			max = num
		}
	}
	return max
}

func concatMultipleSlices(slices []IntVector) IntVector {
	var totalLen int

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make(IntVector, totalLen)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}

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

func isSorted(A IntVector) bool {
	N := len(A)
	for i := 1; i < N; i++ {
		if A[i-1] < A[i] {
			return false
		}
	}
	return true
}

func shuffle(ints IntVector) IntVector {
	n := len(ints)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		ints[i], ints[j] = ints[j], ints[i]
	}
	return ints
}

// //////////////////////////////////////////////////////////////////////////////
//
//	HEAP														  			   //
//
// https://pkg.go.dev/container/heap										   //
// //////////////////////////////////////////////////////////////////////////////
// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
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
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		fmt.Print(n.key, " ")
		printPostOrder(n.right)
	}
}
