package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

type Road struct {
	u, v          int
	tuv, puv, pvu float64
}

//////////////////////////////////////////
//			FUNCIONES AUXIALARES		//
//////////////////////////////////////////

func atoi(s string) int {
	var res int
	fmt.Sscanf(s, "%d", &res)
	return res
}

func atof(s string) float64 {
	var res float64
	fmt.Sscanf(s, "%f", &res)
	return res
}

//////////////////////////////////////////
//			FUNCIONES PRINCIPALES	   //
/////////////////////////////////////////

func buscarCaminos(X int, roads []Road) []Road {
	var result []Road

	for _, road := range roads {
		if road.u == X || road.v == X {
			result = append(result, road)
		}
	}
	return result
}

func selectRandomPath(X int, roads []Road) (int, float64) {
	var auxRoads []Road
	var possiblities []float64
	var auxResult Road
	var result int
	for _, road := range roads {
		if road.u == X {
			auxRoads = append(auxRoads, road)
			possiblities = append(possiblities, road.puv)
		} else if road.v == X {
			auxRoads = append(auxRoads, road)
			possiblities = append(possiblities, road.pvu)
		}
	}
	N := len(auxRoads)
	randomNumber := rand.Float64()
	for i := 0; i < N; i++ {
		randomNumber -= possiblities[i]
		if randomNumber <= 0 {
			auxResult = auxRoads[i]
			break
		}
	}

	if auxResult.u == X {
		result = auxResult.v
	} else {
		result = auxResult.u
	}
	return result, auxResult.tuv
}

func simular(X, C int, roads []Road) float64 {
	var auxt float64
	currentInter := X
	t := 0.0
	var posiblesCarreteras []Road
	for currentInter != C {
		posiblesCarreteras = buscarCaminos(currentInter, roads)
		currentInter, auxt = selectRandomPath(currentInter, posiblesCarreteras)
		t += auxt
	}
	return t
}

func bootstrap(N int, vec []float64) float64 {
	res := 0.0
	for i := 0; i < N; i++ {
		random := rand.Intn(N)
		res += vec[random]
	}
	res = res / float64(N)
	return res
}

/////////////////////////////////////////
//				MAIN				   //
/////////////////////////////////////////

func main() {
	rand.Seed(time.Now().UnixNano())
	Nsimulaciones := 500
	Nboostraps := 50
	file, _ := os.Open("example.txt")
	defer file.Close()

	var N, M, C, A, B int
	var roads []Road

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		input := scanner.Text()
		values := strings.Fields(input)
		N = atoi(values[0])
		M = atoi(values[1])
		C = atoi(values[2])
		A = atoi(values[3])
		B = atoi(values[4])
	}

	for i := 0; i < M; i++ {
		if scanner.Scan() {
			input := scanner.Text()
			values := strings.Fields(input)
			u := atoi(values[0])
			v := atoi(values[1])
			tuv := atof(values[2])
			puv := atof(values[3])
			pvu := atof(values[4])
			roads = append(roads, Road{u, v, tuv, puv, pvu})
		}
	}

	fmt.Printf("N: %d, M: %d, C: %d, A: %d, B: %d\n", N, M, C, A, B)
	fmt.Println("Roads:")
	for _, road := range roads {
		fmt.Printf("u: %d, v: %d, tuv: %f, puv: %f, pvu: %f\n", road.u, road.v, road.tuv, road.puv, road.pvu)
	}

	var tAlist, tBlist []float64
	OA, OB := 0.0, 0.0
	for i := 0; i < Nsimulaciones; i++ {

		//generar mapa aleatorio aqui :D
		tA := simular(A, C, roads)
		tB := simular(B, C, roads)
		tBlist = append(tBlist, tB)
		tAlist = append(tAlist, tA)
		OA += tA
		OB += tB
	}
	OA = OA / float64(Nsimulaciones)
	OB = OB / float64(Nsimulaciones)

	var boostrapsA, boostrapsB []float64
	for i := 0; i < Nboostraps; i++ {
		boostrapsA = append(boostrapsA, bootstrap(Nsimulaciones, tAlist))
		boostrapsB = append(boostrapsB, bootstrap(Nsimulaciones, tBlist))
	}
	sort.Float64s(boostrapsA)

	LA := 2*OA - (boostrapsA[Nboostraps/10*9]+boostrapsA[(Nboostraps/10*9)+1])/2
	RA := 2*OA - (boostrapsA[Nboostraps/10]+boostrapsA[(Nboostraps/10)+1])/2
	fmt.Println(LA, RA)

	LB := 2*OB - (boostrapsB[Nboostraps/10*9]+boostrapsB[(Nboostraps/10*9)+1])/2
	RB := 2*OB - (boostrapsB[Nboostraps/10]+boostrapsB[(Nboostraps/10)+1])/2
	fmt.Println(LB, RB)
}
