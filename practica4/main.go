package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Pedido struct {
	Passenger   int
	StartPos    int
	EndPos      int
	TicketValue int
}

// Pedidos contiene todos los pedidos restantes
// Capacidad disp contiene la capacidad total (capacidad * m) menos la capacidad real usada(suma del vector)
func estimatedMaxIncome(pedidos []Pedido, capacidadDisp int) int {
	aux := capacidadDisp
	for _, pedido := range pedidos {
		if (aux - pedido.TicketValue) > 0 {
			aux -= pedido.TicketValue
		}
	}

	return capacidadDisp - aux
}

// Máximo global para realizar la poda
var absolutMax int

func recursiveSearchDFS(pedidos []Pedido, capacidad []int, capacidadDisp int, alreadyTaken []Pedido, maxIncome int, n int, m int) (int, []Pedido) {

	if maxIncome > absolutMax {
		absolutMax = maxIncome
	}

	if len(pedidos) == 0 { //Hemos alcanzado una hoja del árbol
		return maxIncome, alreadyTaken
	}

	bestIncome := 0
	bestTaken := make([]Pedido, 0)

	//PODA
	estimatedMax := estimatedMaxIncome(pedidos, capacidadDisp) + maxIncome
	if estimatedMax < absolutMax {
		// fmt.Println("Se ha podado con pedidos:", pedidos, alreadyTaken, " los maximos total y estimados eran:", absolutMax, estimatedMax)
		return maxIncome, alreadyTaken
	}

	newCapacidad := make([]int, len(capacidad))
	copy(newCapacidad, capacidad)
	admisible := true

	for i := pedidos[0].StartPos; i < pedidos[0].EndPos; i++ {
		if (newCapacidad[i] + pedidos[0].Passenger) <= n {
			newCapacidad[i] += pedidos[0].Passenger
		} else {
			admisible = false
		}
	}

	if admisible {
		// la solución es válida añadiendo el nuevo pedido
		// fmt.Println("sol admisible", alreadyTaken, pedidos[0])
		newCapacidadDisp := capacidadDisp - pedidos[0].TicketValue
		newMaxIncome := maxIncome
		if maxIncome < (n*m - newCapacidadDisp) {
			newMaxIncome = n*m - newCapacidadDisp
		}
		bestIncome, bestTaken = recursiveSearchDFS(pedidos[1:], newCapacidad, newCapacidadDisp, append(alreadyTaken, pedidos[0]), newMaxIncome, n, m)
	} else {
		// fmt.Println("sol NO admisible", alreadyTaken, pedidos[0], capacidad)
	}

	bestIncome2, bestTaken2 := recursiveSearchDFS(pedidos[1:], capacidad, capacidadDisp, alreadyTaken, maxIncome, n, m)

	if bestIncome > bestIncome2 {
		return bestIncome, bestTaken
	} else {
		return bestIncome2, bestTaken2
	}
}

func solveProblemInstance(enunciado string, pedidosEnunciado []string) (int, float64) {
	// Extraemos los datos del enunciado
	n, _ := strconv.Atoi(strings.Split(enunciado, " ")[0])
	m, _ := strconv.Atoi(strings.Split(enunciado, " ")[1])

	// Vector que almacena el estado de la capacidad en cada punto del recorrido
	capacidad := make([]int, m)
	// Vector con la información de pedidos ya en formato entero
	pedidos := make([]Pedido, len(pedidosEnunciado))
	for i, pedido := range pedidosEnunciado {
		data := strings.Split(pedido, " ")
		startpos, _ := strconv.Atoi(data[0])
		endpos, _ := strconv.Atoi(data[1])
		passenger, _ := strconv.Atoi(data[2])
		ticketvalue := (passenger) * (endpos - startpos)
		pedidos[i] = Pedido{Passenger: passenger, StartPos: startpos, EndPos: endpos, TicketValue: ticketvalue}
	}

	// Ordenamos el vector de pedidos(para agilizar el precalculo de ganancia)
	sort.Slice(pedidos, func(i, j int) bool {
		return pedidos[i].TicketValue > pedidos[j].TicketValue
	})

	// fmt.Println(pedidos)
	// income := estimatedMaxIncome(pedidos, n*m)
	// fmt.Println("max initial income following the estimation:", income)

	start := time.Now()
	bestSol, _ := recursiveSearchDFS(pedidos, capacidad, n*m, make([]Pedido, 0), 0, n, m)
	elapsed := float64(time.Since(start).Nanoseconds()) / 1000000.0
	return bestSol, elapsed
}

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

	f, err := os.Create("outputRamif.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}

	defer f.Close()
	w := bufio.NewWriter(f)

	for i := 0; i < aux; i++ {
		absolutMax = 0
		bestSol, elapsed := solveProblemInstance(enunciados[i], pedidosEnunciados[i])
		w.WriteString(strconv.Itoa(bestSol) + ".0 " + strconv.FormatFloat(elapsed, 'f', -1, 64) + "\n")
	}

	w.Flush()
}
