package main

import (
	"bufio"
	"fmt"
	"os"
	"tp1/operaciones"
)

const ERROR string = "ERROR"

func main() {

	calc := operaciones.CrearCalculadora()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}

		resultado, err := calc.IniciarCalculo(s.Text())

		if err != nil {
			fmt.Fprintf(os.Stdout, "%s\n", ERROR)
		} else {
			fmt.Fprintf(os.Stdout, "%d\n", resultado)
		}
	}
}
