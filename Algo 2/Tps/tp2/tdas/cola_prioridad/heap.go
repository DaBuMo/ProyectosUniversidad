package cola_prioridad

const (
	_TAMAÑO_MIN_HEAP = 10
)

type heap[T any] struct {
	arreglo  []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{arreglo: make([]T, _TAMAÑO_MIN_HEAP), cmp: funcion_cmp}
}
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	arreglo_heap := make([]T, len(arreglo))
	copy(arreglo_heap, arreglo)

	heap := heap[T]{arreglo: arreglo_heap, cantidad: len(arreglo_heap), cmp: funcion_cmp}
	heapify(&heap)
	return &heap
}

func (heap heap[T]) Cantidad() int {
	return heap.cantidad
}

func (heap *heap[T]) Encolar(elem T) {
	if heap.cantidad == 0 {
		redimensionar[T](heap, _TAMAÑO_MIN_HEAP*2)
	} else if heap.cantidad == len(heap.arreglo) {
		redimensionar[T](heap, len(heap.arreglo)*2)
	}
	heap.arreglo[heap.cantidad] = elem
	heap.upHeap(heap.cantidad)
	heap.cantidad++
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := heap.VerMax()
	swap(&heap.arreglo[0], &heap.arreglo[heap.cantidad-1])
	heap.cantidad--
	heap.downHeap(0)
	if heap.cantidad > _TAMAÑO_MIN_HEAP && heap.cantidad*4 < len(heap.arreglo) {
		redimensionar[T](heap, len(heap.arreglo)/2)
	}
	return dato
}

func (heap heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.arreglo[0]
}

func (heap heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) upHeap(indice int) {

	if indice == 0 {
		return
	}

	actual := heap.arreglo[indice]
	indicePadre := (indice - 1) / 2
	padre := heap.arreglo[indicePadre]

	if heap.cmp(actual, padre) > 0 {
		swap(&heap.arreglo[indice], &heap.arreglo[indicePadre])
		heap.upHeap(indicePadre)
	}
}

func (heap *heap[T]) downHeap(indice int) {
	var izq, der T
	indice_izq := 2*indice + 1

	if indice_izq >= heap.cantidad {
		return
	}

	indice_der := 2*indice + 2
	mayor := heap.arreglo[indice]
	indice_mayor := indice
	der_valido := false

	izq = heap.arreglo[indice_izq]

	if indice_der < heap.cantidad {
		der = heap.arreglo[indice_der]
		der_valido = true
	}

	if heap.cmp(izq, mayor) > 0 {
		mayor = izq
		indice_mayor = indice_izq
	}
	if heap.cmp(der, mayor) > 0 && der_valido {
		indice_mayor = indice_der
	}

	if indice_mayor != indice {
		swap(&heap.arreglo[indice], &heap.arreglo[indice_mayor])
		heap.downHeap(indice_mayor)
	}

}

func heapify[T any](heap *heap[T]) {
	for i := (heap.cantidad / 2) - 1; i >= 0; i-- {
		heap.downHeap(i)
	}
}
func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heap := heap[T]{arreglo: elementos, cantidad: len(elementos), cmp: funcion_cmp}
	heapify(&heap)
	for i := heap.cantidad - 1; i > -1; i-- {
		swap(&heap.arreglo[0], &heap.arreglo[i])
		heap.cantidad--
		heap.downHeap(0)
	}
}

func redimensionar[T any](heap *heap[T], tam int) {
	nuevaPila := make([]T, tam)
	copy(nuevaPila, heap.arreglo)
	heap.arreglo = nuevaPila
}

func swap[T any](a *T, b *T) {
	*a, *b = *b, *a
}
