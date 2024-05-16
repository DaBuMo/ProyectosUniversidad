package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	arreglo "tp0/ejercicios"
)

func archivosAVector(rutas []string) []([]int) {

	vect := []([]int){[]int{}, []int{}}

	for i := 0; i < len(rutas); i++ {

		archivo, err := os.Open(rutas[i])
		if err != nil {
			panic(err)
		}

		defer archivo.Close()
		s := bufio.NewScanner(archivo)

		for s.Scan() {
			digit, _ := strconv.Atoi(s.Text())
			vect[i] = append(vect[i], digit)
		}
	}

	return vect
}

func imprimirMayor(vect []([]int), mayor int) {
	elemento := 0

	if mayor == -1 {
		elemento = 1
	}

	arreglo.Seleccion(vect[elemento])

	for _, val := range vect[elemento] {
		fmt.Println(val)
	}
}

func main() {

	rutas := []string{"archivo1.in", "archivo2.in"}
	vect := archivosAVector(rutas)
	mayor := arreglo.Comparar(vect[0], vect[1])
	imprimirMayor(vect, mayor)

}
