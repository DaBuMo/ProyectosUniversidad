package diccionario

import TDAPila "tdas/pila"

type nodoArbol[K comparable, V any] struct {
	clave K
	dato  V
	h_izq *nodoArbol[K, V]
	h_der *nodoArbol[K, V]
}

type abb[K comparable, V any] struct {
	raiz *nodoArbol[K, V]
	cant int
	cmp  func(K, K) int
}

type iteradorExterno[K comparable, V any] struct {
	arbol *abb[K, V]
	pila  TDAPila.Pila[*nodoArbol[K, V]]
	desde *K
	hasta *K
}

func crearNodo[K comparable, V any](clave K, dato V, h_izq *nodoArbol[K, V], h_der *nodoArbol[K, V]) *nodoArbol[K, V] {
	return &nodoArbol[K, V]{clave: clave, dato: dato, h_izq: h_izq, h_der: h_der}
}

func CrearABB[K comparable, V any](cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cant: 0, cmp: cmp}
}

func (arbol *abb[K, V]) Guardar(clave K, dato V) {
	nodo := crearNodo(clave, dato, nil, nil)
	if arbol.raiz == nil {
		arbol.raiz = nodo
		arbol.cant++
	} else {
		act, padre := arbol.buscarClave(clave, arbol.raiz, nil)
		if act == nil {
			mayor := arbol.cmp(clave, padre.clave)
			if mayor < 0 {
				padre.h_izq = nodo
			} else {
				padre.h_der = nodo
			}
			arbol.cant++
		} else {
			act.dato = dato
		}
	}
}
func (arbol abb[K, V]) Pertenece(clave K) bool {
	act, _ := arbol.buscarClave(clave, arbol.raiz, nil)
	return act != nil
}

func (arbol abb[K, V]) Obtener(clave K) V {
	act, _ := arbol.buscarClave(clave, arbol.raiz, nil)
	if act != nil {
		return act.dato
	}
	panic("La clave no pertenece al diccionario")

}
func (arbol *abb[K, V]) Borrar(clave K) V {
	actual, padre := arbol.buscarClave(clave, arbol.raiz, nil)

	if actual == nil {
		panic("La clave no pertenece al diccionario")
	}
	valor := arbol._borrar(actual, padre)
	arbol.cant--
	return valor

}

func (arbol *abb[K, V]) _borrar(actual *nodoArbol[K, V], padre *nodoArbol[K, V]) V {
	dato := actual.dato
	if actual.h_der == nil || actual.h_izq == nil {
		sucesor := actual.h_der
		if sucesor == nil {
			sucesor = actual.h_izq
		}
		if padre != nil {
			if arbol.cmp(actual.clave, padre.clave) >= 0 {
				padre.h_der = sucesor
			} else {
				padre.h_izq = sucesor
			}
		} else {
			arbol.raiz = sucesor
		}
	} else {
		sucesor, padreSucesor := arbol.mayorHijorMenor(actual)
		datoSucesor := sucesor.dato
		claveSucesor := sucesor.clave
		arbol._borrar(sucesor, padreSucesor)
		actual.dato = datoSucesor
		actual.clave = claveSucesor
	}
	return dato
}

func (arbol abb[K, V]) mayorHijorMenor(nodo *nodoArbol[K, V]) (*nodoArbol[K, V], *nodoArbol[K, V]) {
	var padre *nodoArbol[K, V]
	actual := nodo
	siguiente := nodo.h_izq
	for siguiente != nil {
		padre = actual
		actual = siguiente
		siguiente = siguiente.h_der
	}
	return actual, padre
}
func (arbol abb[K, V]) Cantidad() int {
	return arbol.cant
}

func (arbol abb[K, V]) buscarClave(clave K, act *nodoArbol[K, V], padre *nodoArbol[K, V]) (*nodoArbol[K, V], *nodoArbol[K, V]) {
	if act == nil {
		return act, padre
	} else {
		mayor := arbol.cmp(clave, act.clave)

		if mayor < 0 {
			return arbol.buscarClave(clave, act.h_izq, act)
		} else if mayor > 0 {
			return arbol.buscarClave(clave, act.h_der, act)
		}

		return act, padre
	}
}
func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return arbol.IteradorRango(nil, nil)
}

func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoArbol[K, V]]()
	iterador := iteradorExterno[K, V]{arbol: arbol, pila: pila, desde: desde, hasta: hasta}
	apilarNodosRango[K, V](arbol, arbol.raiz, &iterador)
	return &iterador
}
func (iter iteradorExterno[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iter.pila.VerTope()
	return actual.clave, actual.dato
}
func (iter iteradorExterno[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}
func (iter *iteradorExterno[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iter.pila.Desapilar()
	if actual.h_der != nil {
		apilarNodosRango(iter.arbol, actual.h_der, iter)
	}
}
func (arbol *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	arbol.IterarRango(nil, nil, visitar)
}

func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if !_iterarRango(visitar, arbol.raiz, desde, hasta, arbol) {
		return
	}
}

func _iterarRango[K comparable, V any](visitar func(clave K, valor V) bool, actual *nodoArbol[K, V], desde *K, hasta *K, arbol *abb[K, V]) bool {
	if actual == nil {
		return true
	}
	if hasta != nil && arbol.cmp(actual.clave, *hasta) > 0 {
		return _iterarRango(visitar, actual.h_izq, desde, hasta, arbol)
	} else if desde != nil && arbol.cmp(actual.clave, *desde) < 0 {
		return _iterarRango(visitar, actual.h_der, desde, hasta, arbol)
	}

	if (actual.h_izq != nil && !_iterarRango(visitar, actual.h_izq, desde, hasta, arbol)) || !visitar(actual.clave, actual.dato) {
		return false
	}
	return _iterarRango(visitar, actual.h_der, desde, hasta, arbol)
}

func apilarNodosRango[K comparable, V any](arbol *abb[K, V], actual *nodoArbol[K, V], iterador *iteradorExterno[K, V]) {
	if actual == nil {
		return
	}
	if iterador.desde != nil && arbol.cmp(actual.clave, *iterador.desde) < 0 {
		apilarNodosRango(arbol, actual.h_der, iterador)
		return
	}
	if iterador.hasta != nil && arbol.cmp(actual.clave, *iterador.hasta) > 0 {
		apilarNodosRango(arbol, actual.h_izq, iterador)
		return
	}
	iterador.pila.Apilar(actual)
	apilarNodosRango(arbol, actual.h_izq, iterador)
}
