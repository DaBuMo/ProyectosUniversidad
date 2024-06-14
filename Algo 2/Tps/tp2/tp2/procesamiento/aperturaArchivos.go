package procesamiento

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	DB "tp2/almacenamiento"
)

const (
	_MENSAJE_ERROR      string = "Error en comando"
	_AGREGAR_ARCHIVO    string = "agregar_archivo"
	_MAS_VISITADOS      string = "ver_mas_visitados"
	_VER_VISITANTES     string = "ver_visitantes"
	_CANT_MIN_ARG       int    = 1
	_COMANDO            int    = 0
	_ARCHIVO_O_CANT_SOL int    = 1
	_DESDE              int    = 1
	_HASTA              int    = 2
	_IP                 int    = 0
	_HORA               int    = 1
	_RECURSO            int    = 3
)

func LeerInstruccion(instruccion string, db *DB.DataBase) (string, error) {
	comandos := strings.Fields(instruccion)

	if len(comandos) > 0 {
		comando, direc_desde, _hasta, err := abrirArchivo(comandos)

		if err != nil {
			return comando, err
		} else {

			switch comando {

			case _AGREGAR_ARCHIVO:
				err := guardarLog(direc_desde, db)

				if err != nil {

					return comando, err
				}
				db.RecuperarPosiblesDos()

			case _MAS_VISITADOS:
				fmt.Fprintf(os.Stdout, "%s\n", "Sitios más visitados:")
				cant, _ := strconv.Atoi(direc_desde)
				db.RecuperarMasVisitados(cant)

			case _VER_VISITANTES:
				fmt.Fprintf(os.Stdout, "%s\n", "Visitantes:")
				db.VerVisitantes(direc_desde, _hasta)
			}
		}
	}
	return "", nil
}

func verificarComando(comandos []string) bool {
	largo := len(comandos)
	comandosValidos := []string{_AGREGAR_ARCHIVO, _MAS_VISITADOS, _VER_VISITANTES}
	if largo <= _CANT_MIN_ARG || !slices.Contains(comandosValidos, comandos[_COMANDO]) {
		return false
	}
	return true
}

func guardarLog(direccion string, base *DB.DataBase) error {
	archivo, err := os.Open(direccion)
	if err != nil {
		return errors.New(_MENSAJE_ERROR)
	}
	defer archivo.Close()

	var ip, hora, recurso string
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := strings.Fields(s.Text())
		ip = linea[_IP]
		hora = linea[_HORA]
		recurso = linea[_RECURSO]
		base.AlmacenarIp(ip, hora, recurso)

	}
	return nil
}

func abrirArchivo(comandos []string) (string, string, string, error) {
	if !verificarComando(comandos) {
		return comandos[_COMANDO], "", "", errors.New(_MENSAJE_ERROR)
	} else {
		if comandos[_COMANDO] == _VER_VISITANTES {
			return comandos[_COMANDO], comandos[_DESDE], comandos[_HASTA], nil
		} else {
			return comandos[_COMANDO], comandos[_ARCHIVO_O_CANT_SOL], "", nil
		}
	}
}
