package pila

const _INICIO_PILA int = 2

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, _INICIO_PILA)}
}

func (p *pilaDinamica[T]) Apilar(elem T) {
	if p.cantidad == len(p.datos) {
		redimensionar[T](p, len(p.datos)*2)
	}
	p.datos[p.cantidad] = elem
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	if p.cantidad != _INICIO_PILA && p.cantidad*4 <= len(p.datos) {
		redimensionar[T](p, len(p.datos)/2)
	}
	tope := p.datos[p.cantidad-1]
	p.cantidad--
	return tope
}

func (p pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p pilaDinamica[T]) VerTope() T {
	if !p.EstaVacia() {
		return p.datos[p.cantidad-1]
	}
	panic("La pila esta vacia")
}

func redimensionar[T any](p *pilaDinamica[T], tam int) {
	nuevaPila := make([]T, tam)
	copy(nuevaPila, p.datos)
	p.datos = nuevaPila
}
