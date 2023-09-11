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
