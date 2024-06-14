package lista

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iteradorListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func crearNodo[T any](dato T, sig *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato, sig}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{lista: lista, actual: lista.primero}
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {

	_nodo := crearNodo(elem, lista.primero)
	if lista.EstaVacia() {
		lista.ultimo = _nodo
	}
	lista.primero = _nodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {
	_nodo := crearNodo(elem, nil)
	if lista.EstaVacia() {
		lista.primero = _nodo
	} else {
		lista.ultimo.prox = _nodo
	}
	lista.ultimo = _nodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := lista.primero.dato
	lista.primero = lista.primero.prox
	if lista.primero == nil {
		lista.ultimo = nil
	}

	lista.largo--
	return dato
}

func (lista listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	if lista.EstaVacia() {
		return
	}

	for act := lista.primero; act != nil; act = act.prox {
		resultado := visitar(act.dato)
		if !resultado {
			return
		}
	}
}

func (iterador *iteradorListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.actual.dato
}

func (iterador *iteradorListaEnlazada[T]) HaySiguiente() bool {
	if iterador.lista.EstaVacia() {
		return false
	}
	return iterador.actual != nil
}

func (iterador *iteradorListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.prox
}

func (iterador *iteradorListaEnlazada[T]) Insertar(elem T) {
	nuevoNodo := crearNodo(elem, nil)

	if iterador.lista.EstaVacia() {
		iterador.lista.primero = nuevoNodo
		iterador.lista.ultimo = nuevoNodo
	} else if iterador.actual == iterador.lista.primero {
		nuevoNodo.prox = iterador.lista.primero
		iterador.lista.primero = nuevoNodo
	} else {
		iterador.anterior.prox = nuevoNodo
		nuevoNodo.prox = iterador.actual
		if nuevoNodo.prox == nil {
			iterador.lista.ultimo = nuevoNodo
		}
	}

	iterador.actual = nuevoNodo
	iterador.lista.largo++
}

func (iterador *iteradorListaEnlazada[T]) Borrar() T {

	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	borrado := iterador.actual.dato
	if iterador.actual == iterador.lista.primero {
		iterador.lista.primero = iterador.lista.primero.prox
		iterador.actual = iterador.lista.primero
		iterador.anterior = iterador.lista.primero
		if iterador.lista.primero == nil {
			iterador.lista.primero = nil
		}
	} else if iterador.actual == iterador.lista.ultimo {
		iterador.anterior.prox = nil
		iterador.lista.ultimo = iterador.anterior
		iterador.actual = nil
	} else {
		iterador.actual = iterador.actual.prox
		iterador.anterior.prox = iterador.actual
	}
	iterador.lista.largo--
	return borrado

}
