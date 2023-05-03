package main

import (
	"sort"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pedido struct {
	Passenger int
    StartPos int
    EndPos   int
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

func recursiveSearchDFS(pedidos []Pedido, capacidad []int, capacidadDisp int, alreadyTaken []Pedido, maxIncome int, n int, m int) (int, []Pedido){

	if (len(pedidos) == 0) {
		return maxIncome, alreadyTaken
	}
	bestIncome := 0
	bestTaken := make([]Pedido,0)
	//FALTA LA PODA
	//FALTA SIMULAR LA BAJADA??
	newCapacidad := capacidad
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
		newCapacidadDisp := capacidadDisp - pedidos[0].TicketValue
		newMaxIncome := maxIncome
		if (maxIncome < (n*m - newCapacidadDisp)) {
			newMaxIncome = n*m - newCapacidadDisp
		}
		bestIncome, bestTaken = recursiveSearchDFS(pedidos[1:],newCapacidad,newCapacidadDisp,append(alreadyTaken, pedidos[0]), newMaxIncome,n,m)
	}
	bestIncome2, bestTaken2 := recursiveSearchDFS(pedidos[1:],capacidad,capacidadDisp,alreadyTaken,maxIncome,n,m)
	if bestIncome > bestIncome2 {return bestIncome, bestTaken;} else {return bestIncome2, bestTaken2;}
}

func solveProblemInstance(enunciado string, pedidosEnunciado []string) {
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
	fmt.Println(pedidos)

	income := estimatedMaxIncome(pedidos, n*m)
	fmt.Println(income)
	fmt.Println(recursiveSearchDFS(pedidos,capacidad,n*m,make([]Pedido,0),0,n,m))
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
	
	for i:=0; i < aux; i++ {
		solveProblemInstance(enunciados[i],pedidosEnunciados[i])
	}
}
