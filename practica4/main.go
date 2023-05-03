package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Apertura del fichero de datos
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	enunciados := make([]string, 0)
	pedidosEnunciados := make([][]string, 0)
	//Lectura de la frase del fichero
	scanner := bufio.NewScanner(file)
	aux := 0
	scanner.Scan()
	linea := scanner.Text()
	for linea != "0 0 0" {
		// fmt.Println(linea)
		enunciados = append(enunciados, linea)
		datos := strings.Split(enunciados[aux], " ")
		p, _ := strconv.Atoi(datos[2])
		// fmt.Printf("p:%d\n", p)
		pedidosEnunciados = append(pedidosEnunciados, make([]string, p))
		for i := 0; i < p; i++ {
			scanner.Scan()
			pedidosEnunciados[aux][i] = scanner.Text()
			// fmt.Println("\t" + pedidosEnunciados[aux][i])
		}
		scanner.Scan()
		linea = scanner.Text()
		aux += 1
	}
	fmt.Println(pedidosEnunciados, enunciados)
}
