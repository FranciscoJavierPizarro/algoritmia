package main

import "fmt"

//(👉ﾟヮﾟ)👉

func buscarCadena(dic map[string]int, cadenaBuscada string, n int, encontrados []string, pseudomatrix [][]string, alreadyCalculated int) bool {
	if cadenaBuscada == "" {fmt.Println(encontrados); return true}
	if n > len(cadenaBuscada) {return false}
	if n == 1 && len(pseudomatrix[alreadyCalculated]) != 0 {
		for _, element := range pseudomatrix[alreadyCalculated] {
			buscarCadena(dic, cadenaBuscada[len(element)-1:] , 1, append(encontrados,element),pseudomatrix,alreadyCalculated+(len(element)-1))
		}
	}
	_, exists := dic[cadenaBuscada[:n]]
	encontradaGlobal := 0
	encontradaLocal := 0
	encontradosNuevos := encontrados
	if exists {
		encontradosNuevos = append(encontrados,cadenaBuscada[:n])
	}
	// esta asginación a entero en cuenta de hacerlo con booleanos directamente 
	// es porque en golang el || esta cortocircuitado y el | que usa enteros no.
	if (exists && buscarCadena(dic, cadenaBuscada[n:], 1, encontradosNuevos, pseudomatrix, alreadyCalculated + 1)) {
		encontradaLocal = 1
		pseudomatrix[alreadyCalculated + 1 - n] = append(pseudomatrix[alreadyCalculated + 1 - n], cadenaBuscada[:n]) }
	if (buscarCadena(dic, cadenaBuscada, n+1, encontrados, pseudomatrix, alreadyCalculated + 1)) {encontradaGlobal = 1}
	return  ((encontradaLocal) | (encontradaGlobal)) == 1
}

func main() {
	/***************************************************************************
						Definición del diccionario
	/**************************************************************************/
	dic := make(map[string]int)
	dic["me"] = 0
	dic["gusta"] = 1
	dic["megusta"] = 2
	dic["sol"] = 3
	dic["dar"] = 4
	dic["soldar"] = 5
	dic["la"] = 6
	dic["patata"] = 7
	dic["pata"] = 9
	dic["ta"] = 10
	dic["pa"] = 11
	dic[" "]= 8
	/******************************(。﹏。*)************************************/
	cadenaObjetivo := "megustasoldarlapatata"
	
	pseudomatrix := make([][]string, len(cadenaObjetivo) + 1)
	for i := range pseudomatrix {
		pseudomatrix[i] = make([]string, 0)
	}
	if buscarCadena(dic, cadenaObjetivo, 1, make([]string, 0), pseudomatrix, 0) {
		fmt.Println("Sí.")
	} else {
		fmt.Println("No.")
	}
}