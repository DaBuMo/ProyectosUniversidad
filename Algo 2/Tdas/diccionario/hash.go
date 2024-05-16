package diccionario

import (
	"fmt"
)

const (
	CARGA_POS      = 0
	CARGA_NEV      = 1
	INVALIDO       = -1
	VACIO          = 0
	BORRADO        = 1
	OCUPADO        = 2
	TAMAÑO_I       = 20
	COND_REDIM_POS = 0.7
	COND_REDIM_NEG = 0.2
	CONS_REDIM     = 2
)

type celdaHash[K comparable, V any] struct {
	clave     K
	dato      V
	condicion int
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	tamaño   int
	borrados int
	cant     int
}

type iterador[K comparable, V any] struct {
	hash   hashCerrado[K, V]
	pos    int
	actual celdaHash[K, V]
}

func crearTablaDeHash[K comparable, V any](tamaño int) *[]celdaHash[K, V] {
	tabla := make([]celdaHash[K, V], tamaño)
	return &tabla
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla_dic := *crearTablaDeHash[K, V](TAMAÑO_I)
	return &hashCerrado[K, V]{tabla: tabla_dic, tamaño: TAMAÑO_I}
}

func (d hashCerrado[K, V]) Cantidad() int {
	return d.cant
}

func (d *hashCerrado[K, V]) Guardar(clave K, valor V) {
	if d.calcularFactorCarga(CARGA_POS) >= COND_REDIM_POS {
		d.redimensionar(d.tamaño * CONS_REDIM)
	}

	indice := hash[K, V](clave, d.tamaño)
	d._guardar(indice, clave, valor)

}

func (d *hashCerrado[K, V]) _guardar(indice int, clave K, valor V) {
	guardado := false

	for !guardado {
		if d.tabla[indice].condicion == VACIO || (d.tabla[indice].condicion == OCUPADO && d.tabla[indice].clave == clave) {

			if d.tabla[indice].condicion == VACIO {
				d.tabla[indice].condicion = OCUPADO
				d.tabla[indice].clave = clave
				d.cant++
			}

			d.tabla[indice].dato = valor
			guardado = true
		}

		indice++
		if indice == d.tamaño {
			indice = 0
		}
	}
}

func (d *hashCerrado[K, V]) Borrar(clave K) V {
	if d.calcularFactorCarga(CARGA_NEV) < COND_REDIM_NEG && d.tamaño > TAMAÑO_I {
		d.redimensionar(d.tamaño / CONS_REDIM)
	}

	indice := d.encontrarClave(clave)

	if indice == INVALIDO {
		panic("La clave no pertenece al diccionario")
	}

	dato := d.tabla[indice].dato
	d.tabla[indice].condicion = BORRADO
	d.cant--
	d.borrados++
	return dato
}

func (d hashCerrado[K, V]) Obtener(clave K) V {
	indice := d.encontrarClave(clave)
	if indice == INVALIDO {
		panic("La clave no pertenece al diccionario")
	}

	return d.tabla[indice].dato
}

func (d hashCerrado[K, V]) Pertenece(clave K) bool {
	if d.cant == 0 {
		return false
	}

	return d.encontrarClave(clave) != INVALIDO

}

func (d hashCerrado[K, V]) encontrarClave(clave K) int {
	indice := hash[K, V](clave, d.tamaño)

	for d.tabla[indice].condicion != VACIO {
		if d.tabla[indice].condicion == OCUPADO && d.tabla[indice].clave == clave {
			return indice
		}

		indice++
		if indice == d.tamaño {
			indice = 0
		}
	}

	return INVALIDO
}

func (d hashCerrado[K, V]) calcularFactorCarga(v int) float64 {
	if v == CARGA_POS {
		return float64(d.cant+d.borrados) / float64(d.tamaño)
	}
	return float64(d.cant) / float64(d.tamaño)

}

func (d *hashCerrado[K, V]) redimensionar(tam int) {
	nuevoDic := hashCerrado[K, V]{tabla: *crearTablaDeHash[K, V](tam), tamaño: tam}
	for i := 0; i < d.tamaño; i++ {
		if d.tabla[i].condicion == OCUPADO {
			indice := hash[K, V](d.tabla[i].clave, tam)
			nuevoDic._guardar(indice, d.tabla[i].clave, d.tabla[i].dato)
		}
	}
	*d = nuevoDic
}

func (d *hashCerrado[K, V]) Iterar(visitar func(clave K, valor V) bool) {
	for i := 0; i < d.tamaño; i++ {
		if d.tabla[i].condicion == OCUPADO {
			if !visitar(d.tabla[i].clave, d.tabla[i].dato) {
				return
			}
		}
	}
}

func (d hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	nuevoIterador := new(iterador[K, V])
	nuevoIterador.hash = d
	i := 0
	for d.tabla[i].condicion != OCUPADO && d.Cantidad() != 0 {
		i++
	}
	nuevoIterador.actual = d.tabla[i]
	nuevoIterador.pos = i
	return nuevoIterador
}

func (iter iterador[K, V]) HaySiguiente() bool {
	return !(iter.pos >= iter.hash.tamaño || iter.hash.cant == 0)
}

func (iter iterador[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.clave, iter.actual.dato
}

func (iter *iterador[K, V]) Siguiente() {
	if !iter.HaySiguiente() || iter.hash.Cantidad() == 0 {
		panic("El iterador termino de iterar")
	}
	iter.pos++
	for iter.pos < iter.hash.tamaño && iter.hash.tabla[iter.pos].condicion != OCUPADO {
		iter.pos++
	}
	if iter.pos < iter.hash.tamaño {
		iter.actual = iter.hash.tabla[iter.pos]
	}
}

// FUNC HASH FNV
func hash[K comparable, V any](clave K, tam int) int {
	var hash uint32 = 2166136261
	_clave := convertirABytes[K](clave)
	for _, x := range _clave {
		hash ^= uint32(x)
		hash *= 16777619
	}
	return int(hash) % tam
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
