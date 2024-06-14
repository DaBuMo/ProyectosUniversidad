package almacenamiento

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAHEAP "tdas/cola_prioridad"
	TDAHASH "tdas/diccionario"
	"time"
)

const (
	_VALOR_INICIAL_RECURSO = "1"
	_POSIBLE_DOS           = -1
	_RECURSO               = 0
	_APARICIONES           = 1
	_APARICIONES_DOS       = 5
	_APARICION_NUEVA       = 1
	_RESET_DOS             = 0
	_DOS_GUARDADO          = 2
)

type DataBase struct {
	ips          TDAHASH.Diccionario[string, TDAHASH.Diccionario[string, int]]
	visitantes   TDAHASH.DiccionarioOrdenado[string, int]
	recursos     TDAHASH.Diccionario[string, int]
	bannedIps    TDAHASH.Diccionario[string, int]
	dos          []string
	cantRecursos int
	ultimoAcceso string
}

type recurso struct {
	nombre string
	cant   int
}

// Crea y devuelve una base de datos para almacenar las ips
func CrearDB() *DataBase {
	return &DataBase{ips: TDAHASH.CrearHash[string, TDAHASH.Diccionario[string, int]](), visitantes: TDAHASH.CrearABB[string, int](compararIps), recursos: TDAHASH.CrearHash[string, int](), bannedIps: TDAHASH.CrearHash[string, int]()}
}

// Dada una base de datos, almacena en ella las ips visitadas y los recursos utilzados
func (db *DataBase) AlmacenarIp(ip, hora, recurso string) {
	if !db.bannedIps.Pertenece(ip) {

		if db.ips.Pertenece(ip) {
			horasIp := db.ips.Obtener(ip)
			horario := compararHorarios(db.ultimoAcceso, hora)
			if horario < _POSIBLE_DOS {
				db.ultimoAcceso = hora
			}

			if horasIp.Pertenece(db.ultimoAcceso) {
				act := horasIp.Obtener(db.ultimoAcceso) + 1
				horasIp.Guardar(db.ultimoAcceso, act)
			} else {
				horasIp.Guardar(db.ultimoAcceso, _APARICION_NUEVA)
			}

			if horasIp.Obtener(db.ultimoAcceso) == _APARICIONES_DOS {
				db.bannedIps.Guardar(ip, _DOS_GUARDADO)
				db.dos = append(db.dos, ip)
			}

		} else {
			horasIp := TDAHASH.CrearHash[string, int]()
			horasIp.Guardar(hora, _APARICION_NUEVA)
			db.ips.Guardar(ip, horasIp)
			db.visitantes.Guardar(ip, 0)
			db.ultimoAcceso = hora
		}
	}

	if !db.recursos.Pertenece(recurso) {
		db.recursos.Guardar(recurso, _APARICION_NUEVA)
		db.cantRecursos++
	} else {
		act := db.recursos.Obtener(recurso)
		db.recursos.Guardar(recurso, act+1)
	}
}

// Recupera los recursos utilizados junto a sus apariciones y devuelve un slice de ellos convertidos en structs del tipo recurso
func (db *DataBase) recuperarRecursos() []recurso {
	iter := db.recursos.Iterador()
	arr := make([]recurso, db.cantRecursos)
	cont := 0
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		arr[cont] = recurso{nombre: clave, cant: valor}
		cont++
		iter.Siguiente()
	}
	return arr
}

// Dado un rango de ips
func (db DataBase) VerVisitantes(desde, hasta string) {
	iter := db.visitantes.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		fmt.Fprintf(os.Stdout, "	%s\n", clave)
		iter.Siguiente()
	}
}

func (db *DataBase) RecuperarMasVisitados(cantidad int) {
	arr := db.recuperarRecursos()
	heap := TDAHEAP.CrearHeapArr(arr, cmp)
	for i := 0; i < cantidad; i++ {
		if !heap.EstaVacia() {
			valor := heap.Desencolar()
			fmt.Fprintf(os.Stdout, "	%s - %d\n", valor.nombre, valor.cant)
		}
	}
}

func (db *DataBase) RecuperarPosiblesDos() {

	TDAHEAP.HeapSort(db.dos, compararIps)
	for _, ip := range db.dos {
		db.bannedIps.Borrar(ip)
		fmt.Fprintf(os.Stdout, "%s %s\n", "DoS:", ip)
	}
	db.dos = make([]string, _RESET_DOS)
}

func cmp(a, b recurso) int {
	if a.cant > b.cant {
		return 1
	} else if a.cant < b.cant {
		return -1
	}
	return 0
}

// Devuelve 1 si el primer horario es mas viejo, -1 si el primer horario es mas actual, 0 si ambos horarios son iguales
func compararHorarios(_hora1, _hora2 string) int {
	hora1, _ := time.Parse("2006-01-02T15:04:05-07:00", _hora1)
	hora2, _ := time.Parse("2006-01-02T15:04:05-07:00", _hora2)
	return int(hora1.Sub(hora2).Seconds())
}

func compararIps(ip1, ip2 string) int {
	var num1, num2 int
	_ip1 := strings.Split(ip1, ".")
	_ip2 := strings.Split(ip2, ".")

	for i := 0; i < len(_ip1); i++ {
		num1, _ = strconv.Atoi(_ip1[i])
		num2, _ = strconv.Atoi(_ip2[i])

		if num1 > num2 {
			return 1
		} else if num2 > num1 {
			return -1
		}
	}

	return 0
}
