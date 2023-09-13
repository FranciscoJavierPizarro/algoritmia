package main

import (
	"bufio"
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

func countingSort(ints IntVector, size int, place int) {
	// max := 10
	// count := ints[0]
	// output := ints[0]

	// // Calculate count of elements
	// for  i := 0; i < size; i++ {
	// 	count[(ints[i] / place) % 10]++;
	// }

	// // Calculate cumulative count
	// for  i := 0; i < max; i++ {
	// 	count[i] += count[i - 1];
	// }

	// // Place the elements in sorted order
	// for i := size - 1; i >= 0; i-- {
	// 	output[count[(ints[i] / place) % 10] - 1] = ints[i];
	// 	count[(ints[i] / place) % 10]--;
	// }

	// for i := 0; i < size; i++{
	// 	ints[i] = output[i];
	// }

}

func getMax(ints IntVector, n int) int {
	max := ints[0]
	for i := 1; i < n; i++ {
		if ints[i] > max {
			max = ints[i]
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

// //////////////////////////////////////////////////////////////////////////////
//
//	HEAP														  //
//
// https://pkg.go.dev/container/heap											  //
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
