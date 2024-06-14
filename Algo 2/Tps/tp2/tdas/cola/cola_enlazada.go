package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	prim *nodoCola[T]
	ult  *nodoCola[T]
}

func nodoCrear[T any](dato T, sig *nodoCola[T]) *nodoCola[T] {
	return &nodoCola[T]{dato, sig}
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{prim: nil, ult: nil}
}

func (c colaEnlazada[T]) EstaVacia() bool {
	return c.prim == nil

}

func (c colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.prim.dato
}

func (c *colaEnlazada[T]) Encolar(elem T) {
	nodo := nodoCrear[T](elem, nil)

	if c.EstaVacia() {
		c.prim = nodo
		c.ult = nodo
	} else {
		c.ult.prox = nodo
		c.ult = nodo
	}
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := c.prim.dato
	c.prim = c.prim.prox
	if c.prim == nil {
		c.ult = nil
	}
	return dato
}
