package main

import (
	"bufio"
	"fmt"
	"os"
	DB "tp2/almacenamiento"
	PROCESAMIENTO "tp2/procesamiento"
)

const (
	MENSAJE_VALIDACION = "OK"
)

func main() {
	almacenamientoIps := DB.CrearDB()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		comando, err := PROCESAMIENTO.LeerInstruccion(s.Text(), almacenamientoIps)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s %s\n", err, comando)
		} else {
			fmt.Fprintf(os.Stdout, "%s\n", MENSAJE_VALIDACION)
		}
	}
}
